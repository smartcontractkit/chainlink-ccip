use anchor_lang::prelude::*;
use ccip_common::router_accounts::TokenAdminRegistry;
use ccip_common::seed;

use crate::context::*;
use crate::state::*;
use crate::CcipRouterError;

use anchor_spl::token_interface::Mint;

// track state versions
const MAX_TOKEN_REGISTRY_V: u8 = 1;

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
        constraint = uninitialized(token_admin_registry.version) @ CcipRouterError::InvalidVersion,
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
        constraint = valid_version(token_admin_registry.version, MAX_TOKEN_REGISTRY_V) @ CcipRouterError::InvalidVersion,
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

pub mod token_admin_registry_writable {
    use super::TokenAdminRegistry;

    // set writable inserts bits from left to right
    // index 0 is left-most bit
    pub fn set(tar: &mut TokenAdminRegistry, index: u8) {
        match index < 128 {
            true => {
                tar.writable_indexes[0] |= 1 << (127 - index);
            }
            false => {
                tar.writable_indexes[1] |= 1 << (255 - index);
            }
        }
    }

    pub fn reset(tar: &mut TokenAdminRegistry) {
        tar.writable_indexes = [0, 0];
    }

    #[cfg(test)]
    mod tests {
        use super::*;
        use solana_program::pubkey::Pubkey;

        fn is(tar: &TokenAdminRegistry, index: u8) -> bool {
            match index < 128 {
                true => tar.writable_indexes[0] & 1 << (127 - index) != 0,
                false => tar.writable_indexes[1] & 1 << (255 - index) != 0,
            }
        }

        #[test]
        fn set_writable() {
            let state = &mut TokenAdminRegistry {
                version: 0,
                administrator: Pubkey::default(),
                pending_administrator: Pubkey::default(),
                lookup_table: Pubkey::default(),
                writable_indexes: [0, 0],
                mint: Pubkey::default(),
            };

            set(state, 0);
            set(state, 128);
            assert_eq!(state.writable_indexes[0], 2u128.pow(127));
            assert_eq!(state.writable_indexes[1], 2u128.pow(127));

            reset(state);
            assert_eq!(state.writable_indexes[0], 0);
            assert_eq!(state.writable_indexes[1], 0);

            set(state, 0);
            set(state, 2);
            set(state, 127);
            set(state, 128);
            set(state, 2 + 128);
            set(state, 255);
            assert_eq!(
                state.writable_indexes[0],
                2u128.pow(127) + 2u128.pow(127 - 2) + 2u128.pow(0)
            );
            assert_eq!(
                state.writable_indexes[1],
                2u128.pow(127) + 2u128.pow(127 - 2) + 2u128.pow(0)
            );
        }

        #[test]
        fn check_writable() {
            let state = &TokenAdminRegistry {
                version: 0,
                administrator: Pubkey::default(),
                pending_administrator: Pubkey::default(),
                lookup_table: Pubkey::default(),
                writable_indexes: [
                    2u128.pow(127 - 7) + 2u128.pow(127 - 2) + 2u128.pow(127 - 4),
                    2u128.pow(127 - 8) + 2u128.pow(127 - 56) + 2u128.pow(127 - 100),
                ],
                mint: Pubkey::default(),
            };

            assert!(!is(state, 0));
            assert!(!is(state, 128));
            assert!(!is(state, 255));

            assert_eq!(state.writable_indexes[0].count_ones(), 3);
            assert_eq!(state.writable_indexes[1].count_ones(), 3);

            assert!(is(state, 7));
            assert!(is(state, 2));
            assert!(is(state, 4));
            assert!(is(state, 128 + 8));
            assert!(is(state, 128 + 56));
            assert!(is(state, 128 + 100));
        }
    }
}
