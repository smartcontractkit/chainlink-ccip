// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FeeQuoter_applyFeeTokensUpdates is FeeQuoterSetup {
  function test_applyFeeTokensUpdates_singleToken() public {
    address[] memory feeTokensToAdd = new address[](1);
    feeTokensToAdd[0] = vm.addr(1);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(vm.addr(1));

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokensToAdd);

    assertEq(s_feeQuoter.getFeeTokens().length, 4);
    assertEq(s_feeQuoter.getFeeTokens()[3], vm.addr(1));
  }

  function test_applyFeeTokensUpdates_multipleTokens() public {
    address[] memory feeTokensToAdd = new address[](2);
    feeTokensToAdd[0] = vm.addr(1);
    feeTokensToAdd[1] = vm.addr(2);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(vm.addr(1));
    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(vm.addr(2));

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokensToAdd);

    assertEq(s_feeQuoter.getFeeTokens().length, 5);
    assertEq(s_feeQuoter.getFeeTokens()[3], vm.addr(1));
    assertEq(s_feeQuoter.getFeeTokens()[4], vm.addr(2));
  }

  function test_applyFeeTokensUpdates() public {
    address[] memory feeTokensToAdd = new address[](1);
    feeTokensToAdd[0] = s_sourceTokens[1];

    // s_sourceTokens[1] is already in s_sourceFeeTokens, so adding it is a no-op

    // add same feeToken is no-op
    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokensToAdd);
    assertEq(s_feeQuoter.getFeeTokens().length, 3);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenRemoved(feeTokensToAdd[0]);

    s_feeQuoter.applyFeeTokensUpdates(feeTokensToAdd, new address[](0));
    assertEq(s_feeQuoter.getFeeTokens().length, 2);

    // removing already removed feeToken is no-op and does not emit an event
    vm.recordLogs();

    s_feeQuoter.applyFeeTokensUpdates(feeTokensToAdd, new address[](0));
    assertEq(s_feeQuoter.getFeeTokens().length, 2);

    vm.assertEq(vm.getRecordedLogs().length, 0);

    // Removing and adding the same fee token is allowed and emits both events
    // Add it first
    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(feeTokensToAdd[0]);
    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokensToAdd);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenRemoved(feeTokensToAdd[0]);
    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(feeTokensToAdd[0]);

    s_feeQuoter.applyFeeTokensUpdates(feeTokensToAdd, feeTokensToAdd);
  }

  // Reverts

  function test_applyFeeTokensUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), new address[](0));
  }
}
