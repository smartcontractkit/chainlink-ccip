// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPClientExample} from "../../../applications/CCIPClientExample.sol";
import {CCIPClientExampleWithCCVs} from "../../../applications/CCIPClientExampleWithCCVs.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {RouterSetup} from "../../Router/RouterSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCIPClientExampleWithCCVs_applyCCVConfigUpdates is RouterSetup {
  CCIPClientExampleWithCCVs internal s_client;

  bytes internal constant EXTRA_ARGS = abi.encode("extraArgs");

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
      optionalThreshold: optionalThreshold
    });

    vm.expectEmit();
    emit CCIPClientExampleWithCCVs.CCVConfigSet(SOURCE_CHAIN_SELECTOR, requiredCCVs, optionalCCVs, optionalThreshold);
    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (
      address[] memory retRequiredCCVs,
      address[] memory retOptionalCCVs,
      uint8 retOptionalThreshold,
      bytes2 allowedFinalityConfig
    ) = s_client.getCCVsAndFinalityConfig(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequiredCCVs.length, requiredCCVs.length);
    assertEq(retOptionalCCVs.length, optionalCCVs.length);
    assertEq(retOptionalThreshold, optionalThreshold);
    // No finality config set via enableChain, so defaults to bytes2(0) (require full finality).
    assertEq(allowedFinalityConfig, bytes2(0));
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
      optionalThreshold: optionalThreshold
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
      optionalThreshold: 1
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
      optionalThreshold: optionalThreshold
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
      optionalThreshold: optionalThreshold
    });

    vm.expectRevert(
      abi.encodeWithSelector(CCIPClientExampleWithCCVs.DuplicateCCV.selector, SOURCE_CHAIN_SELECTOR, address(0x1))
    );
    s_client.applyCCVConfigUpdates(args);
  }

  function test_getCCVsAndFinalityConfig_RequireFinality_ReturnsZeroAllowedFinalityConfig() public {
    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0
    });

    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (address[] memory retRequired,,, bytes2 allowedFinalityConfig) =
      s_client.getCCVsAndFinalityConfig(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
    // No enableChain call — defaults to bytes2(0) (require full finality).
    assertEq(allowedFinalityConfig, bytes2(0));
  }

  function test_getCCVsAndFinalityConfig_FtfAllowed_ReturnsConfiguredAllowedFinalityConfig() public {
    bytes2 ftfConfig = FinalityCodec._encodeBlockDepthAndSafeFlag(1);

    // Configure finality via the base class enableChain.
    s_client.enableChain(SOURCE_CHAIN_SELECTOR, EXTRA_ARGS, ftfConfig);

    address[] memory requiredCCVs = new address[](1);
    requiredCCVs[0] = address(0x1);

    CCIPClientExampleWithCCVs.CCVConfigArgs[] memory args = new CCIPClientExampleWithCCVs.CCVConfigArgs[](1);
    args[0] = CCIPClientExampleWithCCVs.CCVConfigArgs({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      requiredCCVs: requiredCCVs,
      optionalCCVs: new address[](0),
      optionalThreshold: 0
    });

    s_client.applyCCVConfigUpdates(args);

    bytes memory sender = abi.encodePacked(makeAddr("sender"));

    (address[] memory retRequired,,, bytes2 allowedFinalityConfig) =
      s_client.getCCVsAndFinalityConfig(SOURCE_CHAIN_SELECTOR, sender);
    assertEq(retRequired.length, 1);
    assertEq(retRequired[0], address(0x1));
    assertEq(allowedFinalityConfig, ftfConfig);
  }

  function test_enableChain_AddsToRemoteChainSelectors() public {
    assertEq(s_client.getRemoteChainSelectors().length, 0);

    s_client.enableChain(SOURCE_CHAIN_SELECTOR, EXTRA_ARGS, bytes2(0));

    uint64[] memory selectors = s_client.getRemoteChainSelectors();
    assertEq(selectors.length, 1);
    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
  }

  function test_disableChain_RemovesFromRemoteChainSelectors() public {
    s_client.enableChain(SOURCE_CHAIN_SELECTOR, EXTRA_ARGS, bytes2(0));
    assertEq(s_client.getRemoteChainSelectors().length, 1);

    s_client.disableChain(SOURCE_CHAIN_SELECTOR);
    assertEq(s_client.getRemoteChainSelectors().length, 0);
  }

  function test_enableChain_SetsAllowedFinalityConfig() public {
    bytes2 customFinality = FinalityCodec._encodeBlockDepth(10);
    s_client.enableChain(SOURCE_CHAIN_SELECTOR, EXTRA_ARGS, customFinality);

    bytes2 storedConfig = s_client.getRemoteChainConfig(SOURCE_CHAIN_SELECTOR).allowedFinalityConfig;
    assertEq(storedConfig, customFinality);
  }
}
