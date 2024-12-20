use anchor_lang::prelude::*;

use access_controller::AccessController;

use crate::constants::{TIMELOCK_CONFIG_SEED, TIMELOCK_OPERATION_SEED};
use crate::error::TimelockError;
use crate::event::*;
use crate::state::{Config, Operation};

/// cancels a pending operation
pub fn cancel<'info>(_ctx: Context<'_, '_, '_, 'info, Cancel<'info>>, id: [u8; 32]) -> Result<()> {
    emit!(Cancelled { id });
    // NOTE: PDA is closed - is handled by anchor on exit due to the `close` attribute
    Ok(())
}

#[derive(Accounts)]
#[instruction(id: [u8; 32])]
pub struct Cancel<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, id.as_ref()],
        bump,
        close = authority, // todo: check if we send fund back to the signer, otherwise, we'll have additional destination account
        constraint = operation.id == id @ TimelockError::InvalidId,
        constraint = operation.is_pending() @ TimelockError::OperationNotCancellable,
    )]
    pub operation: Account<'info, Operation>,

    #[account( seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
