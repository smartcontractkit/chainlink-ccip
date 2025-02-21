use anchor_lang::prelude::*;

use crate::{CodeVersion, CurseSubject};

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
    pub local_chain_selector: u64,
}

#[event]
pub struct SubjectCursed {
    pub subject: CurseSubject,
}

#[event]
pub struct SubjectUncursed {
    pub subject: CurseSubject,
}
