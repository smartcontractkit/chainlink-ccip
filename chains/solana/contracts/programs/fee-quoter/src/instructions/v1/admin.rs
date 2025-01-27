use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::context::{
    AddBillingTokenConfig, RemoveBillingTokenConfig, UpdateBillingTokenConfig,
    FEE_BILLING_SIGNER_SEEDS,
};
use crate::event::{FeeTokenAdded, FeeTokenDisabled, FeeTokenEnabled, FeeTokenRemoved};
use crate::state::BillingTokenConfig;

pub fn add_billing_token_config(
    ctx: Context<AddBillingTokenConfig>,
    config: BillingTokenConfig,
) -> Result<()> {
    emit!(FeeTokenAdded {
        fee_token: config.mint,
        enabled: config.enabled
    });
    ctx.accounts.billing_token_config.version = 1; // update this if we change the account struct
    ctx.accounts.billing_token_config.config = config;
    Ok(())
}

pub fn update_billing_token_config(
    ctx: Context<UpdateBillingTokenConfig>,
    config: BillingTokenConfig,
) -> Result<()> {
    if config.enabled != ctx.accounts.billing_token_config.config.enabled {
        // enabled/disabled status has changed
        match config.enabled {
            true => emit!(FeeTokenEnabled {
                fee_token: config.mint
            }),
            false => emit!(FeeTokenDisabled {
                fee_token: config.mint
            }),
        }
    }
    // TODO should we emit an event if the config has changed regardless of the enabled/disabled?

    ctx.accounts.billing_token_config.version = 1; // update this if we change the account struct
    ctx.accounts.billing_token_config.config = config;
    Ok(())
}

pub fn remove_billing_token_config(ctx: Context<RemoveBillingTokenConfig>) -> Result<()> {
    // Close the receiver token account
    // The context constraints already enforce that it holds 0 balance of the target SPL token
    let cpi_accounts = token_interface::CloseAccount {
        account: ctx.accounts.fee_token_receiver.to_account_info(),
        destination: ctx.accounts.authority.to_account_info(),
        authority: ctx.accounts.fee_billing_signer.to_account_info(),
    };
    let cpi_program = ctx.accounts.token_program.to_account_info();
    let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[ctx.bumps.fee_billing_signer]];
    let signer_seeds = &[&seeds[..]];
    let cpi_ctx = CpiContext::new_with_signer(cpi_program, cpi_accounts, signer_seeds);

    token_interface::close_account(cpi_ctx)?;

    emit!(FeeTokenRemoved {
        fee_token: ctx.accounts.fee_token_mint.key()
    });
    Ok(())
}
