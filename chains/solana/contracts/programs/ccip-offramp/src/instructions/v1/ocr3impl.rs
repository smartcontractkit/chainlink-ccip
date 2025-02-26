use anchor_lang::AnchorSerialize;

use crate::context::CommitInput;
use crate::messages::ExecutionReportSingleChain;

use super::ocr3base::{Ocr3Report, ReportContext};

pub(super) struct Ocr3ReportForCommit<'a>(pub &'a CommitInput);

impl Ocr3Report for Ocr3ReportForCommit<'_> {
    fn hash(&self, ctx: &ReportContext) -> [u8; 32] {
        use anchor_lang::solana_program::keccak;
        let mut buffer: Vec<u8> = Vec::new();
        self.0.serialize(&mut buffer).unwrap();
        let report_len = self.len() as u16; // u16 > max tx size, u8 may have overflow
        keccak::hashv(&[&report_len.to_le_bytes(), &buffer, &ctx.as_bytes()]).to_bytes()
    }

    fn len(&self) -> usize {
        4 + (32 + 28) * self.0.price_updates.token_price_updates.len() + // token_price_updates
      4 + (8 + 28) * self.0.price_updates.gas_price_updates.len() + // gas_price_updates
      4 + (32 + 32) * self.0.rmn_signatures.len() + // rmn signatures
      1 + self.0.merkle_root.as_ref().map(|r| r.len()).unwrap_or(0)
        // + 4 + 65 * self.rmn_signatures.len()
    }
}

pub(super) struct Ocr3ReportForExecutionReportSingleChain<'a>(pub &'a ExecutionReportSingleChain);

impl Ocr3Report for Ocr3ReportForExecutionReportSingleChain<'_> {
    fn hash(&self, _: &ReportContext) -> [u8; 32] {
        [0; 32] // not needed, this report is not hashed for signing
    }
    fn len(&self) -> usize {
        let offchain_token_data_len = self
            .0
            .offchain_token_data
            .iter()
            .fold(0, |acc, e| acc + 4 + e.len());

        8 // source chain selector
      + self.0.message.len() // ccip message
      + 4 + offchain_token_data_len// offchain_token_data
      + 4 + self.0.proofs.len() * 32 // count + proofs
      + 4 + self.0.message.token_amounts.len() // token_indexes (not part of report but part of tx size validation)
    }
}
