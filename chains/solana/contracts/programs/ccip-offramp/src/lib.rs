use anchor_lang::error_code;
use anchor_lang::prelude::*;

declare_id!("offqSMQWgQud6WJz694LRzkeN5kMYpCHTpXQr3Rkcjm");

mod event;
mod messages;
mod ocr3base;

mod context;
use crate::context::*;

mod instructions;
use crate::instructions::router;

mod state;
use state::*;

use crate::event::admin::{ConfigSet, ReferenceAddressesSet};

#[program]
pub mod ccip_offramp {

    use super::*;

    #[allow(clippy::empty_line_after_outer_attr)]
    //////////////////////////
    /// Initialization Flow //
    //////////////////////////

    /// Initializes the CCIP Offramp, except for the config account (due to stack size limitations).
    ///
    /// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
    /// initialization methods should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization.
    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        let mut reference_addresses = ctx.accounts.reference_addresses.load_init()?;
        *reference_addresses = ReferenceAddresses {
            version: 1,
            router: ctx.accounts.router.key(),
            fee_quoter: ctx.accounts.fee_quoter.key(),
            rmn_remote: ctx.accounts.rmn_remote.key(),
            offramp_lookup_table: ctx.accounts.offramp_lookup_table.key(),
        };

        ctx.accounts.state.latest_price_sequence_number = 0;

        emit!(ReferenceAddressesSet {
            router: ctx.accounts.router.key(),
            fee_quoter: ctx.accounts.fee_quoter.key(),
            offramp_lookup_table: ctx.accounts.offramp_lookup_table.key(),
            rmn_remote: ctx.accounts.rmn_remote.key(),
        });

