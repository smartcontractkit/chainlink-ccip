use anchor_lang::prelude::*;
use anchor_spl::token::spl_token::native_mint;
use ethnum::U256;
use std::ops::AddAssign;

use crate::context::GetFee;
use crate::instructions::v1::safe_deserialize;
use crate::messages::{
    GetFeeResult, ProcessedExtraArgs, SVM2AnyMessage, SVMTokenAmount, TokenTransferAdditionalData,
};
use crate::state::{BillingTokenConfig, DestChain, PerChainPerTokenConfig, TimestampedPackedU224};
use crate::{FeeQuoterError, LINK_JUEL_DECIMALS};

use super::super::interfaces::Public;

use super::messages::validate_svm2any;
use super::price_math::{get_validated_token_price, Exponential, Usd18Decimals};

/// SVM2EVMRampMessage struct has 10 fields, including 3 variable unnested arrays (data, sender and tokenAmounts).
/// Each variable array takes 1 more slot to store its length.
/// When abi encoded, excluding array contents,
/// SVM2AnyMessage takes up a fixed number of 13 slots, 32 bytes each.
/// We assume sender always takes 1 slot.
/// For structs that contain arrays, 1 more slot is added to the front, reaching a total of 15.
/// The fixed bytes does not cover struct data (this is represented by SVM_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN)
pub const SVM_2_EVM_MESSAGE_FIXED_BYTES: U256 = U256::new(32 * 15);

/// Each token transfer adds 1 RampTokenAmount
/// RampTokenAmount has 5 fields, 2 of which are bytes type, 1 Address, 1 uint256 and 1 uint32.
/// Each bytes type takes 1 slot for length, 1 slot for data and 1 slot for the offset.
/// address
/// uint256 amount takes 1 slot.
/// uint32 destGasAmount takes 1 slot.
pub const SVM_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN: U256 = U256::new(32 * ((2 * 3) + 3));

pub const CCIP_LOCK_OR_BURN_V1_RET_BYTES: u32 = 32;

pub struct Impl;
impl Public for Impl {
    fn get_fee<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, GetFee>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
    ) -> Result<GetFeeResult> {
        let remaining_accounts = &ctx.remaining_accounts;
        let message = &message;
        require_eq!(
            remaining_accounts.len(),
            2 * message.token_amounts.len(),
            FeeQuoterError::InvalidInputsTokenAccounts
        );

        let (token_billing_config_accounts, per_chain_per_token_config_accounts) =
            remaining_accounts.split_at(message.token_amounts.len());

        let token_billing_config_accounts = token_billing_config_accounts
            .iter()
            .zip(message.token_amounts.iter())
            .map(|(a, SVMTokenAmount { token, .. })| {
                safe_deserialize::billing_token_config(a, *token)
            })
            .collect::<Result<Vec<_>>>()?;
        let per_chain_per_token_config_accounts = per_chain_per_token_config_accounts
            .iter()
            .zip(message.token_amounts.iter())
            .map(|(a, SVMTokenAmount { token, .. })| {
                safe_deserialize::per_chain_per_token_config(a, *token, dest_chain_selector)
            })
            .collect::<Result<Vec<_>>>()?;

        let (fee, processed_extra_args) = fee_for_msg(
            message,
            &ctx.accounts.dest_chain,
            &ctx.accounts.billing_token_config.config,
            &token_billing_config_accounts,
            &per_chain_per_token_config_accounts,
        )?;

        let fee_juels = fee_juels(
            &fee,
            &ctx.accounts.billing_token_config.config,
            &ctx.accounts.link_token_config.config,
            ctx.accounts.config.link_token_local_decimals,
            ctx.accounts.config.link_token_mint,
            ctx.accounts.config.max_fee_juels_per_msg,
        )?;

        let token_transfer_additional_data = per_chain_per_token_config_accounts
            .iter()
            .map(
                |per_chain_per_token_config| match per_chain_per_token_config {
                    Some(config) if config.token_transfer_config.is_enabled => {
                        TokenTransferAdditionalData {
                            dest_bytes_overhead: config.token_transfer_config.dest_bytes_overhead,
                            dest_gas_overhead: config.token_transfer_config.dest_gas_overhead,
                        }
                    }
                    _ => TokenTransferAdditionalData {
                        dest_bytes_overhead: CCIP_LOCK_OR_BURN_V1_RET_BYTES,
                        dest_gas_overhead: ctx
                            .accounts
                            .dest_chain
                            .config
                            .default_token_dest_gas_overhead,
                    },
                },
            )
            .collect();

        Ok(GetFeeResult {
            token: fee.token,
            amount: fee.amount,
            juels: fee_juels,
            token_transfer_additional_data,
            processed_extra_args,
        })
    }
}

