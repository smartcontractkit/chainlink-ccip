// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_removePath is LombardTokenPoolSetup {
  bytes32 internal constant L_CHAIN_ID = bytes32("LCHAIN");
  bytes32 internal constant REMOTE_ADAPTER = bytes32("REMOTE_ADAPTER");

  function test_removePath_RemovesConfig() public {
    s_pool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, abi.encode(s_remotePool), REMOTE_ADAPTER);

    vm.expectEmit();
    emit LombardTokenPool.PathRemoved(
      DEST_CHAIN_SELECTOR, L_CHAIN_ID, bytes32(uint256(uint160(s_remotePool))), REMOTE_ADAPTER
    );
    s_pool.removePath(DEST_CHAIN_SELECTOR);

    LombardTokenPool.Path memory path = s_pool.getPath(DEST_CHAIN_SELECTOR);
    assertEq(path.allowedCaller, bytes32(0));
    assertEq(path.lChainId, bytes32(0));
    assertEq(path.remoteAdapter, bytes32(0));
  }

  function test_removePath_RevertWhen_PathMissing() public {
    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.PathNotExist.selector, DEST_CHAIN_SELECTOR));
    s_pool.removePath(DEST_CHAIN_SELECTOR);
  }
}
