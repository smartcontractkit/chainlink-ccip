use anchor_lang::prelude::*;
use anchor_spl::{token::spl_token::native_mint, token_interface};
use ethnum::U256;
use solana_program::{program::invoke_signed, system_instruction};

use crate::{
    utils::{Exponential, Usd18Decimals},
    BillingTokenConfig, CcipRouterError, DestChain, PerChainPerTokenConfig, Solana2AnyMessage,
    SolanaTokenAmount, UnpackedDoubleU224, CCIP_LOCK_OR_BURN_V1_RET_BYTES,
    FEE_BILLING_SIGNER_SEEDS,
};

// TODO change args and implement
pub fn fee_for_msg(
    _dest_chain_selector: u64,
    message: &Solana2AnyMessage,
    dest_chain: &DestChain,
    token_configs: &[&BillingTokenConfig],
    token_configs_for_dest_chain: &[&PerChainPerTokenConfig],
) -> Result<SolanaTokenAmount> {
    let token = if message.fee_token == Pubkey::default() {
        native_mint::ID // Wrapped SOL
    } else {
        message.fee_token
    };
    let fee_token_config = token_configs
        .iter()
        .find(|c| c.mint == token)
        .ok_or(CcipRouterError::UnsupportedToken)?;
    message.validate(dest_chain, fee_token_config)?;

    let fee_token_price = get_validated_token_price(&fee_token_config)?;
    let _packed_gas_price = get_validated_gas_price(dest_chain)?;

    let network_fee = network_fee(
        message,
        dest_chain,
        token_configs,
        token_configs_for_dest_chain,
    )?;

    // TODO un-hardcode
    let execution_cost = 1u32.e(18);
    let data_availability_cost = 1u32.e(18);

    let premium_multiplier = U256::new(fee_token_config.premium_multiplier_wei_per_eth.into());
    let amount = (dbg!(network_fee.premium.0) * dbg!(premium_multiplier)
        + execution_cost
        + data_availability_cost)
        / dbg!(fee_token_price.0);
    let amount: u64 = amount
        .try_into()
        .map_err(|_| CcipRouterError::InvalidTokenPrice)?;

    Ok(SolanaTokenAmount { amount, token })
}

struct NetworkFee {
    premium: Usd18Decimals,
    _token_transfer_gas: U256,
    _token_transfer_bytes_overhead: U256,
}

fn network_fee(
    message: &Solana2AnyMessage,
    dest_chain: &DestChain,
    token_configs: &[&BillingTokenConfig],
    token_configs_for_dest_chain: &[&PerChainPerTokenConfig],
) -> Result<NetworkFee> {
    if message.token_amounts.is_empty() {
        return Ok(NetworkFee {
            premium: Usd18Decimals::from_usd_cents(dest_chain.config.network_fee_usdcents),
            _token_transfer_gas: U256::ZERO,
            _token_transfer_bytes_overhead: U256::ZERO,
        });
    }

    let mut premium = Usd18Decimals::default();
    let mut token_transfer_gas = U256::ZERO;
    let mut token_transfer_bytes_overhead = U256::ZERO;

    for token_amount in &message.token_amounts {
        let config_for_dest_chain = token_configs_for_dest_chain
            .iter()
            .find(|c| c.mint == token_amount.token)
            .ok_or(CcipRouterError::UnsupportedToken)?;

        let (token_fee, token_gas, token_overhead) = if config_for_dest_chain.billing.is_enabled {
            token_network_fees(token_configs, token_amount, config_for_dest_chain)?
        } else {
            // If the token has no specific overrides configured, we use the global defaults.
            global_network_fees(dest_chain)
        };

        premium += token_fee;
        token_transfer_gas += token_gas;
        token_transfer_bytes_overhead += token_overhead;
    }

    Ok(NetworkFee {
        premium,
        _token_transfer_gas: token_transfer_gas,
        _token_transfer_bytes_overhead: token_transfer_bytes_overhead,
    })
}

fn token_network_fees(
    token_configs: &[&BillingTokenConfig],
    token_amount: &SolanaTokenAmount,
    config_for_dest_chain: &PerChainPerTokenConfig,
) -> Result<(Usd18Decimals, U256, U256)> {
    let config = token_configs
        .iter()
        .find(|c| c.mint == token_amount.token)
        .ok_or(CcipRouterError::UnsupportedToken)?;

    let bps_fee = if config_for_dest_chain.billing.deci_bps > 0 {
        let token_price = get_validated_token_price(&config)?;
        // Calculate token transfer value, then apply fee ratio
        // ratio represents multiples of 0.1bps, or 1e-5
        Usd18Decimals(
            token_amount.value(&token_price).0
                * U256::new(config_for_dest_chain.billing.deci_bps.into())
                / 1u32.e(5),
        )
    } else {
        Usd18Decimals::default()
    };
    let min_fee = Usd18Decimals::from_usd_cents(config_for_dest_chain.billing.min_fee_usdcents);
    let max_fee = Usd18Decimals::from_usd_cents(config_for_dest_chain.billing.max_fee_usdcents);
    let (token_fee, token_gas, token_overhead) = (
        bps_fee.clamp(min_fee, max_fee),
        U256::new(config_for_dest_chain.billing.dest_gas_overhead.into()),
        U256::new(config_for_dest_chain.billing.dest_bytes_overhead.into()),
    );
    Ok((token_fee, token_gas, token_overhead))
}

fn global_network_fees(dest_chain: &DestChain) -> (Usd18Decimals, U256, U256) {
    let (global_fee, global_gas, global_overhead) = (
        Usd18Decimals::from_usd_cents(dest_chain.config.default_token_fee_usdcents.into()),
        U256::new(dest_chain.config.default_token_dest_gas_overhead.into()),
        U256::new(CCIP_LOCK_OR_BURN_V1_RET_BYTES.into()),
    );
    (global_fee, global_gas, global_overhead)
}

pub struct PackedPrice {
    pub execution_cost: u128,
    pub gas_price: u128,
}

impl From<UnpackedDoubleU224> for PackedPrice {
    fn from(value: UnpackedDoubleU224) -> Self {
        Self {
            execution_cost: value.low,
            gas_price: value.high,
        }
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

    let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let cpi_ctx = CpiContext::new_with_signer(token_program, transfer, signer_seeds);
    token_interface::transfer_checked(cpi_ctx, fee.amount, decimals)
}

#[cfg(test)]
mod tests {
    use solana_program::{
        entrypoint::SUCCESS,
        program_stubs::{set_syscall_stubs, SyscallStubs},
    };

    use super::*;
    use crate::tests::{sample_billing_config, sample_dest_chain, sample_message};

    struct TestStubs;

    impl SyscallStubs for TestStubs {
        fn sol_get_clock_sysvar(&self, _var_addr: *mut u8) -> u64 {
            // This causes the syscall to return a default-initialized
            // clock when the build target is off-chain. Good enough for tests.
            SUCCESS
        }
    }

    #[test]
    fn retrieving_fee_from_valid_message() {
        set_syscall_stubs(Box::new(TestStubs));
        assert_eq!(
            fee_for_msg(
                0,
                &sample_message(),
                &sample_dest_chain(),
                &[&sample_billing_config()],
                &[]
            )
            .unwrap(),
            SolanaTokenAmount {
                token: native_mint::ID,
                amount: 1
            }
        );
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
                &[&billing_config],
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
                &[&billing_config],
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
                &[&sample_billing_config()],
                &[]
            )
            .unwrap_err(),
            CcipRouterError::StaleGasPrice.into()
        );
    }
}
