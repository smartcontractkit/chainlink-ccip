use anchor_lang::prelude::*;

// zero_copy is used to prevent hitting stack/heap memory limits
#[account(zero_copy)]
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize)]
pub struct Config {
    pub version: u8,
    pub default_allow_out_of_order_execution: u8, // bytemuck::Pod compliant required for zero_copy
    _padding0: [u8; 6],
    pub svm_chain_selector: u64,
    pub default_gas_limit: u128,
    _padding1: [u8; 8],

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    pub enable_manual_execution_after: i64, // Expressed as Unix time (i.e. seconds since the Unix epoch).
    _padding2: [u8; 8],

    pub ocr3: [Ocr3Config; 2],
    // TODO: token pool global config

    // TODO: billing global configs'
    pub fee_aggregator: Pubkey, // Allowed address to withdraw billed fees to (will use ATAs derived from it)
}

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct Ocr3ConfigInfo {
    pub config_digest: [u8; 32], // 32-byte hash of configuration
    pub f: u8,                   // f+1 = number of signatures per report
    pub n: u8,                   // number of signers
    pub is_signature_verification_enabled: u8, // bool -> bytemuck::Pod compliant required for zero_copy
}

// TODO: do we need to verify signers and transmitters are different? (between the two groups)
// signers: pubkey is 20-byte address, secp256k1 curve ECDSA
// transmitters: 32-byte pubkey, ed25519

#[zero_copy]
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Default)]
pub struct Ocr3Config {
    pub plugin_type: u8, // plugin identifier for validation (example: ccip:commit = 0, ccip:execute = 1)
    pub config_info: Ocr3ConfigInfo,
    pub signers: [[u8; 20]; 16], // v0.29.0 - anchor IDL does not build with MAX_SIGNERS
    pub transmitters: [[u8; 32]; 16], // v0.29.0 - anchor IDL does not build with MAX_TRANSMITTERS
}

impl Ocr3Config {
    pub fn new(plugin_type: u8) -> Self {
        Self {
            plugin_type,
            ..Default::default()
        }
    }
}

