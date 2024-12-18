use std::cell::Ref;

use anchor_lang::error_code;
use anchor_lang::prelude::*;
use anchor_spl::token_interface;
use bytemuck::Zeroable;
use solana_program::{instruction::Instruction, program::invoke_signed};

mod context;
use crate::context::*;

mod token_context;
use crate::token_context::*;

mod state;
use crate::state::*;

mod event;
use crate::event::*;

mod merkle;
use crate::merkle::*;

mod messages;
use crate::messages::*;

mod pools;
use crate::pools::*;

mod ocr3base;
use crate::ocr3base::*;

mod fee_quoter;
use crate::fee_quoter::*;

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
    /// * `solana_chain_selector` - The chain selector for Solana.
    /// * `default_gas_limit` - The default gas limit for other destination chains.
    /// * `default_allow_out_of_order_execution` - Whether out-of-order execution is allowed by default for other destination chains.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    pub fn initialize(
        ctx: Context<InitializeCCIPRouter>,
        solana_chain_selector: u64,
        default_gas_limit: u128,
        default_allow_out_of_order_execution: bool,
        enable_execution_after: i64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipRouterError::InvalidInputs); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.solana_chain_selector = solana_chain_selector;
        config.default_gas_limit = default_gas_limit;
        config.enable_manual_execution_after = enable_execution_after;

        if default_allow_out_of_order_execution {
            config.default_allow_out_of_order_execution = 1;
        }

        config.owner = ctx.accounts.authority.key();

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
        let mut config = ctx.accounts.config.load_mut()?;
        require!(
            proposed_owner != config.owner,
            CcipRouterError::InvalidInputs
        );
        config.proposed_owner = proposed_owner;
        Ok(())
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
        let mut config = ctx.accounts.config.load_mut()?;
        config.owner = std::mem::take(&mut config.proposed_owner);
        config.proposed_owner = Pubkey::new_from_array([0; 32]);
        Ok(())
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
        // Set source chain config & state
        let source_chain_state = &mut ctx.accounts.source_chain_state;
        validate_source_chain_config(new_chain_selector, &source_chain_config)?;
        source_chain_state.version = 1;
        source_chain_state.config = source_chain_config.clone();
        source_chain_state.state = SourceChainState { min_seq_nr: 1 };

        // Set dest chain config & state
        let dest_chain_state = &mut ctx.accounts.dest_chain_state;
        validate_dest_chain_config(new_chain_selector, &dest_chain_config)?;
        dest_chain_state.version = 1;
        dest_chain_state.config = dest_chain_config.clone();
        dest_chain_state.state = DestChainState {
            sequence_number: 0,
            usd_per_unit_gas: TimestampedPackedU224 {
                value: [0; 28],
                timestamp: 0,
            },
        };

        emit!(SourceChainAdded {
            source_chain_selector: new_chain_selector,
            source_chain_config,
        });
        emit!(DestChainAdded {
            dest_chain_selector: new_chain_selector,
            dest_chain_config,
        });

        Ok(())
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
        let chain_state = &mut ctx.accounts.source_chain_state;

        chain_state.config.is_enabled = false;

        emit!(SourceChainConfigUpdated {
            source_chain_selector,
            source_chain_config: chain_state.config.clone(),
        });

        Ok(())
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
        let chain_state = &mut ctx.accounts.dest_chain_state;

        chain_state.config.is_enabled = false;

        emit!(DestChainConfigUpdated {
            dest_chain_selector,
            dest_chain_config: chain_state.config.clone(),
        });

        Ok(())
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
        validate_source_chain_config(source_chain_selector, &source_chain_config)?;

        ctx.accounts.source_chain_state.config = source_chain_config.clone();

        emit!(SourceChainConfigUpdated {
            source_chain_selector,
            source_chain_config,
        });
        Ok(())
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
        validate_dest_chain_config(dest_chain_selector, &dest_chain_config)?;

        ctx.accounts.dest_chain_state.config = dest_chain_config.clone();

        emit!(DestChainConfigUpdated {
            dest_chain_selector,
            dest_chain_config,
        });
        Ok(())
    }

    /// Updates the Solana chain selector in the router configuration.
    ///
    /// This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for updating the configuration.
    /// * `new_chain_selector` - The new chain selector for Solana.
    pub fn update_solana_chain_selector(
        ctx: Context<UpdateConfigCCIPRouter>,
        new_chain_selector: u64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_mut()?;

        config.solana_chain_selector = new_chain_selector;

        Ok(())
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
        let mut config = ctx.accounts.config.load_mut()?;

        config.default_gas_limit = new_gas_limit;

        Ok(())
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
        let mut config = ctx.accounts.config.load_mut()?;

        let mut v = 0_u8;
        if new_allow_out_of_order_execution {
            v = 1;
        }
        config.default_allow_out_of_order_execution = v;

        Ok(())
    }

    /// Updates the minimum amount of time required between a message being committed and when it can be manually executed.
    ///
    /// This is part of the OffRamp Configuration for Solana.
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
        let mut config = ctx.accounts.config.load_mut()?;

        config.enable_manual_execution_after = new_enable_manual_execution_after;

        Ok(())
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
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        token_admin_registry.version = 1;
        token_admin_registry.administrator = token_admin_registry_admin;
        token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
        token_admin_registry.lookup_table = Pubkey::new_from_array([0; 32]);

        emit!(AdministratorRegistered {
            token_mint: mint,
            administrator: token_admin_registry_admin,
        });

        Ok(())
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
        let token_mint = ctx.accounts.mint.key().to_owned();
        let mint_authority = ctx.accounts.mint.mint_authority.to_owned();

        require!(
            mint_authority.is_some(),
            CcipRouterError::InvalidInputsPoolAccounts
        );

        let administrator = mint_authority.unwrap();

        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        token_admin_registry.version = 1;
        token_admin_registry.administrator = administrator;
        token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
        token_admin_registry.lookup_table = Pubkey::new_from_array([0; 32]);

        emit!(AdministratorRegistered {
            token_mint,
            administrator,
        });

        Ok(())
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
    pub fn set_pool(
        ctx: Context<ModifyTokenAdminRegistry>,
        mint: Pubkey,
        pool_lookup_table: Pubkey,
    ) -> Result<()> {
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        let previous_pool = token_admin_registry.lookup_table;
        token_admin_registry.lookup_table = pool_lookup_table;

        // TODO: Validate here that the lookup table has everything

        emit!(PoolSet {
            token: mint,
            previous_pool_lookup_table: previous_pool,
            new_pool_lookup_table: pool_lookup_table,
        });

        Ok(())
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
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;

        if new_admin == Pubkey::new_from_array([0; 32]) {
            token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);
        } else {
            token_admin_registry.pending_administrator = new_admin;
        }

        emit!(AdministratorTransferRequested {
            token: mint,
            current_admin: token_admin_registry.administrator,
            new_admin,
        });

        Ok(())
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
        let token_admin_registry = &mut ctx.accounts.token_admin_registry;
        let new_admin = token_admin_registry.pending_administrator;
        token_admin_registry.administrator = new_admin;
        token_admin_registry.pending_administrator = Pubkey::new_from_array([0; 32]);

        emit!(AdministratorTransferred {
            token: mint,
            new_admin,
        });

        Ok(())
    }

    /// Sets the token billing configuration.
    ///
    /// Only CCIP Admin can set the token billing configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for setting the token billing configuration.
    /// * `_chain_selector` - The chain selector.
    /// * `_mint` - The public key of the token mint.
    /// * `cfg` - The token billing configuration.
    pub fn set_token_billing(
        ctx: Context<SetTokenBillingConfig>,
        _chain_selector: u64,
        _mint: Pubkey,
        cfg: TokenBilling,
    ) -> Result<()> {
        ctx.accounts.per_chain_per_token_config.billing = cfg;
        Ok(())
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
        require!(plugin_type < 2, CcipRouterError::InvalidInputs);
        let mut config = ctx.accounts.config.load_mut()?;

        let is_commit = plugin_type == OcrPluginType::Commit as u8;

        config.ocr3[plugin_type as usize].set(
            plugin_type,
            Ocr3ConfigInfo {
                config_digest: config_info.config_digest,
                f: config_info.f,
                n: signers.len() as u8,
                is_signature_verification_enabled: if is_commit { 1 } else { 0 },
            },
            signers,
            transmitters,
        )?;

        if is_commit {
            // When the OCR config changes, we reset the sequence number since it is scoped per config digest.
            // Note that s_minSeqNr/roots do not need to be reset as the roots persist
            // across reconfigurations and are de-duplicated separately.
            ctx.accounts.state.latest_price_sequence_number = 0;
        }

        Ok(())
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
        emit!(FeeTokenAdded {
            fee_token: config.mint,
            enabled: config.enabled
        });
        ctx.accounts.billing_token_config.version = 1; // update this if we change the account struct
        ctx.accounts.billing_token_config.config = config;
        Ok(())
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
        if config.enabled != ctx.accounts.billing_token_config.config.enabled {
            // enabled/disabled status has changed
            match config.enabled {
                true => emit!(FeeTokenEnabled {
                    fee_token: config.mint
                }),
                false => emit!(FeeTokenDisabled {
                    fee_token: config.mint
                }),
            }
        }
        // TODO should we emit an event if the config has changed regardless of the enabled/disabled?

        ctx.accounts.billing_token_config.version = 1; // update this if we change the account struct
        ctx.accounts.billing_token_config.config = config;
        Ok(())
    }

    /// Removes the billing token configuration.
    /// Only CCIP Admin can remove a billing token configuration.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for removing the billing token configuration.
    pub fn remove_billing_token_config(ctx: Context<RemoveBillingTokenConfig>) -> Result<()> {
        // Close the receiver token account
        // The context constraints already enforce that it holds 0 balance of the target SPL token
        let cpi_accounts = token_interface::CloseAccount {
            account: ctx.accounts.fee_token_receiver.to_account_info(),
            destination: ctx.accounts.authority.to_account_info(),
            authority: ctx.accounts.fee_billing_signer.to_account_info(),
        };
        let cpi_program = ctx.accounts.token_program.to_account_info();
        let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[ctx.bumps.fee_billing_signer]];
        let signer_seeds = &[&seeds[..]];
        let cpi_ctx = CpiContext::new_with_signer(cpi_program, cpi_accounts, signer_seeds);

        token_interface::close_account(cpi_ctx)?;

        emit!(FeeTokenRemoved {
            fee_token: ctx.accounts.fee_token_mint.key()
        });
        Ok(())
    }

    /// Calculates the fee for sending a message to the destination chain.
    ///
    /// # Arguments
    ///
    /// * `_ctx` - The context containing the accounts required for the fee calculation.
    /// * `dest_chain_selector` - The chain selector for the destination chain.
    /// * `message` - The message to be sent.
    ///
    /// # Returns
    ///
    /// The fee amount in u64.
    pub fn get_fee(
        ctx: Context<GetFee>,
        dest_chain_selector: u64,
        message: Solana2AnyMessage,
    ) -> Result<u64> {
        Ok(fee_for_msg(
            dest_chain_selector,
            &message,
            &ctx.accounts.dest_chain_state,
            &ctx.accounts.billing_token_config.config,
        )?
        .amount)
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
        desired_amount: u64, // send 0 if transfer_all is false
    ) -> Result<()> {
        let transfer = token_interface::TransferChecked {
            from: ctx.accounts.fee_token_accum.to_account_info(),
            to: ctx.accounts.recipient.to_account_info(),
            mint: ctx.accounts.fee_token_mint.to_account_info(),
            authority: ctx.accounts.fee_billing_signer.to_account_info(),
        };

        let amount = if transfer_all {
            require!(desired_amount == 0, CcipRouterError::InvalidInputs);
            require!(
                ctx.accounts.fee_token_accum.amount > 0,
                CcipRouterError::InsufficientFunds
            );
            ctx.accounts.fee_token_accum.amount
        } else {
            require!(desired_amount > 0, CcipRouterError::InvalidInputs);
            require!(
                desired_amount <= ctx.accounts.fee_token_accum.amount,
                CcipRouterError::InsufficientFunds
            );
            desired_amount
        };

        do_transfer(
            ctx.accounts.token_program.to_account_info(),
            transfer,
            amount,
            ctx.accounts.fee_token_mint.decimals,
            ctx.bumps.fee_billing_signer,
        )
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
        message: Solana2AnyMessage,
    ) -> Result<()> {
        // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
        let config = ctx.accounts.config.load()?;

        let dest_chain = &mut ctx.accounts.dest_chain_state;
        let fee_token_config = &ctx.accounts.fee_token_config.config;
        let fee = fee_for_msg(dest_chain_selector, &message, dest_chain, fee_token_config)?;

        let is_paying_with_native_sol = message.fee_token == Pubkey::zeroed();
        if is_paying_with_native_sol {
            wrap_native_sol(
                &ctx.accounts.fee_token_program.to_account_info(),
                &mut ctx.accounts.authority,
                &mut ctx.accounts.fee_token_receiver,
                fee.amount,
                ctx.bumps.fee_billing_signer,
            )?;
        } else {
            let transfer = token_interface::TransferChecked {
                from: ctx
                    .accounts
                    .fee_token_user_associated_account
                    .to_account_info(),
                to: ctx.accounts.fee_token_receiver.to_account_info(),
                mint: ctx.accounts.fee_token_mint.to_account_info(),
                authority: ctx.accounts.fee_billing_signer.to_account_info(),
            };

            transfer_fee(
                fee,
                ctx.accounts.fee_token_program.to_account_info(),
                transfer,
                ctx.accounts.fee_token_mint.decimals,
                ctx.bumps.fee_billing_signer,
            )?;
        }

        let overflow_add = dest_chain.state.sequence_number.checked_add(1);
        require!(
            overflow_add.is_some(),
            CcipRouterError::ReachedMaxSequenceNumber
        );
        dest_chain.state.sequence_number = overflow_add.unwrap();

        let sender = ctx.accounts.authority.key.to_owned();
        let receiver = message.receiver.clone();
        let source_chain_selector = config.solana_chain_selector;
        let extra_args = extra_args_or_default(config, message.extra_args);

        let nonce_counter_account: &mut Account<'info, Nonce> = &mut ctx.accounts.nonce;
        let final_nonce = bump_nonce(nonce_counter_account, extra_args).unwrap();

        let token_count = message.token_amounts.len();
        require!(
            message.token_indexes.len() == token_count,
            CcipRouterError::InvalidInputs,
        );

        let mut new_message: Solana2AnyRampMessage = Solana2AnyRampMessage {
            sender,
            receiver,
            data: message.data,
            header: RampMessageHeader {
                message_id: [0; 32],
                source_chain_selector,
                dest_chain_selector,
                sequence_number: dest_chain.state.sequence_number,
                nonce: final_nonce,
            },
            extra_args,
            fee_token: message.fee_token,
            token_amounts: vec![Solana2AnyTokenTransfer::default(); token_count],
        };

        let seeds = &[EXTERNAL_TOKEN_POOL_SEED, &[ctx.bumps.token_pools_signer]];
        for (i, token_amount) in message.token_amounts.iter().enumerate() {
            require!(
                token_amount.amount != 0,
                CcipRouterError::InvalidInputsTokenAmount
            );

            // Calculate the indexes for the additional accounts of the current token index `i`
            let (start, end) = calculate_token_pool_account_indices(
                i,
                &message.token_indexes,
                ctx.remaining_accounts.len(),
            )?;

            let current_token_accounts = validate_and_parse_token_accounts(
                ctx.accounts.authority.key(),
                dest_chain_selector,
                ctx.program_id.key(),
                &ctx.remaining_accounts[start..end],
            )?;

            let router_token_pool_signer = &ctx.accounts.token_pools_signer;

            let _token_billing_config = &current_token_accounts._token_billing_config;
            // TODO: Implement charging depending on the token transfer

            // CPI: transfer token amount from user to token pool
            transfer_token(
                token_amount.amount,
                current_token_accounts.token_program,
                current_token_accounts.mint,
                current_token_accounts.user_token_account,
                current_token_accounts.pool_token_account,
                router_token_pool_signer,
                seeds,
            )?;

            // CPI: call lockOrBurn on token pool
            {
                let lock_or_burn = LockOrBurnInV1 {
                    receiver: message.receiver.clone(),
                    remote_chain_selector: dest_chain_selector,
                    original_sender: ctx.accounts.authority.key(),
                    amount: token_amount.amount,
                    local_token: token_amount.token,
                };

                let mut acc_infos = router_token_pool_signer.to_account_infos();
                acc_infos.extend_from_slice(&[
                    current_token_accounts.pool_config.to_account_info(),
                    current_token_accounts.token_program.to_account_info(),
                    current_token_accounts.mint.to_account_info(),
                    current_token_accounts.pool_signer.to_account_info(),
                    current_token_accounts.pool_token_account.to_account_info(),
                    current_token_accounts.pool_chain_config.to_account_info(),
                ]);
                acc_infos.extend_from_slice(current_token_accounts.remaining_accounts);

                let return_data = interact_with_pool(
                    current_token_accounts.pool_program.key(),
                    router_token_pool_signer.key(),
                    acc_infos,
                    lock_or_burn,
                    seeds,
                )?;

                let data = LockOrBurnOutV1::try_from_slice(&return_data)?;
                new_message.token_amounts[i] = Solana2AnyTokenTransfer {
                    source_pool_address: current_token_accounts.pool_config.key(),
                    dest_token_address: data.dest_token_address,
                    extra_data: data.dest_pool_data,
                    amount: u64_to_le_u256(token_amount.amount), // pool on receiver chain handles decimals
                    dest_exec_data: vec![0; 0],                  // TODO: Implement this
                                                                 // Destination chain specific execution data encoded in bytes
                                                                 // for an EVM destination, it consists of the amount of gas available for the releaseOrMint
                                                                 // and transfer calls made by the offRamp
                };
            }
        }

        let message_id = &new_message.hash();
        new_message.header.message_id.clone_from(message_id);

        emit!(CCIPMessageSent {
            dest_chain_selector,
            sequence_number: new_message.header.sequence_number,
            message: new_message,
        });

        Ok(())
    }

    /// OFF RAMP FLOW
    /// Commits a report to the router.
    ///
    /// The method name needs to be commit with Anchor encoding.
    ///
    /// This function is called by the OffChain when committing one Report to the Solana Router.
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
        let report_context = ReportContext::from_byte_words(report_context_byte_words);

        // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
        let config = ctx.accounts.config.load()?;

        // Check if the report contains price updates
        let empty_token_price_updates = report.price_updates.token_price_updates.is_empty();
        let empty_gas_price_updates = report.price_updates.gas_price_updates.is_empty();

        if empty_token_price_updates && empty_gas_price_updates {
            // If the report does not contain any price updates, then there is nothing to update.
            // Thus, as no price accounts have to be updated, the remaining accounts must be empty.
            require_eq!(
                ctx.remaining_accounts.len(),
                0,
                CcipRouterError::InvalidInputs
            );
        } else {
            // There are price updates in the report.
            // Remaining accounts represent:
            // - The state account to store the price sequence updates
            // - the accounts to update BillingTokenConfig for token prices
            // - the accounts to update DestChain for gas prices
            // They must be in order:
            // 1. state_account
            // 2. token_accounts[]
            // 3. gas_accounts[]
            // matching the order of the price updates in the CommitInput.
            // They must also all be writable so they can be updated.
            let minimum_remaining_accounts = 1
                + report.price_updates.token_price_updates.len()
                + report.price_updates.gas_price_updates.len();
            require_eq!(
                ctx.remaining_accounts.len(),
                minimum_remaining_accounts,
                CcipRouterError::InvalidInputs
            );

            let ocr_sequence_number = report_context.sequence_number();

            // The Global state PDA is sent as a remaining_account as it is optional to avoid having the lock when not modifying it, so all validations need to be done manually
            let (expected_state_key, _) = Pubkey::find_program_address(&[STATE_SEED], &crate::ID);
            require_keys_eq!(
                ctx.remaining_accounts[0].key(),
                expected_state_key,
                CcipRouterError::InvalidInputs
            );
            require!(
                ctx.remaining_accounts[0].is_writable,
                CcipRouterError::InvalidInputs
            );

            let mut global_state: Account<GlobalState> =
                Account::try_from(&ctx.remaining_accounts[0])?;

            if global_state.latest_price_sequence_number < ocr_sequence_number {
                // Update the persisted sequence number
                global_state.latest_price_sequence_number = ocr_sequence_number;
                global_state.exit(&crate::ID)?; // as it is manually loaded, it also has to be manually written back

                // For each token price update, unpack the corresponding remaining_account and update the price.
                // Keep in mind that the remaining_accounts are sorted in the same order as tokens and gas price updates in the report.
                for (i, update) in report.price_updates.token_price_updates.iter().enumerate() {
                    apply_token_price_update(update, &ctx.remaining_accounts[i + 1])?;
                }

                // Skip the first state account and the ones for token updates
                let offset = report.price_updates.token_price_updates.len() + 1;

                // Do the same for gas price updates
                for (i, update) in report.price_updates.gas_price_updates.iter().enumerate() {
                    apply_gas_price_update(update, &ctx.remaining_accounts[i + offset])?;
                }
            } else {
                // TODO check if this is really necessary. EVM has this validation checking that the
                // array of merkle roots in the report is not empty. But here, considering we only have 1 root per report,
                // this check is just validating that the root is not zeroed
                // (which should never happen anyway, so it may be redundant).
                require!(
                    report.merkle_root.source_chain_selector > 0,
                    CcipRouterError::StaleCommitReport
                );
            }
        }

        // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
        let source_chain_state = &mut ctx.accounts.source_chain_state;

        require!(
            source_chain_state.config.is_enabled,
            CcipRouterError::UnsupportedSourceChainSelector
        );

        // The Commit Report Account stores the information of 1 Commit Report:
        // - Merkle Root
        // - Timestamp of the Commit Report
        // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
        // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
        let commit_report = &mut ctx.accounts.commit_report;
        let root = &report.merkle_root;

        require!(
            root.min_seq_nr <= root.max_seq_nr,
            CcipRouterError::InvalidSequenceInterval
        );
        require!(
            root.max_seq_nr
                .to_owned()
                .checked_sub(root.min_seq_nr)
                .map_or_else(|| false, |seq_size| seq_size <= 64),
            CcipRouterError::InvalidSequenceInterval
        ); // As we have 64 slots to store the execution state
        require!(
            source_chain_state.state.min_seq_nr == root.min_seq_nr,
            CcipRouterError::InvalidSequenceInterval
        );
        require!(root.merkle_root != [0; 32], CcipRouterError::InvalidProof);
        require!(
            commit_report.timestamp == 0,
            CcipRouterError::ExistingMerkleRoot
        );

        let next_seq_nr = root.max_seq_nr.checked_add(1);

        require!(
            next_seq_nr.is_some(),
            CcipRouterError::ReachedMaxSequenceNumber
        );

        source_chain_state.state.min_seq_nr = next_seq_nr.unwrap();

        let clock: Clock = Clock::get()?;
        commit_report.version = 1;
        commit_report.timestamp = clock.unix_timestamp;
        commit_report.execution_states = 0;
        commit_report.min_msg_nr = root.min_seq_nr;
        commit_report.max_msg_nr = root.max_seq_nr;

        emit!(CommitReportAccepted {
            merkle_root: root.clone(),
            price_updates: report.price_updates.clone(),
        });

        config.ocr3[OcrPluginType::Commit as usize].transmit(
            &ctx.accounts.sysvar_instructions,
            ctx.accounts.authority.key(),
            OcrPluginType::Commit as u8,
            report_context,
            &report,
            &signatures,
        )?;

        Ok(())
    }

    /// OFF RAMP FLOW
    /// Executes a message on the destination chain.
    ///
    /// The method name needs to be execute with Anchor encoding.
    ///
    /// This function is called by the OffChain when executing one Report to the Solana Router.
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
    ) -> Result<()> {
        let report_context = ReportContext::from_byte_words(report_context_byte_words);
        // limit borrowing of ctx
        {
            let config = ctx.accounts.config.load()?;
            config.ocr3[OcrPluginType::Execution as usize].transmit(
                &ctx.accounts.sysvar_instructions,
                ctx.accounts.authority.key(),
                OcrPluginType::Execution as u8,
                report_context,
                &execution_report,
                &[],
            )?;
        }

        internal_execute(ctx, execution_report)
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
    ) -> Result<()> {
        // limit borrowing of ctx
        {
            let config = ctx.accounts.config.load()?;

            // validate time has passed
            let clock: Clock = Clock::get()?;
            let current_timestamp = clock.unix_timestamp;
            require!(
                current_timestamp - ctx.accounts.commit_report.timestamp
                    > config.enable_manual_execution_after,
                CcipRouterError::ManualExecutionNotAllowed
            );
        }
        internal_execute(ctx, execution_report)
    }
}

