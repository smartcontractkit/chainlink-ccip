use anchor_lang::prelude::*;

use access_controller::AccessController;

use crate::error::TimelockError;
use crate::state::{Config, Role};

// NOTE: This macro is used to check if the authority is the owner
// or account that has access to the role
#[macro_export]
macro_rules! only_role_or_admin_role {
    ($ctx:expr, $role:expr) => {
        only_role_or_admin_role(
            &$ctx.accounts.config,
            &$ctx.accounts.role_access_controller,
            &$ctx.accounts.authority,
            $role,
        )
    };
}

pub fn only_role_or_admin_role(
    config: &Account<Config>,
    role_controller: &AccountLoader<AccessController>,
    authority: &Signer,
    role: Role,
) -> Result<()> {
    // early authorization check if the authority is the owner
    if authority.key() == config.owner {
        return Ok(());
    }

    // check the provided access controller is the correct one
    require_keys_eq!(
        config.get_role_controller(&role),
        role_controller.key(),
        TimelockError::InvalidAccessController
    );

    // check if the authority has access to the role
    require!(
        access_controller::has_access(role_controller, &authority.key())?,
        TimelockError::Unauthorized
    );

    Ok(())
}

// NOTE: This macro is used to check if the authority is the owner
// this can be in account contexts, but added also for explicit consistency with other role check
#[macro_export]
macro_rules! only_admin {
    ($ctx:expr) => {
        only_admin(&$ctx.accounts.config, &$ctx.accounts.authority)
    };
}

pub fn only_admin(config: &Account<Config>, authority: &Signer) -> Result<()> {
    require_keys_eq!(authority.key(), config.owner, TimelockError::Unauthorized);
    Ok(())
}
