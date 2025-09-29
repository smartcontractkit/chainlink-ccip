// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_forwardFromRouter is CCVProxySetup {
  function setUp() public virtual override {
    super.setUp();

    vm.startPrank(address(s_sourceRouter));
  }

  function test_forwardFromRouter_oldExtraArgs() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    (
      bytes32 messageId,
      bytes memory encodedMessage,
      CCVProxy.Receipt[] memory verifierReceipts,
      CCVProxy.Receipt memory executorReceipt,
      bytes[] memory receiptBlobs
    ) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: 1,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      messageId: messageId,
      encodedMessage: encodedMessage,
      verifierReceipts: verifierReceipts,
      executorReceipt: executorReceipt,
      receiptBlobs: receiptBlobs
    });

    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_SequenceNumberPersistsAndIncrements() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    // use the stored seq as a running expected value
    CCVProxy.DestChainConfig memory destConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    destConfig.sequenceNumber++;
    // 1) Expect seq to increment for the first message.
    (
      bytes32 messageIdExpected,
      bytes memory encodedMessage,
      CCVProxy.Receipt[] memory verifierReceipts,
      CCVProxy.Receipt memory executorReceipt,
      bytes[] memory receiptBlobs
    ) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: destConfig.sequenceNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: destConfig.sequenceNumber,
      messageId: messageIdExpected,
      encodedMessage: encodedMessage,
      verifierReceipts: verifierReceipts,
      executorReceipt: executorReceipt,
      receiptBlobs: receiptBlobs
    });
    bytes32 messageId1 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // 2) Expect seq to increment again for the next message.
    destConfig.sequenceNumber++;
    (messageIdExpected, encodedMessage, verifierReceipts, executorReceipt, receiptBlobs) = _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: destConfig.sequenceNumber,
      originalSender: STRANGER
    });

    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: destConfig.sequenceNumber,
      messageId: messageIdExpected,
      encodedMessage: encodedMessage,
      verifierReceipts: verifierReceipts,
      executorReceipt: executorReceipt,
      receiptBlobs: receiptBlobs
    });
    bytes32 messageId2 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // Verify sequence numbers and message id are different
    assertTrue(messageId1 != messageId2);
    CCVProxy.DestChainConfig memory finalConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(finalConfig.sequenceNumber, destConfig.sequenceNumber);
  }

  function test_forwardFromRouter_RevertWhen_RouterMustSetOriginalSender() public {
    vm.expectRevert(CCVProxy.RouterMustSetOriginalSender.selector);
    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }
}
