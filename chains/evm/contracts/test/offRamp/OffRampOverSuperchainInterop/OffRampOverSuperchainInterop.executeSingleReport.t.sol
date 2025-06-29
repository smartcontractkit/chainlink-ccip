// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {Router} from "../../../Router.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";

import {LogMessageDataReceiver} from "../../helpers/receivers/LogMessageDataReceiver.sol";
import {MockCrossL2Inbox} from "../../mocks/MockCrossL2Inbox.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract OffRampOverSuperchainInterop_executeSingleReport is OffRampOverSuperchainInteropSetup {
  function setUp() public virtual override {
    super.setUp();

    // Authorize OffRampOverSuperchainInterop to call NonceManager
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_offRampOverSuperchainInterop);
    s_inboundNonceManager.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    // Set up Router to accept OffRampOverSuperchainInterop
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](0);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] =
      Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR_1, offRamp: address(s_offRampOverSuperchainInterop)});
    s_destRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }

  function test_executeSingleReport() public {
    // Setup receiver
    LogMessageDataReceiver receiver = new LogMessageDataReceiver();

    // Generate message
    Internal.Any2EVMRampMessage memory message = _generateValidMessage(SOURCE_CHAIN_SELECTOR_1, 1);
    message.receiver = address(receiver);

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

    // Setup CrossL2Inbox mock for successful validation
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, 1, message);
    s_mockCrossL2Inbox.setValidationSuccess(expectedIdentifier, expectedLogHash);

    // Execute the report as transmitter
    vm.stopPrank();
    vm.prank(s_validTransmitters[0]);

    vm.expectEmit();
    emit LogMessageDataReceiver.MessageReceived(message.data);

    s_offRampOverSuperchainInterop.executeSingleReport(report, new OffRamp.GasLimitOverride[](0));

    // Verify message was executed
    assertTrue(
      Internal.MessageExecutionState.SUCCESS
        == s_offRampOverSuperchainInterop.getExecutionState(SOURCE_CHAIN_SELECTOR_1, 1)
    );
  }

  function test_executeSingleReport_RevertWhen_ValidationFails() public {
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

    // Setup CrossL2Inbox to fail the validation
    bytes32 expectedLogHash = _getLogHash(DEST_CHAIN_SELECTOR, 1, message);
    string memory validationError = "CrossL2Inbox validation failed";
    s_mockCrossL2Inbox.setValidationError(expectedIdentifier, expectedLogHash, validationError);

    // Execute the report as transmitter
    vm.stopPrank();
    vm.prank(s_validTransmitters[0]);

    // Expect revert from CrossL2Inbox validation failure, this revert is not caught by OffRamp
    vm.expectRevert(abi.encodeWithSelector(MockCrossL2Inbox.ValidationFailed.selector, validationError));
    s_offRampOverSuperchainInterop.executeSingleReport(report, new OffRamp.GasLimitOverride[](0));
  }
}
