use anchor_lang::prelude::*;

declare_id!("FeeVB9Q77QvyaENRL1i77BjW6cTkaWwNLjNbZg9JHqpw");

mod context;
use context::*;

mod messages;
use messages::*;

mod state;
use state::*;

mod event;

mod instructions;
use instructions::v1;

#[program]
pub mod fee_quoter {
    use super::*;

    /// Initializes the Fee Quoter.
    ///
    /// The initialization is responsibility of Admin, nothing more than calling this method should be done first.
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
        ctx: Context<Initialize>,
        onramp: Pubkey,
        link_token_mint: Pubkey,
        max_fee_juels_per_msg: u128,
    ) -> Result<()> {
        ctx.accounts.config.set_inner(Config {
            version: 1,
            owner: ctx.accounts.authority.key(),
            proposed_owner: Pubkey::default(),
            onramp,
            max_fee_juels_per_msg,
            link_token_mint,
        });

        // ctx.accounts.state.latest_price_sequence_number = 0; // TODO each offramp has its own price seq_nr?

        Ok(())
    }

    /// Transfers the ownership of the fee quoter to a new proposed owner.
    ///
    /// Shared func signature with other programs
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the transfer.
    /// * `proposed_owner` - The public key of the new proposed owner.
    pub fn transfer_ownership(ctx: Context<UpdateConfig>, new_owner: Pubkey) -> Result<()> {
        v1::admin::transfer_ownership(ctx, new_owner)
    }

    /// Accepts the ownership of the fee quoter by the proposed owner.
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

    /// Adds a new destination chain selector to the fee quoter.
    ///
    /// The Admin needs to add any new chain supported.
    /// When adding a new chain, the Admin needs to specify if it's enabled or not.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for adding the chain selector.
    /// * `chain_selector` - The new chain selector to be added.
    /// * `dest_chain_config` - The configuration for the chain as destination.
    pub fn add_dest_chain(
        ctx: Context<AddDestChain>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        v1::admin::add_dest_chain(ctx, chain_selector, dest_chain_config)
    }

    /// Disables the destination chain selector.
    ///
    /// The Admin is the only one able to disable the chain selector as destination. This method is thought of as an emergency kill-switch.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for disabling the chain selector.
    /// * `chain_selector` - The destination chain selector to be disabled.
    pub fn disable_dest_chain(
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
    ) -> Result<()> {
        v1::admin::disable_dest_chain(ctx, chain_selector)
    }

    /// Updates the configuration of the destination chain selector.
    ///
    /// The Admin is the only one able to update the destination chain config.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the chain selector.
    /// * `chain_selector` - The destination chain selector to be updated.
    /// * `dest_chain_config` - The new configuration for the destination chain.
    pub fn update_dest_chain_config(
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        v1::admin::update_dest_chain_config(ctx, chain_selector, dest_chain_config)
    }

    /// Calculates the fee for sending a message to the destination chain.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for the fee calculation.
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
        v1::public::get_fee(ctx, dest_chain_selector, message)
    }
}

#[error_code]
pub enum FeeQuoterError {
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
}
