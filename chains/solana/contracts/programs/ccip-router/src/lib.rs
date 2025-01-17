use anchor_lang::error_code;
use anchor_lang::prelude::*;

mod context;
use crate::context::*;

mod token_context;
use crate::token_context::*;

mod state;
use crate::state::*;

mod event;
use crate::event::*;

mod messages;
use crate::messages::*;

mod instructions;
use crate::instructions::*;

// Anchor discriminators for CPI calls
const CCIP_RECEIVE_DISCRIMINATOR: [u8; 8] = [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]; // ccip_receive
const TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR: [u8; 8] =
    [0x72, 0xa1, 0x5e, 0x1d, 0x93, 0x19, 0xe8, 0xbf]; // lock_or_burn_tokens
const TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR: [u8; 8] =
    [0x5c, 0x64, 0x96, 0xc6, 0xfc, 0x3f, 0xa4, 0xe4]; // release_or_mint_tokens

declare_id!("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8");

#[program]
/// The `ccip_router` module contains the implementation of the Cross-Chain Interoperability Protocol (CCIP) Router.
///
/// This is the Collapsed Router Program for CCIP.
/// As it's upgradable persisting the same program id, there is no need to have an indirection of a Proxy Program.
/// This Router handles both the OnRamp and OffRamp flow of the CCIP Messages.
///
/// NOTE to devs: This file however should contain *no logic*, only the entrypoints to the different versioned modules,
/// thus making it easier to ensure later on that logic can be changed during upgrades without affecting the interface.
pub mod ccip_router {
    #![warn(missing_docs)]

    use super::*;

    /// Initializes the CCIP Router.
    ///
    /// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization.
    /// * `svm_chain_selector` - The chain selector for SVM.
    /// * `default_gas_limit` - The default gas limit for other destination chains.
    /// * `default_allow_out_of_order_execution` - Whether out-of-order execution is allowed by default for other destination chains.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    #[allow(clippy::too_many_arguments)]
    pub fn initialize(
        ctx: Context<InitializeCCIPRouter>,
        svm_chain_selector: u64,
        default_gas_limit: u128,
        default_allow_out_of_order_execution: bool,
        enable_execution_after: i64,
        fee_aggregator: Pubkey,
        link_token_mint: Pubkey,
        max_fee_juels_per_msg: u128,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipRouterError::InvalidInputs); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.svm_chain_selector = svm_chain_selector;
        config.default_gas_limit = default_gas_limit;
        config.enable_manual_execution_after = enable_execution_after;
        config.link_token_mint = link_token_mint;
        config.max_fee_juels_per_msg = max_fee_juels_per_msg;

        if default_allow_out_of_order_execution {
            config.default_allow_out_of_order_execution = 1;
        }

        config.owner = ctx.accounts.authority.key();

        config.fee_aggregator = fee_aggregator;

        config.ocr3 = [
            Ocr3Config::new(OcrPluginType::Commit as u8),
            Ocr3Config::new(OcrPluginType::Execution as u8),
        ];

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

    /// Updates the fee aggregator in the router configuration.
    /// The Admin is the only one able to update the fee aggregator.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `fee_aggregator` - The new fee aggregator address (ATAs will be derived for it for each token).
    pub fn update_fee_aggregator(
        ctx: Context<UpdateConfigCCIPRouter>,
        fee_aggregator: Pubkey,
    ) -> Result<()> {
        v1::admin::update_fee_aggregator(ctx, fee_aggregator)
    }