fn apply_token_price_update<'info>(
    token_update: &TokenPriceUpdate,
    token_config_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[FEE_BILLING_TOKEN_CONFIG, token_update.source_token.as_ref()],
        &crate::ID,
    );
    require_keys_eq!(
        token_config_account_info.key(),
        expected,
        CcipRouterError::InvalidInputs
    );

    require!(
        token_config_account_info.is_writable,
        CcipRouterError::InvalidInputs
    );

    let token_config_account: &mut Account<BillingTokenConfigWrapper> =
        &mut Account::try_from(token_config_account_info)?;

    require!(
        token_config_account.version == 1,
        CcipRouterError::InvalidInputs
    );

    token_config_account.config.usd_per_token = TimestampedPackedU224 {
        value: token_update.usd_per_token,
        timestamp: Clock::get()?.unix_timestamp,
    };

    emit!(UsdPerTokenUpdated {
        token: token_config_account.config.mint,
        value: token_config_account.config.usd_per_token.value,
        timestamp: token_config_account.config.usd_per_token.timestamp,
    });

    // As the account is manually loaded from the AccountInfo, it also needs to be manually
    // written back to so the changes are persisted.
    token_config_account.exit(&crate::ID)
}

fn apply_gas_price_update<'info>(
    gas_update: &GasPriceUpdate,
    dest_chain_state_account_info: &'info AccountInfo<'info>,
) -> Result<()> {
    let (expected, _) = Pubkey::find_program_address(
        &[
            DEST_CHAIN_STATE_SEED,
            gas_update.dest_chain_selector.to_le_bytes().as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(
        dest_chain_state_account_info.key(),
        expected,
        CcipRouterError::InvalidInputs
    );

    require!(
        dest_chain_state_account_info.is_writable,
        CcipRouterError::InvalidInputs
    );

    // The passed-in chain_state account may refer to the same chain but it only corresponds to source.
    // To update the price that values correspond to the destination, which is a different account.
    // As the account is sent as additional accounts, then Anchor won't automatically (de)serialize the account
    // as it is not the one in the context, so we have to do it manually load it and write it back
    let dest_chain_state_account = &mut Account::try_from(dest_chain_state_account_info)?;
    update_chain_state_gas_price(dest_chain_state_account, gas_update)?;
    dest_chain_state_account.exit(&crate::ID)?;
    Ok(())
}

fn update_chain_state_gas_price(
    chain_state_account: &mut Account<DestChain>,
    gas_update: &GasPriceUpdate,
) -> Result<()> {
    require!(
        chain_state_account.version == 1,
        CcipRouterError::InvalidInputs
    );

    chain_state_account.state.usd_per_unit_gas = TimestampedPackedU224 {
        value: gas_update.usd_per_unit_gas,
        timestamp: Clock::get()?.unix_timestamp,
    };

    emit!(UsdPerUnitGasUpdated {
        dest_chain: gas_update.dest_chain_selector,
        value: chain_state_account.state.usd_per_unit_gas.value,
        timestamp: chain_state_account.state.usd_per_unit_gas.timestamp,
    });

    Ok(())
}

// internal_execute is the base execution logic without any additional validation
fn internal_execute<'info>(
    ctx: Context<'_, '_, 'info, 'info, ExecuteReportContext<'info>>,
    execution_report: ExecutionReportSingleChain,
) -> Result<()> {
    // TODO: Limit send size data to 256

    // The Config Account stores the default values for the Router, the Solana Chain Selector, the Default Gas Limit and the Default Allow Out Of Order Execution and Admin Ownership
    let config = ctx.accounts.config.load()?;
    let solana_chain_selector = config.solana_chain_selector;

    // The Config and State for the Source Chain, containing if it is enabled, the on ramp address and the min sequence number expected for future messages
    let source_chain_state = &ctx.accounts.source_chain_state;
    let on_ramp_address = &source_chain_state.config.on_ramp;

    // The Commit Report Account stores the information of 1 Commit Report:
    // - Merkle Root
    // - Timestamp of the Commit Report
    // - Interval of Messages: The min and max seq num of the messages in the Merkle Tree
    // - Execution State per each Message: 0 for Untouched, 1 for InProgress, 2 for Success and 3 for Failure
    let commit_report = &mut ctx.accounts.commit_report;

    let message_header = execution_report.message.header;

    validate_execution_report(
        &execution_report,
        source_chain_state,
        commit_report,
        &message_header,
        solana_chain_selector,
    )?;

    let original_state = commit_report.get_state(message_header.sequence_number);

    if original_state == MessageExecutionState::Success {
        emit!(SkippedAlreadyExecutedMessage {
            source_chain_selector: message_header.source_chain_selector,
            sequence_number: message_header.sequence_number,
        });
        return Ok(());
    }

    let hashed_leaf = verify_merkle_root(&execution_report, on_ramp_address)?;

    // send tokens any -> SOL
    require!(
        execution_report.token_indexes.len() == execution_report.message.token_amounts.len()
            && execution_report.token_indexes.len() == execution_report.offchain_token_data.len(),
        CcipRouterError::InvalidInputs,
    );
    let seeds = &[EXTERNAL_TOKEN_POOL_SEED, &[ctx.bumps.token_pools_signer]];
    let mut token_amounts =
        vec![SolanaTokenAmount::default(); execution_report.token_indexes.len()];

    // handle tokens
    // note: indexes are used instead of counts in case more accounts need to be passed in remaining_accounts before token accounts
    // token_indexes = [2, 4] where remaining_accounts is [custom_account, custom_account, token1_account1, token1_account2, token2_account1, token2_account2] for example
    for (i, token_amount) in execution_report.message.token_amounts.iter().enumerate() {
        let (start, end) = calculate_token_pool_account_indices(
            i,
            &execution_report.token_indexes,
            ctx.remaining_accounts.len(),
        )?;
        let acc_list = &ctx.remaining_accounts[start..end];
        let accs = validate_and_parse_token_accounts(
            execution_report.message.receiver,
            execution_report.message.header.source_chain_selector,
            ctx.program_id.key(),
            acc_list,
        )?;
        let router_token_pool_signer = &ctx.accounts.token_pools_signer;

        let init_bal = get_balance(accs.user_token_account)?;

        // CPI: call lockOrBurn on token pool
        let release_or_mint = ReleaseOrMintInV1 {
            original_sender: execution_report.message.sender.clone(),
            receiver: execution_report.message.receiver,
            amount: token_amount.amount,
            local_token: token_amount.dest_token_address,
            remote_chain_selector: execution_report.message.header.source_chain_selector,
            source_pool_address: token_amount.source_pool_address.clone(),
            source_pool_data: token_amount.extra_data.clone(),
            offchain_token_data: execution_report.offchain_token_data[i].clone(),
        };
        let mut acc_infos = router_token_pool_signer.to_account_infos();
        acc_infos.extend_from_slice(&[
            accs.pool_config.to_account_info(),
            accs.token_program.to_account_info(),
            accs.mint.to_account_info(),
            accs.pool_signer.to_account_info(),
            accs.pool_token_account.to_account_info(),
            accs.pool_chain_config.to_account_info(),
            accs.user_token_account.to_account_info(),
        ]);
        acc_infos.extend_from_slice(accs.remaining_accounts);
        let return_data = interact_with_pool(
            accs.pool_program.key(),
            router_token_pool_signer.key(),
            acc_infos,
            release_or_mint,
            seeds,
        )?;

        require!(
            return_data.len() == CCIP_POOL_V1_RET_BYTES,
            CcipRouterError::OfframpInvalidDataLength
        );

        token_amounts[i] = SolanaTokenAmount {
            token: accs.mint.key(),
            amount: ReleaseOrMintOutV1::try_from_slice(&return_data)?.destination_amount,
        };

        let post_bal = get_balance(accs.user_token_account)?;
        require!(
            post_bal >= init_bal && post_bal - init_bal == token_amounts[i].amount,
            CcipRouterError::OfframpReleaseMintBalanceMismatch
        );
    }

    let message = Any2SolanaMessage {
        message_id: execution_report.message.header.message_id,
        source_chain_selector: execution_report.source_chain_selector,
        sender: execution_report.message.sender,
        data: execution_report.message.data,
        token_amounts,
    };

    // handle CPI call if there are extra accounts
    // case: no tokens, but there are remaining_accounts passed in
    // case: tokens, but the first token has a non-zero index (indicating extra accounts before token accounts)
    if should_execute_messaging(
        &execution_report.token_indexes,
        ctx.remaining_accounts.is_empty(),
    ) {
        let (msg_program, msg_accounts) = parse_messaging_accounts(
            &execution_report.token_indexes,
            execution_report.message.receiver,
            execution_report.message.extra_args.accounts,
            ctx.remaining_accounts,
        )?;

        // The External Execution Config Account is used to sign the CPI instruction
        let external_execution_config = &ctx.accounts.external_execution_config;

        // The accounts of the user that will be used in the CPI instruction, none of them are signers
        // They need to specify if mutable or not, but none of them is allowed to init, realloc or close
        // note: CPI signer is always first account
        let mut acc_infos = external_execution_config.to_account_infos();
        acc_infos.extend_from_slice(msg_accounts);

        let acc_metas: Vec<AccountMeta> = acc_infos
            .to_vec()
            .iter()
            .flat_map(|acc_info| {
                // Check signer from PDA External Execution config
                let is_signer = acc_info.key() == external_execution_config.key();
                acc_info.to_account_metas(Some(is_signer))
            })
            .collect();

        let data = build_receiver_discriminator_and_data(message)?;

        let instruction = Instruction {
            program_id: msg_program.key(), // The receiver Program Id that will handle the ccip_receive message
            accounts: acc_metas,
            data,
        };

        let seeds = &[
            EXTERNAL_EXECUTION_CONFIG_SEED,
            &[ctx.bumps.external_execution_config],
        ];
        let signer = &[&seeds[..]];

        invoke_signed(&instruction, &acc_infos, signer)?;
    }

    let new_state = MessageExecutionState::Success;
    commit_report.set_state(message_header.sequence_number, new_state.to_owned());

    emit!(ExecutionStateChanged {
        source_chain_selector: message_header.source_chain_selector,
        sequence_number: message_header.sequence_number,
        message_id: message_header.message_id, // Unique identifier for the message, generated with the source chain's encoding scheme
        message_hash: hashed_leaf,             // Hash of the message using Solana encoding
        state: new_state,
    });

    Ok(())
}

// should_execute_messaging checks if there remaining_accounts that are not being used for token pools
// case: no tokens, but there are remaining_accounts passed in
// case: tokens, but the first token has a non-zero index (indicating extra accounts before token accounts)
fn should_execute_messaging(token_indexes: &[u8], remaining_accounts_empty: bool) -> bool {
    (token_indexes.is_empty() && !remaining_accounts_empty)
        || (!token_indexes.is_empty() && token_indexes[0] != 0)
}

/// parse_message_accounts returns all the accounts needed to execute the CPI instruction
/// It also validates that the accounts sent in the message match the ones sent in the source chain
///
/// # Arguments
/// * `token_indexes` - start indexes of token pool accounts, used to determine ending index for arbitrary messaging accounts
/// * `receiver` - receiver address from x-chain message, used to validate `accounts`
/// * `source_accounts` - arbitrary messaging accounts from the x-chain message, used to validate `accounts`. expected order is: [program, ...additional message accounts]
/// * `accounts` - accounts passed via `ctx.remaining_accounts`. expected order is: [program, receiver, ...additional message accounts]
fn parse_messaging_accounts<'info>(
    token_indexes: &[u8],
    receiver: Pubkey,
    source_accounts: Vec<SolanaAccountMeta>,
    accounts: &'info [AccountInfo<'info>],
) -> Result<(&'info AccountInfo<'info>, &'info [AccountInfo<'info>])> {
    let end_ind = if token_indexes.is_empty() {
        accounts.len()
    } else {
        token_indexes[0] as usize
    };

    let msg_program = &accounts[0];
    let msg_accounts = &accounts[1..end_ind];

    let source_program = &source_accounts[0];
    let source_msg_accounts = &source_accounts[1..source_accounts.len()];

    require!(
        source_program.pubkey == msg_program.key(),
        CcipRouterError::InvalidInputs,
    );

    require!(
        msg_accounts[0].key() == receiver,
        CcipRouterError::InvalidInputs
    );

    // assert same number of accounts passed from message and transaction (not including program)
    // source_msg_accounts + 1 to account for separately passed receiver address
    require!(
        source_msg_accounts.len() + 1 == msg_accounts.len(),
        CcipRouterError::InvalidInputs
    );

    // Validate the addresses of all the accounts match the ones in source chain
    if msg_accounts.len() > 1 {
        // Ignore the first account as it's the receiver
        let accounts_to_validate = &msg_accounts[1..msg_accounts.len()];
        require!(
            accounts_to_validate.len() == source_msg_accounts.len(),
            CcipRouterError::InvalidInputs
        );
        for (i, acc) in source_msg_accounts.iter().enumerate() {
            let current_acc = &msg_accounts[i + 1];
            require!(
                acc.pubkey == current_acc.key(),
                CcipRouterError::InvalidInputs
            );
            require!(
                acc.is_writable == current_acc.is_writable,
                CcipRouterError::InvalidInputs
            );
        }
    }

    Ok((msg_program, msg_accounts))
}

