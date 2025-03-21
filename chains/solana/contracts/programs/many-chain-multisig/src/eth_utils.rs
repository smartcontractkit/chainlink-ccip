//! # Ethereum Cryptographic Utilities
//!
//! This module provides cryptographic primitives and utilities required for Ethereum-compatible
//! signature verification and Merkle tree operations. It enables the MCM program to verify
//! ECDSA signatures and validate Merkle proofs across multiple chains.
//!
//! ## Domain Separators
//!
//! The module uses domain separators to ensure cryptographic isolation between different
//! parts of the system:
//!
//! - `MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA_SOLANA`: Used when hashing root metadata
//! - `MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP_SOLANA`: Used when hashing operation data
//!
//! These separators ensure that hashes for different purposes cannot be reused or confused.
use anchor_lang::prelude::*;
use anchor_lang::solana_program::keccak::{hash, hashv, Hash, HASH_BYTES}; // use keccak256 for EVM compatibility
use anchor_lang::solana_program::secp256k1_recover::{
    secp256k1_recover, Secp256k1Pubkey, Secp256k1RecoverError,
};

use crate::error::*;
use crate::state::root::RootMetadataInput;

// Domain separators & evm constants
// NOTE: chain-specific mcm contract should has its own domain separator to avoid ambiguity
// https://github.com/smartcontractkit/ccip-owner-contracts#porting
//

/// Domain separator for metadata leaves in the Merkle tree.
///
/// Result of keccak256("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA_SOLANA").
/// This value is unique to the Solana implementation to prevent cross-chain replay attacks.
pub const MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA: &[u8; HASH_BYTES] = &[
    0x47, 0xfd, 0xed, 0x70, 0x90, 0x1d, 0x27, 0x3, 0x83, 0x94, 0xdb, 0x90, 0x5a, 0x72, 0x56, 0x3c,
    0xad, 0x6f, 0x7, 0x58, 0x1d, 0xbc, 0xdd, 0x14, 0x72, 0xcc, 0xd2, 0xf7, 0x42, 0xaf, 0x63, 0x60,
];

/// Domain separator for operation leaves in the Merkle tree.
///
/// Result of keccak256("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP_SOLANA").
/// This value is unique to the Solana implementation to prevent cross-chain replay attacks.
pub const MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP: &[u8; HASH_BYTES] = &[
    0xfb, 0x98, 0x81, 0x6f, 0xf3, 0xc5, 0x13, 0x8a, 0x68, 0xab, 0xfd, 0x40, 0xb8, 0xd8, 0xfb, 0xc2,
    0x29, 0x72, 0xfe, 0xa1, 0xdd, 0x89, 0x75, 0x73, 0x31, 0x32, 0x7e, 0x6e, 0xa, 0x94, 0x40, 0xb7,
];
/// Size of an Ethereum address in bytes (20 bytes / 160 bits)
pub const EVM_ADDRESS_BYTES: usize = 20;

/// Recovers an Ethereum address from an ECDSA signature on a given message hash.
///
/// This function is used to verify signatures during root setting and retrieves the
/// signer's Ethereum address (the last 20 bytes of the keccak256 hash of the recovered public key).
///
/// # Parameters
///
/// - `eth_signed_msg_hash`: 32-byte hash of the Ethereum signed message
/// - `sig`: The ECDSA signature containing v, r, s components
///
/// # Returns
///
/// - The recovered 20-byte Ethereum address if successful
/// - An error if recovery fails
pub fn ecdsa_recover_evm_addr(
    eth_signed_msg_hash: &[u8; HASH_BYTES],
    sig: &Signature,
) -> Result<[u8; EVM_ADDRESS_BYTES]> {
    // retrieve signer public key
    let public_key = sig
        .secp256k1_recover_from(eth_signed_msg_hash)
        .map_err(|_| McmError::FailedEcdsaRecover)?;

    let public_key_bytes = &public_key.to_bytes();

    // return last 20 bytes of hashed public key as the recovered ethereum address
    let evm_addr: [u8; EVM_ADDRESS_BYTES] = hash(public_key_bytes).to_bytes()
        [(HASH_BYTES - EVM_ADDRESS_BYTES)..]
        .try_into()
        .unwrap();

    Ok(evm_addr)
}

