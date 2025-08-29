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

    vm.expectEmit(false, false, false, false);
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 0,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: 0,
        feeTokenAmount: 1e17,
        feeValueJuels: 1e17,
        originalSender: OWNER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: new bytes[](0)
    });

    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_forwardFromRouter_SequenceNumberPersistsAndIncrements() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    // First call should use sequence number 1
    vm.expectEmit(false, false, false, false);
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: 1,
        feeTokenAmount: 1e17,
        feeValueJuels: 1e17,
        originalSender: OWNER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: new bytes[](0)
    });

    bytes32 messageId1 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // Second call should use sequence number 2.
    vm.expectEmit(false, false, false, false);
    emit CCVProxy.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 2,
      message: _evmMessageToEvent({
        message: message,
        destChainSelector: DEST_CHAIN_SELECTOR,
        seqNum: 2,
        feeTokenAmount: 1e17,
        feeValueJuels: 1e17,
        originalSender: OWNER,
        metadataHash: s_metadataHash
      }),
      receiptBlobs: new bytes[](0)
    });

    bytes32 messageId2 = s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);

    // Verify sequence numbers are different.
    assertTrue(messageId1 != messageId2);

    // Verify the sequence number in storage has been incremented.
    (uint64 sequenceNumber,) = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(sequenceNumber, 2);
  }

  function test_RevertsWhen_forwardFromRouter_RouterMustSetOriginalSender() public {
    vm.expectRevert(CCVProxy.RouterMustSetOriginalSender.selector);
    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }
}
