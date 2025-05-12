use anchor_lang::prelude::*;

use ccip_common::{CommonCcipError, CHAIN_FAMILY_SELECTOR_EVM, CHAIN_FAMILY_SELECTOR_SVM};

use crate::context::{
    AcceptOwnership, AddBillingTokenConfig, AddDestChain, AddPriceUpdater, RemovePriceUpdater,
    SetTokenTransferFeeConfig, UpdateBillingTokenConfig, UpdateConfig, UpdateConfigLinkMint,
    UpdateDestChainConfig,
};
use crate::event::{
    ConfigSet, DestChainAdded, DestChainConfigUpdated, FeeTokenAdded, FeeTokenDisabled,
    FeeTokenEnabled, OwnershipTransferRequested, OwnershipTransferred,
    PremiumMultiplierWeiPerEthUpdated, PriceUpdaterAdded, PriceUpdaterRemoved,
    TokenTransferFeeConfigUpdated, UsdPerTokenUpdated,
};
use crate::instructions::interfaces::Admin;
use crate::instructions::v1::public::CCIP_LOCK_OR_BURN_V1_RET_BYTES;
use crate::state::{
    BillingTokenConfig, CodeVersion, DestChain, DestChainConfig, DestChainState,
    PerChainPerTokenConfig, TimestampedPackedU224, TokenTransferFeeConfig,
};
use crate::FeeQuoterError;

pub struct Impl;
impl Admin for Impl {
    fn transfer_ownership(&self, ctx: Context<UpdateConfig>, proposed_owner: Pubkey) -> Result<()> {
        let config = &mut ctx.accounts.config;
        require!(
            proposed_owner != config.owner,
            FeeQuoterError::RedundantOwnerProposal
        );
        emit!(OwnershipTransferRequested {
            from: config.owner,
            to: proposed_owner,
        });
        ctx.accounts.config.proposed_owner = proposed_owner;
        Ok(())
    }

    fn accept_ownership(&self, ctx: Context<AcceptOwnership>) -> Result<()> {
        let config = &mut ctx.accounts.config;
        emit!(OwnershipTransferred {
            from: config.owner,
            to: config.proposed_owner,
        });
        // NOTE: take() resets proposed_owner to default
        ctx.accounts.config.owner = std::mem::take(&mut ctx.accounts.config.proposed_owner);
        Ok(())
    }

    fn set_default_code_version(
        &self,
        ctx: Context<UpdateConfig>,
        code_version: CodeVersion,
    ) -> Result<()> {
        require_neq!(
            code_version,
            CodeVersion::Default,
            FeeQuoterError::InvalidCodeVersion
        );
        let config = &mut ctx.accounts.config;
        config.default_code_version = code_version;

        emit!(ConfigSet {
            max_fee_juels_per_msg: config.max_fee_juels_per_msg,
            link_token_mint: config.link_token_mint,
            link_token_local_decimals: config.link_token_local_decimals,
            onramp: config.onramp,
            default_code_version: config.default_code_version
        });
        Ok(())
    }

    fn set_max_fee_juels_per_msg(
        &self,
        ctx: Context<UpdateConfig>,
        max_fee_juels_per_msg: u128,
    ) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.max_fee_juels_per_msg = max_fee_juels_per_msg;

