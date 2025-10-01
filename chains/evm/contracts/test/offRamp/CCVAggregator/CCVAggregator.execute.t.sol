// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {ReentrantCCV} from "../../helpers/ReentrantCCV.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";
import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";

contract GasBoundedExecuteCaller {
  CCVAggregator internal immutable i_aggregator;

  constructor(
    address aggregator
  ) {
    i_aggregator = CCVAggregator(aggregator);
  }

  function callExecute(
    bytes memory message,
    address[] calldata ccvs,
    bytes[] calldata ccvData,
    uint256 gasForCall
  ) external {
    address ccvAggregator = address(i_aggregator);
    bytes memory payload = abi.encodeCall(i_aggregator.execute, (message, ccvs, ccvData));

    assembly {
      let success := call(gasForCall, ccvAggregator, 0, add(payload, 0x20), mload(payload), 0, 0)
      if iszero(success) {
        returndatacopy(0, 0, returndatasize())
        revert(0, returndatasize())
      }
    }
  }
}

contract CCVAggregator_execute is CCVAggregatorSetup {
  uint256 internal constant PLENTY_OF_GAS = 1_000_000;

  GasBoundedExecuteCaller internal s_gasBoundedExecuteCaller;

  function setUp() public virtual override {
    super.setUp();

    s_gasBoundedExecuteCaller = new GasBoundedExecuteCaller(address(s_agg));

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Mock validateReport for default message structure.
    bytes memory encodedMessage = MessageV1Codec._encodeMessageV1(message);
    bytes32 messageHash = keccak256(encodedMessage);

    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICrossChainVerifierV1.verifyMessage, (address(s_agg), message, messageHash, abi.encode("mock ccv data"))
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
    emit CCVAggregator.ExecutionStateChanged(
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
        s_agg.getExecutionState(
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
    emit CCVAggregator.ExecutionStateChanged(
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
        s_agg.getExecutionState(
          message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );
  }

  function test_execute_ReentrancyGuardReentrantCall_Fails() public {
    // Create a malicious CCV that will call back into execute during validation.
    ReentrantCCV maliciousCCV = new ReentrantCCV(address(s_agg));

    // Update report to use malicious CCV.
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage,,) = _getReportComponents(message);

    address[] memory ccvs = new address[](2);
    ccvs[0] = address(maliciousCCV);
    ccvs[1] = s_defaultCCV;
    bytes[] memory ccvData = new bytes[](2);

    _setGetCCVsReturnData(address(bytes20(message.receiver)), message.sourceChainSelector, ccvs, new address[](0), 0);

    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      abi.encodeWithSelector(CCVAggregator.ReentrancyGuardReentrantCall.selector)
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

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.CursedByRMN.selector, SOURCE_CHAIN_SELECTOR));
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_SourceChainNotEnabled() public {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    // Configure source chain as disabled.
    _applySourceConfig(abi.encode(makeAddr("onRamp")), false, defaultCCVs, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.SourceChainNotEnabled.selector, SOURCE_CHAIN_SELECTOR));
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    s_gasBoundedExecuteCaller.callExecute(encodedMsg, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_InvalidMessageDestChainSelector() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with wrong destination chain selector.
    message.destChainSelector = DEST_CHAIN_SELECTOR + 1; // Wrong destination.

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    vm.expectRevert(
      abi.encodeWithSelector(CCVAggregator.InvalidMessageDestChainSelector.selector, message.destChainSelector)
    );
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

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.InvalidCCVDataLength.selector, ccvs.length, ccvData.length));
    s_gasBoundedExecuteCaller.callExecute(encodedMessage, ccvs, ccvData, PLENTY_OF_GAS);
  }

  function test_execute_RevertWhen_SkippedAlreadyExecutedMessage() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Execute the message successfully first time.
    s_agg.execute(encodedMessage, ccvs, ccvData);

    // Verify it's in SUCCESS state.
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS),
      uint256(
        s_agg.getExecutionState(
          message.sourceChainSelector, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );

    // Try to execute the same message again - should revert.
    vm.expectRevert(
      abi.encodeWithSelector(
        CCVAggregator.SkippedAlreadyExecutedMessage.selector,
        keccak256(encodedMessage),
        SOURCE_CHAIN_SELECTOR,
        message.sequenceNumber
      )
    );
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_InsufficientGasToCompleteTx_setsToFailure() public {
    uint256 gasForCall = 90_000;
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    // Mock validateReport to pass initial checks.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (address(s_agg), message, messageId, ccvData[0])),
      abi.encode(true)
    );

    // Mock executeSingleMessage to revert with NOT_ENOUGH_GAS_FOR_CALL_SIG.
    vm.mockCallRevert(
      address(s_agg),
      abi.encodeWithSelector(s_agg.executeSingleMessage.selector),
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    // Call from gas estimation sender to trigger the specific error handling. Since we use a contract
    // to set a custom gas limit, we need to etch the code into that address.
    vm.etch(Internal.GAS_ESTIMATION_SENDER, address(s_gasBoundedExecuteCaller).code);

    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
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
    vm.mockCallRevert(address(s_agg), abi.encodeWithSelector(s_agg.executeSingleMessage.selector), revertReason);

    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      revertReason
    );

    // The message should succeed but set execution state to FAILURE due to executeSingleMessage revert.
    // This verifies that execution failures are handled gracefully.
    s_agg.execute(encodedMessage, ccvs, ccvData);

    // Verify message state changed to FAILURE
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE),
      uint256(
        s_agg.getExecutionState(
          SOURCE_CHAIN_SELECTOR, message.sequenceNumber, message.sender, address(bytes20(message.receiver))
        )
      )
    );
  }
}
