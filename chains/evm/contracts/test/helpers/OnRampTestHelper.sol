// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";

/// @notice Test wrapper for OnRamp to expose internal functions for testing.
contract OnRampTestHelper is OnRamp {
  constructor(StaticConfig memory staticConfig, DynamicConfig memory dynamicConfig) OnRamp(staticConfig, dynamicConfig) {}

  /// @notice Exposes the internal _parseExtraArgsWithDefaults function for testing.
  function parseExtraArgsWithDefaults(
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs
  ) external pure returns (Client.EVMExtraArgsV3 memory) {
    return _parseExtraArgsWithDefaults(destChainConfig, extraArgs);
  }

  /// @notice Exposes the internal _mergeCCVLists function for testing.
  function mergeCCVLists(
    Client.CCV[] memory userRequestedOrDefaultCCVs,
    address[] memory laneMandatedCCVs,
    address[] memory poolRequiredCCVs,
    address[] memory defaultCCVs
  ) external pure returns (Client.CCV[] memory ccvs) {
    return _mergeCCVLists(userRequestedOrDefaultCCVs, laneMandatedCCVs, poolRequiredCCVs, defaultCCVs);
  }

  function getCCVsForPool(
    uint64 destChainSelector,
    address token,
    uint256 amount,
    uint16 finality,
    bytes memory tokenArgs
  ) external view returns (address[] memory) {
    return _getCCVsForPool(destChainSelector, token, amount, finality, tokenArgs);
  }
}
