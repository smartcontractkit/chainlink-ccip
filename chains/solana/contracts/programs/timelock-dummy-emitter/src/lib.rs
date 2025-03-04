use anchor_lang::prelude::*;

pub mod events;

use events::*;

declare_id!("8hNnreBcZRQgcWnKaEkzsQwZA6B9ngjXFhSkToVu8V67");

#[program]
pub mod timelock_dummy_emitter {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // 1. Emit CallScheduled
        emit!(CallScheduled {
            id: [0u8; 32],
            index: 1,
            target: Pubkey::default(),
            predecessor: [0u8; 32],
            salt: [0u8; 32],
            delay: 3600,
            data: vec![0u8; 32],
        });

        // 2. Emit CallExecuted
        emit!(CallExecuted {
            id: [0u8; 32],
            index: 1,
            target: Pubkey::default(),
            data: vec![0u8; 32],
        });

        // 3. Emit BypasserCallExecuted
        emit!(BypasserCallExecuted {
            index: 1,
            target: Pubkey::default(),
            data: vec![0u8; 32],
        });

        // 4. Emit Cancelled
        emit!(Cancelled {
            id: [0u8; 32],
        });

        // 5. Emit MinDelayChange
        emit!(MinDelayChange {
            old_duration: 3600,
            new_duration: 7200,
        });

        // 6. Emit Function Selector events
        let selector = [0u8; 8];
        emit!(FunctionSelectorBlocked { selector });
        emit!(FunctionSelectorUnblocked { selector });

        msg!("All events have been emitted!");
        Ok(())
    }

    /// Emit events and reverts. Will mark transaction as failed.
    pub fn trigger_all_events_reverts(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        trigger_all_events(_ctx)?;
        Err(error!(CustomError::IntentionalRevert))
    }
}

#[derive(Accounts)]
pub struct TriggerAllEvents {}

#[error_code]
pub enum CustomError {
    #[msg("Transaction intentionally reverted for testing")]
    IntentionalRevert,
}