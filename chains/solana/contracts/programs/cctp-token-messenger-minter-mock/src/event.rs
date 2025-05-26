use anchor_lang::prelude::*;

#[event]
pub struct DepositForBurn {
    pub nonce: u64,
    pub burn_token: Pubkey,
    pub amount: u64,
    pub depositor: Pubkey,
    pub mint_recipient: Pubkey,
    pub destination_domain: u32,
    pub destination_token_messenger: Pubkey,
    pub destination_caller: Pubkey,
}

#[event]
pub struct MintAndWithdraw {
    pub mint_recipient: Pubkey,
    pub amount: u64,
    pub mint_token: Pubkey,
}

#[event]
pub struct RemoteTokenMessengerAdded {
    pub domain: u32,
    pub token_messenger: Pubkey,
}
#[event]
pub struct SetTokenController {
    pub token_controller: Pubkey,
}

#[event]
pub struct PauserChanged {
    pub new_address: Pubkey,
}

#[event]
pub struct SetBurnLimitPerMessage {
    pub token: Pubkey,
    pub burn_limit_per_message: u64,
}

#[event]
pub struct LocalTokenAdded {
    pub custody: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct LocalTokenRemoved {
    pub custody: Pubkey,
    pub mint: Pubkey,
}

#[event]
pub struct TokenPairLinked {
    pub local_token: Pubkey,
    pub remote_domain: u32,
    pub remote_token: Pubkey,
}

#[event]
pub struct Pause {}

#[event]
pub struct Unpause {}
