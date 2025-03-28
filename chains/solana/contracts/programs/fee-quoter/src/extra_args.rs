use anchor_lang::prelude::*;

use crate::DestChainConfig;

// NOTE: this file contains ExtraArgs for SVM -> remote chains
// these are onramp specific extraArgs for ccipSend only

/// Tag to indicate a gas limit (or dest chain equivalent processing units) and Out Of Order Execution. This tag is
/// available for multiple chain families. If there is no chain family specific tag, this is the default available
/// for a chain.
///
/// Note: not available for Solana VM based chains.
///
/// Fully compatible with the previously existing EVMExtraArgsV2, hence the string used to generate it.
/// bytes4(keccak256("CCIP EVMExtraArgsV2"));
pub const GENERIC_EXTRA_ARGS_V2_TAG: u32 = 0x181dcf10;

/// bytes4(keccak256("CCIP SVMExtraArgsV1"));
pub const SVM_EXTRA_ARGS_V1_TAG: u32 = 0x1f3b3aba;

/// The maximum number of accounts that can be passed in SVMExtraArgs.
pub const SVM_EXTRA_ARGS_MAX_ACCOUNTS: usize = 64;

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct GenericExtraArgsV2 {
    pub gas_limit: u128,                    // message gas limit
    pub allow_out_of_order_execution: bool, // user configurable OOO execution that must match with the DestChainConfig.EnforceOutOfOrderExecution
}

impl GenericExtraArgsV2 {
    pub fn serialize_with_tag(&self) -> Vec<u8> {
        let mut buffer = GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes().to_vec();
        let mut data = self.try_to_vec().unwrap();

        buffer.append(&mut data);
        buffer
    }
    pub fn default_config(cfg: &DestChainConfig) -> GenericExtraArgsV2 {
        GenericExtraArgsV2 {
            gas_limit: cfg.default_tx_gas_limit as u128,
            ..Default::default()
        }
    }
}

// Extra Args of message when SVM -> SVM
#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVMExtraArgsV1 {
    pub compute_units: u32, // compute units are used by the offchain to add computeUnitLimits to the execute stage
    pub account_is_writable_bitmap: u64,
    pub allow_out_of_order_execution: bool, // SVM chains require OOO, but users are required to explicitly state OOO is allowed in extraArgs
    pub token_receiver: [u8; 32],           // cannot be 0-address if tokens are sent in message
    pub accounts: Vec<[u8; 32]>,            // account list for messaging on remote SVM chain
}

impl SVMExtraArgsV1 {
    pub fn serialize_with_tag(&self) -> Vec<u8> {
        let mut buffer = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();
        let mut data = self.try_to_vec().unwrap();

        buffer.append(&mut data);
        buffer
    }
    pub fn default_config(cfg: &DestChainConfig) -> SVMExtraArgsV1 {
        SVMExtraArgsV1 {
            compute_units: cfg.default_tx_gas_limit,
            ..Default::default()
        }
    }
}
