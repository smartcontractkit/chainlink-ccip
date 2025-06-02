// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {LockReleaseTokenPoolSetup} from "./LockReleaseTokenPoolSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LockReleaseTokenPool_setRebalancer is LockReleaseTokenPoolSetup {
  function test_SetRebalancer() public {
    assertEq(address(s_lockReleaseTokenPool.getRebalancer()), OWNER);

    vm.expectEmit();
    emit LockReleaseTokenPool.RebalancerSet(OWNER, STRANGER);

    s_lockReleaseTokenPool.setRebalancer(STRANGER);
    assertEq(address(s_lockReleaseTokenPool.getRebalancer()), STRANGER);
  }

  function test_SetRebalancer_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lockReleaseTokenPool.setRebalancer(STRANGER);
  }
}
