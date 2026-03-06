// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../interfaces/IAdvancedPoolHooks.sol";
import {IExtractor} from "@chainlink/policy-management/interfaces/IExtractor.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../libraries/Pool.sol";

/// @notice Extracts named parameters from AdvancedPoolHooks preflightCheck and postflightCheck calldata
/// for ACE policy evaluation.
contract AdvancedPoolHooksExtractor is IExtractor {
  string public constant override typeAndVersion = "AdvancedPoolHooksExtractor 2.0.0-dev";

  /// @notice Parameter key for the sender address initiating the transfer.
  bytes32 public constant PARAM_FROM = keccak256("from");
  /// @notice Parameter key for the recipient. Raw bytes in preflight, abi.encoded in postflight.
  bytes32 public constant PARAM_TO = keccak256("to");
  /// @notice Parameter key for the transfer amount specified during ccipSend.
  bytes32 public constant PARAM_AMOUNT = keccak256("amount");
  /// @notice Parameter key for the transfer amount after fee deduction. Only present in preflight.
  bytes32 public constant PARAM_AMOUNT_POST_FEE = keccak256("amount_post_fee");
  /// @notice Parameter key for the remote chain selector.
  bytes32 public constant PARAM_REMOTE_CHAIN_SELECTOR = keccak256("remote_chain_selector");
  /// @notice Parameter key for the local token address.
  bytes32 public constant PARAM_TOKEN = keccak256("token");
  /// @notice Parameter key for the requested number of block confirmations.
  bytes32 public constant PARAM_BLOCK_CONFIRMATIONS_REQUESTED = keccak256("block_confirmations_requested");
  /// @notice Parameter key for the source pool address. Only present in postflight.
  bytes32 public constant PARAM_SOURCE_POOL_ADDRESS = keccak256("source_pool_address");
  /// @notice Parameter key for the source pool data. Only present in postflight.
  bytes32 public constant PARAM_SOURCE_POOL_DATA = keccak256("source_pool_data");
  /// @notice Parameter key for the source-denominated transfer amount. Only present in postflight.
  bytes32 public constant PARAM_SOURCE_DENOMINATED_AMOUNT = keccak256("source_denominated_amount");

  /// @inheritdoc IExtractor
  function extract(
    IPolicyEngine.Payload calldata payload
  ) public pure virtual returns (IPolicyEngine.Parameter[] memory) {
    if (payload.selector == IAdvancedPoolHooks.preflightCheck.selector) {
      return _extractPreflightCheck(payload);
    }

    if (payload.selector == IAdvancedPoolHooks.postflightCheck.selector) {
      return _extractPostflightCheck(payload);
    }

    revert IPolicyEngine.UnsupportedSelector(payload.selector);
  }

  /// @dev Decodes preflightCheck arguments: (LockOrBurnInV1, uint16, bytes, uint256).
  function _extractPreflightCheck(
    IPolicyEngine.Payload calldata payload
  ) internal pure returns (IPolicyEngine.Parameter[] memory) {
    // tokenArgs is skipped as it is treated as context in the payload.
    (Pool.LockOrBurnInV1 memory lockOrBurnIn, uint16 blockConfirmationsRequested,, uint256 amountPostFee) =
      abi.decode(payload.data, (Pool.LockOrBurnInV1, uint16, bytes, uint256));

    IPolicyEngine.Parameter[] memory result = new IPolicyEngine.Parameter[](7);
    result[0] = IPolicyEngine.Parameter(PARAM_FROM, abi.encode(lockOrBurnIn.originalSender));
    result[1] = IPolicyEngine.Parameter(PARAM_TO, lockOrBurnIn.receiver);
    result[2] = IPolicyEngine.Parameter(PARAM_AMOUNT, abi.encode(lockOrBurnIn.amount));
    result[3] = IPolicyEngine.Parameter(PARAM_AMOUNT_POST_FEE, abi.encode(amountPostFee));
    result[4] = IPolicyEngine.Parameter(PARAM_REMOTE_CHAIN_SELECTOR, abi.encode(lockOrBurnIn.remoteChainSelector));
    result[5] = IPolicyEngine.Parameter(PARAM_TOKEN, abi.encode(lockOrBurnIn.localToken));
    result[6] = IPolicyEngine.Parameter(PARAM_BLOCK_CONFIRMATIONS_REQUESTED, abi.encode(blockConfirmationsRequested));

    return result;
  }

  /// @dev Decodes postflightCheck arguments: (ReleaseOrMintInV1, uint256, uint16).
  function _extractPostflightCheck(
    IPolicyEngine.Payload calldata payload
  ) internal pure returns (IPolicyEngine.Parameter[] memory) {
    (Pool.ReleaseOrMintInV1 memory releaseOrMintIn, uint256 localAmount, uint16 blockConfirmationsRequested) =
      abi.decode(payload.data, (Pool.ReleaseOrMintInV1, uint256, uint16));

    // offchainTokenData is skipped as it is treated as context in the payload.
    // Note offchainTokenData is no longer used in v2+ TokenPools.
    IPolicyEngine.Parameter[] memory result = new IPolicyEngine.Parameter[](9);
    result[0] = IPolicyEngine.Parameter(PARAM_FROM, releaseOrMintIn.originalSender);
    result[1] = IPolicyEngine.Parameter(PARAM_TO, abi.encode(releaseOrMintIn.receiver));
    result[2] = IPolicyEngine.Parameter(PARAM_AMOUNT, abi.encode(localAmount));
    result[3] = IPolicyEngine.Parameter(PARAM_REMOTE_CHAIN_SELECTOR, abi.encode(releaseOrMintIn.remoteChainSelector));
    result[4] = IPolicyEngine.Parameter(PARAM_TOKEN, abi.encode(releaseOrMintIn.localToken));
    result[5] = IPolicyEngine.Parameter(PARAM_BLOCK_CONFIRMATIONS_REQUESTED, abi.encode(blockConfirmationsRequested));
    result[6] = IPolicyEngine.Parameter(PARAM_SOURCE_POOL_ADDRESS, releaseOrMintIn.sourcePoolAddress);
    result[7] = IPolicyEngine.Parameter(PARAM_SOURCE_POOL_DATA, releaseOrMintIn.sourcePoolData);
    result[8] =
      IPolicyEngine.Parameter(PARAM_SOURCE_DENOMINATED_AMOUNT, abi.encode(releaseOrMintIn.sourceDenominatedAmount));

    return result;
  }
}
