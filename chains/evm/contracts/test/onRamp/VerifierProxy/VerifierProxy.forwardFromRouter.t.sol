// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";

import {Client} from "../../../libraries/Client.sol";
import {VerifierProxy} from "../../../onRamp/VerifierProxy.sol";
import {VerifierProxySetup} from "./VerifierProxySetup.t.sol";

contract VerifierProxy_forwardFromRouter is VerifierProxySetup {
  function setUp() public virtual override {
    super.setUp();

    vm.startPrank(address(s_sourceRouter));
  }

  function test_forwardFromRouter_oldExtraArgs() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    vm.expectEmit(false, false, false, false);
    emit VerifierProxy.CCIPMessageSent({
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

    s_verifierProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 1e17, STRANGER);
  }

  function test_RevertsWhen_forwardFromRouter_RouterMustSetOriginalSender() public {
    vm.expectRevert(VerifierProxy.RouterMustSetOriginalSender.selector);
    s_verifierProxy.forwardFromRouter(DEST_CHAIN_SELECTOR, _generateEmptyMessage(), 1e17, address(0));
  }
}
