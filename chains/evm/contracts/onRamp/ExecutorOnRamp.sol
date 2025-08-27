// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "./BaseOnRamp.sol";
import {Client} from "../libraries/Client.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice The ExecutorOnRamp configures executor fees, destination support, and verifier limits.
contract ExecutorOnRamp is Ownable2StepMsgSender, BaseOnRamp {
  error InvalidConfig();
  error InvalidExtraArgsVersion(bytes4 provided);
  error ExceedsMaxPossibleVerifiers(uint256 provided, uint256 max);
  error ExceedsMaxRequiredVerifiers(uint256 provided, uint256 max);

  event ConfigSet(DynamicConfig dynamicConfig);

  /// @dev Struct that contains the dynamic configuration.
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    // Address that quotes a fee for each message
    address feeQuoter;
    // Address that receives fees, regardless of who calls withdraw
    address feeAggregator; 
    // Max(required ccvs + optional ccvs)
    // Limits the number of ccvs that the executor needs to search for results from
    uint8 maxPossibleVerifiersPerMsg;
    // Max(required ccvs + optional ccv threshold)
    // Limits the number of ccvs that the executor needs to submit on-chain
    uint8 maxRequiredVerifersPerMsg;
  }

  string public constant override typeAndVersion = "ExecutorOnRamp 1.7.0-dev";

  // DYNAMIC CONFIG
  /// @dev The dynamic config for the onRamp.
  DynamicConfig private s_dynamicConfig;

  constructor(DynamicConfig memory dynamicConfig) BaseOnRamp(address(0)) {
    _setDynamicConfig(dynamicConfig);
  }

  /// TODO: Should we instead create an IOnRamp interface that both CCVOnRamp and ExecutorOnRamp can use?
  /// @notice Stub required to implement ICCVOnRamp, not used in ExecutorOnRamp.
  function forwardToVerifier(bytes calldata, uint256) external pure returns (bytes memory) {
    return "";
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the dynamic onRamp config.
  /// @return dynamicConfig the dynamic configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The configuration.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (
      dynamicConfig.feeQuoter == address(0) ||
      dynamicConfig.feeAggregator == address(0) ||
      dynamicConfig.maxRequiredVerifersPerMsg > dynamicConfig.maxPossibleVerifiersPerMsg
    ) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(dynamicConfig);
  }

  /// @notice Updates destination chains specific configs.
  /// @param destChainConfigArgs Array of destination chain specific configs.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] calldata destChainConfigArgs
  ) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory,
    bytes calldata extraArgs
  ) external view returns (uint256) {
    DestChainConfig storage destChainConfig = _getDestChainConfig(destChainSelector);
    if (destChainConfig.ccvProxy == address(0)) {
      revert InvalidDestChainConfig(destChainSelector);
    }

    if (bytes4(extraArgs[0:4]) != Client.GENERIC_EXTRA_ARGS_V3_TAG) {
        revert InvalidExtraArgsVersion(bytes4(extraArgs[0:4]));
    }

    Client.EVMExtraArgsV3 memory resolvedArgs = abi.decode(extraArgs[4:], (Client.EVMExtraArgsV3));

    uint256 possibleVerifiers = resolvedArgs.requiredCCV.length + resolvedArgs.optionalCCV.length;
    if (possibleVerifiers > s_dynamicConfig.maxPossibleVerifiersPerMsg) {
      revert ExceedsMaxPossibleVerifiers(possibleVerifiers, s_dynamicConfig.maxPossibleVerifiersPerMsg);
    }

    uint256 requiredVerifiers = resolvedArgs.requiredCCV.length + resolvedArgs.optionalThreshold;
    if (requiredVerifiers > s_dynamicConfig.maxRequiredVerifersPerMsg) {
      revert ExceedsMaxRequiredVerifiers(requiredVerifiers, s_dynamicConfig.maxRequiredVerifersPerMsg);
    }

    // TODO: get execution fee using extraArgs, for now we just return 0
    return 0;
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @dev This function can be permissionless as it only transfers tokens to the fee aggregator which is a trusted address.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
    _withdrawFeeTokens(feeTokens, s_dynamicConfig.feeAggregator);
  }
}
