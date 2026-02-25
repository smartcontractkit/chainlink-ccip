// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPool_getRequiredCCVs is AdvancedPoolHooksSetup {
  uint16 internal constant WAIT_FOR_FINALITY = 0;
  uint16 internal constant CUSTOM_BLOCK_CONFIRMATION = 10;
  uint16 internal constant DEFAULT_FEE_BPS = 100; // 1%
  uint16 internal constant CUSTOM_FEE_BPS = 200; // 2%

  function setUp() public virtual override {
    super.setUp();

    // Set up token transfer fee config for testing
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationsFeeUSDCents: 100, // $1.00
      customBlockConfirmationsFeeUSDCents: 150, // $1.50
      defaultBlockConfirmationsTransferFeeBps: DEFAULT_FEE_BPS,
      customBlockConfirmationsTransferFeeBps: CUSTOM_FEE_BPS,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }

  function test_getRequiredCCVs_WithDefaultBlockConfirmations_AppliesDefaultFeeBps() public {
    uint256 amount = 10_000e18;
    uint256 expectedAmountAfterFee = amount - (amount * DEFAULT_FEE_BPS) / BPS_DIVIDER; // 9,900e18 (1% fee deducted)

    // Expect the hook to be called with the fee-adjusted amount
    vm.expectCall(
      address(s_advancedPoolHooks),
      abi.encodeCall(
        s_advancedPoolHooks.getRequiredCCVs,
        (
          address(s_token),
          DEST_CHAIN_SELECTOR,
          expectedAmountAfterFee,
          WAIT_FOR_FINALITY,
          "",
          IPoolV2.MessageDirection.Outbound
        )
      )
    );

    s_tokenPool.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, amount, WAIT_FOR_FINALITY, "", IPoolV2.MessageDirection.Outbound
    );
  }

  function test_getRequiredCCVs_WithCustomBlockConfirmations_AppliesCustomFeeBps() public {
    uint256 amount = 10_000e18;
    uint256 expectedAmountAfterFee = amount - (amount * CUSTOM_FEE_BPS) / BPS_DIVIDER; // 9,800e18 (2% fee deducted)

    // Expect the hook to be called with the fee-adjusted amount
    vm.expectCall(
      address(s_advancedPoolHooks),
      abi.encodeCall(
        s_advancedPoolHooks.getRequiredCCVs,
        (
          address(s_token),
          DEST_CHAIN_SELECTOR,
          expectedAmountAfterFee,
          CUSTOM_BLOCK_CONFIRMATION,
          "",
          IPoolV2.MessageDirection.Outbound
        )
      )
    );

    s_tokenPool.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, amount, CUSTOM_BLOCK_CONFIRMATION, "", IPoolV2.MessageDirection.Outbound
    );
  }
}
