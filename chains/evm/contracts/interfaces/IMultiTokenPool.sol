// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IPoolV1V2} from "./IPoolV1V2.sol";

/// @notice Shared interface for CCIP pools that serve multiple tokens.
/// @dev The inherited isSupportedChain(uint64) and getRemoteToken(uint64) views revert with LocalTokenRequired.
/// Use the overloads that take a local token.
interface IMultiTokenPool is IPoolV1V2 {
  error LocalTokenRequired();

  /// @notice Returns whether the pool supports the token and the token has a lane to the remote chain.
  /// @param localToken The local token.
  /// @param remoteChainSelector The selector of the remote chain.
  /// @return supported True if the pool supports the token and a lane exists.
  function isSupportedChain(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bool supported);

  /// @notice Returns the remote token from the resolved lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The selector of the remote chain.
  /// @return remoteToken The encoded remote token address, or empty bytes if no lane exists.
  function getRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory remoteToken);

  /// @notice Returns one page of registered tokens assigned to this pool by the TokenAdminRegistry.
  /// @dev Pages index the pool config manager's token list before filtering by pool assignment. An empty result does
  /// not mark the end. For maxCount > 0, nextIndex - startIndex < maxCount marks the end. Read pages at the same block
  /// for consistent results.
  /// @param startIndex The starting index in the pool config manager's token list.
  /// @param maxCount The maximum number of registrations to read.
  /// @return tokens The tokens assigned to this pool within the page.
  /// @return nextIndex The index at which to continue reading the pool config manager's token list.
  function getTokens(
    uint64 startIndex,
    uint64 maxCount
  ) external view returns (address[] memory tokens, uint64 nextIndex);

  /// @notice Returns the token decimals stored by the pool config manager.
  /// @param localToken The local token.
  /// @return decimals The configured decimals, or zero if the local token is not registered.
  function getTokenDecimals(
    address localToken
  ) external view returns (uint8 decimals);
}
