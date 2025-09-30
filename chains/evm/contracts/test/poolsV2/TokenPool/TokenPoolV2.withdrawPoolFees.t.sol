// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_withdrawPoolFees is TokenPoolV2Setup {
  function test_withdrawPoolFees() public {
    uint256 feeAmount = 20e18;
    s_token.transfer(address(s_tokenPool), feeAmount);
    s_tokenPool.withdrawPoolFees(STRANGER);
    assertEq(s_token.balanceOf(STRANGER), feeAmount);
  }

  // Reverts

  function test_withdrawPoolFees_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.withdrawPoolFees(STRANGER);
  }
}
