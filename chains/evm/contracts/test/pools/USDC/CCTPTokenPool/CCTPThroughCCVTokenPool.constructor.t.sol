// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {CCTPThroughCCVTokenPool} from "../../../../pools/USDC/CCTPThroughCCVTokenPool.sol";
import {CCTPThroughCCVTokenPoolSetup} from "./CCTPThroughCCVTokenPoolSetup.t.sol";

contract CCTPThroughCCVTokenPool_constructor is CCTPThroughCCVTokenPoolSetup {
  function test_constructor_RevertWhen_ZeroAddressCCTPVerifier() public {
    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = s_allowedCaller;

    vm.expectRevert(TokenPool.ZeroAddressInvalid.selector);
    new CCTPThroughCCVTokenPool(s_USDCToken, 6, address(s_mockRMNRemote), address(s_router), address(0), allowedCallers);
  }
}
