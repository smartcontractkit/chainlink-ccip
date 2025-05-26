// Based on CCTP's message.rs inside the message-transmitter program

use anchor_lang::prelude::*;
use solana_program::keccak::{Hash, Hasher};

use crate::CctpMessageTransmitterMockError;

#[derive(Clone, Debug)]
pub struct Message<'a> {
    data: &'a [u8],
}

impl<'a> Message<'a> {
    // Indices of each field in the message
    const VERSION_INDEX: usize = 0;
    const SOURCE_DOMAIN_INDEX: usize = 4;
    const DESTINATION_DOMAIN_INDEX: usize = 8;
    const NONCE_INDEX: usize = 12;
    const SENDER_INDEX: usize = 20;
    const RECIPIENT_INDEX: usize = 52;
    const DESTINATION_CALLER_INDEX: usize = 84;
    const MESSAGE_BODY_INDEX: usize = 116;

    /// Validates source array size and returns a new message
    pub fn new(expected_version: u32, message_bytes: &'a [u8]) -> Result<Self> {
        require_gte!(
            message_bytes.len(),
            Self::MESSAGE_BODY_INDEX,
            CctpMessageTransmitterMockError::MalformedMessage
        );
        let message = Self {
            data: message_bytes,
        };
        require_eq!(
            expected_version,
            message.version()?,
            CctpMessageTransmitterMockError::InvalidMessageVersion
        );
        Ok(message)
    }

    pub fn serialized_len(message_body_len: usize) -> Result<usize> {
        Self::MESSAGE_BODY_INDEX
            .checked_add(message_body_len)
            .ok_or(CctpMessageTransmitterMockError::MalformedMessage.into())
    }

    #[allow(clippy::too_many_arguments)]
    /// Serializes given fields into a message
    pub fn format_message(
        version: u32,
        local_domain: u32,
        destination_domain: u32,
        nonce: u64,
        sender: &Pubkey,
        recipient: &Pubkey,
        destination_caller: &Pubkey,
        message_body: &Vec<u8>,
    ) -> Result<Vec<u8>> {
        let mut output = vec![0; Message::serialized_len(message_body.len())?];

        output[Self::VERSION_INDEX..Self::SOURCE_DOMAIN_INDEX]
            .copy_from_slice(&version.to_be_bytes());
        output[Self::SOURCE_DOMAIN_INDEX..Self::DESTINATION_DOMAIN_INDEX]
            .copy_from_slice(&local_domain.to_be_bytes());
        output[Self::DESTINATION_DOMAIN_INDEX..Self::NONCE_INDEX]
            .copy_from_slice(&destination_domain.to_be_bytes());
        output[Self::NONCE_INDEX..Self::SENDER_INDEX].copy_from_slice(&nonce.to_be_bytes());
        output[Self::SENDER_INDEX..Self::RECIPIENT_INDEX].copy_from_slice(sender.as_ref());
        output[Self::RECIPIENT_INDEX..Self::DESTINATION_CALLER_INDEX]
            .copy_from_slice(recipient.as_ref());
        output[Self::DESTINATION_CALLER_INDEX..Self::MESSAGE_BODY_INDEX]
            .copy_from_slice(destination_caller.as_ref());
        if !message_body.is_empty() {
            output[Self::MESSAGE_BODY_INDEX..].copy_from_slice(message_body.as_slice());
        }

        Ok(output)
    }

    /// Returns Keccak hash of the message
    pub fn hash(&self) -> Hash {
        let mut hasher = Hasher::default();
        hasher.hash(self.data);
        hasher.result()
    }

    /// Returns version field
    pub fn version(&self) -> Result<u32> {
        self.read_u32(Self::VERSION_INDEX)
    }

    /// Returns sender field
    pub fn sender(&self) -> Result<Pubkey> {
        self.read_pubkey(Self::SENDER_INDEX)
    }

    /// Returns recipient field
    pub fn recipient(&self) -> Result<Pubkey> {
        self.read_pubkey(Self::RECIPIENT_INDEX)
    }

    /// Returns source_domain field
    pub fn source_domain(&self) -> Result<u32> {
        self.read_u32(Self::SOURCE_DOMAIN_INDEX)
    }

    /// Returns destination_domain field
    pub fn destination_domain(&self) -> Result<u32> {
        self.read_u32(Self::DESTINATION_DOMAIN_INDEX)
    }

    /// Returns destination_caller field
    pub fn destination_caller(&self) -> Result<Pubkey> {
        self.read_pubkey(Self::DESTINATION_CALLER_INDEX)
    }

    /// Returns nonce field
    pub fn nonce(&self) -> Result<u64> {
        self.read_u64(Self::NONCE_INDEX)
    }

    /// Returns message_body field
    pub fn message_body(&self) -> &[u8] {
        &self.data[Self::MESSAGE_BODY_INDEX..]
    }

    ////////////////////
    // private helpers

    fn read_u32(&self, index: usize) -> Result<u32> {
        let end = index.checked_add(4).unwrap();
        Ok(u32::from_be_bytes(
            self.data[index..end]
                .try_into()
                .map_err(|_| CctpMessageTransmitterMockError::MalformedMessage)?,
        ))
    }

    fn read_u64(&self, index: usize) -> Result<u64> {
        let end = index.checked_add(8).unwrap();
        Ok(u64::from_be_bytes(
            self.data[index..end]
                .try_into()
                .map_err(|_| CctpMessageTransmitterMockError::MalformedMessage)?,
        ))
    }

    /// Reads pubkey field at the given offset
    fn read_pubkey(&self, index: usize) -> Result<Pubkey> {
        let end = index.checked_add(32).unwrap();
        Ok(Pubkey::try_from(&self.data[index..end])
            .map_err(|_| CctpMessageTransmitterMockError::MalformedMessage)?)
    }
}
