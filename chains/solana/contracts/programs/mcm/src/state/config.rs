use anchor_lang::prelude::*;

use crate::{constant::*, EVM_ADDRESS_BYTES};

#[account]
pub struct ConfigSigners {
    pub signer_addresses: Vec<[u8; 20]>,
    pub total_signers: u8,
    pub is_finalized: bool,
    pub bump: u8,
}

impl ConfigSigners {
    // 8 (discriminator) + 4 (vec len) + (20 * total_signers) +1(total_signers) + 1 (is_finalized) + 1 (bump)
    pub const fn space(total_signers: usize) -> usize {
        8 + 4 + (20 * total_signers) + 1 + 1 + 1
    }
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, InitSpace, Debug)]
pub struct McmSigner {
    pub evm_address: [u8; EVM_ADDRESS_BYTES],
    pub index: u8,
    pub group: u8,
}

#[account]
pub struct MultisigConfig {
    pub chain_id: u64,
    pub multisig_name: [u8; MULTISIG_NAME_PADDED],

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub group_quorums: [u8; NUM_GROUPS],
    pub group_parents: [u8; NUM_GROUPS],

    // Keep variable-length data at the end of the account struct
    // https://solana.com/developers/courses/program-optimization/program-architecture#data-order
    pub signers: Vec<McmSigner>, // unable to store as hashmap in Solana
}

impl MultisigConfig {
    pub const INIT_SPACE: usize = 8 + // chain_id (u64)
        NUM_GROUPS + // group_quorums [u8; NUM_GROUPS]
        NUM_GROUPS + // group_parents [u8; NUM_GROUPS]
        32 + // owner (Pubkey)
        32 + // proposed_owner (Pubkey)
        4 + // string prefix for multisig_name
        MULTISIG_NAME_PADDED + // fixed max multisig_name length from initialization
        4; // empty vec prefix for signers

    // for realloc - only need to account for signers
    pub fn space_with_signers(num_signers: usize) -> usize {
        Self::INIT_SPACE +      // Base space including fixed multisig_name
        num_signers * McmSigner::INIT_SPACE // Just add signers space
    }
}
