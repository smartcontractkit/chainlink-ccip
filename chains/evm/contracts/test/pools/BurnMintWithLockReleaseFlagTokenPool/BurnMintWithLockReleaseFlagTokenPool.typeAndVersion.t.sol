// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintWithLockReleaseFlagTokenPoolSetup} from "./BurnMintWithLockReleaseFlagTokenPoolSetup.t.sol";

contract BurnMintWithLockReleaseFlagTokenPool_typeAndVersion is BurnMintWithLockReleaseFlagTokenPoolSetup {
  function test_typeAndVersion() public view {
    assertEq(s_pool.typeAndVersion(), "BurnMintWithLockReleaseFlagTokenPool 2.0.0-dev");
  }
}
