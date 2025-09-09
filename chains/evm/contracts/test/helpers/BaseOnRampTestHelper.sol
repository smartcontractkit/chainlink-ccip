// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {BaseOnRamp} from "../../onRamp/BaseOnRamp.sol";

/// @notice Test helper contract to expose BaseOnRamp's internal functions for testing
contract BaseOnRampTestHelper is BaseOnRamp {
  constructor(
    address rmnRemote
  ) BaseOnRamp(rmnRemote) {}

  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] calldata destChainConfigArgs
  ) external {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external {
    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  function withdrawFeeTokens(address[] calldata feeTokens, address feeAggregator) external {
    _withdrawFeeTokens(feeTokens, feeAggregator);
  }

  function assertSenderIsAllowed(uint64 destChainSelector, address sender) external view {
    _assertSenderIsAllowed(destChainSelector, sender);
  }

  function assertNotCursed(
    uint64 destChainSelector
  ) external view {
    _assertNotCursed(destChainSelector);
  }

  function forwardToVerifier(uint64, bytes32, address, bytes memory, uint256) external pure returns (bytes memory) {
    return "";
  }

  function getFee(uint64, bytes32, address, Client.EVM2AnyMessage memory, bytes memory) external pure returns (uint256) {
    return 0;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "BaseOnRampTestHelper 1.0.0";
  }
}
