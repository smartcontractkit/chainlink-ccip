use anchor_lang::prelude::*;
use solana_program::keccak::Hash;

use crate::context::DISCRIMINATOR_SIZE;
use crate::message::Message;
use crate::CctpMessageTransmitterMockError;

#[account]
#[derive(Debug, InitSpace)]
pub struct MessageSent {
    pub rent_payer: Pubkey,
    #[max_len(1)]
    pub message: Vec<u8>,
}

impl MessageSent {
    pub fn len(message_body_len: usize) -> Result<usize> {
        DISCRIMINATOR_SIZE
            .checked_add(MessageSent::INIT_SPACE)
            .and_then(|x| x.checked_sub(1))
            .and_then(|x| x.checked_add(Message::serialized_len(message_body_len).unwrap()))
            .ok_or(CctpMessageTransmitterMockError::MathError.into())
    }
}

#[account]
#[derive(Debug, InitSpace)]
/// Main state of the MessageTransmitter program
pub struct MessageTransmitter {
    pub owner: Pubkey,
    pub pending_owner: Pubkey,
    pub attester_manager: Pubkey,
    pub pauser: Pubkey,
    pub paused: bool,
    pub local_domain: u32,
    pub version: u32,
    pub signature_threshold: u32,
    #[max_len(1)]
    pub enabled_attesters: Vec<Pubkey>,
    pub max_message_body_size: u64,
    pub next_available_nonce: u64,
}

#[account]
#[derive(Debug, InitSpace)]
pub struct UsedNonces {
    pub remote_domain: u32,
    pub first_nonce: u64,
    used_nonces: [u64; 100], // length = MAX_NONCES / 64 + if MAX_NONCES % 64 != 0 { 1 } else { 0 }
}

impl MessageTransmitter {
    pub const ATTESTATION_SIGNATURE_LENGTH: usize = 65;

    /// Checks if the state is valid
    pub fn validate(&self) -> bool {
        self.owner != Pubkey::default()
            && self.attester_manager != Pubkey::default()
            && self.pauser != Pubkey::default()
            && self.signature_threshold != 0
            && self.signature_threshold as usize <= self.enabled_attesters.len()
            && !self.enabled_attesters.is_empty()
            && self.next_available_nonce > 0
    }

    pub fn verify_attestation_signatures(
        &self,
        message_hash: &Hash,
        attestation: &Vec<u8>,
    ) -> Result<()> {
        require_eq!(
            attestation.len(),
            Self::ATTESTATION_SIGNATURE_LENGTH
                .checked_mul(self.signature_threshold as usize)
                .unwrap(),
            CctpMessageTransmitterMockError::InvalidAttestationLength
        );

        let mut last_attester = Pubkey::default();

        for i in 0..self.signature_threshold as usize {
            // slice next signature
            let signature = &attestation.as_slice()[(i * Self::ATTESTATION_SIGNATURE_LENGTH)
                ..((i * Self::ATTESTATION_SIGNATURE_LENGTH) + Self::ATTESTATION_SIGNATURE_LENGTH)];

            // recover attester's address from the message hash and the signature
            let recovered_attester =
                MessageTransmitter::recover_attester(message_hash.0.as_slice(), signature)?;

            if recovered_attester <= last_attester {
                return err!(CctpMessageTransmitterMockError::InvalidSignatureOrderOrDupe);
            }

            // check if the recovered attester is enabled
            require!(
                self.is_enabled_attester(&recovered_attester),
                CctpMessageTransmitterMockError::InvalidAttesterSignature
            );

            last_attester = recovered_attester;
        }

        Ok(())
    }

    /// Checks if the attester is enabled
    pub fn is_enabled_attester(&self, attester: &Pubkey) -> bool {
        self.enabled_attesters.contains(attester)
    }

