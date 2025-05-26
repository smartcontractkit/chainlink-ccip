// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

abstract contract BurnMintTokenPoolAbstract is TokenPool {
  /// @notice Contains the specific lock or burn token logic for a pool.
  /// @dev overriding this method allows us to create pools with different lock/burn signatures
  /// without duplicating the underlying logic.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual;

  /// @notice Contains the specific release or mint token logic for a pool.
  /// @dev overriding this method allows us to create pools with different release/mint signatures
  /// without duplicating the underlying logic.
  function _releaseOrMint(address receiver, uint256 amount) internal virtual {
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
  }

  /// @notice Burn the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    _lockOrBurn(lockOrBurnIn.amount);

    emit LockedOrBurned(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: _encodeLocalDecimals()
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    // Calculate the local amount
    uint256 localAmount =
      _calculateLocalAmount(releaseOrMintIn.amount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData));

    // Mint to the receiver
    _releaseOrMint(releaseOrMintIn.receiver, localAmount);

    emit ReleasedOrMinted(msg.sender, releaseOrMintIn.receiver, localAmount);

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }
}
