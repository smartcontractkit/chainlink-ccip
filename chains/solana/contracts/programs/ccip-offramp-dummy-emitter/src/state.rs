use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize)]
pub struct PriceUpdates {
    pub token_prices_usd: Vec<u128>,
    pub gas_prices_usd: Vec<u128>,
    pub dest_token_prices: Vec<u128>,
    pub dest_native_prices: Vec<u128>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct SourceChainConfig {
    pub is_enabled: bool,
    pub lane_code_version: CodeVersion,
    pub on_ramp: [[u8; 64]; 2],
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy)]
pub enum CodeVersion {
    Default = 0,
    V1,
}

#[derive(AnchorSerialize, AnchorDeserialize)]
pub enum MessageExecutionState {
    Untouched = 0,
    InProgress = 1,
    Success = 2,
    Failure = 3,
}