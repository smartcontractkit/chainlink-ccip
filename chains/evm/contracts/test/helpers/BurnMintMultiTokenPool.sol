// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../../libraries/Pool.sol";
import {MultiTokenPool} from "./MultiTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintMultiTokenPool is MultiTokenPool {
  constructor(
    IERC20[] memory tokens,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) MultiTokenPool(tokens, allowlist, rmnProxy, router) {}

  /// @notice Burn the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) external virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    IBurnMintERC20(lockOrBurnIn.localToken).burn(lockOrBurnIn.amount);

    emit Burned(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector),
      destPoolData: ""
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) external virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    // Mint to the receiver
    IBurnMintERC20(releaseOrMintIn.localToken).mint(msg.sender, releaseOrMintIn.sourceDenominatedAmount);

    emit Minted(msg.sender, releaseOrMintIn.receiver, releaseOrMintIn.sourceDenominatedAmount);

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.sourceDenominatedAmount});
  }
}