    /// Recovers attester's address from the message hash and the signature
    fn recover_attester(_message_hash: &[u8], _attestation_signature: &[u8]) -> Result<Pubkey> {
        // TODO check if needed
        Err(CctpMessageTransmitterMockError::InvalidAttesterManager.into())
        // // secp256k1_recover doesn't validate input parameters lengths, so manual check is needed
        // require_eq!(
        //     message_hash.len(),
        //     32,
        //     CctpMessageTransmitterMockError::InvalidMessageHash
        // );
        // require_eq!(
        //     attestation_signature.len(),
        //     65,
        //     CctpMessageTransmitterMockError::InvalidAttesterSignature
        // );

        // // extract recovery id from the signature
        // let recovery_id = attestation_signature[Self::ATTESTATION_SIGNATURE_LENGTH - 1];
        // require!(
        //     (27..=30).contains(&recovery_id),
        //     CctpMessageTransmitterMockError::InvalidSignatureRecoveryId
        // );

        // // reject high-s value signatures to prevent malleability
        // let signature = EVMSignature::parse_standard_slice(
        //     &attestation_signature[0..Self::ATTESTATION_SIGNATURE_LENGTH - 1],
        // )
        // .map_err(|_| CctpMessageTransmitterMock::InvalidAttesterSignature)?;
        // require!(
        //     !signature.s.is_high(),
        //     CctpMessageTransmitterMock::InvalidSignatureSValue
        // );

        // // recover attester's public key
        // let pubkey = secp256k1_recover(
        //     message_hash,
        //     recovery_id - 27,
        //     &attestation_signature[0..Self::ATTESTATION_SIGNATURE_LENGTH - 1],
        // )
        // .map_err(|err| match err {
        //     Secp256k1RecoverError::InvalidHash => {
        //         CctpMessageTransmitterMockError::InvalidMessageHash
        //     }
        //     Secp256k1RecoverError::InvalidRecoveryId => {
        //         CctpMessageTransmitterMockError::InvalidSignatureRecoveryId
        //     }
        //     Secp256k1RecoverError::InvalidSignature => {
        //         CctpMessageTransmitterMockError::InvalidAttesterSignature
        //     }
        // })?;

        // // hash public key and return last 20 bytes as Pubkey
        // let mut hasher = Hasher::default();
        // hasher.hash(pubkey.to_bytes().as_slice());
        // let mut address = hasher.result().0;
        // address[0..12].iter_mut().for_each(|x| *x = 0);

        // Ok(Pubkey::new_from_array(address))
    }
}

impl UsedNonces {
    pub const MAX_NONCES: usize = 6400;

    /// Returns the first nonce in the UsedNonces account corresponding to the given nonce.
    /// To minimize on-chain space use, used nonces are stored in a bitset, i.e. 0/1 flags
    /// indicating if the nonce was already used. `first_nonce` represents the very first
    /// nonce in the set. For example, given the first_nonce = 100 and used_nonces set = 0110,
    /// we can tell that nonces 101 and 102 are used, while 100 and 103 are not. Bitset is
    /// implemented as a fixed-size array of 64 bit values that are custom indexed.
    pub fn first_nonce(nonce: u64) -> Result<u64> {
        if nonce == 0 {
            return err!(CctpMessageTransmitterMockError::InvalidNonce);
        }
        nonce
            .checked_sub(1)
            .and_then(|n| n.checked_div(Self::MAX_NONCES as u64))
            .and_then(|n| n.checked_mul(Self::MAX_NONCES as u64))
            .and_then(|n| n.checked_add(1))
            .ok_or(CctpMessageTransmitterMockError::MathError.into())
    }

    /// Marks the nonce as used
    pub fn use_nonce(&mut self, nonce: u64) -> Result<()> {
        let (entry, bit) = self.get_entry_bit(nonce)?;

        require!(
            self.used_nonces[entry] & bit == 0,
            CctpMessageTransmitterMockError::NonceAlreadyUsed
        );

        self.used_nonces[entry] |= bit;

        Ok(())
    }

    /// Checks if nonce is used
    pub fn is_nonce_used(&self, nonce: u64) -> Result<bool> {
        let (entry, bit) = self.get_entry_bit(nonce)?;

        Ok(self.used_nonces[entry] & bit != 0)
    }

    fn get_entry_bit(&self, nonce: u64) -> Result<(usize, u64)> {
        require!(
            nonce >= self.first_nonce && nonce < self.first_nonce + Self::MAX_NONCES as u64,
            CctpMessageTransmitterMockError::InvalidNonce
        );

        let position = nonce.checked_sub(self.first_nonce).unwrap() as usize;
        let entry = position.checked_div(64).unwrap();
        let bit = 1 << (position as u64 % 64);

        Ok((entry, bit))
    }

    /// Adds a delimiter for the used_nonces PDA seeds for domains >= 11
    /// Only add to domains >= 11 to prevent existing (pre-upgrade on mainnet)
    /// PDAs from changing.
    pub fn used_nonces_seed_delimiter(source_domain: u32) -> &'static [u8] {
        if source_domain < 11 {
            b""
        } else {
            b"-"
        }
    }
}
