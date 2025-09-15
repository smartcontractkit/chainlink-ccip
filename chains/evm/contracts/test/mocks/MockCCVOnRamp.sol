// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICCVOnRamp} from "../../interfaces/ICCVOnRamp.sol";

import {Client} from "../../libraries/Client.sol";

contract MockCCVOnRamp is ICCVOnRamp {
  bytes internal s_returnData;

  constructor(
    bytes memory returnData
  ) {
    s_returnData = returnData;
  }

  function forwardToVerifier(uint64, address, bytes memory, uint256) external view returns (bytes memory) {
    return s_returnData;
  }

  function getFee(
    uint64, // destChainSelector,
    address, // originalCaller,
    Client.EVM2AnyMessage memory, // message,
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }
}
