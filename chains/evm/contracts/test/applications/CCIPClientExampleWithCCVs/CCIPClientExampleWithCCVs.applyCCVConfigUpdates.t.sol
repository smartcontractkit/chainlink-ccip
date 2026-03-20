// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPClientExampleWithCCVs} from "../../../applications/CCIPClientExampleWithCCVs.sol";
import {RouterSetup} from "../../Router/RouterSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCIPClientExampleWithCCVs_applyCCVConfigUpdates is RouterSetup {
  CCIPClientExampleWithCCVs internal s_client;

  function setUp() public virtual override {
    super.setUp();

    s_client = new CCIPClientExampleWithCCVs(s_destRouter, IERC20(s_destFeeToken));
  }

  function test_applyCCVConfigUpdates() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = address(0x2);
    optionalCCVs[1] = address(0x3);
    uint8 optionalThreshold = 1;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold,
      allowFasterThanFinality: true
    });

    vm.expectEmit();
    emit CCIPClientExampleWithCCVs.CCVConfigSet(
      SOURCE_CHAIN_SELECTOR, requiredCCVs, optionalCCVs, optionalThreshold, args[0].allowFasterThanFinality
    );
    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (
      address[] memory retRequiredCCVs,
      address[] memory retOptionalCCVs,
      uint8 retOptionalThreshold,
      uint16 minBlockConfirmations
    ) = s_client.getCCVsAndMinBlockConfirmations(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequiredCCVs.length, requiredCCVs.length);
    assertEq(retOptionalCCVs.length, optionalCCVs.length);
    assertEq(retOptionalThreshold, optionalThreshold);
    assertEq(minBlockConfirmations, 1);
    for (uint256 i = 0; i < requiredCCVs.length; ++i) {
      assertEq(retRequiredCCVs[i], requiredCCVs[i]);
    }
    for (uint256 i = 0; i < optionalCCVs.length; ++i) {
      assertEq(retOptionalCCVs[i], optionalCCVs[i]);
    }
  }

  function test_applyCCVConfigUpdates_RevertWhen_InvalidOptionalThreshold() public {
    // Fails when optional threshold > optionalCCVs.length
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
      optionalThreshold: optionalThreshold,
      allowFasterThanFinality: true
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        CCIPClientExampleWithCCVs.InvalidOptionalThreshold.selector, SOURCE_CHAIN_SELECTOR, optionalThreshold
      )
    );
    s_client.applyCCVConfigUpdates(args);

    // Also fails when optionalThreshold == optionalCCVs.length
    optionalThreshold = 1;
    args[0].optionalThreshold = optionalThreshold;
    vm.expectRevert(
      abi.encodeWithSelector(
        CCIPClientExampleWithCCVs.InvalidOptionalThreshold.selector, SOURCE_CHAIN_SELECTOR, optionalThreshold
      )
    );
    s_client.applyCCVConfigUpdates(args);
  }

  function test_applyCCVConfigUpdates_RevertWhen_OptionalThresholdWithNoOptionalCCVs() public {
    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: new address[](1),
      optionalCCVs: new address[](0),
      optionalThreshold: 1,
      allowFasterThanFinality: true
    });

    vm.expectRevert(
      abi.encodeWithSelector(CCIPClientExampleWithCCVs.InvalidOptionalThreshold.selector, SOURCE_CHAIN_SELECTOR, 1)
    );
    s_client.applyCCVConfigUpdates(args);
  }

  function test_applyCCVConfigUpdates_RevertWhen_ZeroAddressNotAllowedAsOptional() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = address(0x0);
    optionalCCVs[0] = address(0x2);
    uint8 optionalThreshold = 1;

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold,
      allowFasterThanFinality: true
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPClientExampleWithCCVs.ZeroAddressNotAllowedAsOptional.selector));
    s_client.applyCCVConfigUpdates(args);
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
      optionalThreshold: optionalThreshold,
      allowFasterThanFinality: true
    });

    vm.expectRevert(
      abi.encodeWithSelector(CCIPClientExampleWithCCVs.DuplicateCCV.selector, SOURCE_CHAIN_SELECTOR, address(0x1))
    );
    s_client.applyCCVConfigUpdates(args);
  }

  function test_getCCVsAndMinBlockConfirmations_RequireFinality_ReturnsZeroMinBlockConfirmations() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0,
      allowFasterThanFinality: false
    });

    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (address[] memory retRequired,,, uint16 minBlockConfirmations) =
      s_client.getCCVsAndMinBlockConfirmations(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
    assertEq(minBlockConfirmations, 0);
  }

  function test_getCCVsAndMinBlockConfirmations_NoRequireFinality_ReturnsNonZeroMinBlockConfirmations() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0,
      allowFasterThanFinality: true
    });

    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (address[] memory retRequired,,, uint16 minBlockConfirmations) =
      s_client.getCCVsAndMinBlockConfirmations(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
    assertEq(minBlockConfirmations, 1);
  }
}
