use anchor_lang::prelude::*;

declare_id!("LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn");

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
pub mod timelock {
    use bytemuck::Zeroable;

    use super::*;

    pub fn initialize(
        ctx: Context<Initialize>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        min_delay: u64,
    ) -> Result<()> {
        initialize::initialize(ctx, timelock_id, min_delay)
    }

    #[access_control(require_only_admin!(ctx))]
    pub fn batch_add_access<'info>(
        ctx: Context<'_, '_, '_, 'info, BatchAddAccess<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        role: Role,
    ) -> Result<()> {
        initialize::batch_add_access(ctx, timelock_id, role)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn initialize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        predecessor: [u8; 32],
        salt: [u8; 32],
        instruction_count: u32,
    ) -> Result<()> {
        operation::initialize_operation(ctx, timelock_id, id, predecessor, salt, instruction_count)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn append_instructions<'info>(
        ctx: Context<'_, '_, '_, 'info, AppendInstructions<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        instructions_batch: Vec<InstructionData>,
    ) -> Result<()> {
        operation::append_instructions(ctx, timelock_id, id, instructions_batch)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn finalize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, FinalizeOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        operation::finalize_operation(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn clear_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, ClearOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        operation::clear_operation(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Proposer))]
    pub fn schedule_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ScheduleBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        delay: u64,
    ) -> Result<()> {
        schedule::schedule_batch(ctx, timelock_id, id, delay)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Canceller))]
    pub fn cancel<'info>(
        ctx: Context<'_, '_, '_, 'info, Cancel<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32], // precalculated id of the tx(instructions)
    ) -> Result<()> {
        cancel::cancel(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Executor))]
    pub fn execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ExecuteBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        execute::execute_batch(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn initialize_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeBypasserOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        salt: [u8; 32],
        instruction_count: u32,
    ) -> Result<()> {
        operation::initialize_bypasser_operation(ctx, timelock_id, id, salt, instruction_count)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn append_bypasser_instructions<'info>(
        ctx: Context<'_, '_, '_, 'info, AppendBypasserInstructions<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
        instructions_batch: Vec<InstructionData>,
    ) -> Result<()> {
        operation::append_bypasser_instructions(ctx, timelock_id, id, instructions_batch)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn finalize_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, FinalizeBypasserOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        operation::finalize_bypasser_operation(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn clear_bypasser_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, ClearBypasserOperation<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        operation::clear_bypasser_operation(ctx, timelock_id, id)
    }

    #[access_control(require_role_or_admin!(ctx, Role::Bypasser))]
    pub fn bypasser_execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, BypasserExecuteBatch<'info>>,
        timelock_id: [u8; TIMELOCK_ID_PADDED],
        id: [u8; 32],
    ) -> Result<()> {
        execute::bypasser_execute_batch(ctx, timelock_id, id)
    }

    #[access_control(require_only_admin!(ctx))]
    pub fn update_delay(
        ctx: Context<UpdateDelay>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        delay: u64,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(delay > 0, TimelockError::InvalidInput);
        emit!(MinDelayChange {
            old_duration: config.min_delay,
            new_duration: delay,
        });
        config.min_delay = delay;
        Ok(())
    }

    #[access_control(require_only_admin!(ctx))]
    pub fn block_function_selector(
        ctx: Context<BlockFunctionSelector>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        selector: [u8; 8],
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.blocked_selectors.block_selector(selector)?;
        emit!(FunctionSelectorBlocked { selector });
        Ok(())
    }

    #[access_control(require_only_admin!(ctx))]
    pub fn unblock_function_selector(
        ctx: Context<UnblockFunctionSelector>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        selector: [u8; 8],
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.blocked_selectors.unblock_selector(selector)?;
        emit!(FunctionSelectorUnblocked { selector });
        Ok(())
    }

    #[access_control(require_only_admin!(ctx))]
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(proposed_owner != config.owner, TimelockError::InvalidInput);
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(
        ctx: Context<AcceptOwnership>,
        _timelock_id: [u8; TIMELOCK_ID_PADDED],
    ) -> Result<()> {
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        ctx.accounts.config.proposed_owner = Pubkey::zeroed();
        Ok(())
    }
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct TransferOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct AcceptOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,
    #[account(address = config.proposed_owner @ AuthError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct UpdateDelay<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct BlockFunctionSelector<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(timelock_id: [u8; TIMELOCK_ID_PADDED])]
pub struct UnblockFunctionSelector<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED, timelock_id.as_ref()], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}