fn fee_juels(
    fee: &SVMTokenAmount,
    source_config: &BillingTokenConfig,
    link_config: &BillingTokenConfig,
    link_local_decimals: u8,
    link_mint: Pubkey,
    max_fee_juels_per_msg: u128,
) -> Result<u128> {
    require_eq!(
        link_mint.key(),
        link_config.mint.key(),
        FeeQuoterError::InvalidInputsMint
    );
    let link_fee = convert(fee, source_config, link_config)?.amount;
    let fee_juels = (link_fee as u128)
        * 1u32.e(LINK_JUEL_DECIMALS
            .checked_sub(link_local_decimals)
            .expect("Guaranteed at construction"));
    let fee_juels: u128 = fee_juels
        .try_into()
        .expect("Impossible to surpass the 128 bit space with a maximum of 18 decimals.");
    require_gte!(
        max_fee_juels_per_msg,
        fee_juels,
        FeeQuoterError::MessageFeeTooHigh
    );
    Ok(fee_juels)
}

// Converts a token amount to one denominated in another token (e.g. from WSOL to LINK)
fn convert(
    source_token_amount: &SVMTokenAmount,
    source_config: &BillingTokenConfig,
    target_config: &BillingTokenConfig,
) -> Result<SVMTokenAmount> {
    assert!(source_config.mint == source_token_amount.token);
    let source_price = get_validated_token_price(source_config)?;
    let target_price = get_validated_token_price(target_config)?;

    Ok(SVMTokenAmount {
        token: target_config.mint,
        amount: ((source_price * source_token_amount.amount).0 / target_price.0)
            .try_into()
            .map_err(|_| FeeQuoterError::InvalidTokenPrice)?,
    })
}

fn fee_for_msg(
    message: &SVM2AnyMessage,
    dest_chain: &DestChain,
    fee_token_config: &BillingTokenConfig,
    additional_token_configs: &[Option<BillingTokenConfig>],
    additional_token_configs_for_dest_chain: &[Option<PerChainPerTokenConfig>],
) -> Result<(SVMTokenAmount, ProcessedExtraArgs)> {
    let fee_token = if message.fee_token == Pubkey::default() {
        native_mint::ID // Wrapped SOL
    } else {
        message.fee_token
    };
    require!(
        additional_token_configs.len() == message.token_amounts.len(),
        FeeQuoterError::InvalidInputsMissingTokenConfig
    );
    require!(
        additional_token_configs_for_dest_chain.len() == message.token_amounts.len(),
        FeeQuoterError::InvalidInputsMissingTokenConfig
    );
    let processed_extra_args = validate_svm2any(message, dest_chain, fee_token_config)?;

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

    let gas_limit = U256::new(processed_extra_args.gas_limit);

    // Calculate calldata gas cost while accounting for EIP-7623 variable calldata gas pricing
    // This logic works for EVMs post Pectra upgrade, while being backwards compatible with pre-Pectra EVMs.
    // This calculation is not exact, the goal is to not lose money on large payloads.
    // The fixed OCR report calldata overhead gas is accounted for in `dest_gas_overhead`.
    // It is not included in the calculation below for simplicity.
    let calldata_length =
        U256::new(message.data.len() as u128) + network_fee.transfer_bytes_overhead;
    let mut calldata_gas =
        calldata_length * U256::new(dest_chain.config.dest_gas_per_payload_byte_base as u128);
    let calldata_threshold =
        U256::new(dest_chain.config.dest_gas_per_payload_byte_threshold as u128);
    if calldata_length > calldata_threshold {
        let base_calldata_gas = U256::new(dest_chain.config.dest_gas_per_payload_byte_base as u128)
            * calldata_threshold;
        let extra_bytes = calldata_length - calldata_threshold;
        let extra_calldata_gas =
            extra_bytes * U256::new(dest_chain.config.dest_gas_per_payload_byte_high as u128);
        calldata_gas = base_calldata_gas + extra_calldata_gas;
    }
    let execution_gas = gas_limit
        + U256::new(dest_chain.config.dest_gas_overhead as u128)
        + calldata_gas
        + network_fee.transfer_gas;

    let execution_cost = execution_gas_price
        * execution_gas
        * U256::new(dest_chain.config.gas_multiplier_wei_per_eth as u128);

    let data_availability_cost: Usd18Decimals = data_availability_cost(
        data_availability_gas_price,
        message,
        network_fee.transfer_bytes_overhead,
        dest_chain,
    );

    let premium_multiplier = U256::new(fee_token_config.premium_multiplier_wei_per_eth.into());
    // At this step, every fee component has been raised to 36 decimals
    let fee_token_value =
        (network_fee.premium * premium_multiplier) + execution_cost + data_availability_cost;

    // Fee token value is in 36 decimals
    // Fee token price is in 18 decimals USD for 1e18 smallest token denominations.
    // The result is the fee in the fee tokens smallest denominations (e.g. lamport for Sol).
    let fee_token_amount = (fee_token_value.0 / fee_token_price.0)
        .try_into()
        .map_err(|_| FeeQuoterError::InvalidTokenPrice)?;

    Ok((
        SVMTokenAmount {
            token: fee_token,
            amount: fee_token_amount,
        },
        processed_extra_args,
    ))
}

