// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Proxy} from "../../Proxy.sol";
import {ProxySetup} from "./ProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Proxy_setFeeAggregator is ProxySetup {
  function test_setFeeAggregator() public {
    address newFeeAggregator = makeAddr("NewFeeAggregator");
    s_proxy.setFeeAggregator(newFeeAggregator);

    assertEq(s_proxy.getFeeAggregator(), newFeeAggregator);
  }

  function test_setFeeAggregator_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_proxy.setFeeAggregator(makeAddr("NewFeeAggregator"));
  }
}

