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
    Pool.LockOrBurnV2 calldata lockOrBurnIn
  ) external returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut);

  // TODO add new methods here for V2. Everything below is a placeholder.
  function getRequiredCCVs(
    address token,
    uint64 sourceChainSelector,
    uint256 amount,
    bytes memory extraData
  ) external view returns (address[] memory requiredCCVs);

  function getRequiredOutboundCCVs(
    uint64 destChainSelector,
    uint256 amount,
    bytes calldata tokenArgs
  ) external view returns (address[] memory);

  function getRequiredInboundCCVs(
    uint64 sourceChainSelector,
    uint256 amount,
    bytes calldata tokenArgs
  ) external view returns (address[] memory);

  function getFee(
    uint64 destChainSelector,
    address sender,
    address feeToken,
    Client.EVMTokenAmount[] calldata tokenAmounts,
    bytes calldata tokenArgs
  ) external view returns (Pool.Quote memory);
}
