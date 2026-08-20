use anchor_lang::prelude::*;

use ccip_common::v1::{
    validate_aptos_address, validate_evm_address, validate_svm_address, validate_tvm_address,
    SUI_PRECOMPILE_SPACE,
};
use ccip_common::{
    CommonCcipError, CHAIN_FAMILY_SELECTOR_APTOS, CHAIN_FAMILY_SELECTOR_EVM,
    CHAIN_FAMILY_SELECTOR_SUI, CHAIN_FAMILY_SELECTOR_SVM, CHAIN_FAMILY_SELECTOR_TVM,
};

use crate::extra_args::{
    GenericExtraArgsV2, SVMExtraArgsV1, SuiExtraArgsV1, GENERIC_EXTRA_ARGS_V2_TAG,
    SUI_EXTRA_ARGS_V1_TAG, SVM_EXTRA_ARGS_MAX_ACCOUNTS, SVM_EXTRA_ARGS_V1_TAG,
};
use crate::messages::{ProcessedExtraArgs, SVM2AnyMessage};
use crate::state::{BillingTokenConfig, DestChain, DestChainConfig};
use crate::FeeQuoterError;

/// The maximum number of receiver object ids that can be passed in SuiExtraArgs.
pub const SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS: usize = 64;

/// Number of overhead accounts needed for message execution on SUI.
/// This is the message.receiver.
pub const SUI_MESSAGING_ACCOUNTS_OVERHEAD: usize = 1;

/// The size of each SUI account address in bytes.
pub const SUI_ACCOUNT_BYTE_SIZE: usize = 32;

/// The expected static payload size of a token transfer when BCS encoded and submitted to SUI.
/// TokenPool extra data and offchain data sizes are dynamic, and should be accounted for separately.
pub const SUI_TOKEN_TRANSFER_DATA_OVERHEAD: usize = (4 + 32) // source_pool, 4 bytes for length, 32 bytes for address
    + 32 // dest_token_address
    + 4 // dest_gas_amount
    + 4 // extra_data length, the contents are calculated separately
    + 32; // amount

/// The expected static payload size of a token transfer when Borsh encoded and submitted to SVM.
/// TokenPool extra data and offchain data sizes are dynamic, and should be accounted for separately.
pub const SVM_TOKEN_TRANSFER_DATA_OVERHEAD: usize = (4 + 32) // source_pool
    + 32 // token_address
    + 4 // gas_amount
    + 4 // extra_data overhead
    + 32 // amount
    + 32 // size of the token lookup table account
    + 32 // token-related accounts in the lookup table, over-estimated to 32, typically between 11 - 13
    + 32 // token account belonging to the token receiver, e.g ATA, not included in the token lookup table
    + 32 // per-chain token pool config, not included in the token lookup table
    + 32 // per-chain token billing config, not always included in the token lookup table
    + 32; // OffRamp pool signer PDA, not included in the token lookup table

/// Number of overhead accounts needed for message execution on SVM.
/// These are message.receiver and the OffRamp Signer PDA specific to the receiver.
pub const SVM_MESSAGING_ACCOUNTS_OVERHEAD: usize = 2;

/// The size of each SVM account address in bytes.
pub const SVM_ACCOUNT_BYTE_SIZE: usize = 32;

pub struct MessageInfo {
    pub number_of_tokens: usize,
    pub contains_receiver: bool,
    pub data_len: usize,
}

pub struct ValidatedMessage {
    pub processed_extra_args: ProcessedExtraArgs,
    pub extra_args_data_len: u32,
}

