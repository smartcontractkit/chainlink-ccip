use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::HASH_BYTES;

use access_controller::AccessController;

use crate::constants::{
    ANCHOR_DISCRIMINATOR, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED, TIMELOCK_BYPASSER_OPERATION_SEED
};
use crate::error::TimelockError;
use crate::state::{Config, InstructionData, InstructionAccount, Operation};

/// Operation management for timelock system, handling both standard (timelock-enforced)
/// and bypass (emergency) operations.
///
/// Standard Operation Flow:
/// - initialize -> append(init_ix, append_ix_data) -> finalize -> schedule -> execute_batch
/// - Enforces timelock delays and predecessor dependencies
///
/// Bypass Operation Flow:
/// - initialize -> append(init_ix, append_ix_data) -> finalize -> bypass_execute_batch
/// - No required delay or additional checks, closes operation account after execution
///
/// Implementation uses separate code paths and PDAs for each operation type
/// to maintain clear security boundaries and audit trails, despite similar logic.
/// All operations enforce state transitions, size limits, and role-based access.
///
pub fn initialize_operation(
    op: &mut Account<Operation>,
    id: [u8; HASH_BYTES],
    predecessor: [u8; HASH_BYTES],
    salt: [u8; 32],
    instruction_count: u32,
) -> Result<()> {
    op.set_inner(Operation {
        timestamp: 0,
        id,
        predecessor,
        salt,
        total_instructions: instruction_count,
        instructions: Vec::with_capacity(instruction_count as usize),
        is_finalized: false,
    });

    Ok(())
}

pub fn initialize_instruction(
    op: &mut Account<Operation>,
    program_id: Pubkey,
    accounts: Vec<InstructionAccount>,
) -> Result<()> {
    op.instructions.push(InstructionData {
        program_id,
        data: Vec::new(),
        accounts,
    });
    Ok(())
}

pub fn append_instruction_data(
    op: &mut Account<Operation>,
    ix_index: u32,
    ix_data_chunk: Vec<u8>,
) -> Result<()> {
    let target_ix = &mut op.instructions[ix_index as usize];
    target_ix.data.extend_from_slice(&ix_data_chunk);

    // NOTE: finalization of each instruction is verified in finalize_operation with verify_id
    Ok(())
}

pub fn finalize_operation(
    op: &mut Account<Operation>,
) -> Result<()> {
    op.is_finalized = true;
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
#[instruction(
    timelock_id: [u8; 32],
    id: [u8; 32],
    program_id: Pubkey,
    accounts: Vec<InstructionAccount>,
)]
pub struct InitializeInstruction<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE
            + operation.instructions.iter().map(|ix| 
                // program_id + 4 + data.len + 4 + accounts.len*34
                32 + 4 + ix.data.len() + 4 + ix.accounts.len() * 34
            ).sum::<usize>()
            // program_id (32) + data Vec overhead (4) + accounts vec overhead (4) + each account is (32 + 1 + 1) = 34
            + 32 + 4 + 4 + (accounts.len() * 34),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = operation.instructions.len() < operation.total_instructions as usize @ TimelockError::TooManyInstructions,
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
#[instruction(
    timelock_id: [u8; 32],
    id: [u8; 32],
    ix_index: u32,
    ix_data_chunk: Vec<u8>,
)]
pub struct AppendInstructionData<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE
            + operation.instructions.iter().enumerate().map(|(i,ix)| {
                if i == ix_index as usize {
                    // old size + ix_data_chunk
                    ix.data.len() + ix_data_chunk.len() +  (32 + 4 + 4 + ix.accounts.len()*34)
                } else {
                    // program_id + 4 + data.len + 4 + accounts.len*34
                    32 + 4 + ix.data.len() + 4 + ix.accounts.len() * 34
                }
            }).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false,

        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = (ix_index as usize) < operation.instructions.len() @ TimelockError::InvalidInput
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
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled, // restrict clearing of scheduled operation
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
#[instruction(
    timelock_id: [u8; 32],
    id: [u8; 32],
    program_id: Pubkey,
    accounts: Vec<InstructionAccount>,
)]
pub struct InitializeBypasserInstruction<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE
            + operation.instructions.iter().map(|ix| 
                // program_id + 4 + data.len + 4 + accounts.len*34
                32 + 4 + ix.data.len() + 4 + ix.accounts.len() * 34
            ).sum::<usize>()
            // program_id (32) + data Vec overhead (4) + accounts vec overhead (4) plus each account is (32 + 1 + 1) = 34
            + 32 + 4 + 4 + (accounts.len() * 34),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() < operation.total_instructions as usize @ TimelockError::TooManyInstructions,
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
#[instruction(
    timelock_id: [u8; 32],
    id: [u8; 32],
    ix_index: u32,
    ix_data_chunk: Vec<u8>,
)]
pub struct AppendBypasserInstructionData<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        realloc = ANCHOR_DISCRIMINATOR + Operation::INIT_SPACE
            + operation.instructions.iter().enumerate().map(|(i,ix)| {
                if i == ix_index as usize {
                    // old size + ix_data_chunk
                    ix.data.len() + ix_data_chunk.len() +  (32 + 4 + 4 + ix.accounts.len()*34)
                } else {
                    // program_id + 4 + data.len + 4 + accounts.len*34
                    32 + 4 + ix.data.len() + 4 + ix.accounts.len() * 34
                }
            }).sum::<usize>(),
        realloc::payer = authority,
        realloc::zero = false,

        constraint = !operation.is_finalized @ TimelockError::OperationAlreadyFinalized,
        constraint = (ix_index as usize) < operation.instructions.len() @ TimelockError::InvalidInput
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
