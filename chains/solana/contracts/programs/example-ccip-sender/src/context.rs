use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};

use crate::{BaseState, CcipSenderError, RemoteChainConfig, CCIP_SENDER, CHAIN_CONFIG_SEED};

const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts, Debug)]
pub struct Initialize<'info> {
    #[account(
        init,
        seeds = [b"state"],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + BaseState::INIT_SPACE,
    )]
    pub state: Account<'info, BaseState>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64)]
pub struct CcipSend<'info> {
    #[account(
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        seeds = [CCIP_SENDER],
        bump,
    )]
    #[account(mut)]
    /// CHECK: sender CPI caller
    pub ccip_sender: UncheckedAccount<'info>,
    #[account(
        mut,
        associated_token::authority = authority.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub authority_fee_token_ata: InterfaceAccount<'info, TokenAccount>,
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    // ------------------------
    // required ccip accounts - accounts are checked by the router
    #[account(
        address = state.router @ CcipSenderError::InvalidRouter,
    )]
    /// CHECK: used as router entry point for CPI
    pub ccip_router: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_config: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: validated during CPI
    pub ccip_dest_chain_state: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: validated during CPI
    pub ccip_sender_nonce: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_token_program: UncheckedAccount<'info>,
    #[account(
        owner = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_mint: InterfaceAccount<'info, Mint>,
    #[account(
        mut,
        associated_token::authority = ccip_sender.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_user_ata: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        associated_token::authority = ccip_fee_billing_signer.key(),
        associated_token::mint = ccip_fee_token_mint.key(),
        associated_token::token_program = ccip_fee_token_program.key(),
    )]
    pub ccip_fee_token_receiver: InterfaceAccount<'info, TokenAccount>,
    /// CHECK: validated during CPI
    pub ccip_fee_billing_signer: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_dest_chain: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_billing_token_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_fee_quoter_link_token_config: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_rmn_remote: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_rmn_remote_curses: UncheckedAccount<'info>,
    /// CHECK: validated during CPI
    pub ccip_rmn_remote_config: UncheckedAccount<'info>,
}

#[derive(Accounts, Debug)]
pub struct UpdateConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, _recipient: Vec<u8>, extra_args_bytes: Vec<u8>)]
pub struct InitChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        init,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        space = ANCHOR_DISCRIMINATOR + RemoteChainConfig::INIT_SPACE + extra_args_bytes.len(),
        payer = authority,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64, _recipient: Vec<u8>, extra_args_bytes: Vec<u8>)]
pub struct UpdateChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + RemoteChainConfig::INIT_SPACE + extra_args_bytes.len(),
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
#[instruction(chain_selector: u64)]
pub struct RemoveChainConfig<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        seeds = [CHAIN_CONFIG_SEED, chain_selector.to_le_bytes().as_ref()],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, RemoteChainConfig>,
    #[account(
        mut,
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts, Debug)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        address = state.proposed_owner @ CcipSenderError::OnlyProposedOwner,
    )]
    pub authority: Signer<'info>,
}

#[derive(Accounts, Debug)]
pub struct WithdrawTokens<'info> {
    #[account(
        mut,
        seeds = [b"state"],
        bump,
    )]
    pub state: Account<'info, BaseState>,
    #[account(
        mut,
        token::mint = mint,
        token::authority = ccip_sender,
        token::token_program = token_program,
    )]
    pub program_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = mint,
        token::token_program = token_program,
    )]
    pub to_token_account: InterfaceAccount<'info, TokenAccount>,
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(address = *mint.to_account_info().owner)]
    /// CHECK: CPI to token program
    pub token_program: AccountInfo<'info>,
    #[account(
        seeds = [CCIP_SENDER],
        bump,
    )]
    /// CHECK: CPI signer for tokens
    pub ccip_sender: UncheckedAccount<'info>,
    #[account(
        address = state.owner @ CcipSenderError::OnlyOwner,
    )]
    pub authority: Signer<'info>,
}
