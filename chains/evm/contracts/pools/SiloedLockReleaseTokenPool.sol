// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../interfaces/ILockBox.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {EnumerableMap} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableMap.sol";

/// @notice A variation on Lock Release token pools where liquidity is shared among some chains, and stored independently
/// for others. Chains which do not share liquidity are known as siloed chains.
contract SiloedLockReleaseTokenPool is TokenPool, ITypeAndVersion {
  using EnumerableMap for EnumerableMap.UintToAddressMap;

  error InsufficientLiquidity(uint256 availableLiquidity, uint256 requestedAmount);
  error LockBoxNotConfigured(uint64 remoteChainSelector);

  struct LockBoxConfig {
    uint64 remoteChainSelector;
    address lockBox;
  }

  /// @notice Lock boxes keyed by chain selector.
  /// @dev We can have a many-to-one mapping of remote chain selectors to lock boxes. This allows for chains to share or isolate liquidity.
  EnumerableMap.UintToAddressMap internal s_lockBoxes;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {}

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
    getLockBox(remoteChainSelector).deposit(address(i_token), remoteChainSelector, amount);
  }

  /// @inheritdoc TokenPool
  /// @dev Withdraws from the lockbox configured for the remote chain selector after a liquidity check.
  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal override {
    ILockBox lockBox = getLockBox(remoteChainSelector);
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
    return i_token.balanceOf(address(getLockBox(remoteChainSelector)));
  }

  /// @notice Returns all configured lockboxes.
  /// @return lockBoxConfigs Array of all configured lockbox configurations.
  function getAllLockBoxConfigs() external view returns (LockBoxConfig[] memory lockBoxConfigs) {
    uint256 length = s_lockBoxes.length();
    lockBoxConfigs = new LockBoxConfig[](length);
    for (uint256 i = 0; i < length; ++i) {
      (uint256 chainSelector, address lockBox) = s_lockBoxes.at(i);
      lockBoxConfigs[i] = LockBoxConfig({remoteChainSelector: uint64(chainSelector), lockBox: lockBox});
    }
    return lockBoxConfigs;
  }

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
      s_lockBoxes.set(lockBoxConfigs[i].remoteChainSelector, lockBox);
      i_token.approve(lockBox, type(uint256).max);
    }
  }

  /// @notice Gets the lockbox for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector to get the lockbox for.
  function getLockBox(
    uint64 remoteChainSelector
  ) public view returns (ILockBox) {
    (bool exists, address lockBox) = s_lockBoxes.tryGet(remoteChainSelector);
    if (!exists) revert LockBoxNotConfigured(remoteChainSelector);
    return ILockBox(lockBox);
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
}
