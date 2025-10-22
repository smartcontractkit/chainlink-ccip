// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract MockVerifier is ICrossChainVerifierV1 {
  bytes private s_verifierResult;

  constructor(
    bytes memory verifierResult
  ) {
    s_verifierResult = verifierResult;
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICrossChainVerifierV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external view returns (bytes memory) {
    return s_verifierResult;
  }

  function getFee(
    uint64, // destChainSelector
    Client.EVM2AnyMessage memory, // message
    bytes memory, // extraArgs
    uint16 // finalityConfig
  ) external pure returns (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) {
    return (0, 0, 0);
  }

  function verifyMessage(
    MessageV1Codec.MessageV1 memory, // message
    bytes32 messageId, // messageId
    bytes memory ccvData // ccvData
  ) external {}

  function getStorageLocation() external pure override returns (string memory) {
    return "mock://ccv";
  }
}
