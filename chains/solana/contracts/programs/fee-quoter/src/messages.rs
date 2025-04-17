use anchor_lang::prelude::*;

use crate::{extra_args::GenericExtraArgsV2, DestChainConfig};

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct SVM2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Debug)]
pub struct TokenTransferAdditionalData {
    pub dest_bytes_overhead: u32,
    pub dest_gas_overhead: u32,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct GetFeeResult {
    pub token: Pubkey,
    pub amount: u64,
    pub juels: u128,
    pub token_transfer_additional_data: Vec<TokenTransferAdditionalData>,
    pub processed_extra_args: ProcessedExtraArgs,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug)]
pub struct ProcessedExtraArgs {
    pub bytes: Vec<u8>,
    pub gas_limit: u128,
    pub allow_out_of_order_execution: bool,
    // If unspecified, the message receiver should be used (e.g. with EVM as destination.)
    // This will also be `None` when there is no token transfer.
    pub token_receiver: Option<Vec<u8>>,
}

impl ProcessedExtraArgs {
    pub fn defaults(config: &DestChainConfig) -> Self {
        let args = GenericExtraArgsV2::default_config(config);

        ProcessedExtraArgs {
            bytes: args.serialize_with_tag(),
            gas_limit: args.gas_limit,
            allow_out_of_order_execution: args.allow_out_of_order_execution,
            token_receiver: None,
        }
    }
}
