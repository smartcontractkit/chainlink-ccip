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
