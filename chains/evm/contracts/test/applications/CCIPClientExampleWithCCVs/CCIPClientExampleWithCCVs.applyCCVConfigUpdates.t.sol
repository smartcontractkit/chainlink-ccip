// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPClientExampleWithCCVs} from "../../../applications/CCIPClientExampleWithCCVs.sol";
import {CCIPClientExampleWithCCVsSetup} from "./CCIPClientExampleWithCCVsSetup.t.sol";

contract CCIPClientExampleWithCCVs_applyCCVConfigUpdates is CCIPClientExampleWithCCVsSetup {
  function test_applyCCVConfigUpdates() public {
    //////////////////////
    // Add a CCV config //
    //////////////////////

    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);
    address[] memory optionalCCVs = new address[](1);
    optionalCCVs[0] = address(0x2);
    uint8 optionalThreshold = 1;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold
    });

    vm.expectEmit();
    emit CCIPClientExampleWithCCVs.CCVConfigSet(SOURCE_CHAIN_SELECTOR, requiredCCVs, optionalCCVs, optionalThreshold);
    s_client.applyCCVConfigUpdates(new uint64[](0), args);

    (address[] memory retRequiredCCVs, address[] memory retOptionalCCVs, uint8 retOptionalThreshold) =
      s_client.getCCVs(SOURCE_CHAIN_SELECTOR);
    assertEq(retRequiredCCVs.length, requiredCCVs.length);
    assertEq(retOptionalCCVs.length, optionalCCVs.length);
    assertEq(retOptionalThreshold, optionalThreshold);
    for (uint256 i = 0; i < requiredCCVs.length; i++) {
      assertEq(retRequiredCCVs[i], requiredCCVs[i]);
    }
    for (uint256 i = 0; i < optionalCCVs.length; i++) {
      assertEq(retOptionalCCVs[i], optionalCCVs[i]);
    }

    /////////////////////////
    // Remove a CCV config //
    /////////////////////////

    uint64[] memory toRemove = new uint64[](1);
    toRemove[0] = SOURCE_CHAIN_SELECTOR;

    vm.expectEmit();
    emit CCIPClientExampleWithCCVs.CCVConfigRemoved(SOURCE_CHAIN_SELECTOR);
    s_client.applyCCVConfigUpdates(toRemove, new CCIPClientExampleWithCCVs.CCVConfigArgs[](0));

    (retRequiredCCVs, retOptionalCCVs, retOptionalThreshold) = s_client.getCCVs(SOURCE_CHAIN_SELECTOR);
    assertEq(retRequiredCCVs.length, 0);
    assertEq(retOptionalCCVs.length, 0);
    assertEq(retOptionalThreshold, 0);
  }

  function test_applyCCVConfigUpdates_RevertWhen_NoCCVsProvided() public {
    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: new address[](0),
      optionalCCVs: new address[](0),
      optionalThreshold: 0
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPClientExampleWithCCVs.NoCCVsProvided.selector, SOURCE_CHAIN_SELECTOR));
    s_client.applyCCVConfigUpdates(new uint64[](0), args);
  }

  function test_applyCCVConfigUpdates_RevertWhen_InvalidOptionalThreshold() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);
    address[] memory optionalCCVs = new address[](1);
    optionalCCVs[0] = address(0x2);
    uint8 optionalThreshold = 2;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        CCIPClientExampleWithCCVs.InvalidOptionalThreshold.selector, SOURCE_CHAIN_SELECTOR, optionalThreshold
      )
    );
    s_client.applyCCVConfigUpdates(new uint64[](0), args);
  }

  function test_applyCCVConfigUpdates_RevertWhen_ZeroAddressNotAllowed() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0);
    address[] memory optionalCCVs = new address[](1);
    optionalCCVs[0] = address(0x2);
    uint8 optionalThreshold = 1;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPClientExampleWithCCVs.ZeroAddressNotAllowed.selector));
    s_client.applyCCVConfigUpdates(new uint64[](0), args);
  }

  function test_applyCCVConfigUpdates_RevertWhen_DuplicateCCV() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = address(0x2);
    optionalCCVs[1] = address(0x1); // duplicate
    uint8 optionalThreshold = 1;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold
    });

    vm.expectRevert(
      abi.encodeWithSelector(CCIPClientExampleWithCCVs.DuplicateCCV.selector, SOURCE_CHAIN_SELECTOR, address(0x1))
    );
    s_client.applyCCVConfigUpdates(new uint64[](0), args);
  }
}
