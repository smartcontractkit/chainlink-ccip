// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../../interfaces/ICrossChainVerifierResolver.sol";

import {CCTPVerifier} from "../../../../ccvs/CCTPVerifier.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {CCTPThroughCCVTokenPool} from "../../../../pools/USDC/CCTPThroughCCVTokenPool.sol";
import {CCTPThroughCCVTokenPoolSetup} from "./CCTPThroughCCVTokenPoolSetup.t.sol";

contract CCTPThroughCCVTokenPool_getTokenTransferFeeConfig is CCTPThroughCCVTokenPoolSetup {
  address internal s_mockCCTPVerifier = makeAddr("mockCCTPVerifier");
  uint16 internal constant FAST_FINALITY_BPS = 2; // 0.02%

  function setUp() public virtual override {
    super.setUp();

    // Mock the resolver's getOutboundImplementation call.
    vm.mockCall(
      address(s_cctpThroughCCVTokenPool.getCCTPVerifier()),
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

  function test_getTokenTransferFeeConfig() public {
    // Set up a fee config first.
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationsFeeUSDCents: 100,
      customBlockConfirmationsFeeUSDCents: 150,
      defaultBlockConfirmationsTransferFeeBps: 0,
      customBlockConfirmationsTransferFeeBps: 0,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    changePrank(OWNER);
    s_cctpThroughCCVTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config with default finality.
    IPoolV2.TokenTransferFeeConfig memory returnedFeeConfig =
      s_cctpThroughCCVTokenPool.getTokenTransferFeeConfig(address(s_USDCToken), DEST_CHAIN_SELECTOR, 0, "");

    assertEq(returnedFeeConfig.destGasOverhead, feeConfig.destGasOverhead);
    assertEq(returnedFeeConfig.destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(returnedFeeConfig.defaultBlockConfirmationsFeeUSDCents, feeConfig.defaultBlockConfirmationsFeeUSDCents);
    assertEq(returnedFeeConfig.customBlockConfirmationsFeeUSDCents, feeConfig.customBlockConfirmationsFeeUSDCents);
    assertEq(returnedFeeConfig.defaultBlockConfirmationsTransferFeeBps, 0);
    // Custom block confirmation transfer fee bps should be overridden by the CCTPVerifier's fast finality bps.
    assertEq(returnedFeeConfig.customBlockConfirmationsTransferFeeBps, FAST_FINALITY_BPS);
    assertEq(returnedFeeConfig.isEnabled, feeConfig.isEnabled);
  }

  function test_getTokenTransferFeeConfig_RevertWhen_CCVNotSetOnResolver() public {
    // Mock the resolver's getOutboundImplementation call to return address(0).
    vm.mockCall(
      address(s_cctpThroughCCVTokenPool.getCCTPVerifier()),
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector, DEST_CHAIN_SELECTOR),
      abi.encode(address(0))
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        CCTPThroughCCVTokenPool.CCVNotSetOnResolver.selector, address(s_cctpThroughCCVTokenPool.getCCTPVerifier())
      )
    );
    s_cctpThroughCCVTokenPool.getTokenTransferFeeConfig(address(s_USDCToken), DEST_CHAIN_SELECTOR, 0, "");
  }
}

