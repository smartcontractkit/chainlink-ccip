use anchor_lang::prelude::*;

use ccip_common::v1::{validate_evm_address, validate_svm_address};
use ccip_common::{CommonCcipError, CHAIN_FAMILY_SELECTOR_EVM, CHAIN_FAMILY_SELECTOR_SVM};

use crate::extra_args::{
    GenericExtraArgsV2, SVMExtraArgsV1, GENERIC_EXTRA_ARGS_V2_TAG, SVM_EXTRA_ARGS_MAX_ACCOUNTS,
    SVM_EXTRA_ARGS_V1_TAG,
};
use crate::messages::{ProcessedExtraArgs, SVM2AnyMessage};
use crate::state::{BillingTokenConfig, DestChain, DestChainConfig};
use crate::FeeQuoterError;

pub fn validate_svm2any(
    msg: &SVM2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
) -> Result<ProcessedExtraArgs> {
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

    let processed_extra_args = process_extra_args(
        &dest_chain.config,
        &msg.extra_args,
        !msg.token_amounts.is_empty(),
    )?;

    require_gte!(
        dest_chain.config.max_per_msg_gas_limit as u128,
        processed_extra_args.gas_limit,
        FeeQuoterError::MessageGasLimitTooHigh,
    );

    require!(
        !dest_chain.config.enforce_out_of_order
            || processed_extra_args.allow_out_of_order_execution,
        FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue,
    );

    validate_dest_family_address(
        msg,
        dest_chain.config.chain_family_selector,
        &processed_extra_args,
    )?;

    Ok(processed_extra_args)
}

fn validate_dest_family_address(
    msg: &SVM2AnyMessage,
    chain_family_selector: [u8; 4],
    msg_extra_args: &ProcessedExtraArgs,
) -> Result<()> {
    let selector = u32::from_be_bytes(chain_family_selector);
    match selector {
        CHAIN_FAMILY_SELECTOR_EVM => validate_evm_address(&msg.receiver),
        CHAIN_FAMILY_SELECTOR_SVM => {
            validate_svm_address(&msg.receiver, msg_extra_args.gas_limit > 0)
        }
        _ => Err(CommonCcipError::InvalidChainFamilySelector.into()),
    }
}

// process_extra_args returns serialized extraArgs, gas_limit, allow_out_of_order_execution
// it calls the chain-specific extra args validation logic
pub fn process_extra_args(
    dest_config: &DestChainConfig,
    extra_args: &[u8],
    message_contains_tokens: bool,
) -> Result<ProcessedExtraArgs> {
    match u32::from_be_bytes(dest_config.chain_family_selector) {
        CHAIN_FAMILY_SELECTOR_SVM => {
            // Extra args are mandatory for a SVM destination, so the tag must exist.
            require_gte!(
                extra_args.len(),
                4,
                FeeQuoterError::InvalidInputsMissingExtraArgs
            );
            let tag: [u8; 4] = extra_args[..4].try_into().unwrap();
            let mut data = &extra_args[4..];
            Ok(parse_and_validate_svm_extra_args(
                dest_config,
                tag,
                &mut data,
                message_contains_tokens,
            )?)
        }
        _ => {
            // Extra args are optional for non-SVM destinations. In case there
            // are extra args, they must be prefixed by a four byte tag
            // -> bytes4(keccak256("CCIP EVMExtraArgsV2"));
            let Some(tag) = extra_args.get(..4) else {
                return Ok(ProcessedExtraArgs::defaults(dest_config));
            };

            let mut data = &extra_args[4..];
            parse_and_validate_generic_extra_args(tag.try_into().unwrap(), &mut data)
        }
    }
}

fn parse_and_validate_generic_extra_args(
    tag: [u8; 4],
    data: &mut &[u8],
) -> Result<ProcessedExtraArgs> {
    match u32::from_be_bytes(tag) {
        GENERIC_EXTRA_ARGS_V2_TAG => {
            if data.is_empty() {
                Err(FeeQuoterError::InvalidInputsMissingDataAfterExtraArgs.into())
            } else {
                let args = GenericExtraArgsV2::deserialize(data)?;
                Ok(ProcessedExtraArgs {
                    bytes: args.serialize_with_tag(),
                    gas_limit: args.gas_limit,
                    allow_out_of_order_execution: args.allow_out_of_order_execution,
                    token_receiver: None,
                })
            }
        }
        _ => Err(FeeQuoterError::InvalidExtraArgsTag.into()),
    }
}

