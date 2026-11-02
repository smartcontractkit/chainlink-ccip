// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.20;

import {Pool} from "../libraries/Pool.sol";

/// @notice Validation hooks for transfers through multi-token pools.
interface IMultiTokenPoolHooks {
  /// @notice Preflight check before the lock or burn operation.
  /// @param lockOrBurnIn The lock or burn input parameters.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param tokenArgs Additional token arguments.
  /// @param amountPostFee The amount after the token fee.
  /// @dev A revert stops the transfer on the source chain.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    uint256 amountPostFee
  ) external;

  /// @notice Postflight check before the release or mint operation.
  /// @param releaseOrMintIn The release or mint input parameters.
  /// @param localAmount The amount in local decimals that the pool releases or mints.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @dev A revert makes the message fail. It can be executed again after the cause is fixed.
  function postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    bytes4 requestedFinalityConfig
  ) external;
}
