pub mod pools {
    use crate::{
        CrossChainAmount, TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR,
        TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR,
    };

    use super::super::pools::ToTxData;
    use anchor_lang::prelude::*;

    #[derive(Clone, AnchorSerialize, AnchorDeserialize)]
    pub(in super::super) struct LockOrBurnInV1 {
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
    pub(in super::super) struct ReleaseOrMintInV1 {
        pub original_sender: Vec<u8>, //          The original sender of the tx on the source chain
        pub remote_chain_selector: u64, // ─╮ The chain ID of the source chain
        pub receiver: Pubkey, // ───────────╯ The Token Associated Account that will receive the tokens on the destination chain.
        pub amount: CrossChainAmount, // LE u256 amount - The amount of tokens to release or mint, denominated in the source token's decimals, pool expected to handle conversation to solana token specifics
        pub local_token: Pubkey,      //            The address of the Token Mint Account on SVM
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
    pub(in super::super) struct LockOrBurnOutV1 {
        pub dest_token_address: Vec<u8>,
        pub dest_pool_data: Vec<u8>,
    }

    #[derive(Clone, AnchorSerialize, AnchorDeserialize)]
    pub(in super::super) struct ReleaseOrMintOutV1 {
        pub destination_amount: u64, // TODO: u256 on EVM?
    }
}

pub mod ramps {
    use anchor_lang::prelude::*;

    use crate::{CcipRouterError, SVMTokenAmount, CCIP_RECEIVE_DISCRIMINATOR};

    #[derive(Clone, AnchorSerialize, AnchorDeserialize)]
    pub(in super::super) struct Any2SVMMessage {
        pub message_id: [u8; 32],
        pub source_chain_selector: u64,
        pub sender: Vec<u8>,
        pub data: Vec<u8>,
        pub token_amounts: Vec<SVMTokenAmount>,
    }

    /// Build the instruction data (discriminator + any other data)
    impl Any2SVMMessage {
        pub fn build_receiver_discriminator_and_data(&self) -> Result<Vec<u8>> {
            let m: std::result::Result<Vec<u8>, std::io::Error> = self.try_to_vec();
            require!(m.is_ok(), CcipRouterError::InvalidMessage);
            let message = m.unwrap();

            let mut data = Vec::with_capacity(8);
            data.extend_from_slice(&CCIP_RECEIVE_DISCRIMINATOR);
            data.extend_from_slice(&message);

            Ok(data)
        }
    }

    pub fn is_writable(bitmap: &u64, index: u8) -> bool {
        index < 64 && (bitmap & 1 << index != 0) // check valid index and that bit at index is 1
    }

    // TODO move this as it is fairly unrelated to other things in this file now
    #[cfg(test)]
    pub mod tests {
        use crate::state::{BillingTokenConfig, DestChain};
        use crate::TimestampedPackedU224;
        use anchor_spl::token::spl_token::native_mint;
        use ethnum::U256;

        pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;

        fn as_u8_28(single: U256) -> [u8; 28] {
            single.to_be_bytes()[4..32].try_into().unwrap()
        }

        // pub fn sample_message() -> SVM2AnyMessage {
        //     let mut receiver = vec![0u8; 32];

        //     // Arbitrary value that pushes the address to the right EVM range
        //     // (above precompile space, under u160::max)
        //     receiver[20] = 0xA;

        //     SVM2AnyMessage {
        //         receiver,
        //         data: vec![],
        //         token_amounts: vec![],
        //         fee_token: native_mint::ID,
        //         extra_args: ExtraArgsInput {
        //             gas_limit: None,
        //             allow_out_of_order_execution: None,
        //         },
        //     }
        // }

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
                state: crate::DestChainState {
                    sequence_number: 0,
                    usd_per_unit_gas,
                },
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
}
