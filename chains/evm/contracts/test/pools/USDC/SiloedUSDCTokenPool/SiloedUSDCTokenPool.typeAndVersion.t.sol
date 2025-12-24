// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_typeAndVersion is SiloedUSDCTokenPoolSetup {
  function test_typeAndVersion() public view {
    assertEq(s_usdcTokenPool.typeAndVersion(), "SiloedUSDCTokenPool 1.7.0-dev");
  }
}
