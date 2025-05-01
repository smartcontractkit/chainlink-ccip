use anchor_lang::prelude::*;

#[account]
#[derive(Debug, InitSpace)]
pub struct BufferedReport {
    #[max_len(0)]
    pub raw_report_data: Vec<u8>,
}
