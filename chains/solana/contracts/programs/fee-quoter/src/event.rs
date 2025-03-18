use anchor_lang::prelude::*;

use crate::state::{CodeVersion, DestChainConfig, TokenTransferFeeConfig};

#[event]
pub struct ConfigSet {
    pub max_fee_juels_per_msg: u128,
    pub link_token_mint: Pubkey,
    pub link_token_local_decimals: u8,
    pub onramp: Pubkey,
    pub default_code_version: CodeVersion,
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

/// Emitted when the price update corresponds to a token that isn't registered
/// for price tracking.
#[event]
pub struct TokenPriceUpdateIgnored {
    pub token: Pubkey,
    pub value: [u8; 28], // EVM uses u256 here
}

#[event]
pub struct TokenTransferFeeConfigUpdated {
    pub dest_chain_selector: u64,
    pub token: Pubkey,
    pub token_transfer_fee_config: TokenTransferFeeConfig,
}

#[event]
pub struct PremiumMultiplierWeiPerEthUpdated {
    pub token: Pubkey,
    pub premium_multiplier_wei_per_eth: u64,
}

#[event]
pub struct PriceUpdaterAdded {
    pub price_updater: Pubkey,
}

#[event]
pub struct PriceUpdaterRemoved {
    pub price_updater: Pubkey,
}
