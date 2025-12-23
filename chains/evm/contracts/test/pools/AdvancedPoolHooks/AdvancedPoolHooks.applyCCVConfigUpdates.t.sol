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
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: inboundCCVs,
      inboundCCVsToAddAboveThreshold: new address[](0)
    });

    vm.expectEmit();
    emit AdvancedPoolHooks.CCVConfigUpdated({
      remoteChainSelector: configArgs[0].remoteChainSelector,
      outboundCCVs: configArgs[0].outboundCCVs,
      outboundCCVsToAddAboveThreshold: configArgs[0].outboundCCVsToAddAboveThreshold,
      inboundCCVs: configArgs[0].inboundCCVs,
      inboundCCVsToAddAboveThreshold: configArgs[0].inboundCCVsToAddAboveThreshold
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
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: validInbound,
      inboundCCVsToAddAboveThreshold: new address[](0)
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
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: duplicateInbound,
      inboundCCVsToAddAboveThreshold: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv2));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_OutboundCCVsToAddAboveThresholdSpecifiedButNoOutboundBaseCCVs()
    public
  {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory additionalOutbound = new address[](1);
    additionalOutbound[0] = s_ccv1;

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      outboundCCVsToAddAboveThreshold: additionalOutbound,
      inboundCCVs: new address[](0),
      inboundCCVsToAddAboveThreshold: new address[](0)
    });

    vm.expectRevert(AdvancedPoolHooks.MustSpecifyUnderThresholdCCVsForAboveThresholdCCVs.selector);
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_InboundCCVsToAddAboveThresholdSpecifiedButNoInboundBaseCCVs() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory additionalInbound = new address[](1);
    additionalInbound[0] = s_ccv2;

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: new address[](0),
      inboundCCVsToAddAboveThreshold: additionalInbound
    });

    vm.expectRevert(AdvancedPoolHooks.MustSpecifyUnderThresholdCCVsForAboveThresholdCCVs.selector);
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCVBetweenOutboundLists() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory outboundBase = new address[](1);
    outboundBase[0] = s_ccv1;

    address[] memory outboundAdditional = new address[](1);
    outboundAdditional[0] = s_ccv1; // Duplicate across lists

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundBase,
      outboundCCVsToAddAboveThreshold: outboundAdditional,
      inboundCCVs: new address[](0),
      inboundCCVsToAddAboveThreshold: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv1));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCVBetweenInboundLists() public {
    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);

    address[] memory inboundBase = new address[](1);
    inboundBase[0] = s_ccv2;

    address[] memory inboundAdditional = new address[](1);
    inboundAdditional[0] = s_ccv2; // Duplicate across lists

    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: inboundBase,
      inboundCCVsToAddAboveThreshold: inboundAdditional
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv2));
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);
  }
}
