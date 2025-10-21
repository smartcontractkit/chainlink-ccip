// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Library for CCIP internal definitions common to multiple contracts.
/// @dev The following is a non-exhaustive list of "known issues" for CCIP:
/// - We could implement yield claiming for Blast. This is not worth the custom code path on non-blast chains.
/// - uint32 is used for timestamps, which will overflow in 2106. This is not a concern for the current use case, as we
/// expect to have migrated to a new version by then.
library Internal {
  // bytes4(keccak256("CCIP ChainFamilySelector EVM"));
  bytes4 public constant CHAIN_FAMILY_SELECTOR_EVM = 0x2812d52c;

  // bytes4(keccak256("CCIP ChainFamilySelector SVM"));
  bytes4 public constant CHAIN_FAMILY_SELECTOR_SVM = 0x1e10bdc4;

  // bytes4(keccak256("CCIP ChainFamilySelector APTOS"));
  bytes4 public constant CHAIN_FAMILY_SELECTOR_APTOS = 0xac77ffec;

  // bytes4(keccak256("CCIP ChainFamilySelector SUI"));
  bytes4 public constant CHAIN_FAMILY_SELECTOR_SUI = 0xc4e05953;

  // byte4(keccak256("CCIP ChainFamilySelector TVM"));
  bytes4 public constant CHAIN_FAMILY_SELECTOR_TVM = 0x647e2ba9;

  error InvalidEVMAddress(bytes encodedAddress);
  error Invalid32ByteAddress(bytes encodedAddress);
  error InvalidTVMAddress(bytes encodedAddress);

  /// @dev We limit return data to a selector plus 4 words. This is to avoid malicious contracts from returning
  /// large amounts of data and causing repeated out-of-gas scenarios.
  uint16 internal constant MAX_RET_BYTES = 4 + 4 * 32;

  /// @dev The address used to send calls for gas estimation.
  /// You only need to use this address if the minimum gas limit specified by the user is not actually enough to execute the
  /// given message and you're attempting to estimate the actual necessary gas limit
  address public constant GAS_ESTIMATION_SENDER = address(0xC11C11C11C11C11C11C11C11C11C11C11C11C1);

  /// @notice A collection of token price and gas price updates.
  /// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
  struct PriceUpdates {
    TokenPriceUpdate[] tokenPriceUpdates;
    GasPriceUpdate[] gasPriceUpdates;
  }

  /// @notice Token price in USD.
  /// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
  struct TokenPriceUpdate {
    address sourceToken; // Source token.
    uint224 usdPerToken; // 1e18 USD per 1e18 of the smallest token denomination.
  }

  /// @notice Gas price for a given chain in USD, its value may contain tightly packed fields.
  /// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
  struct GasPriceUpdate {
    uint64 destChainSelector; // Destination chain selector.
    uint224 usdPerUnitGas; // 1e18 USD per smallest unit (e.g. wei) of destination chain gas.
  }

  /// @notice A timestamped uint224 value that can contain several tightly packed fields.
  struct TimestampedPackedUint224 {
    uint224 value; // ────╮ Value in uint224, packed.
    uint32 timestamp; // ─╯ Timestamp of the most recent price update.
  }

  /// @dev Gas price is stored in 112-bit unsigned int. uint224 can pack 2 prices.
  /// When packing L1 and L2 gas prices, L1 gas price is left-shifted to the higher-order bits.
  /// Using uint8 type, which cannot be higher than other bit shift operands, to avoid shift operand type warning.
  uint8 public constant GAS_PRICE_BITS = 112;

  struct SourceTokenData {
    // The source pool address, abi encoded. This value is trusted as it was obtained through the onRamp. It can be
    // relied upon by the destination pool to validate the source pool.
    bytes sourcePoolAddress;
    // The address of the destination token, abi encoded in the case of EVM chains.
    // This value is UNTRUSTED as any pool owner can return whatever value they want.
    bytes destTokenAddress;
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    bytes extraData;
    uint32 destGasAmount; // The amount of gas available for the releaseOrMint and balanceOf calls on the offRamp
  }

  /// @dev We disallow the first 1024 addresses to avoid calling into a range known for hosting precompiles. Calling
  /// into precompiles probably won't cause any issues, but to be safe we can disallow this range. It is extremely
  /// unlikely that anyone would ever be able to generate an address in this range. There is no official range of
  /// precompiles, but EIP-7587 proposes to reserve the range 0x100 to 0x1ff. Our range is more conservative, even
  /// though it might not be exhaustive for all chains, which is OK. We also disallow the zero address, which is a
  /// common practice.
  uint256 public constant EVM_PRECOMPILE_SPACE = 1024;

  // According to the Aptos docs, the first 0xa addresses are reserved for precompiles.
  // https://github.com/aptos-labs/aptos-core/blob/main/aptos-move/framework/aptos-framework/doc/account.md#function-create_framework_reserved_account-1
  uint256 public constant APTOS_PRECOMPILE_SPACE = 0x0b;

  // According to the Sui docs, a set of non-contiguous addresses below 0xdee9 are reserved for system use.
  // https://github.com/MystenLabs/sui/blob/54ae98508569804127bd73d84aa2fb816bdea402/crates/sui-types/src/lib.rs#L141
  uint256 public constant SUI_PRECOMPILE_SPACE = 0xdee9;

  /// @notice This methods provides validation for parsing abi encoded addresses by ensuring the address is within the
  /// EVM address space. If it isn't it will revert with an InvalidEVMAddress error, which we can catch and handle
  /// more gracefully than a revert from abi.decode.
  function _validateEVMAddress(
    bytes memory encodedAddress
  ) internal pure {
    if (encodedAddress.length != 32) revert InvalidEVMAddress(encodedAddress);
    uint256 encodedAddressUint = abi.decode(encodedAddress, (uint256));
    if (encodedAddressUint > type(uint160).max || encodedAddressUint < EVM_PRECOMPILE_SPACE) {
      revert InvalidEVMAddress(encodedAddress);
    }
  }

  /// @notice This methods provides validation for parsing abi encoded addresses by ensuring the address is within the
  /// bounds of [minValue, uint256.max]. If it isn't it will revert with an Invalid32ByteAddress error.
  function _validate32ByteAddress(bytes memory encodedAddress, uint256 minValue) internal pure {
    if (encodedAddress.length != 32) revert Invalid32ByteAddress(encodedAddress);
    if (minValue > 0) {
      if (abi.decode(encodedAddress, (uint256)) < minValue) {
        revert Invalid32ByteAddress(encodedAddress);
      }
    }
  }

  /// @notice This methods provides validation for TON User-friendly addresses by ensuring the address is 36 bytes long.
  /// @dev The encodedAddress is expected to be the 36-byte raw representation:
  /// - 1 byte: flags (isBounceable, isTestnetOnly, etc.)
  /// - 1 byte: workchain_id (0x00 for BaseChain, 0xff for MasterChain)
  /// - 32 bytes: account_id
  /// - 2 bytes: CRC16 checksum(computationally heavy, validation omitted for simplicity)
  /// @param encodedAddress The 36-byte TON address.
  function _validateTVMAddress(
    bytes memory encodedAddress
  ) internal pure {
    if (encodedAddress.length != 36) revert InvalidTVMAddress(encodedAddress);
    bytes32 accountId;
    assembly {
      accountId := mload(add(encodedAddress, 0x22)) // 0x22 = 0x20 (data start) + 2 (offset for account_id)
    }
    if (accountId == bytes32(0)) revert InvalidTVMAddress(encodedAddress);
  }

  /// @notice Enum listing the possible message execution states within the offRamp contract.
  /// UNTOUCHED never executed.
  /// IN_PROGRESS currently being executed, used a replay protection.
  /// SUCCESS successfully executed. End state.
  /// FAILURE unsuccessfully executed, manual execution is now enabled.
  /// @dev RMN depends on this enum, if changing, please notify the RMN maintainers.
  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }

  struct EVM2AnyTokenTransfer {
    // The source pool EVM address. This value is trusted as it was obtained through the onRamp. It can be relied
    // upon by the destination pool to validate the source pool.
    address sourcePoolAddress;
    // The EVM address of the destination token.
    // This value is UNTRUSTED as any pool owner can return whatever value they want.
    bytes destTokenAddress;
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    bytes extraData;
    uint256 amount; // Amount of tokens.
    // Destination chain data used to execute the token transfer on the destination chain. For an EVM destination, it
    // consists of the amount of gas available for the releaseOrMint and transfer calls made by the offRamp.
    bytes destExecData;
  }

  /// @dev Holds a merkle root and interval for a source chain so that an array of these can be passed in the CommitReport.
  /// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
  /// @dev inefficient struct packing intentionally chosen to maintain order of specificity. Not a storage struct so impact is minimal.
  // solhint-disable-next-line gas-struct-packing
  struct MerkleRoot {
    uint64 sourceChainSelector; // Remote source chain selector that the Merkle Root is scoped to
    bytes onRampAddress; //        Generic onRamp address, to support arbitrary sources; for EVM, use abi.encode
    uint64 minSeqNr; // ─────────╮ Minimum sequence number, inclusive
    uint64 maxSeqNr; // ─────────╯ Maximum sequence number, inclusive
    bytes32 merkleRoot; //         Merkle root covering the interval & source chain messages
  }
}
