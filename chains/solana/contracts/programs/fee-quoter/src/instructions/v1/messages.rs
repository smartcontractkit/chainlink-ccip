use anchor_lang::prelude::*;
use ethnum::U256;

use crate::messages::{SVM2AnyMessage, CHAIN_FAMILY_SELECTOR_EVM};
use crate::state::{BillingTokenConfig, DestChain};
use crate::FeeQuoterError;

const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);

pub fn validate_svm2any(
    msg: &SVM2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
) -> Result<()> {
    require!(
        dest_chain.config.is_enabled,
        FeeQuoterError::DestinationChainDisabled
    );

    require!(token_config.enabled, FeeQuoterError::FeeTokenDisabled);

    require_gte!(
        dest_chain.config.max_data_bytes,
        msg.data.len() as u32,
        FeeQuoterError::MessageTooLarge
    );

    require_gte!(
        dest_chain.config.max_number_of_tokens_per_msg as usize,
        msg.token_amounts.len(),
        FeeQuoterError::UnsupportedNumberOfTokens
    );

    require_gte!(
        dest_chain.config.max_per_msg_gas_limit as u128,
        msg.extra_args
            .gas_limit
            .unwrap_or(dest_chain.config.default_tx_gas_limit as u128),
        FeeQuoterError::MessageGasLimitTooHigh,
    );

    require!(
        !dest_chain.config.enforce_out_of_order
            || msg.extra_args.allow_out_of_order_execution.unwrap_or(false),
        FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue,
    );

    validate_dest_family_address(msg, dest_chain.config.chain_family_selector)
}

fn validate_dest_family_address(
    msg: &SVM2AnyMessage,
    chain_family_selector: [u8; 4],
) -> Result<()> {
    const PRECOMPILE_SPACE: u32 = 1024;

    let selector = u32::from_be_bytes(chain_family_selector);
    // Only EVM is supported as a destination family.
    require_eq!(
        selector,
        CHAIN_FAMILY_SELECTOR_EVM,
        FeeQuoterError::UnsupportedChainFamilySelector
    );

    require_eq!(msg.receiver.len(), 32, FeeQuoterError::InvalidEVMAddress);

    let address: U256 = U256::from_be_bytes(
        msg.receiver
            .clone()
            .try_into()
            .map_err(|_| FeeQuoterError::InvalidEncoding)?,
    );

    require!(address <= U160_MAX, FeeQuoterError::InvalidEVMAddress);

    if let Ok(small_address) = TryInto::<u32>::try_into(address) {
        require_gte!(
            small_address,
            PRECOMPILE_SPACE,
            FeeQuoterError::InvalidEVMAddress
        )
    };

    Ok(())
}
