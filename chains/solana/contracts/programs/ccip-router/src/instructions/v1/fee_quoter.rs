use std::ops::AddAssign;

use anchor_lang::prelude::*;
use anchor_spl::{token::spl_token::native_mint, token_interface};
use ethnum::U256;
use solana_program::{program::invoke_signed, system_instruction};

use crate::{
    BillingTokenConfig, CcipRouterError, DestChain, PerChainPerTokenConfig, Solana2AnyMessage,
    SolanaTokenAmount, UnpackedDoubleU224, FEE_BILLING_SIGNER_SEEDS,
};

use super::messages::ramps::validate_solana2any;
use super::pools::CCIP_LOCK_OR_BURN_V1_RET_BYTES;
use super::utils::{Exponential, Usd18Decimals};

/// Any2EVMRampMessage struct has 10 fields, including 3 variable unnested arrays (data, receiver and tokenAmounts).
/// Each variable array takes 1 more slot to store its length.
/// When abi encoded, excluding array contents,
/// Any2EVMMessage takes up a fixed number of 13 slots, 32 bytes each.
/// For structs that contain arrays, 1 more slot is added to the front, reaching a total of 14.
/// The fixed bytes does not cover struct data (this is represented by ANY_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN)
pub const ANY_2_EVM_MESSAGE_FIXED_BYTES: U256 = U256::new(32 * 14);

/// Each token transfer adds 1 RampTokenAmount
/// RampTokenAmount has 5 fields, 2 of which are bytes type, 1 Address, 1 uint256 and 1 uint32.
/// Each bytes type takes 1 slot for length, 1 slot for data and 1 slot for the offset.
/// address
/// uint256 amount takes 1 slot.
/// uint32 destGasAmount takes 1 slot.
pub const ANY_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN: U256 = U256::new(32 * ((2 * 3) + 3));

pub fn fee_for_msg(
    _dest_chain_selector: u64,
    message: &Solana2AnyMessage,
    dest_chain: &DestChain,
    fee_token_config: &BillingTokenConfig,
    additional_token_configs: &[Option<BillingTokenConfig>],
    additional_token_configs_for_dest_chain: &[PerChainPerTokenConfig],
) -> Result<SolanaTokenAmount> {
    let fee_token = if message.fee_token == Pubkey::default() {
        native_mint::ID // Wrapped SOL
    } else {
        message.fee_token
    };
    require!(
        additional_token_configs.len() == message.token_amounts.len(),
        CcipRouterError::InvalidInputsMissingTokenConfig
    );
    require!(
        additional_token_configs_for_dest_chain.len() == message.token_amounts.len(),
        CcipRouterError::InvalidInputsMissingTokenConfig
    );
    validate_solana2any(message, dest_chain, fee_token_config)?;

    let fee_token_price = get_validated_token_price(fee_token_config)?;
    let PackedPrice {
        execution_gas_price,
        data_availability_gas_price,
    } = get_validated_gas_price(dest_chain)?;

    let network_fee = network_fee(
        message,
        dest_chain,
        additional_token_configs,
        additional_token_configs_for_dest_chain,
    )?;

    let gas_limit = U256::new(
        message
            .extra_args
            .gas_limit
            .unwrap_or_else(|| dest_chain.config.default_tx_gas_limit.into()),
    );

    let execution_gas = gas_limit
        + U256::new(dest_chain.config.dest_gas_overhead as u128)
        + U256::new(message.data.len() as u128)
            * U256::new(dest_chain.config.dest_gas_per_payload_byte as u128)
        + network_fee.transfer_gas;

    let execution_cost = execution_gas_price
        * execution_gas
        * U256::new(dest_chain.config.gas_multiplier_wei_per_eth as u128);

    let data_availability_cost = data_availability_cost(
        data_availability_gas_price,
        message,
        network_fee.transfer_bytes_overhead,
        dest_chain,
    );

    let premium_multiplier = U256::new(fee_token_config.premium_multiplier_wei_per_eth.into());
    let fee_token_value =
        (network_fee.premium * premium_multiplier) + execution_cost + data_availability_cost;
    SolanaTokenAmount::amount(fee_token, fee_token_value, fee_token_price)
}

