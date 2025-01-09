use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::get_associated_token_address_with_program_id, token::spl_token::native_mint,
};
use bytemuck::Zeroable;
use solana_program::{address_lookup_table::state::AddressLookupTable, log::sol_log};

use crate::{
    AcceptAdminRoleTokenAdminRegistry, AdministratorRegistered, AdministratorTransferRequested,
    AdministratorTransferred, CcipRouterError, ModifyTokenAdminRegistry, PoolSet,
    RegisterTokenAdminRegistryViaGetCCIPAdmin, RegisterTokenAdminRegistryViaOwner,
    SetPoolTokenAdminRegistry, CCIP_TOKENPOOL_CONFIG, CCIP_TOKENPOOL_SIGNER,
    FEE_BILLING_TOKEN_CONFIG, TOKEN_ADMIN_REGISTRY_SEED,
};

pub fn register_token_admin_registry_via_get_ccip_admin(
    ctx: Context<RegisterTokenAdminRegistryViaGetCCIPAdmin>,
    mint: Pubkey, // should we validate that this is a real token program?
    token_admin_registry_admin: Pubkey,
) -> Result<()> {
    let token_admin_registry = &mut ctx.accounts.token_admin_registry;
    token_admin_registry.version = 1;
    token_admin_registry.administrator = token_admin_registry_admin;
    token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
    token_admin_registry.lookup_table = Pubkey::new_from_array([0; 32]);

    emit!(AdministratorRegistered {
        token_mint: mint,
        administrator: token_admin_registry_admin,
    });

    Ok(())
}

pub fn register_token_admin_registry_via_owner(
    ctx: Context<RegisterTokenAdminRegistryViaOwner>,
) -> Result<()> {
    let token_mint = ctx.accounts.mint.key().to_owned();
    let mint_authority = ctx.accounts.mint.mint_authority.to_owned();

    require!(
        mint_authority.is_some(),
        CcipRouterError::InvalidInputsPoolAccounts
    );

    let administrator = mint_authority.unwrap();

    let token_admin_registry = &mut ctx.accounts.token_admin_registry;
    token_admin_registry.version = 1;
    token_admin_registry.administrator = administrator;
    token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
    token_admin_registry.lookup_table = Pubkey::new_from_array([0; 32]);

    emit!(AdministratorRegistered {
        token_mint,
        administrator,
    });

    Ok(())
}

pub fn set_pool(
    ctx: Context<SetPoolTokenAdminRegistry>,
    mint: Pubkey,
    writable_indexes: Vec<u8>,
) -> Result<()> {
    // set new lookup table
    let token_admin_registry = &mut ctx.accounts.token_admin_registry;
    let previous_pool = token_admin_registry.lookup_table;
    let new_pool = ctx.accounts.pool_lookuptable.key();
    token_admin_registry.lookup_table = new_pool;

    // set writable indexes
    token_admin_registry.reset_writable();
    for ind in writable_indexes {
        token_admin_registry.set_writable(ind)
    }

    // validate lookup table contains minimum required accounts if not zero address
    if new_pool != Pubkey::zeroed() {
        // deserialize lookup table account
        let lookup_table_data = &mut &ctx.accounts.pool_lookuptable.data.borrow()[..];
        let lookup_table_account: AddressLookupTable =
            AddressLookupTable::deserialize(lookup_table_data)
                .map_err(|_| CcipRouterError::InvalidInputsLookupTableAccounts)?;
        require!(
            lookup_table_account.addresses.len() >= 9,
            CcipRouterError::InvalidInputsLookupTableAccounts
        );

        // calculate or retrieve expected addresses
        let (token_admin_registry, _) = Pubkey::find_program_address(
            &[TOKEN_ADMIN_REGISTRY_SEED, mint.as_ref()],
            ctx.program_id,
        );
        let pool_program = lookup_table_account.addresses[2]; // cannot be calculated, can be custom per pool
        let token_program = lookup_table_account.addresses[6]; // cannot be calculated, can be custom per token
        let (pool_config, _) =
            Pubkey::find_program_address(&[CCIP_TOKENPOOL_CONFIG, mint.as_ref()], &pool_program);
        let (pool_signer, _) =
            Pubkey::find_program_address(&[CCIP_TOKENPOOL_SIGNER, mint.as_ref()], &pool_program);
        let (fee_billing_config, _) = Pubkey::find_program_address(
            &[
                FEE_BILLING_TOKEN_CONFIG,
                if mint.key() == Pubkey::default() {
                    native_mint::ID.as_ref() // pre-2022 WSOL
                } else {
                    mint.as_ref()
                },
            ],
            ctx.program_id,
        );

        let min_accounts = [
            ctx.accounts.pool_lookuptable.key(),
            token_admin_registry,
            pool_program,
            pool_config,
            get_associated_token_address_with_program_id(
                &pool_signer.key(),
                &mint.key(),
                &token_program.key(),
            ),
            pool_signer,
            token_program,
            mint,
            fee_billing_config,
        ];

        for (i, acc) in min_accounts.iter().enumerate() {
            // for easier debugging UX
            if lookup_table_account.addresses[i] != *acc {
                sol_log(&i.to_string());
                sol_log(&acc.to_string());
                sol_log(&lookup_table_account.addresses[i].to_string());
            }
            require!(
                lookup_table_account.addresses[i] == *acc,
                CcipRouterError::InvalidInputsLookupTableAccounts
            );
        }
    }

    emit!(PoolSet {
        token: mint,
        previous_pool_lookup_table: previous_pool,
        new_pool_lookup_table: new_pool,
    });

    Ok(())
}

pub fn transfer_admin_role_token_admin_registry(
    ctx: Context<ModifyTokenAdminRegistry>,
    mint: Pubkey,
    new_admin: Pubkey,
) -> Result<()> {
    let token_admin_registry = &mut ctx.accounts.token_admin_registry;

    if new_admin == Pubkey::new_from_array([0; 32]) {
        token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
    } else {
        token_admin_registry.pending_administrator = new_admin;
    }

    emit!(AdministratorTransferRequested {
        token: mint,
        current_admin: token_admin_registry.administrator,
        new_admin,
    });

    Ok(())
}

pub fn accept_admin_role_token_admin_registry(
    ctx: Context<AcceptAdminRoleTokenAdminRegistry>,
    mint: Pubkey,
) -> Result<()> {
    let token_admin_registry = &mut ctx.accounts.token_admin_registry;
    let new_admin = token_admin_registry.pending_administrator;
    token_admin_registry.administrator = new_admin;
    token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);

    emit!(AdministratorTransferred {
        token: mint,
        new_admin,
    });

    Ok(())
}
