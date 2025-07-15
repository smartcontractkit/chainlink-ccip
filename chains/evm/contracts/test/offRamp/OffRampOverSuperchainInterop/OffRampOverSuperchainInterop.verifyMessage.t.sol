// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {MockCrossL2Inbox} from "../../mocks/MockCrossL2Inbox.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

contract OffRampOverSuperchainInterop_verifyReport is OffRampOverSuperchainInteropSetup {
  function test_verifyReport() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    uint256 blockNumber = 12345;
    uint256 logIndex = 1;
    uint256 timestamp = block.timestamp;

    (bytes32[] memory proofs, Identifier memory expectedIdentifier) =
      _getValidProofsAndIdentifier(blockNumber, logIndex, timestamp);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    // Setup CrossL2Inbox mock expectations
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, 1, message);
    s_mockCrossL2Inbox.setValidationSuccess(expectedIdentifier, expectedLogHash);

    // Call verifyMessage
    (uint256 timestampCommitted, bytes32[] memory hashedLeaves) =
      s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);

    // Verify results
    assertEq(timestampCommitted, timestamp);
    assertEq(hashedLeaves.length, 1);
    assertEq(hashedLeaves[0], SuperchainInterop._hashInteropMessage(message, ON_RAMP_ADDRESS));

    // Verify CrossL2Inbox was called correctly
    assertEq(s_mockCrossL2Inbox.getValidateMessageCallCount(), 1);
    MockCrossL2Inbox.ValidateMessageCall memory call = s_mockCrossL2Inbox.getValidateMessageCall(0);
    assertEq(call.identifier.origin, expectedIdentifier.origin);
    assertEq(call.identifier.blockNumber, expectedIdentifier.blockNumber);
    assertEq(call.identifier.logIndex, expectedIdentifier.logIndex);
    assertEq(call.identifier.timestamp, expectedIdentifier.timestamp);
    assertEq(call.identifier.chainId, expectedIdentifier.chainId);
    assertEq(call.msgHash, expectedLogHash);
  }

  function testFuzz_verifyReport_Success(
    uint256 blockNumber,
    uint256 logIndex,
    uint256 timestamp,
    uint64 sequenceNumber
  ) public {
    // Constrain inputs to avoid InvalidEncodingOfIdentifierInProofs
    vm.assume(blockNumber != 0);
    vm.assume(timestamp != 0);

    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, sequenceNumber);

    (bytes32[] memory proofs, Identifier memory expectedIdentifier) =
      _getValidProofsAndIdentifier(blockNumber, logIndex, timestamp);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    // Mock CrossL2Inbox validation
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, sequenceNumber, message);
    s_mockCrossL2Inbox.setValidationSuccess(expectedIdentifier, expectedLogHash);

    // Should succeed with any valid block number and log index
    (uint256 timestampCommitted, bytes32[] memory hashedLeaves) =
      s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);

    // Verify results
    assertEq(timestampCommitted, timestamp);
    assertEq(hashedLeaves.length, 1);
    assertEq(hashedLeaves[0], SuperchainInterop._hashInteropMessage(message, ON_RAMP_ADDRESS));

    // Verify CrossL2Inbox was called correctly
    assertEq(s_mockCrossL2Inbox.getValidateMessageCallCount(), 1);
    MockCrossL2Inbox.ValidateMessageCall memory call = s_mockCrossL2Inbox.getValidateMessageCall(0);
    assertEq(call.identifier.origin, expectedIdentifier.origin);
    assertEq(call.identifier.blockNumber, expectedIdentifier.blockNumber);
    assertEq(call.identifier.logIndex, expectedIdentifier.logIndex);
    assertEq(call.identifier.timestamp, expectedIdentifier.timestamp);
    assertEq(call.identifier.chainId, expectedIdentifier.chainId);
    assertEq(call.msgHash, expectedLogHash);
  }

  // Reverts

  function test_verifyReport_RevertWhen_InvalidMessageCount() public {
    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](0);

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](0),
      proofs: new bytes32[](5),
      proofFlagBits: 0
    });

    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.ReportMustContainExactlyOneMessage.selector));
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);

    report.messages = new Internal.Any2EVMRampMessage[](2);
    report.messages[0] = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);
    report.messages[1] = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 2);

    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.ReportMustContainExactlyOneMessage.selector));
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_InvalidSourceChainSelector() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_2, 1);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: new bytes32[](5),
      proofFlagBits: 0
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidSourceChainSelector.selector,
        SOURCE_CHAIN_SELECTOR_2,
        SOURCE_CHAIN_SELECTOR_1
      )
    );
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_InvalidDestChainSelector() public {
    uint64 invalidDestChainSelector = 99999;

    Internal.Any2EVMRampMessage memory message = Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256("messageId"),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
        destChainSelector: invalidDestChainSelector,
        sequenceNumber: 1,
        nonce: 1
      }),
      sender: abi.encode(address(0x1234)),
      data: abi.encode("test"),
      receiver: address(s_receiver),
      gasLimit: GAS_LIMIT,
      tokenAmounts: new Internal.Any2EVMTokenTransfer[](0)
    });

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: new bytes32[](5),
      proofFlagBits: 0
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidDestChainSelector.selector, invalidDestChainSelector, DEST_CHAIN_SELECTOR
      )
    );
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_InvalidProofsLength() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    // Test with too few proofs
    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: new bytes32[](4),
      proofFlagBits: 0
    });

    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidProofsWordLength.selector, 4, 5));
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);

    report.proofs = new bytes32[](6);

    vm.expectRevert(abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidProofsWordLength.selector, 6, 5));
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_InvalidSourceOnRamp() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    address wrongOnRamp = makeAddr("wrong_onramp");
    bytes32[] memory proofs = new bytes32[](5);
    proofs[0] = bytes32(uint256(uint160(wrongOnRamp))); // Wrong onRamp
    proofs[1] = bytes32(block.number); // blockNumber
    proofs[2] = bytes32(uint256(1)); // logIndex (can be zero)
    proofs[3] = bytes32(block.timestamp); // timestamp
    proofs[4] = bytes32(CHAIN_ID_1); // chainId

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidSourceOnRamp.selector, SOURCE_CHAIN_SELECTOR_1, wrongOnRamp
      )
    );
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_ChainIdNotConfigured() public {
    // Add a new source chain selector without chainId mapping
    uint64 newSourceChainSelector = 9999;
    address newOnRamp = makeAddr("new_onramp");

    // Configure the new source onramp
    OffRamp.SourceChainConfigArgs[] memory sourceChainConfigs = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: newSourceChainSelector,
      isEnabled: true,
      isRMNVerificationDisabled: false,
      onRamp: abi.encode(newOnRamp)
    });
    s_offRampOverSuperchainInterop.applySourceChainConfigUpdates(sourceChainConfigs);

    Internal.Any2EVMRampMessage memory message = _generateValidMessage(newSourceChainSelector, 1);

    // Identifier uses the new OnRamp
    bytes32[] memory proofs = new bytes32[](5);
    proofs[0] = bytes32(uint256(uint160(newOnRamp)));
    proofs[1] = bytes32(block.number); // blockNumber
    proofs[2] = bytes32(uint256(1)); // logIndex (can be zero)
    proofs[3] = bytes32(block.timestamp); // timestamp
    proofs[4] = bytes32(uint256(99)); // chainId (some non-zero value)

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: newSourceChainSelector,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ChainIdNotConfiguredForSelector.selector, newSourceChainSelector
      )
    );
    s_offRampOverSuperchainInterop.verifyMessage(newSourceChainSelector, report);
  }

  function test_verifyReport_RevertWhen_ChainIdMismatch() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    uint256 invalidChainId = 200;
    (bytes32[] memory proofs,) = _getValidProofsAndIdentifier(block.number, 1, 1);
    proofs[4] = bytes32(invalidChainId);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    // This should revert with SourceChainSelectorMismatch because the chainId doesn't match
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ChainIdMismatch.selector, SOURCE_CHAIN_SELECTOR_1, invalidChainId, CHAIN_ID_1
      )
    );
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyReport_RevertWhen_CrossL2InboxValidationFails() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);

    (bytes32[] memory proofs, Identifier memory expectedIdentifier) = _getValidProofsAndIdentifier(block.number, 1, 1);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 0
    });

    // Setup CrossL2Inbox to fail
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, 1, message);
    string memory validationError = "CrossL2Inbox validation failed";
    s_mockCrossL2Inbox.setValidationError(expectedIdentifier, expectedLogHash, validationError);

    // Expect revert from MockCrossL2Inbox
    vm.expectRevert(abi.encodeWithSelector(MockCrossL2Inbox.ValidationFailed.selector, validationError));
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }

  function test_verifyMessage_RevertWhen_ProofFlagBitsMustBeZero() public {
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);
    (bytes32[] memory proofs,) = _getValidProofsAndIdentifier(block.number, 1, block.timestamp);

    Internal.Any2EVMRampMessage[] memory messages = new Internal.Any2EVMRampMessage[](1);
    messages[0] = message;

    Internal.ExecutionReport memory report = Internal.ExecutionReport({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      messages: messages,
      offchainTokenData: new bytes[][](1),
      proofs: proofs,
      proofFlagBits: 1 // This should trigger ProofFlagBitsMustBeZero
    });

    vm.expectRevert(OffRampOverSuperchainInterop.ProofFlagBitsMustBeZero.selector);
    s_offRampOverSuperchainInterop.verifyMessage(SOURCE_CHAIN_SELECTOR_1, report);
  }
}
