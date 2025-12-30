// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_applyCCVConfigUpdates is AdvancedPoolHooksSetup {
  address internal s_ccv1 = makeAddr("ccv1");
  address internal s_ccv2 = makeAddr("ccv2");

  function test_applyCCVConfigUpdates() public {
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = s_ccv1;

    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = s_ccv2;

    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);
    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      thresholdInboundCCVs: new address[](0)
    });

    vm.expectEmit();
    emit AdvancedPoolHooks.CCVConfigUpdated({
      remoteChainSelector: configArgs[0].remoteChainSelector,
      outboundCCVs: configArgs[0].outboundCCVs,
      thresholdOutboundCCVs: configArgs[0].thresholdOutboundCCVs,
      inboundCCVs: configArgs[0].inboundCCVs,
      thresholdInboundCCVs: configArgs[0].thresholdInboundCCVs
    });
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);

    address[] memory storedOutbound =
      s_tokenPool.getRequiredCCVs(address(0), DEST_CHAIN_SELECTOR, 0, 0, "", IPoolV2.MessageDirection.Outbound);
    address[] memory storedInbound =
      s_tokenPool.getRequiredCCVs(address(0), DEST_CHAIN_SELECTOR, 0, 0, "", IPoolV2.MessageDirection.Inbound);

    assertEq(storedOutbound.length, outboundCCVs.length);
    assertEq(storedOutbound[0], outboundCCVs[0]);

    assertEq(storedInbound.length, inboundCCVs.length);
    assertEq(storedInbound[0], inboundCCVs[0]);
  }

  // Reverts

  function test_applyCCVConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    AdvancedPoolHooks.CCVConfigArg[] memory configArgs;

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCV_Outbound() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);
    address[] memory duplicateOutbound = new address[](3);
    duplicateOutbound[0] = s_ccv1;
    duplicateOutbound[1] = s_ccv2;
    duplicateOutbound[2] = s_ccv1; // Duplicate

    address[] memory validInbound = new address[](1);
    validInbound[0] = s_ccv2;

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: duplicateOutbound,
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: validInbound,
      thresholdInboundCCVs: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv1));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCV_Inbound() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);
    address[] memory validOutbound = new address[](1);
    validOutbound[0] = s_ccv1;

    address[] memory duplicateInbound = new address[](3);
    duplicateInbound[0] = s_ccv2;
    duplicateInbound[1] = s_ccv1;
    duplicateInbound[2] = s_ccv2; // Duplicate

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: validOutbound,
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: duplicateInbound,
      thresholdInboundCCVs: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv2));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_ThresholdOutboundCCVsSpecifiedButNoOutboundBaseCCVs() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory additionalOutbound = new address[](1);
    additionalOutbound[0] = s_ccv1;

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      thresholdOutboundCCVs: additionalOutbound,
      inboundCCVs: new address[](0),
      thresholdInboundCCVs: new address[](0)
    });

    vm.expectRevert(AdvancedPoolHooks.MustSpecifyUnderThresholdCCVsForThresholdCCVs.selector);
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_ThresholdInboundCCVsSpecifiedButNoInboundBaseCCVs() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory additionalInbound = new address[](1);
    additionalInbound[0] = s_ccv2;

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: new address[](0),
      thresholdInboundCCVs: additionalInbound
    });

    vm.expectRevert(AdvancedPoolHooks.MustSpecifyUnderThresholdCCVsForThresholdCCVs.selector);
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCVBetweenOutboundCCVsAndThresholdOutboundCCVs() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory outboundBase = new address[](1);
    outboundBase[0] = s_ccv1;

    address[] memory outboundAdditional = new address[](1);
    outboundAdditional[0] = s_ccv1; // Duplicate across lists

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundBase,
      thresholdOutboundCCVs: outboundAdditional,
      inboundCCVs: new address[](0),
      thresholdInboundCCVs: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv1));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCVBetweenInboundCCVsAndThresholdInboundCCVs() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory inboundBase = new address[](1);
    inboundBase[0] = s_ccv2;

    address[] memory inboundAdditional = new address[](1);
    inboundAdditional[0] = s_ccv2; // Duplicate across lists

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: inboundBase,
      thresholdInboundCCVs: inboundAdditional
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv2));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }
}
