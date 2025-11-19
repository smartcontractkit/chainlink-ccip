// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_forwardFromRouter is OnRampSetup {
  function setUp() public virtual override {
    super.setUp();

    vm.startPrank(address(s_sourceRouter));
    // Router normally forwards the fee token balance before calling the onRamp.
    deal(s_sourceFeeToken, address(s_onRamp), type(uint96).max);
  }

  function test_forwardFromRouter_oldExtraArgs() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    (bytes32 messageId, bytes memory encodedMessage, OnRamp.Receipt[] memory receipts, bytes[] memory verifierBlobs) =
    _evmMessageToEvent({message: message, destChainSelector: DEST_CHAIN_SELECTOR, seqNum: 1, originalSender: STRANGER});

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      messageId: messageId,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });

    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_SequenceNumberPersistsAndIncrements() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    // Use the stored seq as a running expected value.
    OnRamp.DestChainConfig memory destConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    destConfig.sequenceNumber++;
    // 1) Expect seq to increment for the first message.
    (
      bytes32 messageIdExpected,
      bytes memory encodedMessage,
      OnRamp.Receipt[] memory receipts,
      bytes[] memory verifierBlobs
    ) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: destConfig.sequenceNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: destConfig.sequenceNumber,
      messageId: messageIdExpected,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });
    bytes32 messageId1 = s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // 2) Expect seq to increment again for the next message.
    destConfig.sequenceNumber++;
    (messageIdExpected, encodedMessage, receipts, verifierBlobs) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: destConfig.sequenceNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: destConfig.sequenceNumber,
      messageId: messageIdExpected,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });
    bytes32 messageId2 = s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // Verify sequence numbers and message id are different.
    assertTrue(messageId1 != messageId2);
    OnRamp.DestChainConfig memory finalConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(finalConfig.sequenceNumber, destConfig.sequenceNumber);
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
    uint256 totalFee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
    deal(s_sourceFeeToken, address(s_onRamp), totalFee - 1);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InsufficientFeeTokenAmount.selector, totalFee - 1, totalFee));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, totalFee - 1, STRANGER);
  }
}
