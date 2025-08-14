// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";

contract MockExecutor {
  function forwardToVerifier(bytes memory, uint256) external pure returns (bytes memory) {
    return "";
  }

  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256) {
    return 0;
  }
}
