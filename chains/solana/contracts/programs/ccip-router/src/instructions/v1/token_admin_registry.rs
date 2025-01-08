use anchor_lang::prelude::*;

use crate::{
    AcceptAdminRoleTokenAdminRegistry, AdministratorRegistered, AdministratorTransferRequested,
    AdministratorTransferred, CcipRouterError, ModifyTokenAdminRegistry, PoolSet,
    RegisterTokenAdminRegistryViaGetCCIPAdmin, RegisterTokenAdminRegistryViaOwner,
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
    ctx: Context<ModifyTokenAdminRegistry>,
    mint: Pubkey,
    pool_lookup_table: Pubkey,
) -> Result<()> {
    let token_admin_registry = &mut ctx.accounts.token_admin_registry;
    let previous_pool = token_admin_registry.lookup_table;
    token_admin_registry.lookup_table = pool_lookup_table;

    // TODO: Validate here that the lookup table has everything

    emit!(PoolSet {
        token: mint,
        previous_pool_lookup_table: previous_pool,
        new_pool_lookup_table: pool_lookup_table,
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
