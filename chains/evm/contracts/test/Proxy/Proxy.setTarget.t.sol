// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Proxy} from "../../Proxy.sol";
import {ProxySetup} from "./ProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Proxy_setTarget is ProxySetup {
  function test_setTarget() public {
    address newTarget = makeAddr("NewTarget");
    s_proxy.setTarget(newTarget);

    assertEq(s_proxy.getTarget(), newTarget);
  }

  function test_setTarget_RevertWhen_ZeroAddressNotAllowed() public {
    vm.expectRevert(Proxy.ZeroAddressNotAllowed.selector);
    s_proxy.setTarget(address(0));
  }

  function test_setTarget_RevertWhen_NotOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_proxy.setTarget(makeAddr("NewTarget"));
  }
}
