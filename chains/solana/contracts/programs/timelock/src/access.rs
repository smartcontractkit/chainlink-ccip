use anchor_lang::prelude::*;

use access_controller::AccessController;

use crate::error::{AuthError, TimelockError};
use crate::state::{Config, Role};

/// A macro that checks if the authority has admin privileges OR has at least one of the specified roles.
/// This is used as an attribute macro with #[access_control] to guard program instructions.
///
/// # Arguments
/// - `$ctx` - The context containing program accounts
/// - `$role` - One or more roles that are allowed to execute the instruction
///
#[macro_export]
macro_rules! require_role_or_admin {
    ($ctx:expr, $($role:expr),+) => {{
        only_role_or_admin(
            &$ctx.accounts.config,
            &$ctx.accounts.role_access_controller,
            &$ctx.accounts.authority,
            &[$($role),+]
        )
    }};
}

/// Helper function to verify if an authority has admin rights or any of the specified roles.
/// Returns Ok(()) if the authority is either the admin or has at least one of the roles.
///
/// # Arguments
/// - `config` - Account containing program configuration including owner
/// - `role_controller` - Account managing role-based access control
/// - `authority` - The signer attempting to execute the instruction
/// - `roles` - Array of roles being checked for authorization
///
/// # Returns
/// - `Result<()>` - Ok if authorized
/// - `Err(TimelockError::InvalidAccessController)` - If provided controller isn't configured for any roles
/// - `Err(AuthError::Unauthorized)` - If authority lacks admin rights and required roles
pub fn only_role_or_admin(
    config: &AccountLoader<Config>,
    role_controller: &AccountLoader<AccessController>,
    authority: &Signer,
    roles: &[Role],
) -> Result<()> {
    let config = config.load()?;
    if authority.key() == config.owner {
        return Ok(());
    }

    require!(
        roles
            .iter()
            .any(|role| config.get_role_controller(role) == role_controller.key()),
        TimelockError::InvalidAccessController
    );

    let has_valid_role = roles.iter().any(|role| {
        config.get_role_controller(role) == role_controller.key()
            && access_controller::has_access(role_controller, &authority.key()).unwrap_or(false)
    });

    require!(has_valid_role, AuthError::Unauthorized);
    Ok(())
}

/// A macro that checks if the authority is the admin (owner) of the program.
/// This provides a simpler check when only admin access is required.
///
/// # Arguments
/// * `$ctx` - The context containing program accounts
#[macro_export]
macro_rules! require_only_admin {
    ($ctx:expr) => {
        only_admin(&$ctx.accounts.config, &$ctx.accounts.authority)
    };
}

/// Helper function to verify if an authority is the admin (owner).
/// Returns Ok(()) if the authority is the admin, Err otherwise.
///
/// # Arguments
/// - `config` - Account containing program configuration including owner
/// - `authority` - The signer attempting to execute the instruction
///
/// # Returns
/// - `Result<()>` - Ok if authorized, Err with AuthError::Unauthorized otherwise
pub fn only_admin(config: &AccountLoader<Config>, authority: &Signer) -> Result<()> {
    require_keys_eq!(
        authority.key(),
        config.load()?.owner,
        AuthError::Unauthorized
    );
    Ok(())
}
