//! Contains common logic used across ccip programs. Common logic is versioned independently:
//! v1 of the common logic may not necessarily be used by v1 of each particular program.
use anchor_lang::prelude::*;

pub mod seed;
pub mod v1;

#[error_code]
pub enum CommonCcipError {
    #[msg("The given sequence interval is invalid")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 3000 = start at 9000)
    InvalidSequenceInterval = 4000,
    #[msg("Invalid pool accounts")]
    InvalidInputsPoolAccounts,
    #[msg("Invalid token accounts")]
    InvalidInputsTokenAccounts,
    #[msg("Invalid config account")]
    InvalidInputsConfigAccounts,
    #[msg("Invalid Token Admin Registry account")]
    InvalidInputsTokenAdminRegistryAccounts,
    #[msg("Invalid LookupTable account")]
    InvalidInputsLookupTableAccounts,
    #[msg("Invalid LookupTable account writable access")]
    InvalidInputsLookupTableAccountWritable,
}

// Used only to enable the `account` attribute to enable account deserialization, for accounts
// owned by the router (see TokenAdminRegistry below)
declare_id!("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C");
pub mod accounts_for_deserialization {
    /// This mod holds structs that represent accounts owned by other programs,
    /// but are used here just to deserialize data from them. Must be kept in sync
    /// with their originals.
    use super::*;

    // From ccip_router
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
