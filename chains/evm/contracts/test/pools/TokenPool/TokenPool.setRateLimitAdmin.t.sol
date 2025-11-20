// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolSetup} from "./TokenPoolSetup.t.sol";

contract TokenPool_setRateLimitAdmin is TokenPoolSetup {
  function test_SetRateLimitAdmin() public {
    address newRateLimitAdmin = makeAddr("newRateLimitAdmin");
    (,,, address currentRateLimitAdmin) = s_tokenPool.getDynamicConfig();
    assertEq(address(0), currentRateLimitAdmin);

    vm.expectEmit();
    emit TokenPool.DynamicConfigSet(address(s_sourceRouter), 0, 0, newRateLimitAdmin);
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), 0, 0, newRateLimitAdmin);

    (,,, address updatedRateLimitAdmin) = s_tokenPool.getDynamicConfig();
    assertEq(newRateLimitAdmin, updatedRateLimitAdmin);
  }

  // Reverts

  function test_RevertWhen_SetRateLimitAdmin() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), 0, 0, STRANGER);
  }
}
