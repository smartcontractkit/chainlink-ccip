// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

contract OffRampOverSuperchainInterop_constructProofs is OffRampOverSuperchainInteropSetup {
  function test_constructProofs() public view {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    address onRampAddress = ON_RAMP_ADDRESS;
    uint256 blockNumber = 12345;
    uint256 logIndex = 1;
    uint256 timestamp = block.timestamp;
    uint256 chainId = CHAIN_ID_1;

    bytes32[] memory proofs = new bytes32[](5);
    proofs[0] = bytes32(uint256(uint160(onRampAddress)));
    proofs[1] = bytes32(blockNumber);
    proofs[2] = bytes32(logIndex);
    proofs[3] = bytes32(timestamp);
    proofs[4] = bytes32(chainId);

    (Identifier memory identifier, bytes32 logHash) = s_offRampOverSuperchainInterop.constructProofs(message, proofs);

    // Verify identifier
    assertEq(identifier.origin, onRampAddress);
    assertEq(identifier.blockNumber, blockNumber);
    assertEq(identifier.logIndex, logIndex);
    assertEq(identifier.timestamp, timestamp);
    assertEq(identifier.chainId, chainId);

    // Verify log hash
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, 1, message);
    assertEq(logHash, expectedLogHash);
  }

  function testFuzz_constructProofs_Success(
    address onRamp,
    uint256 blockNumber,
    uint256 logIndex,
    uint256 timestamp,
    uint256 chainId,
    uint64 chainSelector,
    uint64 sequenceNumber
  ) public view {
    // Constrain inputs to avoid InvalidEncodingOfIdentifierInProofs
    vm.assume(onRamp != address(0));
    vm.assume(blockNumber != 0);
    vm.assume(timestamp != 0);
    vm.assume(chainId != 0);

    Internal.Any2EVMRampMessage memory message = _generateValidMessage(chainSelector, sequenceNumber);

    bytes32[] memory proofs = new bytes32[](5);
    proofs[0] = bytes32(uint256(uint160(onRamp)));
    proofs[1] = bytes32(blockNumber);
    proofs[2] = bytes32(logIndex);
    proofs[3] = bytes32(timestamp);
    proofs[4] = bytes32(chainId);

    (Identifier memory identifier, bytes32 logHash) = s_offRampOverSuperchainInterop.constructProofs(message, proofs);

    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, sequenceNumber, message);

    assertEq(identifier.origin, onRamp);
    assertEq(identifier.blockNumber, blockNumber);
    assertEq(identifier.logIndex, logIndex);
    assertEq(identifier.timestamp, timestamp);
    assertEq(identifier.chainId, chainId);
    assertEq(expectedLogHash, logHash);
  }

  function test_constructProofs_RevertWhen_InvalidProofsLength() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    // Test with too few proofs
    bytes32[] memory shortProofs = new bytes32[](4);
    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidProofsWordLength.selector, 4, 5));
    s_offRampOverSuperchainInterop.constructProofs(message, shortProofs);

    // Test with too many proofs
    bytes32[] memory longProofs = new bytes32[](6);
    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidProofsWordLength.selector, 6, 5));
    s_offRampOverSuperchainInterop.constructProofs(message, longProofs);

    // Test with empty proofs
    bytes32[] memory emptyProofs = new bytes32[](0);
    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidProofsWordLength.selector, 0, 5));
    s_offRampOverSuperchainInterop.constructProofs(message, emptyProofs);
  }

  function test_constructProofs_RevertWhen_InvalidEncodingOfIdentifierInProofs() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);
    bytes32[] memory proofs = new bytes32[](5);
    proofs[0] = bytes32(uint256(uint160(ON_RAMP_ADDRESS)));
    proofs[1] = bytes32(block.number);
    proofs[2] = bytes32(uint256(1)); // logIndex (can be zero)
    proofs[3] = bytes32(block.timestamp);
    proofs[4] = bytes32(CHAIN_ID_1);

    // Test proofs[0] (origin) cannot be zero
    proofs[0] = bytes32(0);
    vm.expectRevert(
      abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidEncodingOfIdentifierInProofs.selector, proofs)
    );
    s_offRampOverSuperchainInterop.constructProofs(message, proofs);

    // Reset and test proofs[1] (blockNumber) cannot be zero
    proofs[0] = bytes32(uint256(uint160(ON_RAMP_ADDRESS)));
    proofs[1] = bytes32(0);
    vm.expectRevert(
      abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidEncodingOfIdentifierInProofs.selector, proofs)
    );
    s_offRampOverSuperchainInterop.constructProofs(message, proofs);

    // Reset and test proofs[3] (timestamp) cannot be zero
    proofs[1] = bytes32(block.number);
    proofs[3] = bytes32(0);
    vm.expectRevert(
      abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidEncodingOfIdentifierInProofs.selector, proofs)
    );
    s_offRampOverSuperchainInterop.constructProofs(message, proofs);

    // Reset and test proofs[4] (chainId) cannot be zero
    proofs[3] = bytes32(block.timestamp);
    proofs[4] = bytes32(0);
    vm.expectRevert(
      abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidEncodingOfIdentifierInProofs.selector, proofs)
    );
    s_offRampOverSuperchainInterop.constructProofs(message, proofs);
  }
}
