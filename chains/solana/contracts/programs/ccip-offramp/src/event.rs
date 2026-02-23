use anchor_lang::prelude::*;

use crate::context::{MerkleRoot, PriceUpdates};
use crate::state::MessageExecutionState;

#[event]
pub struct CommitReportAccepted {
    pub merkle_root: Option<MerkleRoot>,
    pub price_updates: PriceUpdates,
}

#[event]
pub struct CommitReportPDAClosed {
    pub source_chain_selector: u64,
    pub merkle_root: [u8; 32],
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
        pub rmn_remote: Pubkey,
    }
}

/// Event used only during IDL build to ensure CommitInput and related types
/// are exported in the IDL for Go bindings generation.
/// This event is never emitted at runtime.
#[cfg(feature = "idl-build")]
#[event]
pub struct IdlBuildTypeExport {
    pub commit_input: crate::context::CommitInput,
    pub token_price_update: crate::context::TokenPriceUpdate,
    pub gas_price_update: crate::context::GasPriceUpdate,
    pub any2svm_ramp_message: crate::messages::Any2SVMRampMessage,
    pub ramp_message_header: crate::messages::RampMessageHeader,
    pub any2svm_ramp_extra_args: crate::messages::Any2SVMRampExtraArgs,
    pub any2svm_token_transfer: crate::messages::Any2SVMTokenTransfer,
    pub execution_report_single_chain: crate::messages::ExecutionReportSingleChain,
    pub cross_chain_amount: crate::messages::CrossChainAmount,
    pub svm_token_amount: crate::messages::SVMTokenAmount,
}
