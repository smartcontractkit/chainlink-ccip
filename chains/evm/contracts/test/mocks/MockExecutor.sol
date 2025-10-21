// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IExecutor} from "../../interfaces/IExecutor.sol";

import {Client} from "../../libraries/Client.sol";

contract MockExecutor is IExecutor {
  function getFee(
    uint64, // destChainSelector,
    uint16, // requestedBlockDepth,
    uint32, // dataLength,
    uint8, // numberOfTokens,
    Client.CCV[] memory, // ccvs,
    bytes memory // extraArgs
  ) external pure returns (uint16 usdCents, uint32 gasLimit, uint32 destBytesOverhead) {
    return (0, 0, 0);
  }
}
