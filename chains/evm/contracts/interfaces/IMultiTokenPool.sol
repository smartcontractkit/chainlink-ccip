// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IGetPoolConfigManager} from "./IGetPoolConfigManager.sol";
import {IPoolV1V2} from "./IPoolV1V2.sol";

/// @notice Shared interface for CCIP pools that serve multiple tokens.
/// @dev The inherited isSupportedChain(uint64) and getRemoteToken(uint64) views revert with LocalTokenRequired.
/// Use the overloads that take the local token.
interface IMultiTokenPool is IPoolV1V2, IGetPoolConfigManager {
  error LocalTokenRequired();

  /// @notice Returns whether the pool supports the token and the token has a lane to the remote chain.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  /// @return supported True if the pool supports the token and a lane exists.
  function isSupportedChain(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bool supported);

  /// @notice Returns the remote token from the resolved lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  /// @return remoteToken The encoded remote token, or empty bytes if no lane exists.
  function getRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory remoteToken);

  /// @notice Returns the decimals of the local token.
  /// @param localToken The local token.
  /// @return decimals The configured decimals, or zero if the token is not registered.
  function getTokenDecimals(
    address localToken
  ) external view returns (uint8 decimals);

  /// @notice Withdraws accrued fee token balances to the provided `recipient`.
  /// @dev Only callable by authorized callers. Reverts if the recipient is the zero address.
  /// @param feeTokens The token addresses to withdraw, including the pool token when applicable.
  /// @param recipient The address to withdraw the fee tokens to.
  function withdrawFeeTokens(
    address[] calldata feeTokens,
    address recipient
  ) external;
}
