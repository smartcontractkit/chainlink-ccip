use anchor_lang::prelude::*;
use ethnum::U256;

use crate::{
    BillingTokenConfig, CcipRouterError, DestChain, Solana2AnyMessage, CHAIN_FAMILY_SELECTOR_EVM,
    TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR, TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR,
};

use super::pools::ToTxData;

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

///////////////////
// Ramp Messages //
///////////////////

const U160_MAX: U256 = U256::from_words(u32::MAX as u128, u128::MAX);

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

fn validate_dest_family_address(
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
    use crate::{utils::Exponential, SolanaTokenAmount};
    use anchor_lang::solana_program::pubkey::Pubkey;
    use anchor_spl::token::spl_token::native_mint;

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
            fee_token: Pubkey::default(),
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