fn data_availability_cost(
    data_availability_gas_price: Usd18Decimals,
    message: &Solana2AnyMessage,
    token_transfer_bytes_overhead: U256,
    dest_chain: &DestChain,
) -> Usd18Decimals {
    // Sums up byte lengths of fixed message fields and dynamic message fields.
    // Fixed message fields do account for the offset and length slot of the dynamic fields.
    let data_availability_length_bytes = ANY_2_EVM_MESSAGE_FIXED_BYTES
        + U256::new(message.data.len() as u128)
        + (U256::new(message.token_amounts.len() as u128)
            * ANY_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN)
        + token_transfer_bytes_overhead;

    // dest_data_availability_overhead_gas is a separate config value for flexibility to be updated
    // independently of message cost. Its value is determined by CCIP lane implementation, e.g.
    // the overhead data posted for OCR.
    let data_availability_gas = data_availability_length_bytes
        * U256::new(dest_chain.config.dest_gas_per_data_availability_byte as u128)
        + U256::new(dest_chain.config.dest_data_availability_overhead_gas as u128);

    // data_availability_gas_price is in 18 decimals, dest_data_availability_multiplier_bps is in 4 decimals
    // We pad 14 decimals to bring the result to 36 decimals, in line with token bps and execution fee.
    data_availability_gas_price
        * data_availability_gas
        * U256::new(dest_chain.config.dest_data_availability_multiplier_bps as u128)
        * 1u32.e(14)
}

#[derive(Clone, Default, Debug)]
struct NetworkFee {
    premium: Usd18Decimals,
    transfer_gas: U256,
    transfer_bytes_overhead: U256,
}

impl AddAssign for NetworkFee {
    fn add_assign(&mut self, rhs: Self) {
        self.premium += rhs.premium;
        self.transfer_gas += rhs.transfer_gas;
        self.transfer_bytes_overhead += rhs.transfer_bytes_overhead;
    }
}

fn network_fee(
    message: &Solana2AnyMessage,
    dest_chain: &DestChain,
    token_configs: &[Option<BillingTokenConfig>],
    token_configs_for_dest_chain: &[PerChainPerTokenConfig],
) -> Result<NetworkFee> {
    if message.token_amounts.is_empty() {
        return Ok(NetworkFee {
            premium: Usd18Decimals::from_usd_cents(dest_chain.config.network_fee_usdcents),
            transfer_gas: U256::ZERO,
            transfer_bytes_overhead: U256::ZERO,
        });
    }

    let mut fee = NetworkFee::default();

    for (i, token_amount) in message.token_amounts.iter().enumerate() {
        let config_for_dest_chain = &token_configs_for_dest_chain[i];
        let token_network_fee = if config_for_dest_chain.billing.is_enabled {
            token_network_fees(&token_configs[i], token_amount, config_for_dest_chain)?
        } else {
            // If the token has no specific overrides configured, we use the global defaults.
            global_network_fees(dest_chain)
        };

        fee += token_network_fee;
    }

    Ok(fee)
}

fn token_network_fees(
    billing_config: &Option<BillingTokenConfig>,
    token_amount: &SolanaTokenAmount,
    config_for_dest_chain: &PerChainPerTokenConfig,
) -> Result<NetworkFee> {
    let bps_fee = match billing_config {
        Some(config) if config_for_dest_chain.billing.deci_bps > 0 => {
            let token_price = get_validated_token_price(config)?;
            // Calculate token transfer value, then apply fee ratio
            // ratio represents multiples of 0.1bps, or 1e-5
            Usd18Decimals(
                (token_amount.value(&token_price).0
                    * U256::new(config_for_dest_chain.billing.deci_bps.into()))
                    / 1u32.e(5),
            )
        }
        _ => Usd18Decimals::ZERO,
    };

    let min_fee = Usd18Decimals::from_usd_cents(config_for_dest_chain.billing.min_fee_usdcents);
    let max_fee = Usd18Decimals::from_usd_cents(config_for_dest_chain.billing.max_fee_usdcents);
    let (premium, token_transfer_gas, token_transfer_bytes_overhead) = (
        bps_fee.clamp(min_fee, max_fee),
        U256::new(config_for_dest_chain.billing.dest_gas_overhead.into()),
        U256::new(config_for_dest_chain.billing.dest_bytes_overhead.into()),
    );
    Ok(NetworkFee {
        premium,
        transfer_gas: token_transfer_gas,
        transfer_bytes_overhead: token_transfer_bytes_overhead,
    })
}

