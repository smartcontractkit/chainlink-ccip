// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
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

  /// @notice Verifies that inbound getRequiredCCVs converts the source-denominated amount to local decimals
  /// before passing it to the hooks. The local token is 18 decimals, the remote is 6.
  /// Source amount 1001e6 (1001 tokens) converts to 1001e18 which is above the 1000e18 threshold,
  /// so threshold CCVs must be included.
  /// Without the decimal conversion fix, the raw 1001e6 value would be compared against 1000e18 and
  /// would appear far below the threshold, incorrectly skipping the additional CCVs.
  function test_getRequiredCCVs_Inbound_ConvertsSourceDecimalsToLocal_AboveThreshold() public {
    address ccvBase = makeAddr("inboundCCVBase");
    address ccvThreshold = makeAddr("inboundCCVThreshold");

    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = ccvBase;
    address[] memory thresholdInboundCCVs = new address[](1);
    thresholdInboundCCVs[0] = ccvThreshold;

    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);
    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      thresholdInboundCCVs: thresholdInboundCCVs
    });
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);

    // Remote token has 6 decimals.  1001 tokens = 1001e6 in source denomination.
    // After conversion to local (18 decimals): 1001e6 * 10^12 = 1001e18, which exceeds the 1000e18 threshold.
    uint8 remoteDecimals = 6;
    uint256 sourceAmount = 1001e6;
    bytes memory sourcePoolData = abi.encode(uint256(remoteDecimals));

    // The hook should receive the converted local amount (1001e18), not the raw source amount (1001e6).
    uint256 expectedLocalAmount = 1001e18;
    vm.expectCall(
      address(s_advancedPoolHooks),
      abi.encodeCall(
        s_advancedPoolHooks.getRequiredCCVs,
        (
          address(s_token),
          DEST_CHAIN_SELECTOR,
          expectedLocalAmount,
          WAIT_FOR_FINALITY,
          sourcePoolData,
          IPoolV2.MessageDirection.Inbound
        )
      )
    );

    address[] memory ccvs = s_tokenPool.getRequiredCCVs(
      address(s_token),
      DEST_CHAIN_SELECTOR,
      sourceAmount,
      WAIT_FOR_FINALITY,
      sourcePoolData,
      IPoolV2.MessageDirection.Inbound
    );

    // Should return both base + threshold CCVs because converted amount exceeds threshold.
    assertEq(ccvs.length, 2);
    assertEq(ccvs[0], ccvBase);
    assertEq(ccvs[1], ccvThreshold);
  }

  function test_getRequiredCCVs_Inbound_EmptySourcePoolData_FallsBackToLocalDecimals() public {
    address ccvBase = makeAddr("inboundCCVBase");

    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = ccvBase;

    AdvancedPoolHooks.CCVConfigArg[] memory configArgs = new AdvancedPoolHooks.CCVConfigArg[](1);
    configArgs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      thresholdOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      thresholdInboundCCVs: new address[](0)
    });
    s_advancedPoolHooks.applyCCVConfigUpdates(configArgs);

    uint256 sourceAmount = 500e18;

    // Empty extraData triggers fallback to local decimals, so amount is passed through unchanged.
    vm.expectCall(
      address(s_advancedPoolHooks),
      abi.encodeCall(
        s_advancedPoolHooks.getRequiredCCVs,
        (address(s_token), DEST_CHAIN_SELECTOR, sourceAmount, WAIT_FOR_FINALITY, "", IPoolV2.MessageDirection.Inbound)
      )
    );

    s_tokenPool.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, sourceAmount, WAIT_FOR_FINALITY, "", IPoolV2.MessageDirection.Inbound
    );
  }
}