pub fn validate_svm2any(
    msg: &SVM2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
    dest_bytes_overhead: &u32,
) -> Result<ValidatedMessage> {
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

    let validated_message = process_extra_args_with_data_len(
        &dest_chain.config,
        &msg.extra_args,
        &MessageInfo {
            number_of_tokens: msg.token_amounts.len(),
            contains_receiver: msg.receiver != [0; 32],
            data_len: msg.data.len(),
        },
        dest_bytes_overhead,
    )?;

    require_gte!(
        dest_chain.config.max_per_msg_gas_limit as u128,
        validated_message.processed_extra_args.gas_limit,
        FeeQuoterError::MessageGasLimitTooHigh,
    );

    require!(
        !dest_chain.config.enforce_out_of_order
            || validated_message
                .processed_extra_args
                .allow_out_of_order_execution,
        FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue,
    );

    validate_dest_family_address(
        msg,
        dest_chain.config.chain_family_selector,
        &validated_message.processed_extra_args,
    )?;

    Ok(validated_message)
}

fn validate_dest_family_address(
    msg: &SVM2AnyMessage,
    chain_family_selector: [u8; 4],
    msg_extra_args: &ProcessedExtraArgs,
) -> Result<()> {
    let selector = u32::from_be_bytes(chain_family_selector);
    match selector {
        CHAIN_FAMILY_SELECTOR_APTOS => validate_aptos_address(&msg.receiver),
        CHAIN_FAMILY_SELECTOR_EVM => validate_evm_address(&msg.receiver),
        CHAIN_FAMILY_SELECTOR_SUI => {
            validate_sui_address(&msg.receiver, msg_extra_args.gas_limit > 0)
        }
        CHAIN_FAMILY_SELECTOR_SVM => {
            validate_svm_address(&msg.receiver, msg_extra_args.gas_limit > 0)
        }
        CHAIN_FAMILY_SELECTOR_TVM => validate_tvm_address(&msg.receiver),
        _ => Err(CommonCcipError::InvalidChainFamilySelector.into()),
    }
}

// process_extra_args_with_data_len returns serialized extraArgs, gas_limit, allow_out_of_order_execution
// it calls the chain-specific extra args validation logic
fn process_extra_args_with_data_len(
    dest_config: &DestChainConfig,
    extra_args: &[u8],
    message_info: &MessageInfo,
    dest_bytes_overhead: &u32,
) -> Result<ValidatedMessage> {
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
            parse_and_validate_svm_extra_args(
                dest_config,
                tag,
                &mut data,
                message_info,
                dest_bytes_overhead,
            )
        }
        CHAIN_FAMILY_SELECTOR_SUI => {
            require_gte!(
                extra_args.len(),
                4,
                FeeQuoterError::InvalidInputsMissingExtraArgs
            );

            let tag: [u8; 4] = extra_args[..4].try_into().unwrap();

            let mut data = &extra_args[4..];
            parse_and_validate_sui_extra_args(
                dest_config,
                tag,
                &mut data,
                message_info,
                dest_bytes_overhead,
            )
        }
        _ => {
            // Extra args are optional for other chain family destinations. In case there
            // are extra args, they must be prefixed by a four byte tag
            // -> bytes4(keccak256("CCIP EVMExtraArgsV2"));
            let Some(tag) = extra_args.get(..4) else {
                return Ok(ValidatedMessage {
                    processed_extra_args: ProcessedExtraArgs::defaults(dest_config),
                    extra_args_data_len: 0,
                });
            };

            let mut data = &extra_args[4..];
            parse_and_validate_generic_extra_args(tag.try_into().unwrap(), &mut data)
        }
    }
}

