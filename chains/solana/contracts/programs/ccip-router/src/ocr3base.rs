use anchor_lang::prelude::*;
use anchor_lang::solana_program::sysvar;
use anchor_lang::solana_program::{keccak, secp256k1_recover::*};

#[constant]
pub const MAX_ORACLES: usize = 16; // can set a maximum of 16 transmitters + 16 signers simultaneously in a single set config tx
pub const MAX_SIGNERS: usize = MAX_ORACLES;
pub const MAX_TRANSMITTERS: usize = MAX_ORACLES;

pub const SIGNATURE_LENGTH: usize = SECP256K1_SIGNATURE_LENGTH + 1; // signature + recovery ID

pub const TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES: u128 = 8 // anchor discriminator
    + 3 * 32; // report context
pub const TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES: u128 = 4; // u32 length of signatures vec (borsh serialization)

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct ReportContext {
    // byte_words consists of:
    // [0]: ConfigDigest
    // [1]: 24 byte padding, 8 byte sequence number
    // [2]: ExtraHash
    byte_words: [[u8; 32]; 3], // private, define methods to use it
}

impl ReportContext {
    pub fn sequence_number(&self) -> u64 {
        let sequence_bytes: [u8; 8] = self.byte_words[1][24..].try_into().unwrap();
        u64::from_be_bytes(sequence_bytes)
    }

    pub fn from_byte_words(byte_words: [[u8; 32]; 3]) -> Self {
        Self { byte_words }
    }

    pub fn as_bytes(&self) -> [u8; 32 * 3] {
        self.byte_words.concat().as_slice().try_into().unwrap()
    }
}

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct Ocr3Config {
    pub plugin_type: u8, // plugin identifier for validation (example: ccip:commit = 0, ccip:execute = 1)
    pub config_info: Ocr3ConfigInfo,
    pub signers: [[u8; 20]; 16], // v0.29.0 - anchor IDL does not build with MAX_SIGNERS
    pub transmitters: [[u8; 32]; 16], // v0.29.0 - anchor IDL does not build with MAX_TRANSMITTERS
}

pub trait Ocr3Report {
    fn hash(&self, ctx: &ReportContext) -> [u8; 32];
    fn len(&self) -> usize;
}

// TODO: do we need to verify signers and transmitters are different? (between the two groups)
// signers: pubkey is 20-byte address, secp256k1 curve ECDSA
// transmitters: 32-byte pubkey, ed25519

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct Ocr3ConfigInfo {
    pub config_digest: [u8; 32], // 32-byte hash of configuration
    pub f: u8,                   // f+1 = number of signatures per report
    pub n: u8,                   // number of signers
    pub is_signature_verification_enabled: u8, // bool -> bytemuck::Pod compliant required for zero_copy
}

impl Ocr3Config {
    pub fn new(plugin_type: u8) -> Self {
        Self {
            plugin_type,
            ..Default::default()
        }
    }

    pub fn set(
        &mut self,
        plugin_type: u8,
        cfg: Ocr3ConfigInfo,
        signers: Vec<[u8; 20]>,    // 20-byte EVM address
        transmitters: Vec<Pubkey>, // 32-byte solana pubkey
    ) -> Result<()> {
        require!(
            plugin_type == self.plugin_type,
            Ocr3Error::InvalidPluginType
        );
        require!(cfg.f != 0, Ocr3Error::InvalidConfigFMustBePositive);

        // If F is 0, then the config is not yet set
        if self.config_info.f == 0 {
            self.config_info.is_signature_verification_enabled =
                cfg.is_signature_verification_enabled;
        } else {
            require!(
                self.config_info.is_signature_verification_enabled
                    == cfg.is_signature_verification_enabled,
                Ocr3Error::StaticConfigCannotBeChanged
            )
        };

        require!(
            transmitters.len() <= MAX_TRANSMITTERS,
            Ocr3Error::InvalidConfigTooManyTransmitters
        );

        self.clear_transmitters();

        if cfg.is_signature_verification_enabled != 0 {
            self.clear_signers();
            require!(
                signers.len() <= MAX_SIGNERS,
                Ocr3Error::InvalidConfigTooManySigners
            );
            require!(
                signers.len() > 3 * cfg.f as usize,
                Ocr3Error::InvalidConfigFIsTooHigh
            );

            self.config_info.n = signers.len() as u8;
            assign_oracles(&mut signers.clone(), &mut self.signers)?;
        }

        let transmitter_keys: Vec<[u8; 32]> = transmitters.iter().map(|t| t.to_bytes()).collect();
        assign_oracles(&mut transmitter_keys.clone(), &mut self.transmitters)?;

        // set remaining config
        self.config_info.f = cfg.f;
        self.config_info.config_digest = cfg.config_digest;

        emit!(ConfigSet {
            ocr_plugin_type: self.plugin_type,
            config_digest: self.config_info.config_digest,
            signers,
            transmitters,
            f: self.config_info.f,
        });

        Ok(())
    }

    fn clear_signers(&mut self) {
        self.signers = [[0; 20]; MAX_SIGNERS]
    }

    fn clear_transmitters(&mut self) {
        self.transmitters = [[0; 32]; MAX_TRANSMITTERS]
    }

