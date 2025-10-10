// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_setThresholdAmountForAdditionalCCVs is TokenPoolV2Setup {
  function test_setThresholdAmountForAdditionalCCVs() public {
    uint256 newThreshold = 1000;

    s_tokenPool.setThresholdAmountForAdditionalCCVs(newThreshold);

    assertEq(s_tokenPool.getThresholdAmountForAdditionalCCVs(), newThreshold);
  }

  // Reverts

  function test_setThresholdAmountForAdditionalCCVs_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setThresholdAmountForAdditionalCCVs(1000);
  }
}
