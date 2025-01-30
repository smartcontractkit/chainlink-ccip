use anchor_lang::prelude::*;

use crate::context::*;
use crate::state::*;
use crate::CcipRouterError;

use anchor_spl::token_interface::Mint;

// track state versions
const MAX_TOKEN_REGISTRY_V: u8 = 1;

#[account]
#[derive(InitSpace)]
pub struct TokenAdminRegistry {
    pub version: u8,
    pub administrator: Pubkey,
    pub pending_administrator: Pubkey,
    pub lookup_table: Pubkey,
    // binary representation of indexes that are writable in token pool lookup table
    // lookup table can store 256 addresses
    pub writable_indexes: [u128; 2],
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct RegisterTokenAdminRegistryViaGetCCIPAdmin<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        init,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct RegisterTokenAdminRegistryViaOwner<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        init,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut)]
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(
        mut,
        address = mint.mint_authority.unwrap() @ CcipRouterError::Unauthorized,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct ModifyTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = token_admin_registry.administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct SetPoolTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    /// CHECK: anchor does not support automatic lookup table deserialization
    pub pool_lookuptable: UncheckedAccount<'info>,
    #[account(mut, address = token_admin_registry.administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct AcceptAdminRoleTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = token_admin_registry.pending_administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}
