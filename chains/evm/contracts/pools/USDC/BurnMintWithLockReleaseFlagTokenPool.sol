// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";
import {IPoolV1} from "../../interfaces/IPool.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";

import {Pool} from "../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../BurnMintTokenPool.sol";

bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// @notice A standard BurnMintTokenPool with modified destPoolData so that the remote pool knows to release tokens
/// instead of minting. This enables interoperability with HybridLockReleaseUSDCTokenPool which uses
// the destPoolData to determine whether to mint or release tokens.
/// @dev The only difference between this contract and BurnMintTokenPool is the destPoolData returns the
/// abi-encoded LOCK_RELEASE_FLAG instead of the local token decimals.
contract BurnMintWithLockReleaseFlagTokenPool is BurnMintTokenPool {
  /// @dev Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure override returns (string memory) {
    return "BurnMintWithLockReleaseFlagTokenPool 1.7.0-dev";
  }

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) BurnMintTokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {}

  /// @dev Performs the exact same functionality as BurnMintTokenPool, but returns the LOCK_RELEASE_FLAG
  /// as the destPoolData to signal to the remote pool to release tokens instead of minting them.
  /// @inheritdoc IPoolV1
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public override returns (Pool.LockOrBurnOutV1 memory) {
    Pool.LockOrBurnOutV1 memory out = super.lockOrBurn(lockOrBurnIn);
    out.destPoolData = abi.encode(LOCK_RELEASE_FLAG);
    return out;
  }

  /// @dev Performs the exact same functionality as BurnMintTokenPool, but returns the LOCK_RELEASE_FLAG
  /// as the destPoolData to signal to the remote pool to release tokens instead of minting them.
  /// @inheritdoc IPoolV2
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public override returns (Pool.LockOrBurnOutV1 memory out, uint256 destTokenAmount) {
    (out, destTokenAmount) = super.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    out.destPoolData = abi.encode(LOCK_RELEASE_FLAG);
    return (out, destTokenAmount);
  }

  /// @inheritdoc IPoolV2
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Since the remote token is always canonical USDC, the decimals should always be 6 for remote tokens,
    // which enables potentially local non-canonical USDC with different decimals to be minted.
    uint256 localAmount = _calculateLocalAmount(releaseOrMintIn.sourceDenominatedAmount, 6);

    _validateReleaseOrMint(releaseOrMintIn, localAmount, blockConfirmationRequested);

    _releaseOrMint(releaseOrMintIn.receiver, localAmount, releaseOrMintIn.remoteChainSelector);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }
}
