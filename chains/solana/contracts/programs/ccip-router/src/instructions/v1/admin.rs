use anchor_lang::prelude::*;
use anchor_spl::token_interface;

use crate::{
    AcceptOwnership, AddBillingTokenConfig, AddChainSelector, BillingTokenConfig, CcipRouterError,
    DestChainAdded, DestChainConfig, DestChainConfigUpdated, DestChainState, FeeTokenAdded,
    FeeTokenDisabled, FeeTokenEnabled, FeeTokenRemoved, Ocr3ConfigInfo, OcrPluginType,
    RemoveBillingTokenConfig, SetOcrConfig, SetTokenBillingConfig, SourceChainAdded,
    SourceChainConfig, SourceChainConfigUpdated, SourceChainState, TimestampedPackedU224,
    TokenBilling, TransferOwnership, UpdateBillingTokenConfig, UpdateConfigCCIPRouter,
    UpdateDestChainSelectorConfig, UpdateSourceChainSelectorConfig, WithdrawBilledFunds,
    FEE_BILLING_SIGNER_SEEDS,
};

use super::fee_quoter::do_billing_transfer;

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

pub fn set_token_billing(
    ctx: Context<SetTokenBillingConfig>,
    chain_selector: u64,
    mint: Pubkey,
    cfg: TokenBilling,
) -> Result<()> {
    ctx.accounts.per_chain_per_token_config.version = 1; // update this if we change the account struct
    ctx.accounts.per_chain_per_token_config.billing = cfg;
    ctx.accounts.per_chain_per_token_config.chain_selector = chain_selector;
    ctx.accounts.per_chain_per_token_config.mint = mint;
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

    config.ocr3[plugin_type as usize].set(
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
    config: &SourceChainConfig,
) -> Result<()> {
    require!(config.is_valid_on_ramp(), CcipRouterError::InvalidInputs);
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
