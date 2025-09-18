// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitRamp} from "../../../ccvs/CommitRamp.sol";
import {CommitRampSetup} from "./CommitRampSetup.t.sol";

contract CommitRamp_constructor is CommitRampSetup {
  function test_constructor() public {
    address expectedFeeQuoter = address(s_feeQuoter);
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    // Expect ConfigSet event for the deployment below.
    vm.expectEmit();
    emit CommitRamp.ConfigSet(
      _createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    CommitRamp newOnRamp =
      new CommitRamp(_createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin));

    // Verify dynamic config.
    CommitRamp.DynamicConfig memory dynamicConfig = newOnRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, expectedFeeQuoter);
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidConfig_FeeQuoterZeroAddress() public {
    vm.expectRevert(CommitRamp.InvalidConfig.selector);
    new CommitRamp(
      _createDynamicConfigArgs(address(0), FEE_AGGREGATOR, ALLOWLIST_ADMIN) // Zero fee quoter address
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeAggregatorZeroAddress() public {
    vm.expectRevert(CommitRamp.InvalidConfig.selector);
    new CommitRamp(
      _createDynamicConfigArgs(address(s_feeQuoter), address(0), ALLOWLIST_ADMIN) // Zero fee aggregator address
    );
  }
}
