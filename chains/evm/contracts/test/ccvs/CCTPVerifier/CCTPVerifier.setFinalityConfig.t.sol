// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_setFinalityConfig is CCTPVerifierSetup {
  function test_setFinalityConfig() public {
    uint16[] memory customCCIPFinalities = new uint16[](1);
    customCCIPFinalities[0] = CCIP_FAST_FINALITY_THRESHOLD;

    uint32[] memory customCCTPFinalityThresholds = new uint32[](1);
    customCCTPFinalityThresholds[0] = CCTP_FAST_FINALITY_THRESHOLD;

    uint16[] memory customCCTPFinalityBps = new uint16[](1);
    customCCTPFinalityBps[0] = CCTP_FAST_FINALITY_BPS;

    CCTPVerifier.FinalityConfig memory finalityConfig = CCTPVerifier.FinalityConfig({
      defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
      defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
      customCCIPFinalities: customCCIPFinalities,
      customCCTPFinalityThresholds: customCCTPFinalityThresholds,
      customCCTPFinalityBps: customCCTPFinalityBps
    });

    vm.expectEmit();
    emit CCTPVerifier.FinalityConfigSet(finalityConfig);
    s_cctpVerifier.setFinalityConfig(finalityConfig);

    // Check the finality config.
    CCTPVerifier.FinalityConfig memory got = s_cctpVerifier.getFinalityConfig();
    assertEq(got.defaultCCTPFinalityThreshold, CCTP_STANDARD_FINALITY_THRESHOLD);
    assertEq(got.defaultCCTPFinalityBps, CCTP_STANDARD_FINALITY_BPS);
    assertEq(got.customCCIPFinalities.length, 1);
    assertEq(got.customCCIPFinalities[0], CCIP_FAST_FINALITY_THRESHOLD);
    assertEq(got.customCCTPFinalityThresholds.length, 1);
    assertEq(got.customCCTPFinalityThresholds[0], CCTP_FAST_FINALITY_THRESHOLD);
    assertEq(got.customCCTPFinalityBps.length, 1);
    assertEq(got.customCCTPFinalityBps[0], CCTP_FAST_FINALITY_BPS);
  }

  function test_setFinalityConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: new uint16[](0),
        customCCTPFinalityThresholds: new uint32[](0),
        customCCTPFinalityBps: new uint16[](0)
      })
    );
  }

  function test_setFinalityConfig_RevertWhen_MissingCustomFinalities() public {
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.MissingCustomFinalities.selector));
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: new uint16[](0),
        customCCTPFinalityThresholds: new uint32[](0),
        customCCTPFinalityBps: new uint16[](0)
      })
    );
  }

  function test_setFinalityConfig_RevertWhen_CustomFinalityArraysMustBeSameLength_ThresholdsIncorrect() public {
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.CustomFinalityArraysMustBeSameLength.selector));
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: new uint16[](1),
        customCCTPFinalityThresholds: new uint32[](2),
        customCCTPFinalityBps: new uint16[](1)
      })
    );
  }

  function test_setFinalityConfig_RevertWhen_CustomFinalityArraysMustBeSameLength_BpsIncorrect() public {
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.CustomFinalityArraysMustBeSameLength.selector));
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: new uint16[](1),
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](2)
      })
    );
  }

  function test_setFinalityConfig_RevertWhen_CustomFinalitiesMustBeStrictlyIncreasing() public {
    uint16[] memory customCCIPFinalities = new uint16[](2);
    customCCIPFinalities[0] = CCIP_FAST_FINALITY_THRESHOLD + 1;
    customCCIPFinalities[1] = CCIP_FAST_FINALITY_THRESHOLD;

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.CustomFinalitiesMustBeStrictlyIncreasing.selector));
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](2),
        customCCTPFinalityBps: new uint16[](2)
      })
    );
  }

  function test_setFinalityConfig_RevertWhen_CustomFinalitiesMustBeStrictlyIncreasing_CustomFinalityCannotBeZero()
    public
  {
    uint16[] memory customCCIPFinalities = new uint16[](1);
    customCCIPFinalities[0] = 0;

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.CustomFinalitiesMustBeStrictlyIncreasing.selector));
    s_cctpVerifier.setFinalityConfig(
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      })
    );
  }
}
