use anchor_lang::prelude::*;

#[event]
pub struct FeeTokenAdded {
    pub fee_token: Pubkey,
    pub enabled: bool,
}

#[event]
pub struct FeeTokenEnabled {
    pub fee_token: Pubkey,
}

#[event]
pub struct FeeTokenDisabled {
    pub fee_token: Pubkey,
}

#[event]
pub struct FeeTokenRemoved {
    pub fee_token: Pubkey,
}
