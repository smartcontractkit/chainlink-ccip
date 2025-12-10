// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";

import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {ReentrantCCV} from "../../helpers/ReentrantCCV.sol";

import {
  IAny2EVMMessageReceiver,
  MockReceiverSpoofVerificationFailed
} from "../../mocks/MockReceiverSpoofVerificationFailed.sol";
import {MockReceiverV2} from "../../mocks/MockReceiverV2.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";
import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";

contract GasBoundedExecuteCaller {
  OffRamp internal immutable i_offRamp;

  constructor(
    address offRamp
  ) {
    i_offRamp = OffRamp(offRamp);
  }

  function callExecute(
    bytes memory message,
    address[] calldata ccvs,
    bytes[] calldata verifierResults,
    uint256 gasForCall
  ) external {
    address offRamp = address(i_offRamp);
    bytes memory payload = abi.encodeCall(i_offRamp.execute, (message, ccvs, verifierResults));

    assembly {
      let success := call(gasForCall, offRamp, 0, add(payload, 0x20), mload(payload), 0, 0)
      if iszero(success) {
        returndatacopy(0, 0, returndatasize())
        revert(0, returndatasize())
      }
    }
  }
}

contract OffRamp_execute is OffRampSetup {
  uint256 internal constant PLENTY_OF_GAS = 1_000_000;

  GasBoundedExecuteCaller internal s_gasBoundedExecuteCaller;

  function setUp() public virtual override {
    super.setUp();

    s_gasBoundedExecuteCaller = new GasBoundedExecuteCaller(address(s_offRamp));

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Mock validateReport for default message structure.
    bytes memory encodedMessage = MessageV1Codec._encodeMessageV1(message);
    bytes32 messageHash = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(
        ICrossChainVerifierResolver.getInboundImplementation.selector, abi.encode("mock verifier results")
      ),
      abi.encode(s_defaultCCV)
    );
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageHash, abi.encode("mock verifier results"))),
      abi.encode(true)
    );
  }

  function _getMessage() internal returns (MessageV1Codec.MessageV1 memory message) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: 200_000,
      ccipReceiveGasLimit: 0,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: ON_RAMP,
      offRampAddress: abi.encodePacked(s_offRamp),
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  function _getReportComponents(
    MessageV1Codec.MessageV1 memory message
  ) internal view returns (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) {
    verifierResults = new bytes[](1);
    verifierResults[0] = abi.encode("mock verifier results");
    return (MessageV1Codec._encodeMessageV1(message), _arrayOf(s_defaultCCV), verifierResults);
  }

  function test_execute() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    // Expect execution state change event.
    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.SUCCESS,
      ""
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);

    // Verify final state is SUCCESS.
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(keccak256(encodedMessage)))
    );
  }

  function test_execute_WithReceiver() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    MockReceiverV2 mock = new MockReceiverV2(_arrayOf(s_defaultCCV), new address[](0), 0);
    message.receiver = abi.encodePacked(address(mock)); // Add receiver to message.
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    // Set OffRamp as a valid OffRamp on the Router.
    Router.OffRamp[] memory newRamps = new Router.OffRamp[](1);
    newRamps[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), newRamps);

    // Expect execution state change event.
    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.SUCCESS,
      ""
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);

    // Verify final state is SUCCESS.
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(keccak256(encodedMessage)))
    );
  }

  function test_execute_RunsOutOfGasAndSetsStateToFailure() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, 90000);

    // Verify final state is FAILURE.
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(keccak256(encodedMessage)))
    );
  }

  function test_execute_RevertWhen_InvalidReceiverLength() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes memory shortReceiver = abi.encodePacked(address(bytes20(message.receiver)));
    assembly {
      mstore(shortReceiver, 19)
    }
    message.receiver = shortReceiver;

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, shortReceiver));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  /// @dev This test is a bit undeterministic as the lowGas value needs to be low enough to trigger the revert, but high enough that
  /// it consistently reverts when running forge test, and forge coverage.
  /// Beacuse of this behavior it might not be worth it to keep this test, eventhough it is addressing missing coverage.
  function test_execute_RevertWhen_InsufficientGasToCompleteTx_BufferCheck() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getInboundImplementation.selector, verifierResults[0]),
      abi.encode(s_defaultCCV)
    );
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      abi.encode(true)
    );

    // Use a gas limit that is low enough to trip the buffer check but high enough to cover call overhead.
    uint256 lowGas = 80000; // MAX_GAS_BUFFER_TO_UPDATE_STATE is 12k

    // Lower-level call so we can check the selector without depending on encoded gasleft().
    (bool success, bytes memory err) = address(s_gasBoundedExecuteCaller).call{gas: lowGas}(
      abi.encodeCall(GasBoundedExecuteCaller.callExecute, (encodedMessage, ccvs, verifierResults, lowGas))
    );
    assertFalse(success);
    assertTrue(err.length >= 4, "no revert data");
    assertEq(bytes4(err), OffRamp.InsufficientGasToCompleteTx.selector);
  }

  function test_execute_ReentrancyGuardReentrantCall_Fails() public {
    // Create a malicious CCV that will call back into execute during validation.
    ReentrantCCV maliciousCCV = new ReentrantCCV(address(s_offRamp));

    // Update report to use malicious CCV.
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage,,) = _getReportComponents(message);

    address[] memory ccvs = new address[](2);
    ccvs[0] = address(maliciousCCV);
    ccvs[1] = s_defaultCCV;
    bytes[] memory verifierResults = new bytes[](2);

    _setGetCCVsReturnData(address(bytes20(message.receiver)), message.sourceChainSelector, ccvs, new address[](0), 0);

    vm.expectEmit(true, true, true, false);
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.VERIFICATION_FAILED,
      ""
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);

    assertEq(
      uint256(Internal.MessageExecutionState.VERIFICATION_FAILED),
      uint256(s_offRamp.getExecutionState(keccak256(encodedMessage)))
    );
  }

  function test_execute_InsufficientGasToCompleteTx_setsToFailure() public {
    uint256 gasForCall = 90_000;
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    // Mock validateReport to pass initial checks.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      abi.encode(true)
    );

    // Mock executeSingleMessage to revert with NOT_ENOUGH_GAS_FOR_CALL_SIG.
    vm.mockCallRevert(
      address(s_offRamp),
      abi.encodeWithSelector(s_offRamp.executeSingleMessage.selector),
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    // Call from gas estimation sender to trigger the specific error handling. Since we use a contract
    // to set a custom gas limit, we need to etch the code into that address.
    vm.etch(Internal.GAS_ESTIMATION_SENDER, address(s_gasBoundedExecuteCaller).code);

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    GasBoundedExecuteCaller(Internal.GAS_ESTIMATION_SENDER).callExecute(
      encodedMessage, ccvs, verifierResults, gasForCall
    );
  }

  function test_execute_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed for source chain.
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(SOURCE_CHAIN_SELECTOR))),
      abi.encode(true)
    );

    vm.expectRevert(abi.encodeWithSelector(OffRamp.CursedByRMN.selector, SOURCE_CHAIN_SELECTOR));
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_SourceChainNotEnabled() public {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    // Configure source chain as disabled.
    _applySourceConfig(ON_RAMP, false, defaultCCVs, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(OffRamp.SourceChainNotEnabled.selector, SOURCE_CHAIN_SELECTOR));
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidOnRamp() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with wrong onRamp address.
    message.onRampAddress = abi.encodePacked(makeAddr("invalid onRamp"));

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidOnRamp.selector, message.onRampAddress));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidOffRamp() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    message.offRampAddress = abi.encodePacked(makeAddr("invalid offRamp"));
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidOffRamp.selector, s_offRamp, message.offRampAddress));
    s_offRamp.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_InvalidOffRamp_InvalidLengthOtherwiseCorrectRamp() public {
    // This address has zero's in the first bytes, so if length is not checked properly,
    // it would match the correct offRamp address after casting.
    address offRampWithZeroBytes = 0x22222290dD7278AA3dDD389cc1E1D165cC4bAF00;
    vm.etch(offRampWithZeroBytes, address(s_offRamp).code);

    // Since etching doesn't copy state we have to set the source chain config again for this offRamp.
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    vm.startPrank(address(0));
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = ON_RAMP;

    OffRamp.SourceChainConfigArgs[] memory updates = new OffRamp.SourceChainConfigArgs[](1);
    updates[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCV: defaultCCVs,
      laneMandatedCCVs: new address[](0)
    });
    OffRamp(offRampWithZeroBytes).applySourceChainConfigUpdates(updates);

    // We remove the first byte to make the offRamp address invalid, but it would still match if length wasn't checked.
    // This is because the casting will add leading zeros to make it 20 bytes again.
    bytes memory offRampAddress = abi.encodePacked(offRampWithZeroBytes);
    assembly {
      mstore(offRampAddress, 19)
    }

    // Assert casting would have the same result.
    assertEq(address(bytes20(offRampAddress)), offRampWithZeroBytes);

    MessageV1Codec.MessageV1 memory message = _getMessage();
    message.offRampAddress = offRampAddress;
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InvalidOffRamp.selector, offRampWithZeroBytes, message.offRampAddress)
    );
    OffRamp(offRampWithZeroBytes).execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_InvalidMessageDestChainSelector() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with wrong destination chain selector.
    message.destChainSelector = DEST_CHAIN_SELECTOR + 1; // Wrong destination.

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidMessageDestChainSelector.selector, message.destChainSelector));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidVerifierResultsLength() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory originalCcvs, bytes[] memory originalVerifierResults) =
      _getReportComponents(message);

    // Create mismatched array lengths.
    address[] memory ccvs = new address[](originalCcvs.length + 1);
    for (uint256 i = 0; i < originalCcvs.length; i++) {
      ccvs[i] = originalCcvs[i];
    }
    // Keep verifierResults the same, creating mismatch.
    bytes[] memory verifierResults = originalVerifierResults;

    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InvalidVerifierResultsLength.selector, ccvs.length, verifierResults.length)
    );
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, verifierResults, PLENTY_OF_GAS);
  }

  function test_execute_SetsVerificationFailedAndCanRetry() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getInboundImplementation.selector, verifierResults[0]),
      abi.encode(s_defaultCCV)
    );
    bytes memory revertReason = "CCV validation failed";
    vm.mockCallRevert(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      revertReason
    );

    bytes memory expectedErr =
      abi.encodeWithSelector(OffRamp.VerificationFailed.selector, s_defaultCCV, s_defaultCCV, uint256(0), revertReason);
    if (expectedErr.length > 132) {
      assembly {
        mstore(expectedErr, 132)
      }
    }

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      messageId,
      Internal.MessageExecutionState.VERIFICATION_FAILED,
      expectedErr
    );

    s_offRamp.execute(encodedMessage, ccvs, verifierResults);
    assertEq(
      uint256(Internal.MessageExecutionState.VERIFICATION_FAILED), uint256(s_offRamp.getExecutionState(messageId))
    );

    vm.clearMockedCalls();
    // Re-install mocks cleared above.
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(SOURCE_CHAIN_SELECTOR))),
      abi.encode(false)
    );
    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getInboundImplementation.selector, verifierResults[0]),
      abi.encode(s_defaultCCV)
    );
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      abi.encode(true)
    );

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector, message.messageNumber, messageId, Internal.MessageExecutionState.SUCCESS, ""
    );

    s_offRamp.execute(encodedMessage, ccvs, verifierResults);
    assertEq(uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(messageId)));
  }

  function test_execute_SpoofedVerificationFailedFromReceiverSetsFailure() public {
    MockReceiverSpoofVerificationFailed spoofingReceiver =
      new MockReceiverSpoofVerificationFailed(_arrayOf(s_defaultCCV), new address[](0), 0);

    MessageV1Codec.MessageV1 memory message = _getMessage();
    message.receiver = abi.encodePacked(address(spoofingReceiver));
    message.data = "non-empty";
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getInboundImplementation.selector, verifierResults[0]),
      abi.encode(s_defaultCCV)
    );
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      abi.encode(true)
    );

    bytes memory spoofed =
      abi.encodeWithSelector(OffRamp.VerificationFailed.selector, address(1), address(2), uint256(0), bytes("spoof"));

    // If the router were to actually call the receiver, it would see this revert.
    vm.mockCallRevert(
      address(spoofingReceiver),
      abi.encodeCall(
        IAny2EVMMessageReceiver.ccipReceive,
        (
          Client.Any2EVMMessage({
            messageId: messageId,
            sourceChainSelector: message.sourceChainSelector,
            sender: message.sender,
            data: message.data,
            destTokenAmounts: new Client.EVMTokenAmount[](0)
          })
        )
      ),
      spoofed
    );

    // Router returns the spoofed VerificationFailed payload, OffRamp should classify as ReceiverError => FAILURE.
    vm.mockCall(
      address(s_sourceRouter),
      abi.encodeCall(
        IRouter.routeMessage,
        (
          Client.Any2EVMMessage({
            messageId: messageId,
            sourceChainSelector: message.sourceChainSelector,
            sender: message.sender,
            data: message.data,
            destTokenAmounts: new Client.EVMTokenAmount[](0)
          }),
          GAS_FOR_CALL_EXACT_CHECK,
          message.ccipReceiveGasLimit,
          address(bytes20(message.receiver))
        )
      ),
      abi.encode(false, spoofed, uint256(0))
    );

    bytes memory expectedErr = abi.encodeWithSelector(OffRamp.ReceiverError.selector, spoofed);
    // Truncate returned data to 132 bytes, 4 bytes + 4 words.
    if (expectedErr.length > 132) {
      assembly {
        mstore(expectedErr, 132)
      }
    }

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector, message.messageNumber, messageId, Internal.MessageExecutionState.FAILURE, expectedErr
    );

    s_offRamp.execute(encodedMessage, ccvs, verifierResults);
    assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(messageId)));
  }

  function test_execute_InboundImplementationNotFoundSetsFailureNotVerificationFailed() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getInboundImplementation.selector, verifierResults[0]),
      abi.encode(address(0))
    );

    bytes memory expectedErr =
      abi.encodeWithSelector(OffRamp.InboundImplementationNotFound.selector, s_defaultCCV, verifierResults[0]);
    if (expectedErr.length > 132) {
      assembly {
        mstore(expectedErr, 132)
      }
    }

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector, message.messageNumber, messageId, Internal.MessageExecutionState.FAILURE, expectedErr
    );

    s_offRamp.execute(encodedMessage, ccvs, verifierResults);

    assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(messageId)));
  }

  function test_execute_RevertWhen_SkippedAlreadyExecutedMessage() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);

    // Execute the message successfully first time.
    s_offRamp.execute(encodedMessage, ccvs, verifierResults);

    bytes32 messageId = keccak256(encodedMessage);

    // Verify it's in SUCCESS state.
    assertEq(uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(messageId)));

    // Try to execute the same message again - should revert.
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRamp.SkippedAlreadyExecutedMessage.selector,
        keccak256(encodedMessage),
        SOURCE_CHAIN_SELECTOR,
        message.messageNumber
      )
    );
    s_offRamp.execute(encodedMessage, ccvs, verifierResults);
  }

  function test_execute_RevertWhen_ExecuteSingleMessageFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory verifierResults) = _getReportComponents(message);
    bytes memory revertReason = "ExecuteSingleMessage failed";

    // Mock executeSingleMessage to revert.
    vm.mockCallRevert(address(s_offRamp), abi.encodeWithSelector(s_offRamp.executeSingleMessage.selector), revertReason);

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.messageNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      revertReason
    );

    // The message should succeed but set execution state to FAILURE due to executeSingleMessage revert.
    // This verifies that execution failures are handled gracefully.
    s_offRamp.execute(encodedMessage, ccvs, verifierResults);

    // Verify message state changed to FAILURE
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(keccak256(encodedMessage)))
    );
  }
}
