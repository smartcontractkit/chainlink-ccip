use anchor_lang::prelude::*;

use crate::context::*;
use crate::state::*;
use crate::CcipRouterError;

use anchor_spl::token_interface::Mint;

#[account]
#[derive(InitSpace)]
pub struct TokenAdminRegistry {
    pub version: u8,
    pub administrator: Pubkey,
    pub pending_administrator: Pubkey,
    pub lookup_table: Pubkey,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct RegisterTokenAdminRegistryViaGetCCIPAdmin<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub config: AccountLoader<'info, Config>,
    #[account(
        init,
        seeds = [TOKEN_ADMIN_REGISTRY_SEED, mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
        constraint = uninitialized(token_admin_registry.version) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct RegisterTokenAdminRegistryViaOwner<'info> {
    #[account(
        init,
        seeds = [TOKEN_ADMIN_REGISTRY_SEED, mint.key().as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE,
        constraint = uninitialized(token_admin_registry.version) @ CcipRouterError::InvalidInputs,
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
        mut,
        seeds = [TOKEN_ADMIN_REGISTRY_SEED, mint.as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = token_admin_registry.administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(mint: Pubkey)]
pub struct AcceptAdminRoleTokenAdminRegistry<'info> {
    #[account(
        mut,
        seeds = [TOKEN_ADMIN_REGISTRY_SEED, mint.as_ref()],
        bump,
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidInputs,
    )]
    pub token_admin_registry: Account<'info, TokenAdminRegistry>,
    #[account(mut, address = token_admin_registry.pending_administrator @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64, mint: Pubkey)]
pub struct SetTokenBillingConfig<'info> {
    #[account(
        seeds = [CONFIG_SEED],
        bump,
        constraint = valid_version(config.load()?.version, MAX_CONFIG_V) @ CcipRouterError::InvalidInputs, // validate state version
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(
        init_if_needed,
        seeds = [TOKEN_POOL_BILLING_SEED, chain_selector.to_le_bytes().as_ref(), mint.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + PerChainPerTokenConfig::INIT_SPACE,
        constraint = uninitialized(per_chain_per_token_config.version) || valid_version(per_chain_per_token_config.version, MAX_TOKEN_AND_CHAIN_CONFIG_V) @ CcipRouterError::InvalidInputs,
    )]
    pub per_chain_per_token_config: Account<'info, PerChainPerTokenConfig>,

    // validate signer is registered ccip admin
    #[account(mut, address = config.load()?.owner @ CcipRouterError::Unauthorized)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}