        Ok(())
    }

    /// Initializes the CCIP Offramp Config account.
    ///
    /// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
    /// initialization methods should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization of the config.
    /// * `svm_chain_selector` - The chain selector for SVM.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    pub fn initialize_config(
        ctx: Context<InitializeConfig>,
        svm_chain_selector: u64,
        enable_execution_after: i64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipOfframpError::InvalidVersion); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.default_code_version = CodeVersion::V1.into();
        config.svm_chain_selector = svm_chain_selector;
        config.enable_manual_execution_after = enable_execution_after;
        config.owner = ctx.accounts.authority.key();
        config.ocr3 = [
            Ocr3Config::new(OcrPluginType::Commit),
            Ocr3Config::new(OcrPluginType::Execution),
        ];

        emit!(ConfigSet {
            svm_chain_selector,
            enable_manual_execution_after: enable_execution_after,
        });

        Ok(())
    }

    /// Returns the program type (name) and version.
    /// Used by offchain code to easily determine which program & version is being interacted with.
    ///
    /// # Arguments
    /// * `ctx` - The context
    pub fn type_version(_ctx: Context<Empty>) -> Result<String> {
        let response = env!("CCIP_BUILD_TYPE_VERSION").to_string();
        msg!("{}", response);
        Ok(response)
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).transfer_ownership(ctx, proposed_owner)
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).accept_ownership(ctx)
    }

    #[allow(clippy::empty_line_after_outer_attr)]
    /////////////
    // Config //
    /////////////

    /// Sets the default code version to be used. This is then used by the slim routing layer to determine
    /// which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.
    ///
    /// Shared func signature with other programs
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `code_version` - The new code version to be set as default.
    pub fn set_default_code_version(
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()> {
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).set_default_code_version(ctx, code_version)
    }

    /// Updates reference addresses in the offramp contract, such as
    /// the CCIP router, Fee Quoter, and the Offramp Lookup Table.
    /// Only the Admin may update these addresses.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the reference addresses.
    /// * `router` - The router address to be set.
    /// * `fee_quoter` - The fee_quoter address to be set.
    /// * `offramp_lookup_table` - The offramp_lookup_table address to be set.
    /// * `rmn_remote` - The rmn_remote address to be set.
    pub fn update_reference_addresses(
        ctx: Context<UpdateReferenceAddresses>,
        router: Pubkey,
        fee_quoter: Pubkey,
        offramp_lookup_table: Pubkey,
        rmn_remote: Pubkey,
    ) -> Result<()> {
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).update_reference_addresses(
            ctx,
            router,
            fee_quoter,
            offramp_lookup_table,
            rmn_remote,
        )
    }

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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).add_source_chain(
            ctx,
            new_chain_selector,
            source_chain_config,
        )
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version)
            .disable_source_chain_selector(ctx, source_chain_selector)
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).update_source_chain_config(
            ctx,
            source_chain_selector,
            source_chain_config,
        )
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).update_svm_chain_selector(ctx, new_chain_selector)
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
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version)
            .update_enable_manual_execution_after(ctx, new_enable_manual_execution_after)
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
        plugin_type: OcrPluginType,
        config_info: Ocr3ConfigInfo,
        signers: Vec<[u8; 20]>,
        transmitters: Vec<Pubkey>,
    ) -> Result<()> {
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::admin(default_code_version).set_ocr_config(
            ctx,
            plugin_type,
            config_info,
            signers,
            transmitters,
        )
    }

    #[allow(clippy::empty_line_after_outer_attr)]
    ////////////////////
    /// Off Ramp Flow //
    ////////////////////

    /// Commits a report to the router, containing a Merkle Root.
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
        let lane_code_version = ctx.accounts.source_chain.config.lane_code_version;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::commit(lane_code_version, default_code_version).commit(
            ctx,
            report_context_byte_words,
            raw_report,
            rs,
            ss,
            raw_vs,
        )
    }

    /// Commits a report to the router, with price updates only.
    ///
    /// The method name needs to be commit with Anchor encoding.
    ///
    /// This function is called by the OffChain when committing one Report to the SVM Router,
    /// containing only price updates and no merkle root.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the commit.
    /// * `report_context_byte_words` - consists of:
    ///     * report_context_byte_words[0]: ConfigDigest
    ///     * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
    /// * `raw_report` - The serialized commit input report containing the price updates,
    ///    with no merkle root.
    /// * `rs` - slice of R components of signatures
    /// * `ss` - slice of S components of signatures
    /// * `raw_vs` - array of V components of signatures
    pub fn commit_price_only<'info>(
        ctx: Context<'_, '_, 'info, 'info, PriceOnlyCommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 2],
        raw_report: Vec<u8>,
        rs: Vec<[u8; 32]>,
        ss: Vec<[u8; 32]>,
        raw_vs: [u8; 32],
    ) -> Result<()> {
        // there is no lane here in this case of commit, so use default code version
        let lane_code_version = CodeVersion::Default;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::commit(lane_code_version, default_code_version).commit_price_only(
            ctx,
            report_context_byte_words,
            raw_report,
            rs,
            ss,
            raw_vs,
        )
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
        let lane_code_version = ctx.accounts.source_chain.config.lane_code_version;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::execute(lane_code_version, default_code_version).execute(
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
        let lane_code_version = ctx.accounts.source_chain.config.lane_code_version;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::execute(lane_code_version, default_code_version).manually_execute(
            ctx,
            raw_execution_report,
            &token_indexes,
        )
    }

    /// Initializes and/or inserts a chunk of report data to an execution report buffer.
    ///
    /// When execution reports are too large to fit in a single transaction, they can be chopped
    /// up in chunks first (as a special case, one chunk is also acceptable), and pre-buffered
    /// via multiple calls to this instruction.
    ///
    /// There's no need to pre-initialize the buffer: all chunks can be sent concurrently, and the
    /// first one to arrive will initialize the buffer.
    ///
    /// To benefit from buffering, the eventual call to `execute` or `manually_execute` must
    /// include an additional `remaining_account` with the PDA derived from
    /// ["execution_report_buffer", <buffer_id>, <caller_pubkey>].
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for buffering.
    /// * `buffer_id` - An arbitrary buffer id defined by the caller (could be the message_id). Max 32 bytes.
    /// * `report_length` - Total length in bytes of the execution report.
    /// * `chunk` - The specific chunk to add to the buffer. Chunk must have a consistent size, except
    ///    the last one in the buffer, which may be smaller.
    /// * `chunk_index` - The index of this chunk.
    /// * `num_chunks` - The total number of chunks in the report.
    pub fn buffer_execution_report<'info>(
        ctx: Context<'_, '_, 'info, 'info, BufferExecutionReportContext<'info>>,
        buffer_id: Vec<u8>,
        report_length: u32,
        chunk: Vec<u8>,
        chunk_index: u8,
        num_chunks: u8,
    ) -> Result<()> {
        // Execution report buffering doesn't need to be done differently per lane, so we
        // use the default code version here.
        let lane_code_version = CodeVersion::Default;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::execute(lane_code_version, default_code_version).buffer_execution_report(
            ctx,
            buffer_id,
            report_length,
            chunk,
            chunk_index,
            num_chunks,
        )
    }

    /// Closes the execution report buffer to reclaim funds.
    ///
    /// Note this is only necessary when aborting a buffered transaction, or when a mistake
    /// was made when buffering data. The buffer account will otherwise automatically close
    /// and return funds to the caller whenever buffered execution succeeds.
    pub fn close_execution_report_buffer(
        _ctx: Context<CloseExecutionReportBufferContext>,
        _buffer_id: Vec<u8>,
    ) -> Result<()> {
        Ok(())
    }

    /// Automatically derives all acounts required to call `ccip_execute`.
    ///
    /// This method receives the bare minimum amount of information needed to construct
    /// the entire account list to execute a transaction, and builds it iteratively
    /// over the course of multiple calls.
    ///
    /// The return type contains:
    ///
    /// * `accounts_to_save`: The caller must append these accounts to a list they maintain.
    ///   When complete, this list will contain all accounts needed to call `ccip_execute`.
    /// * `ask_again_with`: When `next_stage` is not empty, the caller must call `derive_accounts_execute`
    ///   again, including exactly these accounts as the `remaining_accounts`.
    /// * `lookup_tables_to_save`: The caller must save those LUTs. They can be used for `ccip_execute`.
    /// * `current_stage`: A string describing the current stage of the derivation process. When the stage
    ///   is "TokenTransferStaticAccounts/<N>/0", it means the `accounts_to_save` block in this response contains
    ///   all accounts relating to the Nth token being transferred. Use this information to construct
    ///   the `token_indexes` vector that `execute` requires.
    /// * `next_stage`: If nonempty, this means the instruction must get called again with this value
    ///   as the `stage` argument.
    ///
    /// Therefore, and starting with an empty `remaining_accounts` list, the caller must repeteadly
    /// call `derive_accounts_execute` until `next_stage` is returned empty.
    ///
    /// # Arguments
    ///
    /// * `ctx`: Context containing only the offramp config.
    /// * `stage`: Requested derivation stage. Pass "Start" the first time, then for each subsequent
    ///   call, pass the value returned in `response.next_stage` until empty.
    /// * `params`:
    ///    * `execute_caller`: Public key of the account that will sign the call to `ccip_execute`.
    ///    * `message_accounts`: If the transaction involves messaging, the message accounts.
    ///    * `source_chain_selector`: CCIP chain selector for the source chain.
    ///    * `mints_of_transferred_token`: List of all token mints for tokens being transferred (i.e.
    ///      the entries in `report.message.token_amounts.destination_address`.)
    ///    * `merkle_root`: Merkle root as per the commit report.
    ///    * `buffer_id`: If the execution will be buffered, the buffer id that will be used by the
    ///      `execute_caller`: If the execution will not be buffered, this should be empty.
    ///    * `token_receiver`: Receiver of token transfers, if any (i.e. report.message.token_receiver)
    pub fn derive_accounts_execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ViewConfigOnly<'info>>,
        params: DeriveAccountsExecuteParams,
        stage: String,
    ) -> Result<DeriveAccountsResponse> {
        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::execute(default_code_version, default_code_version)
            .derive_accounts_execute(ctx, params, stage)
    }

    pub fn close_commit_report_account(
        ctx: Context<CloseCommitReportAccount>,
        source_chain_selector: u64,
        root: Vec<u8>,
    ) -> Result<()> {
        // there is no lane here in this case of commit, so use default code version
        let lane_code_version = CodeVersion::Default;

        let default_code_version: CodeVersion = ctx
            .accounts
            .config
            .load()?
            .default_code_version
            .try_into()?;

        router::commit(lane_code_version, default_code_version).close_commit_report_account(
            ctx,
            source_chain_selector,
            root,
        )
    }
}

