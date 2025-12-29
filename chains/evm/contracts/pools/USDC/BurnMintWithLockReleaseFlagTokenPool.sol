// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../BurnMintTokenPool.sol";

bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// @notice A standard BurnMintTokenPool with modified destPoolData so that the remote pool knows to release tokens
/// instead of minting. This enables interoperability with HybridLockReleaseUSDCTokenPool which uses
// the destPoolData to determine whether to mint or release tokens.
/// @dev The only difference between this contract and BurnMintTokenPool is the destPoolData returns the
/// abi-encoded LOCK_RELEASE_FLAG instead of the local token decimals.
contract BurnMintWithLockReleaseFlagTokenPool is BurnMintTokenPool {
  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) BurnMintTokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {}

  /// @notice Burn the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  /// @dev Performs the exact same functionality as BurnMintTokenPool, but returns the LOCK_RELEASE_FLAG
  /// as the destPoolData to signal to the remote pool to release tokens instead of minting them.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public override returns (Pool.LockOrBurnOutV1 memory) {
    uint256 feeAmount = _getFee(lockOrBurnIn, WAIT_FOR_FINALITY);
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "", feeAmount);

    _lockOrBurn(lockOrBurnIn.amount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    // LOCK_RELEASE_FLAG = bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE"))
    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector), destPoolData: abi.encode(LOCK_RELEASE_FLAG)
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Since the remote token is always canonical USDC, the decimals should always be 6 for remote tokens,
    // which enables potentially local non-canonical USDC with different decimals to be minted.
    uint256 localAmount = _calculateLocalAmount(releaseOrMintIn.sourceDenominatedAmount, 6);

    _validateReleaseOrMint(releaseOrMintIn, localAmount, WAIT_FOR_FINALITY);

    IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, localAmount);

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
