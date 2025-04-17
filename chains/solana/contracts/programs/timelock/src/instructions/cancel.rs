use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::HASH_BYTES;

use access_controller::AccessController;

use crate::constants::{TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED};
use crate::error::TimelockError;
use crate::event::*;
use crate::state::{Config, Operation};

/// cancels a pending operation
pub fn cancel<'info>(
    _ctx: Context<'_, '_, '_, 'info, Cancel<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; HASH_BYTES],
) -> Result<()> {
    emit!(Cancelled { id });
    // NOTE: PDA is closed - is handled by anchor on exit due to the `close` attribute
    Ok(())
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct Cancel<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        close = authority,
        constraint = operation.id == id @ TimelockError::InvalidId,
        constraint = operation.is_pending() @ TimelockError::OperationNotCancellable,
    )]
    pub operation: Account<'info, Operation>,

    #[account( seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

    // NOTE: access controller check happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
