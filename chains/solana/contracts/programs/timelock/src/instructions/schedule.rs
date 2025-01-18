use anchor_lang::prelude::*;

use access_controller::AccessController;

use crate::constants::{
    ANCHOR_DISCRIMINATOR, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED,
};
use crate::error::{AuthError, TimelockError};
use crate::event::*;
use crate::state::{Config, InstructionData, Operation};

/// this function is used to schedule the operation
/// Operation account should be preloaded with instructions and finalized before scheduling
pub fn schedule_batch<'info>(
    ctx: Context<'_, '_, '_, 'info, ScheduleBatch<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; 32],
    delay: u64,
) -> Result<()> {
    // delay should greater than min_delay
    let config = &ctx.accounts.config;
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
    op.timestamp = scheduled_time;

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

/// initialize_operation, append_instructions, finalize_operation functions are used to create a new operation
/// and add instructions to it. finalize_operation is used to mark the operation as finalized.
/// only after the operation is finalized, it can be scheduled.
/// this is due to the fact that the operation PDA cannot be initialized with CPI call from MCM program.
/// This pattern also allows to execute larger transaction(multiple instructions) exceeding 1232 bytes
/// in a single execute_batch transaction ensuring atomicy.
pub fn initialize_operation<'info>(
    ctx: Context<'_, '_, '_, 'info, InitializeOperation<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; 32],
    predecessor: [u8; 32],
    salt: [u8; 32],
    instruction_count: u32,
) -> Result<()> {
    let op = &mut ctx.accounts.operation;

    op.set_inner(Operation {
        timestamp: 0, // not scheduled operation should have timestamp 0, see src/state/operation.rs
        id, // id should be matched with hashed instructions(will be verified in finalize_operation)
        predecessor, // required for dependency check
        salt, // random salt for hash
        total_instructions: instruction_count, // total number of instructions, used for space calculation and finalization validation
        instructions: Vec::with_capacity(instruction_count as usize), // create empty vector with total instruction capacity
        is_finalized: false, // operation should be finalized before scheduling
        authority: ctx.accounts.authority.key(), // authority of the operation
    });

    Ok(())
}

/// append instructions to the operation
pub fn append_instructions<'info>(
    ctx: Context<'_, '_, '_, 'info, AppendInstructions<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; 32],
    instructions_batch: Vec<InstructionData>,
) -> Result<()> {
    let op = &mut ctx.accounts.operation;
    op.instructions.extend(instructions_batch);

    Ok(())
}

/// finalize the operation, this is required to schedule the operation
/// verify the operation status and id before mark it as finalized
pub fn finalize_operation<'info>(
    ctx: Context<'_, '_, '_, 'info, FinalizeOperation<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; 32],
) -> Result<()> {
    let op = &mut ctx.accounts.operation;
    op.is_finalized = true;

    Ok(())
}

pub fn clear_operation(
    _ctx: Context<ClearOperation>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; 32],
) -> Result<()> {
    // NOTE: ctx.accounts.operation is closed to be able to re-initialized,
    // also allow finalized operation to be cleared
    Ok(())
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; 32])]
pub struct ScheduleBatch<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        constraint = operation.is_finalized @ TimelockError::OperationNotFinalized,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled
    )]
    pub operation: Box<Account<'info, Operation>>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; 32],
    predecessor: [u8; 32],
    salt: [u8; 32],
    instruction_count: u32,
)]
pub struct InitializeOperation<'info> {
    #[account(
        init,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; 32], instructions_batch: Vec<InstructionData>)]
pub struct AppendInstructions<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR +
           Operation::INIT_SPACE +
           operation.instructions.iter().map(|ix| {  // space of existing instructions
               InstructionData::space(ix)
            }).sum::<usize>() +
               instructions_batch.iter().map(|ix| {  // add space for new instructions
               InstructionData::space(ix)
           }).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = operation.instructions.len() + instructions_batch.len() <= operation.total_instructions as usize @ TimelockError::TooManyInstructions
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; 32])]
pub struct FinalizeOperation<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = operation.instructions.len() == operation.total_instructions as usize @ TimelockError::TooManyInstructions,
        constraint = operation.verify_id() @ TimelockError::InvalidId
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; 32])]
pub struct ClearOperation<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        close = authority,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled, // restrict clearing of scheduled operation
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in only_role_or_admin_role macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
