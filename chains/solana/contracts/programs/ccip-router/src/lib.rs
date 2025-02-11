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
const TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR: [u8; 8] =
    [0x72, 0xa1, 0x5e, 0x1d, 0x93, 0x19, 0xe8, 0xbf]; // lock_or_burn_tokens

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

    //////////////////////////
    /// Initialization Flow //
    //////////////////////////

    /// Initializes the CCIP Router.
    ///
    /// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization.
    /// * `svm_chain_selector` - The chain selector for SVM.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    #[allow(clippy::too_many_arguments)]
    pub fn initialize(
        ctx: Context<InitializeCCIPRouter>,
        svm_chain_selector: u64,
        fee_aggregator: Pubkey,
        fee_quoter: Pubkey,
        link_token_mint: Pubkey,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipRouterError::InvalidVersion); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.svm_chain_selector = svm_chain_selector;
        config.link_token_mint = link_token_mint;

        config.fee_quoter = fee_quoter;

        config.owner = ctx.accounts.authority.key();

        config.fee_aggregator = fee_aggregator;

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
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        v1::admin::add_chain_selector(ctx, new_chain_selector, dest_chain_config)
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

    /// Add an offramp address to the list of offramps allowed by the router, for a
    /// particular source chain. External users will check this list before accepting
    /// a `ccip_receive` CPI.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for this operation.
    /// * `source_chain_selector` - The source chain for the offramp's lane.
    /// * `offramp` - The offramp's address.
    pub fn add_offramp(
        ctx: Context<AddOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()> {
        v1::admin::add_offramp(ctx, source_chain_selector, offramp)
    }

    /// Remove an offramp address from the list of offramps allowed by the router, for a
    /// particular source chain. External users will check this list before accepting
    /// a `ccip_receive` CPI.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for this operation.
    /// * `source_chain_selector` - The source chain for the offramp's lane.
    /// * `offramp` - The offramp's address.
    pub fn remove_offramp(
        ctx: Context<RemoveOfframp>,
        source_chain_selector: u64,
        offramp: Pubkey,
    ) -> Result<()> {
        v1::admin::remove_offramp(ctx, source_chain_selector, offramp)
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

    /// Bumps the CCIP version for a destination chain.
    /// This effectively just resets the sequence number of the destination chain state.
    ///
    /// # Arguments
    /// * `ctx` - The context containing the accounts required for the bump.
    /// * `dest_chain_selector` - The destination chain selector to bump version for.
    pub fn bump_ccip_version_for_dest_chain(
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::bump_ccip_version_for_dest_chain(ctx, dest_chain_selector)
    }

    /// Rolls back the CCIP version for a destination chain.
    /// This effectively just restores the old version's sequence number of the destination chain state.
    ///
    /// # Arguments
    /// * `ctx` - The context containing the accounts required for the rollback.
    /// * `dest_chain_selector` - The destination chain selector to rollback the version for.
    pub fn rollback_ccip_version_for_dest_chain(
        ctx: Context<UpdateDestChainSelectorConfig>,
        dest_chain_selector: u64,
    ) -> Result<()> {
        v1::admin::rollback_ccip_version_for_dest_chain(ctx, dest_chain_selector)
    }

    ///////////////////////////
    /// Token Admin Registry //
    ///////////////////////////

    /// Registers the Token Admin Registry via the CCIP Admin
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    /// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
    pub fn ccip_admin_propose_administrator(
        ctx: Context<RegisterTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::ccip_admin_propose_administrator(ctx, token_admin_registry_admin)
    }

    /// Overrides the pending admin of the Token Admin Registry
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    /// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
    pub fn ccip_admin_override_pending_administrator(
        ctx: Context<OverridePendingTokenAdminRegistryByCCIPAdmin>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::ccip_admin_override_pending_administrator(
            ctx,
            token_admin_registry_admin,
        )
    }

    /// Registers the Token Admin Registry by the token owner.
    ///
    /// The Authority of the Mint Token can claim the registry of the token.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    /// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
    pub fn owner_propose_administrator(
        ctx: Context<RegisterTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::owner_propose_administrator(ctx, token_admin_registry_admin)
    }

    /// Overrides the pending admin of the Token Admin Registry by the token owner
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for registration.
    /// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
    pub fn owner_override_pending_administrator(
        ctx: Context<OverridePendingTokenAdminRegistryByOwner>,
        token_admin_registry_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::owner_override_pending_administrator(
            ctx,
            token_admin_registry_admin,
        )
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
    ) -> Result<()> {
        v1::token_admin_registry::accept_admin_role_token_admin_registry(ctx)
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
        new_admin: Pubkey,
    ) -> Result<()> {
        v1::token_admin_registry::transfer_admin_role_token_admin_registry(ctx, new_admin)
    }

    /// Sets the pool lookup table for a given token mint.
    ///
    /// The administrator of the token admin registry can set the pool lookup table for a given token mint.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for setting the pool.
    /// * `writable_indexes` - a bit map of the indexes of the accounts in lookup table that are writable
    pub fn set_pool(
        ctx: Context<SetPoolTokenAdminRegistry>,
        writable_indexes: Vec<u8>,
    ) -> Result<()> {
        v1::token_admin_registry::set_pool(ctx, writable_indexes)
    }

    //////////////
    /// Billing //
    //////////////

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

    ///////////////////
    /// On Ramp Flow //
    ///////////////////

    /// Sends a message to the destination chain.
    ///
    /// Request a message to be sent to the destination chain.
    /// The method name needs to be ccip_send with Anchor encoding.
    /// This function is called by the CCIP Sender Contract (or final user) to send a message to the CCIP Router.
    /// The message will be sent to the receiver on the destination chain selector.
    /// This message emits the event CCIPMessageSent with all the necessary data to be retrieved by the OffChain Code
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
    ) -> Result<[u8; 32]> {
        v1::onramp::ccip_send(ctx, dest_chain_selector, message, token_indexes)
    }
}

