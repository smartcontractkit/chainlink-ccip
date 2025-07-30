use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id,
    token_interface::{Mint, TokenAccount, TokenInterface},
};
use base_token_pool::common::*;
use ccip_common::seed;

use crate::{program::BurnmintTokenPool, ChainConfig, State};

const MAX_POOL_STATE_V: u8 = 1;
const MAX_POOL_CONFIG_V: u8 = 1;

#[derive(Accounts)]
pub struct InitGlobalConfig<'info> {
    #[account(
        init,
        seeds = [CONFIG_SEED],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PoolConfig::INIT_SPACE,
    )]
    pub config: Account<'info, PoolConfig>, // Global Config PDA of the Token Pool

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    // Ensures that the provided program is the BurnmintTokenPool program,
    // and that its associated program data account matches the expected one.
    // This guarantees that only the program's upgrade authority can modify the global config.
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, BurnmintTokenPool>,
    // Global Config updates only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct UpdateGlobalConfig<'info> {
    #[account(
        mut,
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_POOL_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub config: Account<'info, PoolConfig>, // Global Config PDA of the Token Pool

    pub authority: Signer<'info>,

    // Ensures that the provided program is the BurnmintTokenPool program,
    // and that its associated program data account matches the expected one.
    // This guarantees that only the program's upgrade authority can modify the global config.
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, BurnmintTokenPool>,
    // Global Config updates only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + State::INIT_SPACE,
    )]
    pub state: Account<'info, State>, // config PDA for token pool
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, BurnmintTokenPool>,

    // The upgrade authority of the token pool program can initialize a token pool or the mint authority of the token
    #[account(constraint = allowed_to_initialize_token_pool(&program_data, &authority, &config, &mint) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,

    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.version, MAX_POOL_CONFIG_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub config: Account<'info, PoolConfig>, // Global Config PDA of the Token Pool
}

#[derive(Accounts)]
pub struct AdminUpdateTokenPool<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>, // config PDA for token pool
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps

    #[account(mut)]
    pub authority: Signer<'info>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, BurnmintTokenPool>,

    // The upgrade authority of the token pool program can update certain values of their token pools only (router and rmn addresses for testing)
    #[account(constraint = allowed_admin_modify_token_pool(&program_data, &authority, &state) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
pub struct TransferMintAuthority<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(address = *mint.to_account_info().owner)]
    pub token_program: Interface<'info, TokenInterface>,
    #[account(
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
        address = state.config.pool_signer,
    )]
    /// CHECK: unchecked CPI signer
    pub pool_signer: UncheckedAccount<'info>,
    pub authority: Signer<'info>,

    /// CHECK: The Multisig can be a Token 2022 Multisig or a SPL Token Multisig and there is no interface for it.
    #[account(owner = *mint.to_account_info().owner @ CcipBnMTokenPoolError::InvalidMultisigOwner)]
    pub new_multisig_mint_authority: UncheckedAccount<'info>, // new mint authority for the underlying token

    // Ensures that the provided program is the BurnmintTokenPool program,
    // and that its associated program data account matches the expected one.
    // This guarantees that only the program's upgrade authority can modify the global config.
    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, BurnmintTokenPool>,
    // Transfer mint authority only allowed by program upgrade authority as it is a critical operation.
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipTokenPoolError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct InitializeStateVersion<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.as_ref()],
        bump,
        constraint = uninitialized(state.version) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
}

