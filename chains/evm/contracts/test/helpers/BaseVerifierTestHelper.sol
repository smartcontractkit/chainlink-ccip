// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

/// @notice Test helper contract to expose BaseVerifier's internal functions for testing
contract BaseVerifierTestHelper is BaseVerifier {
  constructor(
    string memory storageLocation
  ) BaseVerifier(storageLocation) {}

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

  function assertSenderIsAllowed(uint64 destChainSelector, address sender, address caller) external view {
    _assertSenderIsAllowed(destChainSelector, sender, caller);
  }

  function forwardToVerifier(
    address,
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external pure returns (bytes memory) {
    return "";
  }

  function getFee(address, uint64, Client.EVM2AnyMessage memory, bytes memory) external pure returns (uint256) {
    return 0;
  }

  function verifyMessage(
    address originalCaller,
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageId,
    bytes memory ccvData
  ) external {}

  function typeAndVersion() external pure override returns (string memory) {
    return "BaseVerifierTestHelper 1.0.0";
  }
}
