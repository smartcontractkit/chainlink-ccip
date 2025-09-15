// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {BaseOnRamp} from "../../onRamp/BaseOnRamp.sol";

/// @notice Test helper contract to expose BaseOnRamp's internal functions for testing
contract BaseOnRampTestHelper is BaseOnRamp {
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

  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external pure returns (bytes memory) {
    return "";
  }

  function getFee(uint64, Client.EVM2AnyMessage memory, bytes memory) external pure returns (uint256) {
    return 0;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "BaseOnRampTestHelper 1.0.0";
  }
}
