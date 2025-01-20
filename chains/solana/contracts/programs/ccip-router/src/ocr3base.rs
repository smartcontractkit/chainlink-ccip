use anchor_lang::prelude::*;

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
