// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../interfaces/ILockBox.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";

/// @notice A variation on Lock Release token pools where liquidity is shared among some chains, and stored independently
/// for others. Chains which do not share liquidity are known as siloed chains.
contract SiloedLockReleaseTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error InsufficientLiquidity(uint256 availableLiquidity, uint256 requestedAmount);
  error LockBoxNotConfigured(uint64 remoteChainSelector);

  struct LockBoxConfig {
    uint64 remoteChainSelector;
    address lockBox;
  }

  /// @notice Lock boxes keyed by chain selector.
  /// @dev We can have a many-to-one mapping of remote chain selectors to lock boxes. This allows for chains to share or isolate liquidity.
  mapping(uint64 remoteChainSelector => ILockBox lockBox) internal s_lockBoxes;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {
  }

  /// @dev Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "SiloedLockReleaseTokenPool 1.7.0-dev";
  }

  /// @inheritdoc TokenPool
  /// @dev Deposits the amount into the lockbox configured for the remote chain selector.
  function _lockOrBurn(
    uint64 remoteChainSelector,
    uint256 amount
  ) internal override {
    _getLockBox(remoteChainSelector).deposit(address(i_token), remoteChainSelector, amount);
  }

  /// @inheritdoc TokenPool
  /// @dev Withdraws from the lockbox configured for the remote chain selector after a liquidity check.
  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal override {
    ILockBox lockBox = _getLockBox(remoteChainSelector);
    uint256 availableLiquidity = i_token.balanceOf(address(lockBox));
    if (amount > availableLiquidity) revert InsufficientLiquidity(availableLiquidity, amount);

    lockBox.withdraw(address(i_token), remoteChainSelector, amount, receiver);
  }

  /// @notice Returns the amount of tokens in the lock box for a given remote chain selector.
  /// @param remoteChainSelector the remote chain selector to get the lock box for.
  /// @return lockedTokens The tokens locked into the lock box for the given selector.
  function getAvailableTokens(
    uint64 remoteChainSelector
  ) external view returns (uint256) {
    return i_token.balanceOf(address(_getLockBox(remoteChainSelector)));
  }

  /// @notice No-op override to purge the unused code path from the contract.
  function _postFlightCheck(
    Pool.ReleaseOrMintInV1 calldata,
    uint256,
    uint16
  ) internal pure virtual override {}

  /// @notice No-op override to purge the unused code path from the contract.
  function _preFlightCheck(
    Pool.LockOrBurnInV1 calldata,
    uint16,
    bytes memory
  ) internal pure virtual override {}

  /// @notice Configure lockboxes.
  /// @param lockBoxConfigs The lockbox configurations to set.
  function configureLockBoxes(
    LockBoxConfig[] calldata lockBoxConfigs
  ) external onlyOwner {
    for (uint256 i = 0; i < lockBoxConfigs.length; ++i) {
      address lockBox = lockBoxConfigs[i].lockBox;
      if (lockBox == address(0)) revert ZeroAddressInvalid();
      ILockBox lockBoxContract = ILockBox(lockBox);
      if (!lockBoxContract.isTokenSupported(address(i_token))) {
        revert InvalidToken(address(i_token));
      }
      s_lockBoxes[lockBoxConfigs[i].remoteChainSelector] = lockBoxContract;
      i_token.approve(lockBox, type(uint256).max);
    }
  }

  /// @notice Gets the lockbox for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector to get the lockbox for.
  function _getLockBox(
    uint64 remoteChainSelector
  ) internal view returns (ILockBox) {
    ILockBox lockBox = s_lockBoxes[remoteChainSelector];
    if (address(lockBox) == address(0)) revert LockBoxNotConfigured(remoteChainSelector);
    return lockBox;
  }
}
