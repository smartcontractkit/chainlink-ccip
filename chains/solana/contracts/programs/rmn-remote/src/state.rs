use std::fmt::Display;

use anchor_lang::prelude::*;
use borsh::{BorshDeserialize, BorshSerialize};

#[derive(Debug, PartialEq, Eq, Clone, Copy, InitSpace, BorshSerialize, BorshDeserialize)]
pub struct CurseSubject {
    value: [u8; 16],
}

impl CurseSubject {
    pub const fn from_chain_selector(selector: u64) -> Self {
        Self {
            value: (selector as u128).to_le_bytes(),
        }
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy, InitSpace, BorshSerialize, BorshDeserialize)]
#[repr(u8)]
pub enum CodeVersion {
    Default = 0,
    V1,
}

impl Display for CodeVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            CodeVersion::Default => write!(f, "Default"),
            CodeVersion::V1 => write!(f, "V1"),
        }
    }
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Config {
    pub version: u8,
    pub owner: Pubkey,
    pub local_chain_selector: u64,

    pub proposed_owner: Pubkey,
    pub default_code_version: CodeVersion,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Cursed {
    #[max_len(0)]
    pub subjects: Vec<CurseSubject>,
}

impl Cursed {
    pub fn dynamic_len(&self) -> usize {
        Self::INIT_SPACE + self.subjects.len() * CurseSubject::INIT_SPACE
    }
}
