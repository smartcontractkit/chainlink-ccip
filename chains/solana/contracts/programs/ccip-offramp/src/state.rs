use std::fmt::Display;

use anchor_lang::prelude::borsh::{BorshDeserialize, BorshSerialize};
use anchor_lang::prelude::*;

use crate::{CcipOfframpError, ConfigOcrPluginType, OcrPluginType};

// zero_copy is used to prevent hitting stack/heap memory limits
#[account(zero_copy)]
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize)]
pub struct Config {
    pub version: u8,
    pub default_code_version: u8,
    _padding0: [u8; 6],
    pub svm_chain_selector: u64,
    pub enable_manual_execution_after: i64, // Expressed as Unix time (i.e. seconds since the Unix epoch).
    _padding1: [u8; 8],

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    _padding2: [u8; 8],
    pub ocr3: [Ocr3Config; 2],
}

#[account(zero_copy)]
#[derive(InitSpace)]
pub struct ReferenceAddresses {
    pub version: u8,
    pub router: Pubkey,
    pub fee_quoter: Pubkey,
    pub offramp_lookup_table: Pubkey,
    pub rmn_remote: Pubkey,
}

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct Ocr3ConfigInfo {
    pub config_digest: [u8; 32], // 32-byte hash of configuration
    pub f: u8,                   // f+1 = number of signatures per report
    pub n: u8,                   // number of signers
    pub is_signature_verification_enabled: u8, // bool -> bytemuck::Pod compliant required for zero_copy
}

// TODO: do we need to verify signers and transmitters are different? (between the two groups)
// signers: pubkey is 20-byte address, secp256k1 curve ECDSA
// transmitters: 32-byte pubkey, ed25519

#[derive(AnchorSerialize, AnchorDeserialize, InitSpace)]
#[zero_copy]
pub struct Ocr3Config {
    pub plugin_type: ConfigOcrPluginType, // plugin identifier for validation (example: ccip:commit = 0, ccip:execute = 1)
    pub config_info: Ocr3ConfigInfo,
    pub signers: [[u8; 20]; 16], // v0.29.0 - anchor IDL does not build with MAX_SIGNERS
    pub transmitters: [[u8; 32]; 16], // v0.29.0 - anchor IDL does not build with MAX_TRANSMITTERS
}

impl Default for Ocr3Config {
    fn default() -> Self {
        Self {
            plugin_type: OcrPluginType::Commit.into(),
            config_info: Default::default(),
            signers: Default::default(),
            transmitters: Default::default(),
        }
    }
}

impl Ocr3Config {
    pub fn new(plugin_type: OcrPluginType) -> Self {
        Self {
            plugin_type: plugin_type.into(),
            ..Default::default()
        }
    }
}

