use anchor_lang::prelude::*;

use crate::state::DestChainConfig;

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
