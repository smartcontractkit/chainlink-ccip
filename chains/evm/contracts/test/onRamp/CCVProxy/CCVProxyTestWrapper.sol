// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";

/// @notice Test wrapper for CCVProxy to expose internal functions for testing
contract CCVProxyTestWrapper is CCVProxy {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig
  ) CCVProxy(staticConfig, dynamicConfig) {}

  /// @notice Exposes the internal _parseExtraArgsWithDefaults function for testing
  function parseExtraArgsWithDefaults(
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs
  ) external pure returns (Client.EVMExtraArgsV3 memory) {
    return _parseExtraArgsWithDefaults(destChainConfig, extraArgs);
  }
}
