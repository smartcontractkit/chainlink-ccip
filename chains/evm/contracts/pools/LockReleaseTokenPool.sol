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
  event RebalancerSet(address oldRebalancer, address newRebalancer);

  string public constant override typeAndVersion = "LockReleaseTokenPool 1.7.0-dev";

  /// @notice The address of the rebalancer.
  address internal s_rebalancer;

  /// @notice Fee accrued on token transfers.
  uint256 internal s_accruedFees;
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
    token.safeApprove(lockBox, type(uint256).max);
    i_lockBox = ERC20LockBox(lockBox);
  }

  /// @inheritdoc TokenPool
  /// @notice accounts for accrued fees when token is locked.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out, uint256 destTokenAmount) {
    (out, destTokenAmount) = super.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    // Accrue the in-token fee (original amount minus destination amount).
    uint256 feeAmount = lockOrBurnIn.amount - destTokenAmount;
    if (feeAmount != 0) {
      s_accruedFees += feeAmount;
    }

    return (out, destTokenAmount);
  }

  function _lockOrBurn(
    uint256 amount
  ) internal virtual override {
    i_lockBox.deposit(address(i_token), amount);
  }

  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    uint256 availableLiquidity = _availableLiquidity();
    if (amount > availableLiquidity) revert InsufficientLiquidity();
    i_lockBox.withdraw(address(i_token), amount, receiver);
  }

  /// @notice Gets rebalancer, can be address(0) if none is configured.
  /// @return The current liquidity manager.
  function getRebalancer() external view returns (address) {
    return s_rebalancer;
  }

  /// @notice Sets the rebalancer address.
  /// @dev Address(0) can be used to disable the rebalancer.
  /// @dev Only callable by the owner.
  function setRebalancer(
    address rebalancer
  ) external onlyOwner {
    address oldRebalancer = s_rebalancer;

    s_rebalancer = rebalancer;

    emit RebalancerSet(oldRebalancer, rebalancer);
  }

  /// @notice Adds liquidity to the pool. The tokens should be approved first.
  /// @param amount The amount of liquidity to provide.
  function provideLiquidity(
    uint256 amount
  ) external {
    if (s_rebalancer != msg.sender) revert Unauthorized(msg.sender);

    i_token.safeTransferFrom(msg.sender, address(this), amount);
    i_lockBox.deposit(address(i_token), amount);
    emit LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param amount The amount of liquidity to remove.
  function withdrawLiquidity(
    uint256 amount
  ) external {
    if (s_rebalancer != msg.sender) revert Unauthorized(msg.sender);

    if (_availableLiquidity() < amount) revert InsufficientLiquidity();
    i_lockBox.withdraw(address(i_token), amount, msg.sender);
    emit LiquidityRemoved(msg.sender, amount);
  }

  /// @notice This function can be used to transfer liquidity from an older version of the pool to this pool. To do so
  /// this pool will have to be set as the rebalancer in the older version of the pool. This allows it to transfer the
  /// funds in the old pool to the new pool.
  /// @dev When upgrading a LockRelease pool, this function can be called at the same time as the pool is changed in the
  /// TokenAdminRegistry. This allows for a smooth transition of both liquidity and transactions to the new pool.
  /// Alternatively, when no multicall is available, a portion of the funds can be transferred to the new pool before
  /// changing which pool CCIP uses, to ensure both pools can operate. Then the pool should be changed in the
  /// TokenAdminRegistry, which will activate the new pool. All new transactions will use the new pool and its
  /// liquidity. Finally, the remaining liquidity can be transferred to the new pool using this function one more time.
  /// @param from The address of the old pool.
  /// @param amount The amount of liquidity to transfer. If uint256.max is passed, all liquidity will be transferred.
  function transferLiquidity(address from, uint256 amount) external onlyOwner {
    if (amount == type(uint256).max) {
      amount = i_token.balanceOf(from);
    }

    LockReleaseTokenPool(from).withdrawLiquidity(amount);

    emit LiquidityTransferred(from, amount);
  }

  /// @notice Withdraw only accrued fees; user liquidity (including bridged funds) remains untouched.
  function withdrawFeeTokens(address[] calldata feeTokens, address recipient) external override onlyOwner {
    for (uint256 i = 0; i < feeTokens.length; ++i) {
      address feeToken = feeTokens[i];
      uint256 amountToWithdraw;
      if (feeToken == address(i_token)) {
        uint256 balance = i_token.balanceOf(address(this));
        amountToWithdraw = s_accruedFees > balance ? balance : s_accruedFees;
        if (amountToWithdraw != 0) {
          s_accruedFees -= amountToWithdraw;
          i_token.safeTransfer(recipient, amountToWithdraw);
          emit FeeTokenWithdrawn(recipient, feeToken, amountToWithdraw);
        }
      } else {
        amountToWithdraw = IERC20(feeToken).balanceOf(address(this));
        if (amountToWithdraw != 0) {
          IERC20(feeToken).safeTransfer(recipient, amountToWithdraw);
          emit FeeTokenWithdrawn(recipient, feeToken, amountToWithdraw);
        }
      }
    }
  }

  function _availableLiquidity() internal view returns (uint256) {
    uint256 balance = i_token.balanceOf(address(i_lockBox));
    return balance;
  }
}
