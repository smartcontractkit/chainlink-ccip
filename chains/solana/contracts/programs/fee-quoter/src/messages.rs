use anchor_lang::prelude::*;

pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct SVM2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: ExtraArgsInput,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct ExtraArgsInput {
    pub gas_limit: Option<u128>,
    pub allow_out_of_order_execution: Option<bool>,
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
    pub juels: u64,
    pub token_transfer_additional_data: Vec<TokenTransferAdditionalData>,
}
