use anchor_lang::prelude::*;

use crate::context::{
    AcceptOwnership, AddSourceChain, CommitReportContext, ExecuteReportContext,
    PriceOnlyCommitReportContext, SetOcrConfig, TransferOwnership, UpdateConfig,
    UpdateReferenceAddresses, UpdateSourceChain,
};
use crate::state::{CodeVersion, Ocr3ConfigInfo, SourceChainConfig};

pub trait Commit {
    fn commit<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, CommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()>;

    fn commit_price_only<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, PriceOnlyCommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()>;
}

pub trait Execute {
    fn execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        report_context_byte_words: [[u8; 32]; 2],
        token_indexes: &[u8],
    ) -> Result<()>;

    fn manually_execute<'info>(
        &self,
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        token_indexes: &[u8],
    ) -> Result<()>;
}

pub trait Admin {
    fn transfer_ownership(
        &self,
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()>;

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()>;

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()>;

    fn update_reference_addresses(
        &self,
        ctx: Context<UpdateReferenceAddresses>,
        router: Pubkey,
        fee_quoter: Pubkey,
        offramp_lookup_table: Pubkey,
    ) -> Result<()>;

    fn add_source_chain(
        &self,
        ctx: Context<AddSourceChain>,
        new_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()>;

    fn disable_source_chain_selector(
        &self,
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
    ) -> Result<()>;

    fn update_source_chain_config(
        &self,
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()>;

    fn update_svm_chain_selector(
        &self,
        ctx: Context<UpdateConfig>,
        new_chain_selector: u64,
    ) -> Result<()>;

    fn update_enable_manual_execution_after(
        &self,
        ctx: Context<UpdateConfig>,
        new_enable_manual_execution_after: i64,
    ) -> Result<()>;

    fn set_ocr_config(
        &self,
        ctx: Context<SetOcrConfig>,
        plugin_type: u8, // OcrPluginType, u8 used because anchor tests did not work with an enum
        config_info: Ocr3ConfigInfo,
        signers: Vec<[u8; 20]>,
        transmitters: Vec<Pubkey>,
    ) -> Result<()>;
}
