use anchor_lang::prelude::*;

declare_id!("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA");

mod constants;
pub use constants::*;

pub mod access;
pub use access::*;

pub mod error;
pub use error::*;

pub mod event;
pub use event::*;

pub mod state;
pub use state::{config::*, operation::*};

pub mod instructions;
use instructions::*;

#[program]
/// The `timelock` program provides a mechanism to schedule, execute, and (if needed) cancel
/// operations in a delayed fashion. It supports both standard operations (which enforce delays and dependencies)
/// and bypass operations (for emergency cases).
///
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
pub mod timelock {
    #![warn(missing_docs)]
    use super::*;

    /// Initialize the timelock configuration.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the accounts required for initialization.
    /// - `timelock_id`: A unique, padded identifier for this timelock instance.
    /// - `min_delay`: The minimum delay (in seconds) required for scheduled operations.
    pub fn initialize(
        ctx: Context<Initialize>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        min_delay: u64,
    ) -> Result<()> {
        initialize::initialize(ctx, timelock_id, min_delay)
    }

    /// Add a new access role in batch. Only the admin is allowed to perform this operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the accounts required for batch adding access.
    /// - `timelock_id`: A unique, padded identifier for this timelock instance.
    /// - `role`: The role to be added.
    #[access_control(require_only_admin!(ctx))]
    pub fn batch_add_access<'info>(
        ctx: Context<'_, '_, '_, 'info, BatchAddAccess<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        role: Role,
    ) -> Result<()> {
        initialize::batch_add_access(ctx, timelock_id, role)
    }

    //=== Standard Timelock Operation Instructions ===

    /// Initialize a new standard timelock operation.
    ///
    /// This sets up a new operation with the given ID, predecessor, salt, and expected number of instructions.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the operation account.
    /// - `_timelock_id`: A padded identifier for the timelock (unused here but required for PDA derivation).
    /// - `id`: The unique identifier for the operation.
    /// - `predecessor`: The identifier of the predecessor operation.
    /// - `salt`: A salt value to help create unique PDAs.
    /// - `instruction_count`: The total number of instructions that will be added to this operation.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn initialize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        predecessor: [u8; 32],
        salt: [u8; 32],
        instruction_count: u32,
    ) -> Result<()> {
        operation::initialize_operation(
            &mut ctx.accounts.operation,
            id,
            predecessor,
            salt,
            instruction_count,
        )
    }

    /// Append a new instruction to an existing standard operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the operation account.
    /// - `_timelock_id`: The timelock identifier (for PDA derivation).
    /// - `_id`: The operation identifier.
    /// - `program_id`: The target program for the instruction.
    /// - `accounts`: The list of accounts required for the instruction.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn initialize_instruction<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeInstruction<'info>>,
        _timelock_id: [u8; 32],
        _id: [u8; 32],
        program_id: Pubkey,
        accounts: Vec<InstructionAccount>,
    ) -> Result<()> {
        operation::initialize_instruction(&mut ctx.accounts.operation, program_id, accounts)
    }

    /// Append additional instruction data to an instruction of an existing standard operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the operation account.
    /// - `_timelock_id`: The timelock identifier (for PDA derivation).
    /// - `_id`: The operation identifier.
    /// - `ix_index`: The index of the instruction to which the data will be appended.
    /// - `ix_data_chunk`: A chunk of data to be appended.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn append_instruction_data<'info>(
        ctx: Context<'_, '_, '_, 'info, AppendInstructionData<'info>>,
        _timelock_id: [u8; 32],
        _id: [u8; 32],
        ix_index: u32,
        ix_data_chunk: Vec<u8>,
    ) -> Result<()> {
        operation::append_instruction_data(&mut ctx.accounts.operation, ix_index, ix_data_chunk)
    }

    /// Finalize a standard operation.
    ///
    /// Finalizing an operation marks it as ready for scheduling.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the operation account.
    /// - `_timelock_id`: The timelock identifier (for PDA derivation).
    /// - `_id`: The operation identifier.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn finalize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, FinalizeOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        _id: [u8; 32],
    ) -> Result<()> {
        operation::finalize_operation(&mut ctx.accounts.operation)
    }

    /// Clear an operation that has been finalized.
    ///
    /// This effectively closes the operation account so that it can be reinitialized later.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `_id`: The operation identifier.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn clear_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, ClearOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        _id: [u8; 32],
    ) -> Result<()> {
        // NOTE: ctx.accounts.operation is closed to be able to re-initialized, finalized operations are allowed to be cleared as well.
        Ok(())
    }

    /// Schedule a finalized operation to be executed after a delay.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the accounts for scheduling.
    /// - `timelock_id`: The timelock identifier.
    /// - `id`: The operation identifier.
    /// - `delay`: The delay (in seconds) before the operation can be executed.
    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn schedule_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ScheduleBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        delay: u64,
    ) -> Result<()> {
        schedule::schedule_batch(ctx, timelock_id, id, delay)
    }

    /// Cancel a scheduled operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the accounts for cancellation.
    /// - `timelock_id`: The timelock identifier.
    /// - `id`: The operation identifier (precalculated).
    #[access_control(require_role_or_admin!(ctx, Role::Canceller))]
    pub fn cancel<'info>(
        ctx: Context<'_, '_, '_, 'info, Cancel<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32], // precalculated id of the tx(instructions)
    ) -> Result<()> {
        cancel::cancel(ctx, timelock_id, id)
    }

    /// Executes a scheduled batch of operations after validating readiness and predecessor dependencies.
    ///
    /// This function:
    /// 1. Verifies the operation is ready for execution (delay period has passed)
    /// 2. Validates that any predecessor operation has been completed
    /// 3. Executes each instruction in the operation using the timelock signer PDA
    /// 4. Emits events for each executed instruction
    ///
    /// # Parameters
    ///
    /// - `ctx`: Context containing operation accounts and signer information
    /// - `timelock_id`: Identifier for the timelock instance
    /// - `_id`: Operation ID (used for PDA derivation)
    ///
    /// # Security Considerations
    ///
    /// This instruction uses PDA signing to create a trusted execution environment.
    /// The timelock's signer PDA will replace any account marked as a signer in the
    /// original instructions, providing the necessary privileges while maintaining
    /// security through program derivation.
    #[access_control(require_role_or_admin!(ctx, Role::Executor))]
    pub fn execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ExecuteBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        execute::execute_batch(ctx, timelock_id, id)
    }

    //=== Bypasser Operation Instructions ===

    /// Initialize a bypasser operation.
    ///
    /// Bypasser operations have no predecessor and can be executed without delay.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the bypasser operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `id`: The operation identifier.
    /// - `salt`: A salt value for PDA derivation.
    /// - `instruction_count`: The number of instructions to be added.
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn initialize_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeBypasserOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        salt: [u8; 32],
        instruction_count: u32,
    ) -> Result<()> {
        operation::initialize_operation(
            &mut ctx.accounts.operation,
            id,
            [0; 32], // bypasser operations do not have predecessors
            salt,
            instruction_count,
        )
    }

    /// Initialize an instruction for a bypasser operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the bypasser operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `_id`: The operation identifier.
    /// - `program_id`: The target program for the instruction.
    /// - `accounts`: The list of accounts required for the instruction.
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn initialize_bypasser_instruction<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeBypasserInstruction<'info>>,
        _timelock_id: [u8; 32],
        _id: [u8; 32],
        program_id: Pubkey,
        accounts: Vec<InstructionAccount>,
    ) -> Result<()> {
        operation::initialize_instruction(&mut ctx.accounts.operation, program_id, accounts)
    }

    /// Append additional data to an instruction of a bypasser operation.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the bypasser operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `_id`: The operation identifier.
    /// - `ix_index`: The index of the instruction.
    /// - `ix_data_chunk`: The data to append.
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn append_bypasser_instruction_data<'info>(
        ctx: Context<'_, '_, '_, 'info, AppendBypasserInstructionData<'info>>,
        _timelock_id: [u8; 32],
        _id: [u8; 32],
        ix_index: u32,
        ix_data_chunk: Vec<u8>,
    ) -> Result<()> {
        operation::append_instruction_data(&mut ctx.accounts.operation, ix_index, ix_data_chunk)
    }

    /// Finalize a bypasser operation.
    ///
    /// Marks the bypasser operation as finalized, ready for execution.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the bypasser operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `_id`: The operation identifier.
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn finalize_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, FinalizeBypasserOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        _id: [u8; 32],
    ) -> Result<()> {
        operation::finalize_operation(&mut ctx.accounts.operation)
    }

    /// Clear a finalized bypasser operation.
    ///
    /// Closes the bypasser operation account.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the bypasser operation account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `_id`: The operation identifier.
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn clear_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, ClearBypasserOperation<'info>>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        _id: [u8; 32],
    ) -> Result<()> {
        // NOTE: ctx.accounts.operation is closed to be able to re-initialized, finalized operations are allowed to be cleared as well.
        Ok(())
    }

    /// Execute operations immediately using the bypasser flow, bypassing time delays
    /// and predecessor checks.
    ///
    /// This function provides an emergency execution mechanism that:
    /// 1. Skips the timelock waiting period required for standard operations
    /// 2. Does not enforce predecessor dependencies
    /// 3. Closes the operation account after execution
    ///
    /// # Emergency Use Only
    ///
    /// The bypasser flow is intended strictly for emergency situations where
    /// waiting for the standard timelock delay would cause harm. Access to this
    /// function is tightly controlled through the Bypasser role.
    ///
    /// # Parameters
    ///
    /// - `ctx`: Context containing operation accounts and signer information
    /// - `timelock_id`: Identifier for the timelock instance
    /// - `_id`: Operation ID (used for PDA derivation)
    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn bypasser_execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, BypasserExecuteBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        execute::bypasser_execute_batch(ctx, timelock_id, id)
    }

    //=== Admin Instructions ===

    /// Update the minimum delay required for scheduled operations.
    ///
    /// Only the admin can update the delay.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `delay`: The new minimum delay value.
    #[access_control(require_only_admin!(ctx))]
    pub fn update_delay(
        ctx: Context<UpdateDelay>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        delay: u64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        require!(delay > 0, TimelockError::InvalidInput);
        emit!(MinDelayChange {
            old_duration: config.min_delay,
            new_duration: delay,
        });
        config.min_delay = delay;
        Ok(())
    }

    /// Block a function selector from being called.
    ///
    /// Only the admin can block function selectors.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `selector`: The 8-byte function selector(Anchor discriminator) to block.
    #[access_control(require_only_admin!(ctx))]
    pub fn block_function_selector(
        ctx: Context<BlockFunctionSelector>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        selector: [u8; 8],
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        config.blocked_selectors.block_selector(selector)?;
        emit!(FunctionSelectorBlocked { selector });
        Ok(())
    }

    /// Unblock a previously blocked function selector.
    ///
    /// Only the admin can unblock function selectors.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `selector`: The function selector to unblock.
    #[access_control(require_only_admin!(ctx))]
    pub fn unblock_function_selector(
        ctx: Context<UnblockFunctionSelector>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        selector: [u8; 8],
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        config.blocked_selectors.unblock_selector(selector)?;
        emit!(FunctionSelectorUnblocked { selector });
        Ok(())
    }

    /// Propose a new owner for the timelock instance config.
    ///
    /// Only the current owner (admin) can propose a new owner.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_timelock_id`: The timelock identifier.
    /// - `proposed_owner`: The public key of the proposed new owner.
    #[access_control(require_only_admin!(ctx))]
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        require!(
            proposed_owner != config.owner && proposed_owner != Pubkey::default(),
            TimelockError::InvalidInput
        );
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    /// Accept ownership of the timelock config.
    ///
    /// The proposed new owner must call this function to assume ownership.
    ///
    /// # Parameters
    ///
    /// - `ctx`: The context containing the configuration account.
    /// - `_timelock_id`: The timelock identifier.
    pub fn accept_ownership(
        ctx: Context<AcceptOwnership>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;
        // NOTE: take() resets proposed_owner to default
        config.owner = std::mem::take(&mut config.proposed_owner);
        Ok(())
    }
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct TransferOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct AcceptOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,
    #[account(address = config.load()?.proposed_owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct UpdateDelay<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct BlockFunctionSelector<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct UnblockFunctionSelector<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: AccountLoader<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}
