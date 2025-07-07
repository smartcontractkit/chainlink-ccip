use anchor_lang::prelude::Result;

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

    pub fn abi_encode(&self) -> [u8; 64] {
        let mut bytes = [0u8; 64];
        bytes[Self::NONCE_INDEXES.0..Self::NONCE_INDEXES.1]
            .copy_from_slice(&self.nonce.to_be_bytes());
        bytes[Self::SOURCE_DOMAIN_INDEXES.0..Self::SOURCE_DOMAIN_INDEXES.1]
            .copy_from_slice(&self.source_domain.to_be_bytes());
        bytes
    }

    pub fn abi_decode(bytes: &[u8]) -> Result<Self> {
        if bytes.len() != 64 {
            return Err(CctpTokenPoolError::InvalidTokenPoolExtraData.into());
        }
        let nonce = u64::from_be_bytes(
            bytes[Self::NONCE_INDEXES.0..Self::NONCE_INDEXES.1]
                .try_into()
                .unwrap(),
        );
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