fn parse_and_validate_svm_extra_args(
    cfg: &DestChainConfig,
    tag: [u8; 4],
    data: &mut &[u8],
    message_contains_tokens: bool,
) -> Result<ProcessedExtraArgs> {
    match u32::from_be_bytes(tag) {
        SVM_EXTRA_ARGS_V1_TAG => {
            let args = if data.is_empty() {
                SVMExtraArgsV1::default_config(cfg)
            } else {
                let args = SVMExtraArgsV1::deserialize(data)?;
                require_gte!(
                    SVM_EXTRA_ARGS_MAX_ACCOUNTS,
                    args.accounts.len(),
                    FeeQuoterError::InvalidExtraArgsAccounts
                );
                require!(
                    args.account_is_writable_bitmap >> args.accounts.len() == 0,
                    FeeQuoterError::InvalidExtraArgsWritabilityBitmap
                );
                args
            };

            // token_receiver != 0 when tokens are present
            // token_receiver == 0 when tokens are not present
            let receiver_is_zero_address = args.token_receiver == [0; 32];
            require!(
                message_contains_tokens != receiver_is_zero_address,
                FeeQuoterError::InvalidTokenReceiver
            );

            Ok(ProcessedExtraArgs {
                bytes: args.serialize_with_tag(),
                gas_limit: args.compute_units as u128,
                allow_out_of_order_execution: args.allow_out_of_order_execution,
                token_receiver: message_contains_tokens.then_some(args.token_receiver.to_vec()),
            })
        }
        _ => Err(FeeQuoterError::InvalidExtraArgsTag.into()),
    }
}

#[cfg(test)]
pub mod tests {
    use super::*;
    use crate::extra_args::{GenericExtraArgsV2, GENERIC_EXTRA_ARGS_V2_TAG, SVM_EXTRA_ARGS_V1_TAG};
    use crate::{SVMTokenAmount, TimestampedPackedU224};
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;
    use ethnum::U256;

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
                CommonCcipError::InvalidEVMAddress.into()
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

    #[test]
    fn message_exceeds_gas_limit_fails_to_validate() {
        let mut message = sample_message();
        message.extra_args = GenericExtraArgsV2 {
            gas_limit: 1_000_000_000,
            allow_out_of_order_execution: false,
        }
        .serialize_with_tag();
        assert_eq!(
            validate_svm2any(&message, &sample_dest_chain(), &sample_billing_config()).unwrap_err(),
            FeeQuoterError::MessageGasLimitTooHigh.into()
        );
    }

    #[test]
    fn validate_out_of_order_execution() {
        let mut dest_chain_enforce = sample_dest_chain();
        dest_chain_enforce.config.enforce_out_of_order = true;
        let mut dest_chain_not_enforce = sample_dest_chain();
        dest_chain_not_enforce.config.enforce_out_of_order = false;

        let mut message_ooo = sample_message();
        message_ooo.extra_args = GenericExtraArgsV2 {
            gas_limit: 1_000,
            allow_out_of_order_execution: true,
        }
        .serialize_with_tag();
        let mut message_not_ooo = sample_message();
        message_not_ooo.extra_args = GenericExtraArgsV2 {
            gas_limit: 1_000,
            allow_out_of_order_execution: false,
        }
        .serialize_with_tag();

        // allowed cases
        validate_svm2any(&message_ooo, &dest_chain_enforce, &sample_billing_config()).unwrap();
        validate_svm2any(
            &message_ooo,
            &dest_chain_not_enforce,
            &sample_billing_config(),
        )
        .unwrap();
        validate_svm2any(
            &message_not_ooo,
            &dest_chain_not_enforce,
            &sample_billing_config(),
        )
        .unwrap();

        // not allowed cases
        assert_eq!(
            validate_svm2any(
                &message_not_ooo,
                &dest_chain_enforce,
                &sample_billing_config()
            )
            .unwrap_err(),
            FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue.into()
        );
    }

