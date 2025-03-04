use anchor_lang::prelude::*;

pub const NUM_GROUPS: usize = 32;

#[derive(AnchorSerialize, AnchorDeserialize, Clone, InitSpace, Debug)]
pub struct McmSigner {
    pub evm_address: [u8; 20],
    pub index: u8,
    pub group: u8,
}