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

    bytes[] memory expectedReceiptBlobs = new bytes[](1);
    expectedReceiptBlobs[0] = "";

    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: 1,
        feeTokenAmount: 1e17,
        feeValueJuels: 0,
        originalSender: STRANGER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: expectedReceiptBlobs
    });

    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_SequenceNumberPersistsAndIncrements() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    bytes[] memory expectedReceiptBlobs = new bytes[](1);
    expectedReceiptBlobs[0] = "";

    // use the stored seq as a running expected value
    CCVProxy.DestChainConfig memory destConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    uint64 seqNum = destConfig.sequenceNumber;

    // 1) Expect seq to increment for the first message.
    seqNum += 1;
    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: seqNum,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: seqNum,
        feeTokenAmount: 1e17,
        feeValueJuels: 0,
        originalSender: STRANGER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: expectedReceiptBlobs
    });
    bytes32 messageId1 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // 2) Expect seq to increment again for the next message.
    seqNum += 1;
    vm.expectEmit();
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: seqNum,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: seqNum,
        feeTokenAmount: 1e17,
        feeValueJuels: 0,
        originalSender: STRANGER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: expectedReceiptBlobs
    });
    bytes32 messageId2 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // Verify sequence numbers and message id are different
    assertTrue(messageId1 != messageId2);
    CCVProxy.DestChainConfig memory finalConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(finalConfig.sequenceNumber, seqNum);
  }

  function test_forwardFromRouter_RevertWhen_RouterMustSetOriginalSender() public {
    vm.expectRevert(CCVProxy.RouterMustSetOriginalSender.selector);
    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }
}
