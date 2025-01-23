use anchor_lang::prelude::*;
use anchor_lang::solana_program::sysvar;
use anchor_lang::solana_program::{keccak, secp256k1_recover::*};

use crate::ocr3base::{ConfigSet, Ocr3Error, Transmitted};
use crate::state::{Ocr3Config, Ocr3ConfigInfo};

#[constant]
pub const MAX_ORACLES: usize = 16; // can set a maximum of 16 transmitters + 16 signers simultaneously in a single set config tx
pub const MAX_SIGNERS: usize = MAX_ORACLES;
pub const MAX_TRANSMITTERS: usize = MAX_ORACLES;

pub const SIGNATURE_LENGTH: usize = SECP256K1_SIGNATURE_LENGTH + 1; // signature + recovery ID

pub const TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES: u128 = 8 // anchor discriminator
    + 2 // ccip version (major & minor, 1 byte each)
    + 3 * 32; // report context
pub const TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES: u128 = 4; // u32 length of signatures vec (borsh serialization)

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub(super) struct ReportContext {
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

pub fn ocr3_set(
    ocr3_config: &mut Ocr3Config,
    plugin_type: u8,
    cfg: Ocr3ConfigInfo,
    signers: Vec<[u8; 20]>,
    transmitters: Vec<Pubkey>,
) -> Result<()> {
    require!(
        plugin_type == ocr3_config.plugin_type,
        Ocr3Error::InvalidPluginType
    );
    require!(cfg.f != 0, Ocr3Error::InvalidConfigFMustBePositive);

    // If F is 0, then the config is not yet set
    if ocr3_config.config_info.f == 0 {
        ocr3_config.config_info.is_signature_verification_enabled =
            cfg.is_signature_verification_enabled;
    } else {
        require!(
            ocr3_config.config_info.is_signature_verification_enabled
                == cfg.is_signature_verification_enabled,
            Ocr3Error::StaticConfigCannotBeChanged
        )
    };

    require!(
        transmitters.len() <= MAX_TRANSMITTERS,
        Ocr3Error::InvalidConfigTooManyTransmitters
    );

    clear_transmitters(ocr3_config);

    if cfg.is_signature_verification_enabled != 0 {
        clear_signers(ocr3_config);
        require!(
            signers.len() <= MAX_SIGNERS,
            Ocr3Error::InvalidConfigTooManySigners
        );
        require!(
            signers.len() > 3 * cfg.f as usize,
            Ocr3Error::InvalidConfigFIsTooHigh
        );

        ocr3_config.config_info.n = signers.len() as u8;
        assign_oracles(&mut signers.clone(), &mut ocr3_config.signers)?;
    }

    let transmitter_keys: Vec<[u8; 32]> = transmitters.iter().map(|t| t.to_bytes()).collect();
    assign_oracles(&mut transmitter_keys.clone(), &mut ocr3_config.transmitters)?;

    // set remaining config
    ocr3_config.config_info.f = cfg.f;
    ocr3_config.config_info.config_digest = cfg.config_digest;

    emit!(ConfigSet {
        ocr_plugin_type: ocr3_config.plugin_type,
        config_digest: ocr3_config.config_info.config_digest,
        signers,
        transmitters,
        f: ocr3_config.config_info.f,
    });

    Ok(())
}

fn clear_signers(ocr3_config: &mut Ocr3Config) {
    ocr3_config.signers = [[0; 20]; MAX_SIGNERS]
}

fn clear_transmitters(ocr3_config: &mut Ocr3Config) {
    ocr3_config.transmitters = [[0; 32]; MAX_TRANSMITTERS]
}

pub(super) trait Ocr3Report {
    fn hash(&self, ctx: &ReportContext) -> [u8; 32];
    fn len(&self) -> usize;
}

pub(super) fn ocr3_transmit<R: Ocr3Report>(
    ocr3_config: &Ocr3Config,
    instruction_sysvar: &AccountInfo<'_>,
    transmitter: Pubkey,
    plugin_type: u8,
    report_context: ReportContext,
    report: &R,
    signatures: &[[u8; SIGNATURE_LENGTH]],
) -> Result<()> {
    require!(
        plugin_type == ocr3_config.plugin_type,
        Ocr3Error::InvalidPluginType
    );

    // validate raw message length
    let mut expected_data_len: u128 = TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES
        .saturating_add(report.len() as u128);
    if ocr3_config.config_info.is_signature_verification_enabled != 0 {
        expected_data_len = expected_data_len
            .saturating_add(TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES);
        expected_data_len =
            expected_data_len.saturating_add((signatures.len() * SIGNATURE_LENGTH) as u128);
    }
    // validates instruction sysvar is correct address
    let instruction_index = sysvar::instructions::load_current_index_checked(instruction_sysvar)?;
    let tx = sysvar::instructions::load_instruction_at_checked(
        instruction_index as usize,
        instruction_sysvar,
    )?;

    require_eq!(
        tx.data.len() as u128,
        expected_data_len,
        Ocr3Error::WrongMessageLength
    );

    require!(
        ocr3_config.config_info.config_digest == report_context.byte_words[0],
        Ocr3Error::ConfigDigestMismatch
    );

    // TODO: chain fork check - is this even possible?

    // validate transmitter
    ocr3_config
        .transmitters
        .binary_search_by(|probe| transmitter.to_bytes().cmp(probe))
        .map_err(|_| Ocr3Error::UnauthorizedTransmitter)?;

    if ocr3_config.config_info.is_signature_verification_enabled != 0 {
        require!(
            signatures.len() == (ocr3_config.config_info.f + 1) as usize,
            Ocr3Error::WrongNumberOfSignatures
        );

        verify_signatures(ocr3_config, report.hash(&report_context), signatures)?;
    }

    emit!(Transmitted {
        ocr_plugin_type: ocr3_config.plugin_type,
        config_digest: ocr3_config.config_info.config_digest,
        sequence_number: report_context.sequence_number(),
    });

    Ok(())
}

fn verify_signatures(
    ocr3_config: &Ocr3Config,
    hashed_report: [u8; 32],
    signatures: &[[u8; SIGNATURE_LENGTH]],
) -> Result<()> {
    let mut uniques: u16 = 0;
    assert!(uniques.count_ones() + uniques.count_zeros() >= MAX_SIGNERS as u32); // ensure MAX_SIGNERS fit in the bits of uniques

    for raw_sig in signatures {
        let (recovery_id, signature) = raw_sig.split_first().ok_or(Ocr3Error::InvalidSignature)?;

        let signer = secp256k1_recover(&hashed_report, *recovery_id, signature)
            .map_err(|_| Ocr3Error::InvalidSignature)?;
        // convert to a raw 20 byte Ethereum address
        let address: [u8; 20] = keccak::hash(&signer.0).to_bytes()[12..32]
            .try_into()
            .map_err(|_| Ocr3Error::UnauthorizedSigner)?;
        let index = ocr3_config
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
