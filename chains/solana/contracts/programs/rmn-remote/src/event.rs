use anchor_lang::prelude::*;

use crate::CodeVersion;

#[event]
pub struct OwnershipTransferRequested {
    pub from: Pubkey,
    pub to: Pubkey,
}

#[event]
pub struct OwnershipTransferred {
    pub from: Pubkey,
    pub to: Pubkey,
}

#[event]
pub struct ConfigSet {
    pub default_code_version: CodeVersion,
}
