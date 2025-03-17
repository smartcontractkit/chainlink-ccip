use anchor_lang::prelude::*;

pub const CCIP_RECEIVE_DISCRIMINATOR: [u8; 8] = [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]; // ccip_receive
pub const TOKENPOOL_RELEASE_OR_MINT_DISCRIMINATOR: [u8; 8] =
    [0x5c, 0x64, 0x96, 0xc6, 0xfc, 0x3f, 0xa4, 0xe4]; // release_or_mint_tokens

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
/// Report that is submitted by the execution DON at the execution phase. (including chain selector data)
pub struct ExecutionReportSingleChain {
    pub source_chain_selector: u64,
    pub message: Any2SVMRampMessage,
    pub offchain_token_data: Vec<Vec<u8>>, // https://github.com/smartcontractkit/chainlink/blob/885baff9479e935e0fc34d9f52214a32c158eac5/contracts/src/v0.8/ccip/libraries/Internal.sol#L72
    pub proofs: Vec<[u8; 32]>,
}

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
    }
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

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, Default, Debug)]
pub struct CrossChainAmount {
    le_bytes: [u8; 32],
}

impl<T: Into<ethnum::U256>> From<T> for CrossChainAmount {
    fn from(value: T) -> Self {
        Self {
            le_bytes: value.into().to_le_bytes(),
        }
    }
}
