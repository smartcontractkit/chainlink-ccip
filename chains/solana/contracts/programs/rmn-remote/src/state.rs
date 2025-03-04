use std::fmt::Display;

use anchor_lang::prelude::*;

/// Abstract curse subject.
///
/// In particular, a curse subject can be constructed from a chain
/// selector to signify that any lane involving that chain as `destination` or `source` is
/// cursed.
///
/// The above is not exhaustive: there may be other ways to define subjects.
#[derive(Debug, PartialEq, Eq, Clone, Copy, InitSpace, AnchorDeserialize, AnchorSerialize)]
pub struct CurseSubject {
    pub value: [u8; 16],
}

impl CurseSubject {
    // Global curse subject, standardized across chains and chain families. If this subject is
    // cursed, all lanes starting to or ending in this chain are disabled.
    pub const GLOBAL: Self = {
        Self {
            value: [
                0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x00, 0x01,
            ],
        }
    };

    pub const fn from_chain_selector(selector: u64) -> Self {
        Self {
            value: (selector as u128).to_le_bytes(),
        }
    }

    pub const fn from_bytes(bytes: [u8; 16]) -> Self {
        Self { value: bytes }
    }
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
    #[max_len(0)]
    pub cursed_subjects: Vec<CurseSubject>,
}

impl Curses {
    pub fn dynamic_len(&self) -> usize {
        Self::INIT_SPACE + self.cursed_subjects.len() * CurseSubject::INIT_SPACE
    }
}
