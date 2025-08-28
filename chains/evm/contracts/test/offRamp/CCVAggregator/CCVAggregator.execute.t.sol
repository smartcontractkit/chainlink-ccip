// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../../../interfaces/ICCVOffRampV1.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

contract CCVAggregator_execute is CCVAggregatorSetup {
  function setUp() public virtual override {
    super.setUp();

    // Mock validateReport for default message structure
    CCVAggregator.AggregatedReport memory defaultReport = _getReport();
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.validateReport,
        (
          defaultReport.message,
          keccak256(abi.encode(defaultReport.message)),
          defaultReport.ccvData[0],
          Internal.MessageExecutionState.UNTOUCHED
        )
      ),
      abi.encode(true)
    );
  }

  function _getReport() internal returns (CCVAggregator.AggregatedReport memory) {
    Internal.Any2EVMMessage memory message = Internal.Any2EVMMessage({
      header: Internal.Header({
        messageId: keccak256("test-message-1"),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: 1
      }),
      sender: abi.encode(makeAddr("sender")),
      data: "",
      receiver: makeAddr("receiver"),
      gasLimit: 0,
      tokenAmounts: new Internal.TokenTransfer[](0)
    });

    bytes[] memory ccvData = new bytes[](1);
    ccvData[0] = abi.encode("mock ccv data");

    return CCVAggregator.AggregatedReport({message: message, ccvs: _arrayOf(s_defaultCCV), ccvData: ccvData});
  }

  function test_execute() public {
    CCVAggregator.AggregatedReport memory report = _getReport();

    // Expect execution state change event
    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
      report.message.header.sourceChainSelector,
      report.message.header.sequenceNumber,
      report.message.header.messageId,
      Internal.MessageExecutionState.SUCCESS,
      ""
    );

    s_agg.execute(report);

    // Verify final state is SUCCESS
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS),
      uint256(
        s_agg.getExecutionState(
          report.message.header.sourceChainSelector,
          report.message.header.sequenceNumber,
          report.message.sender,
          report.message.receiver
        )
      )
    );
  }

  function test_execute_RevertWhen_ReentrancyGuardReentrantCall() public {
    // Create a malicious CCV that will call back into execute during validation
    ReentrantCCV maliciousCCV = new ReentrantCCV(address(s_agg));

    // Update report to use malicious CCV
    CCVAggregator.AggregatedReport memory report = _getReport();
    report.ccvs = new address[](2);
    report.ccvs[0] = address(maliciousCCV);
    report.ccvs[1] = s_defaultCCV;
    report.ccvData = new bytes[](2);

    _setGetCCVsReturnData(
      report.message.receiver, report.message.header.sourceChainSelector, report.ccvs, new address[](0), 0
    );

    vm.expectRevert(CCVAggregator.ReentrancyGuardReentrantCall.selector);
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed for source chain
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(SOURCE_CHAIN_SELECTOR))),
      abi.encode(true)
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.CursedByRMN.selector, SOURCE_CHAIN_SELECTOR));
    s_agg.execute(_getReport());
  }

  function test_execute_RevertWhen_SourceChainNotEnabled() public {
    // Configure source chain as disabled
    _applySourceConfig(
      s_sourceRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), false, new address[](1), new address[](0)
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.SourceChainNotEnabled.selector, SOURCE_CHAIN_SELECTOR));
    s_agg.execute(_getReport());
  }

  function test_execute_RevertWhen_InvalidMessageDestChainSelector() public {
    CCVAggregator.AggregatedReport memory report = _getReport();
    // Modify message with wrong destination chain selector
    report.message.header.destChainSelector = DEST_CHAIN_SELECTOR + 1; // Wrong destination

    vm.expectRevert(
      abi.encodeWithSelector(
        CCVAggregator.InvalidMessageDestChainSelector.selector, report.message.header.destChainSelector
      )
    );
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_InvalidCCVDataLength() public {
    CCVAggregator.AggregatedReport memory report = _getReport();

    // Modify report to have mismatched array lengths
    report.ccvs = new address[](report.ccvs.length + 1);
    // Keep ccvData the same, creating mismatch

    vm.expectRevert(
      abi.encodeWithSelector(CCVAggregator.InvalidCCVDataLength.selector, report.ccvs.length, report.ccvData.length)
    );
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_SkippedAlreadyExecutedMessage() public {
    CCVAggregator.AggregatedReport memory report = _getReport();

    // Execute the message successfully first time
    s_agg.execute(report);

    // Verify it's in SUCCESS state
    assertEq(
      uint256(Internal.MessageExecutionState.SUCCESS),
      uint256(
        s_agg.getExecutionState(
          report.message.header.sourceChainSelector,
          report.message.header.sequenceNumber,
          report.message.sender,
          report.message.receiver
        )
      )
    );

    // Try to execute the same message again - should revert
    vm.expectRevert(
      abi.encodeWithSelector(
        CCVAggregator.SkippedAlreadyExecutedMessage.selector,
        SOURCE_CHAIN_SELECTOR,
        report.message.header.sequenceNumber
      )
    );
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_InvalidNumberOfTokens() public {
    CCVAggregator.AggregatedReport memory report = _getReport();

    // Modify message with multiple tokens (more than 1)
    report.message.tokenAmounts = new Internal.TokenTransfer[](2);

    vm.expectRevert(
      abi.encodeWithSelector(CCVAggregator.InvalidNumberOfTokens.selector, report.message.tokenAmounts.length)
    );
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_ReceiverCCV() public {
    address receiver = makeAddr("receiver");
    address requiredCCV = makeAddr("requiredCCV");

    // Set up receiver to require a specific CCV
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, _arrayOf(requiredCCV), new address[](0), 0);

    CCVAggregator.AggregatedReport memory report = _getReport();
    report.message.receiver = receiver;
    // Keep default CCV in report, but don't include the required CCV

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, requiredCCV, false));
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_PoolCCV() public {
    address poolRequiredCCV = makeAddr("poolRequiredCCV");
    address token = makeAddr("token");
    address pool = makeAddr("pool");
    uint256 tokenAmount = 100;

    CCVAggregator.AggregatedReport memory report = _getReport();
    // Modify message with token transfer
    Internal.TokenTransfer[] memory tokenAmounts = new Internal.TokenTransfer[](1);
    tokenAmounts[0] = Internal.TokenTransfer({
      sourcePoolAddress: abi.encode(pool),
      destTokenAddress: token,
      extraData: "",
      amount: tokenAmount
    });
    report.message.tokenAmounts = tokenAmounts;

    // Mock token admin registry to return the pool
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (token)), abi.encode(pool));

    // Mock pool supportsInterface for IPoolV2
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool to require a specific CCV
    vm.mockCall(
      pool,
      abi.encodeCall(IPoolV2.getRequiredCCVs, (token, SOURCE_CHAIN_SELECTOR, tokenAmount, "")),
      abi.encode(_arrayOf(poolRequiredCCV))
    );

    // Keep default CCV in report, but don't include the pool required CCV

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, poolRequiredCCV, true));
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_RequiredCCVMissing_LaneMandatedCCV() public {
    address laneMandatedCCV = makeAddr("laneMandatedCCV");

    // Configure source chain with lane mandated CCV
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

    CCVAggregator.AggregatedReport memory report = _getReport();
    // Report doesn't include the lane mandated CCV

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, laneMandatedCCV, true));
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_OptionalCCVQuorumNotReached() public {
    CCVAggregator.AggregatedReport memory report = _getReport();
    address receiver = report.message.receiver;

    address optionalCCV2 = makeAddr("optionalCCV2");
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = s_defaultCCV; // This will be found in the report
    optionalCCVs[1] = optionalCCV2; // This won't be found

    uint8 optionalThreshold = 2; // Need 2 optional CCVs

    // Set up receiver to return optional CCVs with threshold 2
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, new address[](0), optionalCCVs, optionalThreshold);
    // Report only includes one CCV, but threshold requires 2

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.OptionalCCVQuorumNotReached.selector, optionalThreshold, 1));
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_InsufficientGasToCompleteTx() public {
    CCVAggregator.AggregatedReport memory report = _getReport();

    // Mock validateReport to pass initial checks
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.validateReport,
        (
          report.message,
          keccak256(abi.encode(report.message)),
          report.ccvData[0],
          Internal.MessageExecutionState.UNTOUCHED
        )
      ),
      abi.encode(true)
    );

    // Mock executeSingleMessage to revert with NOT_ENOUGH_GAS_FOR_CALL_SIG
    vm.mockCallRevert(
      address(s_agg),
      abi.encodeCall(s_agg.executeSingleMessage, (report.message)),
      abi.encodeWithSelector(CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG)
    );

    // Call from gas estimation sender to trigger the specific error handling
    vm.startPrank(Internal.GAS_ESTIMATION_SENDER);
    vm.expectRevert(
      abi.encodeWithSelector(
        CCVAggregator.InsufficientGasToCompleteTx.selector, CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG
      )
    );
    s_agg.execute(report);
    vm.stopPrank();
  }

  function test_execute_RevertWhen_CCVValidationFails() public {
    CCVAggregator.AggregatedReport memory report = _getReport();
    bytes memory revertReason = "CCV validation failed";

    // Mock CCV validateReport to fail/revert
    vm.mockCallRevert(
      s_defaultCCV,
      abi.encodeCall(
        ICCVOffRampV1.validateReport,
        (
          report.message,
          keccak256(abi.encode(report.message)),
          report.ccvData[0],
          Internal.MessageExecutionState.UNTOUCHED
        )
      ),
      revertReason
    );

    vm.expectRevert(revertReason);
    s_agg.execute(report);
  }

  function test_execute_RevertWhen_ExecuteSingleMessageFails() public {
    CCVAggregator.AggregatedReport memory report = _getReport();
    bytes memory revertReason = "ExecuteSingleMessage failed";

    // Mock executeSingleMessage to revert
    vm.mockCallRevert(address(s_agg), abi.encodeCall(s_agg.executeSingleMessage, (report.message)), revertReason);

    vm.expectEmit();
    emit CCVAggregator.ExecutionStateChanged(
      report.message.header.sourceChainSelector,
      report.message.header.sequenceNumber,
      report.message.header.messageId,
      Internal.MessageExecutionState.FAILURE,
      revertReason
    );

    // The message should succeed but set execution state to FAILURE due to executeSingleMessage revert
    // This verifies that execution failures are handled gracefully
    s_agg.execute(report);

    // Verify message state changed to FAILURE
    assertEq(
      uint256(Internal.MessageExecutionState.FAILURE),
      uint256(
        s_agg.getExecutionState(
          SOURCE_CHAIN_SELECTOR, report.message.header.sequenceNumber, report.message.sender, report.message.receiver
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

  function validateReport(
    Internal.Any2EVMMessage memory message,
    bytes32, /* messageHash */
    bytes memory ccvData,
    Internal.MessageExecutionState /* originalState */
  ) external override {
    // Create a dummy report to trigger reentrancy
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(this);
    bytes[] memory ccvDataArray = new bytes[](1);
    ccvDataArray[0] = ccvData;

    // This should trigger the reentrancy guard
    i_aggregator.execute(CCVAggregator.AggregatedReport({message: message, ccvs: ccvs, ccvData: ccvDataArray}));
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId;
  }
}