/// Computes the Ethereum-compatible message hash for root validation.
///
/// Creates a hash that matches Ethereum's personal sign message format:
/// "\x19Ethereum Signed Message:\n32" + keccak256(root || valid_until)
///
/// # Parameters
///
/// - `root`: The 32-byte Merkle root
/// - `valid_until`: Timestamp until which the root is valid
///
/// # Returns
///
/// - The 32-byte message hash ready for ECDSA verification
pub fn compute_eth_message_hash(root: &[u8; HASH_BYTES], valid_until: u32) -> Hash {
    // Use big-endian encoding for EVM compatibility
    let valid_until_bytes = left_pad_vec(&valid_until.to_be_bytes());
    let hashed_encoded_params = hashv(&[root, &valid_until_bytes]);

    hashv(&[
        b"\x19Ethereum Signed Message:\n32",
        &hashed_encoded_params.to_bytes(),
    ])
}

/// Calculates a Merkle root from a leaf node and a proof path.
///
/// This function iteratively combines a leaf hash with the provided proof elements
/// to reconstruct the Merkle root. It allows verification that a specific leaf
/// (operation or metadata) is included in the tree with the given root.
///
/// # Parameters
///
/// - `proof`: Ordered list of sibling hashes needed to reconstruct the path to the root
/// - `leaf`: The leaf hash being verified
///
/// # Returns
///
/// - The computed 32-byte Merkle root
pub fn calculate_merkle_root(
    proof: Vec<[u8; HASH_BYTES]>,
    leaf: &[u8; HASH_BYTES],
) -> [u8; HASH_BYTES] {
    let mut computed_hash = *leaf;
    for proof_element in proof {
        computed_hash = hash_pair(&computed_hash, &proof_element);
    }
    computed_hash
}

/// Hashes two nodes to create their parent node in the Merkle tree.
///
/// Ensures consistent ordering of hashing by sorting the nodes lexicographically first.
/// This matches the OpenZeppelin implementation used in the Ethereum contract.
///
/// # Parameters
///
/// - `a`: First 32-byte hash
/// - `b`: Second 32-byte hash
///
/// # Returns
///
/// - The 32-byte hash of the parent node
fn hash_pair(a: &[u8; HASH_BYTES], b: &[u8; HASH_BYTES]) -> [u8; HASH_BYTES] {
    let (left, right) = if a < b { (a, b) } else { (b, a) };
    hashv(&[left, right]).to_bytes()
}

/// Left-pads a byte array to the target size with zeros.
///
/// # Parameters
///
/// - `input`: The input byte array
/// - `num_bytes`: Target size for the padded result
///
/// # Returns
///
/// - A vector containing the left-padded result
fn _left_pad_vec(input: &[u8], num_bytes: usize) -> Vec<u8> {
    let mut padded: Vec<u8> = Vec::with_capacity(num_bytes);
    padded.resize(num_bytes - input.len(), 0);
    padded.extend_from_slice(input);
    padded
}

/// Left-pads a byte array to 32 bytes with zeros.
///
/// Used for consistent encoding of numeric values in the Merkle tree.
///
/// # Parameters
///
/// - `input`: The input byte array
///
/// # Returns
///
/// - A vector containing the left-padded result (32 bytes)
pub fn left_pad_vec(input: &[u8]) -> Vec<u8> {
    _left_pad_vec(input, HASH_BYTES)
}

/// ECDSA signature with components used in Ethereum signature verification.
#[derive(AnchorSerialize, AnchorDeserialize, Clone, InitSpace, Debug)]
pub struct Signature {
    pub v: u8,
    pub r: [u8; 32],
    pub s: [u8; 32],
}

impl Signature {
    /// Recovers the public key from this signature and a message hash.
    ///
    /// Used internally for ECDSA recovery of signer addresses.
    ///
    /// # Parameters
    ///
    /// - `eth_signed_msg_hash`: The 32-byte hash of the signed message
    ///
    /// # Returns
    ///
    /// - The recovered secp256k1 public key if successful
    /// - A Secp256k1RecoverError if recovery fails
    fn secp256k1_recover_from(
        &self,
        eth_signed_msg_hash: &[u8; HASH_BYTES],
    ) -> std::result::Result<Secp256k1Pubkey, Secp256k1RecoverError> {
        // Ref: https://github.com/anza-xyz/agave/blob/c8685ce0e1bb9b26014f1024de2cd2b8c308cbde/curves/secp256k1-recover/src/lib.rs#L106-L115
        let v = self
            .v
            .checked_sub(27)
            .ok_or(Secp256k1RecoverError::InvalidRecoveryId)?;
        let rs = [self.r, self.s].concat();
        secp256k1_recover(eth_signed_msg_hash, v, rs.as_slice())
    }
}