#[account]
#[derive(InitSpace)]
pub struct GlobalState {
    // This holds global variables for the contract that are not manually set by the admin.
    // They are auto-updated by the contract during regular usage of CCIP.
    pub latest_price_sequence_number: u64,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainConfig {
    pub is_enabled: bool, // Flag whether the source chain is enabled or not
    // OnRamp addresses supported from the source chain, each of them has a 64 byte address. So this can hold 2 addresses.
    // If only one address is configured, then the space for the second address must be zeroed.
    // Each address must be right padded with zeros if it is less than 64 bytes.
    pub on_ramp: [[u8; 64]; 2],
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct SourceChainState {
    pub min_seq_nr: u64, // The min sequence number expected for future messages
}

#[account]
#[derive(InitSpace, Debug)]
pub struct SourceChain {
    // Config for Any2SVM
    pub version: u8,
    pub chain_selector: u64,       // Chain selector used for the seed
    pub state: SourceChainState,   // values that are updated automatically
    pub config: SourceChainConfig, // values configured by an admin
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct DestChainState {
    pub sequence_number: u64, // The last used sequence number
    pub usd_per_unit_gas: TimestampedPackedU224,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, InitSpace, Debug)]
pub struct DestChainConfig {
    pub is_enabled: bool, // Whether this destination chain is enabled

    pub max_number_of_tokens_per_msg: u16, // Maximum number of distinct ERC20 token transferred per message
    pub max_data_bytes: u32,               // Maximum payload data size in bytes
    pub max_per_msg_gas_limit: u32,        // Maximum gas limit for messages targeting EVMs
    pub dest_gas_overhead: u32, //  Gas charged on top of the gasLimit to cover destination chain costs
    pub dest_gas_per_payload_byte: u16, // Destination chain gas charged for passing each byte of `data` payload to receiver
    pub dest_data_availability_overhead_gas: u32, // Extra data availability gas charged on top of the message, e.g. for OCR
    pub dest_gas_per_data_availability_byte: u16, // Amount of gas to charge per byte of message data that needs availability
    pub dest_data_availability_multiplier_bps: u16, // Multiplier for data availability gas, multiples of bps, or 0.0001

    // The following three properties are defaults, they can be overridden by setting the TokenTransferFeeConfig for a token
    pub default_token_fee_usdcents: u16, // Default token fee charged per token transfer
    pub default_token_dest_gas_overhead: u32, //  Default gas charged to execute the token transfer on the destination chain
    pub default_tx_gas_limit: u32,            // Default gas limit for a tx
    pub gas_multiplier_wei_per_eth: u64, // Multiplier for gas costs, 1e18 based so 11e17 = 10% extra cost.
    pub network_fee_usdcents: u32, // Flat network fee to charge for messages, multiples of 0.01 USD
    pub gas_price_staleness_threshold: u32, // The amount of time a gas price can be stale before it is considered invalid (0 means disabled)
    pub enforce_out_of_order: bool, // Whether to enforce the allowOutOfOrderExecution extraArg value to be true.
    pub chain_family_selector: [u8; 4], // Selector that identifies the destination chain's family. Used to determine the correct validations to perform for the dest chain.
}

#[account]
#[derive(InitSpace, Debug)]
pub struct DestChain {
    // Config for SVM2Any
    pub version: u8,
    pub chain_selector: u64,     // Chain selector used for the seed
    pub state: DestChainState,   // values that are updated automatically
    pub config: DestChainConfig, // values configured by an admin
}

#[account]
#[derive(InitSpace)]
pub struct Nonce {
    pub version: u8,  // version to check if nonce account is already initialized
    pub counter: u64, // Counter per user and per lane to use as nonce for all the messages to be executed in order
}

#[account]
#[derive(InitSpace)]
pub struct ExternalExecutionConfig {}

#[account]
#[derive(InitSpace)]
pub struct CommitReport {
    pub version: u8,
    pub chain_selector: u64,
    pub merkle_root: [u8; 32],
    pub timestamp: i64, // Expressed as Unix time (i.e. seconds since the Unix epoch).
    pub min_msg_nr: u64,
    pub max_msg_nr: u64, // TODO: Change this to [u128; 2] when supporting commit reports with 256 messages
    pub execution_states: u128,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct PerChainPerTokenConfig {
    pub version: u8,         // schema version
    pub chain_selector: u64, // remote chain
    pub mint: Pubkey,        // token on solana

    pub billing: TokenBilling, // EVM: configurable in router only by ccip admins
}

#[derive(InitSpace, Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct TokenBilling {
    pub min_fee_usdcents: u32, // Minimum fee to charge per token transfer, multiples of 0.01 USD
    pub max_fee_usdcents: u32, // Maximum fee to charge per token transfer, multiples of 0.01 USD
    pub deci_bps: u16, // Basis points charged on token transfers, multiples of 0.1bps, or 1e-5
    pub dest_gas_overhead: u32, // Gas charged to execute the token transfer on the destination chain
    // Extra data availability bytes that are returned from the source pool and sent
    pub dest_bytes_overhead: u32, // to the destination pool. Must be >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
    pub is_enabled: bool,         // Whether this token has custom transfer fees
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct RateLimitTokenBucket {
    pub tokens: u128,      // Current number of tokens that are in the bucket.
    pub last_updated: u32, // Timestamp in seconds of the last token refill, good for 100+ years.
    pub is_enabled: bool,  // Indication whether the rate limiting is enabled or not
    pub capacity: u128,    // Maximum number of tokens that can be in the bucket.
    pub rate: u128,        // Number of tokens per second that the bucket is refilled.
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize, Debug)]
pub struct BillingTokenConfig {
    // NOTE: when modifying this struct, make sure to update the version in the wrapper
    pub enabled: bool,
    pub mint: Pubkey,

    // price tracking
    pub usd_per_token: TimestampedPackedU224,
    // billing configs
    pub premium_multiplier_wei_per_eth: u64,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct BillingTokenConfigWrapper {
    pub version: u8,
    pub config: BillingTokenConfig,
}

#[derive(InitSpace, Clone, AnchorSerialize, AnchorDeserialize, Debug)]
pub struct TimestampedPackedU224 {
    pub value: [u8; 28],
    pub timestamp: i64, // maintaining the type that SVM returns for the time (solana_program::clock::UnixTimestamp = i64)
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize, Debug, PartialEq)]
// used in the commit report execution_states field
pub enum MessageExecutionState {
    Untouched = 0,
    InProgress = 1, // Not used in SVM, but used in EVM
    Success = 2,
    Failure = 3,
}

impl TryFrom<u128> for MessageExecutionState {
    type Error = &'static str;

    fn try_from(value: u128) -> std::result::Result<MessageExecutionState, &'static str> {
        match value {
            0 => Ok(MessageExecutionState::Untouched),
            1 => Ok(MessageExecutionState::InProgress),
            2 => Ok(MessageExecutionState::Success),
            3 => Ok(MessageExecutionState::Failure),
            _ => Err("Invalid ExecutionState"),
        }
    }
}

#[cfg(test)]
mod tests {

    use super::*;
    use std::convert::TryFrom;

    #[test]
    fn test_execution_state_try_from() {
        assert_eq!(
            MessageExecutionState::try_from(0).unwrap(),
            MessageExecutionState::Untouched
        );
        assert_eq!(
            MessageExecutionState::try_from(1).unwrap(),
            MessageExecutionState::InProgress
        );
        assert_eq!(
            MessageExecutionState::try_from(2).unwrap(),
            MessageExecutionState::Success
        );
        assert_eq!(
            MessageExecutionState::try_from(3).unwrap(),
            MessageExecutionState::Failure
        );
        assert!(MessageExecutionState::try_from(4).is_err());
    }
}
