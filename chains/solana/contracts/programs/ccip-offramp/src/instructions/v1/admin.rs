use anchor_lang::prelude::*;

use super::ocr3base::ocr3_set;

use crate::context::{
    AcceptOwnership, AddSourceChain, OcrPluginType, SetOcrConfig, TransferOwnership, UpdateConfig,
    UpdateSourceChain,
};
use crate::event::admin::{
    OwnershipTransferRequested, OwnershipTransferred, SourceChainAdded, SourceChainConfigUpdated,
};
use crate::state::{
    Ocr3ConfigInfo, Ocr3ConfigInfoInput, SourceChain, SourceChainConfig, SourceChainState,
};
use crate::CcipOfframpError;

pub fn transfer_ownership(ctx: Context<TransferOwnership>, proposed_owner: Pubkey) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    require!(
        proposed_owner != config.owner,
        CcipOfframpError::InvalidInputs
    );
    emit!(OwnershipTransferRequested {
        from: config.owner,
        to: proposed_owner,
    });
    config.proposed_owner = proposed_owner;
    Ok(())
}

pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    emit!(OwnershipTransferred {
        from: config.owner,
        to: config.proposed_owner,
    });
    config.owner = std::mem::take(&mut config.proposed_owner);
    config.proposed_owner = Pubkey::new_from_array([0; 32]);
    Ok(())
}

pub fn add_source_chain(
    ctx: Context<AddSourceChain>,
    new_chain_selector: u64,
    source_chain_config: SourceChainConfig,
) -> Result<()> {
    // Set source chain config & state
    let source_chain = &mut ctx.accounts.source_chain;
    validate_source_chain_config(new_chain_selector, &source_chain_config)?;
    source_chain.set_inner(SourceChain {
        version: 1,
        chain_selector: new_chain_selector,
        state: SourceChainState { min_seq_nr: 1 },
        config: source_chain_config.clone(),
    });

    emit!(SourceChainAdded {
        source_chain_selector: new_chain_selector,
        source_chain_config,
    });

    Ok(())
}

pub fn disable_source_chain_selector(
    ctx: Context<UpdateSourceChain>,
    source_chain_selector: u64,
) -> Result<()> {
    let source_chain = &mut ctx.accounts.source_chain;

    source_chain.config.is_enabled = false;

    emit!(SourceChainConfigUpdated {
        source_chain_selector,
        source_chain_config: source_chain.config.clone(),
    });

    Ok(())
}

pub fn update_source_chain_config(
    ctx: Context<UpdateSourceChain>,
    source_chain_selector: u64,
    source_chain_config: SourceChainConfig,
) -> Result<()> {
    validate_source_chain_config(source_chain_selector, &source_chain_config)?;

    ctx.accounts.source_chain.config = source_chain_config.clone();

    emit!(SourceChainConfigUpdated {
        source_chain_selector,
        source_chain_config,
    });
    Ok(())
}

pub fn update_svm_chain_selector(
    ctx: Context<UpdateConfig>,
    new_chain_selector: u64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.svm_chain_selector = new_chain_selector;

    Ok(())
}

pub fn update_enable_manual_execution_after(
    ctx: Context<UpdateConfig>,
    new_enable_manual_execution_after: i64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.enable_manual_execution_after = new_enable_manual_execution_after;

    Ok(())
}

pub fn set_ocr_config(
    ctx: Context<SetOcrConfig>,
    plugin_type: u8, // OcrPluginType, u8 used because anchor tests did not work with an enum
    config_info: Ocr3ConfigInfoInput,
    signers: Vec<[u8; 20]>,
    transmitters: Vec<Pubkey>,
) -> Result<()> {
    require!(plugin_type < 2, CcipOfframpError::InvalidInputs);
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

fn validate_source_chain_config(
    _source_chain_selector: u64,
    _config: &SourceChainConfig,
) -> Result<()> {
    // As of now, the config has very few properties and there is nothing to validate yet.
    // This is a placeholder to add validations as that config object grows.
    Ok(())
}
