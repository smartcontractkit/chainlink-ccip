use anchor_lang::prelude::*;

pub mod events;
pub mod state;

use events::*;
use state::*;

declare_id!("9ykZ4KUXJUtACe5Cg3UuTM14t5bxk1Amf6uaawGpGR5d");

#[program]
pub mod fee_quoter_dummy_emitter {
    use super::*;

    /// Emit events gracefully. Will mark transaction as succeeded.
    pub fn trigger_all_events(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        // 1. Emit ConfigSet
        emit!(ConfigSet {
            max_fee_juels_per_msg: 1000000,
            link_token_mint: Pubkey::default(),
            onramp: Pubkey::default(),
            default_code_version: CodeVersion::V1,
        });

        // 2. Emit FeeToken events
        let fee_token = Pubkey::default();
        emit!(FeeTokenAdded {
            fee_token,
            enabled: true,
        });
        emit!(FeeTokenEnabled { fee_token });
        emit!(FeeTokenDisabled { fee_token });
        emit!(FeeTokenRemoved { fee_token });

        // 3. Emit DestChain events
        let dest_chain_config = DestChainConfig {
            is_enabled: true,
            lane_code_version: CodeVersion::V1,
            max_number_of_tokens_per_msg: 5,
            max_data_bytes: 1000,
            max_per_msg_gas_limit: 200000,
            dest_gas_overhead: 50000,
            dest_gas_per_payload_byte_base: 16,
            dest_gas_per_payload_byte_high: 32,
            dest_gas_per_payload_byte_threshold: 1000,
            dest_data_availability_overhead_gas: 50000,
            dest_gas_per_data_availability_byte: 16,
            dest_data_availability_multiplier_bps: 10000,
            default_token_fee_usdcents: 100,
            default_token_dest_gas_overhead: 50000,
            default_tx_gas_limit: 200000,
            gas_multiplier_wei_per_eth: 110000000000000000,
            network_fee_usdcents: 100,
            gas_price_staleness_threshold: 3600,
            enforce_out_of_order: false,
            chain_family_selector: [0u8; 4],
        };

        emit!(DestChainAdded {
            dest_chain_selector: 1,
            dest_chain_config: dest_chain_config.clone(),
        });

        emit!(DestChainConfigUpdated {
            dest_chain_selector: 1,
            dest_chain_config,
        });

        // 4. Emit Ownership events
        let owner = Pubkey::default();
        let new_owner = Pubkey::default();
        emit!(OwnershipTransferRequested {
            from: owner,
            to: new_owner,
        });
        emit!(OwnershipTransferred {
            from: owner,
            to: new_owner,
        });

        // 5. Emit Price Update events
        emit!(UsdPerUnitGasUpdated {
            dest_chain: 1,
            value: [0u8; 28],
            timestamp: 1677777777,
        });

        emit!(UsdPerTokenUpdated {
            token: Pubkey::default(),
            value: [0u8; 28],
            timestamp: 1677777777,
        });

        // 6. Emit Token Transfer Fee Config events
        emit!(TokenTransferFeeConfigUpdated {
            dest_chain_selector: 1,
            token: Pubkey::default(),
            token_transfer_fee_config: TokenTransferFeeConfig {
                min_fee_usdcents: 100,
                max_fee_usdcents: 1000,
                deci_bps: 1000,
                dest_gas_overhead: 50000,
                dest_bytes_overhead: 1000,
                is_enabled: true,
            },
        });

        emit!(PremiumMultiplierWeiPerEthUpdated {
            token: Pubkey::default(),
            premium_multiplier_wei_per_eth: 110000000000000000,
        });

        // 7. Emit Price Updater events
        let price_updater = Pubkey::default();
        emit!(PriceUpdaterAdded { price_updater });
        emit!(PriceUpdaterRemoved { price_updater });

        msg!("All events have been emitted!");
        Ok(())
    }

    /// Emit events and reverts. Will mark transaction as failed.
    pub fn trigger_all_events_reverts(_ctx: Context<TriggerAllEvents>) -> Result<()> {
        trigger_all_events(_ctx)?;
        Err(error!(CustomError::IntentionalRevert))
    }
}

#[derive(Accounts)]
pub struct TriggerAllEvents {}

#[error_code]
pub enum CustomError {
    #[msg("Transaction intentionally reverted for testing")]
    IntentionalRevert,
}