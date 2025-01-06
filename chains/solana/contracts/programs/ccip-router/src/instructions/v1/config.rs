use anchor_lang::prelude::*;

use crate::{
    AcceptOwnership, AddChainSelector, CcipRouterError, DestChainAdded, DestChainConfig,
    DestChainConfigUpdated, DestChainState, SourceChainAdded, SourceChainConfig,
    SourceChainConfigUpdated, SourceChainState, TimestampedPackedU224, TransferOwnership,
    UpdateConfigCCIPRouter, UpdateDestChainSelectorConfig, UpdateSourceChainSelectorConfig,
};

pub fn transfer_ownership(ctx: Context<TransferOwnership>, proposed_owner: Pubkey) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
    require!(
        proposed_owner != config.owner,
        CcipRouterError::InvalidInputs
    );
    config.proposed_owner = proposed_owner;
    Ok(())
}

pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;
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
    dest_chain_state.state = DestChainState {
        sequence_number: 0,
        usd_per_unit_gas: TimestampedPackedU224 {
            value: [0; 28],
            timestamp: 0,
        },
    };

    emit!(SourceChainAdded {
        source_chain_selector: new_chain_selector,
        source_chain_config,
    });
    emit!(DestChainAdded {
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

    emit!(SourceChainConfigUpdated {
        source_chain_selector,
        source_chain_config: chain_state.config.clone(),
    });

    Ok(())
}

pub fn disable_dest_chain_selector(
    ctx: Context<UpdateDestChainSelectorConfig>,
    dest_chain_selector: u64,
) -> Result<()> {
    let chain_state = &mut ctx.accounts.dest_chain_state;

    chain_state.config.is_enabled = false;

    emit!(DestChainConfigUpdated {
        dest_chain_selector,
        dest_chain_config: chain_state.config.clone(),
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

    emit!(SourceChainConfigUpdated {
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

    emit!(DestChainConfigUpdated {
        dest_chain_selector,
        dest_chain_config,
    });
    Ok(())
}

pub fn update_solana_chain_selector(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_chain_selector: u64,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.solana_chain_selector = new_chain_selector;

    Ok(())
}

pub fn update_default_gas_limit(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_gas_limit: u128,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    config.default_gas_limit = new_gas_limit;

    Ok(())
}

pub fn update_default_allow_out_of_order_execution(
    ctx: Context<UpdateConfigCCIPRouter>,
    new_allow_out_of_order_execution: bool,
) -> Result<()> {
    let mut config = ctx.accounts.config.load_mut()?;

    let mut v = 0_u8;
    if new_allow_out_of_order_execution {
        v = 1;
    }
    config.default_allow_out_of_order_execution = v;

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

fn validate_source_chain_config(
    _source_chain_selector: u64,
    _config: &SourceChainConfig,
) -> Result<()> {
    // As of now, the config has very few properties and there is nothing to validate yet.
    // This is a placeholder to add validations as that config object grows.
    Ok(())
}

fn validate_dest_chain_config(dest_chain_selector: u64, config: &DestChainConfig) -> Result<()> {
    // TODO improve errors
    require!(dest_chain_selector != 0, CcipRouterError::InvalidInputs);
    require!(
        config.default_tx_gas_limit != 0,
        CcipRouterError::InvalidInputs
    );
    require!(
        config.default_tx_gas_limit <= config.max_per_msg_gas_limit,
        CcipRouterError::InvalidInputs
    );
    require!(
        config.chain_family_selector != [0; 4],
        CcipRouterError::InvalidInputs
    );
    Ok(())
}
