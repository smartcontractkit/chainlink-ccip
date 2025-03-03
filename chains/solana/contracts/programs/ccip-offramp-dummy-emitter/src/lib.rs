use anchor_lang::prelude::*;

pub mod events;
pub mod state;

use events::*;
use events::admin::*;
use state::*;

declare_id!("7h44xjUiJHH5wJCNpewaEDmgYLbUd7DDp6URuBKEenMT");

#[program]
pub mod ccip_offramp_dummy_emitter {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // 1. Emit CommitReportAccepted
        emit!(CommitReportAccepted {
            merkle_root: [0u8; 32],
            price_updates: PriceUpdates {
                token_prices_usd: vec![],
                gas_prices_usd: vec![],
                dest_token_prices: vec![],
                dest_native_prices: vec![],
            },
        });

        // 2. Emit SkippedAlreadyExecutedMessage
        emit!(SkippedAlreadyExecutedMessage {
            source_chain_selector: 1,
            sequence_number: 1,
        });

        // 3. Emit ExecutionStateChanged
        emit!(ExecutionStateChanged {
            source_chain_selector: 1,
            sequence_number: 1,
            message_id: [0u8; 32],
            message_hash: [0u8; 32],
            state: MessageExecutionState::Success,
        });

        // 4. Emit Admin Events
        let source_chain_config = SourceChainConfig {
            is_enabled: true,
            lane_code_version: CodeVersion::V1,
            on_ramp: [[0u8; 64]; 2],
        };

        emit!(SourceChainConfigUpdated {
            source_chain_selector: 1,
            source_chain_config: source_chain_config.clone(),
        });

        emit!(SourceChainAdded {
            source_chain_selector: 2,
            source_chain_config,
        });

        let owner = Pubkey::default();
        let new_owner = Pubkey::default();
        emit!(OwnershipTransferRequested {
            from: owner,
            to: new_owner,
        });

        emit!(OwnershipTransferred {
            from: owner,
            to: new_owner,
        });

        emit!(ConfigSet {
            svm_chain_selector: 1,
            enable_manual_execution_after: 3600, // 1 hour
        });

        emit!(ReferenceAddressesSet {
            router: Pubkey::default(),
            fee_quoter: Pubkey::default(),
            offramp_lookup_table: Pubkey::default(),
        });

        msg!("All events have been emitted!");
        Ok(())
    }

    /// Emit events and reverts. Will mark transaction as failed.
    pub fn trigger_all_events_reverts(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // Same events as trigger_all_events
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