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
  /// @param tokenArgs Additional token arguments.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes calldata tokenArgs
  ) external returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut);

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param localToken The address of the local token.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token arguments.
  /// @return requiredCCVs A set of addresses representing the required outbound CCVs.
  function getRequiredOutboundCCVs(
    address localToken,
    uint64 destChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns the set of required CCVs for incoming messages from a source chain.
  /// @param localToken The address of the local token.
  /// @param sourceChainSelector The chain selector of the source chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param finality The finality configuration from the CCIP message.
  /// @param sourcePoolData The data received from the source pool to process the release or mint.
  /// @return requiredCCVs A set of addresses representing the required inbound CCVs.
  function getRequiredInboundCCVs(
    address localToken,
    uint64 sourceChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata sourcePoolData
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns a fee quote for transferring tokens to a destination chain.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param message The message to be sent to the destination chain.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token argument from the CCIP message.
  /// @return feeTokenAmount The amount of fee token needed for the fee.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (uint256 feeTokenAmount);
}
