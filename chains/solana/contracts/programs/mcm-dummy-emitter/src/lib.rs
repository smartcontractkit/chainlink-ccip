use anchor_lang::prelude::*;

pub mod events;
pub mod state;

use events::*;
use state::*;

declare_id!("EqaAbT4NkoDU7WeKTHK9DrJEZ6xgSmZzoufpZQ7GPQE6");

#[program]
pub mod mcm_dummy_emitter {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // 1. Emit NewRoot event
        emit!(NewRoot {
            root: [0u8; 32],
            valid_until: 1677777777,
            metadata_chain_id: 1,
            metadata_multisig: Pubkey::default(),
            metadata_pre_op_count: 0,
            metadata_post_op_count: 1,
            metadata_override_previous_root: false,
        });

        // 2. Emit ConfigSet event
        let signers = vec![McmSigner {
            evm_address: [0u8; 20],
            index: 0,
            group: 0,
        }];

        emit!(ConfigSet {
            group_parents: [0u8; NUM_GROUPS],
            group_quorums: [1u8; NUM_GROUPS],
            is_root_cleared: true,
            signers,
        });

        // 3. Emit OpExecuted event
        emit!(OpExecuted {
            nonce: 1,
            to: Pubkey::default(),
            data: vec![0u8; 32],
        });

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