fn data_availability_cost(
    data_availability_gas_price: Usd18Decimals,
    message: &SVM2AnyMessage,
    token_transfer_bytes_overhead: U256,
    dest_chain: &DestChain,
) -> Usd18Decimals {
    // Sums up byte lengths of fixed message fields and dynamic message fields.
    // Fixed message fields do account for the offset and length slot of the dynamic fields.
    let data_availability_length_bytes = SVM_2_EVM_MESSAGE_FIXED_BYTES
        + U256::new(message.data.len() as u128)
        + (U256::new(message.token_amounts.len() as u128)
            * SVM_2_EVM_MESSAGE_FIXED_BYTES_PER_TOKEN)
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
    message: &SVM2AnyMessage,
    dest_chain: &DestChain,
    token_configs: &[Option<BillingTokenConfig>],
    token_configs_for_dest_chain: &[Option<PerChainPerTokenConfig>],
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

        let token_network_fee = match config_for_dest_chain {
            Some(config) if config.token_transfer_config.is_enabled => {
                token_network_fees(&token_configs[i], token_amount, config)?
            }
            // If the token has no specific overrides configured, we use the global defaults.
            _ => default_token_network_fees(dest_chain),
        };

        fee += token_network_fee;
    }

    Ok(fee)
}

fn token_network_fees(
    billing_config: &Option<BillingTokenConfig>,
    token_amount: &SVMTokenAmount,
    config_for_dest_chain: &PerChainPerTokenConfig,
) -> Result<NetworkFee> {
    let bps_fee = match billing_config {
        Some(config) if config_for_dest_chain.token_transfer_config.deci_bps > 0 => {
            let token_price = get_validated_token_price(config)?;
            // Calculate token transfer value, then apply fee ratio
            // ratio represents multiples of 0.1bps, or 1e-5
            Usd18Decimals(
                Usd18Decimals::from_token_amount(token_amount, &token_price).0
                    * U256::new(config_for_dest_chain.token_transfer_config.deci_bps.into())
                    / 1u32.e(5),
            )
        }
        _ => Usd18Decimals::ZERO,
    };

    let min_fee =
        Usd18Decimals::from_usd_cents(config_for_dest_chain.token_transfer_config.min_fee_usdcents);
    let max_fee =
        Usd18Decimals::from_usd_cents(config_for_dest_chain.token_transfer_config.max_fee_usdcents);
    let (premium, token_transfer_gas, token_transfer_bytes_overhead) = (
        bps_fee.clamp(min_fee, max_fee),
        U256::new(
            config_for_dest_chain
                .token_transfer_config
                .dest_gas_overhead
                .into(),
        ),
        U256::new(
            config_for_dest_chain
                .token_transfer_config
                .dest_bytes_overhead
                .into(),
        ),
    );
    Ok(NetworkFee {
        premium,
        transfer_gas: token_transfer_gas,
        transfer_bytes_overhead: token_transfer_bytes_overhead,
    })
}

