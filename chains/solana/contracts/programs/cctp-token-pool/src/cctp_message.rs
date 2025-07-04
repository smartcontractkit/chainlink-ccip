use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct CctpMessage {
    pub data: Vec<u8>,
}

impl CctpMessage {
    const MAX_NONCES: u64 = 6400;

    // Indices of each field in the message
    const VERSION: (usize, usize) = (0, 4); // u32 (4 bytes)
    const SOURCE_DOMAIN: (usize, usize) = (4, 8); // u32 (4 bytes)
    const DESTINATION_DOMAIN: (usize, usize) = (8, 12); // u32 (4 bytes)
    const NONCE: (usize, usize) = (12, 20); // u64 (8 bytes)
    const SENDER: (usize, usize) = (20, 52); // Pubkey (32 bytes)
    const RECIPIENT: (usize, usize) = (52, 84); // Pubkey (32 bytes)
    const DESTINATION_CALLER: (usize, usize) = (84, 116); // Pubkey (32 bytes)
    const MESSAGE_BODY_START: usize = 116;

    // From CCIP's point of view, the CCTP message must be at least long enough to contain the metadata fields,
    // so that CCIP's own checks can be performed without panicking. Whenever CCTP is invoked, CCTP will perform
    // further checks on the message.
    pub const MINIMUM_LENGTH: usize = Self::MESSAGE_BODY_START;

    pub fn version(&self) -> u32 {
        u32::from_be_bytes(
            self.data[Self::VERSION.0..Self::VERSION.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn source_domain(&self) -> u32 {
        u32::from_be_bytes(
            self.data[Self::SOURCE_DOMAIN.0..Self::SOURCE_DOMAIN.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn destination_domain(&self) -> u32 {
        u32::from_be_bytes(
            self.data[Self::DESTINATION_DOMAIN.0..Self::DESTINATION_DOMAIN.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn nonce(&self) -> u64 {
        u64::from_be_bytes(self.data[Self::NONCE.0..Self::NONCE.1].try_into().unwrap())
    }

    pub fn sender(&self) -> Pubkey {
        Pubkey::new_from_array(
            self.data[Self::SENDER.0..Self::SENDER.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn recipient(&self) -> Pubkey {
        Pubkey::new_from_array(
            self.data[Self::RECIPIENT.0..Self::RECIPIENT.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn destination_caller(&self) -> Pubkey {
        Pubkey::new_from_array(
            self.data[Self::DESTINATION_CALLER.0..Self::DESTINATION_CALLER.1]
                .try_into()
                .unwrap(),
        )
    }

    pub fn first_nonce(&self) -> u64 {
        self.nonce()
            .checked_sub(1)
            .and_then(|n| n.checked_div(Self::MAX_NONCES))
            .and_then(|n| n.checked_mul(Self::MAX_NONCES))
            .and_then(|n| n.checked_add(1))
            .unwrap()
    }
}
