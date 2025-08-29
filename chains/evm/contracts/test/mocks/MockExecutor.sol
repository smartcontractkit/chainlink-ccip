// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IExecutorOnRamp} from "../../interfaces/IExecutorOnRamp.sol";

import {Client} from "../../libraries/Client.sol";

contract MockExecutor is IExecutorOnRamp {
  function getFee(
    uint64, // destChainSelector,
    Client.EVM2AnyMessage memory, // message,
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }
}
