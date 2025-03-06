use std::fmt::Display;

use anchor_lang::prelude::*;

// Global curse subject, standardized across chains and chain families. If this subject is
// cursed, all lanes starting to or ending in this chain are disabled.
pub fn global_curse_subject() -> Vec<u8> {
    vec![
        0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x01,
    ]
}

pub const SUBJECT_SPACE: usize = 4 + 16;

pub fn curse_from_chain_selector(selector: u64) -> Vec<u8> {
    (selector as u128).to_le_bytes().to_vec()
}

#[derive(Debug, PartialEq, Eq, Clone, Copy, InitSpace, AnchorDeserialize, AnchorSerialize)]
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

    pub proposed_owner: Pubkey,
    pub default_code_version: CodeVersion,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Curses {
    pub version: u8,
    #[max_len(0, 0)]
    pub cursed_subjects: Vec<Vec<u8>>,
}

impl Curses {
    pub fn dynamic_len(&self) -> usize {
        Self::INIT_SPACE + self.cursed_subjects.len() * SUBJECT_SPACE
    }
}
