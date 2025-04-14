use anchor_lang::prelude::*;
use anchor_lang::solana_program::sysvar;
use anchor_lang::solana_program::{keccak, secp256k1_recover::*};

use crate::ocr3base::{ConfigSet, Transmitted, MAX_ORACLES};
use crate::state::{Ocr3Config, Ocr3ConfigInfo};
use crate::CcipOfframpError;
use crate::OcrPluginType;

pub const MAX_SIGNERS: usize = MAX_ORACLES;
pub const MAX_TRANSMITTERS: usize = MAX_ORACLES;

pub const TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES: u128 = 8 // anchor discriminator
    + 2 * 32 // report context
    + 4; // length component of serialized report
pub const TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES: u128 = 4 // u32 length of rs vec
    + 4 // u32 length of ss vec
    + 32; // length of rawVs

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub(super) struct ReportContext {
    // byte_words consists of:
    // [0]: ConfigDigest
    // [1]: 24 byte padding, 8 byte sequence number
    byte_words: [[u8; 32]; 2], // private, define methods to use it
}

// Signature struct used to reduce total number of input parameters to function
pub struct Signatures {
    pub rs: Vec<[u8; 32]>,
    pub ss: Vec<[u8; 32]>,
    pub raw_vs: [u8; 32],
}

impl ReportContext {
    pub fn sequence_number(&self) -> u64 {
        let sequence_bytes: [u8; 8] = self.byte_words[1][24..].try_into().unwrap();
        u64::from_be_bytes(sequence_bytes)
    }

    pub fn from_byte_words(byte_words: [[u8; 32]; 2]) -> Self {
        Self { byte_words }
    }

    pub fn as_bytes(&self) -> [u8; 32 * 2] {
        self.byte_words.concat().as_slice().try_into().unwrap()
    }
}

