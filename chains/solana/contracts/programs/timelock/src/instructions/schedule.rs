use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::HASH_BYTES;

use access_controller::AccessController;

use crate::constants::{
    ANCHOR_DISCRIMINATOR, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED,
};
use crate::error::TimelockError;
use crate::event::*;
use crate::state::{Config, Operation};

/// this function is used to schedule the operation
/// Operation account should be preloaded with instructions and finalized before scheduling
pub fn schedule_batch<'info>(
    ctx: Context<'_, '_, '_, 'info, ScheduleBatch<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
    delay: u64,
) -> Result<()> {
    // delay should greater than min_delay
    let config = ctx.accounts.config.load()?;
    require!(delay >= config.min_delay, TimelockError::DelayInsufficient);

    let current_time = Clock::get()?.unix_timestamp as u64;
    let scheduled_time = current_time
        .checked_add(delay)
        .ok_or(TimelockError::Overflow)?;

    require!(
        scheduled_time > current_time && scheduled_time < u64::MAX,
        TimelockError::Overflow
    );

    let op = &mut ctx.accounts.operation;
    op.schedule(scheduled_time);

    for (i, ix) in op.instructions.iter().enumerate() {
        // check for 8-byte Anchor discriminator (since our internally developed programs use Anchor)
        // non-Anchor instructions may have different/empty data formats and are allowed
        if ix.data.len() >= ANCHOR_DISCRIMINATOR {
            // extract the first 8 bytes from ix.data as the selector
            let selector: [u8; ANCHOR_DISCRIMINATOR] =
                ix.data[..ANCHOR_DISCRIMINATOR].try_into().unwrap();

            if config.blocked_selectors.is_blocked(&selector) {
                return err!(TimelockError::BlockedSelector);
            }
        }

        emit!(CallScheduled {
            id: op.id,
            index: i as u64,
            target: ix.program_id,
            predecessor: op.predecessor,
            salt: op.salt,
            delay,
            data: ix.data.clone(),
        });
    }

    Ok(())
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct ScheduleBatch<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        // NOTE: order matters here, we need to check if the operation is already scheduled before checking if it is finalized
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = operation.is_finalized() @ TimelockError::OperationNotFinalized,
    )]
    pub operation: Box<Account<'info, Operation>>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
