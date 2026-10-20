// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IGetPoolConfigManager} from "./IGetPoolConfigManager.sol";
import {IPoolV1V2} from "./IPoolV1V2.sol";

/// @notice Shared interface for CCIP pools that serve multiple tokens.
/// @dev The inherited isSupportedChain(uint64) and getRemoteToken(uint64) views revert with LocalTokenRequired.
/// Use the overloads that take local tokens.
interface IMultiTokenPool is IPoolV1V2, IGetPoolConfigManager {
  error LocalTokenRequired();

  /// @notice Returns whether the pool supports the token and the token has a lane to the remote chain.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The selector of the remote chain.
  /// @return supported True if the pool supports the token and a lane exists, aligned with localTokens by index.
  function isSupportedChain(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (bool[] memory supported);

  /// @notice Returns the remote token from the resolved lane.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The selector of the remote chain.
  /// @return remoteTokens The encoded remote token addresses, or empty bytes if no lane exists, aligned with
  /// localTokens by index.
  function getRemoteToken(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (bytes[] memory remoteTokens);

  /// @notice Returns the token decimals stored by the pool config manager.
  /// @param localTokens The local tokens.
  /// @return decimals The configured decimals, or zero if the local token is not registered, aligned with localTokens
  /// by index.
  function getTokenDecimals(
    address[] calldata localTokens
  ) external view returns (uint8[] memory decimals);
}
