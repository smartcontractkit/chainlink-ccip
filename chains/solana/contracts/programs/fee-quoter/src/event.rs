use anchor_lang::prelude::*;

use crate::state::{DestChainConfig, TokenTransferFeeConfig};

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
pub struct DestChainAdded {
    pub dest_chain_selector: u64,
    pub dest_chain_config: DestChainConfig,
}

#[event]
pub struct DestChainConfigUpdated {
    pub dest_chain_selector: u64,
    pub dest_chain_config: DestChainConfig,
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
    pub token_transfer_fee_config: TokenTransferFeeConfig,
}
