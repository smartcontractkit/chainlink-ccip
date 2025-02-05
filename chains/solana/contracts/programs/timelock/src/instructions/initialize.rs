use anchor_lang::prelude::*;

use access_controller::{AccessController, AccessControllerProgram};

use crate::constants::{ANCHOR_DISCRIMINATOR, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED};
use crate::error::{AuthError, TimelockError};
use crate::program::Timelock;
use crate::state::{Config, Role};

/// initialize Timelock config with owner(admin),
/// role access controller keys and global configuration value.
pub fn initialize(
    ctx: Context<Initialize>,
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    min_delay: u64,
) -> Result<()> {
    // assign owner(owner is admin)
    let config = &mut ctx.accounts.config;
    config.timelock_id = timelock_id;
    config.owner = ctx.accounts.authority.key();
    config.min_delay = min_delay;

    config.proposer_role_access_controller = ctx.accounts.proposer_role_access_controller.key();
    config.executor_role_access_controller = ctx.accounts.executor_role_access_controller.key();
    config.canceller_role_access_controller = ctx.accounts.canceller_role_access_controller.key();
    config.bypasser_role_access_controller = ctx.accounts.bypasser_role_access_controller.key();

    Ok(())
}

/// wrapper function that calls access_controller::cpi::add_access via batch CPI call.
/// target addresses should be provided in remaining_accounts,
/// tested with up to 24 addresses per each transaction.
pub fn batch_add_access<'info>(
    ctx: Context<'_, '_, '_, 'info, BatchAddAccess<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    role: Role,
) -> Result<()> {
    let expected = ctx.accounts.config.get_role_controller(&role);
    require_keys_eq!(
        expected,
        ctx.accounts.role_access_controller.key(),
        TimelockError::InvalidAccessController
    );

    require!(
        !ctx.remaining_accounts.is_empty(),
        TimelockError::InvalidInput
    );

    for account_info in ctx.remaining_accounts.iter() {
        let cpi_accounts = access_controller::cpi::accounts::AddAccess {
            state: ctx.accounts.role_access_controller.to_account_info(),
            owner: ctx.accounts.authority.to_account_info(),
            address: account_info.clone(),
        };
        let cpi_ctx = CpiContext::new(
            ctx.accounts.access_controller_program.to_account_info(),
            cpi_accounts,
        );
        access_controller::cpi::add_access(cpi_ctx)?;
    }
    Ok(())
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct Initialize<'info> {
    #[account(
        init,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
        seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()],
        bump,
        payer = authority,
    )]
    pub config: Account<'info, Config>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, Timelock>,
    // NOTE: initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ AuthError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,

    // access controller program and states per role
    pub access_controller_program: Program<'info, AccessControllerProgram>,
    #[account(owner = access_controller_program.key())]
    pub proposer_role_access_controller: AccountLoader<'info, AccessController>,
    #[account(owner = access_controller_program.key())]
    pub executor_role_access_controller: AccountLoader<'info, AccessController>,
    #[account(owner = access_controller_program.key())]
    pub canceller_role_access_controller: AccountLoader<'info, AccessController>,
    #[account(owner = access_controller_program.key())]
    pub bypasser_role_access_controller: AccountLoader<'info, AccessController>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], role: Role)]
pub struct BatchAddAccess<'info> {
    #[account(
        seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()],
        bump,
    )]
    pub config: Account<'info, Config>,

    pub access_controller_program: Program<'info, AccessControllerProgram>,

    // NOTE: access controller for the role of access list
    #[account(
        mut,
        owner = access_controller_program.key(),
    )]
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut, address = config.owner)]
    pub authority: Signer<'info>,
}
