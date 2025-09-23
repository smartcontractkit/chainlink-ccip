// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_applyDestChainConfigUpdates is BaseVerifierSetup {
  function test_applyDestChainConfigUpdates() public {
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);

    destChainConfigs[0] = _getDestChainConfig(s_router, 12345, true);

    vm.expectEmit();
    emit BaseVerifier.DestChainConfigSet(
      destChainConfigs[0].destChainSelector, address(destChainConfigs[0].router), destChainConfigs[0].allowlistEnabled
    );

    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);

    // Verify config was set.
    (bool allowlistEnabled, address router,) = s_baseVerifier.getDestChainConfig(destChainConfigs[0].destChainSelector);
    assertEq(router, address(destChainConfigs[0].router));
    assertEq(allowlistEnabled, destChainConfigs[0].allowlistEnabled);
  }

  // Reverts

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroDestChainSelector() public {
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, 0, false);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidDestChainConfig.selector, 0));
    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);
  }
}
