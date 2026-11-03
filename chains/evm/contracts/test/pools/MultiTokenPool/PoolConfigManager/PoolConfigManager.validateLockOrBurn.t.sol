// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

/// @notice The manager resolves the lane, computes the token fee and consumes the outbound rate limit of the amount
/// after the fee. The store-level rules are covered by the store tests, so these tests cover the orchestration: the
/// arguments that reach each store, the returned lane, the fee deduction and the caller check.
contract PoolConfigManager_validateLockOrBurn is PoolConfigManagerTransferSetup {
  uint256 internal constant AMOUNT = 1e18;

  address internal s_additionalPoolDataProvider = makeAddr("additionalPoolDataProvider");

  /// @dev The connection store accepts a provider only if it has code. The tests mock each call to it.
  function setUp() public virtual override {
    super.setUp();
    vm.etch(s_additionalPoolDataProvider, hex"00");
  }

  /// @notice Token A has an inline connection record, so the lane carries its remote token and the pool data of the
  /// token. The DEFAULT fee record has zero fees and the security record is empty, so the fee is zero and there is no
  /// hook.
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

  /// @notice The security store consumes the outbound rate limit of the amount after the token fee.
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

  /// @notice The V1 path charges no token fee, so the outbound rate limit consumes the whole amount.
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

  /// @notice The requested finality and the token args reach each store. The connection store passes them to the
  /// additional pool data provider. The fee store takes the FTF fee bps, and the security store consumes the FTF
  /// outbound bucket of the amount after that fee. The default outbound bucket keeps its tokens.
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

  /// @notice The lane carries the hook of the security record.
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

  /// @notice The PCM passes the length of the pool data to the fee store, which rejects pool data longer than the
  /// destination bytes overhead. The additional pool data makes the pool data 224 bytes long.
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

  /// @dev Sets the additional pool data provider and keeps the other values of the token profile.
  function _setAdditionalPoolDataProvider(
    address localToken,
    address additionalPoolDataProvider
  ) internal {
    PolicyConfig.TokenProfileArgs memory tokenProfile = _getStoredTokenProfile(localToken);
    tokenProfile.additionalPoolDataProvider = additionalPoolDataProvider;
    _setTokenProfile(tokenProfile);
  }
}
