// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";

contract CCTPTokenPool_constructor is CCTPTokenPoolSetup {
  function test_constructor() public {
    new CCTPTokenPool(s_USDCToken, 6, address(0), address(s_rmnProxy), address(s_router), address(s_cctpVerifier));
  }

  function test_constructor_RevertWhen_InvalidCCTPVerifier() public {
    vm.expectRevert(abi.encodeWithSelector(CCTPTokenPool.InvalidCCTPVerifier.selector, address(0)));
    new CCTPTokenPool(s_USDCToken, 6, address(0), address(s_rmnProxy), address(s_router), address(0));
  }
}
