// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_setFeeAggregator is VersionedVerifierResolverSetup {
  function test_setFeeAggregator() public {
    address newFeeAggregator = makeAddr("NewFeeAggregator");
    s_versionedVerifierResolver.setFeeAggregator(newFeeAggregator);

    assertEq(s_versionedVerifierResolver.getFeeAggregator(), newFeeAggregator);
  }

  function test_setFeeAggregator_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.setFeeAggregator(makeAddr("NewFeeAggregator"));
  }
}

