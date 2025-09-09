// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICCVOnRamp} from "../../interfaces/ICCVOnRamp.sol";
import {Client} from "../../libraries/Client.sol";

contract MockCCVOnRamp is ICCVOnRamp {
  function forwardToVerifier(uint64, bytes32, address, bytes memory, uint256) external pure returns (bytes memory) {
    return "";
  }

  function getFee(
    uint64, // destChainSelector,
    bytes32, // version,
    address, // caller,
    Client.EVM2AnyMessage memory, // message,
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }
}
