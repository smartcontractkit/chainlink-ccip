use anchor_lang::prelude::*;

use crate::{
    messages::{
        CrossChainAmount, SVMTokenAmount, CCIP_RECEIVE_DISCRIMINATOR,
        TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR,
    },
    CcipOfframpError,
};

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub(super) struct ReleaseOrMintInV1 {
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

impl ReleaseOrMintInV1 {
    pub fn to_tx_data(&self) -> Vec<u8> {
        let mut data = Vec::new();
        data.extend_from_slice(&TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR);
        data.extend_from_slice(&self.try_to_vec().unwrap());
        data
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub(super) struct ReleaseOrMintOutV1 {
    pub destination_amount: u64, // TODO: u256 on EVM?
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub(super) struct Any2SVMMessage {
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
        require!(m.is_ok(), CcipOfframpError::InvalidMessage);
        let message = m.unwrap();

        let mut data = Vec::with_capacity(8 + message.len());
        data.extend_from_slice(&CCIP_RECEIVE_DISCRIMINATOR);
        data.extend_from_slice(&message);

        Ok(data)
    }
}

pub fn is_writable(bitmap: &u64, index: u8) -> bool {
    index < 64 && (bitmap & 1 << index != 0) // check valid index and that bit at index is 1
}
