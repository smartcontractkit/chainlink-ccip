use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, Debug, InitSpace)]
pub enum CodeVersion {
    Default = 0,
    V1,
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize, Debug)]
pub struct DestChainConfig {
    pub is_enabled: bool,
    pub lane_code_version: CodeVersion,
    pub max_number_of_tokens_per_msg: u16,
    pub max_data_bytes: u32,
    pub max_per_msg_gas_limit: u32,
    pub dest_gas_overhead: u32,
    pub dest_gas_per_payload_byte_base: u32,
    pub dest_gas_per_payload_byte_high: u32,
    pub dest_gas_per_payload_byte_threshold: u32,
    pub dest_data_availability_overhead_gas: u32,
    pub dest_gas_per_data_availability_byte: u16,
    pub dest_data_availability_multiplier_bps: u16,
    pub default_token_fee_usdcents: u16,
    pub default_token_dest_gas_overhead: u32,
    pub default_tx_gas_limit: u32,
    pub gas_multiplier_wei_per_eth: u64,
    pub network_fee_usdcents: u32,
    pub gas_price_staleness_threshold: u32,
    pub enforce_out_of_order: bool,
    pub chain_family_selector: [u8; 4],
}

#[derive(InitSpace, Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct TokenTransferFeeConfig {
    pub min_fee_usdcents: u32,
    pub max_fee_usdcents: u32,
    pub deci_bps: u16,
    pub dest_gas_overhead: u32,
    pub dest_bytes_overhead: u32,
    pub is_enabled: bool,
}