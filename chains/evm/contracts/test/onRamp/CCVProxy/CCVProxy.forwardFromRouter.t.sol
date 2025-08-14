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

  function test_RevertsWhen_forwardFromRouter_RouterMustSetOriginalSender() public {
    vm.expectRevert(CCVProxy.RouterMustSetOriginalSender.selector);
    s_ccvProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }
}