#[account]
#[derive(InitSpace)]
pub struct GlobalState {
    // This holds global variables for the contract that are not manually set by the admin.
    // They are auto-updated by the contract during regular usage of CCIP.
    pub latest_price_sequence_number: u64,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainConfig {
    pub is_enabled: bool, // Flag whether the source chain is enabled or not

    pub is_rmn_verification_disabled: bool, // Currently a placeholder.

    pub lane_code_version: CodeVersion, // The code version of the lane, which may override the global default code version

    pub on_ramp: OnRampAddress, // OnRamp addresses supported from the source chain
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct OnRampAddress {
    bytes: [u8; 64],
    len: u32,
}

impl OnRampAddress {
    pub fn bytes(&self) -> &[u8] {
        &self.bytes[0..self.len as usize]
    }

    pub fn is_empty(&self) -> bool {
        self.len == 0
    }

    pub fn is_zero(&self) -> bool {
        self.bytes().iter().all(|b| *b == 0)
    }

    pub const EMPTY: Self = Self {
        bytes: [0u8; 64],
        len: 0,
    };
}

impl TryInto<OnRampAddress> for Vec<u8> {
    type Error = CcipOfframpError;

    fn try_into(self) -> std::result::Result<OnRampAddress, Self::Error> {
        let mut address = OnRampAddress {
            bytes: [0u8; 64],
            len: 0,
        };
        if self.is_empty() || self.len() > address.bytes.len() {
            return Err(CcipOfframpError::InvalidOnrampAddress);
        }

        address.len = self.len() as u32;
        address.bytes[0..address.len as usize].copy_from_slice(&self);

        Ok(address)
    }
}

impl<const N: usize> From<[u8; N]> for OnRampAddress {
    fn from(val: [u8; N]) -> Self {
        assert!(N <= 64 && N > 0);
        let mut address = Self {
            bytes: [0u8; 64],
            len: N as u32,
        };

        address.bytes[0..N].copy_from_slice(&val);
        address
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainState {
    pub min_seq_nr: u64, // The min sequence number expected for future messages
}

#[account]
#[derive(InitSpace, Debug)]
pub struct SourceChain {
    // Config for Any2SVM
    pub version: u8,
    pub chain_selector: u64,       // Chain selector used for the seed
    pub state: SourceChainState,   // values that are updated automatically
    pub config: SourceChainConfig, // values configured by an admin
}

#[account]
#[derive(InitSpace)]
pub struct CommitReport {
    pub version: u8,
    pub chain_selector: u64,
    pub merkle_root: [u8; 32],
    pub timestamp: i64, // Expressed as Unix time (i.e. seconds since the Unix epoch).
    pub min_msg_nr: u64,
    pub max_msg_nr: u64, // TODO: Change this to [u128; 2] when supporting commit reports with 256 messages
    pub execution_states: u128,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Debug, PartialEq)]
// used in the commit report execution_states field
pub enum MessageExecutionState {
    Untouched = 0,
    InProgress = 1, // Not used in SVM, but used in EVM
    Success = 2,
    Failure = 3, // Not used in SVM, but used in EVM
}

impl TryFrom<u128> for MessageExecutionState {
    type Error = &'static str;

    fn try_from(value: u128) -> std::result::Result<MessageExecutionState, &'static str> {
        match value {
            0 => Ok(MessageExecutionState::Untouched),
            1 => Ok(MessageExecutionState::InProgress),
            2 => Ok(MessageExecutionState::Success),
            3 => Ok(MessageExecutionState::Failure),
            _ => Err("Invalid ExecutionState"),
        }
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy, InitSpace, BorshSerialize, BorshDeserialize)]
#[repr(u8)]
pub enum CodeVersion {
    Default = 0,
    V1,
}

impl Display for CodeVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            CodeVersion::Default => write!(f, "Default"),
            CodeVersion::V1 => write!(f, "V1"),
        }
    }
}

impl From<CodeVersion> for u8 {
    fn from(value: CodeVersion) -> u8 {
        value as u8
    }
}

impl TryFrom<u8> for CodeVersion {
    type Error = CcipOfframpError;

    fn try_from(value: u8) -> std::result::Result<Self, Self::Error> {
        match value {
            0 => Ok(CodeVersion::Default),
            1 => Ok(CodeVersion::V1),
            _ => Err(CcipOfframpError::InvalidCodeVersion),
        }
    }
}

#[cfg(test)]
mod tests {

    use super::*;
    use std::convert::TryFrom;

    #[test]
    fn test_execution_state_try_from() {
        assert_eq!(
            MessageExecutionState::try_from(0).unwrap(),
            MessageExecutionState::Untouched
        );
        assert_eq!(
            MessageExecutionState::try_from(1).unwrap(),
            MessageExecutionState::InProgress
        );
        assert_eq!(
            MessageExecutionState::try_from(2).unwrap(),
            MessageExecutionState::Success
        );
        assert_eq!(
            MessageExecutionState::try_from(3).unwrap(),
            MessageExecutionState::Failure
        );
        assert!(MessageExecutionState::try_from(4).is_err());
    }

    #[test]
    fn test_code_version_try_from_u8() {
        assert_eq!(CodeVersion::try_from(0).unwrap(), CodeVersion::Default);
        assert_eq!(CodeVersion::try_from(1).unwrap(), CodeVersion::V1);
        assert!(CodeVersion::try_from(2).is_err()); // this should be updated if new code versions are added
    }

    #[test]
    fn test_u8_from_code_version() {
        assert_eq!(u8::from(CodeVersion::Default), 0);
        assert_eq!(u8::from(CodeVersion::V1), 1);
    }
}
