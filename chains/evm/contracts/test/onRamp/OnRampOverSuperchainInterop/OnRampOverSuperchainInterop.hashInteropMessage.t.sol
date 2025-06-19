// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_hashInteropMessage is OnRampOverSuperchainInteropSetup {
  function test_BasicMessageHash() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash = SuperchainInterop._hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getOffRampMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }

  function test_TokenMessageHash() public {
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessageWithTokens();

    bytes32 messageHash = SuperchainInterop._hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getOffRampMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }

  function test_SameMessageProducesSameHash() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash1 = SuperchainInterop._hashInteropMessage(message);
    bytes32 messageHash2 = SuperchainInterop._hashInteropMessage(message);

    assertEq(messageHash1, messageHash2);
  }

  function test_DifferentMessageProduceDifferentHashes() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash1 = SuperchainInterop._hashInteropMessage(message);

    message.header.destChainSelector = DEST_CHAIN_SELECTOR + 1;
    bytes32 messageHash2 = SuperchainInterop._hashInteropMessage(message);

    assertTrue(messageHash1 != messageHash2);
  }

  function testFuzz_DifferentMessageData_Success(
    bytes memory sender,
    bytes memory data,
    address receiver,
    uint256 gasLimit,
    uint64 sequenceNumber,
    uint64 nonce
  ) public view {
    Internal.Any2EVMRampMessage memory message =
      _generateAny2EVMMessageWithCustomFields(sender, data, receiver, gasLimit, sequenceNumber, nonce);

    bytes32 messageHash = SuperchainInterop._hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getOffRampMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }
}
