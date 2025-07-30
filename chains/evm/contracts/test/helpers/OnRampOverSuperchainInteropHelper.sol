// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {OnRampOverSuperchainInterop} from "../../onRamp/OnRampOverSuperchainInterop.sol";

contract OnRampOverSuperchainInteropHelper is OnRampOverSuperchainInterop {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    DestChainConfigArgs[] memory destChainConfigs
  ) OnRampOverSuperchainInterop(staticConfig, dynamicConfig, destChainConfigs) {}

  /// @notice Exposes the internal _postProcessMessage function as public for testing
  function postProcessMessage(
    Internal.EVM2AnyRampMessage memory message
  ) public returns (Internal.EVM2AnyRampMessage memory) {
    return super._postProcessMessage(message);
  }

  /// @notice Helper function to access the sent interop message hash storage for testing
  function getSentInteropMessageHash(uint64 destChainSelector, uint64 sequenceNumber) public view returns (bytes32) {
    return s_sentInteropMessageHashes[destChainSelector][sequenceNumber];
  }
}
