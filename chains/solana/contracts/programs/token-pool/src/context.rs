use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id,
    token_interface::{Mint, TokenAccount},
};

use crate::{
    ChainConfig, Config, ExternalExecutionConfig, LockOrBurnInV1, RateLimitConfig,
    ReleaseOrMintInV1, RemoteAddress, RemoteConfig,
};

const ANCHOR_DISCRIMINATOR: usize = 8; // 8-byte anchor discriminator length
const CCIP_TOKENPOOL_CONFIG: &[u8] = b"ccip_tokenpool_config";
pub const CCIP_TOKENPOOL_SIGNER: &[u8] = b"ccip_tokenpool_signer";
pub const CCIP_TOKENPOOL_CHAINCONFIG: &[u8] = b"ccip_tokenpool_chainconfig";
pub const EXTERNAL_TOKENPOOL_SIGNER: &[u8] = b"external_token_pools_signer";
pub const RELEASE_MINT: [u8; 8] = [0x14, 0x94, 0x71, 0xc6, 0xe5, 0xaa, 0x47, 0x30];
pub const LOCK_BURN: [u8; 8] = [0xc8, 0x0e, 0x32, 0x09, 0x2c, 0x5b, 0x79, 0x25];
pub const ALLOWED_OFFRAMP: &[u8] = b"allowed_offramp";

#[derive(Accounts)]
pub struct InitializeTokenPool<'info> {
    #[account(
        init,
        seeds = [CCIP_TOKENPOOL_CONFIG, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: Account<'info, Config>, // config PDA for token pool
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(
        init,
        seeds = [CCIP_TOKENPOOL_SIGNER, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ExternalExecutionConfig::INIT_SPACE,
    )]
    pub pool_signer: Account<'info, ExternalExecutionConfig>, // PDA for managing tokens
    #[account(
        mut,
        address = mint.mint_authority.unwrap() @ CcipTokenPoolError::InvalidInitPoolPermissions, // TODO - what if mint_authority is 0-address (no more minting)
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct SetConfig<'info> {
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(address = config.owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(address = config.proposed_owner @ CcipTokenPoolError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(release_or_mint: ReleaseOrMintInV1)]
pub struct TokenOfframp<'info> {
    // CCIP accounts ------------------------
    pub authority: Signer<'info>,

    /// CHECK offramp program: exists only to derive the allowed offramp PDA
    /// and the signer PDA. This constraint verifies that the authority is
    /// correctly derived from the offramp program.
    #[account(
        constraint = {
            let (pda, _) = Pubkey::find_program_address(
                &[
                    EXTERNAL_TOKENPOOL_SIGNER,
                ],
                &offramp_program.key()
            );
            pda == authority.key() && authority.owner == &offramp_program.key() } @ CcipTokenPoolError::InvalidPoolCaller
    )]
    pub offramp_program: UncheckedAccount<'info>,

    /// CHECK PDA of the router program verifying the signer is an allowed offramp.
    /// If PDA does not exist, the router doesn't allow this offramp
    #[account(
        constraint = {
        let (pda, _) = Pubkey::find_program_address(
            &[
                ALLOWED_OFFRAMP,
                release_or_mint.remote_chain_selector.to_le_bytes().as_ref(),
                offramp_program.key().as_ref(),
            ],
            &config.ccip_router,
        );
        allowed_offramp.key() == pda && allowed_offramp.owner == &config.ccip_router
        } @ CcipTokenPoolError::InvalidPoolCaller
    )]
    pub allowed_offramp: UncheckedAccount<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(address = *mint.to_account_info().owner)]
    /// CHECK: CPI to token program
    pub token_program: AccountInfo<'info>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [CCIP_TOKENPOOL_SIGNER, config.mint.key().as_ref()],
        bump,
    )]
    pub pool_signer: Account<'info, ExternalExecutionConfig>,
    #[account(mut, address = get_associated_token_address_with_program_id(&pool_signer.key(), &mint.key(), &token_program.key()))]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, release_or_mint.remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,

    // User specific accounts ---------------
    #[account(mut, address = get_associated_token_address_with_program_id(&release_or_mint.receiver, &mint.key(), &token_program.key()))]
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
    #[account(address = config.ramp_authority @ CcipTokenPoolError::InvalidPoolCaller)]
    pub authority: Signer<'info>,

    // Token pool accounts ------------------
    // consistent set + token pool program
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(address = *mint.to_account_info().owner)]
    /// CHECK: CPI to underlying token program
    pub token_program: AccountInfo<'info>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>,
    #[account(
        seeds = [CCIP_TOKENPOOL_SIGNER, config.mint.key().as_ref()],
        bump,
    )]
    pub pool_signer: Account<'info, ExternalExecutionConfig>,
    #[account(mut, address = get_associated_token_address_with_program_id(&pool_signer.key(), &mint.key(), &token_program.key()))]
    pub pool_token_account: InterfaceAccount<'info, TokenAccount>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, lock_or_burn.remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
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
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(
        init,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct EditChainConfigStaticSize<'info> {
    #[account(
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, cfg: RemoteConfig)]
pub struct EditChainConfigDynamicSize<'info> {
    #[account(
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE + cfg.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey, addresses: Vec<RemoteAddress>)]
pub struct AppendRemotePoolAddresses<'info> {
    #[account(
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + ChainConfig::INIT_SPACE
            + chain_config.remote.pool_addresses.iter().map(RemoteAddress::space).sum::<usize>()
            + addresses.iter().map(RemoteAddress::space).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(remote_chain_selector: u64, mint: Pubkey)]