fn parse_and_validate_generic_extra_args(
    tag: [u8; 4],
    data: &mut &[u8],
) -> Result<ValidatedMessage> {
    match u32::from_be_bytes(tag) {
        GENERIC_EXTRA_ARGS_V2_TAG => {
            if data.is_empty() {
                Err(FeeQuoterError::InvalidInputsMissingDataAfterExtraArgs.into())
            } else {
                let args = GenericExtraArgsV2::deserialize(data)?;
                Ok(ValidatedMessage {
                    processed_extra_args: ProcessedExtraArgs {
                        bytes: args.serialize_with_tag(),
                        gas_limit: args.gas_limit,
                        allow_out_of_order_execution: args.allow_out_of_order_execution,
                        token_receiver: None,
                    },
                    extra_args_data_len: 0,
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
    message_info: &MessageInfo,
    dest_bytes_overhead: &u32,
) -> Result<ValidatedMessage> {
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
                let has_writable_bits_beyond_accounts = (args.accounts.len()..u64::BITS as usize)
                    .any(|bit| args.account_is_writable_bitmap & (1u64 << bit) != 0);
                require!(
                    !has_writable_bits_beyond_accounts,
                    FeeQuoterError::InvalidExtraArgsWritabilityBitmap
                );
                args
            };

            // token_receiver != 0 when tokens are present
            // token_receiver == 0 when tokens are not present
            let receiver_is_zero_address = args.token_receiver == [0; 32];
            let message_contains_tokens = message_info.number_of_tokens != 0;
            require!(
                message_contains_tokens != receiver_is_zero_address,
                FeeQuoterError::InvalidTokenReceiver
            );

            let accounts_len = args.accounts.len();
            let mut svm_expanded_data_len = message_info.data_len;
            let mut extra_args_data_len = 0;

            if !message_info.contains_receiver {
                require_eq!(accounts_len, 0, FeeQuoterError::InvalidExtraArgsAccounts);
            } else {
                extra_args_data_len =
                    (accounts_len + SVM_MESSAGING_ACCOUNTS_OVERHEAD) * SVM_ACCOUNT_BYTE_SIZE;
                svm_expanded_data_len += extra_args_data_len;
            }

            let token_transfer_data_overhead =
                message_info.number_of_tokens * SVM_TOKEN_TRANSFER_DATA_OVERHEAD;
            extra_args_data_len += token_transfer_data_overhead;
            svm_expanded_data_len += token_transfer_data_overhead;

            svm_expanded_data_len += *dest_bytes_overhead as usize;

            require_gte!(
                cfg.max_data_bytes as usize,
                svm_expanded_data_len,
                FeeQuoterError::MessageTooLarge,
            );

            Ok(ValidatedMessage {
                processed_extra_args: ProcessedExtraArgs {
                    bytes: args.serialize_with_tag(),
                    gas_limit: args.compute_units as u128,
                    allow_out_of_order_execution: args.allow_out_of_order_execution,
                    token_receiver: message_contains_tokens.then_some(args.token_receiver.to_vec()),
                },
                extra_args_data_len: extra_args_data_len as u32,
            })
        }
        _ => Err(FeeQuoterError::InvalidExtraArgsTag.into()),
    }
}

fn parse_and_validate_sui_extra_args(
    cfg: &DestChainConfig,
    tag: [u8; 4],
    data: &mut &[u8],
    message_info: &MessageInfo,
    dest_bytes_overhead: &u32,
) -> Result<ValidatedMessage> {
    match u32::from_be_bytes(tag) {
        SUI_EXTRA_ARGS_V1_TAG => {
            let args = SuiExtraArgsV1::deserialize(data)?;

            let mut sui_expanded_data_len = message_info.data_len;
            let mut extra_args_data_len = 0;

            // token_receiver != 0 when tokens are present
            let receiver_is_present = args.token_receiver != [0; 32];
            let message_contains_tokens = message_info.number_of_tokens > 0;
            require!(
                !message_contains_tokens || receiver_is_present,
                FeeQuoterError::InvalidTokenReceiver
            );

            let receiver_object_ids_len = args.receiver_object_ids.len();

            if !message_info.contains_receiver {
                require_eq!(
                    receiver_object_ids_len,
                    0,
                    FeeQuoterError::InvalidSuiReceiverObjectIds,
                );
            } else {
                extra_args_data_len = (receiver_object_ids_len + SUI_MESSAGING_ACCOUNTS_OVERHEAD)
                    * SUI_ACCOUNT_BYTE_SIZE;
                sui_expanded_data_len += extra_args_data_len;
            }

            require_gte!(
                SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS,
                args.receiver_object_ids.len(),
                FeeQuoterError::InvalidSuiReceiverObjectIds,
            );

            let token_transfer_data_overhead =
                message_info.number_of_tokens * SUI_TOKEN_TRANSFER_DATA_OVERHEAD;
            extra_args_data_len += token_transfer_data_overhead;
            sui_expanded_data_len += token_transfer_data_overhead;

            // The token destBytesOverhead can be very different per token so we have to take it into account as well.
            sui_expanded_data_len += *dest_bytes_overhead as usize;

            require_gte!(
                cfg.max_data_bytes as usize,
                sui_expanded_data_len,
                FeeQuoterError::MessageTooLarge,
            );

            Ok(ValidatedMessage {
                processed_extra_args: ProcessedExtraArgs {
                    bytes: args.serialize_with_tag(),
                    gas_limit: args.gas_limit,
                    allow_out_of_order_execution: args.allow_out_of_order_execution,
                    token_receiver: message_contains_tokens.then_some(args.token_receiver.to_vec()),
                },
                extra_args_data_len: extra_args_data_len as u32,
            })
        }
        _ => Err(FeeQuoterError::InvalidExtraArgsTag.into()),
    }
}

fn validate_sui_address(
    address: &[u8],
    address_must_be_outside_precompile_space: bool,
) -> Result<()> {
    require_eq!(address.len(), 32, CommonCcipError::InvalidSuiAddress);

    let addr_32: &[u8; 32] = address
        .try_into()
        .expect("slice length is guaranteed to be 32");

    require!(
        !address_must_be_outside_precompile_space || addr_32 >= &SUI_PRECOMPILE_SPACE,
        CommonCcipError::InvalidSuiAddress
    );

    Ok(())
}

#[cfg(test)]
pub mod tests {
    use super::*;
    use crate::extra_args::{GenericExtraArgsV2, GENERIC_EXTRA_ARGS_V2_TAG, SVM_EXTRA_ARGS_V1_TAG};
    use crate::{SVMTokenAmount, TimestampedPackedU224};
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;
    use ethnum::U256;

    fn process_extra_args(
        dest_config: &DestChainConfig,
        extra_args: &[u8],
        message_info: &MessageInfo,
        dest_bytes_overhead: &u32,
    ) -> Result<ProcessedExtraArgs> {
        Ok(process_extra_args_with_data_len(
            dest_config,
            extra_args,
            message_info,
            dest_bytes_overhead,
        )?
        .processed_extra_args)
    }

    fn validation_err<T>(result: Result<T>) -> anchor_lang::error::Error {
        match result {
            Ok(_) => panic!("expected validation to fail"),
            Err(err) => err,
        }
    }

    #[test]
    fn message_not_validated_for_disabled_destination_chain() {
        let mut chain = sample_dest_chain();
        chain.config.is_enabled = false;

        assert_eq!(
            validation_err(validate_svm2any(
                &sample_message(),
                &chain,
                &sample_billing_config(),
                &0
            )),
            FeeQuoterError::DestinationChainDisabled.into()
        );
    }

    #[test]
    fn message_not_validated_for_disabled_token() {
        let mut billing_config = sample_billing_config();
        billing_config.enabled = false;

        assert_eq!(
            validation_err(validate_svm2any(
                &sample_message(),
                &sample_dest_chain(),
                &billing_config,
                &0
            )),
            FeeQuoterError::FeeTokenDisabled.into()
        );
    }

    #[test]
    fn large_message_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.data = vec![0; dest_chain.config.max_data_bytes as usize + 1];
        assert_eq!(
            validation_err(validate_svm2any(
                &message,
                &sample_dest_chain(),
                &sample_billing_config(),
                &0
            )),
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
                validation_err(validate_svm2any(
                    &message,
                    &sample_dest_chain(),
                    &sample_billing_config(),
                    &0
                )),
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
            validation_err(validate_svm2any(
                &message,
                &sample_dest_chain(),
                &sample_billing_config(),
                &0
            )),
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
            validation_err(validate_svm2any(
                &message,
                &sample_dest_chain(),
                &sample_billing_config(),
                &0
            )),
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
        validate_svm2any(
            &message_ooo,
            &dest_chain_enforce,
            &sample_billing_config(),
            &0,
        )
        .unwrap();
        validate_svm2any(
            &message_ooo,
            &dest_chain_not_enforce,
            &sample_billing_config(),
            &0,
        )
        .unwrap();
        validate_svm2any(
            &message_not_ooo,
            &dest_chain_not_enforce,
            &sample_billing_config(),
            &0,
        )
        .unwrap();

        // not allowed cases
        assert_eq!(
            validation_err(validate_svm2any(
                &message_not_ooo,
                &dest_chain_enforce,
                &sample_billing_config(),
                &0
            )),
            FeeQuoterError::ExtraArgOutOfOrderExecutionMustBeTrue.into()
        );
    }

    #[test]
    fn process_extra_args_matches_family_evm_and_unknown() {
        let evm_dest_chain = sample_dest_chain();
        let mut none_dest_chain = sample_dest_chain();
        none_dest_chain.config.chain_family_selector = [0; 4];

        let svm_tag_bytes = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();

        // EVM behaves like generic.
        assert_generic_family_behaviour(&evm_dest_chain, &svm_tag_bytes);

        // Unknown family selector also behaves like generic.
        assert_generic_family_behaviour(&none_dest_chain, &svm_tag_bytes);
    }

    #[test]
    fn process_extra_args_matches_family_svm() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let generic_tag_bytes = GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes().to_vec();
        assert_svm_family_behaviour(&svm_dest_chain, &generic_tag_bytes);
    }

    #[test]
    fn process_extra_args_matches_family_tvm() {
        let mut tvm_dest_chain = sample_dest_chain();
        tvm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_TVM.to_be_bytes();

        let svm_tag_bytes = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();
        assert_generic_family_behaviour(&tvm_dest_chain, &svm_tag_bytes);
    }

    #[test]
    fn process_extra_args_matches_family_aptos() {
        let mut aptos_dest_chain = sample_dest_chain();
        aptos_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_APTOS.to_be_bytes();

        let svm_tag_bytes = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();
        assert_generic_family_behaviour(&aptos_dest_chain, &svm_tag_bytes);
    }

    fn assert_generic_family_behaviour(dest_chain: &DestChain, other_family_tag: &[u8]) {
        let generic_tag_bytes = GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes().to_vec();

        // tag but no data fails
        assert_eq!(
            process_extra_args(
                &dest_chain.config,
                &generic_tag_bytes,
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidInputsMissingDataAfterExtraArgs.into()
        );

        // default case: (no data or tag)
        let extra_args = process_extra_args(
            &dest_chain.config,
            &[],
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: false,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
        assert_eq!(
            extra_args.bytes[..4],
            GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes()
        );
        assert_eq!(
            extra_args.gas_limit,
            dest_chain.config.default_tx_gas_limit as u128
        );
        assert!(!extra_args.allow_out_of_order_execution);
        assert_eq!(extra_args.token_receiver, None);

        // passed in data
        let extra_args = process_extra_args(
            &dest_chain.config,
            &GenericExtraArgsV2 {
                gas_limit: 100,
                allow_out_of_order_execution: true,
            }
            .serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: false,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
        assert_eq!(
            extra_args.bytes[..4],
            GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes()
        );
        assert_eq!(extra_args.gas_limit, 100);
        assert!(extra_args.allow_out_of_order_execution);
        assert_eq!(extra_args.token_receiver, None);

        // fail to match an unrelated family's tag
        assert_eq!(
            process_extra_args(
                &dest_chain.config,
                other_family_tag,
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidExtraArgsTag.into()
        );
    }

    fn assert_svm_family_behaviour(svm_dest_chain: &DestChain, generic_tag_bytes: &[u8]) {
        let svm_tag_bytes = SVM_EXTRA_ARGS_V1_TAG.to_be_bytes().to_vec();

        // default case requires the SVM tag bytes
        let extra_args = process_extra_args(
            &svm_dest_chain.config,
            &svm_tag_bytes,
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: false,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
        assert_eq!(extra_args.bytes[..4], SVM_EXTRA_ARGS_V1_TAG.to_be_bytes());
        assert_eq!(
            extra_args.gas_limit,
            svm_dest_chain.config.default_tx_gas_limit as u128
        );
        // Reflects the no token case
        assert_eq!(extra_args.token_receiver, None);
        assert!(!extra_args.allow_out_of_order_execution);

        // empty tag (no data) fails
        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                &[],
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidInputsMissingExtraArgs.into()
        );

        // contains tokens but no receiver address
        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                &svm_tag_bytes,
                &MessageInfo {
                    number_of_tokens: 1,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidTokenReceiver.into(),
        );

        // passed in data
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
        let extra_args = process_extra_args(
            &svm_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: true,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
        assert_eq!(extra_args.bytes, args.serialize_with_tag());
        assert_eq!(extra_args.gas_limit, 100);
        assert_eq!(extra_args.token_receiver, Some(token_receiver.to_vec()));
        assert!(extra_args.allow_out_of_order_execution);

        // fail to match generic tag
        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                generic_tag_bytes,
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidExtraArgsTag.into()
        );
    }

    #[test]
    fn svm_extra_args_data_len_matches_destination_payload_overhead() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let token_receiver = [1; 32];
        let args = SVMExtraArgsV1 {
            compute_units: 100,
            allow_out_of_order_execution: true,
            token_receiver,
            accounts: vec![[1; 32], [2; 32]],
            ..Default::default()
        };

        let validated_message = process_extra_args_with_data_len(
            &svm_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: true,
                data_len: 0,
            },
            &100,
        )
        .unwrap();

        assert_eq!(
            validated_message.extra_args_data_len as usize,
            (args.accounts.len() + SVM_MESSAGING_ACCOUNTS_OVERHEAD) * SVM_ACCOUNT_BYTE_SIZE
                + SVM_TOKEN_TRANSFER_DATA_OVERHEAD
        );
    }

    #[test]
    fn svm_expanded_payload_counts_towards_max_data_bytes() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let args = SVMExtraArgsV1 {
            token_receiver: [1; 32],
            accounts: vec![[1; 32], [2; 32]],
            ..Default::default()
        };

        let dest_bytes_overhead = 100_u32;
        let account_overhead =
            (args.accounts.len() + SVM_MESSAGING_ACCOUNTS_OVERHEAD) * SVM_ACCOUNT_BYTE_SIZE;
        let token_overhead = SVM_TOKEN_TRANSFER_DATA_OVERHEAD + dest_bytes_overhead as usize;
        let total_overhead = account_overhead + token_overhead;
        let data_len = svm_dest_chain.config.max_data_bytes as usize - total_overhead;

        process_extra_args(
            &svm_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: true,
                data_len,
            },
            &dest_bytes_overhead,
        )
        .unwrap();

        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 1,
                    contains_receiver: true,
                    data_len: data_len + 1,
                },
                &dest_bytes_overhead,
            )
            .unwrap_err(),
            FeeQuoterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn svm_writable_bitmap_allows_all_bits_for_max_accounts() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let args = SVMExtraArgsV1 {
            accounts: vec![[1; 32]; SVM_EXTRA_ARGS_MAX_ACCOUNTS],
            account_is_writable_bitmap: u64::MAX,
            ..Default::default()
        };

        process_extra_args(
            &svm_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: true,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
    }

    #[test]
    fn svm_writable_bitmap_rejects_bits_beyond_accounts() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let args = SVMExtraArgsV1 {
            accounts: vec![[1; 32]; SVM_EXTRA_ARGS_MAX_ACCOUNTS - 1],
            account_is_writable_bitmap: u64::MAX,
            ..Default::default()
        };

        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: true,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidExtraArgsWritabilityBitmap.into()
        );
    }

    #[test]
    fn svm_accounts_require_message_receiver() {
        let mut svm_dest_chain = sample_dest_chain();
        svm_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SVM.to_be_bytes();

        let args = SVMExtraArgsV1 {
            accounts: vec![[1; 32]],
            ..Default::default()
        };

        assert_eq!(
            process_extra_args(
                &svm_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidExtraArgsAccounts.into()
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

    /// SUI Tests

    #[test]
    fn process_extra_args_matches_family_sui() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        // Empty extra args are invalid.
        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &[],
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidInputsMissingExtraArgs.into()
        );

        // Wrong tag fails.
        let generic_tag = GENERIC_EXTRA_ARGS_V2_TAG.to_be_bytes();

        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &generic_tag,
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidExtraArgsTag.into()
        );
    }

    #[test]
    fn sui_tokens_require_token_receiver() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            token_receiver: [0; 32],
            ..Default::default()
        };

        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 1,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidTokenReceiver.into()
        );
    }

    #[test]
    fn sui_token_receiver_is_allowed_without_tokens() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            token_receiver: [1; 32],
            ..Default::default()
        };

        let extra_args = process_extra_args(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: false,
                data_len: 0,
            },
            &0,
        )
        .unwrap();

        assert_eq!(extra_args.bytes, args.serialize_with_tag());
        assert_eq!(extra_args.token_receiver, None);
    }

    #[test]
    fn sui_receiver_object_ids_require_message_receiver() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            receiver_object_ids: vec![[1; 32]],
            ..Default::default()
        };

        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: false,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidSuiReceiverObjectIds.into()
        );
    }

    #[test]
    fn sui_max_receiver_object_ids_is_allowed() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            receiver_object_ids: vec![[1; 32]; SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS],
            ..Default::default()
        };

        process_extra_args(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: true,
                data_len: 0,
            },
            &0,
        )
        .unwrap();
    }

    #[test]
    fn sui_too_many_receiver_object_ids_fails() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            receiver_object_ids: vec![[1; 32]; SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS + 1],
            ..Default::default()
        };

        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: true,
                    data_len: 0,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::InvalidSuiReceiverObjectIds.into()
        );
    }

    #[test]
    fn sui_receiver_accounts_count_towards_max_data_bytes() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            receiver_object_ids: vec![[1; 32], [2; 32]],
            ..Default::default()
        };

        let receiver_overhead = (args.receiver_object_ids.len() + SUI_MESSAGING_ACCOUNTS_OVERHEAD)
            * SUI_ACCOUNT_BYTE_SIZE;

        let data_len = sui_dest_chain.config.max_data_bytes as usize - receiver_overhead;

        // Exactly max_data_bytes is allowed.
        process_extra_args(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 0,
                contains_receiver: true,
                data_len,
            },
            &0,
        )
        .unwrap();

        // One additional byte exceeds max_data_bytes.
        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 0,
                    contains_receiver: true,
                    data_len: data_len + 1,
                },
                &0,
            )
            .unwrap_err(),
            FeeQuoterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn sui_token_transfer_overhead_counts_towards_max_data_bytes() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            token_receiver: [1; 32],
            ..Default::default()
        };

        let dest_bytes_overhead = 100_u32;

        let total_token_overhead = SUI_TOKEN_TRANSFER_DATA_OVERHEAD + dest_bytes_overhead as usize;

        let data_len = sui_dest_chain.config.max_data_bytes as usize - total_token_overhead;

        // Exactly max_data_bytes is allowed.
        process_extra_args(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: false,
                data_len,
            },
            &dest_bytes_overhead,
        )
        .unwrap();

        // One additional byte exceeds max_data_bytes.
        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 1,
                    contains_receiver: false,
                    data_len: data_len + 1,
                },
                &dest_bytes_overhead,
            )
            .unwrap_err(),
            FeeQuoterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn sui_combined_overheads_count_towards_max_data_bytes() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            token_receiver: [1; 32],
            receiver_object_ids: vec![[1; 32], [2; 32]],
            ..Default::default()
        };

        let dest_bytes_overhead = 100_u32;

        let receiver_overhead = (args.receiver_object_ids.len() + SUI_MESSAGING_ACCOUNTS_OVERHEAD)
            * SUI_ACCOUNT_BYTE_SIZE;

        let token_overhead = SUI_TOKEN_TRANSFER_DATA_OVERHEAD + dest_bytes_overhead as usize;

        let total_overhead = receiver_overhead + token_overhead;

        let data_len = sui_dest_chain.config.max_data_bytes as usize - total_overhead;

        // All overheads combined result in exactly max_data_bytes.
        process_extra_args(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: true,
                data_len,
            },
            &dest_bytes_overhead,
        )
        .unwrap();

        // One additional byte exceeds max_data_bytes.
        assert_eq!(
            process_extra_args(
                &sui_dest_chain.config,
                &args.serialize_with_tag(),
                &MessageInfo {
                    number_of_tokens: 1,
                    contains_receiver: true,
                    data_len: data_len + 1,
                },
                &dest_bytes_overhead,
            )
            .unwrap_err(),
            FeeQuoterError::MessageTooLarge.into()
        );
    }

    #[test]
    fn sui_extra_args_data_len_matches_destination_payload_overhead() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let args = SuiExtraArgsV1 {
            token_receiver: [1; 32],
            receiver_object_ids: vec![[1; 32], [2; 32]],
            ..Default::default()
        };

        let validated_message = process_extra_args_with_data_len(
            &sui_dest_chain.config,
            &args.serialize_with_tag(),
            &MessageInfo {
                number_of_tokens: 1,
                contains_receiver: true,
                data_len: 0,
            },
            &100,
        )
        .unwrap();

        assert_eq!(
            validated_message.extra_args_data_len as usize,
            (args.receiver_object_ids.len() + SUI_MESSAGING_ACCOUNTS_OVERHEAD)
                * SUI_ACCOUNT_BYTE_SIZE
                + SUI_TOKEN_TRANSFER_DATA_OVERHEAD
        );
    }

    #[test]
    fn sui_zero_receiver_is_allowed_when_gas_limit_is_zero() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let mut msg = sample_message();
        msg.receiver = vec![0; 32];
        msg.token_amounts = vec![SVMTokenAmount {
            token: Pubkey::new_unique(),
            amount: 1,
        }];
        msg.extra_args = SuiExtraArgsV1 {
            gas_limit: 0,
            token_receiver: [1; 32],
            ..Default::default()
        }
        .serialize_with_tag();

        validate_svm2any(&msg, &sui_dest_chain, &sample_billing_config(), &0).unwrap();
    }

    #[test]
    fn sui_zero_receiver_is_rejected_when_gas_limit_is_nonzero() {
        let mut sui_dest_chain = sample_dest_chain();
        sui_dest_chain.config.chain_family_selector = CHAIN_FAMILY_SELECTOR_SUI.to_be_bytes();

        let mut msg = sample_message();
        msg.receiver = vec![0; 32];
        msg.extra_args = SuiExtraArgsV1 {
            gas_limit: 1,
            ..Default::default()
        }
        .serialize_with_tag();

        assert_eq!(
            validation_err(validate_svm2any(
                &msg,
                &sui_dest_chain,
                &sample_billing_config(),
                &0
            )),
            CommonCcipError::InvalidSuiAddress.into()
        );
    }
}
