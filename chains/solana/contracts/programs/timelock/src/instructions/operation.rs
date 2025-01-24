use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::HASH_BYTES;

use access_controller::AccessController;

use crate::constants::{
    ANCHOR_DISCRIMINATOR, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED,
};
use crate::error::TimelockError;
use crate::state::{Config, InstructionData, Operation};
use crate::TIMELOCK_BYPASSER_OPERATION_SEED;

/// Operation management for timelock system, handling both standard (timelock-enforced)
/// and bypass (emergency) operations.
///
/// Standard Operation Flow:
/// - initialize -> append -> finalize -> schedule -> execute_batch
/// - Enforces timelock delays and predecessor dependencies
///
/// Bypass Operation Flow:
/// - initialize -> append -> finalize -> bypass_execute_batch
/// - No delay, immediate execution
///
/// Implementation uses separate code paths and PDAs for each operation type
/// to maintain clear security boundaries and audit trails, despite similar logic.
/// All operations enforce state transitions, size limits, and role-based access.
///
pub fn initialize_operation<'info>(
    ctx: Context<'_, '_, '_, 'info, InitializeOperation<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; HASH_BYTES],
    predecessor: [u8; HASH_BYTES],
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
    });

    Ok(())
}

/// append instructions to the operation
pub fn append_instructions<'info>(
    ctx: Context<'_, '_, '_, 'info, AppendInstructions<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
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
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    let op = &mut ctx.accounts.operation;
    op.is_finalized = true;

    Ok(())
}

pub fn clear_operation(
    _ctx: Context<ClearOperation>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    // NOTE: ctx.accounts.operation is closed to be able to re-initialized,
    // also allow finalized operation to be cleared
    Ok(())
}

pub fn initialize_bypasser_operation<'info>(
    ctx: Context<'_, '_, '_, 'info, InitializeBypasserOperation<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; HASH_BYTES],
    salt: [u8; 32],
    instruction_count: u32,
) -> Result<()> {
    let op = &mut ctx.accounts.operation;

    op.set_inner(Operation {
        timestamp: 0, // not scheduled operation should have timestamp 0, see src/state/operation.rs
        id, // id should be matched with hashed instructions(will be verified in finalize_operation)
        predecessor: [0; 32], // required for dependency check
        salt, // random salt for hash
        total_instructions: instruction_count, // total number of instructions, used for space calculation and finalization validation
        instructions: Vec::with_capacity(instruction_count as usize), // create empty vector with total instruction capacity
        is_finalized: false, // operation should be finalized before scheduling
    });

    Ok(())
}

/// append instructions to the operation
pub fn append_bypasser_instructions<'info>(
    ctx: Context<'_, '_, '_, 'info, AppendBypasserInstructions<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
    instructions_batch: Vec<InstructionData>,
) -> Result<()> {
    let op = &mut ctx.accounts.operation;
    op.instructions.extend(instructions_batch);

    Ok(())
}

/// finalize the operation, this is required to schedule the operation
/// verify the operation status and id before mark it as finalized
pub fn finalize_bypasser_operation<'info>(
    ctx: Context<'_, '_, '_, 'info, FinalizeBypasserOperation<'info>>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    let op = &mut ctx.accounts.operation;
    op.is_finalized = true;

    Ok(())
}

pub fn clear_bypasser_operation(
    _ctx: Context<ClearBypasserOperation>,
    _timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    // NOTE: ctx.accounts.operation is closed to be able to re-initialized,
    // also allow finalized operation to be cleared
    Ok(())
}

#[derive(Accounts)]
#[instruction(
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; HASH_BYTES],
    predecessor: [u8; HASH_BYTES],
    salt: [u8; HASH_BYTES],
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

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES], instructions_batch: Vec<InstructionData>)]
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

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
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

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
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

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    id: [u8; HASH_BYTES],
    salt: [u8; HASH_BYTES],
    instruction_count: u32,
)]
pub struct InitializeBypasserOperation<'info> {
    #[account(
        init,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE,
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES], instructions_batch: Vec<InstructionData>)]
pub struct AppendBypasserInstructions<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
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
        constraint = operation.instructions.len() + instructions_batch.len() <= operation.total_instructions as usize @ TimelockError::TooManyInstructions
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct FinalizeBypasserOperation<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() == operation.total_instructions as usize @ TimelockError::TooManyInstructions,
        constraint = operation.verify_id() @ TimelockError::InvalidId
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct ClearBypasserOperation<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        close = authority,
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
