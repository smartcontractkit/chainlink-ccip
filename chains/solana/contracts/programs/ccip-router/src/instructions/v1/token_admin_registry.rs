use crate::{
    events::token_admin_registry as events,
    instructions::interfaces::TokenAdminRegistry,
    token_admin_registry_writable,
    token_context::{
        OverridePendingTokenAdminRegistryByCCIPAdmin, OverridePendingTokenAdminRegistryByOwner,
    },
};
use anchor_lang::prelude::*;
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use ccip_common::seed;
use solana_program::{address_lookup_table::state::AddressLookupTable, log::sol_log};

use crate::{
    token_context::{RegisterTokenAdminRegistryByCCIPAdmin, RegisterTokenAdminRegistryByOwner},
    AcceptAdminRoleTokenAdminRegistry, CcipRouterError, ModifyTokenAdminRegistry,
    SetPoolTokenAdminRegistry,
};

const MINIMUM_TOKEN_POOL_ACCOUNTS: usize = 9;

pub struct Impl;
impl TokenAdminRegistry for Impl {
    fn ccip_admin_propose_administrator(
        &self,
        ctx: Context<RegisterTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        let token_mint = ctx.accounts.mint.key().to_owned();
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        init_with_pending_administrator(
            token_admin_registry,
            token_mint,
            token_admin_registry_admin,
        )?;

        Ok(())
    }

    fn ccip_admin_override_pending_administrator(
        &self,
        ctx: Context<OverridePendingTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        require_eq!(
            token_admin_registry.administrator,
            Pubkey::new_from_array([0; 32]),
            CcipRouterError::InvalidTokenAdminRegistryProposedAdmin
        );

        token_admin_registry.pending_administrator = token_admin_registry_admin;

        Ok(())
    }

    fn owner_propose_administrator(
        &self,
        ctx: Context<RegisterTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        let token_mint = ctx.accounts.mint.key().to_owned();
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        init_with_pending_administrator(
            token_admin_registry,
            token_mint,
            token_admin_registry_admin,
        )?;

        Ok(())
    }

    fn owner_override_pending_administrator(
        &self,
        ctx: Context<OverridePendingTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        require_eq!(
            token_admin_registry.administrator,
            Pubkey::new_from_array([0; 32]),
            CcipRouterError::InvalidTokenAdminRegistryProposedAdmin
        );

        token_admin_registry.pending_administrator = token_admin_registry_admin;

        Ok(())
    }

    fn transfer_admin_role_token_admin_registry(
        &self,
        ctx: Context<ModifyTokenAdminRegistry>,
        new_admin: Pubkey,
    ) -> Result<()> {
        let token_mint = ctx.accounts.mint.key().to_owned();
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        token_admin_registry.pending_administrator = new_admin;

        emit!(events::AdministratorTransferRequested {
            token: token_mint,
            current_admin: token_admin_registry.administrator,
            new_admin,
        });

        Ok(())
    }

    fn accept_admin_role_token_admin_registry(
        &self,
        ctx: Context<AcceptAdminRoleTokenAdminRegistry>,
    ) -> Result<()> {
        let token_mint = ctx.accounts.mint.key().to_owned();
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        let new_admin = token_admin_registry.pending_administrator;

        token_admin_registry.administrator = new_admin;
        token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);

        emit!(events::AdministratorTransferred {
            token: token_mint,
            new_admin,
        });

