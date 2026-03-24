// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPool_getFee is AdvancedPoolHooksSetup {
  function test_getFee_DefaultFinality() public {
    uint16 defaultFeeBps = 250; // 2.50%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      finalityFeeUSDCents: 75,
      fastFinalityFeeUSDCents: 125,
      finalityTransferFeeBps: defaultFeeBps,
      fastFinalityTransferFeeBps: 400,
      isEnabled: true
    });

    _applyFeeConfig(feeConfig);

    uint256 amount = 1_000e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), bytes2(0), "");

    assertEq(usdFeeCents, feeConfig.finalityFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, defaultFeeBps);
    assertEq(isEnabled, true);
  }

  function test_getFee_CustomFinality() public {
    uint16 customFeeBps = 400; // 4%
    bytes2 minFinality = bytes2(uint16(5));
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      finalityFeeUSDCents: 80,
      fastFinalityFeeUSDCents: 150,
      finalityTransferFeeBps: 0,
      fastFinalityTransferFeeBps: customFeeBps,
      isEnabled: true
    });

    vm.startPrank(OWNER);
    // Enable custom block confirmations by setting minFinality > 0.
    s_tokenPool.setFinalityConfig(minFinality);
    _applyFeeConfig(feeConfig);

    uint256 amount = 1_500e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), minFinality, "");

    assertEq(usdFeeCents, feeConfig.fastFinalityFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, customFeeBps);
    assertEq(isEnabled, true);
  }

  function test_getFee_NoConfig() public view {
    uint256 amount = 777e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), bytes2(0), "");

    assertEq(usdFeeCents, 0);
    assertEq(destGasOverhead, 0);
    assertEq(destBytesOverhead, 0);
    assertEq(tokenFeeBps, 0);
    assertEq(isEnabled, false);
  }

  // Reverts

  function test_getFee_RevertWhen_FastFinalityNotEnabled() public {
    bytes2 requestedFinality = bytes2(uint16(1)); // Any non-zero value triggers custom finality path

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requestedFinality, bytes2(0))
    );
    s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, 1e18, address(0), requestedFinality, "");
  }

  function test_getFee_RevertWhen_InvalidFinalityConfig() public {
    bytes2 minFinality = bytes2(uint16(10));

    // Set custom block confirmation config with minimum of 10 blocks
    s_tokenPool.setFinalityConfig(minFinality);

    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      finalityFeeUSDCents: 75,
      fastFinalityFeeUSDCents: 125,
      finalityTransferFeeBps: 100,
      fastFinalityTransferFeeBps: 200,
      isEnabled: true
    });
    _applyFeeConfig(feeConfig);

    uint256 amount = 1_000e6;
    bytes2 requestedFinality = bytes2(uint16(5)); // Less than minimum of 10

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requestedFinality, minFinality)
    );
    s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), requestedFinality, "");
  }

  function test_getFee_DisabledConfig_ReturnsZeros() public {
    // First enable a config
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      finalityFeeUSDCents: 75,
      fastFinalityFeeUSDCents: 125,
      finalityTransferFeeBps: 250,
      fastFinalityTransferFeeBps: 400,
      isEnabled: true
    });

    _applyFeeConfig(feeConfig);

    // Now disable it
    uint64[] memory disableConfigs = new uint64[](1);
    disableConfigs[0] = DEST_CHAIN_SELECTOR;
    s_tokenPool.applyTokenTransferFeeConfigUpdates(new TokenPool.TokenTransferFeeConfigArgs[](0), disableConfigs);

    uint256 amount = 1_000e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), bytes2(0), "");

    // Should return all zeros with isEnabled=false when disabled
    assertEq(usdFeeCents, 0, "Fee should be zero");
    assertEq(destGasOverhead, 0, "Gas overhead should be zero");
    assertEq(destBytesOverhead, 0, "Bytes overhead should be zero");
    assertEq(tokenFeeBps, 0, "Token fee bps should be zero");
    assertEq(isEnabled, false, "isEnabled should be false");
  }

  function test_getFee_DisabledConfig_CustomFinality_ReturnsZeros() public {
    uint16 minFinality = 5;

    vm.startPrank(OWNER);
    s_tokenPool.setFinalityConfig(bytes2(uint16(minFinality)));

    // First enable a config
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      finalityFeeUSDCents: 80,
      fastFinalityFeeUSDCents: 150,
      finalityTransferFeeBps: 100,
      fastFinalityTransferFeeBps: 400,
      isEnabled: true
    });

    _applyFeeConfig(feeConfig);

    // Now disable it
    uint64[] memory disableConfigs = new uint64[](1);
    disableConfigs[0] = DEST_CHAIN_SELECTOR;
    s_tokenPool.applyTokenTransferFeeConfigUpdates(new TokenPool.TokenTransferFeeConfigArgs[](0), disableConfigs);

    uint256 amount = 1_500e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), bytes2(uint16(minFinality)), "");

    // Should return all zeros with isEnabled=false when disabled, even for custom finality
    assertEq(usdFeeCents, 0, "Fee should be zero");
    assertEq(destGasOverhead, 0, "Gas overhead should be zero");
    assertEq(destBytesOverhead, 0, "Bytes overhead should be zero");
    assertEq(tokenFeeBps, 0, "Token fee bps should be zero");
    assertEq(isEnabled, false, "isEnabled should be false");
  }

  function _applyFeeConfig(
    IPoolV2.TokenTransferFeeConfig memory feeConfig
  ) internal {
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }
}