pub fn ocr3_set(
    ocr3_config: &mut Ocr3Config,
    plugin_type: OcrPluginType,
    cfg: Ocr3ConfigInfo,
    signers: Vec<[u8; 20]>,
    transmitters: Vec<Pubkey>,
) -> Result<()> {
    require!(
        plugin_type == ocr3_config.plugin_type.try_into()?,
        CcipOfframpError::Ocr3InvalidPluginType
    );
    require!(
        cfg.f != 0,
        CcipOfframpError::Ocr3InvalidConfigFMustBePositive
    );

    // If F is 0, then the config is not yet set
    if ocr3_config.config_info.f == 0 {
        ocr3_config.config_info.is_signature_verification_enabled =
            cfg.is_signature_verification_enabled;
    } else {
        require!(
            ocr3_config.config_info.is_signature_verification_enabled
                == cfg.is_signature_verification_enabled,
            CcipOfframpError::Ocr3StaticConfigCannotBeChanged
        )
    };

    require!(
        !transmitters.is_empty(),
        CcipOfframpError::Ocr3InvalidConfigNoTransmitters
    );

    require!(
        transmitters.len() <= MAX_TRANSMITTERS,
        CcipOfframpError::Ocr3InvalidConfigTooManyTransmitters
    );

    clear_transmitters(ocr3_config);

    if cfg.is_signature_verification_enabled != 0 {
        clear_signers(ocr3_config);
        require!(
            signers.len() <= MAX_SIGNERS,
            CcipOfframpError::Ocr3InvalidConfigTooManySigners
        );
        require!(
            signers.len() > 3 * cfg.f as usize,
            CcipOfframpError::Ocr3InvalidConfigFIsTooHigh
        );

        // NOTE: Transmitters cannot exceed signers. Transmitters do not have to be >= 3F + 1 because they can
        // match >= 3fChain + 1, where fChain <= F. fChain is not represented in MultiOCR3Base - so we skip this check.
        require_gte!(
            signers.len(),
            transmitters.len(),
            CcipOfframpError::Ocr3InvalidConfigTooManyTransmitters
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
        ocr_plugin_type: ocr3_config.plugin_type.try_into()?,
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
    plugin_type: OcrPluginType,
    report_context: ReportContext,
    report: &R,
    signatures: Signatures,
) -> Result<()> {
    require!(
        plugin_type == ocr3_config.plugin_type.try_into()?,
        CcipOfframpError::Ocr3InvalidPluginType
    );

    // validate raw message length
    let mut expected_data_len: u128 = TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT_NO_SIGNATURES
        .saturating_add(report.len() as u128);
    if ocr3_config.config_info.is_signature_verification_enabled != 0 {
        expected_data_len = expected_data_len
            .saturating_add(TRANSMIT_MSGDATA_EXTRA_CONSTANT_LENGTH_COMPONENT_FOR_SIGNATURES);
        expected_data_len = expected_data_len
            .saturating_add((signatures.rs.len() * SECP256K1_SIGNATURE_LENGTH) as u128);
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
        CcipOfframpError::Ocr3WrongMessageLength
    );

    require!(
        ocr3_config.config_info.config_digest == report_context.byte_words[0],
        CcipOfframpError::Ocr3ConfigDigestMismatch
    );

    // TODO: chain fork check - is this even possible?

    // validate transmitter
    ocr3_config
        .transmitters
        .binary_search_by(|probe| transmitter.to_bytes().cmp(probe))
        .map_err(|_| CcipOfframpError::Ocr3UnauthorizedTransmitter)?;

    if ocr3_config.config_info.is_signature_verification_enabled != 0 {
        require_eq!(
            signatures.rs.len(),
            (ocr3_config.config_info.f + 1) as usize,
            CcipOfframpError::Ocr3WrongNumberOfSignatures
        );
        require_eq!(
            signatures.rs.len(),
            signatures.ss.len(),
            CcipOfframpError::Ocr3SignaturesOutOfRegistration
        );

        verify_signatures(ocr3_config, report.hash(&report_context), signatures)?;
    }

    emit!(Transmitted {
        ocr_plugin_type: ocr3_config.plugin_type.try_into()?,
        config_digest: ocr3_config.config_info.config_digest,
        sequence_number: report_context.sequence_number(),
    });

    Ok(())
}

fn verify_signatures(
    ocr3_config: &Ocr3Config,
    hashed_report: [u8; 32],
    signatures: Signatures,
) -> Result<()> {
    let mut uniques: u16 = 0;
    // verify that uniques has no bits set (all zeros) at initialization
    assert_eq!(uniques.count_ones(), 0);
    // ensure the u16 type has enough bits to represent all MAX_SIGNERS
    assert!(uniques.count_zeros() >= MAX_SIGNERS as u32);

    for (i, _) in signatures.rs.iter().enumerate() {
        let signer = secp256k1_recover(
            &hashed_report,
            signatures.raw_vs[i],
            [signatures.rs[i], signatures.ss[i]].concat().as_ref(),
        )
        .map_err(|_| CcipOfframpError::Ocr3InvalidSignature)?;
        // convert to a raw 20 byte Ethereum address
        let address: [u8; 20] = keccak::hash(&signer.0).to_bytes()[12..32]
            .try_into()
            .map_err(|_| CcipOfframpError::Ocr3UnauthorizedSigner)?;
        let index = ocr3_config
            .signers
            .binary_search_by(|probe| address.cmp(probe))
            .map_err(|_| CcipOfframpError::Ocr3UnauthorizedSigner)?;

        uniques |= 1 << index;
    }

    require!(
        uniques.count_ones() as usize == signatures.rs.len(),
        CcipOfframpError::Ocr3NonUniqueSignatures
    );

    Ok(())
}

// assign_oracles (generic) takes a vector of byte arrays sorts (for binary searching), validates, and writes it to `location`
fn assign_oracles<const A: usize>(oracles: &mut [[u8; A]], location: &mut [[u8; A]]) -> Result<()> {
    // sort + verify not duplicated
    oracles.sort_unstable_by(|a, b| b.cmp(a)); // reverse sort to ensure appended 0-values in correct
    let duplicate = oracles.windows(2).any(|pair| pair[0] == pair[1]);
    require!(
        !duplicate,
        CcipOfframpError::Ocr3InvalidConfigRepeatedOracle
    );

    for (i, o) in oracles.iter().enumerate() {
        require!(
            *o != [0; A],
            CcipOfframpError::Ocr3OracleCannotBeZeroAddress
        );
        location[i] = *o;
    }
    Ok(())
}
