// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IExecutor} from "../../interfaces/IExecutor.sol";

import {Client} from "../../libraries/Client.sol";

contract MockExecutor is IExecutor {
  function getFee(
    address, // originalCaller
    uint64, // destChainSelector,
    Client.EVM2AnyMessage memory, // message,
    Client.CCV[] memory, // ccvs,
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }
}
