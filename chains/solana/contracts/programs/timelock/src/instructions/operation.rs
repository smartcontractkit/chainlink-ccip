use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::HASH_BYTES;

use access_controller::AccessController;

use crate::constants::{
    ANCHOR_DISCRIMINATOR, TIMELOCK_BYPASSER_OPERATION_SEED, TIMELOCK_CONFIG_SEED,
    TIMELOCK_ID_PADDED, TIMELOCK_OPERATION_SEED,
};
use crate::error::TimelockError;
use crate::state::{Config, InstructionAccount, InstructionData, Operation, OperationState};

pub fn initialize_operation(
    op: &mut Account<Operation>,
    id: [u8; HASH_BYTES],
    predecessor: [u8; HASH_BYTES],
    salt: [u8; 32],
    instruction_count: u32,
) -> Result<()> {
    op.set_inner(Operation {
        state: OperationState::Initialized,
        timestamp: 0,
        id,
        predecessor,
        salt,
        total_instructions: instruction_count,
        instructions: Vec::with_capacity(instruction_count as usize),
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

pub fn finalize_operation(op: &mut Account<Operation>) -> Result<()> {
    op.finalize();
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
    pub config: AccountLoader<'info, Config>,

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
        realloc = operation.space_after_init_instruction(&accounts),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() < operation.total_instructions as usize @ TimelockError::TooManyInstructions,
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
        realloc = operation.space_after_append_instruction(&ix_data_chunk),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = (ix_index as usize) < operation.instructions.len() @ TimelockError::InvalidInput
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
        constraint = !operation.is_scheduled() @ TimelockError::OperationAlreadyScheduled,
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() == operation.total_instructions as usize @ TimelockError::TooManyInstructions,
        constraint = operation.verify_id() @ TimelockError::InvalidId
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
    pub config: AccountLoader<'info, Config>,

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
    pub config: AccountLoader<'info, Config>,

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
        realloc = operation.space_after_init_instruction(&accounts),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() < operation.total_instructions as usize @ TimelockError::TooManyInstructions,
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
        realloc = operation.space_after_append_instruction(&ix_data_chunk),
        realloc::payer = authority,
        realloc::zero = false,
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = (ix_index as usize) < operation.instructions.len() @ TimelockError::InvalidInput
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
        constraint = !operation.is_finalized() @ TimelockError::OperationAlreadyFinalized,
        constraint = operation.instructions.len() == operation.total_instructions as usize @ TimelockError::TooManyInstructions,
        constraint = operation.verify_id() @ TimelockError::InvalidId
    )]
    pub operation: Account<'info, Operation>,

    #[account(seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,

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
    pub config: AccountLoader<'info, Config>,

    // NOTE: access controller check and access happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
