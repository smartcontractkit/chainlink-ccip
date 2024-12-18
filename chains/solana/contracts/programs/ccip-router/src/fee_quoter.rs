use anchor_lang::prelude::*;
use anchor_spl::token_interface;
use bytemuck::Zeroable;
use spl_token_2022::native_mint;
use token_interface::spl_token_2022;

use crate::{
    BillingTokenConfig, CcipRouterError, DestChain, Solana2AnyMessage, SolanaTokenAmount,
    FEE_BILLING_SIGNER_SEEDS,
};

// TODO change args and implement
pub fn fee_for_msg(
    _dest_chain_selector: u64,
    message: Solana2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
) -> Result<SolanaTokenAmount> {
    let token = if message.fee_token == Pubkey::zeroed() {
        native_mint::ID // Wrapped SOL
    } else {
        message.fee_token
    };

    require!(
        dest_chain.config.is_enabled,
        CcipRouterError::DestinationChainDisabled
    );
    require!(token_config.enabled, CcipRouterError::FeeTokenDisabled);

    // TODO validate that dest_chain_selector is whitelisted and not cursed

    //? /Users/tobi/dev/work/ccip/contracts/src/v0.8/ccip/FeeQuoter.sol:518
    //? 1. Get dest chain config
    //? 1. validate message
    //? 1. retrieve fee token's current price
    //? 1. retrieve gas price on the dest chain
    //? 1. get cost of transfer _of each token in the message_ to the dest chain, including premium, gas, and bytes overhead
    //? 1. get data availability cost on the dest chain
    //? 1. Calculate and return the total fee

    // TODO un-hardcode
    Ok(SolanaTokenAmount { amount: 1, token })
}

pub fn transfer_fee<'info>(
    fee: SolanaTokenAmount,
    token_program: AccountInfo<'info>,
    transfer: token_interface::TransferChecked<'info>,
    decimals: u8,
    signer_bump: u8,
) -> Result<()> {
    require!(
        fee.token == transfer.mint.key(),
        CcipRouterError::InvalidInputs
    ); // TODO use more specific error

    let seeds = &[FEE_BILLING_SIGNER_SEEDS, &[signer_bump]];
    let signer_seeds = &[&seeds[..]];
    let cpi_ctx = CpiContext::new_with_signer(token_program, transfer, signer_seeds);
    token_interface::transfer_checked(cpi_ctx, fee.amount, decimals)
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::DestChain;

    #[test]
    fn fees_not_returned_for_disabled_destination_chain() {
        let mut chain = sample_dest_chain();
        chain.config.is_enabled = false;

        assert!(fee_for_msg(0, sample_message(), &chain, &sample_billing_config()).is_err());
    }

    #[test]
    fn fees_not_returned_for_disabled_token() {
        let mut billing_config = sample_billing_config();
        billing_config.enabled = false;

        assert!(fee_for_msg(0, sample_message(), &sample_dest_chain(), &billing_config).is_err());
    }

    fn sample_message() -> Solana2AnyMessage {
        Solana2AnyMessage {
            receiver: vec![],
            data: vec![],
            token_amounts: vec![],
            fee_token: Pubkey::new_unique(),
            extra_args: crate::ExtraArgsInput {
                gas_limit: None,
                allow_out_of_order_execution: None,
            },
            token_indexes: vec![],
        }
    }

    fn sample_billing_config() -> BillingTokenConfig {
        BillingTokenConfig {
            enabled: true,
            mint: Pubkey::new_unique(),
            usd_per_token: crate::TimestampedPackedU224 {
                value: [0; 28],
                timestamp: 0,
            },
            premium_multiplier_wei_per_eth: 0,
        }
    }

    fn sample_dest_chain() -> DestChain {
        DestChain {
            state: crate::DestChainState {
                sequence_number: 0,
                usd_per_unit_gas: crate::TimestampedPackedU224 {
                    value: [0; 28],
                    timestamp: 0,
                },
            },
            config: crate::DestChainConfig {
                is_enabled: true,
                max_number_of_tokens_per_msg: 0,
                max_data_bytes: 0,
                max_per_msg_gas_limit: 0,
                dest_gas_overhead: 0,
                dest_gas_per_payload_byte: 0,
                dest_data_availability_overhead_gas: 0,
                dest_gas_per_data_availability_byte: 0,
                dest_data_availability_multiplier_bps: 0,
                default_token_fee_usdcents: 0,
                default_token_dest_gas_overhead: 0,
                default_tx_gas_limit: 0,
                gas_multiplier_wei_per_eth: 0,
                network_fee_usdcents: 0,
                gas_price_staleness_threshold: 0,
                enforce_out_of_order: false,
                chain_family_selector: [0; 4],
            },
        }
    }
}
