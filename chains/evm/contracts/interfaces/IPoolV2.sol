// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPoolV1} from "./IPool.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";

// TODO Milestone 2: implement.
/// @notice Shared public interface for multiple V2 pool types.
/// Each pool type handles a different child token model e.g. lock/unlock, mint/burn.
interface IPoolV2 is IPoolV1 {
  /// @notice Lock tokens into the pool or burn the tokens.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes calldata tokenExtraData
  ) external returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut);

  // TODO add new methods here for V2. Everything below is a placeholder.

  /// @notice Returns the list of required outbound CCVs (CCV OnRamps) for a given destination chain and amount.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param tokenArgs Additional token arguments.
  /// @return An array of addresses representing the required outbound CCVs.
  function getRequiredOutboundCCVs(
    uint64 destChainSelector,
    uint256 amount,
    bytes calldata tokenArgs
  ) external view returns (address[] memory);

  /// @notice Returns the list of required inbound CCVs (CCV OffRamps) for a given source chain and amount.
  /// @param sourceChainSelector The chain selector of the source chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param tokenArgs Additional token arguments.
  /// @return An array of addresses representing the required inbound CCVs.
  function getRequiredInboundCCVs(
    uint64 sourceChainSelector,
    uint256 amount,
    bytes calldata tokenArgs
  ) external view returns (address[] memory);

  /// @notice Returns a fee quote for transferring tokens to a destination chain.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param sender The address of the sender on the source chain.
  /// @param feeToken The address of the token to be used for fee payment.
  /// @param tokenAmounts An array of token amounts to be transferred.
  /// @param tokenArgs Additional token arguments.
  /// @return A Pool.Quote struct containing the fee breakdown.
  function getFee(
    uint64 destChainSelector,
    address sender,
    address feeToken,
    Client.EVMTokenAmount[] calldata tokenAmounts,
    bytes calldata tokenArgs
  ) external view returns (Pool.Quote memory);
}
