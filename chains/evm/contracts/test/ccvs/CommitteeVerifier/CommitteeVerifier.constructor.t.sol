// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_constructor is CommitteeVerifierSetup {
  function test_constructor() public {
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    // Expect ConfigSet event for the deployment below.
    vm.expectEmit();
    emit CommitteeVerifier.ConfigSet(_createDynamicConfigArgs(expectedFeeAggregator, expectedAllowlistAdmin));

    CommitteeVerifier committeeVerifier = new CommitteeVerifier(
      _createDynamicConfigArgs(expectedFeeAggregator, expectedAllowlistAdmin),
      s_storageLocations,
      address(s_mockRMNRemote)
    );

    CommitteeVerifier.DynamicConfig memory dynamicConfig = committeeVerifier.getDynamicConfig();
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);

    string[] memory storageLocations = committeeVerifier.getStorageLocations();
    assertEq(storageLocations.length, 1);
    assertEq(storageLocations[0], storageLocations[0]);
  }
}