    pub fn transmit<R: Ocr3Report>(
        &self,
        instruction_sysvar: &AccountInfo<'_>,
        transmitter: Pubkey,
        plugin_type: u8,
        report_context: ReportContext,
        report: &R,
        signatures: &[[u8; SIGNATURE_LENGTH]],
    ) -> Result<()> {
        require!(
            plugin_type == self.plugin_type,
            Ocr3Error::InvalidPluginType
        );

        // validate raw message length
        let mut expected_data_len: u128 = TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES
            .saturating_add(report.len() as u128);
        if self.config_info.is_signature_verification_enabled != 0 {
            expected_data_len = expected_data_len
                .saturating_add(TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES);
            expected_data_len =
                expected_data_len.saturating_add((signatures.len() * SIGNATURE_LENGTH) as u128);
        }
        // validates instruction sysvar is correct address
        let instruction_index =
            sysvar::instructions::load_current_index_checked(instruction_sysvar)?;
        let tx = sysvar::instructions::load_instruction_at_checked(
            instruction_index as usize,
            instruction_sysvar,
        )?;

        require!(
            tx.data.len() as u128 == expected_data_len,
            Ocr3Error::WrongMessageLength
        );

        require!(
            self.config_info.config_digest == report_context.byte_words[0],
            Ocr3Error::ConfigDigestMismatch
        );

        // TODO: chain fork check - is this even possible?

        // validate transmitter
        self.transmitters
            .binary_search_by(|probe| transmitter.to_bytes().cmp(probe))
            .map_err(|_| Ocr3Error::UnauthorizedTransmitter)?;

        if self.config_info.is_signature_verification_enabled != 0 {
            require!(
                signatures.len() == (self.config_info.f + 1) as usize,
                Ocr3Error::WrongNumberOfSignatures
            );

            self.verify_signatures(report.hash(&report_context), signatures)?;
        }

        emit!(Transmitted {
            ocr_plugin_type: self.plugin_type,
            config_digest: self.config_info.config_digest,
            sequence_number: report_context.sequence_number(),
        });

        Ok(())
    }

    fn verify_signatures(
        &self,
        hashed_report: [u8; 32],
        signatures: &[[u8; SIGNATURE_LENGTH]],
    ) -> Result<()> {
        let mut uniques: u16 = 0;
        assert!(uniques.count_ones() + uniques.count_zeros() >= MAX_SIGNERS as u32); // ensure MAX_SIGNERS fit in the bits of uniques

        for raw_sig in signatures {
            let (recovery_id, signature) =
                raw_sig.split_first().ok_or(Ocr3Error::InvalidSignature)?;

            let signer = secp256k1_recover(&hashed_report, *recovery_id, signature)
                .map_err(|_| Ocr3Error::InvalidSignature)?;
            // convert to a raw 20 byte Ethereum address
            let address: [u8; 20] = keccak::hash(&signer.0).to_bytes()[12..32]
                .try_into()
                .map_err(|_| Ocr3Error::UnauthorizedSigner)?;
            let index = self
                .signers
                .binary_search_by(|probe| address.cmp(probe))
                .map_err(|_| Ocr3Error::UnauthorizedSigner)?;

            uniques |= 1 << index;
        }

        require!(
            uniques.count_ones() as usize == signatures.len(),
            Ocr3Error::NonUniqueSignatures
        );

        Ok(())
    }
}

// assign_oracles (generic) takes a vector of byte arrays sorts (for binary searching), validates, and writes it to `location`
fn assign_oracles<const A: usize>(oracles: &mut [[u8; A]], location: &mut [[u8; A]]) -> Result<()> {
    // sort + verify not duplicated
    oracles.sort_unstable_by(|a, b| b.cmp(a)); // reverse sort to ensure appended 0-values in correct
    let duplicate = oracles.windows(2).any(|pair| pair[0] == pair[1]);
    require!(!duplicate, Ocr3Error::InvalidConfigRepeatedOracle);

    for (i, o) in oracles.iter().enumerate() {
        require!(*o != [0; A], Ocr3Error::OracleCannotBeZeroAddress);
        location[i] = *o;
    }
    Ok(())
}

#[error_code]
pub enum Ocr3Error {
    #[msg("Invalid config: F must be positive")]
    InvalidConfigFMustBePositive,
    #[msg("Invalid config: Too many transmitters")]
    InvalidConfigTooManyTransmitters,
    #[msg("Invalid config: Too many signers")]
    InvalidConfigTooManySigners,
    #[msg("Invalid config: F is too high")]
    InvalidConfigFIsTooHigh,
    #[msg("Invalid config: Repeated oracle address")]
    InvalidConfigRepeatedOracle,
    #[msg("Wrong message length")]
    WrongMessageLength,
    #[msg("Config digest mismatch")]
    ConfigDigestMismatch,
    #[msg("Wrong number signatures")]
    WrongNumberOfSignatures,
    #[msg("Unauthorized transmitter")]
    UnauthorizedTransmitter,
    #[msg("Unauthorized signer")]
    UnauthorizedSigner,
    #[msg("Non unique signatures")]
    NonUniqueSignatures,
    #[msg("Oracle cannot be zero address")]
    OracleCannotBeZeroAddress,
    #[msg("Static config cannot be changed")]
    StaticConfigCannotBeChanged,
    #[msg("Incorrect plugin type")]
    InvalidPluginType,
    #[msg("Invalid signature")]
    InvalidSignature,
}

#[event]
pub struct ConfigSet {
    pub ocr_plugin_type: u8,
    pub config_digest: [u8; 32],
    pub signers: Vec<[u8; 20]>,
    pub transmitters: Vec<Pubkey>,
    pub f: u8,
}

#[event]
pub struct Transmitted {
    pub ocr_plugin_type: u8,
    pub config_digest: [u8; 32],
    pub sequence_number: u64,
}
