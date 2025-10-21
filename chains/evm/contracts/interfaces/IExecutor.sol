// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface IExecutor {
  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  function getFee(
    uint64 destChainSelector,
    uint16 requestedBlockDepth,
    uint32 dataLength,
    uint8 numberOfTokens,
    Client.CCV[] memory ccvs,
    bytes memory extraArgs
  ) external view returns (uint16 usdCents, uint32 gasLimit, uint32 destBytesOverhead);
}
