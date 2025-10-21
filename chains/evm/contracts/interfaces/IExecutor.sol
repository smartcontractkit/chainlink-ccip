// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface IExecutor {
  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  function getFee(
    address originalCaller,
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    Client.CCV[] memory ccvs,
    bytes memory extraArgs
  ) external view returns (uint256);
}
