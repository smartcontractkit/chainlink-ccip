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
    pub mint: Pubkey,
}

#[derive(Accounts)]
pub struct RegisterTokenAdminRegistryByCCIPAdmin<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        init,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    // The following validation is the only difference between the two contexts
    // Only CCIP Admin can propose an owner for the token admin registry if not the mint authority
    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct OverridePendingTokenAdminRegistryByCCIPAdmin<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    // The following validation is the only difference between the two contexts
    // Only CCIP Admin can propose an owner for the token admin registry if not the mint authority
    #[account(mut, address = config.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct RegisterTokenAdminRegistryByOwner<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        init,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
        constraint = uninitialized(token_admin_registry.version) @ CcipRouterError::InvalidVersion,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    // The following validation is the only difference between the two contexts
    // Only the mint authority can propose an owner for the token admin registry
    #[account(
        mut,
        address = mint.mint_authority.unwrap() @ CcipRouterError::Unauthorized,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct OverridePendingTokenAdminRegistryByOwner<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidVersion,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    // The following validation is the only difference between the two contexts
    // Only the mint authority can propose an owner for the token admin registry
    #[account(
        mut,
        address = mint.mint_authority.unwrap() @ CcipRouterError::Unauthorized,
    )]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct ModifyTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidVersion,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(mut, address = token_admin_registry.administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct SetPoolTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidVersion,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    /// CHECK: anchor does not support automatic lookup table deserialization
    pub pool_lookuptable: UncheckedAccount<'info>,
    #[account(mut, address = token_admin_registry.administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptAdminRoleTokenAdminRegistry<'info> {
    #[account(
        seeds = [seed::CONFIG],
        bump,
        constraint = valid_version(config.version, MAX_CONFIG_V) @ CcipRouterError::InvalidVersion,
    )]
    pub config: Account<'info, Config>,
    #[account(
        mut,
        seeds = [seed::TOKEN_ADMIN_REGISTRY, mint.key().as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidVersion,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    pub mint: InterfaceAccount<'info, Mint>, // underlying token that the pool wraps
    #[account(mut, address = token_admin_registry.pending_administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}
