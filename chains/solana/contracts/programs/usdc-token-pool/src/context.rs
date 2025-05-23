use crate::MESSAGE_TRANSMITTER;
use crate::TOKEN_MESSENGER_MINTER;
use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenAccount};
use base_token_pool::common::*;
use burnmint_token_pool::ChainConfig;

const MAX_POOL_STATE_V: u8 = 1;
const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        space = ANCHOR_DISCRIMINATOR + burnmint_token_pool::State::INIT_SPACE,
        payer = authority,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = valid_version(1, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, crate::program::UsdcTokenPool>,

    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()))]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.proposed_owner == authority.key() @ CcipTokenPoolError::Unauthorized
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct TokenOfframp<'info> {
    #[account(constraint = authority.key() == offramp_program.key() || authority.key() == allowed_offramp.key() @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,

    #[account(constraint = offramp_program.key() != Pubkey::default() @ CcipTokenPoolError::InvalidInputs)]
    pub offramp_program: UncheckedAccount<'info>,

    #[account(
        seeds = [
            ALLOWED_OFFRAMP,
            state.config.router.as_ref(),
            offramp_program.key().as_ref(),
        ],
        bump,
        seeds::program = state.config.router
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.router != Pubkey::default() @ CcipTokenPoolError::InvalidInputs,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub token_program: Program<'info, anchor_spl::token::Token>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(
        seeds = [
            POOL_SIGNER_SEED,
            mint.key().as_ref()
        ],
        bump,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    #[account()]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    #[account(
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &release_or_mint.remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(
        constraint = rmn_remote.key() == state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    pub rmn_remote_curses: UncheckedAccount<'info>,

    pub rmn_remote_config: UncheckedAccount<'info>,

    #[account(
        mut,
        constraint = receiver_token_account.mint == state.config.mint @ CcipTokenPoolError::InvalidInputs
    )]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,

    #[account(constraint = message_transmitter.key() == MESSAGE_TRANSMITTER @ crate::UsdcTokenPoolError::InvalidMessageTransmitter)]
    pub message_transmitter: UncheckedAccount<'info>,

    #[account(mut)]
    pub used_nonces: UncheckedAccount<'info>,

    pub receiver: UncheckedAccount<'info>,

    #[account(mut)]
    pub cctp_payer: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(lock_or_burn: LockOrBurnInV1)]
pub struct TokenOnramp<'info> {
    pub authority: Signer<'info>,

    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.router == authority.key() || state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub token_program: Program<'info, anchor_spl::token::Token>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(
        seeds = [
            POOL_SIGNER_SEED,
            mint.key().as_ref()
        ],
        bump,
    )]
    pub pool_signer: UncheckedAccount<'info>,

    #[account(mut)]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,

    #[account(
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &lock_or_burn.remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(
        constraint = rmn_remote.key() == state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    pub rmn_remote_curses: UncheckedAccount<'info>,

    pub rmn_remote_config: UncheckedAccount<'info>,

    #[account(constraint = token_messenger.key() == TOKEN_MESSENGER_MINTER @ crate::UsdcTokenPoolError::InvalidTokenMessenger)]
    pub token_messenger: UncheckedAccount<'info>,

    #[account(constraint = message_transmitter.key() == MESSAGE_TRANSMITTER @ crate::UsdcTokenPoolError::InvalidMessageTransmitter)]
    pub message_transmitter: UncheckedAccount<'info>,

    #[account(mut)]
    pub cctp_event_rent_payer: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct InitializeChainConfig<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        init,
        payer = authority,
        space = 8 + ChainConfig::INIT_SPACE,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
        realloc = 8 + ChainConfig::INIT_SPACE,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, addresses: Vec<RemoteAddress>)]
pub struct AppendRemotePoolAddresses<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
        realloc = 8 + ChainConfig::INIT_SPACE + addresses.len() * RemoteAddress::INIT_SPACE,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct SetChainRateLimit<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct DeleteChainConfig<'info> {
    #[account(
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    #[account(
        mut,
        close = authority,
        seeds = [
            POOL_CHAINCONFIG_SEED,
            state.key().as_ref(),
            &remote_chain_selector.to_le_bytes(),
        ],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(add: Vec<Pubkey>)]
pub struct AddToAllowList<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
        realloc = 8 + burnmint_token_pool::State::INIT_SPACE + add.len() * 32,
        realloc::payer = authority,
        realloc::zero = true,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remove: Vec<Pubkey>)]
pub struct RemoveFromAllowlist<'info> {
    #[account(
        mut,
        seeds = [
            POOL_STATE_SEED,
            mint.key().as_ref()
        ],
        bump,
        constraint = state.config.owner == authority.key() @ CcipTokenPoolError::Unauthorized,
    )]
    pub state: Account<'info, burnmint_token_pool::State>,

    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct Empty<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}
