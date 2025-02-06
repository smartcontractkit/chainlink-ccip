pub mod pools {
    use crate::TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR;

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
    pub(in super::super) struct LockOrBurnOutV1 {
        pub dest_token_address: Vec<u8>,
        pub dest_pool_data: Vec<u8>,
    }

    #[derive(Clone, AnchorSerialize, AnchorDeserialize)]
    pub(in super::super) struct ReleaseOrMintOutV1 {
        pub destination_amount: u64, // TODO: u256 on EVM?
    }
}
