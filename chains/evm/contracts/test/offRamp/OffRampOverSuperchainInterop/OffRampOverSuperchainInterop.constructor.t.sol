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
    assertEq(staticConfig.chainSelector, retrievedStaticConfig.chainSelector);
    assertEq(staticConfig.gasForCallExactCheck, retrievedStaticConfig.gasForCallExactCheck);
    assertEq(address(staticConfig.rmnRemote), address(retrievedStaticConfig.rmnRemote));
    assertEq(staticConfig.tokenAdminRegistry, retrievedStaticConfig.tokenAdminRegistry);
    assertEq(staticConfig.nonceManager, retrievedStaticConfig.nonceManager);
    assertEq("OffRampOverSuperchainInterop 1.6.2-dev", offRamp.typeAndVersion());

    // Verify dynamic config
    OffRamp.DynamicConfig memory retrievedDynamicConfig = offRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, retrievedDynamicConfig.feeQuoter);
    assertEq(
      dynamicConfig.permissionLessExecutionThresholdSeconds,
      retrievedDynamicConfig.permissionLessExecutionThresholdSeconds
    );
    assertEq(dynamicConfig.messageInterceptor, retrievedDynamicConfig.messageInterceptor);
  }

  // Reverts

  function test_constructor_RevertWhen_ZeroCrossL2Inbox() public {
    OffRamp.StaticConfig memory staticConfig = OffRamp.StaticConfig({
      chainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry),
      nonceManager: address(s_inboundNonceManager)
    });

    vm.expectRevert(OffRampOverSuperchainInterop.CrossL2InboxCannotBeZero.selector);
    new OffRampOverSuperchainInterop(
      staticConfig,
      _generateDynamicOffRampConfig(address(s_feeQuoter)),
      new OffRamp.SourceChainConfigArgs[](0),
      address(0), // Invalid crossL2Inbox
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0)
    );
  }
}
