// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {IFeeQuoter} from "../../../interfaces/IFeeQuoter.sol";
import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {Vm} from "forge-std/Vm.sol";

contract OnRamp_forwardFromRouter is OnRampSetup {
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    keccak256("CCIPMessageSent(uint64,address,bytes32,address,bytes,(address,uint32,uint32,uint256,bytes)[],bytes[])");

  function setUp() public virtual override {
    super.setUp();

    vm.startPrank(address(s_sourceRouter));
    // Router normally forwards the fee token balance before calling the onRamp.
    deal(s_sourceFeeToken, address(s_onRamp), type(uint96).max);
  }

  function test_forwardFromRouter_oldExtraArgs() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    (bytes32 messageId, bytes memory encodedMessage, OnRamp.Receipt[] memory receipts, bytes[] memory verifierBlobs) = _evmMessageToEvent({
      message: message, destChainSelector: DEST_CHAIN_SELECTOR, msgNum: 1, originalSender: STRANGER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sender: STRANGER,
      messageId: messageId,
      feeToken: s_sourceFeeToken,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });

    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_TokenReceiverNotAllowed() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encode(OWNER);

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createV3ExtraArgs(new address[](0), new bytes[](0));
    extraArgs.tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    message.extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.TokenReceiverNotAllowed.selector, DEST_CHAIN_SELECTOR));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e18, STRANGER);
  }

  function test_forwardFromRouter_messageNumberPersistsAndIncrements() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    // Use the stored msgNum as a running expected value.
    OnRamp.DestChainConfig memory destConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    destConfig.messageNumber++;
    // 1) Expect msgNum to increment for the first message.
    (
      bytes32 messageIdExpected,
      bytes memory encodedMessage,
      OnRamp.Receipt[] memory receipts,
      bytes[] memory verifierBlobs
    ) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      msgNum: destConfig.messageNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sender: STRANGER,
      messageId: messageIdExpected,
      feeToken: s_sourceFeeToken,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });
    bytes32 messageId1 = s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);

    // 2) Expect msgNum to increment again for the next message.
    destConfig.messageNumber++;
    (messageIdExpected, encodedMessage, receipts, verifierBlobs) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      msgNum: destConfig.messageNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sender: STRANGER,
      messageId: messageIdExpected,
      feeToken: s_sourceFeeToken,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });
    bytes32 messageId2 = s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);

    // Verify message numbers and message id are different.
    assertTrue(messageId1 != messageId2);
    OnRamp.DestChainConfig memory finalConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(finalConfig.messageNumber, destConfig.messageNumber);
  }

  function test_getExpectedNextMessageNumber_TracksDestChainCounter() public {
    // Before any messages are sent, the next message number should be 1.
    assertEq(s_onRamp.getExpectedNextMessageNumber(DEST_CHAIN_SELECTOR), 1);

    // Send a message through the router.
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);

    // After sending one message, the next expected number should increment.
    assertEq(s_onRamp.getExpectedNextMessageNumber(DEST_CHAIN_SELECTOR), 2);
  }

  function test_forwardFromRouter_UsesMessageNetworkFeeWhenNoTokens() public {
    uint256 feeTokenPrice = 1e18;
    uint256 percentMultiplier = 100;
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(IFeeQuoter.quoteGasForExec.selector),
      abi.encode(uint32(0), uint256(0), feeTokenPrice, percentMultiplier)
    );

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    vm.recordLogs();
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);
    OnRamp.Receipt[] memory receipts = _getReceiptsFromLogs(vm.getRecordedLogs());

    uint256 expectedFee = (uint256(MESSAGE_NETWORK_FEE_USD_CENTS) * percentMultiplier * 1e32) / feeTokenPrice;
    assertEq(receipts[receipts.length - 1].feeTokenAmount, expectedFee);
  }

  function test_forwardFromRouter_UsesTokenNetworkFeeWhenTokens() public {
    uint256 feeTokenPrice = 1e18;
    uint256 percentMultiplier = 100;
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(IFeeQuoter.quoteGasForExec.selector),
      abi.encode(uint32(0), uint256(0), feeTokenPrice, percentMultiplier)
    );

    address token = s_sourceTokens[0];
    address pool = s_sourcePoolByToken[token];
    uint256 amount = 1 ether;
    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: abi.encode(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: STRANGER,
      amount: amount,
      localToken: token
    });
    Pool.LockOrBurnOutV1 memory returnData =
      Pool.LockOrBurnOutV1({destTokenAddress: abi.encode(address(s_destTokenBySourceToken[token])), destPoolData: ""});
    vm.mockCall(pool, abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, expectedInput), abi.encode(returnData));
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(false)
    );
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeCall(IFeeQuoter.getTokenTransferFee, (DEST_CHAIN_SELECTOR, token)),
      abi.encode(uint256(0), uint32(0), uint32(0))
    );

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, amount);
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    vm.recordLogs();
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);
    OnRamp.Receipt[] memory receipts = _getReceiptsFromLogs(vm.getRecordedLogs());

    uint256 expectedFee = (uint256(TOKEN_NETWORK_FEE_USD_CENTS) * percentMultiplier * 1e32) / feeTokenPrice;
    assertEq(receipts[receipts.length - 1].feeTokenAmount, expectedFee);
  }

  function test_forwardFromRouter_RevertWhen_RouterMustSetOriginalSender() public {
    vm.expectRevert(OnRamp.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }

  function test_forwardFromRouter_RevertWhen_DestChainNotSupportedByCCV() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    vm.mockCall(
      address(s_defaultCCV),
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector, DEST_CHAIN_SELECTOR),
      abi.encode(address(0))
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRamp.DestinationChainNotSupportedByCCV.selector, address(s_defaultCCV), DEST_CHAIN_SELECTOR
      )
    );
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_CursedByRMN() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    // Set a curse on the specific destination chain (subject-specific, not global).
    _setMockRMNChainCurse(DEST_CHAIN_SELECTOR, true);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.CursedByRMN.selector, DEST_CHAIN_SELECTOR));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_InsufficientFeeTokenAmount() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InsufficientFeeTokenAmount.selector));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee - 1, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_CanOnlySendOneTokenPerMessage() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    message.tokenAmounts = new Client.EVMTokenAmount[](2);
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: makeAddr("token1"), amount: 123 ether});
    message.tokenAmounts[1] = Client.EVMTokenAmount({token: makeAddr("token2"), amount: 456 ether});

    vm.expectRevert(abi.encodeWithSelector(OnRamp.CanOnlySendOneTokenPerMessage.selector));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_UnsupportedToken() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    message.tokenAmounts = new Client.EVMTokenAmount[](1);
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: makeAddr("unsupportedToken"), amount: 123 ether});

    vm.expectRevert(abi.encodeWithSelector(OnRamp.UnsupportedToken.selector, message.tokenAmounts[0].token));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_RevertWhen_SourceTokenDataTooLarge() public {
    address token = s_sourceTokens[0];
    uint256 amount = 1 ether;
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, amount);

    uint32 paidBytesOverhead = 32;
    uint256 actualBytes = uint256(paidBytesOverhead) + 1;
    address pool = s_sourcePoolByToken[token];

    // Make the pool quote a small destBytesOverhead in getFee (so this is what the sender "paid for").
    vm.mockCall(
      pool,
      abi.encodeWithSelector(IPoolV2.getFee.selector, token, DEST_CHAIN_SELECTOR, amount, s_sourceFeeToken, 0, ""),
      abi.encode(uint256(0), uint32(0), paidBytesOverhead, uint16(0), true)
    );

    // Make lockOrBurn return a larger destPoolData than was quoted.
    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: message.receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: STRANGER,
      amount: amount,
      localToken: token
    });
    Pool.LockOrBurnOutV1 memory returnData = Pool.LockOrBurnOutV1({
      destTokenAddress: abi.encode(address(s_destTokenBySourceToken[token])), destPoolData: new bytes(actualBytes)
    });
    vm.mockCall(
      pool, abi.encodeWithSelector(IPoolV2.lockOrBurn.selector, expectedInput, 0, ""), abi.encode(returnData, amount)
    );

    // Quote fee and attempt send. The send should revert due to oversized pool payload.
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
    vm.expectRevert(
      abi.encodeWithSelector(OnRamp.SourceTokenDataTooLarge.selector, token, actualBytes, paidBytesOverhead)
    );
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, STRANGER);
  }

  function _getReceiptsFromLogs(
    Vm.Log[] memory logs
  ) private pure returns (OnRamp.Receipt[] memory receipts) {
    for (uint256 i = 0; i < logs.length; ++i) {
      if (logs[i].topics.length != 0 && logs[i].topics[0] == CCIP_MESSAGE_SENT_TOPIC) {
        (,, receipts,) = abi.decode(logs[i].data, (address, bytes, OnRamp.Receipt[], bytes[]));
        break;
      }
    }
    return receipts;
  }
}
