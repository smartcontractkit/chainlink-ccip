// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeRamp} from "../../../ccvs/CommitteeRamp.sol";
import {CommitteeRampSetup} from "./CommitteeRampSetup.t.sol";

contract CommitteeRamp_constructor is CommitteeRampSetup {
  function test_constructor() public {
    address expectedFeeQuoter = address(s_feeQuoter);
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    // Expect ConfigSet event for the deployment below.
    vm.expectEmit();
    emit CommitteeRamp.ConfigSet(
      _createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    CommitteeRamp newOnRamp =
      new CommitteeRamp(_createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin));

    // Verify dynamic config.
    CommitteeRamp.DynamicConfig memory dynamicConfig = newOnRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, expectedFeeQuoter);
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidConfig_FeeQuoterZeroAddress() public {
    vm.expectRevert(CommitteeRamp.InvalidConfig.selector);
    new CommitteeRamp(
      _createDynamicConfigArgs(address(0), FEE_AGGREGATOR, ALLOWLIST_ADMIN) // Zero fee quoter address
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeAggregatorZeroAddress() public {
    vm.expectRevert(CommitteeRamp.InvalidConfig.selector);
    new CommitteeRamp(
      _createDynamicConfigArgs(address(s_feeQuoter), address(0), ALLOWLIST_ADMIN) // Zero fee aggregator address
    );
  }
}
