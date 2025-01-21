use anchor_lang::prelude::*;

use access_controller::AccessController;

use crate::error::{AuthError, TimelockError};
use crate::state::{Config, Role};

/// A macro that checks if the authority has admin privileges OR has at least one of the specified roles.
/// This is used as an attribute macro with #[access_control] to guard program instructions.
///
/// # Arguments
/// * `$ctx` - The context containing program accounts
/// * `$role` - One or more roles that are allowed to execute the instruction
///
/// # Examples
/// ```rust
/// #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
/// pub fn single_role_instruction(ctx: Context<...>) -> Result<()> {
///     // ...
/// }
///
/// #[access_control(require_role_or_admin!(ctx, Role::Proposer, Role::Bypasser))]
/// pub fn multi_role_instruction(ctx: Context<...>) -> Result<()> {
///     // ...
/// }
/// ```
#[macro_export]
macro_rules! require_role_or_admin {
    ($ctx:expr, $($role:expr),+) => {
        {
            let mut result = Err(anchor_lang::error::Error::from(AuthError::Unauthorized));
            $(
                if only_role_or_admin(
                    &$ctx.accounts.config,
                    &$ctx.accounts.role_access_controller,
                    &$ctx.accounts.authority,
                    $role
                ).is_ok() {
                    result = Ok(());
                }
            )*
            result
        }
    };
}

/// Helper function to verify if an authority has admin rights or a specific role.
/// Returns Ok(()) if the authority is either the admin or has the specified role.
///
/// # Arguments
/// * `config` - Account containing program configuration including owner
/// * `role_controller` - Account managing role-based access control
/// * `authority` - The signer attempting to execute the instruction
/// * `role` - The role being checked for authorization
///
/// # Returns
/// * `Result<()>` - Ok if authorized, Err with AuthError::Unauthorized otherwise
pub fn only_role_or_admin(
    config: &Account<Config>,
    role_controller: &AccountLoader<AccessController>,
    authority: &Signer,
    role: Role,
) -> Result<()> {
    // Check if the authority is the admin (owner) first
    // This provides a bypass for the owner regardless of role assignments
    if authority.key() == config.owner {
        return Ok(());
    }

    // Verify that we're using the correct access controller for this role
    // This prevents using a different access controller than what's registered
    require_keys_eq!(
        config.get_role_controller(&role),
        role_controller.key(),
        TimelockError::InvalidAccessController
    );

    // Finally, check if the authority has been granted the required role
    require!(
        access_controller::has_access(role_controller, &authority.key())?,
        AuthError::Unauthorized
    );

    Ok(())
}

/// A macro that checks if the authority is the admin (owner) of the program.
/// This provides a simpler check when only admin access is required.
///
/// # Arguments
/// * `$ctx` - The context containing program accounts
///
/// # Examples
/// ```rust
/// #[access_control(require_only_admin!(ctx))]
/// pub fn admin_only_instruction(ctx: Context<...>) -> Result<()> {
///     // ...
/// }
/// ```
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
/// * `config` - Account containing program configuration including owner
/// * `authority` - The signer attempting to execute the instruction
///
/// # Returns
/// * `Result<()>` - Ok if authorized, Err with AuthError::Unauthorized otherwise
pub fn only_admin(config: &Account<Config>, authority: &Signer) -> Result<()> {
    require_keys_eq!(authority.key(), config.owner, AuthError::Unauthorized);
    Ok(())
}
