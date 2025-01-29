use anchor_lang::prelude::*;
use ethnum::U256;

use crate::messages::{SVM2AnyMessage, CHAIN_FAMILY_SELECTOR_EVM};
use crate::state::{BillingTokenConfig, DestChain};
use crate::FeeQuoterError;

const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);

pub fn validate_svm2any(
    msg: &SVM2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
) -> Result<()> {
    require!(
        dest_chain.config.is_enabled,
        FeeQuoterError::DestinationChainDisabled
    );

    require!(token_config.enabled, FeeQuoterError::FeeTokenDisabled);

    require_gte!(
        dest_chain.config.max_data_bytes,
        msg.data.len() as u32,
        FeeQuoterError::MessageTooLarge
    );

    require_gte!(
        dest_chain.config.max_number_of_tokens_per_msg as usize,
        msg.token_amounts.len(),
        FeeQuoterError::UnsupportedNumberOfTokens
    );

    require_gte!(
        dest_chain.config.max_per_msg_gas_limit as u128,
        msg.extra_args
            .gas_limit
            .unwrap_or(dest_chain.config.default_tx_gas_limit as u128),
        FeeQuoterError::MessageGasLimitTooHigh,
    );

    require!(
        !dest_chain.config.enforce_out_of_order
            || msg.extra_args.allow_out_of_order_execution.unwrap_or(false),
        FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue,
    );

    validate_dest_family_address(msg, dest_chain.config.chain_family_selector)
}

fn validate_dest_family_address(
    msg: &SVM2AnyMessage,
    chain_family_selector: [u8; 4],
) -> Result<()> {
    const PRECOMPILE_SPACE: u32 = 1024;

    let selector = u32::from_be_bytes(chain_family_selector);
    // Only EVM is supported as a destination family.
    require_eq!(
        selector,
        CHAIN_FAMILY_SELECTOR_EVM,
        FeeQuoterError::UnsupportedChainFamilySelector
    );

    require_eq!(msg.receiver.len(), 32, FeeQuoterError::InvalidEVMAddress);

    let address: U256 = U256::from_be_bytes(
        msg.receiver
            .clone()
            .try_into()
            .map_err(|_| FeeQuoterError::InvalidEncoding)?,
    );

    require!(address <= U160_MAX, FeeQuoterError::InvalidEVMAddress);

    if let Ok(small_address) = TryInto::<u32>::try_into(address) {
        require_gte!(
            small_address,
            PRECOMPILE_SPACE,
            FeeQuoterError::InvalidEVMAddress
        )
    };

    Ok(())
}

#[cfg(test)]
pub mod tests {
    use super::*;
    use crate::{ExtraArgsInput, SVMTokenAmount, TimestampedPackedU224};
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;

    fn as_u8_28(single: U256) -> [u8; 28] {
        single.to_be_bytes()[4..32].try_into().unwrap()
    }

    #[test]
    fn message_not_validated_for_disabled_destination_chain() {
        let mut chain = sample_dest_chain();
        chain.config.is_enabled = false;

        assert_eq!(
            validate_svm2any(&sample_message(), &chain, &sample_billing_config()).unwrap_err(),
            FeeQuoterError::DestinationChainDisabled.into()
        );
    }

    #[test]
    fn message_not_validated_for_disabled_token() {
        let mut billing_config = sample_billing_config();
        billing_config.enabled = false;

        assert_eq!(
            validate_svm2any(&sample_message(), &sample_dest_chain(), &billing_config).unwrap_err(),
            FeeQuoterError::FeeTokenDisabled.into()
        );
    }

