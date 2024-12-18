use crate::{ReportContext, TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR};
use anchor_lang::prelude::*;

use crate::{ocr3base::Ocr3Report, ToTxData, TOKENPOOL_LOCK_OR_BURN_DISCRIMINATOR};

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
// Family-agnostic header for OnRamp & OffRamp messages.
// The messageId is not expected to match hash(message), since it may originate from another ramp family
pub struct RampMessageHeader {
    pub message_id: [u8; 32], // Unique identifier for the message, generated with the source chain's encoding scheme
    pub source_chain_selector: u64, // the chain selector of the source chain, note: not chainId
    pub dest_chain_selector: u64, // the chain selector of the destination chain, note: not chainId
    pub sequence_number: u64, // sequence number, not unique across lanes
    pub nonce: u64, // nonce for this lane for this sender, not unique across senders/lanes
}

impl RampMessageHeader {
    pub fn len(&self) -> usize {
        32 // message_id
        + 8 // source_chain
        + 8 // dest_chain
        + 8 // sequence
        + 8 // nonce
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
/// Report that is submitted by the execution DON at the execution phase. (including chain selector data)
pub struct ExecutionReportSingleChain {
    pub source_chain_selector: u64,
    pub message: Any2SolanaRampMessage,
    pub offchain_token_data: Vec<Vec<u8>>, // https://github.com/smartcontractkit/chainlink/blob/885baff9479e935e0fc34d9f52214a32c158eac5/contracts/src/v0.8/ccip/libraries/Internal.sol#L72
    pub root: [u8; 32],
    pub proofs: Vec<[u8; 32]>,

    // NOT HASHED
    pub token_indexes: Vec<u8>, // outside of message because this is not available during commit stage
}

impl Ocr3Report for ExecutionReportSingleChain {
    fn hash(&self, _: &ReportContext) -> [u8; 32] {
        [0; 32] // not needed, this report is not hashed for signing
    }
    fn len(&self) -> usize {
        let offchain_token_data_len = self
            .offchain_token_data
            .iter()
            .fold(0, |acc, e| acc + 4 + e.len());

        8 // source chain selector
        + self.message.len() // ccip message
        + 4 + offchain_token_data_len// offchain_token_data
        + 32 // root
        + 4 + self.proofs.len() * 32 // count + proofs
        + 4 + self.token_indexes.len() // token_indexes
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace)]
pub struct SolanaAccountMeta {
    pub pubkey: Pubkey,
    pub is_writable: bool,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct SolanaExtraArgs {
    pub compute_units: u32,
    pub accounts: Vec<SolanaAccountMeta>,
}

impl SolanaExtraArgs {
    pub fn len(&self) -> usize {
        4 // compute units
        + 4 + self.accounts.len() * SolanaAccountMeta::INIT_SPACE // additional accounts
    }
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct EvmExtraArgs {
    pub gas_limit: u128,
    pub allow_out_of_order_execution: bool,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaRampMessage {
    pub header: RampMessageHeader,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    // receiver is used as the target for the two main functionalities
    // token transfers: recipient of token transfers (associated token addresses are validated against this address)
    // arbitrary messaging: expected account in the declared arbitrary messaging accounts (2nd in the list of the accounts)
    pub receiver: Pubkey,
    pub token_amounts: Vec<Any2SolanaTokenTransfer>,
    pub extra_args: SolanaExtraArgs,
}

impl Any2SolanaRampMessage {
    pub fn hash(&self, on_ramp_address: &[u8]) -> [u8; 32] {
        use anchor_lang::solana_program::hash;

        // Calculate vectors size to ensure that the hash is unique
        let sender_size = [self.sender.len() as u8];
        let on_ramp_address_size = [on_ramp_address.len() as u8];
        let data_size = self.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

        // RampMessageHeader struct
        let header_source_chain_selector = self.header.source_chain_selector.to_be_bytes();
        let header_dest_chain_selector = self.header.dest_chain_selector.to_be_bytes();
        let header_sequence_number = self.header.sequence_number.to_be_bytes();
        let header_nonce = self.header.nonce.to_be_bytes();

        // Extra Args struct
        let extra_args_compute_units = self.extra_args.compute_units.to_be_bytes();
        let extra_args_accounts_len = [self.extra_args.accounts.len() as u8];
        let extra_args_accounts = self.extra_args.accounts.try_to_vec().unwrap();

        // TODO: Hash token amounts

        // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
        // As similar as https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L402
        let result = hash::hashv(&[
            "Any2SolanaMessageHashV1".as_bytes(),
            &header_source_chain_selector,
            &header_dest_chain_selector,
            &on_ramp_address_size,
            on_ramp_address,
            &self.header.message_id,
            &self.receiver.to_bytes(),
            &header_sequence_number,
            &extra_args_compute_units,
            &extra_args_accounts_len,
            &extra_args_accounts,
            &header_nonce,
            &sender_size,
            &self.sender,
            &data_size.to_be_bytes(),
            &self.data,
        ]);

        result.to_bytes()
    }

    pub fn len(&self) -> usize {
        let token_len = self.token_amounts.iter().fold(0, |acc, e| acc + e.len());

        self.header.len() // header
        + 4 + self.sender.len() // sender
        + 4 + self.data.len() // data
        + 32 // receiver
        + 4 + token_len // token_amount
        + self.extra_args.len() // extra_args
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
// Family-agnostic message emitted from the OnRamp
// Note: hash(Any2SolanaRampMessage) != hash(Solana2AnyRampMessage) due to encoding & parameter differences
// messageId = hash(Solana2AnyRampMessage) using the source EVM chain's encoding format
pub struct Solana2AnyRampMessage {
    pub header: RampMessageHeader, // Message header
    pub sender: Pubkey,            // sender address on the source chain
    pub data: Vec<u8>,             // arbitrary data payload supplied by the message sender
    pub receiver: Vec<u8>,         // receiver address on the destination chain
    pub extra_args: EvmExtraArgs, // destination-chain specific extra args, such as the gasLimit for EVM chains
    pub fee_token: Pubkey,
    pub token_amounts: Vec<Solana2AnyTokenTransfer>,
}

impl Solana2AnyRampMessage {
    pub fn hash(&self) -> [u8; 32] {
        use anchor_lang::solana_program::hash;

        // Push Data Size to ensure that the hash is unique
        let data_size = self.data.len() as u16; // u16 > maximum transaction size, u8 may have overflow

        // RampMessageHeader struct
        let header_source_chain_selector = self.header.source_chain_selector.to_be_bytes();
        let header_dest_chain_selector = self.header.dest_chain_selector.to_be_bytes();
        let header_sequence_number = self.header.sequence_number.to_be_bytes();
        let header_nonce = self.header.nonce.to_be_bytes();

        // Extra Args struct
        let extra_args_gas_limit = self.extra_args.gas_limit.to_be_bytes();
        let extra_args_allow_out_of_order_execution =
            [self.extra_args.allow_out_of_order_execution as u8];

        // NOTE: calling hash::hashv is orders of magnitude cheaper than using Hasher::hashv
        let result = hash::hashv(&[
            &self.sender.to_bytes(),
            &self.receiver,
            &data_size.to_be_bytes(),
            &self.data,
            &header_source_chain_selector,
            &header_dest_chain_selector,
            &header_sequence_number,
            &header_nonce,
            &extra_args_gas_limit,
            &extra_args_allow_out_of_order_execution,
        ]);

        result.to_bytes()
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct Solana2AnyTokenTransfer {
    // The source pool address. This value is trusted as it was obtained through the onRamp. It can be relied
    // upon by the destination pool to validate the source pool.
    pub source_pool_address: Pubkey,
    // The address of the destination token.
    // This value is UNTRUSTED as any pool owner can return whatever value they want.
    pub dest_token_address: Vec<u8>,
    // Optional pool data to be transferred to the destination chain. By default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    pub extra_data: Vec<u8>,
    pub amount: [u8; 32], // LE encoded u256 -  cross-chain token amount is always u256
    // Destination chain data used to execute the token transfer on the destination chain. For an EVM destination, it
    // consists of the amount of gas available for the releaseOrMint and transfer calls made by the offRamp.
    pub dest_exec_data: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaTokenTransfer {
    // The source pool address encoded to bytes. This value is trusted as it is obtained through the onRamp. It can be
    // relied upon by the destination pool to validate the source pool.
    pub source_pool_address: Vec<u8>,
    pub dest_token_address: Pubkey, // Address of destination token
    pub dest_gas_amount: u32, // The amount of gas available for the releaseOrMint and transfer calls on the offRamp.
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    pub extra_data: Vec<u8>,
    pub amount: [u8; 32], // LE encoded u256, any cross-chain token amounts are u256
}

impl Any2SolanaTokenTransfer {
    pub fn len(&self) -> usize {
        4 + self.source_pool_address.len() // source_pool
        + 32 // token_address
        + 4  // gas_amount
        + 4 + self.extra_data.len()  // extra_data
        + 32 // amount
    }
}

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

#[cfg(test)]
mod tests {

    use super::*;
    use anchor_lang::solana_program::pubkey::Pubkey;

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
        let hash_result = message.hash(on_ramp_address);

        assert_eq!(
            "03da97f96c82237d8a8ab0f68d4f7ba02afe188b4a876f348278fbf2226312ed",
            hex::encode(hash_result)
        );
    }
}
