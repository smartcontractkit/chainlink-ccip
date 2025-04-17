//! Contains common logic used across ccip programs. Common logic is versioned independently:
//! v1 of the common logic may not necessarily be used by v1 of each particular program.
use anchor_lang::prelude::*;

pub mod context;
pub mod seed;
pub mod v1;

#[error_code]
pub enum CommonCcipError {
    #[msg("The given sequence interval is invalid")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 4000 = start at 10000)
    InvalidSequenceInterval = 4000,
    #[msg("Invalid pool accounts")]
    InvalidInputsPoolAccounts,
    #[msg("Invalid token accounts")]
    InvalidInputsTokenAccounts,
    #[msg("Invalid Token Admin Registry account")]
    InvalidInputsTokenAdminRegistryAccounts,
    #[msg("Invalid LookupTable account")]
    InvalidInputsLookupTableAccounts,
    #[msg("Invalid LookupTable account writable access")]
    InvalidInputsLookupTableAccountWritable,
    #[msg("Invalid pool signer account")]
    InvalidInputsPoolSignerAccounts,

    #[msg("Invalid chain family selector")]
    InvalidChainFamilySelector,
    #[msg("Invalid encoding")]
    InvalidEncoding,
    #[msg("Invalid EVM address")]
    InvalidEVMAddress,
    #[msg("Invalid SVM address")]
    InvalidSVMAddress,
}

// https://github.com/smartcontractkit/chainlink/blob/ff8a597fd9df653f8967427498eaa5a04b19febb/contracts/src/v0.8/ccip/libraries/Internal.sol#L276
pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;
pub const CHAIN_FAMILY_SELECTOR_SVM: u32 = 0x1e10bdc4;

// Duplicates the router ID to declare router accounts that must be visible from the common crate,
// avoiding a circular dependency. This means this crate may only declare accounts that belong
// to the router, and no other program.
declare_id!("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C");
// null contract required for IDL + gobinding generation
#[program]
pub mod ccip_common {}

pub mod router_accounts {
    use super::*;

    #[account]
    #[derive(InitSpace)]
    pub struct TokenAdminRegistry {
        pub version: u8,
        pub administrator: Pubkey,
        pub pending_administrator: Pubkey,
        pub lookup_table: Pubkey,
        // binary representation of indexes that are writable in token pool lookup table
        // lookup table can store 256 addresses
        pub writable_indexes: [u128; 2],
        pub mint: Pubkey,
    }
}
