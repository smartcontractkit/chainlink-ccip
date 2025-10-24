//! # EIP-712 Typed Structured Data Hashing
//!
//! This module implements EIP-712 typed structured data hashing for Ethereum-compatible
//! signature verification.
//!
//! ## Domain Separator
//!
//! The domain separator uniquely identifies this contract instance and prevents
//! signature replay attacks across different chains or contracts:
//!
//! - **name**: "ManyChainMultiSig" - The contract name
//! - **version**: "1" - The contract version
//! - **chainId**: Solana chain identifier (e.g., hash of "solana:localnet")
//! - **salt**: The Solana Program ID (32 bytes) - Uniquely identifies this program instance
//!
//! ## Message Structure
//!
//! Users sign a `RootValidation` message containing:
//! - **root**: The 32-byte Merkle root of operations
//! - **validUntil**: Timestamp (uint32) until which the root is valid

use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::{hashv, Hash, HASH_BYTES};

/// Type hash for EIP712Domain
/// keccak256("EIP712Domain(string name,string version,uint256 chainId,bytes32 salt)")
pub const EIP712_DOMAIN_TYPE_HASH: &[u8; HASH_BYTES] = &[
    0xa6, 0x04, 0xff, 0xf5, 0xa2, 0x7d, 0x59, 0x51, 0xf3, 0x34, 0xcc, 0xda, 0x7a, 0xbf, 0xf3, 0x28,
    0x6a, 0x8a, 0xf2, 0x9c, 0xae, 0xeb, 0x19, 0x6a, 0x6f, 0x2b, 0x40, 0xa1, 0xdc, 0xe7, 0x61, 0x2b,
];

/// Type hash for RootValidation message
/// keccak256("RootValidation(bytes32 root,uint32 validUntil)")
pub const ROOT_VALIDATION_TYPE_HASH: &[u8; HASH_BYTES] = &[
    0x07, 0x02, 0x19, 0xd6, 0x56, 0xf4, 0x72, 0x71, 0xfe, 0x6e, 0x2e, 0x8c, 0xff, 0x49, 0x55, 0xfe,
    0xcb, 0xca, 0xb4, 0xc9, 0x42, 0x22, 0x83, 0x09, 0xe9, 0x0a, 0xae, 0x20, 0xbd, 0x1f, 0xa8, 0x95,
];

/// Hash of the domain name "ManyChainMultiSig"
/// keccak256("ManyChainMultiSig")
pub const EIP712_DOMAIN_NAME_HASH: &[u8; HASH_BYTES] = &[
    0x62, 0x7b, 0x9a, 0x5c, 0xae, 0x29, 0x68, 0x84, 0x2a, 0x54, 0x1a, 0x6b, 0xa8, 0x61, 0xa7, 0x15,
    0x0a, 0xb5, 0xec, 0x3c, 0x42, 0x02, 0xa6, 0x24, 0xbc, 0xa4, 0x6b, 0xff, 0x04, 0x11, 0xce, 0x9a,
];

/// Hash of the domain version "1"
/// keccak256("1")
pub const EIP712_DOMAIN_VERSION_HASH: &[u8; HASH_BYTES] = &[
    0xc8, 0x9e, 0xfd, 0xaa, 0x54, 0xc0, 0xf2, 0x0c, 0x7a, 0xdf, 0x61, 0x28, 0x82, 0xdf, 0x09, 0x50,
    0xf5, 0xa9, 0x51, 0x63, 0x7e, 0x03, 0x07, 0xcd, 0xcb, 0x4c, 0x67, 0x2f, 0x29, 0x8b, 0x8b, 0xc6,
];

/// Computes the EIP-712 message hash for root validation.
///
/// This is the final hash that gets signed by the user. It follows the EIP-712
/// format: `keccak256("\x19\x01" || domainSeparator || structHash)`
///
/// # Parameters
///
/// - `root`: The 32-byte Merkle root
/// - `valid_until`: Timestamp until which the root is valid
/// - `chain_id`: The chain identifier
/// - `program_id`: The Solana Program ID
///
/// # Returns
///
/// - The 32-byte message hash ready for ECDSA verification
pub fn compute_message_hash(
    root: &[u8; HASH_BYTES],
    valid_until: u32,
    chain_id: u64,
    program_id: &Pubkey,
) -> Hash {
    let domain_separator = compute_domain_hash(chain_id, program_id);
    let struct_hash = compute_struct_hash(root, valid_until);

    hashv(&[
        b"\x19\x01",
        &domain_separator.to_bytes(),
        &struct_hash.to_bytes(),
    ])
}

/// Computes the EIP-712 domain separator hash.
///
/// The domain separator ensures that signatures are unique to this specific
/// contract instance and chain, preventing replay attacks.
///
/// # Parameters
///
/// - `chain_id`: The chain identifier (e.g., keccak256("solana:localnet"))
/// - `program_id`: The Solana Program ID (32 bytes), used as the salt
///
/// # Returns
///
/// - The 32-byte domain separator hash
pub fn compute_domain_hash(chain_id: u64, program_id: &Pubkey) -> Hash {
    // Chain ID as bytes32 (big-endian, left-padded)
    let chain_id_bytes = left_pad_to_32(&chain_id.to_be_bytes());

    // Salt = Program ID (32 bytes)
    let salt = program_id.to_bytes();

    hashv(&[
        EIP712_DOMAIN_TYPE_HASH,
        EIP712_DOMAIN_NAME_HASH,
        EIP712_DOMAIN_VERSION_HASH,
        &chain_id_bytes,
        &salt,
    ])
}

