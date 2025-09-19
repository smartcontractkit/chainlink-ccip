// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICCVRampV1} from "../../interfaces/ICCVRampV1.sol";

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

contract MockCCVRamp is ICCVRampV1 {
  bytes private s_verifierResult;

  constructor(
    bytes memory verifierResult
  ) {
    s_verifierResult = verifierResult;
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICCVRampV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  function forwardToVerifier(
    address,
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external view returns (bytes memory) {
    return s_verifierResult;
  }

  function getFee(
    address, // originalSender
    uint64, // destChainSelector
    Client.EVM2AnyMessage memory, // message
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    return 0;
  }

  function verifyMessage(
    address, // originalCaller
    MessageV1Codec.MessageV1 memory, // message
    bytes32 messageId, // messageId
    bytes memory ccvData // ccvData
  ) external {}
}
