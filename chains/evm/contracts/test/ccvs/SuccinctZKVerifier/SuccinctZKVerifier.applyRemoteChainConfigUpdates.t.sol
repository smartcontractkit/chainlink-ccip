// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_applyRemoteChainConfigUpdates is SuccinctZKVerifierSetup {
  function test_applyRemoteChainConfigUpdates() public {
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, SOURCE_CHAIN_SELECTOR, true);

    vm.expectEmit();
    emit BaseVerifier.RemoteChainConfigSet(SOURCE_CHAIN_SELECTOR, address(s_router), true);

    s_zkVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    (BaseVerifier.RemoteChainConfigArgs memory config,) = s_zkVerifier.getRemoteChainConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(address(s_router), address(config.router));
    assertTrue(config.allowlistEnabled);
    assertEq(DEFAULT_CCV_FEE_USD_CENTS, config.feeUSDCents);
    assertEq(DEFAULT_CCV_GAS_LIMIT, config.gasForVerification);
    assertEq(DEFAULT_CCV_PAYLOAD_SIZE, config.payloadSizeBytes);
  }

  // Reverts

  function test_applyRemoteChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.applyRemoteChainConfigUpdates(new BaseVerifier.RemoteChainConfigArgs[](0));
  }
}
