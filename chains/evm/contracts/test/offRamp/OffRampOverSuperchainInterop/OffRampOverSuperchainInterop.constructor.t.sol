// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

contract OffRampOverSuperchainInterop_constructor is OffRampOverSuperchainInteropSetup {
  function test_constructor() public {
    // Deploy with valid parameters
    OffRamp.StaticConfig memory staticConfig = OffRamp.StaticConfig({
      chainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry),
      nonceManager: address(s_inboundNonceManager)
    });

    OffRamp.DynamicConfig memory dynamicConfig = _generateDynamicOffRampConfig(address(s_feeQuoter));

    OffRamp.SourceChainConfigArgs[] memory sourceChainConfigs = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      isEnabled: true,
      isRMNVerificationDisabled: false,
      onRamp: abi.encode(ON_RAMP_ADDRESS)
    });

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainIdConfigs =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    chainIdConfigs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: CHAIN_ID_1
    });

    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_1, CHAIN_ID_1);

    OffRampOverSuperchainInterop offRamp = new OffRampOverSuperchainInterop(
      staticConfig, dynamicConfig, sourceChainConfigs, address(s_mockCrossL2Inbox), chainIdConfigs
    );

    // Verify deployment
    assertEq(address(s_mockCrossL2Inbox), offRamp.getCrossL2Inbox());

    // Verify chainId config
    assertEq(CHAIN_ID_1, offRamp.getChainId(SOURCE_CHAIN_SELECTOR_1));

    // Verify static config
    OffRamp.StaticConfig memory retrievedStaticConfig = offRamp.getStaticConfig();
    assertEq(DEST_CHAIN_SELECTOR, retrievedStaticConfig.chainSelector);
    assertEq(GAS_FOR_CALL_EXACT_CHECK, retrievedStaticConfig.gasForCallExactCheck);
    assertEq(address(s_mockRMNRemote), address(retrievedStaticConfig.rmnRemote));
    assertEq(address(s_tokenAdminRegistry), retrievedStaticConfig.tokenAdminRegistry);
    assertEq(address(s_inboundNonceManager), retrievedStaticConfig.nonceManager);

    // Verify dynamic config
    OffRamp.DynamicConfig memory retrievedDynamicConfig = offRamp.getDynamicConfig();
    assertEq(address(s_feeQuoter), retrievedDynamicConfig.feeQuoter);
    assertEq(60 * 60, retrievedDynamicConfig.permissionLessExecutionThresholdSeconds);
    assertEq(address(0), retrievedDynamicConfig.messageInterceptor);
  }
}
