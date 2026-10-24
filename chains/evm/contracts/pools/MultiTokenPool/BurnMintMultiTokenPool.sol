// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {MultiTokenPool} from "./MultiTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";

/// @notice A multi-token pool that burns 3rd-party tokens on send and mints them on receive.
/// @dev Serves the four burn flavors, which match BurnMintTokenPool, BurnFromMintTokenPool, BurnWithFromMintTokenPool
/// and BurnToAddressMintTokenPool. Rejects a token with the Lock flavor, so it never mints a lock and release token.
/// The pool needs the minter role of each token, and the burner role of each token without the BurnToAddress flavor.
contract BurnMintMultiTokenPool is MultiTokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "BurnMintMultiTokenPool 2.1.0-dev";
  }

  constructor(
    address rmnProxy,
    address router,
    address poolConfigManager
  ) MultiTokenPool(rmnProxy, router, poolConfigManager) {}

  /// @inheritdoc MultiTokenPool
  function _isSupportedFlavor(
    IPoolConfigManager.LockOrBurnFlavor flavor
  ) internal pure virtual override returns (bool) {
    return flavor != IPoolConfigManager.LockOrBurnFlavor.Lock;
  }

  /// @inheritdoc MultiTokenPool
  /// @dev Burns tokens held by the pool. The token fee stays in the pool. BurnToAddress sends the tokens to an address
  /// which can never transfer them, making the tokens unrecoverable without reducing the total supply.
  /// Some tokens allow burning from the sender without approval, but not all do. To be safe, BurnFrom and BurnWithFrom
  /// approve the pool to burn from the pool. The single-token pools approve in the constructor. This pool learns its
  /// tokens later, so it approves the maximum on a send when the allowance is below the amount.
  function _lockOrBurn(
    address token,
    uint64, // remoteChainSelector
    uint256 amount,
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal virtual override {
    if (flavor == IPoolConfigManager.LockOrBurnFlavor.Burn) {
      IBurnMintERC20(token).burn(amount);
    } else if (flavor == IPoolConfigManager.LockOrBurnFlavor.BurnToAddress) {
      IERC20(token).safeTransfer(lockOrBurnAddress, amount);
    } else if (
      flavor == IPoolConfigManager.LockOrBurnFlavor.BurnFrom
        || flavor == IPoolConfigManager.LockOrBurnFlavor.BurnWithFrom
    ) {
      if (IERC20(token).allowance(address(this), address(this)) < amount) {
        IERC20(token).forceApprove(address(this), type(uint256).max);
      }
      if (flavor == IPoolConfigManager.LockOrBurnFlavor.BurnFrom) {
        IBurnMintERC20(token).burnFrom(address(this), amount);
      } else {
        IBurnMintERC20(token).burn(address(this), amount);
      }
    } else {
      revert UnsupportedLockOrBurnFlavor(token, flavor);
    }
  }

  /// @inheritdoc MultiTokenPool
  function _releaseOrMint(
    address token,
    address receiver,
    uint256 amount,
    uint64, // remoteChainSelector
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address // lockOrBurnAddress
  ) internal virtual override {
    if (flavor == IPoolConfigManager.LockOrBurnFlavor.Lock) revert UnsupportedLockOrBurnFlavor(token, flavor);
    IBurnMintERC20(token).mint(receiver, amount);
  }
}
