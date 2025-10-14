// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {ReentrantCCV} from "../../helpers/ReentrantCCV.sol";
import {ExactGasReceiver} from "../../helpers/receivers/ExactGasReceiver.sol";
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
    bytes[] calldata ccvData,
    uint256 gasForCall
  ) external {
    address offRamp = address(i_offRamp);
    bytes memory payload = abi.encodeCall(i_offRamp.execute, (message, ccvs, ccvData));

    assembly {
      let success := call(gasForCall, offRamp, 0, add(payload, 0x20), mload(payload), 0, 0)
      if iszero(success) {
        returndatacopy(0, 0, returndatasize())
        revert(0, returndatasize())
      }
    }
  }
}

contract OffRamp_execute is OffRampSetup, OffRamp {
  uint256 internal constant PLENTY_OF_GAS = 1_000_000;

  GasBoundedExecuteCaller internal s_gasBoundedExecuteCaller;

  // We have the constructor here simply to access & test some internal functions.
  constructor()
    OffRamp(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: IRMNRemote(makeAddr("rmnRemote")),
        tokenAdminRegistry: makeAddr("tokenAdminRegistry")
      })
    )
  {}

  function setUp() public virtual override {
    super.setUp();

    s_gasBoundedExecuteCaller = new GasBoundedExecuteCaller(address(s_offRamp));

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Mock validateReport for default message structure.
    bytes memory encodedMessage = MessageV1Codec._encodeMessageV1(message);
    bytes32 messageHash = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICrossChainVerifierV1.verifyMessage, (address(s_offRamp), message, messageHash, abi.encode("mock ccv data"))
      ),
      abi.encode(true)
    );
  }

  function _getMessage() internal returns (MessageV1Codec.MessageV1 memory message) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      onRampAddress: abi.encodePacked(makeAddr("onRamp")),
      offRampAddress: abi.encodePacked(makeAddr("offRamp")),
      //
      finality: 0,
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  function _getReportComponents(
    MessageV1Codec.MessageV1 memory message
  ) internal view returns (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) {
    ccvData = new bytes[](1);
    ccvData[0] = abi.encode("mock ccv data");
    return (MessageV1Codec._encodeMessageV1(message), _arrayOf(s_defaultCCV), ccvData);
  }

  function test_execute() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Expect execution state change event.
    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.SUCCESS,
      ""
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, PLENTY_OF_GAS);

    // Verify final state is SUCCESS.
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS),
      uint256(
        s_offRamp.getExecutionState(
          message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );
  }

  function test_execute_RunsOutOfGasAndSetsStateToFailure() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Expect execution state change event.
    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      "" // empty because there is no error when a tx runs out of gas.
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, 85_000);

    // Verify final state is FAILURE.
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE),
      uint256(
        s_offRamp.getExecutionState(
          message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );
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
    bytes[] memory ccvData = new bytes[](2);

    _setGetCCVsReturnData(address(bytes20(message.receiver)), message.sourceChainSelector, ccvs, new address[](0), 0);

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      abi.encodeWithSelector(OffRamp.ReentrancyGuardReentrantCall.selector)
    );

    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, PLENTY_OF_GAS);
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
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_SourceChainNotEnabled() public {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    // Configure source chain as disabled.
    _applySourceConfig(abi.encode(makeAddr("onRamp")), false, defaultCCVs, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(OffRamp.SourceChainNotEnabled.selector, SOURCE_CHAIN_SELECTOR));
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidMessageDestChainSelector() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with wrong destination chain selector.
    message.destChainSelector = DEST_CHAIN_SELECTOR + 1; // Wrong destination.

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidMessageDestChainSelector.selector, message.destChainSelector));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidCCVDataLength() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory originalCcvs, bytes[] memory originalCcvData) =
      _getReportComponents(message);

    // Create mismatched array lengths.
    address[] memory ccvs = new address[](originalCcvs.length + 1);
    for (uint256 i = 0; i < originalCcvs.length; i++) {
      ccvs[i] = originalCcvs[i];
    }
    // Keep ccvData the same, creating mismatch.
    bytes[] memory ccvData = originalCcvData;

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidCCVDataLength.selector, ccvs.length, ccvData.length));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_SkippedAlreadyExecutedMessage() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Execute the message successfully first time.
    s_offRamp.execute(encodedMessage, ccvs, ccvData);

    // Verify it's in SUCCESS state.
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS),
      uint256(
        s_offRamp.getExecutionState(
          message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );

    // Try to execute the same message again - should revert.
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRamp.SkippedAlreadyExecutedMessage.selector,
        keccak256(encodedMessage),
        SOURCE_CHAIN_SELECTOR,
        message.sequenceNumber
      )
    );
    s_offRamp.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_InsufficientGasToCompleteTx_setsToFailure() public {
    uint256 gasForCall = 90_000;
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    // Mock validateReport to pass initial checks.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (address(s_offRamp), message, messageId, ccvData[0])),
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
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    GasBoundedExecuteCaller(Internal.GAS_ESTIMATION_SENDER).callExecute(encodedMessage, ccvs, ccvData, gasForCall);
  }

  function test_execute_RevertWhen_ExecuteSingleMessageFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes memory revertReason = "ExecuteSingleMessage failed";

    // Mock executeSingleMessage to revert.
    vm.mockCallRevert(address(s_offRamp), abi.encodeWithSelector(s_offRamp.executeSingleMessage.selector), revertReason);

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      revertReason
    );

    // The message should succeed but set execution state to FAILURE due to executeSingleMessage revert.
    // This verifies that execution failures are handled gracefully.
    s_offRamp.execute(encodedMessage, ccvs, ccvData);

    // Verify message state changed to FAILURE
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE),
      uint256(
        s_offRamp.getExecutionState(
          SOURCE_CHAIN_SELECTOR, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );
  }

  function test_padTo32() public pure {
    assertEq(OffRamp._padTo32(0), 0);
    assertEq(OffRamp._padTo32(1), 32);
    assertEq(OffRamp._padTo32(31), 32);
    assertEq(OffRamp._padTo32(32), 32);
    assertEq(OffRamp._padTo32(33), 64);
  }

  function test_execute_WithReceiver_HighBaseExecuteGas() public {
    callExecuteWithReceiver(50_000, 500, 20_000_000);
  }

  /// forge-config: default.fuzz.runs = 1000
  function testFuzz_execute_WithReceiver(uint32 _gasUsedByCCIPReceive, uint16 _calldataLength) public {
    uint256 gasUsedByCCIPReceive = bound(_gasUsedByCCIPReceive, 1_000, PLENTY_OF_GAS);
    uint256 calldataLength = bound(_calldataLength, 0, 1_000);

    callExecuteWithReceiver(gasUsedByCCIPReceive, calldataLength, 120_000);
  }

  function callExecuteWithReceiver(
    uint256 gasUsedByCCIPReceive, // The total gas that the receiver should consume during ccipReceive.
    uint256 calldataLength, // The length of the calldata passed to ccipReceive.
    uint256 baseExecuteGas // Accounts for logic preceeding and following the call to routeMessage.
  ) public {
    // Create message with exact gas receiver and calldata.
    MessageV1Codec.MessageV1 memory message = _getMessage();
    message.receiver = abi.encodePacked(address(new ExactGasReceiver(gasUsedByCCIPReceive)));
    bytes memory data = new bytes(calldataLength);
    assembly {
      let dataPtr := add(data, 0x20) // Skip the length field of the bytes array.
      let endPtr := add(dataPtr, calldataLength) // Calculate end pointer.

      // Fill the array with 0x01 bytes, filling 32 bytes at a time.
      // This forces the worst-case calldata gas scenario for the given calldata size.
      for { let i := dataPtr } lt(i, endPtr) { i := add(i, 32) } {
        mstore(i, 0x0101010101010101010101010101010101010101010101010101010101010101)
      }
    }

    // Test with 0 and 1 token transfers.
    for (uint256 tokenTransferLength = 0; tokenTransferLength < 2; ++tokenTransferLength) {
      message.sequenceNumber++;
      uint256 routeMessageCalldataLen =
        388 + _padTo32(message.sender.length) + _padTo32(calldataLength) + tokenTransferLength * 64;

      // gasForExecute reverses the gas computation peformed by the OffRamp when calling routeMessage.
      uint256 gasForExecute = baseExecuteGas
        + ((gasUsedByCCIPReceive * 64 / 63) + 16 * calldataLength + GAS_FOR_CALL_EXACT_CHECK) * 64 / 63
        + GAS_FOR_CALL_EXACT_CHECK + 2 * (16 * routeMessageCalldataLen);

      message.data = data;
      (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

      // Execute and verify final state is SUCCESS.
      s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, gasForExecute);
      assertEq(
        uint256(Internal.MessageExecutionState.SUCCESS),
        uint256(
          s_offRamp.getExecutionState(
            message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
          )
        )
      );
    }
  }
}
