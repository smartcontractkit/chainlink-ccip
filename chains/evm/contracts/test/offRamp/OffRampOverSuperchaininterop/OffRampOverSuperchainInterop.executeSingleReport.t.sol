// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Client} from "../../../libraries/Client.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {INonceManager} from "../../../interfaces/INonceManager.sol";
import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";
import {MockCrossL2Inbox} from "../../helpers/MockCrossL2Inbox.sol";
import {ICrossL2Inbox} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";

contract OffRampOverSuperchainInterop_executeSingleReport is OffRampOverSuperchainInteropSetup {
  function test_executeSingleReport() public {
    // Create message
    uint64 sequenceNumber = 100;
    uint64 nonce = 1;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      nonce,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Calculate expected message hash
    bytes32 expectedMessageHash = SuperchainInterop._hashInteropMessage(message, ON_RAMP_ADDRESS_1);

    // Expect ExecutionStateChanged event with correct message hash
    // Check all fields except gas (last parameter)
    vm.expectEmit(true, true, true, false);
    emit OffRampOverSuperchainInterop.ExecutionStateChanged(
      SOURCE_CHAIN_SELECTOR_1,
      sequenceNumber,
      message.header.messageId,
      expectedMessageHash,
      Internal.MessageExecutionState.SUCCESS,
      "",
      0 // Gas value not checked
    );

    // Execute as allowed transmitter
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify state transition
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_executeSingleReport_OrderedMessage() public {
    // Create ordered message (nonce != 0)
    uint64 sequenceNumber = 100;
    uint64 nonce = 1; // Use nonce 1 as the first nonce
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      nonce,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Get initial nonce (should be 0)
    uint64 inboundNonceBefore = s_nonceManager.getInboundNonce(SOURCE_CHAIN_SELECTOR_1, message.sender);
    assertEq(inboundNonceBefore, 0);

    // Execute and verify nonce increment
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify execution succeeded
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );

    // Verify nonce was incremented
    uint64 inboundNonceAfter = s_nonceManager.getInboundNonce(SOURCE_CHAIN_SELECTOR_1, message.sender);
    assertEq(inboundNonceAfter, nonce);
  }

  function test_executeSingleReport_UnorderedMessage() public {
    // Create unordered message (nonce = 0)
    uint64 sequenceNumber = 100;
    uint64 nonce = 0;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      nonce,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Execute - should succeed without nonce check
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify execution succeeded
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_executeSingleReport_RevertWhen_InvalidDestChainSelector() public {
    // Create message with wrong dest chain selector
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );
    message.header.destChainSelector = DEST_CHAIN_SELECTOR + 1; // Wrong selector

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.MismatchedDestChainSelector.selector,
        DEST_CHAIN_SELECTOR + 1,
        DEST_CHAIN_SELECTOR + 1
      )
    );
    s_offRamp.execute(report);
  }

  function test_executeSingleReport_RevertWhen_InvalidSourceOnRamp() public {
    // Create message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report with wrong origin
    address wrongOnRamp = makeAddr("wrongOnRamp");
    Identifier memory identifier = _createIdentifier(
      wrongOnRamp, // Wrong OnRamp address
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidSourceOnRamp.selector,
        wrongOnRamp
      )
    );
    s_offRamp.execute(report);
  }

  function test_executeSingleReport_RevertWhen_ChainIdNotConfigured() public {
    // Setup a source chain without chain ID mapping
    // Already pranked as OWNER in BaseTest.setUp()
    OffRampOverSuperchainInterop.SourceChainConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.SourceChainConfigArgs[](1);
    configs[0] = OffRampOverSuperchainInterop.SourceChainConfigArgs({
      router: s_router,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_2,
      isEnabled: true,
      onRamp: ON_RAMP_ENCODED_2
    });
    s_offRamp.applySourceChainConfigUpdates(configs);
    // Don't set chain ID mapping

    // Create message from unconfigured chain
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_2,
      ON_RAMP_ADDRESS_2,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_2,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_2
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ChainIdNotConfiguredForSelector.selector,
        SOURCE_CHAIN_SELECTOR_2
      )
    );
    s_offRamp.execute(report);
  }

  function test_executeSingleReport_RevertWhen_ChainIdMismatch() public {
    // Create message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report with wrong chain ID
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_2 // Wrong chain ID
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.SourceChainSelectorMismatch.selector,
        SOURCE_CHAIN_ID_1,
        SOURCE_CHAIN_ID_2
      )
    );
    s_offRamp.execute(report);
  }

  function test_executeSingleReport_CrossL2InboxValidation() public {
    // Create message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Calculate expected parameters for validateMessage
    bytes memory logData = _encodeLogData(
      message.header.destChainSelector,
      message.header.sequenceNumber,
      message
    );
    bytes32 expectedMsgHash = keccak256(logData);

    // Set mock validation expectation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Expect the ExecutingMessage event from CrossL2Inbox
    vm.expectEmit(address(s_mockCrossL2Inbox));
    emit ICrossL2Inbox.ExecutingMessage(expectedMsgHash, identifier);

    // Execute
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);
  }

  function test_executeSingleReport_RevertWhen_TokenDataMismatch() public {
    // Create message with tokens
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({
      token: address(s_destTokens[0]),
      amount: 100
    });
    tokenAmounts[1] = Client.EVMTokenAmount({
      token: address(s_destTokens[1]),
      amount: 200
    });

    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      tokenAmounts,
      true
    );

    // Create execution report with wrong number of offchain token data
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](1) // Wrong length - should be 2
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.TokenDataMismatch.selector,
        SOURCE_CHAIN_SELECTOR_1,
        sequenceNumber
      )
    );
    s_offRamp.execute(report);
  }


  function test_executeSingleReport_RevertWhen_CrossL2InboxReverts() public {
    // Create message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Create execution report
    Identifier memory identifier = _createIdentifier(
      ON_RAMP_ADDRESS_1,
      block.number,
      1,
      block.timestamp,
      SOURCE_CHAIN_ID_1
    );
    SuperchainInterop.ExecutionReport memory report = _createExecutionReport(
      message,
      identifier,
      new bytes[](0)
    );

    // Set CrossL2Inbox to revert
    s_mockCrossL2Inbox.setShouldRevert(true);

    // Execute should revert
    vm.stopPrank();
    vm.prank(s_transmitter1);
    vm.expectRevert("CrossL2Inbox: validation failed");
    s_offRamp.execute(report);
  }
}