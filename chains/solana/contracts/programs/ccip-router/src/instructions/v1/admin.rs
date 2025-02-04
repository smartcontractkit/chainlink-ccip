use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::events::admin as events;
use crate::{
    AcceptOwnership, AddChainSelector, CcipRouterError, DestChainConfig, DestChainState,
    Ocr3ConfigInfo, OcrPluginType, SetOcrConfig, SourceChainConfig, SourceChainState,
    TransferOwnership, UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig,
    UpdateSourceChainSelectorConfig, WithdrawBilledFunds,
};

use super::fees::do_billing_transfer;
use super::ocr3base::ocr3_set;

pub fn transfer_ownership(ctx: Context<TransferOwnership>, proposed_owner: Pubkey) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    require!(
        proposed_owner != config.owner,
        CcipRouterError::InvalidInputs
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
    source_chain_config: SourceChainConfig,
    dest_chain_config: DestChainConfig,
) -> Result<()> {
    // Set source chain config & state
    let source_chain_state = &mut ctx.accounts.source_chain_state;
    validate_source_chain_config(new_chain_selector, &source_chain_config)?;
    source_chain_state.version = 1;
    source_chain_state.chain_selector = new_chain_selector;
    source_chain_state.config = source_chain_config.clone();
    source_chain_state.state = SourceChainState { min_seq_nr: 1 };

    // Set dest chain config & state
    let dest_chain_state = &mut ctx.accounts.dest_chain_state;
    validate_dest_chain_config(new_chain_selector, &dest_chain_config)?;
    dest_chain_state.version = 1;
    dest_chain_state.chain_selector = new_chain_selector;
    dest_chain_state.config = dest_chain_config.clone();
    dest_chain_state.state = DestChainState { sequence_number: 0 };

    emit!(events::SourceChainAdded {
        source_chain_selector: new_chain_selector,
        source_chain_config,
    });
    emit!(events::DestChainAdded {
        dest_chain_selector: new_chain_selector,
        dest_chain_config,
    });

    Ok(())
}

pub fn disable_source_chain_selector(
    ctx: Context<UpdateSourceChainSelectorConfig>,
    source_chain_selector: u64,
) -> Result<()> {
    let chain_state = &mut ctx.accounts.source_chain_state;

    chain_state.config.is_enabled = false;

    emit!(events::SourceChainConfigUpdated {
        source_chain_selector,
        source_chain_config: chain_state.config.clone(),
    });

    Ok(())
}

pub fn update_source_chain_config(
    ctx: Context<UpdateSourceChainSelectorConfig>,
    source_chain_selector: u64,
    source_chain_config: SourceChainConfig,
) -> Result<()> {
    validate_source_chain_config(source_chain_selector, &source_chain_config)?;

    ctx.accounts.source_chain_state.config = source_chain_config.clone();

    emit!(events::SourceChainConfigUpdated {
        source_chain_selector,
        source_chain_config,
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

pub fn update_svm_chain_selector(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_chain_selector: u64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.svm_chain_selector = new_chain_selector;

    Ok(())
}

pub fn update_enable_manual_execution_after(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_enable_manual_execution_after: i64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.enable_manual_execution_after = new_enable_manual_execution_after;

    Ok(())
}

pub fn set_ocr_config(
    ctx: Context<SetOcrConfig>,
    plugin_type: u8, // OcrPluginType, u8 used because anchor tests did not work with an enum
    config_info: Ocr3ConfigInfo,
    signers: Vec<[u8; 20]>,
    transmitters: Vec<Pubkey>,
) -> Result<()> {
    require!(plugin_type < 2, CcipRouterError::InvalidInputs);
    let mut config = ctx.accounts.config.load_mut()?;

    let is_commit = plugin_type == OcrPluginType::Commit as u8;

    ocr3_set(
        &mut config.ocr3[plugin_type as usize],
        plugin_type,
        Ocr3ConfigInfo {
            config_digest: config_info.config_digest,
            f: config_info.f,
            n: signers.len() as u8,
            is_signature_verification_enabled: if is_commit { 1 } else { 0 },
        },
        signers,
        transmitters,
    )?;

    if is_commit {
        // When the OCR config changes, we reset the sequence number since it is scoped per config digest.
        // Note that s_minSeqNr/roots do not need to be reset as the roots persist
        // across reconfigurations and are de-duplicated separately.
        ctx.accounts.state.latest_price_sequence_number = 0;
    }

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
        require!(desired_amount == 0, CcipRouterError::InvalidInputs);
        require!(
            ctx.accounts.fee_token_accum.amount > 0,
            CcipRouterError::InsufficientFunds
        );
        ctx.accounts.fee_token_accum.amount
    } else {
        require!(desired_amount > 0, CcipRouterError::InvalidInputs);
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

fn validate_source_chain_config(
    _source_chain_selector: u64,
    _config: &SourceChainConfig,
) -> Result<()> {
    // As of now, the config has very few properties and there is nothing to validate yet.
    // This is a placeholder to add validations as that config object grows.
    Ok(())
}

fn validate_dest_chain_config(dest_chain_selector: u64, _config: &DestChainConfig) -> Result<()> {
    // As of now, the config has very few properties and there is very little to validate yet.
    // This is mainly a placeholder to add validations as that config object grows.
    // TODO improve errors
    require!(dest_chain_selector != 0, CcipRouterError::InvalidInputs);
    Ok(())
}
