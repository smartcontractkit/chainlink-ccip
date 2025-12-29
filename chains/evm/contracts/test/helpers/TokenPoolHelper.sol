// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract TokenPoolHelper is TokenPool {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {}

  function encodeLocalDecimals() external view returns (bytes memory) {
    return _encodeLocalDecimals();
  }

  function parseRemoteDecimals(
    bytes memory sourcePoolData
  ) external view returns (uint256) {
    return _parseRemoteDecimals(sourcePoolData);
  }

  function calculateLocalAmount(
    uint256 remoteAmount,
    uint8 remoteDecimals
  ) external view returns (uint256) {
    return _calculateLocalAmount(remoteAmount, remoteDecimals);
  }

  function validateLockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) external {
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "", 0);
  }

  function validateLockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality,
    bytes calldata tokenArgs,
    uint256 feeAmount
  ) external {
    _validateLockOrBurn(lockOrBurnIn, finality, tokenArgs, feeAmount);
  }

  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 finality
  ) external returns (uint256) {
    _validateReleaseOrMint(releaseOrMintIn, localAmount, finality);
    return localAmount;
  }

  function applyFee(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality
  ) external view returns (uint256) {
    return lockOrBurnIn.amount - _getFee(lockOrBurnIn, finality);
  }

  function getFee(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality
  ) external view returns (uint256) {
    return _getFee(lockOrBurnIn, finality);
  }

  function onlyOnRampModifier(
    uint64 remoteChainSelector
  ) external view {
    _onlyOnRamp(remoteChainSelector);
  }

  function onlyOffRampModifier(
    uint64 remoteChainSelector
  ) external view {
    _onlyOffRamp(remoteChainSelector);
  }

  function getCustomMinBlockConfirmation() external view returns (uint16 minBlockConfirmation) {
    return s_minBlockConfirmation;
  }

  function getFastOutboundBucket(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory bucket) {
    return s_outboundRateLimiterConfig[remoteChainSelector];
  }

  function getFastInboundBucket(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory bucket) {
    return s_inboundRateLimiterConfig[remoteChainSelector];
  }
}
