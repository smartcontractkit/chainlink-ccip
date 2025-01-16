use anchor_lang::prelude::*;

use crate::{
    context::MerkleRoot, DestChainConfig, MessageExecutionState, PriceUpdates, SVM2AnyRampMessage,
    SourceChainConfig, TokenBilling,
};

#[event]
pub struct CCIPMessageSent {
    pub dest_chain_selector: u64,
    pub sequence_number: u64,
    pub message: SVM2AnyRampMessage,
}

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
pub struct AlreadyAttempted {
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
pub struct PoolSet {
    pub token: Pubkey,
    pub previous_pool_lookup_table: Pubkey,
    pub new_pool_lookup_table: Pubkey,
}

#[event]
pub struct AdministratorTransferRequested {
    pub token: Pubkey,
    pub current_admin: Pubkey,
    pub new_admin: Pubkey,
}

#[event]
pub struct AdministratorTransferred {
    pub token: Pubkey,
    pub new_admin: Pubkey,
}

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

#[event]
pub struct TokenTransferFeeConfigUpdated {
    pub dest_chain_selector: u64,
    pub token: Pubkey,
    pub token_transfer_fee_config: TokenBilling,
}

#[event]
pub struct PremiumMultiplierWeiPerEthUpdated {
    pub token: Pubkey,
    pub premium_multiplier_wei_per_eth: u64,
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
pub struct DestChainConfigUpdated {
    pub dest_chain_selector: u64,
    pub dest_chain_config: DestChainConfig,
}

#[event]
pub struct DestChainAdded {
    pub dest_chain_selector: u64,
    pub dest_chain_config: DestChainConfig,
}

#[event]
pub struct AdministratorRegistered {
    pub token_mint: Pubkey,
    pub administrator: Pubkey,
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
