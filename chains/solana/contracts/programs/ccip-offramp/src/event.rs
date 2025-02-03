use anchor_lang::prelude::*;

use crate::context::{MerkleRoot, PriceUpdates};
use crate::state::MessageExecutionState;

#[event]
pub struct CommitReportAccepted {
    pub merkle_root: MerkleRoot,
    pub price_updates: PriceUpdates,
}

#[event]
pub struct SkippedAlreadyExecutedMessage {
    pub source_chain_selector: u64,
    pub sequence_number: u64,
}

#[event]
pub struct ExecutionStateChanged {
    pub source_chain_selector: u64,
    pub sequence_number: u64,
    pub message_id: [u8; 32],
    pub message_hash: [u8; 32],
    pub state: MessageExecutionState,
}

#[event]
pub struct UsdPerUnitGasUpdated {
    pub dest_chain: u64,
    pub value: [u8; 28], // EVM uses u256 here
    pub timestamp: i64,  // EVM uses u256 here
}

#[event]
pub struct UsdPerTokenUpdated {
    pub token: Pubkey,
    pub value: [u8; 28], // EVM uses u256 here
    pub timestamp: i64,  // EVM uses u256 here
}

pub mod admin {
    use crate::state::SourceChainConfig;

    use super::*;

    #[event]
    pub struct FeeTokenAdded {
        pub fee_token: Pubkey,
        pub enabled: bool,
    }

    #[event]
    pub struct FeeTokenEnabled {
        pub fee_token: Pubkey,
    }

    #[event]
    pub struct FeeTokenDisabled {
        pub fee_token: Pubkey,
    }

    #[event]
    pub struct FeeTokenRemoved {
        pub fee_token: Pubkey,
    }

    #[event]
    pub struct SourceChainConfigUpdated {
        pub source_chain_selector: u64,
        pub source_chain_config: SourceChainConfig,
    }

    #[event]
    pub struct SourceChainAdded {
        pub source_chain_selector: u64,
        pub source_chain_config: SourceChainConfig,
    }

    #[event]
    pub struct OwnershipTransferRequested {
        pub from: Pubkey,
        pub to: Pubkey,
    }

    #[event]
    pub struct OwnershipTransferred {
        pub from: Pubkey,
        pub to: Pubkey,
    }
}
