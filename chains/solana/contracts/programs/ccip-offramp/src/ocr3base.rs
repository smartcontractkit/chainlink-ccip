use anchor_lang::prelude::*;

use crate::OcrPluginType;

#[constant]
pub const MAX_ORACLES: usize = 16; // can set a maximum of 16 transmitters + 16 signers simultaneously in a single set config tx

#[event]
pub struct ConfigSet {
    pub ocr_plugin_type: OcrPluginType,
    pub config_digest: [u8; 32],
    pub signers: Vec<[u8; 20]>,
    pub transmitters: Vec<Pubkey>,
    pub f: u8,
}

#[event]
pub struct Transmitted {
    pub ocr_plugin_type: OcrPluginType,
    pub config_digest: [u8; 32],
    pub sequence_number: u64,
}
