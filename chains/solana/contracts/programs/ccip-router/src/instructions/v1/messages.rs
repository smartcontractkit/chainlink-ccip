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
}
