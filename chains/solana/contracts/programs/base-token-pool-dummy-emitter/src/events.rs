use anchor_lang::prelude::*;
use crate::state::*;

#[event]
pub struct TokensConsumed {
    pub tokens: u64,
}

#[event]
pub struct ConfigChanged {
    pub config: RateLimitConfig,
}

#[event]
pub struct Burned {
    pub sender: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Minted {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Locked {
    pub sender: Pubkey,
    pub amount: u64,
}

#[event]
pub struct Released {
    pub sender: Pubkey,
    pub recipient: Pubkey,
    pub amount: u64,
}

#[event]
pub struct RemoteChainConfigured {
    pub chain_selector: u64,
    pub token: RemoteAddress,
    pub previous_token: RemoteAddress,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RateLimitConfigured {
    pub chain_selector: u64,
    pub outbound_rate_limit: RateLimitConfig,
    pub inbound_rate_limit: RateLimitConfig,
}

#[event]
pub struct RemotePoolsAppended {
    pub chain_selector: u64,
    pub pool_addresses: Vec<RemoteAddress>,
    pub previous_pool_addresses: Vec<RemoteAddress>,
}

#[event]
pub struct RemoteChainRemoved {
    pub chain_selector: u64,
}

#[event]
pub struct RouterUpdated {
    pub old_router: Pubkey,
    pub new_router: Pubkey,
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