// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// End consumer library.
library Client {
  struct EVMTokenAmount {
    address token; // token address on the local chain.
    uint256 amount; // Amount of tokens.
  }

  struct Any2EVMMessage {
    bytes32 messageId; // MessageId corresponding to ccipSend on source.
    uint64 sourceChainSelector; // Source chain selector.
    bytes sender; // abi.decode(sender) if coming from an EVM chain.
    bytes data; // payload sent in original message.
    EVMTokenAmount[] destTokenAmounts; // Tokens and their amounts in their destination chain representation.
  }

  // If extraArgs is empty bytes, the default is 200k gas limit.
  struct EVM2AnyMessage {
    bytes receiver; // abi.encode(receiver address) for dest EVM chains.
    bytes data; // Data payload.
    EVMTokenAmount[] tokenAmounts; // Token transfers.
    address feeToken; // Address of feeToken. address(0) means you will send msg.value.
    bytes extraArgs; // Populate this with _argsToBytes(EVMExtraArgsV3).
  }

  /// @notice Tag to indicate no execution on the destination chain. Execution will need to be done manually.
  /// @dev Preimage for this tag is: keccak256("NO_EXECUTION_TAG")[:4]
  bytes4 public constant NO_EXECUTION_TAG = 0xeba517d2;
  address public constant NO_EXECUTION_ADDRESS = address(bytes20(NO_EXECUTION_TAG));

  bytes4 public constant GENERIC_EXTRA_ARGS_V3_TAG = 0x302326cb;

  /// @notice The GenericExtraArgsV3 struct is used to pass extra arguments for all destination chain families. There
  /// are bytes fields inside that could contain encoded structs specific to each chain family. The primary field that
  /// contains chain family specific data is the executorArgs field. This field gets passed to the destination chain
  /// unmodified and is expected to be encoded in a way that the destination chain can understand.
  struct GenericExtraArgsV3 {
    /// @notice An array of CCV structs representing the cross-chain verifiers to be used for the message, and optional
    /// arguments that are passed into the CCV without modification or inspection. CCIP itself does not interpret these
    /// arguments: they are in the format of the CCV contracts being used.
    CCV[] ccvs;
    /// @notice The finality config, 0 means the default finality that the CCV considers final. Any non-zero value means
    /// a block depth. CCVs, Pools and the executor may all reject this value by reverting the transaction on the source
    /// chain if they do now want to take on the risk of the block depth specified.
    uint16 finalityConfig;
    /// @notice Gas limit for the callback on the destination chain. If the gas limit is zero, no callback will be
    /// performed, even if a receiver is specified. A gas limit of zero is useful when only token transfers are desired,
    /// or when the receiver is an EOA account instead of a contract. Besides this gasLimit check, there are other
    /// checks on the destination chain that may prevent the callback from being executed, depending on the destination
    /// chain family.
    /// @dev The sender is billed for the gas specified, not the gas actually used. Any unspent gas is not refunded.
    /// There are various ways to estimate the gas required for a callback on the destination chain, depending on the
    /// chain family. Please refer to the documentation for each chain family for more details.
    uint32 gasLimit;
    /// @notice Address of the executor contract on the source chain. The executor is responsible for executing the
    /// message on the destination chains once a quorum of CCVs have verified the message.
    address executor;
    /// @notice Destination chain family specific arguments for the executor. This field is passed to the destination
    /// chain as part of the message itself and these args are therefore fully protected through the message ID. The
    /// format of this field is specific to each chain family and is not interpreted by CCIP itself, only by the
    /// executor. Things that may be included here are Solana accounts or Sui object IDs, which must be secured through
    /// the message ID as passing in incorrect values can lead to loss of funds.
    bytes executorArgs;
    /// @notice Address of the token receiver on the destination chain, in bytes format. If an empty bytes array is
    /// provided, the receiver address from the message itself is used for token transfers. This field allows for
    /// scenarios where the token receiver is different from the message receiver.
    bytes tokenReceiver;
    /// @notice Additional arguments for token transfers. This field is passed into the token pool on the source chain
    /// and is not inspected by CCIP itself. The format of this field is therefore specific to the token pool being used
    /// and may vary between different pools.
    bytes tokenArgs;
  }

  function _argsToBytes(
    GenericExtraArgsV3 memory extraArgs
  ) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(GENERIC_EXTRA_ARGS_V3_TAG, extraArgs);
  }

  struct SVMExecutorArgsV1 {
    // TODO Use ATA or raw account flag
    bool useATA;
    uint64 accountIsWritableBitmap;
    // Additional accounts needed for execution of CCIP receiver. Must be empty if message.receiver is zero.
    // Token transfer related accounts are specified in the token pool lookup table on SVM.
    bytes32[] accounts;
  }

  struct SuiExecutorArgsV1 {
    bytes32[] receiverObjectIds;
  }

  // ================================================================
  // │                           Legacy                             │
  // ================================================================

  // Tag to indicate only a gas limit. Only usable for EVM as destination chain.
  bytes4 public constant EVM_EXTRA_ARGS_V1_TAG = 0x97a657c9;

  struct EVMExtraArgsV1 {
    uint256 gasLimit;
  }

  function _argsToBytes(
    EVMExtraArgsV1 memory extraArgs
  ) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(EVM_EXTRA_ARGS_V1_TAG, extraArgs);
  }

  // Tag to indicate a gas limit (or dest chain equivalent processing units) and Out Of Order Execution. This tag is
  // available for multiple chain families. If there is no chain family specific tag, this is the default available
  // for a chain.
  // Note: not available for Solana or Sui VM based chains.
  bytes4 public constant GENERIC_EXTRA_ARGS_V2_TAG = 0x181dcf10;

  /// @param gasLimit: gas limit for the callback on the destination chain.
  /// @param allowOutOfOrderExecution: if true, it indicates that the message can be executed in any order relative to
  /// other messages from the same sender. This value's default varies by chain. On some chains, a particular value is
  /// enforced, meaning if the expected value is not set, the message request will revert.
  /// @dev Fully compatible with the previously existing EVMExtraArgsV2.
  struct GenericExtraArgsV2 {
    uint256 gasLimit;
    bool allowOutOfOrderExecution;
  }

  // Extra args tag for chains that use the Sui VM.
  bytes4 public constant SUI_EXTRA_ARGS_V1_TAG = 0x21ea4ca9;

  // Extra args tag for chains that use the Solana VM.
  bytes4 public constant SVM_EXTRA_ARGS_V1_TAG = 0x1f3b3aba;

  struct SVMExtraArgsV1 {
    uint32 computeUnits;
    uint64 accountIsWritableBitmap;
    bool allowOutOfOrderExecution;
    bytes32 tokenReceiver;
    // Additional accounts needed for execution of CCIP receiver. Must be empty if message.receiver is zero.
    // Token transfer related accounts are specified in the token pool lookup table on SVM.
    bytes32[] accounts;
  }

  /// @dev The maximum number of accounts that can be passed in SVMExtraArgs.
  uint256 public constant SVM_EXTRA_ARGS_MAX_ACCOUNTS = 64;

  /// @dev The expected static payload size of a token transfer when Borsh encoded and submitted to SVM.
  /// TokenPool extra data and offchain data sizes are dynamic, and should be accounted for separately.
  uint256 public constant SVM_TOKEN_TRANSFER_DATA_OVERHEAD = (4 + 32) // source_pool
    + 32 // token_address
    + 4 // gas_amount
    + 4 // extra_data overhead
    + 32 // amount
    + 32 // size of the token lookup table account
    + 32 // token-related accounts in the lookup table, over-estimated to 32, typically between 11 - 13
    + 32 // token account belonging to the token receiver, e.g ATA, not included in the token lookup table
    + 32 // per-chain token pool config, not included in the token lookup table
    + 32 // per-chain token billing config, not always included in the token lookup table
    + 32; // OffRamp pool signer PDA, not included in the token lookup table

  /// @dev Number of overhead accounts needed for message execution on SVM.
  /// @dev These are message.receiver, and the OffRamp Signer PDA specific to the receiver.
  uint256 public constant SVM_MESSAGING_ACCOUNTS_OVERHEAD = 2;

  /// @dev The size of each SVM account address in bytes.
  uint256 public constant SVM_ACCOUNT_BYTE_SIZE = 32;

  struct SuiExtraArgsV1 {
    uint256 gasLimit;
    bool allowOutOfOrderExecution;
    bytes32 tokenReceiver;
    bytes32[] receiverObjectIds;
  }

  /// @dev The expected static payload size of a token transfer when BCS encoded and submitted to SUI.
  /// TokenPool extra data and offchain data sizes are dynamic, and should be accounted for separately.
  uint256 public constant SUI_TOKEN_TRANSFER_DATA_OVERHEAD = (4 + 32) // source_pool, 4 bytes for length, 32 bytes for address
    + 32 // dest_token_address
    + 4 // dest_gas_amount
    + 4 // extra_data length, the contents are calculated separately
    + 32; // amount

  /// @dev Number of overhead accounts needed for message execution on SUI.
  /// @dev This is the message.receiver.
  uint256 public constant SUI_MESSAGING_ACCOUNTS_OVERHEAD = 1;

  /// @dev The maximum number of receiver object ids that can be passed in SuiExtraArgs.
  uint256 public constant SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS = 64;

  /// @dev The size of each SUI account address in bytes.
  uint256 public constant SUI_ACCOUNT_BYTE_SIZE = 32;

  function _argsToBytes(
    GenericExtraArgsV2 memory extraArgs
  ) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(GENERIC_EXTRA_ARGS_V2_TAG, extraArgs);
  }

  function _svmArgsToBytes(
    SVMExtraArgsV1 memory extraArgs
  ) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(SVM_EXTRA_ARGS_V1_TAG, extraArgs);
  }

  function _suiArgsToBytes(
    SuiExtraArgsV1 memory extraArgs
  ) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(SUI_EXTRA_ARGS_V1_TAG, extraArgs);
  }

  /// @notice The CCV struct is used to represent a cross-chain verifier.
  struct CCV {
    /// @param The ccvAddress is the address of the verifier contract on the source chain
    address ccvAddress;
    /// @param args The args are the arguments that the verifier contract expects. They are opaque to CCIP and are only
    /// used in the CCV.
    bytes args;
  }
}