/// Computes the EIP-712 struct hash for a RootValidation message.
///
/// # Parameters
///
/// - `root`: The 32-byte Merkle root
/// - `valid_until`: Timestamp until which the root is valid
///
/// # Returns
///
/// - The 32-byte struct hash
pub fn compute_struct_hash(root: &[u8; HASH_BYTES], valid_until: u32) -> Hash {
    // valid_until as bytes32 (big-endian, left-padded)
    let valid_until_bytes = left_pad_to_32(&valid_until.to_be_bytes());

    hashv(&[ROOT_VALIDATION_TYPE_HASH, root, &valid_until_bytes])
}

/// Left-pads a byte array to 32 bytes with zeros.
///
/// # Parameters
///
/// - `input`: The input byte array
///
/// # Returns
///
/// - A 32-byte array with the input left-padded with zeros
fn left_pad_to_32(input: &[u8]) -> [u8; 32] {
    let mut padded = [0u8; 32];
    let start = 32 - input.len();
    padded[start..].copy_from_slice(input);
    padded
}

#[cfg(test)]
mod tests {
    use super::*;
    use anchor_lang::solana_program::keccak::hash;

    /// EIP-712 domain name for the ManyChainMultiSig contract
    pub const EIP712_DOMAIN_NAME: &str = "ManyChainMultiSig";

    /// EIP-712 domain version
    pub const EIP712_DOMAIN_VERSION: &str = "1";

    // Last 8 bytes of keccak256("solana:localnet") as big-endian
    const CHAIN_ID: u64 = 5190648258797659666;

    fn decode32(s: &str) -> [u8; HASH_BYTES] {
        hex::decode(s).unwrap().try_into().unwrap()
    }

    #[test]
    fn verify_domain_type_hash() {
        let expected =
            hash(b"EIP712Domain(string name,string version,uint256 chainId,bytes32 salt)");
        assert_eq!(&expected.to_bytes(), EIP712_DOMAIN_TYPE_HASH);
    }

    #[test]
    fn verify_domain_name_hash() {
        let expected = hash(EIP712_DOMAIN_NAME.as_bytes());
        assert_eq!(&expected.to_bytes(), EIP712_DOMAIN_NAME_HASH);
    }

    #[test]
    fn verify_domain_version_hash() {
        let expected = hash(EIP712_DOMAIN_VERSION.as_bytes());
        assert_eq!(&expected.to_bytes(), EIP712_DOMAIN_VERSION_HASH);
    }

    #[test]
    fn verify_root_validation_type_hash() {
        let expected = hash(b"RootValidation(bytes32 root,uint32 validUntil)");
        assert_eq!(&expected.to_bytes(), ROOT_VALIDATION_TYPE_HASH);
    }

    #[test]
    fn test_left_pad_to_32() {
        let input = [1, 2, 3, 4];
        let result = left_pad_to_32(&input);
        let expected: [u8; 32] = [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
            2, 3, 4,
        ];
        assert_eq!(result, expected);
    }

    #[test]
    fn test_compute_domain_hash() {
        let program_id = Pubkey::new_from_array(decode32(
            "b870e12dd379891561d2e9fa8f26431834eb736f2f24fc2a2a4dff1fd5dca4df",
        ));

        let domain_hash = compute_domain_hash(CHAIN_ID, &program_id);

        let chain_id_bytes = left_pad_to_32(&CHAIN_ID.to_be_bytes());
        let salt = program_id.to_bytes();

        let expected = hashv(&[
            EIP712_DOMAIN_TYPE_HASH,
            EIP712_DOMAIN_NAME_HASH,
            EIP712_DOMAIN_VERSION_HASH,
            &chain_id_bytes,
            &salt,
        ]);

        assert_eq!(domain_hash.to_bytes(), expected.to_bytes());
    }

    #[test]
    fn test_compute_struct_hash() {
        let root = decode32("d5ef592d1ad183db43b4980d7ab7ee43a6f6a284988c3e3a23d38c07beb520c7");
        let valid_until: u32 = 1748317727;

        let struct_hash = compute_struct_hash(&root, valid_until);

        // Verify the structure
        let valid_until_bytes = left_pad_to_32(&valid_until.to_be_bytes());
        let expected = hashv(&[ROOT_VALIDATION_TYPE_HASH, &root, &valid_until_bytes]);

        assert_eq!(struct_hash.to_bytes(), expected.to_bytes());
    }

    #[test]
    fn test_compute_message_hash() {
        let root = decode32("d5ef592d1ad183db43b4980d7ab7ee43a6f6a284988c3e3a23d38c07beb520c7");
        let valid_until: u32 = 1748317727;
        let program_id = Pubkey::new_from_array(decode32(
            "b870e12dd379891561d2e9fa8f26431834eb736f2f24fc2a2a4dff1fd5dca4df",
        ));

        let message_hash = compute_message_hash(&root, valid_until, CHAIN_ID, &program_id);

        // Verify it follows EIP-712 format: \x19\x01 || domainSeparator || structHash
        let domain_separator = compute_domain_hash(CHAIN_ID, &program_id);
        let struct_hash = compute_struct_hash(&root, valid_until);

        let expected = hashv(&[
            b"\x19\x01",
            &domain_separator.to_bytes(),
            &struct_hash.to_bytes(),
        ]);

        assert_eq!(message_hash.to_bytes(), expected.to_bytes());
    }
}
