use anchor_lang::{prelude::*, system_program, Discriminator};
use anchor_spl::associated_token::get_associated_token_address_with_program_id;
use solana_program::{address_lookup_table::state::AddressLookupTable, log::sol_log};

use ccip_common::router_accounts::TokenAdminRegistry;
use ccip_common::seed::{self};

use crate::context::ANCHOR_DISCRIMINATOR;
use crate::events::token_admin_registry as events;
use crate::instructions::interfaces::TokenAdminRegistry as TokenAdminRegistryTrait;
use crate::token_admin_registry_writable;
use crate::token_context::{
    OverridePendingTokenAdminRegistryByCCIPAdmin, OverridePendingTokenAdminRegistryByOwner,
    RegisterTokenAdminRegistryByCCIPAdmin, RegisterTokenAdminRegistryByOwner,
    SetPoolTokenAdminRegistry, UpgradeTokenAdminRegistry,
};
use crate::{
    AcceptAdminRoleTokenAdminRegistry, CcipRouterError, EditPoolTokenAdminRegistry,
    ModifyTokenAdminRegistry,
};

const MINIMUM_TOKEN_POOL_ACCOUNTS: usize = 9;

pub struct Impl;
impl TokenAdminRegistryTrait for Impl {
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

    fn set_pool_supports_auto_derivation(
        &self,
        ctx: Context<EditPoolTokenAdminRegistry>,
        mint: Pubkey,
        supports_auto_derivation: bool,
    ) -> Result<()> {
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        token_admin_registry.supports_auto_derivation = supports_auto_derivation;

        emit!(events::PoolEdited {
            token: mint,
            supports_auto_derivation,
        });

        Ok(())
    }

    fn upgrade_token_admin_registry_from_v1(
        &self,
        ctx: Context<UpgradeTokenAdminRegistry>,
    ) -> Result<()> {
        let acc_info = ctx.accounts.token_admin_registry.to_account_info();

        // Validate the account to upgrade (the context already checks initialization and size)
        {
            msg!("Validating account to upgrade...");
            let data = acc_info.try_borrow_data()?;
            require!(
                data[..8] == TokenAdminRegistry::DISCRIMINATOR,
                CcipRouterError::InvalidInputsTokenAdminRegistryAccounts
            );
            require_eq!(
                data[8], // version byte
                1,
                CcipRouterError::InvalidInputsTokenAdminRegistryAccounts
            );
        }

        // Extend the account to the new size (realloc)
        {
            msg!("Extending account...");
            let required_space = ANCHOR_DISCRIMINATOR + TokenAdminRegistry::INIT_SPACE;
            let minimum_balance = Rent::get()?.minimum_balance(required_space);
            let current_lamports = acc_info.lamports();

            if current_lamports < minimum_balance {
                system_program::transfer(
                    CpiContext::new(
                        ctx.accounts.system_program.to_account_info(),
                        system_program::Transfer {
                            from: ctx.accounts.authority.to_account_info(),
                            to: acc_info.clone(),
                        },
                    ),
                    minimum_balance.checked_sub(current_lamports).unwrap(),
                )?;
            }
            acc_info.realloc(required_space, false)?;
        }

        let mut token_admin_registry: TokenAdminRegistry;

        // Load realloc'd data
        {
            msg!("Deserialize account...");
            let data = acc_info.try_borrow_data()?;
            let mut data: &[u8] = &data;

            token_admin_registry = TokenAdminRegistry::try_deserialize(&mut data)?;
        }

        // Upgrade data from v1 to v2
        token_admin_registry.version = 2;
        token_admin_registry.supports_auto_derivation = false; // Safe default value.
                                                               // Upgradeable by its admin later with other methods

        // Persist the upgraded account data
        msg!("Re-serialize account...");
        token_admin_registry.try_serialize(&mut &mut acc_info.try_borrow_mut_data()?[..])?;

        emit!(crate::events::PdaUpgraded {
            address: acc_info.key(),
            old_version: 1,
            new_version: 2,
            name: "TokenAdminRegistry".to_string(),
            seeds: [
                seed::TOKEN_ADMIN_REGISTRY,
                token_admin_registry.mint.as_ref()
            ]
            .concat(),
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

            // The mandatory accounts (PDAs) are stored in the lookup table to save space even if they can be inferred
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
    token_admin_registry: &mut Account<'_, TokenAdminRegistry>,
    token_mint: Pubkey,
    new_admin: Pubkey,
) -> Result<()> {
    require_neq!(
        new_admin,
        Pubkey::new_from_array([0; 32]),
        CcipRouterError::InvalidTokenAdminRegistryInputsZeroAddress
    );

    token_admin_registry.set_inner(TokenAdminRegistry {
        version: 2,
        administrator: Pubkey::new_from_array([0; 32]),
        pending_administrator: new_admin,
        lookup_table: Pubkey::new_from_array([0; 32]),
        writable_indexes: [0, 0],
        mint: token_mint,
        supports_auto_derivation: false,
    });

    emit!(events::AdministratorTransferRequested {
        token: token_mint,
        current_admin: Pubkey::new_from_array([0; 32]),
        new_admin,
    });

    Ok(())
}