pub struct DeleteChainConfig<'info> {
    #[account(
        seeds = [CCIP_TOKENPOOL_CONFIG, config.mint.key().as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [CCIP_TOKENPOOL_CHAINCONFIG, remote_chain_selector.to_le_bytes().as_ref(), mint.key().as_ref()],
        bump,
        close = authority,
    )]
    pub chain_config: Account<'info, ChainConfig>,
    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[error_code]
pub enum CcipTokenPoolError {
    #[msg("Pool authority does not match token mint owner")]
    InvalidInitPoolPermissions,
    #[msg("Unauthorized")]
    Unauthorized,
    #[msg("Invalid inputs")]
    InvalidInputs,
    #[msg("Caller is not ramp on router")]
    InvalidPoolCaller,
    #[msg("Invalid source pool address")]
    InvalidSourcePoolAddress,
    #[msg("Invalid token")]
    InvalidToken,
    #[msg("Invalid token amount conversion")]
    InvalidTokenAmountConversion,

    // Rate limit errors
    #[msg("RateLimit: bucket overfilled")]
    RLBucketOverfilled,
    #[msg("RateLimit: max capacity exceeded")]
    RLMaxCapacityExceeded,
    #[msg("RateLimit: rate limit reached")]
    RLRateLimitReached,
    #[msg("RateLimit: invalid rate limit rate")]
    RLInvalidRateLimitRate,
    #[msg("RateLimit: disabled non-zero rate limit")]
    RLDisabledNonZeroRateLimit,
}

#[event]
pub struct Burned {
    pub sender: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Minted {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Locked {
    pub sender: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Released {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

// note: configuration events are slightly different than EVM chains because configuration follows different steps
#[event]
pub struct RemoteChainConfigured {
    pub chain_selector: u64,
    pub token: RemoteAddress,
    pub previous_token: RemoteAddress,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RateLimitConfigured {
    pub chain_selector: u64,
    pub outbound_rate_limit: RateLimitConfig,
    pub inbound_rate_limit: RateLimitConfig,
}

#[event]
pub struct RemotePoolsAppended {
    pub chain_selector: u64,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RemoteChainRemoved {
    pub chain_selector: u64,
}

#[event]
pub struct RampAuthorityUpdated {
    pub old_authority: Pubkey,
    pub new_authority: Pubkey,
}

#[event]
pub struct RouterUpdated {
    pub old_router: Pubkey,
    pub new_router: Pubkey,
}