fn default_token_network_fees(dest_chain: &DestChain) -> NetworkFee {
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
struct PackedPrice {
    // L1 gas price (encoded in the lower 112 bits)
    pub execution_gas_price: Usd18Decimals,
    // L2 gas price (encoded in the higher 112 bits)
    pub data_availability_gas_price: Usd18Decimals,
}

impl From<&TimestampedPackedU224> for PackedPrice {
    fn from(value: &TimestampedPackedU224) -> Self {
        UnpackedDoubleU224::from(value).into()
    }
}

#[derive(Debug, Clone)]
struct UnpackedDoubleU224 {
    pub high: u128,
    pub low: u128,
}

impl From<&TimestampedPackedU224> for UnpackedDoubleU224 {
    fn from(packed: &TimestampedPackedU224) -> Self {
        let mut u128_buffer = [0u8; 16];
        u128_buffer[2..16].clone_from_slice(&packed.value[14..]);
        let low = u128::from_be_bytes(u128_buffer);
        u128_buffer[2..16].clone_from_slice(&packed.value[..14]);
        let high = u128::from_be_bytes(u128_buffer);
        Self { high, low }
    }
}

#[allow(clippy::from_over_into)] // we don't want to implement methods for state structs, only for internal ones
impl Into<PackedPrice> for UnpackedDoubleU224 {
    fn into(self) -> PackedPrice {
        PackedPrice {
            execution_gas_price: Usd18Decimals(self.low.into()),
            data_availability_gas_price: Usd18Decimals(self.high.into()),
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
                .map_err(|_| FeeQuoterError::InvalidTokenPrice)?,
            low: value
                .execution_gas_price
                .0
                .try_into()
                .map_err(|_| FeeQuoterError::InvalidTokenPrice)?,
        })
    }
}

fn get_validated_gas_price(dest_chain: &DestChain) -> Result<PackedPrice> {
    let timestamp = dest_chain.state.usd_per_unit_gas.timestamp;
    let price = PackedPrice::from(&dest_chain.state.usd_per_unit_gas);
    let threshold = dest_chain.config.gas_price_staleness_threshold as i64;
    let elapsed_time = Clock::get()?.unix_timestamp - timestamp;

    require!(
        threshold == 0 || threshold > elapsed_time,
        FeeQuoterError::StaleGasPrice
    );

    Ok(price)
}

#[cfg(test)]
mod tests {
    use solana_program::{
        entrypoint::SUCCESS,
        program_stubs::{set_syscall_stubs, SyscallStubs},
    };

    use super::super::messages::tests::{sample_billing_config, sample_dest_chain, sample_message};
    use crate::TokenTransferFeeConfig;

    use super::*;

    struct TestStubs;

    impl SyscallStubs for TestStubs {
        fn sol_get_clock_sysvar(&self, _var_addr: *mut u8) -> u64 {
            // This causes the syscall to return a default-initialized
            // clock when the build target is off-chain. Good enough for tests.
            SUCCESS
        }
    }

    impl UnpackedDoubleU224 {
        pub fn pack(self, timestamp: i64) -> TimestampedPackedU224 {
            let mut value = [0u8; 28];
            value[14..].clone_from_slice(&self.low.to_be_bytes()[2..16]);
            value[..14].clone_from_slice(&self.high.to_be_bytes()[2..16]);
            TimestampedPackedU224 { value, timestamp }
        }
    }

