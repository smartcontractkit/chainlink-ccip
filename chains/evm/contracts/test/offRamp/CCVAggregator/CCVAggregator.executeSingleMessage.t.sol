// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// forge-config: default.allow_internal_expect_revert = true

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract CCVAggregator_executeSingleMessage is CCVAggregatorSetup {
  function setUp() public virtual override {
    super.setUp();

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Mock validateReport for default message structure.
    bytes32 messageHash = keccak256(MessageV1Codec._encodeMessageV1(message));

    bytes memory defaultCcvData = abi.encode("mock ccv data");
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (address(s_agg), message, messageHash, defaultCcvData)),
      abi.encode(true)
    );

    vm.startPrank(address(s_agg));
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

  function test_executeSingleMessage_RevertWhen_CanOnlySelfCall() public {
    vm.stopPrank();
    MessageV1Codec.MessageV1 memory message;

    vm.expectRevert(CCVAggregator.CanOnlySelfCall.selector);
    s_agg.executeSingleMessage(message, bytes32(0), new address[](0), new bytes[](0));
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_ReceiverCCV() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address receiver = makeAddr("receiver");
    address requiredCCV = makeAddr("requiredCCV");

    message.receiver = abi.encodePacked(receiver);

    // Set up receiver to require a specific CCV.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, _arrayOf(requiredCCV), new address[](0), 0);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Keep default CCV, but don't include the required CCV.
    bytes[] memory ccvData = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, requiredCCV));

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_PoolCCV() public {
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

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Keep default CCV, but don't include the pool required CCV.
    bytes[] memory ccvData = new bytes[](1);

    // Mock token admin registry to return the pool.
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (token)), abi.encode(pool));

    // Mock pool supportsInterface for IPoolV2.
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool to require a specific CCV.
    vm.mockCall(
      pool,
      abi.encodeCall(IPoolV2.getRequiredInboundCCVs, (token, SOURCE_CHAIN_SELECTOR, tokenAmount, 0, "")),
      abi.encode(_arrayOf(poolRequiredCCV))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, poolRequiredCCV));

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_LaneMandatedCCV() public {
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

    vm.stopPrank();
    vm.startPrank(OWNER);
    s_agg.applySourceChainConfigUpdates(sourceChainConfigArgs);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Report doesn't include the lane mandated CCV.
    bytes[] memory ccvData = new bytes[](1);

    vm.startPrank(address(s_agg));

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, laneMandatedCCV));

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_OptionalCCVQuorumNotReached() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    address receiver = address(bytes20(message.receiver));

    address optionalCCV2 = makeAddr("optionalCCV2");
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = s_defaultCCV; // This will be found in the report.
    optionalCCVs[1] = optionalCCV2; // This won't be found.

    uint8 optionalThreshold = 2; // Need 2 optional CCVs.

    // Set up receiver to return optional CCVs with threshold 2.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, new address[](0), optionalCCVs, optionalThreshold);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Report only includes one CCV, but threshold requires 2.
    bytes[] memory ccvData = new bytes[](1);

    vm.expectRevert(
      abi.encodeWithSelector(CCVAggregator.OptionalCCVQuorumNotReached.selector, optionalThreshold, ccvs.length)
    );

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_CCVValidationFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory ccvData = new bytes[](1);
    bytes memory revertReason = "CCV validation failed";

    // Mock CCV validateReport to fail/revert.
    vm.mockCallRevert(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (address(s_agg), message, messageId, ccvData[0])),
      revertReason
    );

    vm.expectRevert(revertReason);

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_InvalidNumberOfTokens() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Create message with multiple token transfers (invalid) - but we need to manually set this
    // since encoding would fail. We'll create a valid single token transfer first, then modify the array.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool1")),
      sourceTokenAddress: abi.encodePacked(makeAddr("sourceToken1")),
      destTokenAddress: abi.encodePacked(makeAddr("token1")),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    // Encode the message with single token transfer first.
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    // Now manually create the invalid message with 2 token transfers for the executeSingleMessage call.
    MessageV1Codec.TokenTransferV1[] memory invalidTokenAmounts = new MessageV1Codec.TokenTransferV1[](2);
    invalidTokenAmounts[0] = tokenAmounts[0];
    invalidTokenAmounts[1] = tokenAmounts[0];
    message.tokenTransfer = invalidTokenAmounts;

    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory ccvData = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.InvalidNumberOfTokens.selector, invalidTokenAmounts.length));

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }

  function test_executeSingleMessage_RevertWhen_InvalidEVMAddress_TokenAddress() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Create message with invalid token address length.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("sourceToken")),
      destTokenAddress: hex"1234", // Invalid length (not 20 bytes).
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory ccvData = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, tokenAmounts[0].destTokenAddress));

    s_agg.executeSingleMessage(message, messageId, ccvs, ccvData);
  }
}
