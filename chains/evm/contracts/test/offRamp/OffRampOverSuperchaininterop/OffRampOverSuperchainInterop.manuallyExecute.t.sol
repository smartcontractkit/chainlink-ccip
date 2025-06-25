// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Client} from "../../../libraries/Client.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";
import {MaybeRevertMessageReceiver} from "../../helpers/receivers/MaybeRevertMessageReceiver.sol";

contract OffRampOverSuperchainInterop_manuallyExecute is OffRampOverSuperchainInteropSetup {
  function test_manuallyExecute_AfterThreshold() public {
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

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Create gas override
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride;

    // Execute manually
    vm.stopPrank();
    vm.prank(STRANGER);
    s_offRamp.manuallyExecute(report, gasOverride);

    // Verify execution succeeded
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_manuallyExecute_PreviouslyFailed() public {
    // Create message that will fail
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );
    message.receiver = address(s_reverting_receiver);

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

    // First execution - will fail
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify failure state
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.FAILURE)
    );

    // Fix the receiver
    message.receiver = address(s_receiver);
    report = _createExecutionReport(message, identifier, new bytes[](0));
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Manual execution should succeed even without waiting
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride;
    vm.stopPrank();
    vm.prank(STRANGER);
    s_offRamp.manuallyExecute(report, gasOverride);

    // Verify state changed to success
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_manuallyExecute_GasLimitOverride() public {
    // Create message with low gas limit
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );
    message.gasLimit = 100_000; // Low gas limit

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

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Create gas override with higher limit
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride = 
      OffRampOverSuperchainInterop.GasLimitOverride({
        receiverExecutionGasLimit: 300_000,
        tokenGasOverrides: new uint32[](0)
      });

    // Execute manually with gas override
    vm.stopPrank();
    vm.prank(STRANGER);
    s_offRamp.manuallyExecute(report, gasOverride);

    // Verify execution succeeded with overridden gas
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }


  function test_manuallyExecute_RevertWhen_BeforeThreshold() public {
    // Create message with sequenceNumber close to current timestamp
    // This ensures block.timestamp - sequenceNumber < threshold
    uint64 sequenceNumber = uint64(block.timestamp - 100); // 100 seconds ago, less than 500 threshold
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

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Don't warp time - still within threshold
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride;

    // Execute manually should revert
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ManualExecutionNotYetEnabled.selector,
        SOURCE_CHAIN_SELECTOR_1
      )
    );
    s_offRamp.manuallyExecute(report, gasOverride);
  }

  function test_manuallyExecute_RevertWhen_InvalidGasLimit() public {
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
    message.gasLimit = 200_000;

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

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Create gas override with lower limit
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride = 
      OffRampOverSuperchainInterop.GasLimitOverride({
        receiverExecutionGasLimit: 100_000, // Lower than original
        tokenGasOverrides: new uint32[](0)
      });

    // Execute manually should revert
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidManualExecutionGasLimit.selector,
        SOURCE_CHAIN_SELECTOR_1,
        message.header.messageId,
        100_000
      )
    );
    s_offRamp.manuallyExecute(report, gasOverride);
  }

  function test_manuallyExecute_RevertWhen_InvalidTokenGasOverride() public {
    // Create message with tokens
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({
      token: address(s_destTokens[0]),
      amount: 100
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
    message.tokenAmounts[0].destGasAmount = 50_000;

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
      new bytes[](1)
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Create gas override with lower token gas
    uint32[] memory tokenGasOverrides = new uint32[](1);
    tokenGasOverrides[0] = 25_000; // Lower than original

    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride = 
      OffRampOverSuperchainInterop.GasLimitOverride({
        receiverExecutionGasLimit: 0,
        tokenGasOverrides: tokenGasOverrides
      });

    // Execute manually should revert
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.InvalidManualExecutionTokenGasOverride.selector,
        message.header.messageId,
        0,
        50_000,
        25_000
      )
    );
    s_offRamp.manuallyExecute(report, gasOverride);
  }

  function test_manuallyExecute_RevertWhen_GasAmountCountMismatch() public {
    // Create message with 2 tokens
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
      new bytes[](2)
    );

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Create gas override with wrong number of token overrides
    uint32[] memory tokenGasOverrides = new uint32[](1); // Only 1 override for 2 tokens
    tokenGasOverrides[0] = 50_000;

    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride = 
      OffRampOverSuperchainInterop.GasLimitOverride({
        receiverExecutionGasLimit: 0,
        tokenGasOverrides: tokenGasOverrides
      });

    // Execute manually should revert
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ManualExecutionGasAmountCountMismatch.selector,
        message.header.messageId,
        sequenceNumber
      )
    );
    s_offRamp.manuallyExecute(report, gasOverride);
  }


  function test_manuallyExecute_SkipAlreadyExecuted() public {
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

    // Set mock validation
    _setMockCrossL2InboxValidMessage(message, ON_RAMP_ADDRESS_1);

    // First execution - success
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify success state
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );

    // Warp time beyond threshold
    vm.warp(block.timestamp + PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS + 1);

    // Manual execution should skip with event
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride;
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.SkippedAlreadyExecutedMessage(SOURCE_CHAIN_SELECTOR_1, sequenceNumber);
    vm.stopPrank();
    vm.prank(STRANGER);
    s_offRamp.manuallyExecute(report, gasOverride);
  }

  function test_manuallyExecute_RevertWhen_AlreadyFailed() public {
    // Create message that will fail
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      true
    );
    message.receiver = address(s_reverting_receiver);

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

    // First execution - will fail
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify failure state
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.FAILURE)
    );

    // Manual execution should revert
    OffRampOverSuperchainInterop.GasLimitOverride memory gasOverride;
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.ExecutionError.selector,
        message.header.messageId,
        abi.encodeWithSelector(
          OffRampOverSuperchainInterop.ReceiverError.selector,
          abi.encodeWithSelector(MaybeRevertMessageReceiver.CustomError.selector, "")
        )
      )
    );
    s_offRamp.manuallyExecute(report, gasOverride);
  }
}