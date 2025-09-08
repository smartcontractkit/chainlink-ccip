// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_constructor is CommitOnRampSetup {
  function test_constructor() public {
    address expectedRmnRemote = address(s_mockRMNRemote);
    address expectedNonceManager = address(s_nonceManager);
    address expectedFeeQuoter = address(s_feeQuoter);
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    CommitOnRamp newOnRamp = new CommitOnRamp(
      expectedRmnRemote,
      expectedNonceManager,
      _getDynamicConfig(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    // Verify static config
    CommitOnRamp.StaticConfig memory staticConfig = newOnRamp.getStaticConfig();
    assertEq(staticConfig.rmnRemote, expectedRmnRemote);
    assertEq(staticConfig.nonceManager, expectedNonceManager);

    // Verify dynamic config
    CommitOnRamp.DynamicConfig memory dynamicConfig = newOnRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, expectedFeeQuoter);
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);

    // Verify ConfigSet event was emitted
    vm.expectEmit();
    emit CommitOnRamp.ConfigSet(
      CommitOnRamp.StaticConfig({rmnRemote: expectedRmnRemote, nonceManager: expectedNonceManager}),
      _getDynamicConfig(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    new CommitOnRamp(
      expectedRmnRemote,
      expectedNonceManager,
      _getDynamicConfig(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );
  }

  function test_constructor_WithValidAllowlistAdmin() public {
    address validAllowlistAdmin = makeAddr("validAllowlistAdmin");

    CommitOnRamp newOnRamp = new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      _getDynamicConfig(address(s_feeQuoter), FEE_AGGREGATOR, validAllowlistAdmin)
    );

    CommitOnRamp.DynamicConfig memory dynamicConfig = newOnRamp.getDynamicConfig();
    assertEq(dynamicConfig.allowlistAdmin, validAllowlistAdmin);
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidConfig_RmnRemoteZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      address(0), // Zero RMN remote address
      address(s_nonceManager),
      _getDynamicConfig(address(s_feeQuoter), FEE_AGGREGATOR, ALLOWLIST_ADMIN)
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_NonceManagerZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      address(s_mockRMNRemote),
      address(0), // Zero nonce manager address
      _getDynamicConfig(address(s_feeQuoter), FEE_AGGREGATOR, ALLOWLIST_ADMIN)
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeQuoterZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      _getDynamicConfig(address(0), FEE_AGGREGATOR, ALLOWLIST_ADMIN) // Zero fee quoter address
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeAggregatorZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      _getDynamicConfig(address(s_feeQuoter), address(0), ALLOWLIST_ADMIN) // Zero fee aggregator address
    );
  }
}
