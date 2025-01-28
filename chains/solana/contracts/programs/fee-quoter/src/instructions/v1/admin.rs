use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::context::{
    AcceptOwnership, AddBillingTokenConfig, AddDestChain, RemoveBillingTokenConfig,
    SetTokenBillingConfig, UpdateBillingTokenConfig, UpdateConfig, UpdateDestChainConfig,
    FEE_BILLING_SIGNER_SEEDS,
};
use crate::event::{
    DestChainAdded, DestChainConfigUpdated, FeeTokenAdded, FeeTokenDisabled, FeeTokenEnabled,
    FeeTokenRemoved, OwnershipTransferRequested, OwnershipTransferred,
    TokenTransferFeeConfigUpdated,
};
use crate::state::{
    BillingTokenConfig, DestChain, DestChainConfig, DestChainState, PerChainPerTokenConfig,
    TimestampedPackedU224, TokenBilling,
};
use crate::FeeQuoterError;

pub fn transfer_ownership(ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
    let config = &mut ctx.accounts.config;
    require!(
        proposed_owner != config.owner,
        FeeQuoterError::InvalidInputs
    );
    emit!(OwnershipTransferRequested {
        from: config.owner,
        to: proposed_owner,
    });
    ctx.accounts.config.proposed_owner = proposed_owner;
    Ok(())
}

pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
    let config = &mut ctx.accounts.config;
    emit!(OwnershipTransferred {
        from: config.owner,
        to: config.proposed_owner,
    });
    ctx.accounts.config.owner = ctx.accounts.config.proposed_owner;
    ctx.accounts.config.proposed_owner = Pubkey::default();
    Ok(())
}

// pub fn update_onramp(ctx: Context<UpdateConfig>, onramp: Pubkey) -> Result<()> {
//     ctx.accounts.config.onramp = onramp;
//     // TODO emit event
//     Ok(())
// }

// pub fn update_offramp(ctx: Context<UpdateConfig>, offramp: Pubkey) -> Result<()> {
//     ctx.accounts.config.offramp = offramp;
//     // TODO emit event
//     Ok(())
// }

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

pub fn add_dest_chain(
    ctx: Context<AddDestChain>,
    chain_selector: u64,
    dest_chain_config: DestChainConfig,
) -> Result<()> {
    validate_dest_chain_config(chain_selector, &dest_chain_config)?;
    ctx.accounts.dest_chain.set_inner(DestChain {
        version: 1,
        chain_selector,
        state: DestChainState {
            usd_per_unit_gas: TimestampedPackedU224 {
                timestamp: 0,
                value: [0; 28],
            },
        },
        config: dest_chain_config.clone(),
    });
    emit!(DestChainAdded {
        dest_chain_selector: chain_selector,
        dest_chain_config,
    });
    Ok(())
}

pub fn disable_dest_chain(ctx: Context<UpdateDestChainConfig>, chain_selector: u64) -> Result<()> {
    ctx.accounts.dest_chain.config.is_enabled = false;
    emit!(DestChainConfigUpdated {
        dest_chain_selector: chain_selector,
        dest_chain_config: ctx.accounts.dest_chain.config.clone(),
    });
    Ok(())
}

pub fn update_dest_chain_config(
    ctx: Context<UpdateDestChainConfig>,
    chain_selector: u64,
    dest_chain_config: DestChainConfig,
) -> Result<()> {
    validate_dest_chain_config(chain_selector, &dest_chain_config)?;
    ctx.accounts.dest_chain.config = dest_chain_config.clone();
    emit!(DestChainConfigUpdated {
        dest_chain_selector: chain_selector,
        dest_chain_config,
    });
    Ok(())
}

pub fn set_token_billing(
    ctx: Context<SetTokenBillingConfig>,
    chain_selector: u64,
    mint: Pubkey,
    cfg: TokenBilling,
) -> Result<()> {
    ctx.accounts
        .per_chain_per_token_config
        .set_inner(PerChainPerTokenConfig {
            version: 1, // update this if we change the account struct
            chain_selector,
            mint,
            billing: cfg.clone(),
        });
    emit!(TokenTransferFeeConfigUpdated {
        dest_chain_selector: chain_selector,
        token: mint,
        token_transfer_fee_config: cfg,
    });
    Ok(())
}

// --- helpers ---

fn validate_dest_chain_config(dest_chain_selector: u64, config: &DestChainConfig) -> Result<()> {
    // TODO improve errors
    require!(dest_chain_selector != 0, FeeQuoterError::InvalidInputs);
    require!(
        config.default_tx_gas_limit != 0,
        FeeQuoterError::InvalidInputs
    );
    require!(
        config.default_tx_gas_limit <= config.max_per_msg_gas_limit,
        FeeQuoterError::InvalidInputs
    );
    require!(
        config.chain_family_selector != [0; 4],
        FeeQuoterError::InvalidInputs
    );
    Ok(())
}
