// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_setFeeAggregator is USDCTokenPoolProxySetup {
  address internal s_feeAggregator = makeAddr("feeAggregator");

  function test_setFeeAggregator() public {
    address newFeeAggregator = makeAddr("newFeeAggregator");

    s_usdcTokenPoolProxy.setFeeAggregator(newFeeAggregator);

    assertEq(s_usdcTokenPoolProxy.getFeeAggregator(), newFeeAggregator);
  }

  // Reverts

  function test_setFeeAggregator_RevertWhen_CallerIsNotOwner() public {
    address nonOwner = makeAddr("nonOwner");
    address newFeeAggregator = makeAddr("newFeeAggregator");

    vm.startPrank(nonOwner);
    vm.expectRevert();
    s_usdcTokenPoolProxy.setFeeAggregator(newFeeAggregator);
  }
}