fn extra_args_or_default(default_config: Ref<Config>, extra_args: ExtraArgsInput) -> AnyExtraArgs {
    let mut result_args = AnyExtraArgs {
        gas_limit: default_config.default_gas_limit.to_owned(),
        allow_out_of_order_execution: default_config.default_allow_out_of_order_execution != 0,
    };

    if let Some(gas_limit) = extra_args.gas_limit {
        gas_limit.clone_into(&mut result_args.gas_limit)
    }

    if let Some(allow_out_of_order_execution) = extra_args.allow_out_of_order_execution {
        allow_out_of_order_execution.clone_into(&mut result_args.allow_out_of_order_execution)
    }

    result_args
}

/// Build the instruction data (discriminator + any other data)
fn build_receiver_discriminator_and_data(ramp_message: Any2SolanaMessage) -> Result<Vec<u8>> {
    let m: std::result::Result<Vec<u8>, std::io::Error> = ramp_message.try_to_vec();
    require!(m.is_ok(), CcipRouterError::InvalidMessage);
    let message = m.unwrap();

    let mut data = Vec::with_capacity(8);
    data.extend_from_slice(&CCIP_RECEIVE_DISCRIMINATOR);
    data.extend_from_slice(&message);

    Ok(data)
}

fn validate_source_chain_config(
    _source_chain_selector: u64,
    _config: &SourceChainConfig,
) -> Result<()> {
    // As of now, the config has very few properties and there is nothing to validate yet.
    // This is a placeholder to add validations as that config object grows.
    Ok(())
}

