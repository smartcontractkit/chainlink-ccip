// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

/// @notice Checks the converted amount, returned profile, hook, and caller authorization. Store tests cover the
/// rules within each store.
contract PoolConfigManager_validateReleaseOrMint is PoolConfigManagerTransferSetup {
  uint8 internal constant REMOTE_DECIMALS = 6;

  address internal s_receiver = makeAddr("receiver");

  /// @notice Converts the amount to local token decimals before applying the inbound rate limit.
  function test_validateReleaseOrMint() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setCaller(address(s_pool));

    IPoolConfigManager.ResolvedInboundTransfer memory transfer = s_poolConfigManager.validateReleaseOrMint(
      _getReleaseOrMintIn(
        s_tokenA, DEST_CHAIN_SELECTOR, 1e6, abi.encode(uint256(REMOTE_DECIMALS), abi.encode(s_remoteTokenA), bytes(""))
      ),
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    assertEq(1e18, transfer.localAmount);
    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.Burn), uint8(transfer.lockOrBurnFlavor));
    assertEq(address(0), transfer.lockOrBurnAddress);
    assertEq(address(0), transfer.hook);
    assertEq(CAPACITY - 1e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).inbound.tokens);
  }

  /// @notice Returns the token profile and hook.
  function test_validateReleaseOrMint_ReturnsProfileAndHook() public {
    address hook = makeAddr("hook");
    address burnAddress = makeAddr("burnAddress");
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), new address[](0), hook);
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    PolicyConfig.TokenProfileArgs memory tokenProfile = _getStoredTokenProfile(s_tokenA);
    tokenProfile.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    tokenProfile.lockOrBurnAddress = burnAddress;
    _setTokenProfile(tokenProfile);
    _setCaller(address(s_pool));

    IPoolConfigManager.ResolvedInboundTransfer memory transfer = s_poolConfigManager.validateReleaseOrMint(
      _getReleaseOrMintIn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, _getPoolData(s_remoteTokenA)),
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.BurnToAddress), uint8(transfer.lockOrBurnFlavor));
    assertEq(burnAddress, transfer.lockOrBurnAddress);
    assertEq(hook, transfer.hook);
  }

  /// @notice The requested finality reaches the additional pool data provider. FTF transfers consume only the FTF
  /// inbound bucket. The default finality bucket is unchanged. Inbound transfers do not check finality against the
  /// token profile.
  function test_validateReleaseOrMint_FastFinality() public {
    bytes4 requestedFinalityConfig = bytes4(uint32(10));
    _setAdditionalPoolDataProvider(s_tokenA, s_additionalPoolDataProvider);
    _setInlineRateLimitConfigByFinality(s_tokenA, CAPACITY, 2 * CAPACITY);
    bytes memory providerCallData = abi.encodeCall(
      IAdditionalPoolDataProvider.isAdditionalPoolDataValid,
      (s_tokenA, DEST_CHAIN_SELECTOR, requestedFinalityConfig, bytes(""))
    );
    vm.mockCall(s_additionalPoolDataProvider, providerCallData, abi.encode(true));
    _setCaller(address(s_pool));

    vm.expectCall(s_additionalPoolDataProvider, providerCallData);
    IPoolConfigManager.ResolvedInboundTransfer memory transfer = s_poolConfigManager.validateReleaseOrMint(
      _getReleaseOrMintIn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, _getPoolData(s_remoteTokenA)), requestedFinalityConfig
    );

    assertEq(1e18, transfer.localAmount);
    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(2 * CAPACITY - 1e18, state.fastFinalityInbound.tokens);
    assertEq(CAPACITY, state.inbound.tokens);
  }

  // Reverts

  /// @notice Only the pool that the TokenAdminRegistry assigns to the token can consume its inbound rate limit.
  function test_validateReleaseOrMint_RevertWhen_CallerIsNotTokenPool() public {
    _setCaller(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(PoolConfigManager.CallerIsNotTokenPool.selector, STRANGER, s_tokenA));
    s_poolConfigManager.validateReleaseOrMint(
      _getReleaseOrMintIn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, _getPoolData(s_remoteTokenA)),
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  /// @dev A message from the remote pool.
  function _getReleaseOrMintIn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 sourceAmount,
    bytes memory sourcePoolData
  ) internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(OWNER),
      remoteChainSelector: remoteChainSelector,
      receiver: s_receiver,
      sourceDenominatedAmount: sourceAmount,
      localToken: localToken,
      sourcePoolAddress: abi.encode(s_remotePool),
      sourcePoolData: sourcePoolData,
      offchainTokenData: ""
    });
  }
}
