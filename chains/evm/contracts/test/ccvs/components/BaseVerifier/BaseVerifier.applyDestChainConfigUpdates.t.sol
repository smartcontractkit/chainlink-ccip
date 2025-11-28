// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_applyRemoteChainConfigUpdates is BaseVerifierSetup {
  function test_applyRemoteChainConfigUpdates() public {
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);

    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, 12345, true);

    vm.expectEmit();
    emit BaseVerifier.RemoteChainConfigSet(
      remoteChainConfigs[0].remoteChainSelector,
      address(remoteChainConfigs[0].router),
      remoteChainConfigs[0].allowlistEnabled
    );

    s_baseVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    // Verify config was set.
    (bool allowlistEnabled, address router,) =
      s_baseVerifier.getRemoteChainConfig(remoteChainConfigs[0].remoteChainSelector);
    assertEq(router, address(remoteChainConfigs[0].router));
    assertEq(allowlistEnabled, remoteChainConfigs[0].allowlistEnabled);
  }

  // Reverts

  function test_applyRemoteChainConfigUpdates_RevertWhen_InvalidRemoteChainConfig_ZeroRemoteChainSelector() public {
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, 0, false);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidRemoteChainConfig.selector, 0));
    s_baseVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);
  }
}
