// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {CCVProxy} from "../../onRamp/CCVProxy.sol";

/// @notice Test wrapper for CCVProxy to expose internal functions for testing.
contract CCVProxyTestHelper is CCVProxy {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig
  ) CCVProxy(staticConfig, dynamicConfig) {}

  /// @notice Exposes the internal _parseExtraArgsWithDefaults function for testing.
  function parseExtraArgsWithDefaults(
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs
  ) external pure returns (Client.EVMExtraArgsV3 memory) {
    return _parseExtraArgsWithDefaults(destChainConfig, extraArgs);
  }

  /// @notice Exposes the internal _mergeCCVsWithPoolAndLaneMandated function for testing.
  /// Note: This assumes the test has already set up the DestChainConfig via applyDestChainConfigUpdates
  function mergeCCVsWithPoolAndLaneMandated(
    uint64 destChainSelector,
    address[] memory poolRequiredCCVs,
    Client.CCV[] memory requiredCCV,
    Client.CCV[] memory optionalCCV,
    uint8 optionalThreshold
  )
    external
    view
    returns (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold)
  {
    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
    return
      _mergeCCVsWithPoolAndLaneMandated(destChainConfig, poolRequiredCCVs, requiredCCV, optionalCCV, optionalThreshold);
  }
}
