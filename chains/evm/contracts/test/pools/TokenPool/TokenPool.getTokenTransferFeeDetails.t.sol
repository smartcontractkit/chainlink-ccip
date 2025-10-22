// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getTokenTransferFeeDetails is TokenPoolV2Setup {
  function test_getTokenTransferFeeDetails_DefaultFinality() public {
    uint256 amount = 1000e6;
    uint16 defaultFeeBps = 250; // 2.50%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultFinalityFeeUSDCents: 75,
      customFinalityFeeUSDCents: 125,
      defaultFinalityTransferFeeBps: defaultFeeBps,
      customFinalityTransferFeeBps: 400,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    vm.startPrank(OWNER);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    IPoolV2.TokenTransferFeeDetails memory feeDetails =
      s_tokenPool.getTokenTransferFeeDetails(address(s_token), DEST_CHAIN_SELECTOR, amount, 0, "");

    assertEq(feeDetails.tokenFeeBps, defaultFeeBps);
    uint256 expectedFeeAmount = (amount * defaultFeeBps) / BPS_DIVIDER;
    assertEq(feeDetails.tokenFeeAmount, expectedFeeAmount);
    assertEq(feeDetails.destinationAmount, amount - expectedFeeAmount);
    assertEq(feeDetails.usdFeeCents, feeConfig.defaultFinalityFeeUSDCents);
  }

  function test_getTokenTransferFeeDetails_CustomFinality() public {
    uint256 amount = 1_500e6;
    uint16 customFeeBps = 400; // 4%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultFinalityFeeUSDCents: 80,
      customFinalityFeeUSDCents: 150,
      defaultFinalityTransferFeeBps: 0,
      customFinalityTransferFeeBps: customFeeBps,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    vm.startPrank(OWNER);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    uint16 finality = 5;
    IPoolV2.TokenTransferFeeDetails memory feeDetails =
      s_tokenPool.getTokenTransferFeeDetails(address(s_token), DEST_CHAIN_SELECTOR, amount, finality, "");

    assertEq(feeDetails.tokenFeeBps, customFeeBps);
    uint256 expectedFeeAmount = (amount * customFeeBps) / BPS_DIVIDER;
    assertEq(feeDetails.tokenFeeAmount, expectedFeeAmount);
    assertEq(feeDetails.destinationAmount, amount - expectedFeeAmount);
    assertEq(feeDetails.usdFeeCents, feeConfig.customFinalityFeeUSDCents);
  }

  function test_getTokenTransferFeeDetails_DisabledConfig() public view {
    uint256 amount = 777e6;
    IPoolV2.TokenTransferFeeDetails memory feeDetails =
      s_tokenPool.getTokenTransferFeeDetails(address(s_token), DEST_CHAIN_SELECTOR, amount, 0, "");

    assertEq(feeDetails.tokenFeeBps, 0);
    assertEq(feeDetails.tokenFeeAmount, 0);
    assertEq(feeDetails.destinationAmount, amount);
    assertEq(feeDetails.usdFeeCents, 0);
  }
}