impl RootMetadataInput {
    /// Computes the Merkle leaf hash for this metadata.
    ///
    /// Combines the domain separator with the metadata fields to create a unique
    /// leaf hash that can be included in the Merkle tree.
    ///
    /// # Returns
    ///
    /// - The 32-byte hash representing this metadata as a Merkle leaf
    pub fn hash_leaf(&self) -> [u8; HASH_BYTES] {
        let chain_id = left_pad_vec(&self.chain_id.to_le_bytes());
        let pre_op_count = left_pad_vec(&self.pre_op_count.to_le_bytes());
        let post_op_count = left_pad_vec(&self.post_op_count.to_le_bytes());
        let override_previous_root: &[u8] = &[if self.override_previous_root { 1 } else { 0 }];
        let override_previous_root_bytes = left_pad_vec(override_previous_root);

        hashv(&[
            MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA,
            chain_id.as_slice(),
            &self.multisig.to_bytes(),
            pre_op_count.as_slice(),
            post_op_count.as_slice(),
            override_previous_root_bytes.as_slice(),
        ])
        .to_bytes()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // Last 8 bytes of keccak256("solana:localnet") as big-endian
    // This is 0x4808e31713a26612 --> in little-endian, it is "1266a21317e30848"
    const CHAIN_ID: u64 = 5190648258797659666;

    fn _decode<const N: usize>(s: &str) -> [u8; N] {
        hex::decode(s).unwrap().to_owned().try_into().unwrap()
    }

    fn decode32(s: &str) -> [u8; HASH_BYTES] {
        _decode::<HASH_BYTES>(s)
    }

    fn decode20(s: &str) -> [u8; EVM_ADDRESS_BYTES] {
        _decode::<EVM_ADDRESS_BYTES>(s)
    }

    mod domain_separators {
        use anchor_lang::solana_program::keccak;

        use super::*;

        #[test]
        fn verify_domain_separators() {
            let metadata =
                keccak::hash("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA_SOLANA".as_bytes());
            assert_eq!(
                &metadata.to_bytes(),
                MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA
            );
            let op = keccak::hash("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP_SOLANA".as_bytes());
            assert_eq!(&op.to_bytes(), MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP);
        }
    }

    mod test_hash_pair {
        use super::*;

        #[test]
        fn basic_keccak() {
            let a = &[0; 32];
            let b = &[0; 32];
            let result = hash_pair(a, b);

            assert_eq!(
                result,
                decode32("ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5")
            );
        }

        #[test]
        fn ordering() {
            let a = &[0; 32];
            let b = &[1; 32];

            let expected =
                decode32("d5f4f7e1d989848480236fb0a5f808d5877abf778364ae50845234dd6c1e80fc");

            assert_eq!(hash_pair(a, b), expected);
            assert_eq!(hash_pair(b, a), expected);
        }
    }

    mod test_left_pad_vec {
        use super::*;

        #[test]
        fn exact() {
            let input = [1, 2, 3];
            let result = _left_pad_vec(&input, input.len());
            assert_eq!(result.as_slice(), input);
        }

        #[test]
        fn single() {
            let input = [1, 2, 3];
            let result = _left_pad_vec(&input, input.len() + 1);
            assert_eq!(result.as_slice(), [0, 1, 2, 3]);
        }

        #[test]
        fn multiple() {
            let input = [1, 2, 3];
            let result = _left_pad_vec(&input, 32);
            let expected: [u8; 32] = [
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 1, 2, 3,
            ];
            assert_eq!(result.as_slice(), expected);
        }
    }
    mod test_compute_merkle_root {
        // All test values constructed from a tree with 3 leaves. See below "diagram" (all in hex):
        //
        // 0000000000000000000000000000000000000000000000000000000000000000
        //                 |
        //                 |--> 8e4b8e18156a1c7271055ce5b7ef53bb370294ebd631a3b95418a92da46e681f
        //                 |                                                 |
        // 1111111111111111111111111111111111111111111111111111111111111111  |
        //                                                                   |
        // 2222222222222222222222222222222222222222222222222222222222222222 ----> 888aba2887457beba19643fd1c5e5be943d3f0b910d418c1ab49c057c06f6738

        use super::*;

        #[test]
        fn identity() {
            let hashed_leaf = [0; HASH_BYTES];
            let proofs = vec![];
            let expected_root = [0; HASH_BYTES];

            let result = calculate_merkle_root(proofs, &hashed_leaf);
            assert_eq!(expected_root, result);
        }

        #[test]
        fn single_step() {
            let hashed_leaf =
                decode32("0000000000000000000000000000000000000000000000000000000000000000");

            let proofs = vec![decode32(
                "1111111111111111111111111111111111111111111111111111111111111111",
            )];
            let expected_root =
                decode32("8e4b8e18156a1c7271055ce5b7ef53bb370294ebd631a3b95418a92da46e681f");

            let result = calculate_merkle_root(proofs, &hashed_leaf);
            assert_eq!(expected_root, result);
        }

        #[test]
        fn multi_step() {
            let hashed_leaf =
                decode32("1111111111111111111111111111111111111111111111111111111111111111");
            let proofs = vec![
                decode32("0000000000000000000000000000000000000000000000000000000000000000"),
                decode32("2222222222222222222222222222222222222222222222222222222222222222"),
            ];

            let expected_root: [u8; HASH_BYTES] =
                decode32("888aba2887457beba19643fd1c5e5be943d3f0b910d418c1ab49c057c06f6738");

            let result = calculate_merkle_root(proofs, &hashed_leaf);
            assert_eq!(expected_root, result);
        }
    }

    mod test_compute_eth_message_hash {
        use super::*;

        #[test]
        fn basic() {
            let root =
                &decode32("d5ef592d1ad183db43b4980d7ab7ee43a6f6a284988c3e3a23d38c07beb520c7");
            let valid_until: u32 = 1748317727;

            let result = compute_eth_message_hash(root, valid_until);

            assert_eq!(
                result.to_bytes(),
                decode32("032705bd71839baef725154f00f87ddcc1d95c4b5189c9fb5983f26ad6c95102")
            );
        }
    }

    mod test_hash_leaf {
        use super::*;

        #[test]
        fn valid() {
            let md = RootMetadataInput {
                chain_id: CHAIN_ID,
                multisig: Pubkey::new_from_array(decode32(
                    "b870e12dd379891561d2e9fa8f26431834eb736f2f24fc2a2a4dff1fd5dca4df",
                )),
                pre_op_count: 0,
                post_op_count: 1,
                override_previous_root: false,
            };
            assert_eq!(
                md.hash_leaf(),
                decode32("c31496e313ba769f8c9f061dd35c6aa06c1c51ec9111f54be7be307a0de6b556")
            );

            // This is the metadata in hex:
            // 47fded70901d27038394db905a72563cad6f07581dbcdd1472ccd2f742af6360 <-- METADATA_DOMAIN_SEPARATOR
            // 0000000000000000000000000000000000000000000000001266a21317e30848 <-- chain_id
            // b870e12dd379891561d2e9fa8f26431834eb736f2f24fc2a2a4dff1fd5dca4df <-- multisig
            // 0000000000000000000000000000000000000000000000000000000000000000 <-- pre_op_count
            // 0000000000000000000000000000000000000000000000000100000000000000 <-- post_op_count
            // 0000000000000000000000000000000000000000000000000000000000000000 <-- override_previous_root
        }
    }

    mod test_ecdsa_recover_evm_addr {
        use super::*;

        #[test]
        fn valid() {
            let eth_signed_message_hash =
                &decode32("910cd291f5281f5bf25d8a83962f282b6c2bdf831f079dfcb84480f922abd2e1");
            let sig = &Signature {
                v: 28, // 1c
                r: decode32("45283a6239b1b559a910e97f79a52bab1605e8bd952c4b4e0720ed9b1e9e9671"),
                s: decode32("2acab6f5f946bfa3dfa61f47705aff6e2f17f6ad83d484857bb119a06ba1f0e7"),
            };
            let recovered_addr = ecdsa_recover_evm_addr(eth_signed_message_hash, sig);
            assert_eq!(
                recovered_addr.unwrap(),
                decode20("16c9fACed8a1e3C6aEA2B654EEca5617eb900EFf")
            );
        }
    }
}