        Ok(())
    }

    // TODO: Add a test with look up table as 0 address
    fn set_pool(
        &self,
        ctx: Context<SetPoolTokenAdminRegistry>,
        writable_indexes: Vec<u8>,
    ) -> Result<()> {
        let token_mint = ctx.accounts.mint.key().to_owned();

        // Enable to overwrite the lookup table pool with 0 to delist from CCIP
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        let previous_pool = token_admin_registry.lookup_table;
        let new_pool = ctx.accounts.pool_lookuptable.key();
        token_admin_registry.lookup_table = new_pool;

        // set writable indexes
        // used to build offchain tokenpool additional accounts in the token transactions
        // indexes can be checked to indicate is_writable for the specific account
        // this is also used to validate to ensure the correct write permissions are used according the admin
        token_admin_registry_writable::reset(token_admin_registry);
        for ind in writable_indexes {
            token_admin_registry_writable::set(token_admin_registry, ind)
        }

        // validate lookup table contains minimum required accounts if not zero address
        if new_pool != Pubkey::default() {
            // deserialize lookup table account
            let lookup_table_data = &mut &ctx.accounts.pool_lookuptable.data.borrow()[..];
            let lookup_table_account: AddressLookupTable =
                AddressLookupTable::deserialize(lookup_table_data)
                    .map_err(|_| CcipRouterError::InvalidInputsLookupTableAccounts)?;
            require_gte!(
                lookup_table_account.addresses.len(),
                MINIMUM_TOKEN_POOL_ACCOUNTS,
                CcipRouterError::InvalidInputsLookupTableAccounts
            );

            // The mandatory accounts (PDAs) are stored in the lookup table to save space even if they can be infered
            let (token_admin_registry, _) = Pubkey::find_program_address(
                &[seed::TOKEN_ADMIN_REGISTRY, token_mint.as_ref()],
                ctx.program_id,
            );
            let pool_program = lookup_table_account.addresses[2]; // cannot be calculated, can be custom per pool
            let token_program = lookup_table_account.addresses[6]; // cannot be calculated, can be custom per token
            let (pool_config, _) = Pubkey::find_program_address(
                &[seed::CCIP_TOKENPOOL_CONFIG, token_mint.as_ref()],
                &pool_program,
            );
            let (pool_signer, _) = Pubkey::find_program_address(
                &[seed::CCIP_TOKENPOOL_SIGNER, token_mint.as_ref()],
                &pool_program,
            );
            let (fee_billing_config, _) = Pubkey::find_program_address(
                &[seed::FEE_BILLING_TOKEN_CONFIG, token_mint.as_ref()],
                &ctx.accounts.config.fee_quoter,
            );

            // verify token mint owner matches token program
            require_keys_eq!(
                *ctx.accounts.mint.to_account_info().owner,
                token_program,
                CcipRouterError::InvalidInputsMint
            );

            let min_accounts = [
                ctx.accounts.pool_lookuptable.key(),
                token_admin_registry,
                pool_program,
                pool_config,
                get_associated_token_address_with_program_id(
                    &pool_signer.key(),
                    &token_mint.key(),
                    &token_program.key(),
                ),
                pool_signer,
                token_program,
                token_mint,
                fee_billing_config,
            ];

            for (i, acc) in min_accounts.iter().enumerate() {
                // for easier debugging UX
                if lookup_table_account.addresses[i] != *acc {
                    sol_log(&i.to_string());
                    sol_log(&acc.to_string());
                    sol_log(&lookup_table_account.addresses[i].to_string());
                }
                require_eq!(
                    lookup_table_account.addresses[i],
                    *acc,
                    CcipRouterError::InvalidInputsLookupTableAccounts
                );
            }
        }

        emit!(events::PoolSet {
            token: token_mint,
            previous_pool_lookup_table: previous_pool,
            new_pool_lookup_table: new_pool,
        });

        Ok(())
    }
}

fn init_with_pending_administrator(
    token_admin_registry: &mut Account<'_, ccip_common::router_accounts::TokenAdminRegistry>,
    token_mint: Pubkey,
    new_admin: Pubkey,
) -> Result<()> {
    require_neq!(
        new_admin,
        Pubkey::new_from_array([0; 32]),
        CcipRouterError::InvalidTokenAdminRegistryInputsZeroAddress
    );

    token_admin_registry.version = 1;
    token_admin_registry.administrator = Pubkey::new_from_array([0; 32]);
    token_admin_registry.pending_administrator = new_admin;
    token_admin_registry.lookup_table = Pubkey::new_from_array([0; 32]);
    token_admin_registry.mint = token_mint;

    emit!(events::AdministratorTransferRequested {
        token: token_mint,
        current_admin: Pubkey::new_from_array([0; 32]),
        new_admin,
    });

    Ok(())
}
