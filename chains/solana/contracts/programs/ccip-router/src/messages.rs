use std::convert::Into;

use anchor_lang::prelude::*;

// https://github.com/smartcontractkit/chainlink/blob/ff8a597fd9df653f8967427498eaa5a04b19febb/contracts/src/v0.8/ccip/libraries/Internal.sol#L276
pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;
pub const CHAIN_FAMILY_SELECTOR_SVM: u32 = 0x1e10bdc4;

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
    pub message: Any2SVMRampMessage,
    pub offchain_token_data: Vec<Vec<u8>>, // https://github.com/smartcontractkit/chainlink/blob/885baff9479e935e0fc34d9f52214a32c158eac5/contracts/src/v0.8/ccip/libraries/Internal.sol#L72
    pub root: [u8; 32],
    pub proofs: Vec<[u8; 32]>,
}

// Any2SVMRampExtraArgs is used during the execute or manual execute calls (offramp only)
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMRampExtraArgs {
    pub compute_units: u32,
    pub is_writable_bitmap: u64, // part of the message to avoid calculating it onchain
}

impl Any2SVMRampExtraArgs {
    pub fn len(&self) -> usize {
        4 // compute units
        + 8 // isWritable bitmap
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMRampMessage {
    pub header: RampMessageHeader,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    // Token Receiver is the address which the ATA will be calculated from.
    // If token receiver and message execution, then the token receiver must be a PDA from the logic receiver
    // (Logic receiver is passed into relevant instructions through `remaining_accounts`)
    pub token_receiver: Pubkey,
    pub token_amounts: Vec<Any2SVMTokenTransfer>,
    pub extra_args: Any2SVMRampExtraArgs,
    pub on_ramp_address: Vec<u8>,
}

impl Any2SVMRampMessage {
    pub fn len(&self) -> usize {
        let token_len = self.token_amounts.iter().fold(0, |acc, e| acc + e.len());

        self.header.len() // header
        + 4 + self.sender.len() // sender
        + 4 + self.data.len() // data
        + 32 // token receiver
        + 4 + token_len // token_amount
        + self.extra_args.len() // extra_args
        + 4 + self.on_ramp_address.len() // on_ramp_address
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
// Family-agnostic message emitted from the OnRamp
// Note: hash(Any2SVMRampMessage) != hash(SVM2AnyRampMessage) due to encoding & parameter differences
// messageId = hash(SVM2AnyRampMessage) using the source EVM chain's encoding format
pub struct SVM2AnyRampMessage {
    pub header: RampMessageHeader, // Message header
    pub sender: Pubkey,            // sender address on the source chain
    pub data: Vec<u8>,             // arbitrary data payload supplied by the message sender
    pub receiver: Vec<u8>,         // receiver address on the destination chain
    pub extra_args: Vec<u8>, // destination-chain specific extra args, such as the gasLimit for EVM chains
    pub fee_token: Pubkey,
    pub token_amounts: Vec<SVM2AnyTokenTransfer>,
    pub fee_token_amount: CrossChainAmount,
    pub fee_value_juels: CrossChainAmount,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVM2AnyTokenTransfer {
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
    pub amount: CrossChainAmount, // LE encoded u256 -  cross-chain token amount is always u256
    // Destination chain data used to execute the token transfer on the destination chain. For an EVM destination, it
    // consists of the amount of gas available for the releaseOrMint and transfer calls made by the offRamp.
    pub dest_exec_data: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMTokenTransfer {
    // The source pool address encoded to bytes. This value is trusted as it is obtained through the onRamp. It can be
    // relied upon by the destination pool to validate the source pool.
    pub source_pool_address: Vec<u8>,
    pub dest_token_address: Pubkey, // Address of destination token
    pub dest_gas_amount: u32, // The amount of gas available for the releaseOrMint and transfer calls on the offRamp.
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    pub extra_data: Vec<u8>,
    pub amount: CrossChainAmount,
}

impl Any2SVMTokenTransfer {
    pub fn len(&self) -> usize {
        4 + self.source_pool_address.len() // source_pool
        + 32 // token_address
        + 4  // gas_amount
        + 4 + self.extra_data.len()  // extra_data
        + 32 // amount
    }
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct SVM2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, Default, Debug)]
pub struct CrossChainAmount {
    le_bytes: [u8; 32],
}

impl CrossChainAmount {
    pub fn to_bytes(self) -> [u8; 32] {
        self.le_bytes
    }
}

impl<T: Into<ethnum::U256>> From<T> for CrossChainAmount {
    fn from(value: T) -> Self {
        Self {
            le_bytes: value.into().to_le_bytes(),
        }
    }
}
