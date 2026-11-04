// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

/// @notice Checks store arguments, returned lane data, fee deduction, and caller authorization. Store tests cover
/// the rules within each store.
contract PoolConfigManager_validateLockOrBurn is PoolConfigManagerTransferSetup {
  uint256 internal constant AMOUNT = 1e18;

  /// @notice Token A uses its inline remote token address and pool data. The default fee record charges no fees,
  /// and the security record has no hook.
  function test_validateLockOrBurn() public {
    _setCaller(address(s_pool));

    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );

    assertEq(abi.encode(s_remoteTokenA), lane.remoteToken);
    assertEq(_getPoolData(s_tokenA), lane.destPoolData);
    assertEq(DEFAULT_TOKEN_DECIMALS, lane.localDecimals);
    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.Burn), uint8(lane.lockOrBurnFlavor));
    assertEq(0, lane.feeAmount);
    assertEq(address(0), lane.hook);
  }

  /// @notice Outbound rate limits apply to the amount after the token fee.
  function test_validateLockOrBurn_ConsumesRateLimitAfterFee() public {
    _setInlineFeeConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getFeeConfig(0, 1000));
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setCaller(address(s_pool));

    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );

    assertEq(1e17, lane.feeAmount);
    assertEq(CAPACITY - 9e17, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice The V1 path charges no token fee, so rate limits apply to the full amount.
  function test_validateLockOrBurn_WithoutTokenFee() public {
    _setInlineFeeConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getFeeConfig(0, 1000));
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setCaller(address(s_pool));

    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", false
    );

    assertEq(0, lane.feeAmount);
    assertEq(CAPACITY - AMOUNT, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice The requested finality reaches all three stores. Token arguments reach the additional pool data
  /// provider. The fee store applies the FTF fee, and the security store consumes the FTF outbound bucket for the
  /// remaining amount. The default finality bucket is unchanged.
  function test_validateLockOrBurn_FastFinality() public {
    bytes4 requestedFinalityConfig = bytes4(uint32(10));
    bytes memory tokenArgs = abi.encode("tokenArgs");
    bytes memory additionalPoolData = abi.encode(bytes32("uid"));
    _setAllowedFinalityConfig(s_tokenA, bytes4(uint32(5)));
    _setAdditionalPoolDataProvider(s_tokenA, s_additionalPoolDataProvider);
    // The pool data with additional pool data is longer than the minimum bytes overhead.
    _setInlineDestOverheadConfig(s_tokenA, DEST_CHAIN_SELECTOR, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, 256);
    _setInlineFeeConfig(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      FeePolicyStore.FeeConfig({
        finalityFeeUSDCents: 1, fastFinalityFeeUSDCents: 2, finalityTransferFeeBps: 100, fastFinalityTransferFeeBps: 300
      })
    );
    _setInlineRateLimitConfigByFinality(s_tokenA, CAPACITY, 2 * CAPACITY);
    bytes memory providerCallData = abi.encodeCall(
      IAdditionalPoolDataProvider.getAdditionalPoolData,
      (s_tokenA, DEST_CHAIN_SELECTOR, requestedFinalityConfig, tokenArgs)
    );
    vm.mockCall(s_additionalPoolDataProvider, providerCallData, abi.encode(additionalPoolData));
    _setCaller(address(s_pool));

    vm.expectCall(s_additionalPoolDataProvider, providerCallData);
    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, requestedFinalityConfig, tokenArgs, true
    );

    uint256 expectedFee = 3e16;
    assertEq(expectedFee, lane.feeAmount);
    assertEq(abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), abi.encode(s_tokenA), additionalPoolData), lane.destPoolData);
    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(2 * CAPACITY - (AMOUNT - expectedFee), state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY, state.outbound.tokens);
  }

  /// @notice Returns the hook from the resolved security record.
  function test_validateLockOrBurn_ReturnsHook() public {
    address hook = makeAddr("hook");
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), new address[](0), hook);
    _setCaller(address(s_pool));

    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );

    assertEq(hook, lane.hook);
  }

  // Reverts

  /// @notice Only the pool that the TokenAdminRegistry assigns to the token can consume its outbound rate limit.
  function test_validateLockOrBurn_RevertWhen_CallerIsNotTokenPool() public {
    _setCaller(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(PoolConfigManager.CallerIsNotTokenPool.selector, STRANGER, s_tokenA));
    s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );
  }

  /// @notice The fee store rejects pool data that exceeds the destination bytes overhead. Additional pool data
  /// increases the total length to 224 bytes.
  function test_validateLockOrBurn_RevertWhen_PoolDataTooLarge() public {
    _setAdditionalPoolDataProvider(s_tokenA, s_additionalPoolDataProvider);
    vm.mockCall(
      s_additionalPoolDataProvider,
      abi.encodeWithSelector(IAdditionalPoolDataProvider.getAdditionalPoolData.selector),
      abi.encode(abi.encode(bytes32("uid")))
    );
    _setCaller(address(s_pool));

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector, s_tokenA, DEST_CHAIN_SELECTOR, 224, MIN_DEST_BYTES_OVERHEAD
      )
    );
    s_poolConfigManager.validateLockOrBurn(
      s_tokenA, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );
  }
}
