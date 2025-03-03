use anchor_lang::prelude::*;

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitConfig {
    pub enabled: bool,
    pub capacity: u64,
    pub rate: u64,
}

#[derive(Default, InitSpace, Clone, AnchorDeserialize, AnchorSerialize, PartialEq, Eq)]
pub struct RemoteAddress {
    #[max_len(64)]
    pub address: Vec<u8>,
}