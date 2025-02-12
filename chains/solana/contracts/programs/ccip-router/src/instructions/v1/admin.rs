use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::events::admin as events;
use crate::state::RestoreOnAction;
use crate::{
    AcceptOwnership, AddChainSelector, CcipRouterError, DestChainConfig, DestChainState,
    TransferOwnership, UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig, WithdrawBilledFunds,
};
use crate::{AddOfframp, RemoveOfframp};

use super::fees::do_billing_transfer;

pub fn transfer_ownership(ctx: Context<TransferOwnership>, proposed_owner: Pubkey) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    require!(
        proposed_owner != config.owner,
        CcipRouterError::RedundantOwnerProposal
    );
    emit!(events::OwnershipTransferRequested {
        from: config.owner,
        to: proposed_owner,
    });
    config.proposed_owner = proposed_owner;
    Ok(())
}

pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    emit!(events::OwnershipTransferred {
        from: config.owner,
        to: config.proposed_owner,
    });
    config.owner = std::mem::take(&mut config.proposed_owner);
    config.proposed_owner = Pubkey::new_from_array([0; 32]);
    Ok(())
}

pub fn update_fee_aggregator(
    ctx: Context<UpdateConfigCCIPRouter>,
    fee_aggregator: Pubkey,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    config.fee_aggregator = fee_aggregator;
    Ok(())
}

pub fn add_chain_selector(
    ctx: Context<AddChainSelector>,
    new_chain_selector: u64,
    dest_chain_config: DestChainConfig,
) -> Result<()> {
    // Set dest chain config & state
    let dest_chain_state = &mut ctx.accounts.dest_chain_state;
    validate_dest_chain_config(new_chain_selector, &dest_chain_config)?;
    dest_chain_state.version = 1;
    dest_chain_state.chain_selector = new_chain_selector;
    dest_chain_state.config = dest_chain_config.clone();
    dest_chain_state.state = DestChainState {
        sequence_number: 0,
        restore_on_action: RestoreOnAction::None,
        sequence_number_to_restore: 0,
    };

    emit!(events::DestChainAdded {
        dest_chain_selector: new_chain_selector,
        dest_chain_config,
    });

    Ok(())
}

pub fn update_dest_chain_config(
    ctx: Context<UpdateDestChainSelectorConfig>,
    dest_chain_selector: u64,
    dest_chain_config: DestChainConfig,
) -> Result<()> {
    validate_dest_chain_config(dest_chain_selector, &dest_chain_config)?;

    ctx.accounts.dest_chain_state.config = dest_chain_config.clone();

    emit!(events::DestChainConfigUpdated {
        dest_chain_selector,
        dest_chain_config,
    });
    Ok(())
}

pub fn bump_ccip_version_for_dest_chain(
    ctx: Context<UpdateDestChainSelectorConfig>,
    dest_chain_selector: u64,
) -> Result<()> {
    let dest_chain_state = &mut ctx.accounts.dest_chain_state.state;

    emit!(events::CcipVersionForDestChainVersionBumped {
        dest_chain_selector,
        previous_sequence_number: dest_chain_state.sequence_number,
        new_sequence_number: 0,
    });

    let current_seq_nr = dest_chain_state.sequence_number;

    dest_chain_state.sequence_number = match dest_chain_state.restore_on_action {
        RestoreOnAction::Upgrade => dest_chain_state.sequence_number_to_restore,
        _ => 0,
    };
    dest_chain_state.sequence_number_to_restore = current_seq_nr;

    // restore on next rollback, as seq nr was of the previous CCIP version
    dest_chain_state.restore_on_action = RestoreOnAction::Rollback;

    Ok(())
}

pub fn rollback_ccip_version_for_dest_chain(
    ctx: Context<UpdateDestChainSelectorConfig>,
    dest_chain_selector: u64,
) -> Result<()> {
    let dest_chain_state = &mut ctx.accounts.dest_chain_state.state;

    // If there was loss of information, we can't rollback. We support at most 1 consecutive rollback.
    // So, once a rollback has happened, the admin must bump the CCIP version before another rollback.
    require_eq!(
        dest_chain_state.restore_on_action,
        RestoreOnAction::Rollback,
        CcipRouterError::InvalidCcipVersionRollback
    );

    emit!(events::CcipVersionForDestChainVersionRolledBack {
        dest_chain_selector,
        previous_sequence_number: dest_chain_state.sequence_number,
        new_sequence_number: dest_chain_state.sequence_number_to_restore,
    });

    let current_seq_nr = dest_chain_state.sequence_number;

    dest_chain_state.sequence_number = dest_chain_state.sequence_number_to_restore;
    dest_chain_state.sequence_number_to_restore = current_seq_nr;

    // restore on next upgrade, as seq nr was of the previously-bumped CCIP version
    dest_chain_state.restore_on_action = RestoreOnAction::Upgrade;

    Ok(())
}

pub fn add_offramp(
    _ctx: Context<AddOfframp>,
    source_chain_selector: u64,
    offramp: Pubkey,
) -> Result<()> {
    emit!(events::OfframpAdded {
        source_chain_selector,
        offramp,
    });
    Ok(())
}

pub fn remove_offramp(
    _ctx: Context<RemoveOfframp>,
    source_chain_selector: u64,
    offramp: Pubkey,
) -> Result<()> {
    emit!(events::OfframpRemoved {
        source_chain_selector,
        offramp,
    });
    Ok(())
}

pub fn update_svm_chain_selector(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_chain_selector: u64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.svm_chain_selector = new_chain_selector;

    Ok(())
}

pub fn withdraw_billed_funds(
    ctx: Context<WithdrawBilledFunds>,
    transfer_all: bool,
    desired_amount: u64, // if transfer_all is false, this value must be 0
) -> Result<()> {
    let transfer = token_interface::TransferChecked {
        from: ctx.accounts.fee_token_accum.to_account_info(),
        to: ctx.accounts.recipient.to_account_info(),
        mint: ctx.accounts.fee_token_mint.to_account_info(),
        authority: ctx.accounts.fee_billing_signer.to_account_info(),
    };

    let amount = if transfer_all {
        require!(
            desired_amount == 0,
            CcipRouterError::InvalidInputsTransferAllAmount
        );
        require!(
            ctx.accounts.fee_token_accum.amount > 0,
            CcipRouterError::InsufficientFunds
        );
        ctx.accounts.fee_token_accum.amount
    } else {
        require!(
            desired_amount > 0,
            CcipRouterError::InvalidInputsTokenAmount
        );
        require!(
            desired_amount <= ctx.accounts.fee_token_accum.amount,
            CcipRouterError::InsufficientFunds
        );
        desired_amount
    };

    do_billing_transfer(
        ctx.accounts.token_program.to_account_info(),
        transfer,
        amount,
        ctx.accounts.fee_token_mint.decimals,
        ctx.bumps.fee_billing_signer,
    )
}

/////////////
// Helpers //
/////////////

fn validate_dest_chain_config(dest_chain_selector: u64, _config: &DestChainConfig) -> Result<()> {
    // As of now, the config has very few properties and there is very little to validate yet.
    // This is mainly a placeholder to add validations as that config object grows.
    require!(
        dest_chain_selector != 0,
        CcipRouterError::InvalidInputsChainSelector
    );
    Ok(())
}
