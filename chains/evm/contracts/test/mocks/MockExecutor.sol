// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutor} from "../../interfaces/IExecutor.sol";

import {Client} from "../../libraries/Client.sol";

contract MockExecutor is IExecutor {
  function getMinBlockConfirmations() external view virtual returns (uint16) {
    return 0;
  }

  function getFee(
    uint64, // destChainSelector,
    uint16, // requestedBlockDepth,
    Client.CCV[] memory, // ccvs,
    bytes memory // extraArgs
  ) external pure returns (uint16 usdCents) {
    return 0;
  }
}