    /// Adds a new chain selector to the router.
    ///
    /// The Admin needs to add any new chain supported (this means both OnRamp and OffRamp).
    /// When adding a new chain, the Admin needs to specify if it's enabled or not.
    /// They may enable only source, or only destination, or neither, or both.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for adding the chain selector.
    /// * `new_chain_selector` - The new chain selector to be added.
    /// * `source_chain_config` - The configuration for the chain as source.
    /// * `dest_chain_config` - The configuration for the chain as destination.
    pub fn add_chain_selector(
        ctx: Context<AddChainSelector>,
        new_chain_selector: u64,
        source_chain_config: SourceChainConfig,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        v1::admin::add_chain_selector(
            ctx,
            new_chain_selector,
            source_chain_config,
            dest_chain_config,
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
        ctx: Context<UpdateSourceChainSelectorConfig>,
        source_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::disable_source_chain_selector(ctx, source_chain_selector)
    }

    /// Disables the destination chain selector.
    ///
    /// The Admin is the only one able to disable the chain selector as destination. This method is thought of as an emergency kill-switch.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for disabling the chain selector.
    /// * `dest_chain_selector` - The destination chain selector to be disabled.
    pub fn disable_dest_chain_selector(
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::disable_dest_chain_selector(ctx, dest_chain_selector)
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
        ctx: Context<UpdateSourceChainSelectorConfig>,
        source_chain_selector: u64,
        source_chain_config: SourceChainConfig,
    ) -> Result<()> {
        v1::admin::update_source_chain_config(ctx, source_chain_selector, source_chain_config)
    }

    /// Updates the configuration of the destination chain selector.
    ///
    /// The Admin is the only one able to update the destination chain config.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the chain selector.
    /// * `dest_chain_selector` - The destination chain selector to be updated.
    /// * `dest_chain_config` - The new configuration for the destination chain.
    pub fn update_dest_chain_config(
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        v1::admin::update_dest_chain_config(ctx, dest_chain_selector, dest_chain_config)
    }

    /// Updates the SVM chain selector in the router configuration.
    ///
    /// This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_chain_selector` - The new chain selector for SVM.
    pub fn update_svm_chain_selector(
        ctx: Context<UpdateConfigCCIPRouter>,
        new_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::update_svm_chain_selector(ctx, new_chain_selector)
    }

    /// Updates the default gas limit in the router configuration.
    ///
    /// This change affects the default value for gas limit on every other destination chain.
    /// The Admin is the only one able to update the default gas limit.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_gas_limit` - The new default gas limit.
    pub fn update_default_gas_limit(
        ctx: Context<UpdateConfigCCIPRouter>,
        new_gas_limit: u128,
    ) -> Result<()> {
        v1::admin::update_default_gas_limit(ctx, new_gas_limit)
    }

    /// Updates the default setting for allowing out-of-order execution for other destination chains.
    /// The Admin is the only one able to update this config.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_allow_out_of_order_execution` - The new setting for allowing out-of-order execution.
    pub fn update_default_allow_out_of_order_execution(
        ctx: Context<UpdateConfigCCIPRouter>,
        new_allow_out_of_order_execution: bool,
    ) -> Result<()> {
        v1::admin::update_default_allow_out_of_order_execution(
            ctx,
            new_allow_out_of_order_execution,
        )
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
        ctx: Context<UpdateConfigCCIPRouter>,
        new_enable_manual_execution_after: i64,
    ) -> Result<()> {
        v1::admin::update_enable_manual_execution_after(ctx, new_enable_manual_execution_after)
    }

    /// Registers the Token Admin Registry via the CCIP Admin
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    /// * `mint` - The public key of the token mint.
    /// * `token_admin_registry_admin` - The public key of the token admin registry admin.
    pub fn register_token_admin_registry_via_get_ccip_admin(
        ctx: Context<RegisterTokenAdminRegistryViaGetCCIPAdmin>,
        mint: Pubkey, // should we validate that this is a real token program?
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::register_token_admin_registry_via_get_ccip_admin(
            ctx,
            mint,
            token_admin_registry_admin,
        )
    }

    /// Registers the Token Admin Registry via the token owner.
    ///
    /// The Authority of the Mint Token can claim the registry of the token.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    pub fn register_token_admin_registry_via_owner(
        ctx: Context<RegisterTokenAdminRegistryViaOwner>,
    ) -> Result<()> {
        v1::token_admin_registry::register_token_admin_registry_via_owner(ctx)
    }

    /// Sets the pool lookup table for a given token mint.
    ///
    /// The administrator of the token admin registry can set the pool lookup table for a given token mint.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for setting the pool.
    /// * `mint` - The public key of the token mint.
    /// * `pool_lookup_table` - The public key of the pool lookup table, this address will be used for validations when interacting with the pool.
    /// * `is_writable` - index of account in lookup table that is writable
    pub fn set_pool(
        ctx: Context<SetPoolTokenAdminRegistry>,
        mint: Pubkey,
        writable_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::token_admin_registry::set_pool(ctx, mint, writable_indexes)
    }

    /// Transfers the admin role of the token admin registry to a new admin.
    ///
    /// Only the Admin can transfer the Admin Role of the Token Admin Registry, this setups the Pending Admin and then it's their responsibility to accept the role.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the transfer.
    /// * `mint` - The public key of the token mint.
    /// * `new_admin` - The public key of the new admin.
    pub fn transfer_admin_role_token_admin_registry(
        ctx: Context<ModifyTokenAdminRegistry>,
        mint: Pubkey,
        new_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::transfer_admin_role_token_admin_registry(ctx, mint, new_admin)
    }

    /// Accepts the admin role of the token admin registry.
    ///
    /// The Pending Admin must call this function to accept the admin role of the Token Admin Registry.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for accepting the admin role.
    /// * `mint` - The public key of the token mint.
    pub fn accept_admin_role_token_admin_registry(
        ctx: Context<AcceptAdminRoleTokenAdminRegistry>,
        mint: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::accept_admin_role_token_admin_registry(ctx, mint)
    }

