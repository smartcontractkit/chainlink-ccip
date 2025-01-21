use anchor_lang::solana_program::keccak;

pub const LEAF_DOMAIN_SEPARATOR: [u8; 32] = [0; 32];
const MAX_NUM_HASHES: usize = 128; // TODO: Change this to 256 when supporting commit reports with 256 messages

#[derive(Debug)]
pub enum MerkleError {
    InvalidProof,
}

fn hash_pair(hash1: &[u8; 32], hash2: &[u8; 32]) -> [u8; 32] {
    if hash1 < hash2 {
        keccak::hashv(&[hash1, hash2]).to_bytes()
    } else {
        keccak::hashv(&[hash2, hash1]).to_bytes()
    }
}

pub fn calculate_merkle_root(
    hashed_leaf: [u8; 32],
    proofs: Vec<[u8; 32]>,
) -> Result<[u8; 32], MerkleError> {
    let proofs_len = proofs.len();

    if proofs_len > MAX_NUM_HASHES {
        return Err(MerkleError::InvalidProof);
    }

    let mut hash = hashed_leaf;

    for p in proofs {
        hash = hash_pair(&hash, &p);
    }

    Ok(hash)
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_calculate_merkle_root_valid() {
        let hexa_leaf = "7e1ff3c10bacb7a70bd9dbaa1b2ddeb4c860c6db3c3557d31baff96222505e2a";
        let hashed_leaf: [u8; 32] = hex::decode(hexa_leaf)
            .unwrap()
            .to_owned()
            .try_into()
            .unwrap();

        let proofs =
            vec![
                hex::decode("22ae9b57dfb3f830622fb5ee07a795961532dc9ab7f641271ac7cf1b89cb39f6")
                    .unwrap()
                    .to_owned()
                    .try_into()
                    .unwrap(),
            ];
        let expected_root: [u8; 32] =
            hex::decode("94b949ca8fd6307aa72481fe44eca36c63686f8e85acac99e4f0cd2b36a99d33")
                .unwrap()
                .to_owned()
                .try_into()
                .unwrap();

        let result = calculate_merkle_root(hashed_leaf, proofs);
        assert!(result.is_ok());
        assert_eq!(hex::encode(expected_root), hex::encode(result.unwrap()));
    }

    #[test]
    fn test_calculate_merkle_root_from_size_1() {
        let hexa_leaf = "7e1ff3c10bacb7a70bd9dbaa1b2ddeb4c860c6db3c3557d31baff96222505e2a";
        let hashed_leaf: [u8; 32] = hex::decode(hexa_leaf)
            .unwrap()
            .to_owned()
            .try_into()
            .unwrap();

        let proofs = vec![];
        let expected_root: [u8; 32] =
            hex::decode("7e1ff3c10bacb7a70bd9dbaa1b2ddeb4c860c6db3c3557d31baff96222505e2a")
                .unwrap()
                .to_owned()
                .try_into()
                .unwrap();

        let result = calculate_merkle_root(hashed_leaf, proofs);
        assert!(result.is_ok());
        assert_eq!(hex::encode(expected_root), hex::encode(result.unwrap()));
    }

    #[test]
    fn test_calculate_merkle_root_invalid_proof() {
        let hexa_leaf = "7e1ff3c10bacb7a70bd9dbaa1b2ddeb4c860c6db3c3557d31baff96222505e2a";
        let hashed_leaf: [u8; 32] = hex::decode(hexa_leaf)
            .unwrap()
            .to_owned()
            .try_into()
            .unwrap();
        let proofs = vec![[0x44; 32]; 129]; // Array size greater than 128

        let result = calculate_merkle_root(hashed_leaf, proofs);
        assert!(result.is_err());
    }

    #[test]
    fn test_create_merkle_validate_proof() {
        let a = hex::decode("1ac2f192702849e03dfe5c31ec66a4f6408b5eb16cc02f1583ce713b22be92ed")
            .unwrap()
            .to_owned()
            .try_into()
            .unwrap();
        let proofs = vec![
            hex::decode("81caa6284d8f53a5cd06190bd30c33c4622e6dde8b1204e953f98d768eeab615")
                .unwrap()
                .to_owned()
                .try_into()
                .unwrap(),
            hex::decode("16e4eb7487ed6b9476a0aca9294a0d5e6c7fe7bf7d5ca71908d2e46802843135")
                .unwrap()
                .to_owned()
                .try_into()
                .unwrap(),
        ];

        let expected_root: [u8; 32] =
            hex::decode("577252413aa3c3c02bca5a8e30ad69fdf1b138d4ccc3d834d3c6934775ceaf87")
                .unwrap()
                .to_owned()
                .try_into()
                .unwrap();

        let result = calculate_merkle_root(a, proofs);
        assert!(result.is_ok());
        assert_eq!(hex::encode(expected_root), hex::encode(result.unwrap()));
    }
}
