use anchor_lang::prelude::*;

#[event]
pub struct CallScheduled {
    pub id: [u8; 32],
    pub index: u64,
    pub target: Pubkey,
    pub predecessor: [u8; 32],
    pub salt: [u8; 32],
    pub delay: u64,
    pub data: Vec<u8>,
}

#[event]
pub struct CallExecuted {
    pub id: [u8; 32],
    pub index: u64,
    pub target: Pubkey,
    pub data: Vec<u8>,
}

#[event]
pub struct BypasserCallExecuted {
    pub index: u64,
    pub target: Pubkey,
    pub data: Vec<u8>,
}

#[event]
pub struct Cancelled {
    pub id: [u8; 32],
}

#[event]
pub struct MinDelayChange {
    pub old_duration: u64,
    pub new_duration: u64,
}

#[event]
pub struct FunctionSelectorBlocked {
    pub selector: [u8; 8],
}

#[event]
pub struct FunctionSelectorUnblocked {
    pub selector: [u8; 8],
}