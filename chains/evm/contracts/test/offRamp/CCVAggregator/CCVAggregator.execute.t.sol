// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../../../interfaces/ICCVOffRampV1.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

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

    bytes memory defaultCcvData = abi.encode("mock ccv data");
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.verifyMessage, (message, messageHash, defaultCcvData, Internal.MessageExecutionState.UNTOUCHED)
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

  function test_execute_runsOutOfGasAndSetsStateToFailure() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Expect execution state change event.
    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
      message.sourceChainSelector,
      message.sequenceNumber,
      keccak256(encodedMessage),
      Internal.MessageExecutionState.FAILURE,
      abi.encodeWithSelector(CCVAggregator.OutOfGas.selector)
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

  function test_execute_RevertWhen_ReentrancyGuardReentrantCall() public {
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

    vm.expectRevert(CCVAggregator.ReentrancyGuardReentrantCall.selector);
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
    // Configure source chain as disabled.
    _applySourceConfig(
      s_sourceRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), false, new address[](1), new address[](0)
    );

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
        CCVAggregator.SkippedAlreadyExecutedMessage.selector, SOURCE_CHAIN_SELECTOR, message.sequenceNumber
      )
    );
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_ReceiverCCV() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address receiver = makeAddr("receiver");
    address requiredCCV = makeAddr("requiredCCV");

    message.receiver = abi.encodePacked(receiver);

    // Set up receiver to require a specific CCV.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, _arrayOf(requiredCCV), new address[](0), 0);

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, requiredCCV, false));

    // Keep default CCV in report, but don't include the required CCV.
    (bytes memory encodedMsg, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    s_agg.execute(encodedMsg, ccvs, ccvData);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_PoolCCV() public {
    address poolRequiredCCV = makeAddr("poolRequiredCCV");
    address sourceToken = makeAddr("sourceToken");
    address token = makeAddr("token");
    address pool = makeAddr("pool");
    uint256 tokenAmount = 100;

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with token transfer.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: tokenAmount,
      sourcePoolAddress: abi.encodePacked(pool),
      sourceTokenAddress: abi.encodePacked(sourceToken),
      destTokenAddress: abi.encodePacked(token),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);

    // Mock token admin registry to return the pool.
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (token)), abi.encode(pool));

    // Mock pool supportsInterface for IPoolV2.
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool to require a specific CCV.
    vm.mockCall(
      pool,
      abi.encodeCall(IPoolV2.getRequiredCCVs, (token, SOURCE_CHAIN_SELECTOR, tokenAmount, "")),
      abi.encode(_arrayOf(poolRequiredCCV))
    );

    // Keep default CCV in report, but don't include the pool required CCV.

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, poolRequiredCCV, true));
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_LaneMandatedCCV() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address laneMandatedCCV = makeAddr("laneMandatedCCV");

    // Configure source chain with lane mandated CCV.
    CCVAggregator.SourceChainConfigArgs[] memory sourceChainConfigArgs = new CCVAggregator.SourceChainConfigArgs[](1);
    sourceChainConfigArgs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: _arrayOf(s_defaultCCV),
      laneMandatedCCVs: _arrayOf(laneMandatedCCV)
    });

    s_agg.applySourceChainConfigUpdates(sourceChainConfigArgs);

    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    // Report doesn't include the lane mandated CCV.

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, laneMandatedCCV, true));
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_OptionalCCVQuorumNotReached() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    address receiver = address(bytes20(message.receiver));

    address optionalCCV2 = makeAddr("optionalCCV2");
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = s_defaultCCV; // This will be found in the report.
    optionalCCVs[1] = optionalCCV2; // This won't be found.

    uint8 optionalThreshold = 2; // Need 2 optional CCVs.

    // Set up receiver to return optional CCVs with threshold 2.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, new address[](0), optionalCCVs, optionalThreshold);
    // Report only includes one CCV, but threshold requires 2.

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.OptionalCCVQuorumNotReached.selector, optionalThreshold, 1));
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_InsufficientGasToCompleteTx() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    // Mock validateReport to pass initial checks.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.verifyMessage, (message, messageId, ccvData[0], Internal.MessageExecutionState.UNTOUCHED)
      ),
      abi.encode(true)
    );

    // Mock executeSingleMessage to revert with NOT_ENOUGH_GAS_FOR_CALL_SIG.
    vm.mockCallRevert(
      address(s_agg),
      abi.encodeCall(s_agg.executeSingleMessage, (message, messageId)),
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    // Call from gas estimation sender to trigger the specific error handling. Since we use a contract
    // to set a custom gas limit, we need to etch the code into that address.
    vm.etch(Internal.GAS_ESTIMATION_SENDER, address(s_gasBoundedExecuteCaller).code);

    vm.expectRevert(
      abi.encodeWithSelector(
        CCVAggregator.InsufficientGasToCompleteTx.selector, CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG
      )
    );
    GasBoundedExecuteCaller(Internal.GAS_ESTIMATION_SENDER).callExecute(encodedMessage, ccvs, ccvData, 150_000);
  }

  function test_execute_RevertWhen_CCVValidationFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);
    bytes memory revertReason = "CCV validation failed";

    // Mock CCV validateReport to fail/revert.
    vm.mockCallRevert(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.verifyMessage, (message, messageId, ccvData[0], Internal.MessageExecutionState.UNTOUCHED)
      ),
      revertReason
    );

    vm.expectRevert(revertReason);
    s_agg.execute(encodedMessage, ccvs, ccvData);
  }

  function test_execute_RevertWhen_ExecuteSingleMessageFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    (bytes memory encodedMessage, address[] memory ccvs, bytes[] memory ccvData) = _getReportComponents(message);
    bytes32 messageId = keccak256(encodedMessage);

    bytes memory revertReason = "ExecuteSingleMessage failed";

    // Mock executeSingleMessage to revert.
    vm.mockCallRevert(address(s_agg), abi.encodeCall(s_agg.executeSingleMessage, (message, messageId)), revertReason);

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

contract ReentrantCCV is ICCVOffRampV1 {
  CCVAggregator internal immutable i_aggregator;

  constructor(
    address aggregator
  ) {
    i_aggregator = CCVAggregator(aggregator);
  }

  function verifyMessage(
    MessageV1Codec.MessageV1 memory message,
    bytes32, /* messageHash */
    bytes memory ccvData,
    Internal.MessageExecutionState /* originalState */
  ) external override {
    // Create a dummy report to trigger reentrancy.
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(this);
    bytes[] memory ccvDataArray = new bytes[](1);
    ccvDataArray[0] = ccvData;

    // This should trigger the reentrancy guard.
    i_aggregator.execute(MessageV1Codec._encodeMessageV1(message), ccvs, ccvDataArray);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId;
  }
}
