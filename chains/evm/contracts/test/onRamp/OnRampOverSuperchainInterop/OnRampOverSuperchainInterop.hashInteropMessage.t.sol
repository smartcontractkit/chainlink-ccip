// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_hashInteropMessage is OnRampOverSuperchainInteropSetup {
  function _getMetadataHash() internal view returns (bytes32) {
    return keccak256(
      abi.encode(
        Internal.ANY_2_EVM_MESSAGE_HASH,
        SOURCE_CHAIN_SELECTOR,
        DEST_CHAIN_SELECTOR,
        keccak256(abi.encode(address(s_onRampOverSuperchainInterop)))
      )
    );
  }

  function test_BasicMessageHash() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash = s_onRampOverSuperchainInterop.hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }

  function test_TokenMessageHash() public {
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessageWithTokens();

    bytes32 messageHash = s_onRampOverSuperchainInterop.hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }

  function test_SameMessageProducesSameHash() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash1 = s_onRampOverSuperchainInterop.hashInteropMessage(message);
    bytes32 messageHash2 = s_onRampOverSuperchainInterop.hashInteropMessage(message);

    assertEq(messageHash1, messageHash2);
  }

  function test_DifferentMessageProduceDifferentHashes() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    bytes32 messageHash1 = s_onRampOverSuperchainInterop.hashInteropMessage(message);

    message.header.destChainSelector = DEST_CHAIN_SELECTOR + 1;
    bytes32 messageHash2 = s_onRampOverSuperchainInterop.hashInteropMessage(message);

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

    bytes32 messageHash = s_onRampOverSuperchainInterop.hashInteropMessage(message);
    bytes32 expectedMessageHash = Internal._hash(message, _getMetadataHash());

    assertEq(messageHash, expectedMessageHash);
  }
}
