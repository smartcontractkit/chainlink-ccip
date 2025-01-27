use anchor_lang::prelude::*;
use anchor_spl::token::spl_token::native_mint;
use ethnum::U256;
use std::ops::AddAssign;

use crate::context::GetFee;
use crate::instructions::v1::safe_deserialize;
use crate::messages::{SVM2AnyMessage, SVMTokenAmount};
use crate::state::{BillingTokenConfig, DestChain, PerChainPerTokenConfig, TimestampedPackedU224};
use crate::FeeQuoterError;

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

pub fn get_fee<'info>(
    ctx: Context<'_, '_, 'info, 'info, GetFee>,
    dest_chain_selector: u64,
    message: SVM2AnyMessage,
) -> Result<u64> {
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
        .map(|(a, SVMTokenAmount { token, .. })| safe_deserialize::billing_token_config(a, *token))
        .collect::<Result<Vec<_>>>()?;
    let per_chain_per_token_config_accounts = per_chain_per_token_config_accounts
        .iter()
        .zip(message.token_amounts.iter())
        .map(|(a, SVMTokenAmount { token, .. })| {
            safe_deserialize::per_chain_per_token_config(a, *token, dest_chain_selector)
        })
        .collect::<Result<Vec<_>>>()?;

    Ok(fee_for_msg(
        message,
        &ctx.accounts.dest_chain,
        &ctx.accounts.billing_token_config.config,
        &token_billing_config_accounts,
        &per_chain_per_token_config_accounts,
    )?
    .amount)
}

fn fee_for_msg(
    message: &SVM2AnyMessage,
    dest_chain: &DestChain,
    fee_token_config: &BillingTokenConfig,
    additional_token_configs: &[Option<BillingTokenConfig>],
    additional_token_configs_for_dest_chain: &[PerChainPerTokenConfig],
) -> Result<SVMTokenAmount> {
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
    validate_svm2any(message, dest_chain, fee_token_config)?;

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
        + (U256::new(message.data.len() as u128) + network_fee.transfer_bytes_overhead)
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

    let fee_token_amount = (fee_token_value.0 / fee_token_price.0)
        .try_into()
        .map_err(|_| FeeQuoterError::InvalidTokenPrice)?;

    Ok(SVMTokenAmount {
        token: fee_token,
        amount: fee_token_amount,
    })
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
            default_token_network_fees(dest_chain)
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
        Some(config) if config_for_dest_chain.billing.deci_bps > 0 => {
            let token_price = get_validated_token_price(config)?;
            // Calculate token transfer value, then apply fee ratio
            // ratio represents multiples of 0.1bps, or 1e-5
            Usd18Decimals(
                Usd18Decimals::from_token_amount(token_amount, &token_price).0
                    * U256::new(config_for_dest_chain.billing.deci_bps.into())
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
pub(super) struct PackedPrice {
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
pub(super) struct UnpackedDoubleU224 {
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

// TODO move unit tests here