    #[test]
    // NOTE: This test is unique in that the return value of `fee_for_msg` has been
    // directly validated against a `getFee` query in the Ethereum mainnet -> Avalanche Fuji
    // lane, using the same configuration. This ensures that at least on a simple execution
    // path, the SVM fee quoter behaves identically to the EVM implementation. Further
    // tests after this one simply observe the impact of modifying certain parameters on the
    // output.
    fn retrieving_fee_from_valid_message() {
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &sample_message(),
                &sample_dest_chain(),
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
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
                &sample_message(),
                &chain,
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the network fee component of the sum
                amount: 298071755652939846
            }
        );
    }

    #[test]
    fn network_fee_for_an_unsupported_token_fails() {
        let mut message = sample_message();
        message.token_amounts = vec![SVMTokenAmount {
            token: Pubkey::new_unique(),
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &message,
                &sample_dest_chain(),
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap_err(),
            FeeQuoterError::InvalidInputsMissingTokenConfig.into()
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_disabled_billing() {
        let mut chain = sample_dest_chain();

        // Will have no effect because we're not using the network fee
        chain.config.network_fee_usdcents = 0;

        let (token_config, mut per_chain_per_token) = sample_additional_token();

        // Not enabled == no overrides
        per_chain_per_token.token_transfer_config.is_enabled = false;

        // Will have no effect since billing overrides are disabled
        per_chain_per_token.token_transfer_config.min_fee_usdcents = 0;
        per_chain_per_token.token_transfer_config.max_fee_usdcents = 0;

        let mut message = sample_message();
        message.token_amounts = vec![SVMTokenAmount {
            token: per_chain_per_token.mint,
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(token_config)],
                &[Some(per_chain_per_token)]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                amount: 52911699750913573,
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

        another_per_chain_per_token_config
            .token_transfer_config
            .min_fee_usdcents = 800;
        another_per_chain_per_token_config
            .token_transfer_config
            .max_fee_usdcents = 1600;

        let mut message = sample_message();
        message.token_amounts = vec![SVMTokenAmount {
            token: another_token_config.mint,
            amount: 1,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config)],
                &[Some(another_per_chain_per_token_config)]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the min_fee
                amount: 398634738352169990
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
        another_per_chain_per_token_config
            .token_transfer_config
            .min_fee_usdcents = 1;
        another_per_chain_per_token_config
            .token_transfer_config
            .max_fee_usdcents = 1000;
        another_per_chain_per_token_config
            .token_transfer_config
            .deci_bps = 10000;

        let mut message = sample_message();
        message.token_amounts = vec![SVMTokenAmount {
            token: another_token_config.mint,
            amount: 15_000_000_000_000_000,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config.clone())],
                &[Some(another_per_chain_per_token_config.clone())]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                amount: 36654452956230811
            }
        );

        // changing deci_bps affects the outcome.
        another_per_chain_per_token_config
            .token_transfer_config
            .deci_bps = 20000;

        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[Some(another_token_config)],
                &[Some(another_per_chain_per_token_config)]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                // Slight increase in price
                amount: 38004452956230811
            }
        );
    }

    #[test]
    fn network_fee_for_a_supported_token_with_no_fee_token_config() {
        let mut chain = sample_dest_chain();

        chain.config.network_fee_usdcents *= 0;
        let (_, mut another_per_chain_per_token_config) = sample_additional_token();

        // Will have no effect, as we cannot know the price of the token
        another_per_chain_per_token_config
            .token_transfer_config
            .min_fee_usdcents = 1;
        another_per_chain_per_token_config
            .token_transfer_config
            .max_fee_usdcents = 1000;
        another_per_chain_per_token_config
            .token_transfer_config
            .deci_bps = 10000;

        let mut message = sample_message();
        message.token_amounts = vec![SVMTokenAmount {
            token: another_per_chain_per_token_config.mint,
            amount: 15_000_000_000_000_000,
        }];
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[None],
                &[Some(another_per_chain_per_token_config.clone())]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                amount: 35758615812975735
            }
        );

        // Will have no effect, as we cannot know the price of the token
        another_per_chain_per_token_config
            .token_transfer_config
            .deci_bps = 20000;

        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &[None],
                &[Some(another_per_chain_per_token_config.clone())]
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                amount: 35758615812975735
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
            .map(|t| SVMTokenAmount {
                token: t.mint,
                amount: 1,
            })
            .collect();

        let tokens: Vec<_> = tokens.into_iter().map(Some).collect();
        let per_chains: Vec<_> = per_chains.into_iter().map(Some).collect();
        set_syscall_stubs(Box::new(TestStubs));

        let mut chain = sample_dest_chain();
        chain.config.max_number_of_tokens_per_msg = 5;
        assert_eq!(
            fee_for_msg(
                &message,
                &chain,
                &sample_billing_config(),
                &tokens,
                &per_chains
            )
            .unwrap()
            .0,
            SVMTokenAmount {
                token: native_mint::ID,
                // Increases proportionally to the number of tokens
                amount: 155328258355951652
            }
        );
    }

    pub fn sample_additional_token() -> (BillingTokenConfig, PerChainPerTokenConfig) {
        let mint = Pubkey::new_unique();
        (
            sample_billing_config(),
            PerChainPerTokenConfig {
                version: 1,
                chain_selector: 0,
                mint,
                token_transfer_config: TokenTransferFeeConfig {
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
                &sample_message(),
                &sample_dest_chain(),
                &billing_config,
                &[],
                &[]
            )
            .unwrap_err(),
            FeeQuoterError::InvalidTokenPrice.into()
        );
    }

    #[test]
    fn fee_cannot_be_retrieved_when_token_price_is_zero() {
        let mut billing_config = sample_billing_config();
        billing_config.usd_per_token.value = [0u8; 28];
        assert_eq!(
            fee_for_msg(
                &sample_message(),
                &sample_dest_chain(),
                &billing_config,
                &[],
                &[]
            )
            .unwrap_err(),
            FeeQuoterError::InvalidTokenPrice.into()
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
                &sample_message(),
                &chain,
                &sample_billing_config(),
                &[],
                &[]
            )
            .unwrap_err(),
            FeeQuoterError::StaleGasPrice.into()
        );
    }

    #[test]
    fn packing_unpacking_price() {
        let price = PackedPrice {
            execution_gas_price: Usd18Decimals::from_usd_cents(100),
            data_availability_gas_price: Usd18Decimals::from_usd_cents(200),
        };

        let unpacked: UnpackedDoubleU224 = price.clone().try_into().unwrap();
        let ts_packed: TimestampedPackedU224 = unpacked.pack(0);
        let roundtrip: PackedPrice = (&ts_packed).into();

        assert_eq!(price, roundtrip);
    }

    #[test]
    fn test_packed_price_from_bytes() {
        // 2gwei DA encoded in higher order 112 bits, 1gwei Exec gas price encoded in lower order 112 bits
        let bytes = [
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1b, 0xc1, 0x6d, 0x67, 0x4e, 0xc8, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0xe0, 0xb6, 0xb3, 0xa7, 0x64, 0x00, 0x00,
        ];

        let ts_packed = TimestampedPackedU224 {
            value: bytes,
            timestamp: 0,
        };

        let price: PackedPrice = (&ts_packed).into();

        assert_eq!(price.execution_gas_price, Usd18Decimals(1u32.e(18)));
        assert_eq!(price.data_availability_gas_price, Usd18Decimals(2u32.e(18)));
    }

    #[test]
    fn converting_fee_to_juels() {
        let fee_billing_token_config = sample_billing_config();
        let fee = SVMTokenAmount {
            token: fee_billing_token_config.mint,
            amount: 100,
        };

        let link_mint = Pubkey::new_unique();
        let mut link_billing_config = sample_billing_config();
        link_billing_config.mint = link_mint;

        let link_local_decimals = 9;

        let juels = fee_juels(
            &fee,
            &fee_billing_token_config,
            &link_billing_config,
            link_local_decimals,
            link_mint,
            (200u128 * 1u32.e(9)).try_into().unwrap(),
        )
        .unwrap();

        // As per this test, both the fee token and LINK have the same value, so the difference in amount
        // will be proportional to the difference in decimals between solana's representation (1e9 divisions = 1 LINK)
        // and the canonical one in evm (1e18 juels = 1 LINK).
        assert_eq!(juels, fee.amount as u128 * 1u32.e(9));

        // Fails if it exceeds the maximum
        fee_juels(
            &fee,
            &fee_billing_token_config,
            &link_billing_config,
            link_local_decimals,
            link_mint,
            (50u128 * 1u32.e(9)).try_into().unwrap(),
        )
        .unwrap_err();

        // Returns the same value when decimals are equal
        let juels = fee_juels(
            &fee,
            &fee_billing_token_config,
            &link_billing_config,
            LINK_JUEL_DECIMALS,
            link_mint,
            (200u128 * 1u32.e(9)).try_into().unwrap(),
        )
        .unwrap();

        assert_eq!(juels, fee.amount as u128);

        // 0 decimals edge case
        let juels = fee_juels(
            &fee,
            &fee_billing_token_config,
            &link_billing_config,
            0,
            link_mint,
            (200u128 * 1u32.e(18)).try_into().unwrap(),
        )
        .unwrap();

        assert_eq!(juels, fee.amount as u128 * 1u32.e(18));
    }

    #[test]
    fn test_max_fee_juels_per_msg_validation() {
        set_syscall_stubs(Box::new(TestStubs));

        // create custom billing configs with identical prices for deterministic testing
        let mut fee_token_config = sample_billing_config();
        let mut link_config = sample_billing_config();

        // set a specific timestamp and price value
        let timestamp = 100;

        // create a value for 1 USD per token with 18 decimals precision
        // this is a simplified value (1 * 10^18) for easier calculations
        let price_bytes = [
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0xe0, 0xb6, 0xb3, 0xa7, 0x64, 0x00,
        ]; // 1 * 10^18 (1.0 with 18 decimals)

        fee_token_config.usd_per_token.value = price_bytes;
        fee_token_config.usd_per_token.timestamp = timestamp;

        link_config.usd_per_token.value = price_bytes;
        link_config.usd_per_token.timestamp = timestamp;

        // define LINK mint pubkey
        let link_mint = Pubkey::new_unique();
        link_config.mint = link_mint;

        // use 9 decimals for local LINK representation (typical for Solana)
        let link_local_decimals = 9;

        // set max fee to exactly 1 LINK in juels (18 decimals)
        let max_fee_juels = 1_000_000_000_000_000_000u128;

        // test with fee amount just below max (should pass)
        // 0.999999999 LINK in 9 decimals = 999,999,999
        // converts to 999,999,999,000,000,000 juels (< max)
        let under_fee = SVMTokenAmount {
            token: fee_token_config.mint,
            amount: 999_999_999,
        };

        let result = fee_juels(
            &under_fee,
            &fee_token_config,
            &link_config,
            link_local_decimals,
            link_mint,
            max_fee_juels,
        );

        assert!(result.is_ok(), "under max fee should be accepted");

        // test with fee amount over max (should fail)
        // 1.000000001 LINK in 9 decimals = 1,000,000,001
        // converts to 1,000,000,001,000,000,000 juels (> max)
        let over_fee = SVMTokenAmount {
            token: fee_token_config.mint,
            amount: 1_000_000_001,
        };

        let result = fee_juels(
            &over_fee,
            &fee_token_config,
            &link_config,
            link_local_decimals,
            link_mint,
            max_fee_juels,
        );

        assert!(result.is_err(), "over max fee should be rejected");
        assert_eq!(
            result.unwrap_err(),
            FeeQuoterError::MessageFeeTooHigh.into()
        );
    }

    #[test]
    fn test_fee_juels_decimal_conversions() {
        set_syscall_stubs(Box::new(TestStubs));

        // create test billing configs with identical prices
        let mut fee_token_config = sample_billing_config();
        let mut link_config = sample_billing_config();

        // set a specific timestamp and price value
        let timestamp = 100;

        // create a value for 1 USD per token with 18 decimals precision
        let price_bytes = [
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0xe0, 0xb6, 0xb3, 0xa7, 0x64, 0x00,
        ]; // 1 * 10^18 (1.0 with 18 decimals)

        fee_token_config.usd_per_token.value = price_bytes;
        fee_token_config.usd_per_token.timestamp = timestamp;

        link_config.usd_per_token.value = price_bytes;
        link_config.usd_per_token.timestamp = timestamp;

        // define LINK mint pubkey
        let link_mint = Pubkey::new_unique();
        link_config.mint = link_mint;

        // case 1: 9 decimals local, fee under limit
        {
            let link_local_decimals = 9;
            let fee_amount = 100_000_000; // 0.1 LINK in 9 decimals
            let max_juels = 100_000_000_000_000_000; // 0.1 LINK in 18 decimals

            let fee = SVMTokenAmount {
                token: fee_token_config.mint,
                amount: fee_amount,
            };

            let result = fee_juels(
                &fee,
                &fee_token_config,
                &link_config,
                link_local_decimals,
                link_mint,
                max_juels,
            );

            if result.is_err() {
                println!("Case 1 error: {:?}", result.unwrap_err());
                unreachable!("Case 1 should pass");
            } else {
                let juels = result.unwrap();
                println!(
                    "Case 1: {} with {} decimals converted to {} juels",
                    fee_amount, link_local_decimals, juels
                );
            }
        }

        // case 2: 9 decimals local, fee over limit
        {
            let link_local_decimals = 9;
            let fee_amount = 200_000_000; // 0.2 LINK in 9 decimals
            let max_juels = 100_000_000_000_000_000; // 0.1 LINK in 18 decimals

            let fee = SVMTokenAmount {
                token: fee_token_config.mint,
                amount: fee_amount,
            };

            let result = fee_juels(
                &fee,
                &fee_token_config,
                &link_config,
                link_local_decimals,
                link_mint,
                max_juels,
            );

            assert!(result.is_err(), "Case 2 should fail");
            if let Err(err) = result {
                assert_eq!(err, FeeQuoterError::MessageFeeTooHigh.into());
            }
        }

        // case 3: same decimals (18), no scaling needed
        {
            let link_local_decimals = 18;
            let fee_amount = 50_000_000_000_000_000; // 0.05 LINK in 18 decimals
            let max_juels = 100_000_000_000_000_000; // 0.1 LINK in 18 decimals

            let fee = SVMTokenAmount {
                token: fee_token_config.mint,
                amount: fee_amount,
            };

            let result = fee_juels(
                &fee,
                &fee_token_config,
                &link_config,
                link_local_decimals,
                link_mint,
                max_juels,
            );

            if result.is_err() {
                println!("Case 3 error: {:?}", result.unwrap_err());
                unreachable!("Case 3 should pass");
            } else {
                let juels = result.unwrap();
                println!(
                    "Case 3: {} with {} decimals converted to {} juels",
                    fee_amount, link_local_decimals, juels
                );
                assert_eq!(juels, fee_amount as u128);
            }
        }

        // case 4: 6 decimals local (scale up by 12)
        {
            let link_local_decimals = 6;
            let fee_amount = 50_000; // 0.05 LINK in 6 decimals
            let max_juels = 100_000_000_000_000_000; // 0.1 LINK in 18 decimals

            let fee = SVMTokenAmount {
                token: fee_token_config.mint,
                amount: fee_amount,
            };

            let result = fee_juels(
                &fee,
                &fee_token_config,
                &link_config,
                link_local_decimals,
                link_mint,
                max_juels,
            );

            if result.is_err() {
                println!("Case 4 error: {:?}", result.unwrap_err());
                unreachable!("Case 4 should pass");
            } else {
                let juels = result.unwrap();
                println!(
                    "Case 4: {} with {} decimals converted to {} juels",
                    fee_amount, link_local_decimals, juels
                );
                assert_eq!(juels, 50_000_000_000_000_000, "Case 4 juels mismatch");
            }
        }

        // case 5: 0 decimals local (scale up by 18)
        {
            let link_local_decimals = 0;
            let fee_amount = 1; // 1 LINK in 0 decimals
            let max_juels = 2_000_000_000_000_000_000; // 2 LINK in 18 decimals

            let fee = SVMTokenAmount {
                token: fee_token_config.mint,
                amount: fee_amount,
            };

            let result = fee_juels(
                &fee,
                &fee_token_config,
                &link_config,
                link_local_decimals,
                link_mint,
                max_juels,
            );

            if result.is_err() {
                println!("Case 5 error: {:?}", result.unwrap_err());
                unreachable!("Case 5 should pass");
            } else {
                let juels = result.unwrap();
                println!(
                    "Case 5: {} with {} decimals converted to {} juels",
                    fee_amount, link_local_decimals, juels
                );
                assert_eq!(juels, 1_000_000_000_000_000_000, "Case 5 juels mismatch");
            }
        }
    }
}
