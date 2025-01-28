use anchor_lang::prelude::*;

use crate::DestChainConfig;

// bytes4(keccak256("CCIP EVMExtraArgsV2"));
pub const EVM_EXTRA_ARGS_V2_TAG: u32 = 0x181dcf10;

// bytes4(keccak256("CCIP SVMExtraArgsV1"));
pub const SVM_EXTRA_ARGS_V1_TAG: u32 = 0x1f3b3aba;

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct EVMExtraArgsV2 {
    pub gas_limit: u128,
    pub allow_out_of_order_execution: bool,
}

impl EVMExtraArgsV2 {
    pub fn serialize_with_tag(&self) -> Vec<u8> {
        let mut buffer = EVM_EXTRA_ARGS_V2_TAG.to_be_bytes().to_vec();
        let mut data = self.try_to_vec().unwrap();

        buffer.append(&mut data);
        buffer
    }
    pub fn default_config(cfg: &DestChainConfig) -> EVMExtraArgsV2 {
        EVMExtraArgsV2 {
            gas_limit: cfg.default_tx_gas_limit as u128,
            ..Default::default()
        }
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVMExtraArgsV1 {
    pub compute_units: u32,
    pub account_is_writable_bitmap: u64,
    pub allow_out_of_order_execution: bool,
    pub token_receiver: [u8; 32],
    pub accounts: Vec<[u8; 32]>,
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