fn validate_dest_chain_config(dest_chain_selector: u64, config: &DestChainConfig) -> Result<()> {
    // TODO improve errors
    require!(dest_chain_selector != 0, CcipRouterError::InvalidInputs);
    require!(
        config.default_tx_gas_limit != 0,
        CcipRouterError::InvalidInputs
    );
    require!(
        config.default_tx_gas_limit <= config.max_per_msg_gas_limit,
        CcipRouterError::InvalidInputs
    );
    require!(
        config.chain_family_selector != [0; 4],
        CcipRouterError::InvalidInputs
    );
    Ok(())
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
}

// TODO: Refactor this to use the same structure as messages: execution_report.validate(..)
pub fn validate_execution_report<'info>(
    execution_report: &ExecutionReportSingleChain,
    source_chain_state: &Account<'info, SourceChain>,
    commit_report: &Account<'info, CommitReport>,
    message_header: &RampMessageHeader,
    solana_chain_selector: u64,
) -> Result<()> {
    require!(
        execution_report.message.header.nonce == 0,
        CcipRouterError::InvalidInputs
    );

    require!(
        source_chain_state.config.is_enabled,
        CcipRouterError::UnsupportedSourceChainSelector
    );

    require!(
        execution_report.message.header.sequence_number >= commit_report.min_msg_nr
            && execution_report.message.header.sequence_number <= commit_report.max_msg_nr,
        CcipRouterError::InvalidSequenceInterval
    );

    require!(
        message_header.source_chain_selector == execution_report.source_chain_selector,
        CcipRouterError::UnsupportedSourceChainSelector
    );
    require!(
        message_header.dest_chain_selector == solana_chain_selector,
        CcipRouterError::UnsupportedDestinationChainSelector
    );
    require!(
        commit_report.timestamp != 0,
        CcipRouterError::RootNotCommitted
    );

    Ok(())
}

