use anchor_lang::error_code;
use anchor_lang::prelude::*;

declare_id!("offRPDpDxT5MGFNmMh99QKTZfPWTkqYUrStEriAS1H5");

mod event;
mod messages;
mod ocr3base;

mod context;
use crate::context::*;

mod instructions;
use crate::instructions::v1;

mod state;
use state::*;

#[program]
pub mod ccip_offramp {
    use super::*;

    //////////////////////////
    /// Initialization Flow //
    //////////////////////////

    /// Initializes the CCIP Offramp.
    ///
    /// The initialization of the Offramp is responsibility of Admin, nothing more than calling this method should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization.
    /// * `svm_chain_selector` - The chain selector for SVM.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    pub fn initialize(
        ctx: Context<Initialize>,
        svm_chain_selector: u64,
        enable_execution_after: i64,
    ) -> Result<()> {
        {
            let mut config = ctx.accounts.config.load_init()?;
            require!(config.version == 0, CcipOfframpError::InvalidInputs); // assert uninitialized state - AccountLoader doesn't work with constraint
            config.version = 1;
            config.svm_chain_selector = svm_chain_selector;
            config.enable_manual_execution_after = enable_execution_after;
            config.owner = ctx.accounts.authority.key();
            config.ocr3 = [
                Ocr3Config::new(OcrPluginType::Commit as u8),
                Ocr3Config::new(OcrPluginType::Execution as u8),
            ];
        }

        ctx.accounts
            .reference_addresses
            .set_inner(ReferenceAddresses {
                version: 1,
                router: ctx.accounts.router.key(),
                fee_quoter: ctx.accounts.fee_quoter.key(),
                offramp_lookup_table: ctx.accounts.offramp_lookup_table.key(),
            });

        ctx.accounts.state.latest_price_sequence_number = 0;

        Ok(())
    }

    /// Transfers the ownership of the router to a new proposed owner.
    ///
    /// Shared func signature with other programs
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the transfer.
    /// * `proposed_owner` - The public key of the new proposed owner.
    pub fn transfer_ownership(
        ctx: Context<TransferOwnership>,
        proposed_owner: Pubkey,
    ) -> Result<()> {
        v1::admin::transfer_ownership(ctx, proposed_owner)
    }

    /// Accepts the ownership of the router by the proposed owner.
    ///
    /// Shared func signature with other programs
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for accepting ownership.
    /// The new owner must be a signer of the transaction.
    pub fn accept_ownership(ctx: Context<AcceptOwnership>) -> Result<()> {
        v1::admin::accept_ownership(ctx)
    }

    /////////////
    /// Config //
    /////////////

    /// Adds a new source chain selector with its config to the offramp.
    ///
    /// The Admin needs to add any new chain supported.
    /// When adding a new chain, the Admin needs to specify if it's enabled or not.
    ///
    /// # Arguments
    pub fn add_source_chain(
        ctx: Context<AddSourceChain>,
        new_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()> {
        v1::admin::add_source_chain(ctx, new_chain_selector, source_chain_config)
    }

    /// Disables the source chain selector.
    ///
    /// The Admin is the only one able to disable the chain selector as source. This method is thought of as an emergency kill-switch.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for disabling the chain selector.
    /// * `source_chain_selector` - The source chain selector to be disabled.
    pub fn disable_source_chain_selector(
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::disable_source_chain_selector(ctx, source_chain_selector)
    }

    /// Updates the configuration of the source chain selector.
    ///
    /// The Admin is the only one able to update the source chain config.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the chain selector.
    /// * `source_chain_selector` - The source chain selector to be updated.
    /// * `source_chain_config` - The new configuration for the source chain.
    pub fn update_source_chain_config(
        ctx: Context<UpdateSourceChain>,
        source_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()> {
        v1::admin::update_source_chain_config(ctx, source_chain_selector, source_chain_config)
    }

    /// Updates the SVM chain selector in the offramp configuration.
    ///
    /// This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_chain_selector` - The new chain selector for SVM.
    pub fn update_svm_chain_selector(
        ctx: Context<UpdateConfig>,
        new_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::update_svm_chain_selector(ctx, new_chain_selector)
    }

    /// Updates the minimum amount of time required between a message being committed and when it can be manually executed.
    ///
    /// This is part of the OffRamp Configuration for SVM.
    /// The Admin is the only one able to update this config.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_enable_manual_execution_after` - The new minimum amount of time required.
    pub fn update_enable_manual_execution_after(
        ctx: Context<UpdateConfig>,
        new_enable_manual_execution_after: i64,
    ) -> Result<()> {
        v1::admin::update_enable_manual_execution_after(ctx, new_enable_manual_execution_after)
    }

    /// Sets the OCR configuration.
    /// Only CCIP Admin can set the OCR configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for setting the OCR configuration.
    /// * `plugin_type` - The type of OCR plugin [0: Commit, 1: Execution].
    /// * `config_info` - The OCR configuration information.
    /// * `signers` - The list of signers.
    /// * `transmitters` - The list of transmitters.
    pub fn set_ocr_config(
        ctx: Context<SetOcrConfig>,
        plugin_type: u8, // OcrPluginType, u8 used because anchor tests did not work with an enum
        config_info: Ocr3ConfigInfo,
        signers: Vec<[u8; 20]>,
        transmitters: Vec<Pubkey>,
    ) -> Result<()> {
        v1::admin::set_ocr_config(ctx, plugin_type, config_info, signers, transmitters)
    }

    ////////////////////
    /// Off Ramp Flow //
    ////////////////////

    /// Commits a report to the router.
    ///
    /// The method name needs to be commit with Anchor encoding.
    ///
    /// This function is called by the OffChain when committing one Report to the SVM Router.
    /// In this Flow only one report is sent, the Commit Report. This is different as EVM does,
    /// this is because here all the chain state is stored in one account per Merkle Tree Root.
    /// So, to avoid having to send a dynamic size array of accounts, in this message only one Commit Report Account is sent.
    /// This message validates the signatures of the report and stores the Merkle Root in the Commit Report Account.
    /// The Report must contain an interval of messages, and the min of them must be the next sequence number expected.
    /// The max size of the interval is 64.
    /// This message emits two events: CommitReportAccepted and Transmitted.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the commit.
    /// * `report_context_byte_words` - consists of:
    ///     * report_context_byte_words[0]: ConfigDigest
    ///     * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
    /// * `raw_report` - The serialized commit input report, single merkle root with RMN signatures and price updates
    /// * `rs` - slice of R components of signatures
    /// * `ss` - slice of S components of signatures
    /// * `raw_vs` - array of V components of signatures
    pub fn commit<'info>(
        ctx: Context<'_, '_, 'info, 'info, CommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()> {
        v1::commit::commit(ctx, report_context_byte_words, raw_report, rs, ss, raw_vs)
    }

    /// Executes a message on the destination chain.
    ///
    /// The method name needs to be execute with Anchor encoding.
    ///
    /// This function is called by the OffChain when executing one Report to the SVM Router.
    /// In this Flow only one message is sent, the Execution Report. This is different as EVM does,
    /// this is because there is no try/catch mechanism to allow batch execution.
    /// This message validates that the Merkle Tree Proof of the given message is correct and is stored in the Commit Report Account.
    /// The message must be untouched to be executed.
    /// This message emits the event ExecutionStateChanged with the new state of the message.
    /// Finally, executes the CPI instruction to the receiver program in the ccip_receive message.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the execute.
    /// * `raw_execution_report` - the serialized execution report containing only one message and proofs
    /// * `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)
    /// *  consists of:
    ///     * report_context_byte_words[0]: ConfigDigest
    ///     * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
    pub fn execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        report_context_byte_words: [[u8; 32]; 2],
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::execute::execute(
            ctx,
            raw_execution_report,
            report_context_byte_words,
            &token_indexes,
        )
    }

    /// Manually executes a report to the router.
    ///
    /// When a message is not being executed, then the user can trigger the execution manually.
    /// No verification over the transmitter, but the message needs to be in some commit report.
    /// It validates that the required time has passed since the commit and then executes the report.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the execution.
    /// * `raw_execution_report` - The serialized execution report containing the message and proofs.
    pub fn manually_execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        raw_execution_report: Vec<u8>,
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::execute::manually_execute(ctx, raw_execution_report, &token_indexes)
    }
}

