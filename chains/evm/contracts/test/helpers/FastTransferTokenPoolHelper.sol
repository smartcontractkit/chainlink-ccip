// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../pools/FastTransferTokenPoolAbstract.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

contract FastTransferTokenPoolHelper is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "FastTransferTokenPoolHelper 1.6.1";

  /// @dev Accumulated pool fees for lock/release pool accounting
  uint256 internal s_accumulatedPoolFees;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    uint64 sourceChainSelector
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router, sourceChainSelector) {}

  // Implementation of abstract functions
  function _handleFastTransferLockOrBurn(uint64, address sender, uint256 amount) internal override {
    // For testing, we'll just transfer tokens from sender to this contract
    getToken().safeTransferFrom(sender, address(this), amount);
  }

  /// @notice Validates settlement prerequisites - simple implementation for testing
  function _validateSettlement(uint64, bytes memory) internal view override {
    // For testing, we'll do minimal validation
    // Real implementations would check RMN curse and source pool validation
  }

  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    getToken().safeTransfer(receiver, amount);
  }

  /// @notice Override for lock/release pools - use accounting instead of minting
  /// @dev Since this is a lock/release pool that cannot mint tokens, we need to use
  /// the accounting-based approach for pool fee management.
  function _handleFastFillReimbursement(
    bytes32,
    uint64,
    address filler,
    uint256 fillerReimbursementAmount,
    uint256 poolReimbursementAmount
  ) internal override {
    // Reimburse the filler with their original amount plus their fee
    _releaseOrMint(filler, fillerReimbursementAmount);

    if (poolReimbursementAmount > 0) {
      // For lock/release pools: accumulate pool fees in storage since we can't mint new tokens
      s_accumulatedPoolFees += poolReimbursementAmount;
    }
  }

  /// @notice Override to return accumulated pool fees from storage for lock/release pools
  function getAccumulatedPoolFees() public view override returns (uint256) {
    return s_accumulatedPoolFees;
  }

  /// @notice Override to withdraw accumulated pool fees from storage for lock/release pools
  function withdrawPoolFees(
    address recipient
  ) external override onlyOwner {
    uint256 amount = s_accumulatedPoolFees;
    if (amount > 0) {
      s_accumulatedPoolFees = 0;
      _releaseOrMint(recipient, amount);
      emit PoolFeeWithdrawn(recipient, amount);
    }
  }
}
