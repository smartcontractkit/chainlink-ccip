// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {MultiTokenPool} from "./MultiTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";

/// @notice A multi-token pool that burns on send and mints on receive.
/// @dev Serves the four burn flavors, which match the four single-token burn and mint pools. Rejects a token with the
/// Lock flavor, so it never mints a lock and release token. The pool needs the mint and burn roles of each token.
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
  /// @dev BurnFrom and BurnWithFrom spend an allowance of the pool for itself. The pool approves itself for the maximum
  /// when its allowance is below the amount. The approval is never revoked, and no other address can use it.
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
    } else {
      if (IERC20(token).allowance(address(this), address(this)) < amount) {
        IERC20(token).forceApprove(address(this), type(uint256).max);
      }
      if (flavor == IPoolConfigManager.LockOrBurnFlavor.BurnFrom) {
        IBurnMintERC20(token).burnFrom(address(this), amount);
      } else {
        IBurnMintERC20(token).burn(address(this), amount);
      }
    }
  }

  /// @inheritdoc MultiTokenPool
  function _releaseOrMint(
    address token,
    address receiver,
    uint256 amount,
    uint64, // remoteChainSelector
    IPoolConfigManager.LockOrBurnFlavor, // flavor
    address // lockOrBurnAddress
  ) internal virtual override {
    IBurnMintERC20(token).mint(receiver, amount);
  }
}
