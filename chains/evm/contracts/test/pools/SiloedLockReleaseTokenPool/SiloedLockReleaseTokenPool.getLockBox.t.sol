// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

contract SiloedLockReleaseTokenPool_getLockBox is SiloedLockReleaseTokenPoolSetup {
  function test_getLockBox() public view {
    assertEq(address(s_siloedLockReleaseTokenPool.getLockBox(DEST_CHAIN_SELECTOR)), address(s_lockBox));
    assertEq(address(s_siloedLockReleaseTokenPool.getLockBox(SILOED_CHAIN_SELECTOR)), address(s_siloLockBox));
  }

  function test_getLockBox_RevertWhen_LockBoxNotConfigured() public {
    uint64 unconfiguredChainSelector = 999;
    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.LockBoxNotConfigured.selector, unconfiguredChainSelector)
    );
    s_siloedLockReleaseTokenPool.getLockBox(unconfiguredChainSelector);
  }
}