// TODO this is a hack because Anchor + Anchor-Go fail to include all errors in the IDL and the gobindings.
// By having this first (though unused) error enum here, it does pick up the actual (second) error enum
#[error_code]
pub enum AnchorErrorHack {
    Something,
    Else,
}

#[error_code]
pub enum CcipRouterError {
    #[msg("The signer is unauthorized")]
    Unauthorized,
    #[msg("Mint account input is invalid")]
    InvalidInputsMint,
    #[msg("Invalid version of the onchain state")]
    InvalidVersion,
    #[msg("Fee token doesn't match transfer token")]
    FeeTokenMismatch,
    #[msg("Proposed owner is the current owner")]
    RedundantOwnerProposal,
    #[msg("Reached max sequence number")]
    ReachedMaxSequenceNumber,
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
    #[msg("Must specify zero amount to send alongside transfer_all")]
    InvalidInputsTransferAllAmount,
    #[msg("Invalid Associated Token Account address")]
    InvalidInputsAtaAddress,
    #[msg("Invalid Associated Token Account writable flag")]
    InvalidInputsAtaWritable,
    #[msg("Chain selector is invalid")]
    InvalidInputsChainSelector,
    #[msg("Insufficient lamports")]
    InsufficientLamports,
    #[msg("Insufficient funds")]
    InsufficientFunds,
    #[msg("Source token data is too large")]
    SourceTokenDataTooLarge,
    #[msg("New Admin can not be zero address")]
    InvalidTokenAdminRegistryInputsZeroAddress,
    #[msg("An already owned registry can not be proposed")]
    InvalidTokenAdminRegistryProposedAdmin,
    #[msg("Sender not allowed for that destination chain")]
    SenderNotAllowed,
    #[msg("Invalid rollback attempt on the CCIP version of the onramp to the destination chain")]
    InvalidCcipVersionRollback,
}
