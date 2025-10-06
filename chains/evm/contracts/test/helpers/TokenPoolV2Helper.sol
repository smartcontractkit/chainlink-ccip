// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {TokenPool} from "../../poolsV2/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice Helper contract for testing TokenPool V2 functionality.
contract TokenPoolV2Helper is TokenPool {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function getFastFinalityConfig()
    external
    view
    returns (uint16 finalityThreshold, uint16 fastTransferFeeBps, uint256 maxAmountPerRequest)
  {
    FastFinalityConfig storage config = s_finalityConfig;
    return (config.finalityThreshold, config.fastTransferFeeBps, config.maxAmountPerRequest);
  }

  function getFastOutboundBucket(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory bucket) {
    return s_finalityConfig.outboundRateLimiterConfig[remoteChainSelector];
  }

  function getFastInboundBucket(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory bucket) {
    return s_finalityConfig.inboundRateLimiterConfig[remoteChainSelector];
  }

  function validateLockOrBurn(Pool.LockOrBurnInV1 calldata lockOrBurnIn, uint16 finality) external {
    _validateLockOrBurn(lockOrBurnIn, finality);
  }

  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 finality
  ) external returns (uint256) {
    _validateReleaseOrMint(releaseOrMintIn, localAmount, finality);
    return localAmount;
  }

  function applyFee(Pool.LockOrBurnInV1 calldata lockOrBurnIn, uint16 finality) external view returns (uint256) {
    return _applyFee(lockOrBurnIn, finality);
  }
}