    /// Sets the token billing configuration.
    ///
    /// Only CCIP Admin can set the token billing configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for setting the token billing configuration.
    /// * `chain_selector` - The chain selector.
    /// * `mint` - The public key of the token mint.
    /// * `cfg` - The token billing configuration.
    pub fn set_token_billing(
        ctx: Context<SetTokenBillingConfig>,
        chain_selector: u64,
        mint: Pubkey,
        cfg: TokenBilling,
    ) -> Result<()> {
        v1::admin::set_token_billing(ctx, chain_selector, mint, cfg)
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

    /// Adds a billing token configuration.
    /// Only CCIP Admin can add a billing token configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for adding the billing token configuration.
    /// * `config` - The billing token configuration to be added.
    pub fn add_billing_token_config(
        ctx: Context<AddBillingTokenConfig>,
        config: BillingTokenConfig,
    ) -> Result<()> {
        v1::admin::add_billing_token_config(ctx, config)
    }

    /// Updates the billing token configuration.
    /// Only CCIP Admin can update a billing token configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the billing token configuration.
    /// * `config` - The new billing token configuration.
    pub fn update_billing_token_config(
        ctx: Context<UpdateBillingTokenConfig>,
        config: BillingTokenConfig,
    ) -> Result<()> {
        v1::admin::update_billing_token_config(ctx, config)
    }

    /// Removes the billing token configuration.
    /// Only CCIP Admin can remove a billing token configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for removing the billing token configuration.
    pub fn remove_billing_token_config(ctx: Context<RemoveBillingTokenConfig>) -> Result<()> {
        v1::admin::remove_billing_token_config(ctx)
    }

    /// Calculates the fee for sending a message to the destination chain.
    ///
    /// # Arguments
    ///
    /// * `_ctx` - The context containing the accounts required for the fee calculation.
    /// * `dest_chain_selector` - The chain selector for the destination chain.
    /// * `message` - The message to be sent.
    ///
    /// # Additional accounts
    ///
    /// In addition to the fixed amount of accounts defined in the `GetFee` context,
    /// the following accounts must be provided:
    ///
    /// * First, the billing token config accounts for each token sent with the message, sequentially.
    ///   For each token with no billing config account (i.e. tokens that cannot be possibly used as fee
    ///   tokens, which also have no BPS fees enabled) the ZERO address must be provided instead.
    /// * Then, the per chain / per token config of every token sent with the message, sequentially
    ///   in the same order.
    ///
    /// # Returns
    ///
    /// The fee amount in u64.
    pub fn get_fee<'info>(
        ctx: Context<'_, '_, 'info, 'info, GetFee>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
    ) -> Result<u64> {
        v1::onramp::get_fee(ctx, dest_chain_selector, message)
    }

    /// Transfers the accumulated billed fees in a particular token to an arbitrary token account.
    /// Only the CCIP Admin can withdraw billed funds.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the transfer of billed fees.
    /// * `transfer_all` - A flag indicating whether to transfer all the accumulated fees in that token or not.
    /// * `desired_amount` - The amount to transfer. If `transfer_all` is true, this value must be 0.
    pub fn withdraw_billed_funds(
        ctx: Context<WithdrawBilledFunds>,
        transfer_all: bool,
        desired_amount: u64, // if transfer_all is false, this value must be 0
    ) -> Result<()> {
        v1::admin::withdraw_billed_funds(ctx, transfer_all, desired_amount)
    }

    /// ON RAMP FLOW
    /// Sends a message to the destination chain.
    ///
    /// Request a message to be sent to the destination chain.
    /// The method name needs to be ccip_send with Anchor encoding.
    /// This function is called by the CCIP Sender Contract (or final user) to send a message to the CCIP Router.
    /// The message will be sent to the receiver on the destination chain selector.
    /// This message emits the event CCIPSendRequested with all the necessary data to be retrieved by the OffChain Code
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for sending the message.
    /// * `dest_chain_selector` - The chain selector for the destination chain.
    /// * `message` - The message to be sent. The size limit of data is 256 bytes.
    pub fn ccip_send<'info>(
        ctx: Context<'_, '_, 'info, 'info, CcipSend<'info>>,
        dest_chain_selector: u64,
        message: SVM2AnyMessage,
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::onramp::ccip_send(ctx, dest_chain_selector, message, token_indexes)
    }

    /// OFF RAMP FLOW
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
    ///     * report_context_byte_words[2]: ExtraHash
    /// * `report` - The commit input report, single merkle root with RMN signatures and price updates
    /// * `signatures` - The list of signatures. v0.29.0 - anchor idl does not build with ocr3base::SIGNATURE_LENGTH
    pub fn commit<'info>(
        ctx: Context<'_, '_, 'info, 'info, CommitReportContext<'info>>,
        report_context_byte_words: [[u8; 32]; 3],
        report: CommitInput,
        signatures: Vec<[u8; 65]>,
    ) -> Result<()> {
        v1::offramp::commit(ctx, report_context_byte_words, report, signatures)
    }

    /// OFF RAMP FLOW
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
    /// * `execution_report` - the execution report containing only one message and proofs
    /// * `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)
    /// *  consists of:
    ///     * report_context_byte_words[0]: ConfigDigest
    ///     * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
    ///     * report_context_byte_words[2]: ExtraHash
    pub fn execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        execution_report: ExecutionReportSingleChain,
        report_context_byte_words: [[u8; 32]; 3],
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::offramp::execute(
            ctx,
            execution_report,
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
    /// * `execution_report` - The execution report containing the message and proofs.
    pub fn manually_execute<'info>(
        ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
        execution_report: ExecutionReportSingleChain,
        token_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::offramp::manually_execute(ctx, execution_report, &token_indexes)
    }
}

#[error_code]
pub enum CcipRouterError {
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
}
