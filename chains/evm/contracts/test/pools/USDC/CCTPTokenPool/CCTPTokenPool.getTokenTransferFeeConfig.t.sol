// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../../interfaces/ICrossChainVerifierResolver.sol";

import {CCTPVerifier} from "../../../../ccvs/CCTPVerifier.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";

contract CCTPTokenPool_getTokenTransferFeeConfig is CCTPTokenPoolSetup {
  address internal s_mockCCTPVerifier = makeAddr("mockCCTPVerifier");
  uint16 internal constant FAST_FINALITY_BPS = 2; // 0.02%

  function setUp() public virtual override {
    super.setUp();

    // Mock the resolver's getOutboundImplementation call.
    vm.mockCall(
      address(s_cctpTokenPool.getCCTPVerifier()),
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector, DEST_CHAIN_SELECTOR),
      abi.encode(s_mockCCTPVerifier)
    );

    // Mock the CCTPVerifier's getDynamicConfig call.
    CCTPVerifier.DynamicConfig memory dynamicConfig = CCTPVerifier.DynamicConfig({
      feeAggregator: makeAddr("feeAggregator"),
      allowlistAdmin: makeAddr("allowlistAdmin"),
      fastFinalityBps: FAST_FINALITY_BPS
    });
    vm.mockCall(
      s_mockCCTPVerifier, abi.encodeWithSelector(CCTPVerifier.getDynamicConfig.selector), abi.encode(dynamicConfig)
    );
  }

  function test_getTokenTransferFeeConfig_WithDefaultFinality() public {
    // Set up a fee config first.
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 123,
      customBlockConfirmationTransferFeeBps: 456,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    changePrank(OWNER);
    s_cctpTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config with default finality.
    IPoolV2.TokenTransferFeeConfig memory returnedFeeConfig =
      s_cctpTokenPool.getTokenTransferFeeConfig(address(s_USDCToken), DEST_CHAIN_SELECTOR, 0, "");

    assertEq(returnedFeeConfig.destGasOverhead, feeConfig.destGasOverhead);
    assertEq(returnedFeeConfig.destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(returnedFeeConfig.defaultBlockConfirmationFeeUSDCents, feeConfig.defaultBlockConfirmationFeeUSDCents);
    assertEq(returnedFeeConfig.customBlockConfirmationFeeUSDCents, feeConfig.customBlockConfirmationFeeUSDCents);
    // When default finality, defaultBlockConfirmationTransferFeeBps should be set to 0.
    assertEq(returnedFeeConfig.defaultBlockConfirmationTransferFeeBps, 0);
    assertEq(returnedFeeConfig.customBlockConfirmationTransferFeeBps, feeConfig.customBlockConfirmationTransferFeeBps);
    assertEq(returnedFeeConfig.isEnabled, feeConfig.isEnabled);
  }

  function test_getTokenTransferFeeConfig_WithCustomFinality() public {
    // Set up a fee config first.
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 123,
      customBlockConfirmationTransferFeeBps: 456,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    changePrank(OWNER);
    s_cctpTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config with custom finality (blockConfirmationRequested > 0).
    // This should set customBlockConfirmationTransferFeeBps to dynamicConfig.fastFinalityBps.
    IPoolV2.TokenTransferFeeConfig memory returnedFeeConfig =
      s_cctpTokenPool.getTokenTransferFeeConfig(address(s_USDCToken), DEST_CHAIN_SELECTOR, 1, "");

    assertEq(returnedFeeConfig.destGasOverhead, feeConfig.destGasOverhead);
    assertEq(returnedFeeConfig.destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(returnedFeeConfig.defaultBlockConfirmationFeeUSDCents, feeConfig.defaultBlockConfirmationFeeUSDCents);
    assertEq(returnedFeeConfig.customBlockConfirmationFeeUSDCents, feeConfig.customBlockConfirmationFeeUSDCents);
    assertEq(returnedFeeConfig.defaultBlockConfirmationTransferFeeBps, feeConfig.defaultBlockConfirmationTransferFeeBps);
    // When custom finality, customBlockConfirmationTransferFeeBps should be set to fastFinalityBps.
    assertEq(returnedFeeConfig.customBlockConfirmationTransferFeeBps, FAST_FINALITY_BPS);
    assertEq(returnedFeeConfig.isEnabled, feeConfig.isEnabled);
  }
}
