// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {BurnMintMultiTokenPool} from "../../../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

/// @notice Adds a pool that serves tokens A and B through the pool config manager. The token admin registry names the
/// pool for both tokens, so the tests call the validate functions as the pool.
abstract contract PoolConfigManagerTransferSetup is PoolConfigManagerRolesSetup {
  BurnMintMultiTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();
    s_pool = new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(s_sourceRouter), address(s_poolConfigManager));
    _registerPool(s_tokenA);
    _registerPool(s_tokenB);
  }

  function _registerPool(
    address token
  ) internal {
    s_tokenAdminRegistry.proposeAdministrator(token, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(token);
    s_tokenAdminRegistry.setPool(token, address(s_pool));
  }

  /// @dev Pool data with the source token, in the decimals of the test tokens, and no additional pool data.
  function _getPoolData(
    address sourceToken
  ) internal pure returns (bytes memory) {
    return abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), abi.encode(sourceToken), bytes(""));
  }

  /// @dev An inline rate limit record with one capacity for the default finality buckets and another for the FTF
  /// buckets, so a consumed bucket shows which class it belongs to.
  function _setInlineRateLimitConfigByFinality(
    address localToken,
    uint128 capacity,
    uint128 fastFinalityCapacity
  ) internal {
    RateLimiter.Config memory config = _getRateLimitConfig(capacity);
    RateLimiter.Config memory fastFinalityConfig = _getRateLimitConfig(fastFinalityCapacity);
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config,
        inbound: config,
        fastFinalityOutbound: fastFinalityConfig,
        fastFinalityInbound: fastFinalityConfig
      })
    });
    s_poolConfigManager.applyInlineRateLimitConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
