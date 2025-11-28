// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../libraries/Pool.sol";
import {IPoolV2} from "./IPoolV2.sol";

/// @notice Interface for AdvancedPoolHooks contract. Implementations may contain no-op logic.
interface IAdvancedPoolHooks {
  /// @notice Preflight check before lock or burn operation.
  /// @param lockOrBurnIn The lock or burn input parameters.
  /// @param blockConfirmationRequested The block confirmation requested.
  /// @param tokenArgs Additional token arguments.
  /// @dev This function may revert if the preflight check fails. This means the transaction is rolled back on source.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external;

  /// @notice Postflight check before releasing or minting tokens.
  /// @param releaseOrMintOut The release or mint output parameters.
  /// @param localAmount The local amount to be released or minted.
  /// @param blockConfirmationRequested The block confirmation requested.
  /// @dev This function may revert if the postflight check fails. This means the transaction is unexecutable until
  /// the issue is resolved.
  function postFlightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintOut,
    uint256 localAmount,
    uint16 blockConfirmationRequested
  ) external;

  /// @notice Returns the set of required CCVs for transfers in a specific direction.
  /// @param remoteChainSelector The remote chain selector for this transfer.
  /// @param amount The amount being transferred.
  /// @param direction The direction of the transfer (Inbound or Outbound).
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16 blockConfirmationRequested,
    bytes calldata extraData,
    IPoolV2.MessageDirection direction
  ) external view returns (address[] memory requiredCCVs);
}
