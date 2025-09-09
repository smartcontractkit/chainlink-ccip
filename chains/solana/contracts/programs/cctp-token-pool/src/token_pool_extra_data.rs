use anchor_lang::prelude::*;

use crate::CctpTokenPoolError;

pub struct TokenPoolExtraData {
    pub nonce: u64,         // The nonce of the message being locked or burned
    pub source_domain: u32, // The source chain domain ID, which for Solana is always 5
}

impl TokenPoolExtraData {
    // ABI-encoding left-pads each number to 32-bytes, and uses big-endian encoding.

    // The nonce is in bytes 24..32 in the serialized data (u64 is 8 bytes)
    const NONCE_INDEXES: (usize, usize) = (24, 32);
    // The source domain is in bytes 60..64 in the serialized data (u32 is 4 bytes)
    const SOURCE_DOMAIN_INDEXES: (usize, usize) = (60, 64);
    // Total size of the serialized data (2 fields of 32 bytes each)
    const TOTAL_SIZE: usize = 64;

    pub fn abi_encode(&self) -> [u8; 64] {
        let mut bytes = [0u8; 64];
        bytes[Self::NONCE_INDEXES.0..Self::NONCE_INDEXES.1]
            .copy_from_slice(&self.nonce.to_be_bytes());
        bytes[Self::SOURCE_DOMAIN_INDEXES.0..Self::SOURCE_DOMAIN_INDEXES.1]
            .copy_from_slice(&self.source_domain.to_be_bytes());
        bytes
    }

    pub fn abi_decode(bytes: &[u8]) -> Result<Self> {
        require_eq!(
            bytes.len(),
            Self::TOTAL_SIZE,
            CctpTokenPoolError::InvalidTokenPoolExtraData
        );

        // First value (nonce) in ABI-encoding is left-padded with zeros,
        // so we check that the first bytes are all zero.
        require!(
            bytes[0..Self::NONCE_INDEXES.0] == [0u8; Self::NONCE_INDEXES.0],
            CctpTokenPoolError::InvalidTokenPoolExtraData
        );
        // Then, the last bytes in the first 32 bytes are the actual nonce (big-endian).
        let nonce = u64::from_be_bytes(
            bytes[Self::NONCE_INDEXES.0..Self::NONCE_INDEXES.1]
                .try_into()
                .unwrap(),
        );

        // Second value (source_domain) in ABI-encoding is left-padded with zeros,
        // so we check that the first bytes are all zero.
        require!(
            bytes[Self::NONCE_INDEXES.1..Self::SOURCE_DOMAIN_INDEXES.0]
                == [0u8; Self::SOURCE_DOMAIN_INDEXES.0 - Self::NONCE_INDEXES.1],
            CctpTokenPoolError::InvalidTokenPoolExtraData
        );
        // Then, the last bytes in the last 32 bytes are the actual source domain (big-endian).
        let source_domain = u32::from_be_bytes(
            bytes[Self::SOURCE_DOMAIN_INDEXES.0..Self::SOURCE_DOMAIN_INDEXES.1]
                .try_into()
                .unwrap(),
        );

        Ok(Self {
            nonce,
            source_domain,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_abi_encode_decode() {
        let extra_data = TokenPoolExtraData {
            nonce: 123456789,
            source_domain: 5,
        };
        let encoded = extra_data.abi_encode();
        let decoded = TokenPoolExtraData::abi_decode(&encoded).unwrap();
        assert_eq!(decoded.nonce, extra_data.nonce);
        assert_eq!(decoded.source_domain, extra_data.source_domain);
    }

    #[test]
    fn test_abi_decode_with_invalid_length() {
        let invalid_length_data = [0u8; 63]; // One byte short
        let result = TokenPoolExtraData::abi_decode(&invalid_length_data);
        assert!(result.is_err());
    }

    #[test]
    fn test_abi_decode_with_invalid_padding() {
        let mut invalid_data = [0u8; 64];
        invalid_data[20] = 1; // Change one byte in the nonce padding
        let result = TokenPoolExtraData::abi_decode(&invalid_data);
        assert!(result.is_err());
    }

    #[test]
    fn test_abi_decode_with_invalid_source_domain_padding() {
        let mut invalid_data = [0u8; 64];
        invalid_data[50] = 1; // Change one byte in the source domain padding
        let result = TokenPoolExtraData::abi_decode(&invalid_data);
        assert!(result.is_err());
    }
}
