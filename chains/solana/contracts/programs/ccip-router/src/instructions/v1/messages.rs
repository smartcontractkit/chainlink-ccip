use anchor_lang::prelude::*;
use ethnum::U256;

use crate::{
    Any2SolanaRampMessage, BillingTokenConfig, CcipRouterError, DestChain, Solana2AnyMessage,
    Solana2AnyRampMessage, CHAIN_FAMILY_SELECTOR_EVM, TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR,
    TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR,
};

use super::pools::ToTxData;

const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);

///////////////////
// Pool Messages //
///////////////////

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnInV1 {
    pub receiver: Vec<u8>, //  The recipient of the tokens on the destination chain
    pub remote_chain_selector: u64, // The chain ID of the destination chain
    pub original_sender: Pubkey, // The original sender of the tx on the source chain
    pub amount: u64, // solana-specific amount uses u64 - the amount of tokens to lock or burn, denominated in the source token's decimals
    pub local_token: Pubkey, //  The address on this chain of the mint account to lock or burn
}

impl ToTxData for LockOrBurnInV1 {
    fn to_tx_data(&self) -> Vec<u8> {
        let mut data = Vec::new();
        data.extend_from_slice(&TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR);
        data.extend_from_slice(&self.try_to_vec().unwrap());
        data
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintInV1 {
    pub original_sender: Vec<u8>, //          The original sender of the tx on the source chain
    pub remote_chain_selector: u64, // ─╮ The chain ID of the source chain
    pub receiver: Pubkey, // ───────────╯ The Token Associated Account that will receive the tokens on the destination chain.
    pub amount: [u8; 32], // LE u256 amount - The amount of tokens to release or mint, denominated in the source token's decimals, pool expected to handle conversation to solana token specifics
    pub local_token: Pubkey, //            The address of the Token Mint Account on Solana
    /// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
    /// expected pool address for the given remoteChainSelector.
    pub source_pool_address: Vec<u8>, //       The address bytes of the source pool
    pub source_pool_data: Vec<u8>, //          The data received from the source pool to process the release or mint
    /// @dev WARNING: offchainTokenData is untrusted data.
    pub offchain_token_data: Vec<u8>, //       The offchain data to process the release or mint
}

impl ToTxData for ReleaseOrMintInV1 {
    fn to_tx_data(&self) -> Vec<u8> {
        let mut data = Vec::new();
        data.extend_from_slice(&TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR);
        data.extend_from_slice(&self.try_to_vec().unwrap());
        data
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct LockOrBurnOutV1 {
    pub dest_token_address: Vec<u8>,
    pub dest_pool_data: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct ReleaseOrMintOutV1 {
    pub destination_amount: u64, // TODO: u256 on EVM?
}

/////////////////////////////
// Ramp messages functions //
/////////////////////////////

pub fn hash_any2solana(msg: &Any2SolanaRampMessage, on_ramp_address: &[u8]) -> [u8; 32] {
    use anchor_lang::solana_program::hash;

    // Calculate vectors size to ensure that the hash is unique
    let sender_size = [msg.sender.len() as u8];
    let on_ramp_address_size = [on_ramp_address.len() as u8];
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    // Extra Args struct
    let extra_args_compute_units = msg.extra_args.compute_units.to_be_bytes();
    let extra_args_accounts_len = [msg.extra_args.accounts.len() as u8];
    let extra_args_accounts = msg.extra_args.accounts.try_to_vec().unwrap();

    // TODO: Hash token amounts

    // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
    // As similar as https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L402
    let result = hash::hashv(&[
        "Any2SolanaMessageHashV1".as_bytes(),
        &header_source_chain_selector,
        &header_dest_chain_selector,
        &on_ramp_address_size,
        on_ramp_address,
        &msg.header.message_id,
        &msg.receiver.to_bytes(),
        &header_sequence_number,
        &extra_args_compute_units,
        &extra_args_accounts_len,
        &extra_args_accounts,
        &header_nonce,
        &sender_size,
        &msg.sender,
        &data_size.to_be_bytes(),
        &msg.data,
    ]);

    result.to_bytes()
}

pub fn hash_solana2any(msg: &Solana2AnyRampMessage) -> [u8; 32] {
    // TODO: Modify this hash to be similar to the one in EVM
    // https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/libraries/Internal.sol#L129
    // Fixed-size message fields are included in nested hash to reduce stack pressure.
    // - metadata_hash =  sha256("Solana2AnyMessageHashV1", solana_chain_selector, dest_chain_selector, ccip_router_program_id))
    // - first_part = sha256(sender, sequence_number, nonce, fee_token, fee_token_amount)
    // - receiver
    // - message.data
    // - token_amounts
    // - extra_args

    use anchor_lang::solana_program::hash;

    // Push Data Size to ensure that the hash is unique
    let data_size = msg.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

    // RampMessageHeader struct
    let header_source_chain_selector = msg.header.source_chain_selector.to_be_bytes();
    let header_dest_chain_selector = msg.header.dest_chain_selector.to_be_bytes();
    let header_sequence_number = msg.header.sequence_number.to_be_bytes();
    let header_nonce = msg.header.nonce.to_be_bytes();

    // Extra Args struct
    let extra_args_gas_limit = msg.extra_args.gas_limit.to_be_bytes();
    let extra_args_allow_out_of_order_execution =
        [msg.extra_args.allow_out_of_order_execution as u8];

    // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
    let result = hash::hashv(&[
        &msg.sender.to_bytes(),
        &msg.receiver,
        &data_size.to_be_bytes(),
        &msg.data,
        &header_source_chain_selector,
        &header_dest_chain_selector,
        &header_sequence_number,
        &header_nonce,
        &extra_args_gas_limit,
        &extra_args_allow_out_of_order_execution,
    ]);

    result.to_bytes()
}

pub fn validate_solana2any(
    msg: &Solana2AnyMessage,
    dest_chain: &DestChain,
    token_config: &BillingTokenConfig,
) -> Result<()> {
    require!(
        dest_chain.config.is_enabled,
        CcipRouterError::DestinationChainDisabled
    );

    require!(token_config.enabled, CcipRouterError::FeeTokenDisabled);

    require_gte!(
        dest_chain.config.max_data_bytes,
        msg.data.len() as u32,
        CcipRouterError::MessageTooLarge
    );

    require_gte!(
        dest_chain.config.max_number_of_tokens_per_msg as usize,
        msg.token_amounts.len(),
        CcipRouterError::UnsupportedNumberOfTokens
    );

    validate_dest_family_address(msg, dest_chain.config.chain_family_selector)
}

pub fn validate_dest_family_address(
    msg: &Solana2AnyMessage,
    chain_family_selector: [u8; 4],
) -> Result<()> {
    const PRECOMPILE_SPACE: u32 = 1024;

    let selector = u32::from_be_bytes(chain_family_selector);
    // Only EVM is supported as a destination family.
    require_eq!(
        selector,
        CHAIN_FAMILY_SELECTOR_EVM,
        CcipRouterError::UnsupportedChainFamilySelector
    );

    require_eq!(msg.receiver.len(), 32, CcipRouterError::InvalidEVMAddress);

    let address: U256 = U256::from_be_bytes(
        msg.receiver
            .clone()
            .try_into()
            .map_err(|_| CcipRouterError::InvalidEncoding)?,
    );

    require!(address <= U160_MAX, CcipRouterError::InvalidEVMAddress);

    if let Ok(small_address) = TryInto::<u32>::try_into(address) {
        require_gte!(
            small_address,
            PRECOMPILE_SPACE,
            CcipRouterError::InvalidEVMAddress
        )
    };

    Ok(())
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;
    use crate::{
        utils::Exponential, RampMessageHeader, SolanaAccountMeta, SolanaExtraArgs,
        SolanaTokenAmount,
    };
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;
    use bytemuck::Zeroable;

    /// Builds a message and hash it, it's compared with a known hash
    #[test]
    fn test_hash() {
        let message = Any2SolanaRampMessage {
            sender: [
                1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0,
            ]
            .to_vec(),
            receiver: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb").unwrap(),
            data: vec![4, 5, 6],
            header: RampMessageHeader {
                message_id: [
                    8, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                    0, 0, 0, 0, 0, 0,
                ],
                source_chain_selector: 67,
                dest_chain_selector: 78,
                sequence_number: 89,
                nonce: 90,
            },
            token_amounts: [].to_vec(), // TODO: hash token amounts
            extra_args: SolanaExtraArgs {
                compute_units: 1000,
                accounts: vec![SolanaAccountMeta {
                    pubkey: Pubkey::try_from("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
                        .unwrap(),
                    is_writable: true,
                }],
            },
        };

        let on_ramp_address = &[1, 2, 3].to_vec();
        let hash_result = hash_any2solana(&message, on_ramp_address);

        assert_eq!(
            "03da97f96c82237d8a8ab0f68d4f7ba02afe188b4a876f348278fbf2226312ed",
            hex::encode(hash_result)
        );
    }

    #[test]
    fn message_not_validated_for_disabled_destination_chain() {
        let mut chain = sample_dest_chain();
        chain.config.is_enabled = false;

        assert_eq!(
            validate_solana2any(&sample_message(), &chain, &sample_billing_config()).unwrap_err(),
            CcipRouterError::DestinationChainDisabled.into()
        );
    }

    #[test]
    fn message_not_validated_for_disabled_token() {
        let mut billing_config = sample_billing_config();
        billing_config.enabled = false;

        assert_eq!(
            validate_solana2any(&sample_message(), &sample_dest_chain(), &billing_config)
                .unwrap_err(),
            CcipRouterError::FeeTokenDisabled.into()
        );
    }

    #[test]
    fn large_message_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.data = vec![0; dest_chain.config.max_data_bytes as usize + 1];
        assert_eq!(
            validate_solana2any(&message, &sample_dest_chain(), &sample_billing_config())
                .unwrap_err(),
            CcipRouterError::MessageTooLarge.into()
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
                validate_solana2any(&message, &sample_dest_chain(), &sample_billing_config())
                    .unwrap_err(),
                CcipRouterError::InvalidEVMAddress.into()
            );
        }
    }

    #[test]
    fn message_with_too_many_tokens_fails_to_validate() {
        let dest_chain = sample_dest_chain();
        let mut message = sample_message();
        message.token_amounts = vec![
            SolanaTokenAmount {
                token: Pubkey::new_unique(),
                amount: 1
            };
            dest_chain.config.max_number_of_tokens_per_msg as usize + 1
        ];
        assert_eq!(
            validate_solana2any(&message, &sample_dest_chain(), &sample_billing_config())
                .unwrap_err(),
            CcipRouterError::UnsupportedNumberOfTokens.into()
        );
    }

    pub fn sample_message() -> Solana2AnyMessage {
        let mut receiver = vec![0u8; 32];

        // Arbitrary value that pushes the address to the right EVM range
        // (above precompile space, under u160::max)
        receiver[20] = 0xA;

        Solana2AnyMessage {
            receiver,
            data: vec![],
            token_amounts: vec![],
            fee_token: Pubkey::zeroed(),
            extra_args: crate::ExtraArgsInput {
                gas_limit: None,
                allow_out_of_order_execution: None,
            },
            token_indexes: vec![],
        }
    }

    pub fn sample_billing_config() -> BillingTokenConfig {
        let mut value = [0; 28];
        value.clone_from_slice(&3u32.e(18).to_be_bytes()[4..]);
        BillingTokenConfig {
            enabled: true,
            mint: native_mint::ID,
            usd_per_token: crate::TimestampedPackedU224 {
                value,
                timestamp: 100,
            },
            premium_multiplier_wei_per_eth: 1,
        }
    }

    pub fn sample_dest_chain() -> DestChain {
        let mut value = [0; 28];
        // L1 gas price
        value[0..14].clone_from_slice(&1u32.e(18).to_be_bytes()[18..]);
        // L2 gas price
        value[14..].clone_from_slice(&U256::new(22u128).to_be_bytes()[18..]);
        DestChain {
            version: 1,
            chain_selector: 1,
            state: crate::DestChainState {
                sequence_number: 0,
                usd_per_unit_gas: crate::TimestampedPackedU224 {
                    value,
                    timestamp: 100,
                },
            },
            config: crate::DestChainConfig {
                is_enabled: true,
                max_number_of_tokens_per_msg: 5,
                max_data_bytes: 200,
                max_per_msg_gas_limit: 0,
                dest_gas_overhead: 1,
                dest_gas_per_payload_byte: 0,
                dest_data_availability_overhead_gas: 0,
                dest_gas_per_data_availability_byte: 1,
                dest_data_availability_multiplier_bps: 1,
                default_token_fee_usdcents: 100,
                default_token_dest_gas_overhead: 0,
                default_tx_gas_limit: 0,
                gas_multiplier_wei_per_eth: 1,
                network_fee_usdcents: 100,
                gas_price_staleness_threshold: 10,
                enforce_out_of_order: false,
                chain_family_selector: CHAIN_FAMILY_SELECTOR_EVM.to_be_bytes(),
            },
        }
    }
}
