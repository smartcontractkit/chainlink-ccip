use anchor_lang::prelude::*;

use crate::context::{MerkleRoot, PriceUpdates};
use crate::state::MessageExecutionState;

#[event]
pub struct CommitReportAccepted {
    // Will be zero'd (MerkleRoot::Default() in case the commit report included no
    // merkle root. This is to circumvent a Go deserialization issue.)
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

pub mod admin {
    use crate::state::SourceChainConfig;

    use super::*;

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

    #[event]
    pub struct ConfigSet {
        pub svm_chain_selector: u64,
        pub enable_manual_execution_after: i64,
    }

    #[event]
    pub struct ReferenceAddressesSet {
        pub router: Pubkey,
        pub fee_quoter: Pubkey,
        pub offramp_lookup_table: Pubkey,
    }
}
