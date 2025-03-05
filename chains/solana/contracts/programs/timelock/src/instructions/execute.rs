use anchor_lang::solana_program::instruction::Instruction;
use anchor_lang::solana_program::keccak::HASH_BYTES;
use anchor_lang::{prelude::*, solana_program};
use solana_program::program::invoke_signed;

use access_controller::AccessController;
use bytemuck::Zeroable;

use crate::constants::{
    EMPTY_PREDECESSOR, TIMELOCK_BYPASSER_OPERATION_SEED, TIMELOCK_CONFIG_SEED, TIMELOCK_ID_PADDED,
    TIMELOCK_OPERATION_SEED, TIMELOCK_SIGNER_SEED,
};
use crate::error::TimelockError;
use crate::event::*;
use crate::state::{Config, InstructionData, Operation};

/// execute scheduled operation(instructions) in batch
/// operation can be executed only if it's ready and all predecessors are done
pub fn execute_batch<'info>(
    ctx: Context<'_, '_, '_, 'info, ExecuteBatch<'info>>,
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    let op = &mut ctx.accounts.operation;

    // check if the operation is ready
    let current_time = Clock::get()?.unix_timestamp as u64;
    require!(op.is_ready(current_time), TimelockError::OperationNotReady);

    if op.predecessor != EMPTY_PREDECESSOR {
        // force predecessor to be provided
        let (expected_address, _) = Pubkey::find_program_address(
            &[
                TIMELOCK_OPERATION_SEED,
                timelock_id.as_ref(),
                op.predecessor.as_ref(),
            ],
            ctx.program_id,
        );

        require_keys_eq!(
            ctx.accounts.predecessor_operation.key(),
            expected_address,
            TimelockError::InvalidInput
        );

        let predecessor_data = ctx.accounts.predecessor_operation.try_borrow_data()?;
        let mut predecessor_data_slice: &[u8] = &predecessor_data;
        let predecessor_acc = Operation::try_deserialize(&mut predecessor_data_slice)?;

        require!(predecessor_acc.is_done(), TimelockError::MissingDependency);
    } else {
        require_keys_eq!(
            ctx.accounts.predecessor_operation.key(),
            Pubkey::zeroed(),
            TimelockError::InvalidInput
        );
    }

    let seeds = &[
        TIMELOCK_SIGNER_SEED,
        timelock_id.as_ref(),
        &[ctx.bumps.timelock_signer],
    ];
    let signer = &[&seeds[..]];

    for (i, instruction_data) in op.instructions.iter().enumerate() {
        execute(
            instruction_data,
            ctx.remaining_accounts,
            signer,
            ctx.accounts.timelock_signer.key(),
        )?;

        emit!(CallExecuted {
            id: op.id,
            index: i as u64,
            target: instruction_data.program_id,
            data: instruction_data.data.clone(),
        });
    }

    // all executed, update the operation state
    op.mark_done();

    Ok(())
}

/// execute operation(instructions) w/o checking predecessors and readiness
/// bypasser_execute also need the operation to be uploaded formerly
/// NOTE: operation is closed after execution
pub fn bypasser_execute_batch<'info>(
    ctx: Context<'_, '_, '_, 'info, BypasserExecuteBatch<'info>>,
    timelock_id: [u8; TIMELOCK_ID_PADDED],
    _id: [u8; HASH_BYTES],
) -> Result<()> {
    let op = &mut ctx.accounts.operation;

    let seeds = &[
        TIMELOCK_SIGNER_SEED,
        timelock_id.as_ref(),
        &[ctx.bumps.timelock_signer],
    ];
    let signer = &[&seeds[..]];

    for (i, instruction_data) in op.instructions.iter().enumerate() {
        execute(
            instruction_data,
            ctx.remaining_accounts,
            signer,
            ctx.accounts.timelock_signer.key(),
        )?;
        emit!(BypasserCallExecuted {
            index: i as u64,
            target: instruction_data.program_id,
            data: instruction_data.data.clone(),
        });
    }

    Ok(())
}

/// Executes each instruction with the timelock signer as the signing authority.
///
/// This function is the core execution mechanism for both standard and bypasser operations.
/// It performs the critical replacement of signers in the instruction, ensuring that:
/// 1. The timelock signer PDA becomes the only valid signer in the CPI
/// 2. Original signer designations are preserved only for the timelock PDA
/// 3. The instruction is executed with the proper signing authority
///
/// # Parameters
///
/// - `instruction`: The instruction data to be executed
/// - `remaining_accounts`: All accounts needed for the instruction execution
/// - `signer_seeds`: Seeds used to derive and sign with the PDA
/// - `timelock_signer`: Public key of the timelock signer PDA
///
/// # Note
///
/// This signer substitution is critical for security - it ensures that only
/// properly authorized and verified operations can execute with the
/// timelock's signing authority.
fn execute(
    instruction: &InstructionData,
    remaining_accounts: &[AccountInfo],
    signer_seeds: &[&[&[u8]]],
    timelock_signer: Pubkey,
) -> Result<()> {
    let mut ix: Instruction = instruction.into();

    ix.accounts = ix
        .accounts
        .iter()
        .map(|acc| {
            let mut acc = acc.clone();
            acc.is_signer = acc.pubkey == timelock_signer;
            acc
        })
        .collect();

    invoke_signed(&ix, remaining_accounts, signer_seeds)?;
    Ok(())
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct ExecuteBatch<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        constraint = operation.is_scheduled() @ TimelockError::InvalidId,
    )]
    pub operation: Account<'info, Operation>,

    /// CHECK: Will be validated in handler if predecessor exists
    pub predecessor_operation: UncheckedAccount<'info>,

    #[account( seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref(),], bump)]
    pub config: AccountLoader<'info, Config>,

    /// CHECK: program signer PDA that can hold balance
    #[account(
        seeds = [TIMELOCK_SIGNER_SEED, timelock_id.as_ref()],
        bump
    )]
    pub timelock_signer: UncheckedAccount<'info>,

    // NOTE: access controller check happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED], id: [u8; HASH_BYTES])]
pub struct BypasserExecuteBatch<'info> {
    #[account(
        mut,
        seeds = [TIMELOCK_BYPASSER_OPERATION_SEED, timelock_id.as_ref(), id.as_ref()],
        bump,
        constraint = operation.is_finalized() @ TimelockError::OperationNotFinalized,
        close = authority, // close the operation after bypasser execution
    )]
    pub operation: Account<'info, Operation>,

    #[account( seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref(),], bump)]
    pub config: AccountLoader<'info, Config>,

    /// CHECK: program signer PDA that can hold balance
    #[account(
        seeds = [TIMELOCK_SIGNER_SEED, timelock_id.as_ref()],
        bump
    )]
    pub timelock_signer: UncheckedAccount<'info>,

    // NOTE: access controller check happens in require_role_or_admin macro
    pub role_access_controller: AccountLoader<'info, AccessController>,

    #[account(mut)]
    pub authority: Signer<'info>,
}
