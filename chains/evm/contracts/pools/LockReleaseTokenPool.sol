// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

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

  string public constant override typeAndVersion = "LockReleaseTokenPool 1.6.3-dev";

  /// @notice The address of the rebalancer.
  address internal s_rebalancer;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    i_token.safeTransfer(receiver, amount);
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
    emit LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param amount The amount of liquidity to remove.
  function withdrawLiquidity(
    uint256 amount
  ) external {
    if (s_rebalancer != msg.sender) revert Unauthorized(msg.sender);

    if (i_token.balanceOf(address(this)) < amount) revert InsufficientLiquidity();
    i_token.safeTransfer(msg.sender, amount);
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
}
