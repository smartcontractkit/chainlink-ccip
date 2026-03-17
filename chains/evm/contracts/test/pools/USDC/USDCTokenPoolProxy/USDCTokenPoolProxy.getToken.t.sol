// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getToken is USDCTokenPoolProxySetup {
  function test_getToken() public view {
    assertEq(address(s_usdcTokenPoolProxy.getToken()), address(s_USDCToken));
  }
}
