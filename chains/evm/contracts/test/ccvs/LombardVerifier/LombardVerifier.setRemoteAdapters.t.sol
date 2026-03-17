// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LombardVerifier_setRemoteAdapters is LombardVerifierSetup {
  bytes32 internal constant REMOTE_ADAPTER_A = bytes32("REMOTE_ADAPTER_A");
  bytes32 internal constant REMOTE_ADAPTER_B = bytes32("REMOTE_ADAPTER_B");
  uint64 internal constant OTHER_CHAIN_SELECTOR = 999;
  address internal s_tokenA;
  address internal s_tokenB;

  function setUp() public virtual override {
    super.setUp();
    s_tokenA = address(s_testToken);
    s_tokenB = makeAddr("tokenB");
  }

  function test_setRemoteAdapters() public {
    LombardVerifier.RemoteAdapterArgs[] memory args = new LombardVerifier.RemoteAdapterArgs[](1);
    args[0] = LombardVerifier.RemoteAdapterArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: s_tokenA, remoteAdapter: REMOTE_ADAPTER_A
    });

    vm.expectEmit();
    emit LombardVerifier.RemoteAdapterSet(DEST_CHAIN_SELECTOR, s_tokenA, REMOTE_ADAPTER_A);

    s_lombardVerifier.setRemoteAdapters(args);

    assertEq(s_lombardVerifier.getRemoteAdapter(DEST_CHAIN_SELECTOR, s_tokenA), REMOTE_ADAPTER_A);
  }

  function test_setRemoteAdapters_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    LombardVerifier.RemoteAdapterArgs[] memory args = new LombardVerifier.RemoteAdapterArgs[](1);
    args[0] = LombardVerifier.RemoteAdapterArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: s_tokenA, remoteAdapter: REMOTE_ADAPTER_A
    });

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.setRemoteAdapters(args);
  }
}
