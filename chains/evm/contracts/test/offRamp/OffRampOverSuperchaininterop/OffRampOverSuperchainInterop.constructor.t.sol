// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {MockCrossL2Inbox} from "../../helpers/MockCrossL2Inbox.sol";

contract OffRampOverSuperchainInterop_constructor is OffRampOverSuperchainInteropSetup {
  function test_Constructor() public {
    // Verify static config
    OffRampOverSuperchainInterop.StaticConfig memory staticConfig = s_offRamp.getStaticConfig();
    assertEq(staticConfig.chainSelector, DEST_CHAIN_SELECTOR);
    assertEq(staticConfig.gasForCallExactCheck, GAS_FOR_CALL_EXACT_CHECK);
    assertEq(address(staticConfig.crossL2Inbox), address(s_mockCrossL2Inbox));
    assertEq(staticConfig.tokenAdminRegistry, address(s_tokenAdminRegistry));
    assertEq(staticConfig.nonceManager, address(s_nonceManager));

    // Verify dynamic config
    OffRampOverSuperchainInterop.DynamicConfig memory dynamicConfig = s_offRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, makeAddr("mockFeeQuoter"));
    assertEq(dynamicConfig.permissionLessExecutionThresholdSeconds, PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS);
    assertEq(dynamicConfig.messageInterceptor, address(0));

    // Verify allowed transmitters
    address[] memory transmitters = s_offRamp.getAllAllowedTransmittes();
    assertEq(transmitters.length, 2);
    assertEq(transmitters[0], s_transmitter1);
    assertEq(transmitters[1], s_transmitter2);
  }

  function test_Constructor_RevertWhen_ZeroCrossL2Inbox() public {
    vm.expectRevert(OffRampOverSuperchainInterop.ZeroAddressNotAllowed.selector);
    new OffRampOverSuperchainInterop(
      OffRampOverSuperchainInterop.StaticConfig({
        chainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        crossL2Inbox: MockCrossL2Inbox(address(0)),
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        nonceManager: address(s_nonceManager)
      }),
      OffRampOverSuperchainInterop.DynamicConfig({
        feeQuoter: makeAddr("mockFeeQuoter"),
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS,
        messageInterceptor: address(0)
      }),
      new OffRampOverSuperchainInterop.SourceChainConfigArgs[](0),
      new address[](0),
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0)
    );
  }

  function test_Constructor_RevertWhen_ZeroChainSelector() public {
    vm.expectRevert(OffRampOverSuperchainInterop.ZeroChainSelectorNotAllowed.selector);
    new OffRampOverSuperchainInterop(
      OffRampOverSuperchainInterop.StaticConfig({
        chainSelector: 0,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        crossL2Inbox: s_mockCrossL2Inbox,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        nonceManager: address(s_nonceManager)
      }),
      OffRampOverSuperchainInterop.DynamicConfig({
        feeQuoter: makeAddr("mockFeeQuoter"),
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS,
        messageInterceptor: address(0)
      }),
      new OffRampOverSuperchainInterop.SourceChainConfigArgs[](0),
      new address[](0),
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0)
    );
  }
}