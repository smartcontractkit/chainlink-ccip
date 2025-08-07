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
pub const CHAIN_FAMILY_SELECTOR_TVM: u32 = 0x647e2ba9;

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
    #[derive(InitSpace, PartialEq, Debug)]
    pub struct TokenAdminRegistry {
        pub version: u8,
        pub administrator: Pubkey,
        pub pending_administrator: Pubkey,
        pub lookup_table: Pubkey,
        // binary representation of indexes that are writable in token pool lookup table
        // lookup table can store 256 addresses
        pub writable_indexes: [u128; 2],
        pub mint: Pubkey,
        // if true, the pool supports calling `derive_accounts_lock_or_burn` and
        // `derive_accounts_release_or_mint` to derive the list of accounts and LUTs
        // needed to interact with it.
        pub supports_auto_derivation: bool,
    }

    #[derive(AnchorDeserialize)]
    pub(super) struct TokenAdminRegistryV1 {
        pub version: u8,
        pub administrator: Pubkey,
        pub pending_administrator: Pubkey,
        pub lookup_table: Pubkey,
        pub writable_indexes: [u128; 2],
        pub mint: Pubkey,
    }

    impl TryFrom<TokenAdminRegistryV1> for TokenAdminRegistry {
        type Error = anchor_lang::error::Error;

        fn try_from(
            v1: TokenAdminRegistryV1,
        ) -> std::result::Result<TokenAdminRegistry, Self::Error> {
            require_eq!(
                v1.version,
                1, // this deserialization is only valid for v1
                CommonCcipError::InvalidInputsTokenAdminRegistryAccounts
            );

            Ok(TokenAdminRegistry {
                version: 1,
                administrator: v1.administrator,
                pending_administrator: v1.pending_administrator,
                lookup_table: v1.lookup_table,
                writable_indexes: v1.writable_indexes,
                mint: v1.mint,
                supports_auto_derivation: false, // this is not part of the v1 data, it defaults to false
            })
        }
    }
}
