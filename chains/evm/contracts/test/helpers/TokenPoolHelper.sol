// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract TokenPoolHelper is TokenPool {
  using EnumerableSet for EnumerableSet.Bytes32Set;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function encodeLocalDecimals() external view returns (bytes memory) {
    return _encodeLocalDecimals();
  }

  function parseRemoteDecimals(
    bytes memory sourcePoolData
  ) external view returns (uint256) {
    return _parseRemoteDecimals(sourcePoolData);
  }

  function calculateLocalAmount(uint256 remoteAmount, uint8 remoteDecimals) external view returns (uint256) {
    return _calculateLocalAmount(remoteAmount, remoteDecimals);
  }

  function validateLockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) external {
    _validateLockOrBurn(lockOrBurnIn);
  }

  function validateReleaseOrMint(Pool.ReleaseOrMintInV1 calldata releaseOrMintIn, uint256 localAmount) external {
    _validateReleaseOrMint(releaseOrMintIn, localAmount);
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
}
