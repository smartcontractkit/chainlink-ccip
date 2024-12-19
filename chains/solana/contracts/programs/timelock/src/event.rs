use anchor_lang::prelude::*;

#[event]
/// @dev Emitted when a call is scheduled as part of operation `id`.
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
/// @dev Emitted when a call is performed as part of operation `id`.
pub struct CallExecuted {
    pub id: [u8; 32],
    pub index: u64,
    pub target: Pubkey,
    pub data: Vec<u8>,
}

#[event]
/// @dev Emitted when a call is performed via bypasser.
pub struct BypasserCallExecuted {
    pub index: u64,
    pub target: Pubkey,
    pub data: Vec<u8>,
}

#[event]
/// @dev Emitted when operation `id` is cancelled.
pub struct Cancelled {
    pub id: [u8; 32],
}

#[event]
/// @dev Emitted when the minimum delay for future operations is modified.
pub struct MinDelayChange {
    pub old_duration: u64,
    pub new_duration: u64,
}

#[event]
/// @dev Emitted when a function selector is blocked.
pub struct FunctionSelectorBlocked {
    pub selector: [u8; 8],
}

#[event]
/// @dev Emitted when a function selector is unblocked.
pub struct FunctionSelectorUnblocked {
    pub selector: [u8; 8],
}
