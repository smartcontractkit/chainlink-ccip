// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_hashInteropMessage is OnRampOverSuperchainInteropSetup {
  function test_hashInteropMessage_BasicMessageHash() public {
    Internal.Any2EVMRampMessage memory basicMessage = _generateBasicAny2EVMMessage();
    Internal.Any2EVMRampMessage memory tokenMessage = _generateAny2EVMMessageWithTokens();

    assertEq(
      Internal._hash(basicMessage, _getOffRampMetadataHash()),
      SuperchainInterop._hashInteropMessage(basicMessage, address(s_onRampOverSuperchainInterop))
    );
    assertEq(
      Internal._hash(tokenMessage, _getOffRampMetadataHash()),
      SuperchainInterop._hashInteropMessage(tokenMessage, address(s_onRampOverSuperchainInterop))
    );
  }

  function test_hashInteropMessage_SameMessageProducesSameHash() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash1 = SuperchainInterop._hashInteropMessage(message, address(s_onRampOverSuperchainInterop));

    // MessageHash should remain the same when calculated in a different block
    vm.roll(block.number + 100);
    vm.warp(block.timestamp + 100);

    bytes32 messageHash2 = SuperchainInterop._hashInteropMessage(message, address(s_onRampOverSuperchainInterop));

    assertEq(messageHash1, messageHash2);
  }

  function testFuzz_hashInteropMessage_DifferentMessageFields(
    bytes memory sender,
    bytes memory data,
    address receiver,
    uint256 gasLimit,
    uint64 sequenceNumber,
    uint64 nonce
  ) public view {
    Internal.Any2EVMRampMessage memory message =
      _generateAny2EVMMessageWithCustomFields(sender, data, receiver, gasLimit, sequenceNumber, nonce);

    bytes32 messageHash = SuperchainInterop._hashInteropMessage(message, address(s_onRampOverSuperchainInterop));
    bytes32 expectedMessageHash = Internal._hash(message, _getOffRampMetadataHash());

    assertEq(expectedMessageHash, messageHash);
  }
}