    #[test]
    fn large_message_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.data = vec![0; dest_chain.config.max_data_bytes as usize + 1];
        assert_eq!(
            validate_svm2any(&message, &sample_dest_chain(), &sample_billing_config()).unwrap_err(),
            FeeQuoterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn invalid_addresses_fail_to_validate() {
        let mut address_bigger_than_u160_max = vec![0u8; 32];
        address_bigger_than_u160_max[11] = 1;
        let mut address_in_precompile_space = vec![0u8; 32];
        address_in_precompile_space[30] = 1;
        let incorrect_length_address = vec![1u8, 12];

        let invalid_addresses = [
            address_bigger_than_u160_max,
            address_in_precompile_space,
            incorrect_length_address,
        ];

        let mut message = sample_message();
        for address in invalid_addresses {
            message.receiver = address;
            assert_eq!(
                validate_svm2any(&message, &sample_dest_chain(), &sample_billing_config())
                    .unwrap_err(),
                FeeQuoterError::InvalidEVMAddress.into()
            );
        }
    }

    #[test]
    fn message_with_too_many_tokens_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.token_amounts = vec![
            SVMTokenAmount {
                token: Pubkey::new_unique(),
                amount: 1
            };
            dest_chain.config.max_number_of_tokens_per_msg as usize + 1
        ];
        assert_eq!(
            validate_svm2any(&message, &sample_dest_chain(), &sample_billing_config()).unwrap_err(),
            FeeQuoterError::UnsupportedNumberOfTokens.into()
        );
    }

    pub fn sample_message() -> SVM2AnyMessage {
        let mut receiver = vec![0u8; 32];

        // Arbitrary value that pushes the address to the right EVM range
        // (above precompile space, under u160::max)
        receiver[20] = 0xA;

        SVM2AnyMessage {
            receiver,
            data: vec![],
            token_amounts: vec![],
            fee_token: native_mint::ID,
            extra_args: ExtraArgsInput {
                gas_limit: None,
                allow_out_of_order_execution: None,
            },
        }
    }

    pub fn sample_billing_config() -> BillingTokenConfig {
        // All values are derived from inspecting the Ethereum Mainnet -> Avalanche lane as of Jan 9 2025,
        // using USDC as transfer token and LINK as fee token wherever applicable (these are not important,
        // they were used to retrieve correctly dimensioned values)

        let arbitrary_timestamp = 100;
        let usd_per_token = TimestampedPackedU224 {
            timestamp: arbitrary_timestamp,
            value: as_u8_28(U256::new(19816680000000000000)),
        };

        BillingTokenConfig {
            enabled: true,
            mint: native_mint::ID,
            usd_per_token,
            premium_multiplier_wei_per_eth: 900000000000000000,
        }
    }

    pub fn sample_dest_chain() -> DestChain {
        // All values are derived from inspecting the Ethereum Mainnet -> Avalanche lane as of Jan 9 2025,
        // using USDC as transfer token and LINK as fee token wherever applicable (these are not important,
        // they were used to retrieve correctly dimensioned values)
        let usd_per_unit_gas = TimestampedPackedU224 {
            // value encodes execution_gas_price of Usd18Decimals(U256::new(921441088750)) and data_availability_gas_price of 0
            value: as_u8_28(U256::new(921441088750)),
            timestamp: 100, // arbitrary value
        };

        DestChain {
            version: 1,
            chain_selector: 1,
            state: crate::DestChainState { usd_per_unit_gas },
            config: crate::DestChainConfig {
                is_enabled: true,
                max_number_of_tokens_per_msg: 1,
                max_data_bytes: 30000,
                max_per_msg_gas_limit: 3000000,
                dest_gas_overhead: 300000,
                dest_gas_per_payload_byte_base: 16,
                dest_gas_per_payload_byte_high: 40,
                dest_gas_per_payload_byte_threshold: 3000,
                dest_data_availability_overhead_gas: 0,
                dest_gas_per_data_availability_byte: 16,
                dest_data_availability_multiplier_bps: 0,
                default_token_fee_usdcents: 50,
                default_token_dest_gas_overhead: 90000,
                default_tx_gas_limit: 200000,
                gas_multiplier_wei_per_eth: 1100000000000000000,
                network_fee_usdcents: 50,
                gas_price_staleness_threshold: 90000,
                enforce_out_of_order: false,
                chain_family_selector: CHAIN_FAMILY_SELECTOR_EVM.to_be_bytes(),
            },
        }
    }
}
