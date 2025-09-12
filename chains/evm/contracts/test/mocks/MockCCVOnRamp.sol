// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICCVOnRampV1} from "../../interfaces/ICCVOnRampV1.sol";

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

contract MockCCVOnRamp is ICCVOnRampV1 {
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external pure returns (bytes memory) {
    return "";
  }

  function getFee(
    uint64, // destChainSelector,
    Client.EVM2AnyMessage memory, // message,
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }
}
