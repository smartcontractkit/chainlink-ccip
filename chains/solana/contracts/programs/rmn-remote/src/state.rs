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

    pub proposed_owner: Pubkey,
    pub default_code_version: CodeVersion,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Curses {
    #[max_len(0)]
    pub cursed_subjects: Vec<CurseSubject>,

    // Set during initialization of the PDA. Defines the subject that represents
    // this entire chain being cursed, if present on the vector above.
    global_subject: CurseSubject,
}

impl Curses {
    pub fn dynamic_len(&self) -> usize {
        Self::INIT_SPACE + self.cursed_subjects.len() * CurseSubject::INIT_SPACE
    }

    /// sets the chain selector of the current chain, which will be saved as a curse
    /// subject that represents the entire chain being cursed.
    pub fn set_local_chain_selector(&mut self, local_chain_selector: u64) {
        self.global_subject = CurseSubject::from_chain_selector(local_chain_selector);
    }

    pub fn is_subject_cursed(&self, subject: CurseSubject) -> bool {
        self.cursed_subjects.contains(&subject)
    }

    pub fn is_chain_globally_cursed(&self) -> bool {
        self.cursed_subjects.contains(&self.global_subject)
    }
}
