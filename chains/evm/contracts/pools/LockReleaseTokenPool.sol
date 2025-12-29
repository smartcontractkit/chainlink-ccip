// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../libraries/Pool.sol";
import {ERC20LockBox} from "./ERC20LockBox.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
/// Because of lock/unlock requiring liquidity, this pool contract also has function to add and remove
/// liquidity. This allows for proper bookkeeping for both user and liquidity provider balances.
/// @dev One token per LockReleaseTokenPool.
contract LockReleaseTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error InsufficientLiquidity();

  event LiquidityTransferred(address indexed from, uint256 amount);
  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed provider, uint256 indexed amount);

  string public constant override typeAndVersion = "LockReleaseTokenPool 1.7.0-dev";

  /// @notice The lock box for the token pool.
  ERC20LockBox internal immutable i_lockBox;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    address lockBox
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {
    if (lockBox == address(0)) revert ZeroAddressInvalid();

    ERC20LockBox lockBoxContract = ERC20LockBox(lockBox);
    if (!lockBoxContract.isTokenSupported(address(token))) {
      revert InvalidToken(address(token));
    }
    token.safeApprove(lockBox, type(uint256).max);
    i_lockBox = lockBoxContract;
  }

  /// @notice Locks the tokens in the lockBox.
  /// @dev The router has already transferred the full amount to this contract before calling lockOrBurn.
  /// For V1 the amount = full amount. For V2 the amount = destTokenAmount (after fees), and fees remain on this contract.
  /// @param lockOrBurnIn The lock or burn input parameters.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out) {
    // super.lockOrBurn will validate the lockOrBurnIn and revert if invalid.
    out = super.lockOrBurn(lockOrBurnIn);
    i_lockBox.deposit(address(i_token), lockOrBurnIn.remoteChainSelector, lockOrBurnIn.amount);
    return out;
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256) {
    (Pool.LockOrBurnOutV1 memory out, uint256 destTokenAmount) =
      super.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    i_lockBox.deposit(address(i_token), lockOrBurnIn.remoteChainSelector, destTokenAmount);
    return (out, destTokenAmount);
  }

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    Pool.ReleaseOrMintOutV1 memory out = super.releaseOrMint(releaseOrMintIn, blockConfirmationRequested);
    // Release tokens from the lock box to the receiver.
    i_lockBox.withdraw(
      address(i_token), releaseOrMintIn.remoteChainSelector, out.destinationAmount, releaseOrMintIn.receiver
    );
    return out;
  }

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }
}
