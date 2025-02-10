use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id,
    token_interface::{Mint, TokenAccount},
};
use base_token_pool::common::*;

use crate::{ChainConfig, State};

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        seeds = [POOL_STATE_SEED, &mint.key().to_bytes()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + State::INIT_SPACE,
    )]
    pub state: Account<'info, State>, // config PDA for token pool
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(mut)]
    pub authority: Signer<'info>, // anyone can init token pool for a token, but ccip token admin registry controlls who can register pool for token
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(address = state.config.proposed_owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct TokenOfframp<'info> {
    // CCIP accounts ------------------------
    #[account(
        seeds = [EXTERNAL_TOKENPOOL_SIGNER],
        bump,
        seeds::program = offramp_program.key(),
    )]
    pub authority: Signer<'info>,

    /// CHECK offramp program: exists only to derive the allowed offramp PDA
    /// and the authority PDA.
    pub offramp_program: UncheckedAccount<'info>,

    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp
    #[account(
        owner = state.config.router @ CcipTokenPoolError::InvalidPoolCaller, // this guarantees that it was initialized
        seeds = [
            ALLOWED_OFFRAMP,
            release_or_mint.remote_chain_selector.to_le_bytes().as_ref(),
            &offramp_program.key().to_bytes()
        ],
        bump,
        seeds::program = state.config.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,
    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    /// CHECK: CPI to token program
    pub token_program: AccountInfo<'info>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [POOL_SIGNER_SEED, &state.config.mint.key().to_bytes()],
        bump,
        address = state.config.pool_signer,
    )]
    /// CHECK: unchecked CPI signer
    pub pool_signer: UncheckedAccount<'info>,
    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, release_or_mint.remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    // User specific accounts ---------------
    #[account(mut)]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,
    // remaining accounts -----------------
    // LockAndRelease: []
    // BurnAndMint: []
    // Wrapped: [wrapped program, ..remaining_accounts]
}

#[derive(Accounts)]
#[instruction(lock_or_burn: LockOrBurnInV1)]
pub struct TokenOnramp<'info> {
    // CCIP accounts ------------------------
    #[account(address = state.config.router_onramp_authority @ CcipTokenPoolError::InvalidPoolCaller)]
    pub authority: Signer<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    /// CHECK: CPI to underlying token program
    pub token_program: AccountInfo<'info>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [POOL_SIGNER_SEED, &state.config.mint.key().to_bytes()],
        bump,
        address = state.config.pool_signer,
    )]
    /// CHECK: unchecked CPI signer
    pub pool_signer: UncheckedAccount<'info>,
    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, lock_or_burn.remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    // remaining accounts -----------------
    // LockAndRelease: []
    // BurnAndMint: []
    // Wrapped: [wrapped program, ..remaining_accounts]
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct InitializeChainConfig<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(
        init,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct SetChainRateLimit<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, constraint = authority.key() == state.config.owner || authority.key() == state.config.rate_limit_admin)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, cfg: RemoteConfig)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE + cfg.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, addresses: Vec<RemoteAddress>)]
pub struct AppendRemotePoolAddresses<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE
            + chain_config.base.remote.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>()
            + addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct DeleteChainConfig<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, &state.config.mint.key().to_bytes()],
        bump,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), &mint.key().to_bytes()],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}