#[error_code]
pub enum CcipOfframpError {
    #[msg("The given sequence interval is invalid")]
    // offset error code so that they don't clash with other programs
    // (Anchor's base custom error code 6000 + offset 3000 = start at 9000)
    InvalidSequenceInterval = 3000,
    #[msg("The given Merkle Root is missing")]
    RootNotCommitted,
    #[msg("Invalid RMN Remote Address")]
    InvalidRMNRemoteAddress,
    #[msg("The given Merkle Root is already committed")]
    ExistingMerkleRoot,
    #[msg("The signer is unauthorized")]
    Unauthorized,
    #[msg("Invalid Nonce")]
    InvalidNonce,
    #[msg("Account should be writable")]
    InvalidInputsMissingWritable,
    #[msg("Onramp was not configured")]
    OnrampNotConfigured,
    #[msg("Failed to deserialize report")]
    FailedToDeserializeReport,
    #[msg("Invalid plugin type")]
    InvalidPluginType,
    #[msg("Invalid version of the onchain state")]
    InvalidVersion,
    #[msg("Commit report is missing expected price updates")]
    MissingExpectedPriceUpdates,
    #[msg("Commit report is missing expected merkle root")]
    MissingExpectedMerkleRoot,
    #[msg("Commit report contains unexpected merkle root")]
    UnexpectedMerkleRoot,
    #[msg("Proposed owner is the current owner")]
    RedundantOwnerProposal,
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
    #[msg("Number of accounts is invalid")]
    InvalidInputsNumberOfAccounts,
    #[msg("Invalid global state account address")]
    InvalidInputsGlobalStateAccount,
    #[msg("Invalid pool account account indices")]
    InvalidInputsTokenIndices,
    #[msg("Invalid pool accounts")]
    InvalidInputsPoolAccounts,
    #[msg("Invalid token accounts")]
    InvalidInputsTokenAccounts,
    #[msg("Invalid sysvar instructions account")]
    InvalidInputsSysvarAccount,
    #[msg("Invalid fee quoter account")]
    InvalidInputsFeeQuoterAccount,
    #[msg("Invalid offramp authorization account")]
    InvalidInputsAllowedOfframpAccount,
    #[msg("Invalid Token Admin Registry account")]
    InvalidInputsTokenAdminRegistryAccounts,
    #[msg("Invalid LookupTable account")]
    InvalidInputsLookupTableAccounts,
    #[msg("Invalid LookupTable account writable access")]
    InvalidInputsLookupTableAccountWritable,
    #[msg("Release or mint balance mismatch")]
    OfframpReleaseMintBalanceMismatch,
    #[msg("Invalid data length")]
    OfframpInvalidDataLength,
    #[msg("Stale commit report")]
    StaleCommitReport,
    #[msg("Invalid writability bitmap")]
    InvalidWritabilityBitmap,
    #[msg("Invalid code version")]
    InvalidCodeVersion,
    #[msg("Invalid config: F must be positive")]
    Ocr3InvalidConfigFMustBePositive,
    #[msg("Invalid config: Too many transmitters")]
    Ocr3InvalidConfigTooManyTransmitters,
    #[msg("Invalid config: No transmitters")]
    Ocr3InvalidConfigNoTransmitters,
    #[msg("Invalid config: Too many signers")]
    Ocr3InvalidConfigTooManySigners,
    #[msg("Invalid config: F is too high")]
    Ocr3InvalidConfigFIsTooHigh,
    #[msg("Invalid config: Repeated oracle address")]
    Ocr3InvalidConfigRepeatedOracle,
    #[msg("Wrong message length")]
    Ocr3WrongMessageLength,
    #[msg("Config digest mismatch")]
    Ocr3ConfigDigestMismatch,
    #[msg("Wrong number signatures")]
    Ocr3WrongNumberOfSignatures,
    #[msg("Unauthorized transmitter")]
    Ocr3UnauthorizedTransmitter,
    #[msg("Unauthorized signer")]
    Ocr3UnauthorizedSigner,
    #[msg("Non unique signatures")]
    Ocr3NonUniqueSignatures,
    #[msg("Oracle cannot be zero address")]
    Ocr3OracleCannotBeZeroAddress,
    #[msg("Static config cannot be changed")]
    Ocr3StaticConfigCannotBeChanged,
    #[msg("Incorrect plugin type")]
    Ocr3InvalidPluginType,
    #[msg("Invalid signature")]
    Ocr3InvalidSignature,
    #[msg("Signatures out of registration")]
    Ocr3SignaturesOutOfRegistration,
    #[msg("Invalid onramp address")]
    InvalidOnrampAddress,
    #[msg("Invalid external execution signer account")]
    InvalidInputsExternalExecutionSignerAccount,
    #[msg("Commit report has pending messages")]
    CommitReportHasPendingMessages,
    #[msg("The execution report buffer already contains that chunk")]
    ExecutionReportBufferAlreadyContainsChunk,
    #[msg("The execution report buffer is already initialized")]
    ExecutionReportBufferAlreadyInitialized,
    #[msg("Invalid length for execution report buffer")]
    ExecutionReportBufferInvalidLength,
    #[msg("Chunk lies outside the execution report buffer")]
    ExecutionReportBufferInvalidChunkIndex,
    #[msg("Total number of chunks is not consistent")]
    ExecutionReportBufferInvalidChunkNumber,
    #[msg("Chunk size is too small")]
    ExecutionReportBufferChunkSizeTooSmall,
    #[msg("Invalid chunk size")]
    ExecutionReportBufferInvalidChunkSize,
    #[msg("Invalid ID size for buffer")]
    ExecutionReportBufferInvalidIdSize,
    #[msg("Execution report buffer is not complete: chunks are missing")]
    ExecutionReportBufferIncomplete,
    #[msg("Execution report wasn't provided either directly or via buffer")]
    ExecutionReportUnavailable,
    #[msg("Invalid account list for PDA derivation")]
    InvalidAccountListForPdaDerivation,
    #[msg("Unexpected account derivation stage")]
    InvalidDerivationStage,
    #[msg("Token pool returned an unexpected derivation response")]
    InvalidTokenPoolAccountDerivationResponse,
    #[msg("Can't fit account derivation response.")]
    AccountDerivationResponseTooLarge,
    #[msg("Proposed owner is the default pubkey")]
    DefaultOwnerProposal,
}