#[derive(Accounts)]
pub struct Empty<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(add: Vec<Pubkey>)]
pub struct AddToAllowList<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + State::INIT_SPACE + 32 * (state.config.allow_list.len() + add.len()),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remove: Vec<Pubkey>)]
pub struct RemoveFromAllowlist<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + State::INIT_SPACE + 32 * (state.config.allow_list.len().saturating_sub(remove.len())),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(address = state.config.proposed_owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct TokenOfframp<'info> {
    // CCIP accounts ------------------------
    #[account(
        seeds = [EXTERNAL_TOKEN_POOLS_SIGNER, crate::ID.as_ref()],
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
            offramp_program.key().as_ref()
        ],
        bump,
        seeds::program = state.config.router,
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(address = *mint.to_account_info().owner)]
    pub token_program: Interface<'info, TokenInterface>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
        address = state.config.pool_signer,
    )]
    /// CHECK: unchecked CPI signer
    pub pool_signer: UncheckedAccount<'info>,
    #[account(mut, address = get_associated_token_address_with_program_id(&pool_signer.key(), &mint.key(), &token_program.key()))]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, release_or_mint.remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,

    // User specific accounts ---------------
    #[account(mut, address = get_associated_token_address_with_program_id(&release_or_mint.receiver, &mint.key(), &token_program.key()))]
    pub receiver_token_account: InterfaceAccount<'info, TokenAccount>,
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
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(address = *mint.to_account_info().owner)]
    pub token_program: Interface<'info, TokenInterface>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [POOL_SIGNER_SEED, mint.key().as_ref()],
        bump,
        address = state.config.pool_signer,
    )]
    /// CHECK: unchecked CPI signer
    pub pool_signer: UncheckedAccount<'info>,
    #[account(mut, address = get_associated_token_address_with_program_id(&pool_signer.key(), &mint.key(), &token_program.key()))]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    ////////////////////
    // RMN Remote CPI //
    ////////////////////
    /// CHECK: This is the account for the RMN Remote program
    #[account(
        address = state.config.rmn_remote @ CcipTokenPoolError::InvalidRMNRemoteAddress,
    )]
    pub rmn_remote: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CURSES],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_curses: UncheckedAccount<'info>,

    /// CHECK: This account is just used in the CPI to the RMN Remote program
    #[account(
        seeds = [seed::CONFIG],
        bump,
        seeds::program = state.config.rmn_remote,
    )]
    pub rmn_remote_config: UncheckedAccount<'info>,

    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, lock_or_burn.remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct InitializeChainConfig<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(
        init,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
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
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, constraint = authority.key() == state.config.owner || authority.key() == state.config.rate_limit_admin @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct SetRateLimitAdmin<'info> {
    #[account(
        mut,
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(mut, address = state.config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, cfg: RemoteConfig)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
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
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
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
        seeds = [POOL_STATE_SEED, mint.key().as_ref()],
        bump,
        constraint = valid_version(state.version, MAX_POOL_STATE_V) @ CcipTokenPoolError::InvalidVersion,
    )]
    pub state: Account<'info, State>,
    #[account(
        mut,
        seeds = [POOL_CHAINCONFIG_SEED, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = state.config.owner)]
    pub authority: Signer<'info>,
}

#[error_code]
pub enum CcipBnMTokenPoolError {
    #[msg("Invalid Multisig Mint")]
    InvalidMultisig,
    #[msg("Mint Authority already set")]
    MintAuthorityAlreadySet,
    #[msg("Token with no Mint Authority")]
    FixedMintToken,
    #[msg("Unsupported Token Program")]
    UnsupportedTokenProgram,
    #[msg("Invalid Multisig Account Data for Token 2022")]
    InvalidToken2022Multisig,
    #[msg("Invalid Multisig Account Data for SPL Token")]
    InvalidSPLTokenMultisig,
    #[msg("Token Pool Signer PDA must be m times a signer of the Multisig")]
    PoolSignerNotInMultisig,
    #[msg("Multisig must have more than 2 valid signers")]
    MultisigMustHaveAtLeastTwoSigners,
    #[msg("Multisig must have more than one required signer")]
    MultisigMustHaveMoreThanOneSigner,
    #[msg("Multisig Owner must match Token Program ID")]
    InvalidMultisigOwner,
    #[msg("Invalid multisig threshold: required signatures cannot exceed total signers")]
    InvalidMultisigThreshold,
    #[msg(
        "Invalid multisig m: required signatures cannot exceed the available for outside signers"
    )]
    InvalidMultisigThresholdTooHigh,
}

// This account can not be declared in the common crate, the program ID for that Account would be incorrect.
#[account]
#[derive(InitSpace)]
pub struct PoolConfig {
    pub version: u8,
    pub self_served_allowed: bool,

    // Default adress values for the token pool
    pub router: Pubkey,
    pub rmn_remote: Pubkey,
}

/// Checks if the given authority is allowed to initialize the token pool.
pub fn allowed_to_initialize_token_pool(
    program_data: &Account<ProgramData>,
    authority: &Signer,
    config: &Account<PoolConfig>,
    mint: &InterfaceAccount<Mint>,
) -> bool {
    program_data.upgrade_authority_address == Some(authority.key()) || // The upgrade authority of the token pool program can initialize a token pool
    (config.self_served_allowed && Some(authority.key()) == mint.mint_authority.into() )
    // or the mint authority of the token
}

/// Checks that the authority and the token pool owner are the upgrade authority
pub fn allowed_admin_modify_token_pool(
    program_data: &Account<ProgramData>,
    authority: &Signer,
    state: &Account<State>,
) -> bool {
    program_data.upgrade_authority_address == Some(authority.key()) && // Only the upgrade authority of the token pool program can modify certain values of a given token pool
    state.config.owner == authority.key() // if only if the token pool is owned by the upgrade authority
}
