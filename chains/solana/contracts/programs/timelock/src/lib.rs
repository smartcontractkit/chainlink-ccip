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
pub use state::*;

pub mod instructions;
use instructions::*;

#[program]
pub mod timelock {
    use bytemuck::Zeroable;

    use super::*;

    pub fn initialize(ctx: Context<Initialize>, min_delay: u64) -> Result<()> {
        initialize::initialize(ctx, min_delay)
    }

    #[access_control(only_admin!(ctx))]
    pub fn batch_add_access<'info>(
        ctx: Context<'_, '_, '_, 'info, BatchAddAccess<'info>>,
        role: Role,
    ) -> Result<()> {
        initialize::batch_add_access(ctx, role)
    }

    #[access_control(only_role_or_admin_role!(ctx, Role::Proposer))]
    pub fn schedule_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ScheduleBatch<'info>>,
        id: [u8; 32],
        delay: u64,
    ) -> Result<()> {
        schedule::schedule_batch(ctx, id, delay)
    }

    pub fn initialize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, InitializeOperation<'info>>,
        id: [u8; 32],
        predecessor: [u8; 32],
        salt: [u8; 32],
        instruction_count: u32,
    ) -> Result<()> {
        schedule::initialize_operation(ctx, id, predecessor, salt, instruction_count)
    }

    pub fn append_instructions<'info>(
        ctx: Context<'_, '_, '_, 'info, AppendInstructions<'info>>,
        id: [u8; 32],
        instructions_batch: Vec<InstructionData>,
    ) -> Result<()> {
        schedule::append_instructions(ctx, id, instructions_batch)
    }

    pub fn clear_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, ClearOperation<'info>>,
        id: [u8; 32],
    ) -> Result<()> {
        schedule::clear_operation(ctx, id)
    }

    pub fn finalize_operation<'info>(
        ctx: Context<'_, '_, '_, 'info, FinalizeOperation<'info>>,
        id: [u8; 32],
    ) -> Result<()> {
        schedule::finalize_operation(ctx, id)
    }

    #[access_control(only_role_or_admin_role!(ctx, Role::Canceller))]
    pub fn cancel<'info>(
        ctx: Context<'_, '_, '_, 'info, Cancel<'info>>,
        id: [u8; 32], // precalculated id of the tx(instructions)
    ) -> Result<()> {
        cancel::cancel(ctx, id)
    }

    #[access_control(only_role_or_admin_role!(ctx, Role::Executor))]
    pub fn execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, ExecuteBatch<'info>>,
        id: [u8; 32],
    ) -> Result<()> {
        execute::execute_batch(ctx, id)
    }

    #[access_control(only_role_or_admin_role!(ctx, Role::Bypasser))]
    pub fn bypasser_execute_batch<'info>(
        ctx: Context<'_, '_, '_, 'info, BypasserExecuteBatch<'info>>,
        id: [u8; 32],
    ) -> Result<()> {
        execute::bypasser_execute_batch(ctx, id)
    }

    #[access_control(only_admin!(ctx))]
    pub fn update_delay(ctx: Context<UpdateDelay>, delay: u64) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(delay > 0, TimelockError::InvalidInput);
        emit!(MinDelayChange {
            old_duration: config.min_delay,
            new_duration: delay,
        });
        config.min_delay = delay;

        Ok(())
    }

    #[access_control(only_admin!(ctx))]
    pub fn block_function_selector(
        ctx: Context<BlockFunctionSelector>,
        selector: [u8; 8],
    ) -> Result<()> {
        // todo: implement function blocker related methods
        emit!(FunctionSelectorBlocked { selector });
        Ok(())
    }

    #[access_control(only_admin!(ctx))]
    pub fn unblock_function_selector(
        ctx: Context<UnblockFunctionSelector>,
        selector: [u8; 8],
    ) -> Result<()> {
        // todo: implement function blocker related methods
        emit!(FunctionSelectorUnblocked { selector });
        Ok(())
    }

    #[access_control(only_admin!(ctx))]
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(proposed_owner != config.owner, TimelockError::InvalidInput);
        config.proposed_owner = proposed_owner;
        Ok(())
    }

    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        ctx.accounts.config.proposed_owner = Pubkey::zeroed();
        Ok(())
    }
}

#[derive(Accounts)]
pub struct TransferOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct AcceptOwnership<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,
    #[account(address = config.proposed_owner @ TimelockError::Unauthorized)]
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UpdateDelay<'info> {
    #[account(mut, seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct BlockFunctionSelector<'info> {
    #[account(seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}

#[derive(Accounts)]
pub struct UnblockFunctionSelector<'info> {
    #[account(seeds = [TIMELOCK_CONFIG_SEED], bump)]
    pub config: Account<'info, Config>,
    // owner(admin) only, access control with only_admin macro
    pub authority: Signer<'info>,
}
