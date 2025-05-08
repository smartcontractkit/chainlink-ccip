use anchor_lang::prelude::*;

// Identifies a buffer. Has no specific meaning: it's only used
// by the caller to track which buffer is which in case of uploading several.
#[derive(Clone, Debug, AnchorDeserialize, AnchorSerialize, Eq, PartialEq)]
pub struct BufferId {
    pub bytes: [u8; 32],
}

#[account]
#[derive(Debug, InitSpace)]
pub struct BufferedReport {
    #[max_len(0)]
    pub raw_report_data: Vec<u8>,
}