pub fn verify_merkle_root(
    execution_report: &ExecutionReportSingleChain,
    on_ramp_address: &[u8],
) -> Result<[u8; 32]> {
    let hashed_leaf = execution_report.message.hash(on_ramp_address);
    let verified_root: std::result::Result<[u8; 32], MerkleError> =
        calculate_merkle_root(hashed_leaf, execution_report.proofs.clone());
    require!(
        verified_root.is_ok() && verified_root.unwrap() == execution_report.root,
        CcipRouterError::InvalidProof
    );
    Ok(hashed_leaf)
}

fn bump_nonce(nonce_counter_account: &mut Account<Nonce>, extra_args: AnyExtraArgs) -> Result<u64> {
    // Avoid Re-initialization attack as the account is init_if_needed
    // https://solana.com/developers/courses/program-security/reinitialization-attacks#add-is-initialized-check
    if nonce_counter_account.version == 0 {
        // The authority must be the owner of the PDA, as it's their Public Key in the seed
        // If the account is not initialized, initialize it
        nonce_counter_account.version = 1;
        nonce_counter_account.counter = 0;
    }

    // TODO: Require config.enforce_out_of_order => extra_args.allow_out_of_order_execution
    let mut final_nonce = 0;
    if !extra_args.allow_out_of_order_execution {
        nonce_counter_account.counter = nonce_counter_account.counter.checked_add(1).unwrap();
        final_nonce = nonce_counter_account.counter;
    }
    Ok(final_nonce)
}
