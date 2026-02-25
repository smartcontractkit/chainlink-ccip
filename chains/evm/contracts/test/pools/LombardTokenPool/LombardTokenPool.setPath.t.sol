// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_setPath is LombardTokenPoolSetup {
  bytes32 internal constant L_CHAIN_ID = bytes32("LCHAIN");
  bytes32 internal constant REMOTE_ADAPTER = bytes32("REMOTE_ADAPTER");

  function test_setPath() public {
    bytes32 expectedAllowedCaller = bytes32(uint256(uint160(s_remotePool)));

    vm.expectEmit();
    emit LombardTokenPool.PathSet(DEST_CHAIN_SELECTOR, L_CHAIN_ID, expectedAllowedCaller, REMOTE_ADAPTER);
    s_pool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, abi.encode(s_remotePool), REMOTE_ADAPTER);

    LombardTokenPool.Path memory path = s_pool.getPath(DEST_CHAIN_SELECTOR);
    assertEq(path.lChainId, L_CHAIN_ID);
    assertEq(path.allowedCaller, expectedAllowedCaller);
    assertEq(path.remoteAdapter, REMOTE_ADAPTER);
  }

  function test_setPath_RevertWhen_ChainNotSupported() public {
    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.ChainNotSupported.selector, 999));
    s_pool.setPath(999, L_CHAIN_ID, abi.encode(s_remotePool), bytes32(0));
  }

  function test_setPath_RevertWhen_ZeroLChainId() public {
    vm.expectRevert(LombardTokenPool.ZeroLombardChainId.selector);
    s_pool.setPath(DEST_CHAIN_SELECTOR, bytes32(0), abi.encode(s_remotePool), bytes32(0));
  }

  function test_setPath_RevertWhen_InvalidRemotePoolForChain() public {
    vm.expectRevert(
      abi.encodeWithSelector(TokenPool.InvalidRemotePoolForChain.selector, DEST_CHAIN_SELECTOR, hex"1234")
    );
    s_pool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, hex"1234", bytes32(0));
  }

  function test_setPath_RevertWhen_InvalidAllowedCaller() public {
    bytes memory remotePoolAddress = hex"1234";
    s_pool.addRemotePool(DEST_CHAIN_SELECTOR, remotePoolAddress);

    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.InvalidAllowedCaller.selector, remotePoolAddress));
    s_pool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, remotePoolAddress, bytes32(0));
  }
}
