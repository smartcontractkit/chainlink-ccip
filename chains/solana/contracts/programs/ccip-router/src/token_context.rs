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
    // binary representation of indexes that are writable in token pool lookup table
    // lookup table can store 256 addresses
    pub writable_indexes: [u128; 2],
}

impl TokenAdminRegistry {
    // set writable inserts bits from left to right
    // index 0 is left-most bit
    pub fn set_writable(&mut self, index: u8) {
        match index < 128 {
            true => {
                self.writable_indexes[0] |= 1 << (127 - index);
            }
            false => {
                self.writable_indexes[1] |= 1 << (255 - index);
            }
        }
    }

    pub fn is_writable(&self, index: u8) -> bool {
        match index < 128 {
            true => self.writable_indexes[0] & 1 << (127 - index) != 0,
            false => self.writable_indexes[1] & 1 << (255 - index) != 0,
        }
    }

    pub fn reset_writable(&mut self) {
        self.writable_indexes = [0, 0];
    }
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

#[cfg(test)]
mod tests {
    use bytemuck::Zeroable;
    use solana_program::pubkey::Pubkey;

    use crate::TokenAdminRegistry;
    #[test]
    fn set_writable() {
        let mut state = TokenAdminRegistry {
            version: 0,
            administrator: Pubkey::zeroed(),
            pending_administrator: Pubkey::zeroed(),
            lookup_table: Pubkey::zeroed(),
            writable_indexes: [0, 0],
        };

        state.set_writable(0);
        state.set_writable(128);
        assert_eq!(state.writable_indexes[0], 2u128.pow(127));
        assert_eq!(state.writable_indexes[1], 2u128.pow(127));

        state.reset_writable();
        assert_eq!(state.writable_indexes[0], 0);
        assert_eq!(state.writable_indexes[1], 0);

        state.set_writable(0);
        state.set_writable(2);
        state.set_writable(127);
        state.set_writable(128);
        state.set_writable(2 + 128);
        state.set_writable(255);
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
        let state = TokenAdminRegistry {
            version: 0,
            administrator: Pubkey::zeroed(),
            pending_administrator: Pubkey::zeroed(),
            lookup_table: Pubkey::zeroed(),
            writable_indexes: [
                2u128.pow(127 - 7) + 2u128.pow(127 - 2) + 2u128.pow(127 - 4),
                2u128.pow(127 - 8) + 2u128.pow(127 - 56) + 2u128.pow(127 - 100),
            ],
        };

        assert_eq!(state.is_writable(0), false);
        assert_eq!(state.is_writable(128), false);
        assert_eq!(state.is_writable(255), false);

        assert_eq!(state.writable_indexes[0].count_ones(), 3);
        assert_eq!(state.writable_indexes[1].count_ones(), 3);

        assert_eq!(state.is_writable(7), true);
        assert_eq!(state.is_writable(2), true);
        assert_eq!(state.is_writable(4), true);
        assert_eq!(state.is_writable(128 + 8), true);
        assert_eq!(state.is_writable(128 + 56), true);
        assert_eq!(state.is_writable(128 + 100), true);
    }
}
