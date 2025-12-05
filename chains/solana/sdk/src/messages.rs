use borsh::{BorshDeserialize, BorshSerialize};
use solana_sdk::pubkey::Pubkey;

/// Message struct for SVM -> Any chain.
#[derive(BorshSerialize, BorshDeserialize, Clone, Debug, Default, PartialEq)]
pub struct SVM2AnyMessage {
    /// Receiver address (format depends on destination chain).
    pub receiver: Vec<u8>,
    /// Arbitrary payload data.
    pub data: Vec<u8>,
    /// List of tokens and amounts attached to the message.
    pub token_amounts: Vec<SVMTokenAmount>,
    /// Fee token mint (or `Pubkey::default()` for native SOL).
    pub fee_token: Pubkey,
    /// Extra args (tag + borsh-encoded struct, e.g. EVMExtraArgsV2).
    pub extra_args: Vec<u8>,
}

/// Single token + amount tuple.
#[derive(BorshSerialize, BorshDeserialize, Clone, Debug, Default, PartialEq)]
pub struct SVMTokenAmount {
    /// Token mint or native address.
    pub token: Pubkey,
    /// Amount to transfer.
    pub amount: u64,
}

/// Args struct for `ccip_send`.
#[derive(BorshSerialize, BorshDeserialize, Clone, Debug, Default, PartialEq)]
pub struct CcipSendArgs {
    /// Destination chain selector (unique numeric ID for the target chain).
    pub dest_chain_selector: u64,
    /// Message containing receiver, data, tokens, etc.
    pub message: SVM2AnyMessage,
    /// Indexes pointing into `remaining_accounts` that correspond to each token.
    pub token_indexes: Vec<u8>,
}


/// Tag to indicate a gas limit (or dest chain equivalent processing units) and Out Of Order Execution. This tag is
/// available for multiple chain families. If there is no chain family specific tag, this is the default available
/// for a chain.
///
/// Note: not available for Solana VM based chains.
///
/// Fully compatible with the previously existing EVMExtraArgsV2, hence the string used to generate it.
/// bytes4(keccak256("CCIP EVMExtraArgsV2"));
pub const GENERIC_EXTRA_ARGS_V2_TAG: u32 = 0x181dcf10;

/// bytes4(keccak256("CCIP SVMExtraArgsV1"));
pub const SVM_EXTRA_ARGS_V1_TAG: u32 = 0x1f3b3aba;


// Extra Args of message when SVM -> SVM
#[derive(BorshSerialize, BorshDeserialize, Clone, Debug, Default, PartialEq)]
pub struct SVMExtraArgsV1 {
    pub compute_units: u32, // compute units are used by the offchain to add computeUnitLimits to the execute stage
    pub account_is_writable_bitmap: u64,
    pub allow_out_of_order_execution: bool, // SVM chains require OOO, but users are required to explicitly state OOO is allowed in extraArgs
    pub token_receiver: [u8; 32],           // cannot be 0-address if tokens are sent in message
    pub accounts: Vec<[u8; 32]>,            // account list for messaging on remote SVM chain
}

#[derive(BorshSerialize, BorshDeserialize, Clone, Debug, Default, PartialEq)]
pub struct GenericExtraArgsV2 {
    pub gas_limit: u128,                    // message gas limit
    pub allow_out_of_order_execution: bool, // user configurable OOO execution that must match with the DestChainConfig.EnforceOutOfOrderExecution
}