#[error_code]
pub enum CcipOfframpError {
    // TODO review unused ones and remove them
    #[msg("The given sequence interval is invalid")]
    InvalidSequenceInterval,
    #[msg("The given Merkle Root is missing")]
    RootNotCommitted,
    #[msg("The given Merkle Root is already committed")]
    ExistingMerkleRoot,
    #[msg("The signer is unauthorized")]
    Unauthorized,
    #[msg("Invalid inputs")]
    InvalidInputs,
    #[msg("Source chain selector not supported")]
    UnsupportedSourceChainSelector,
    #[msg("Destination chain selector not supported")]
    UnsupportedDestinationChainSelector,
    #[msg("Invalid Proof for Merkle Root")]
    InvalidProof,
    #[msg("Invalid message format")]
    InvalidMessage,
    #[msg("Reached max sequence number")]
    ReachedMaxSequenceNumber,
    #[msg("Manual execution not allowed")]
    ManualExecutionNotAllowed,
    #[msg("Invalid pool account account indices")]
    InvalidInputsTokenIndices,
    #[msg("Invalid pool accounts")]
    InvalidInputsPoolAccounts,
    #[msg("Invalid token accounts")]
    InvalidInputsTokenAccounts,
    #[msg("Invalid config account")]
    InvalidInputsConfigAccounts,
    #[msg("Invalid Token Admin Registry account")]
    InvalidInputsTokenAdminRegistryAccounts,
    #[msg("Invalid LookupTable account")]
    InvalidInputsLookupTableAccounts,
    #[msg("Invalid LookupTable account writable access")]
    InvalidInputsLookupTableAccountWritable,
    #[msg("Cannot send zero tokens")]
    InvalidInputsTokenAmount,
    #[msg("Release or mint balance mismatch")]
    OfframpReleaseMintBalanceMismatch,
    #[msg("Invalid data length")]
    OfframpInvalidDataLength,
    #[msg("Stale commit report")]
    StaleCommitReport,
    #[msg("Destination chain disabled")]
    DestinationChainDisabled,
    #[msg("Fee token disabled")]
    FeeTokenDisabled,
    #[msg("Message exceeds maximum data size")]
    MessageTooLarge,
    #[msg("Message contains an unsupported number of tokens")]
    UnsupportedNumberOfTokens,
    #[msg("Chain family selector not supported")]
    UnsupportedChainFamilySelector,
    #[msg("Invalid EVM address")]
    InvalidEVMAddress,
    #[msg("Invalid encoding")]
    InvalidEncoding,
    #[msg("Invalid Associated Token Account address")]
    InvalidInputsAtaAddress,
    #[msg("Invalid Associated Token Account writable flag")]
    InvalidInputsAtaWritable,
    #[msg("Invalid token price")]
    InvalidTokenPrice,
    #[msg("Stale gas price")]
    StaleGasPrice,
    #[msg("Insufficient lamports")]
    InsufficientLamports,
    #[msg("Insufficient funds")]
    InsufficientFunds,
    #[msg("Unsupported token")]
    UnsupportedToken,
    #[msg("Inputs are missing token configuration")]
    InvalidInputsMissingTokenConfig,
    #[msg("Message fee is too high")]
    MessageFeeTooHigh,
    #[msg("Source token data is too large")]
    SourceTokenDataTooLarge,
    #[msg("Message gas limit too high")]
    MessageGasLimitTooHigh,
    #[msg("Extra arg out of order execution must be true")]
    ExtraArgOutOfOrderExecutionMustBeTrue,
    #[msg("New Admin can not be zero address")]
    InvalidTokenAdminRegistryInputsZeroAddress,
    #[msg("An already owned registry can not be proposed")]
    InvalidTokenAdminRegistryProposedAdmin,
    #[msg("Invalid writability bitmap")]
    InvalidWritabilityBitmap,
    #[msg("Invalid extra args tag")]
    InvalidExtraArgsTag,
    #[msg("Invalid chain family selector")]
    InvalidChainFamilySelector,
    #[msg("Invalid token receiver")]
    InvalidTokenReceiver,
    #[msg("Invalid SVM address")]
    InvalidSVMAddress,
    #[msg("Sender not allowed for that destination chain")]
    SenderNotAllowed,
}