    #[test]
    fn process_extra_args_matches_family() {
        let evm_dest_chain = sample_dest_chain();
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();
        let evm_tag_bytes = GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes().to_vec();
        let svm_tag_bytes = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();
        let mut none_dest_chain = sample_dest_chain();
        none_dest_chain.config.chain_family_selector = [0; 4];

        // evm - tag but no data fails
        assert_eq!(
            process_extra_args(&evm_dest_chain.config, &evm_tag_bytes, false).unwrap_err(),
            FeeQuoterError::InvalidInputsMissingDataAfterExtraArgs.into()
        );

        // evm - default case: (no data or tag)
        let extra_args = process_extra_args(&evm_dest_chain.config, &[], false).unwrap();
        assert_eq!(
            extra_args.bytes[..4],
            GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes()
        );
        assert_eq!(
            extra_args.gas_limit,
            evm_dest_chain.config.default_tx_gas_limit as u128
        );
        assert!(!extra_args.allow_out_of_order_execution);

        // evm - passed in data
        let extra_args = process_extra_args(
            &evm_dest_chain.config,
            &GenericExtraArgsV2 {
                gas_limit: 100,
                allow_out_of_order_execution: true,
            }
            .serialize_with_tag(),
            false,
        )
        .unwrap();
        assert_eq!(
            extra_args.bytes[..4],
            GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes()
        );
        assert_eq!(extra_args.gas_limit, 100);
        assert!(extra_args.allow_out_of_order_execution);
        assert_eq!(extra_args.token_receiver, None);

        // unknown family - uses generic (same as evm) tag
        let extra_args = process_extra_args(
            &none_dest_chain.config,
            &GenericExtraArgsV2 {
                gas_limit: 100,
                allow_out_of_order_execution: true,
            }
            .serialize_with_tag(),
            false,
        )
        .unwrap();
        assert_eq!(
            extra_args.bytes[..4],
            GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes()
        );
        assert_eq!(extra_args.gas_limit, 100);
        assert!(extra_args.allow_out_of_order_execution);
        assert_eq!(extra_args.token_receiver, None);

        // evm - fail to match
        assert_eq!(
            process_extra_args(&evm_dest_chain.config, &svm_tag_bytes, false).unwrap_err(),
            FeeQuoterError::InvalidExtraArgsTag.into()
        );

        // svm - default case
        let extra_args = process_extra_args(&svm_dest_chain.config, &svm_tag_bytes, false).unwrap();
        assert_eq!(extra_args.bytes[..4], SVM_EXTRA_ARGS_V1_TAG.to_be_bytes());
        assert_eq!(
            extra_args.gas_limit,
            svm_dest_chain.config.default_tx_gas_limit as u128
        );
        // Reflects the no token case
        assert_eq!(extra_args.token_receiver, None);
        assert!(!extra_args.allow_out_of_order_execution);

        // svm - empty tag (no data) fails
        assert_eq!(
            process_extra_args(&svm_dest_chain.config, &[], false).unwrap_err(),
            FeeQuoterError::InvalidInputsMissingExtraArgs.into()
        );

        // svm - contains tokens but no receiver address
        assert_eq!(
            process_extra_args(&svm_dest_chain.config, &svm_tag_bytes, true).unwrap_err(),
            FeeQuoterError::InvalidTokenReceiver.into(),
        );

        // svm - passed in data
        let token_receiver = Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa")
            .unwrap()
            .to_bytes();
        let args = SVMExtraArgsV1 {
            compute_units: 100,
            account_is_writable_bitmap: 3,
            allow_out_of_order_execution: true,
            token_receiver,
            accounts: vec![
                Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa")
                    .unwrap()
                    .to_bytes(),
                Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTa")
                    .unwrap()
                    .to_bytes(),
            ],
        };
        let extra_args =
            process_extra_args(&svm_dest_chain.config, &args.serialize_with_tag(), true).unwrap();
        assert_eq!(extra_args.bytes, args.serialize_with_tag());
        assert_eq!(extra_args.gas_limit, 100);
        assert_eq!(extra_args.token_receiver, Some(token_receiver.to_vec()));
        assert!(extra_args.allow_out_of_order_execution);

        // svm - fail to match
        assert_eq!(
            process_extra_args(&svm_dest_chain.config, &evm_tag_bytes, false).unwrap_err(),
            FeeQuoterError::InvalidExtraArgsTag.into()
        );
    }

    pub fn as_u8_28(single: U256) -> [u8; 28] {
        single.to_be_bytes()[4..32].try_into().unwrap()
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
            extra_args: vec![], // empty extraArgs, use defaults
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
                lane_code_version: crate::state::CodeVersion::Default,
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