fn global_network_fees(dest_chain: &DestChain) -> NetworkFee {
    let (premium, global_gas, global_overhead) = (
        Usd18Decimals::from_usd_cents(dest_chain.config.default_token_fee_usdcents.into()),
        U256::new(dest_chain.config.default_token_dest_gas_overhead.into()),
        U256::new(CCIP_LOCK_OR_BURN_V1_RET_BYTES.into()),
    );
    NetworkFee {
        premium,
        transfer_gas: global_gas,
        transfer_bytes_overhead: global_overhead,
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct PackedPrice {
    // L1 gas price (encoded in the lower 112 bits)
    pub execution_gas_price: Usd18Decimals,
    // L2 gas price (encoded in the higher 112 bits)
    pub data_availability_gas_price: Usd18Decimals,
}

impl From<UnpackedDoubleU224> for PackedPrice {
    fn from(value: UnpackedDoubleU224) -> Self {
        Self {
            execution_gas_price: Usd18Decimals(value.low.into()),
            data_availability_gas_price: Usd18Decimals(value.high.into()),
        }
    }
}

impl TryFrom<PackedPrice> for UnpackedDoubleU224 {
    type Error = Error;

    fn try_from(value: PackedPrice) -> Result<Self> {
        Ok(Self {
            high: value
                .data_availability_gas_price
                .0
                .try_into()
                .map_err(|_| CcipRouterError::InvalidTokenPrice)?,
            low: value
                .execution_gas_price
                .0
                .try_into()
                .map_err(|_| CcipRouterError::InvalidTokenPrice)?,
        })
    }
}

fn get_validated_gas_price(dest_chain: &DestChain) -> Result<PackedPrice> {
    let timestamp = dest_chain.state.usd_per_unit_gas.timestamp;
    let price = dest_chain.state.usd_per_unit_gas.unpack().into();
    let threshold = dest_chain.config.gas_price_staleness_threshold as i64;
    let elapsed_time = Clock::get()?.unix_timestamp - timestamp;

    require!(
        threshold == 0 || threshold > elapsed_time,
        CcipRouterError::StaleGasPrice
    );

    Ok(price)
}

fn get_validated_token_price(token_config: &BillingTokenConfig) -> Result<Usd18Decimals> {
    let timestamp = token_config.usd_per_token.timestamp;
    let price = token_config.usd_per_token.as_single();

    // NOTE: There's no validation done with respect to token price staleness since data feeds are not
    // supported in solana. Only the existence of `any` timestamp is checked, to ensure the price
    // was set at least once.
    require!(
        price != 0 && timestamp != 0,
        CcipRouterError::InvalidTokenPrice
    );

    Ok(Usd18Decimals(price))
}

pub fn wrap_native_sol<'info>(
    token_program: &AccountInfo<'info>,
    from: &mut Signer<'info>,
    to: &mut InterfaceAccount<'info, token_interface::TokenAccount>,
    amount: u64,
    signer_bump: u8,
) -> Result<()> {
    require!(
        // guarantee that if caller is a PDA with data it won't get garbage-collected
        from.data_is_empty() || from.get_lamports() > amount,
        CcipRouterError::InsufficientLamports
    );

    invoke_signed(
        &system_instruction::transfer(&from.key(), &to.key(), amount),
        &[from.to_account_info(), to.to_account_info()],
        &[&[FEE_BILLING_SIGNER_SEEDS, &[signer_bump]]],
    )?;

    let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let account = to.to_account_info();
    let sync: anchor_spl::token_2022::SyncNative = anchor_spl::token_2022::SyncNative { account };
    let cpi_ctx = CpiContext::new_with_signer(token_program.to_account_info(), sync, signer_seeds);

    token_interface::sync_native(cpi_ctx)
}

pub fn transfer_fee<'info>(
    fee: SolanaTokenAmount,
    token_program: AccountInfo<'info>,
    transfer: token_interface::TransferChecked<'info>,
    decimals: u8,
    signer_bump: u8,
) -> Result<()> {
    require!(
        fee.token == transfer.mint.key(),
        CcipRouterError::InvalidInputs
    ); // TODO use more specific error

    do_billing_transfer(token_program, transfer, fee.amount, decimals, signer_bump)
}

