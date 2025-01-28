use anchor_lang::prelude::*;

use crate::context::seed::{DEST_CHAIN, FEE_BILLING_TOKEN_CONFIG};
use crate::context::{GasPriceUpdate, TokenPriceUpdate, UpdatePrices};
use crate::event::{UsdPerTokenUpdated, UsdPerUnitGasUpdated};
use crate::state::{BillingTokenConfigWrapper, DestChain, TimestampedPackedU224};
use crate::FeeQuoterError;

pub fn update_prices<'info>(
    ctx: Context<'_, '_, 'info, 'info, UpdatePrices<'info>>,
    token_updates: Vec<TokenPriceUpdate>,
    gas_updates: Vec<GasPriceUpdate>,
) -> Result<()> {
    require!(
        !token_updates.is_empty() || !gas_updates.is_empty(),
        FeeQuoterError::InvalidInputs
    );

    // Remaining accounts represent:
    // - the accounts to update BillingTokenConfig for token prices
    // - the accounts to update DestChain for gas prices
    // They must be in order:
    // 1. token_accounts[]
    // 2. gas_accounts[]
    // matching the order of the price updates.
    // They must also all be writable so they can be updated.
    require_eq!(
        ctx.remaining_accounts.len(),
        token_updates.len() + gas_updates.len(),
        FeeQuoterError::InvalidInputs
    );

    // For each token price update, unpack the corresponding remaining_account and update the price.
    // Keep in mind that the remaining_accounts are sorted in the same order as tokens and gas price updates in the report.
    for (i, update) in token_updates.iter().enumerate() {
        apply_token_price_update(update, &ctx.remaining_accounts[i])?;
    }

    // Skip the first state account and the ones for token updates
    let offset = token_updates.len();

    // Do the same for gas price updates
    for (i, update) in gas_updates.iter().enumerate() {
        apply_gas_price_update(update, &ctx.remaining_accounts[i + offset])?;
    }

    Ok(())
}

/////////////
// Helpers //
/////////////

fn apply_token_price_update<'info>(
    token_update: &TokenPriceUpdate,
    token_config_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[FEE_BILLING_TOKEN_CONFIG, token_update.source_token.as_ref()],
        &crate::ID,
    );
    require_keys_eq!(
        token_config_account_info.key(),
        expected,
        FeeQuoterError::InvalidInputs
    );

    require!(
        token_config_account_info.is_writable,
        FeeQuoterError::InvalidInputs
    );

    // As the account is sent as remaining accounts, then Anchor won't automatically (de)serialize the account
    // as it is not the one in the context, so we have to do it manually load it and write it back
    let token_config_account = &mut Account::try_from(token_config_account_info)?;
    update_billing_token_config_price(token_config_account, token_update)?;
    token_config_account.exit(&crate::ID)
}

fn update_billing_token_config_price(
    token_config_account: &mut Account<BillingTokenConfigWrapper>,
    token_update: &TokenPriceUpdate,
) -> Result<()> {
    require!(
        token_config_account.version == 1,
        FeeQuoterError::InvalidInputs
    );
    token_config_account.config.usd_per_token = TimestampedPackedU224 {
        value: token_update.usd_per_token,
        timestamp: Clock::get()?.unix_timestamp,
    };
    emit!(UsdPerTokenUpdated {
        token: token_config_account.config.mint,
        value: token_config_account.config.usd_per_token.value,
        timestamp: token_config_account.config.usd_per_token.timestamp,
    });
    Ok(())
}

fn apply_gas_price_update<'info>(
    gas_update: &GasPriceUpdate,
    dest_chain_state_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[
            DEST_CHAIN,
            gas_update.dest_chain_selector.to_le_bytes().as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(
        dest_chain_state_account_info.key(),
        expected,
        FeeQuoterError::InvalidInputs
    );

    require!(
        dest_chain_state_account_info.is_writable,
        FeeQuoterError::InvalidInputs
    );

    // As the account is sent as remaining accounts, then Anchor won't automatically (de)serialize the account
    // as it is not the one in the context, so we have to do it manually load it and write it back
    let dest_chain_state_account = &mut Account::try_from(dest_chain_state_account_info)?;
    update_chain_state_gas_price(dest_chain_state_account, gas_update)?;
    dest_chain_state_account.exit(&crate::ID)
}

fn update_chain_state_gas_price(
    chain_state_account: &mut Account<DestChain>,
    gas_update: &GasPriceUpdate,
) -> Result<()> {
    require!(
        chain_state_account.version == 1,
        FeeQuoterError::InvalidInputs
    );

    chain_state_account.state.usd_per_unit_gas = TimestampedPackedU224 {
        value: gas_update.usd_per_unit_gas,
        timestamp: Clock::get()?.unix_timestamp,
    };

    emit!(UsdPerUnitGasUpdated {
        dest_chain: gas_update.dest_chain_selector,
        value: chain_state_account.state.usd_per_unit_gas.value,
        timestamp: chain_state_account.state.usd_per_unit_gas.timestamp,
    });

    Ok(())
}
