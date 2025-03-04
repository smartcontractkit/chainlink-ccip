use anchor_lang::prelude::*;
use crate::state::*;

#[event]
pub struct NewRoot {
    pub root: [u8; 32],
    pub valid_until: u32,
    pub metadata_chain_id: u64,
    pub metadata_multisig: Pubkey,
    pub metadata_pre_op_count: u64,
    pub metadata_post_op_count: u64,
    pub metadata_override_previous_root: bool,
}

#[event]
pub struct ConfigSet {
    pub group_parents: [u8; NUM_GROUPS],
    pub group_quorums: [u8; NUM_GROUPS],
    pub is_root_cleared: bool,
    pub signers: Vec<McmSigner>,
}

#[event]
pub struct OpExecuted {
    pub nonce: u64,
    pub to: Pubkey,
    pub data: Vec<u8>,
}