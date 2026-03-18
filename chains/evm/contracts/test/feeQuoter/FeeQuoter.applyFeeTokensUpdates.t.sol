// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FeeQuoter_applyFeeTokensUpdates is FeeQuoterSetup {
  function test_applyFeeTokensUpdates() public {
    address feeToken = s_sourceTokens[1];
    assertEq(s_feeQuoter.getFeeTokens().length, 6);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenRemoved(feeToken);

    address[] memory feeTokensToRemove = new address[](1);
    feeTokensToRemove[0] = feeToken;
    s_feeQuoter.removeFeeTokens(feeTokensToRemove);
    assertEq(s_feeQuoter.getFeeTokens().length, 5);

    assertEq(s_feeQuoter.getTokenPrice(feeToken).value, 0);
    vm.expectRevert(abi.encodeWithSelector(FeeQuoter.TokenNotSupported.selector, feeToken));
    s_feeQuoter.getValidatedTokenPrice(feeToken);

    // removing already removed feeToken is no-op and does not emit an event
    vm.recordLogs();

    s_feeQuoter.removeFeeTokens(feeTokensToRemove);
    assertEq(s_feeQuoter.getFeeTokens().length, 5);

    vm.assertEq(vm.getRecordedLogs().length, 0);
  }

  // Reverts

  function test_applyFeeTokensUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_feeQuoter.removeFeeTokens(new address[](0));
  }
}
