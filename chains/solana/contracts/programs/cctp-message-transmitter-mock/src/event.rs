use anchor_lang::prelude::*;

#[event]
pub struct OwnershipTransferStarted {
    pub previous_owner: Pubkey,
    pub new_owner: Pubkey,
}

#[event]
pub struct OwnershipTransferred {
    pub previous_owner: Pubkey,
    pub new_owner: Pubkey,
}

#[event]
pub struct PauserChanged {
    pub new_address: Pubkey,
}

#[event]
pub struct AttesterManagerUpdated {
    pub previous_attester_manager: Pubkey,
    pub new_attester_manager: Pubkey,
}

#[event]
pub struct MessageReceived {
    pub caller: Pubkey,
    pub source_domain: u32,
    pub nonce: u64,
    pub sender: Pubkey,
    pub message_body: Vec<u8>,
}

#[event]
pub struct SignatureThresholdUpdated {
    pub old_signature_threshold: u32,
    pub new_signature_threshold: u32,
}

#[event]
pub struct AttesterEnabled {
    pub attester: Pubkey,
}

#[event]
pub struct AttesterDisabled {
    pub attester: Pubkey,
}

#[event]
pub struct MaxMessageBodySizeUpdated {
    pub new_max_message_body_size: u64,
}

#[event]
pub struct Pause {}

#[event]
pub struct Unpause {}
