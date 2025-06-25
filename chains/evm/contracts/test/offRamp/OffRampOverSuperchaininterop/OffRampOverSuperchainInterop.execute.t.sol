// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Client} from "../../../libraries/Client.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

contract OffRampOverSuperchainInterop_execute is OffRampOverSuperchainInteropSetup {
  function test_execute_SingleMessage() public {
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

    // Execute as allowed transmitter
    vm.stopPrank();
    vm.prank(s_transmitter1);
    s_offRamp.execute(report);

    // Verify state
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_execute_RevertWhen_UnauthorizedTransmitter() public {
    // Create message
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      100,
      0,
      new Client.EVMTokenAmount[](0),
      false
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

    // Execute as unauthorized address
    address unauthorizedTransmitter = makeAddr("unauthorized");
    vm.stopPrank();
    vm.prank(unauthorizedTransmitter);
    vm.expectRevert(OffRampOverSuperchainInterop.UnauthorizedTransmitter.selector);
    s_offRamp.execute(report);
  }

  function test_execute_GasEstimationSender() public {
    // Create message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      sequenceNumber,
      0,
      new Client.EVMTokenAmount[](0),
      false
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

    // Execute as GAS_ESTIMATION_SENDER
    vm.stopPrank();
    vm.prank(Internal.GAS_ESTIMATION_SENDER);
    s_offRamp.execute(report);

    // Verify execution succeeded
    assertEq(
      uint256(s_offRamp.getExecutionState(SOURCE_CHAIN_SELECTOR_1, sequenceNumber)),
      uint256(Internal.MessageExecutionState.SUCCESS)
    );
  }

  function test_execute_RevertWhen_InvalidDestChainSelector() public {
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

  function test_execute_RevertWhen_ChainIdNotConfigured() public {
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
}