// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

contract SiloedLockReleaseTokenPool_getAvailableTokens is SiloedLockReleaseTokenPoolSetup {
  function test_getAvailableTokens_ReturnsLockBoxBalance() public {
    uint256 amount = 1e24;

    // Directly fund the lockbox to simulate available liquidity
    deal(address(s_token), address(s_lockBox), amount);

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), amount);
  }

  function test_getAvailableTokens_RevertWhen_LockBoxNotConfigured() public {
    vm.expectRevert(abi.encodeWithSelector(SiloedLockReleaseTokenPool.LockBoxNotConfigured.selector, type(uint64).max));

    s_siloedLockReleaseTokenPool.getAvailableTokens(type(uint64).max);
  }
}