        emit!(ConfigSet {
            max_fee_juels_per_msg: config.max_fee_juels_per_msg,
            link_token_mint: config.link_token_mint,
            link_token_local_decimals: config.link_token_local_decimals,
            onramp: config.onramp,
            default_code_version: config.default_code_version
        });
        Ok(())
    }

    fn set_link_token_mint(&self, ctx: Context<UpdateConfigLinkMint>) -> Result<()> {
        let config = &mut ctx.accounts.config;
        config.link_token_mint = ctx.accounts.link_token_mint.key();
        config.link_token_local_decimals = ctx.accounts.link_token_mint.decimals;

        emit!(ConfigSet {
            max_fee_juels_per_msg: config.max_fee_juels_per_msg,
            link_token_mint: config.link_token_mint,
            link_token_local_decimals: config.link_token_local_decimals,
            onramp: config.onramp,
            default_code_version: config.default_code_version
        });
        Ok(())
    }

    fn add_billing_token_config(
        &self,
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

    fn update_billing_token_config(
        &self,
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

        // emit an event if the premium multiplier has changed, before updating the config
        if config.premium_multiplier_wei_per_eth
            != ctx
                .accounts
                .billing_token_config
                .config
                .premium_multiplier_wei_per_eth
        {
            emit!(PremiumMultiplierWeiPerEthUpdated {
                token: config.mint,
                premium_multiplier_wei_per_eth: config.premium_multiplier_wei_per_eth,
            });
        }

        if config.usd_per_token != ctx.accounts.billing_token_config.config.usd_per_token {
            emit!(UsdPerTokenUpdated {
                token: config.mint,
                value: config.usd_per_token.value,
                timestamp: config.usd_per_token.timestamp,
            });
        }

        ctx.accounts.billing_token_config.version = 1; // update this if we change the account struct
        ctx.accounts.billing_token_config.config = config;

        Ok(())
    }

    fn add_dest_chain(
        &self,
        ctx: Context<AddDestChain>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        validate_dest_chain_config(chain_selector, &dest_chain_config)?;
        ctx.accounts.dest_chain.set_inner(DestChain {
            version: 1,
            chain_selector,
            state: DestChainState {
                usd_per_unit_gas: TimestampedPackedU224 {
                    timestamp: 0,
                    value: [0; 28],
                },
            },
            config: dest_chain_config.clone(),
        });
        emit!(DestChainAdded {
            dest_chain_selector: chain_selector,
            dest_chain_config,
        });
        Ok(())
    }

    fn disable_dest_chain(
        &self,
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
    ) -> Result<()> {
        ctx.accounts.dest_chain.config.is_enabled = false;
        emit!(DestChainConfigUpdated {
            dest_chain_selector: chain_selector,
            dest_chain_config: ctx.accounts.dest_chain.config.clone(),
        });
        Ok(())
    }

    fn update_dest_chain_config(
        &self,
        ctx: Context<UpdateDestChainConfig>,
        chain_selector: u64,
        dest_chain_config: DestChainConfig,
    ) -> Result<()> {
        validate_dest_chain_config(chain_selector, &dest_chain_config)?;
        ctx.accounts.dest_chain.config = dest_chain_config.clone();
        emit!(DestChainConfigUpdated {
            dest_chain_selector: chain_selector,
            dest_chain_config,
        });
        Ok(())
    }

    fn add_price_updater(
        &self,
        _ctx: Context<AddPriceUpdater>,
        price_updater: Pubkey,
    ) -> Result<()> {
        emit!(PriceUpdaterAdded { price_updater });
        Ok(())
    }

    fn remove_price_updater(
        &self,
        _ctx: Context<RemovePriceUpdater>,
        price_updater: Pubkey,
    ) -> Result<()> {
        emit!(PriceUpdaterRemoved { price_updater });
        Ok(())
    }

    fn set_token_transfer_fee_config(
        &self,
        ctx: Context<SetTokenTransferFeeConfig>,
        chain_selector: u64,
        mint: Pubkey,
        cfg: TokenTransferFeeConfig,
    ) -> Result<()> {
        require!(
            cfg.max_fee_usdcents > cfg.min_fee_usdcents,
            FeeQuoterError::InvalidTokenTransferFeeMaxMin
        );

        require_gte!(
            cfg.dest_bytes_overhead,
            CCIP_LOCK_OR_BURN_V1_RET_BYTES,
            FeeQuoterError::InvalidTokenTransferFeeDestBytesOverhead
        );

        ctx.accounts
            .per_chain_per_token_config
            .set_inner(PerChainPerTokenConfig {
                version: 1, // update this if we change the account struct
                chain_selector,
                mint,
                token_transfer_config: cfg.clone(),
            });
        emit!(TokenTransferFeeConfigUpdated {
            dest_chain_selector: chain_selector,
            token: mint,
            token_transfer_fee_config: cfg,
        });
        Ok(())
    }
}

// --- helpers ---

fn validate_dest_chain_config(dest_chain_selector: u64, config: &DestChainConfig) -> Result<()> {
    // check if the lane code version is supported
    require!(
        matches!(
            config.lane_code_version,
            CodeVersion::Default | CodeVersion::V1
        ),
        FeeQuoterError::InvalidVersion
    );
    require!(
        dest_chain_selector != 0,
        FeeQuoterError::InvalidInputsChainSelector
    );
    require!(
        config.default_tx_gas_limit != 0,
        FeeQuoterError::ZeroGasLimit
    );
    require!(
        config.default_tx_gas_limit <= config.max_per_msg_gas_limit,
        FeeQuoterError::DefaultGasLimitExceedsMaximum
    );

    require!(
        matches!(
            u32::from_be_bytes(config.chain_family_selector),
            CHAIN_FAMILY_SELECTOR_EVM | CHAIN_FAMILY_SELECTOR_SVM
        ),
        CommonCcipError::InvalidChainFamilySelector
    );

    Ok(())
}
