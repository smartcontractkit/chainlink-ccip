// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";

contract MaybeRevertingBurnMintTokenPool is BurnMintTokenPool {
  bytes public s_revertReason = "";
  bytes public s_sourceTokenData = "";
  uint256 public s_releaseOrMintMultiplier = 1;

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) BurnMintTokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function setShouldRevert(
    bytes calldata revertReason
  ) external {
    s_revertReason = revertReason;
  }

  function setSourceTokenData(
    bytes calldata sourceTokenData
  ) external {
    s_sourceTokenData = sourceTokenData;
  }

  function setReleaseOrMintMultiplier(
    uint256 multiplier
  ) external {
    s_releaseOrMintMultiplier = multiplier;
  }

  function _lockOrBurn(
    uint256 amount
  ) internal override {
    IBurnMintERC20(address(i_token)).burn(amount);
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out) {
    out = super.lockOrBurn(lockOrBurnIn);

    bytes memory revertReason = s_revertReason;
    if (revertReason.length != 0) {
      assembly {
        revert(add(32, revertReason), mload(revertReason))
      }
    }

    if (s_sourceTokenData.length != 0) {
      out.destPoolData = s_sourceTokenData;
    }
    return out;
  }

  function _releaseOrMint(address receiver, uint256 amount) internal override {
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
  }

  /// @notice Reverts depending on the value of `s_revertReason`
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory out) {
    out = super.releaseOrMint(releaseOrMintIn);

    bytes memory revertReason = s_revertReason;
    if (revertReason.length != 0) {
      assembly {
        revert(add(32, revertReason), mload(revertReason))
      }
    }

    return out;
  }

  function _calculateLocalAmount(uint256 remoteAmount, uint8 remoteDecimals) internal view override returns (uint256) {
    return super._calculateLocalAmount(remoteAmount, remoteDecimals) * s_releaseOrMintMultiplier;
  }
}