pub fn do_billing_transfer<'info>(
    token_program: AccountInfo<'info>,
    transfer: token_interface::TransferChecked<'info>,
    amount: u64,
    decimals: u8,
    signer_bump: u8,
) -> Result<()> {
    let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let cpi_ctx = CpiContext::new_with_signer(token_program, transfer, signer_seeds);
    token_interface::transfer_checked(cpi_ctx, amount, decimals)
}

#[cfg(test)]
mod tests {
    use solana_program::{
        entrypoint::SUCCESS,
        program_stubs::{set_syscall_stubs, SyscallStubs},
    };

    use crate::TokenBilling;

    use super::super::messages::ramps::tests::*;
    use super::*;

    struct TestStubs;

    impl SyscallStubs for TestStubs {
        fn sol_get_clock_sysvar(&self, _var_addr: *mut u8) -> u64 {
            // This causes the syscall to return a default-initialized
            // clock when the build target is off-chain. Good enough for tests.
            SUCCESS
        }
    }

    #[test]
    // NOTE: This test is unique in that the return value of `fee_for_msg` has been
    // directly validated against a `getFee` query in the Ethereum mainnet -> Avalanche Fuji
    // lane, using the same configuration. This ensures that at least on a simple execution
    // path, the Solana fee quoter behaves identically to the EVM implementation. Further
    // tests after this one simply observe the impact of modifying certain parameters on the
    // output.
    fn retrieving_fee_from_valid_message() {
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &sample_dest_chain(),
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 48282184443231661
            }
        );
    }

    #[test]
    fn network_fee_config_is_reflected_on_fee_retrieval() {
        set_syscall_stubs(Box::new(TestStubs));
        let mut chain = sample_dest_chain();
        chain.config.network_fee_usdcents *= 12;
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &chain,
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the network fee component of the sum
                amount: 298071755652939846
            }
        );
    }

    #[test]
    fn network_fee_for_an_unsupported_token_fails() {
        let mut message = sample_message();
        message.token_amounts = vec![SolanaTokenAmount {
            token: Pubkey::new_unique(),
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &sample_dest_chain(),
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap_err(),
            CcipRouterError::InvalidInputsMissingTokenConfig.into()
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_disabled_billing() {
        let mut chain = sample_dest_chain();

        // Will have no effect because we're not using the network fee
        chain.config.network_fee_usdcents = 0;

        let (token_config, mut per_chain_per_token) = sample_additional_token();

        // Not enabled == no overrides
        per_chain_per_token.billing.is_enabled = false;

        // Will have no effect since billing overrides are disabled
        per_chain_per_token.billing.min_fee_usdcents = 0;
        per_chain_per_token.billing.max_fee_usdcents = 0;

        let mut message = sample_message();
        message.token_amounts = vec![SolanaTokenAmount {
            token: per_chain_per_token.mint,
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(token_config)],
                &[per_chain_per_token]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 52885511932309044,
            }
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_enabled_billing() {
        let mut chain = sample_dest_chain();

        // Will have no effect because we're not using the network fee
        chain.config.network_fee_usdcents *= 0;
        let (another_token_config, mut another_per_chain_per_token_config) =
            sample_additional_token();

        another_per_chain_per_token_config.billing.min_fee_usdcents = 800;
        another_per_chain_per_token_config.billing.max_fee_usdcents = 1600;

        let mut message = sample_message();
        message.token_amounts = vec![SolanaTokenAmount {
            token: another_token_config.mint,
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config)],
                &[another_per_chain_per_token_config]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the min_fee
                amount: 398110981980079407
            }
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_bps() {
        let mut chain = sample_dest_chain();

        // Will have no effect because we're not using the network fee
        chain.config.network_fee_usdcents *= 0;
        let (another_token_config, mut another_per_chain_per_token_config) =
            sample_additional_token();

        // Set some arbitrary values so the bps fee lands between min and max
        another_per_chain_per_token_config.billing.min_fee_usdcents = 1;
        another_per_chain_per_token_config.billing.max_fee_usdcents = 1000;
        another_per_chain_per_token_config.billing.deci_bps = 10000;

        let mut message = sample_message();
        message.token_amounts = vec![SolanaTokenAmount {
            token: another_token_config.mint,
            amount: 15_000_000_000_000_000,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config.clone())],
                &[another_per_chain_per_token_config.clone()]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 36130696584140229
            }
        );

        // changing deci_bps affects the outcome.
        another_per_chain_per_token_config.billing.deci_bps = 20000;

        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config)],
                &[another_per_chain_per_token_config]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                // Slight increase in price
                amount: 37480696584140229
            }
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_no_fee_token_config() {
        let mut chain = sample_dest_chain();

        chain.config.network_fee_usdcents *= 0;
        let (_, mut another_per_chain_per_token_config) = sample_additional_token();

        // Will have no effect, as we cannot know the price of the token
        another_per_chain_per_token_config.billing.min_fee_usdcents = 1;
        another_per_chain_per_token_config.billing.max_fee_usdcents = 1000;
        another_per_chain_per_token_config.billing.deci_bps = 10000;

        let mut message = sample_message();
        message.token_amounts = vec![SolanaTokenAmount {
            token: another_per_chain_per_token_config.mint,
            amount: 15_000_000_000_000_000,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[None],
                &[another_per_chain_per_token_config.clone()]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 35234859440885153
            }
        );

        // Will have no effect, as we cannot know the price of the token
        another_per_chain_per_token_config.billing.deci_bps = 20000;

        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &[None],
                &[another_per_chain_per_token_config.clone()]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 35234859440885153
            }
        );
    }

    #[test]
    fn network_fee_for_multiple_tokens() {
        let (tokens, per_chains): (Vec<_>, Vec<_>) =
            (0..4).map(|_| sample_additional_token()).unzip();

        let mut message = sample_message();
        message.token_amounts = tokens
            .iter()
            .map(|t| SolanaTokenAmount {
                token: t.mint,
                amount: 1,
            })
            .collect();

        let tokens: Vec<_> = tokens.into_iter().map(|t| Some(t)).collect();
        let per_chains: Vec<_> = per_chains.into_iter().collect();
        set_syscall_stubs(Box::new(TestStubs));

        let mut chain = sample_dest_chain();
        chain.config.max_number_of_tokens_per_msg = 5;
        assert_eq!(
            fee_for_msg(
                0,
                &message,
                &chain,
                &sample_billing_config(),
                &tokens,
                &per_chains
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the number of tokens
                amount: 153233232867589323
            }
        );
    }

    fn sample_additional_token() -> (BillingTokenConfig, PerChainPerTokenConfig) {
        let mint = Pubkey::new_unique();
        (
            sample_billing_config(),
            PerChainPerTokenConfig {
                version: 1,
                chain_selector: 0,
                mint,
                billing: TokenBilling {
                    min_fee_usdcents: 50,
                    max_fee_usdcents: 4294967295,
                    deci_bps: 0,
                    dest_gas_overhead: 180000,
                    dest_bytes_overhead: 640,
                    is_enabled: true,
                },
            },
        )
    }

    #[test]
    fn fee_cannot_be_retrieved_when_token_price_is_not_timestamped() {
        let mut billing_config = sample_billing_config();
        billing_config.usd_per_token.timestamp = 0;
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &sample_dest_chain(),
                &billing_config,
                &[],
                &[]
            )
            .unwrap_err(),
            CcipRouterError::InvalidTokenPrice.into()
        );
    }

    #[test]
    fn fee_cannot_be_retrieved_when_token_price_is_zero() {
        let mut billing_config = sample_billing_config();
        billing_config.usd_per_token.value = [0u8; 28];
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &sample_dest_chain(),
                &billing_config,
                &[],
                &[]
            )
            .unwrap_err(),
            CcipRouterError::InvalidTokenPrice.into()
        );
    }

    #[test]
    fn fee_cannot_be_retrieved_when_gas_price_is_stale() {
        // This will make the unix timestamp be zero, so we'll adjust
        // the timestamp accordingly to a negative one.
        set_syscall_stubs(Box::new(TestStubs));
        let mut chain = sample_dest_chain();
        chain.state.usd_per_unit_gas.timestamp =
            -2 * chain.config.gas_price_staleness_threshold as i64;
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &chain,
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap_err(),
            CcipRouterError::StaleGasPrice.into()
        );
    }

    #[test]
    fn packing_unpacking_price() {
        let price = PackedPrice {
            execution_gas_price: Usd18Decimals::from_usd_cents(100),
            data_availability_gas_price: Usd18Decimals::from_usd_cents(200),
        };

        let roundtrip = PackedPrice::from(
            TryInto::<UnpackedDoubleU224>::try_into(price.clone())
                .unwrap()
                .pack(0)
                .unpack(),
        );
        assert_eq!(price, roundtrip);
    }
}
