use anchor_lang::prelude::*;

pub const CHAIN_FAMILY_SELECTOR_EVM: u32 = 0x2812d52c;

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
pub struct AnyExtraArgs {
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
    pub on_ramp_address: Vec<u8>,
}

impl Any2SolanaRampMessage {
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
    pub extra_args: AnyExtraArgs, // destination-chain specific extra args, such as the gasLimit for EVM chains
    pub fee_token: Pubkey,
    pub token_amounts: Vec<Solana2AnyTokenTransfer>,
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
pub struct Solana2AnyMessage {
    pub receiver: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
    pub fee_token: Pubkey, // pass zero address if native SOL
    pub extra_args: ExtraArgsInput,

    // solana specific parameter for mapping tokens to set of accounts
    pub token_indexes: Vec<u8>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Default, Debug, PartialEq, Eq)]
pub struct SolanaTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // u64 - amount local to solana
}

#[derive(Clone, Copy, AnchorSerialize, AnchorDeserialize)]
pub struct ExtraArgsInput {
    pub gas_limit: Option<u128>,
    pub allow_out_of_order_execution: Option<bool>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SolanaMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SolanaTokenAmount>,
}
