use anchor_lang::prelude::*;

use crate::eth_utils::*;

#[account]
pub struct RootSignatures {
    pub total_signatures: u8,
    pub signatures: Vec<Signature>,
    pub is_finalized: bool,
    pub bump: u8,
}

impl RootSignatures {
    // 8 (discriminator) + 4 (vec len) + (65 * max_sigs) + 32 (root) + 4 (valid_until) + 1 (is_finalized) + 1 (bump)
    pub const fn space(total_signatures: usize) -> usize {
        8 + 4 + (65 * total_signatures) + 32 + 4 + 1 + 1 + 1
    }
}

#[account]
#[derive(InitSpace, Debug)]
pub struct RootMetadata {
    pub chain_id: u64,
    pub multisig: Pubkey,
    pub pre_op_count: u64,
    pub post_op_count: u64,
    pub override_previous_root: bool,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct RootMetadataInput {
    pub chain_id: u64,
    pub multisig: Pubkey,
    pub pre_op_count: u64,
    pub post_op_count: u64,
    pub override_previous_root: bool,
}

#[account]
#[derive(InitSpace)]
pub struct ExpiringRootAndOpCount {
    pub root: [u8; 32],
    pub valid_until: u32,
    pub op_count: u64,
}

#[account]
#[derive(InitSpace)]
pub struct SeenSignedHash {
    pub seen: bool, // always set to true in practice when initialized, hacky way of defining a Set via PDAs.
}
