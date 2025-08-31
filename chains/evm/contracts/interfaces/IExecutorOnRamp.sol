// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface IExecutorOnRamp {
  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    Client.CCV[] memory requiredCCVs,
    Client.CCV[] memory optionalCCVs,
    bytes memory extraArgs
  ) external view returns (uint256);
}
