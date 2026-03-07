// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPClientExampleWithCCVs} from "../../../applications/CCIPClientExampleWithCCVs.sol";
import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";
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
      requireFinality: false
    });

    vm.expectEmit();
    emit CCIPClientExampleWithCCVs.CCVConfigSet(
      SOURCE_CHAIN_SELECTOR, requiredCCVs, optionalCCVs, optionalThreshold, false
    );
    s_client.applyCCVConfigUpdates(args);

    (address[] memory retRequiredCCVs, address[] memory retOptionalCCVs, uint8 retOptionalThreshold) =
      s_client.getCCVs(SOURCE_CHAIN_SELECTOR, 0);
    assertEq(retRequiredCCVs.length, requiredCCVs.length);
    assertEq(retOptionalCCVs.length, optionalCCVs.length);
    assertEq(retOptionalThreshold, optionalThreshold);
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
      requireFinality: false
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
      requireFinality: false
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
      requireFinality: false
    });

    vm.expectRevert(
      abi.encodeWithSelector(CCIPClientExampleWithCCVs.DuplicateCCV.selector, SOURCE_CHAIN_SELECTOR, address(0x1))
    );
    s_client.applyCCVConfigUpdates(args);
  }

  function test_getCCVs_RequireFinality_SucceedsWhenFinalized() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0,
      requireFinality: true
    });

    s_client.applyCCVConfigUpdates(args);

    (address[] memory retRequired,,) = s_client.getCCVs(SOURCE_CHAIN_SELECTOR, 0);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
  }

  function test_getCCVs_RequireFinality_RevertWhen_NotFinalized() public {
    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: new address[](1),
      optionalCCVs: new address[](0),
      optionalThreshold: 0,
      requireFinality: true
    });

    s_client.applyCCVConfigUpdates(args);

    uint16 nonZeroFinality = 5;
    vm.expectRevert(
      abi.encodeWithSelector(CCIPReceiver.BlockDepthNotSupported.selector, SOURCE_CHAIN_SELECTOR, nonZeroFinality)
    );
    s_client.getCCVs(SOURCE_CHAIN_SELECTOR, nonZeroFinality);
  }

  function test_getCCVs_NoRequireFinality_SucceedsWithAnyFinality() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0,
      requireFinality: false
    });

    s_client.applyCCVConfigUpdates(args);

    (address[] memory retRequired,,) = s_client.getCCVs(SOURCE_CHAIN_SELECTOR, 0);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));

    (retRequired,,) = s_client.getCCVs(SOURCE_CHAIN_SELECTOR, 5);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
  }
}
