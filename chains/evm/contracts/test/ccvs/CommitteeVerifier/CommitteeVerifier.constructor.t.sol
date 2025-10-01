// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_constructor is CommitteeVerifierSetup {
  function test_constructor() public {
    address expectedFeeQuoter = address(s_feeQuoter);
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    // Expect ConfigSet event for the deployment below.
    vm.expectEmit();
    emit CommitteeVerifier.ConfigSet(
      _createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    CommitteeVerifier committeeVerifier = new CommitteeVerifier(
      _createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin), STORAGE_LOCATION
    );

    // Verify dynamic config.
    CommitteeVerifier.DynamicConfig memory dynamicConfig = committeeVerifier.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, expectedFeeQuoter);
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);

    assertEq(committeeVerifier.getStorageLocation(), STORAGE_LOCATION);
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidConfig_FeeQuoterZeroAddress() public {
    vm.expectRevert(CommitteeVerifier.InvalidConfig.selector);
    new CommitteeVerifier(
      _createDynamicConfigArgs(address(0), FEE_AGGREGATOR, ALLOWLIST_ADMIN),
      STORAGE_LOCATION // Zero fee quoter address
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeAggregatorZeroAddress() public {
    vm.expectRevert(CommitteeVerifier.InvalidConfig.selector);
    new CommitteeVerifier(
      _createDynamicConfigArgs(address(s_feeQuoter), address(0), ALLOWLIST_ADMIN),
      STORAGE_LOCATION // Zero fee aggregator address
    );
  }
}
