// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";

contract ReentrantCCV is ICrossChainVerifierV1, ICrossChainVerifierResolver {
  OffRamp internal immutable i_offRamp;

  constructor(
    address offRamp
  ) {
    i_offRamp = OffRamp(offRamp);
  }

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
    uint64, // destChainSelector
    Client.EVM2AnyMessage memory, // message
    bytes memory, // extraArgs
    uint16 // blockConfirmations
  ) external pure returns (uint16 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) {
    return (0, 0, 0);
  }

  function verifyMessage(
    MessageV1Codec.MessageV1 memory message,
    bytes32, // messageHash
    bytes memory verifierResults
  ) external override {
    // Create a dummy report to trigger reentrancy.
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(this);
    bytes[] memory verifierResultsArray = new bytes[](1);
    verifierResultsArray[0] = verifierResults;

    // This should trigger the reentrancy guard.
    i_offRamp.execute(MessageV1Codec._encodeMessageV1(message), ccvs, verifierResultsArray);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(ICrossChainVerifierV1).interfaceId;
  }

  function getStorageLocation() external pure override returns (string memory) {
    return "reentrant://ccv";
  }

  function getInboundImplementation(
    bytes calldata // verifierResults
  ) external view returns (address) {
    return address(this);
  }

  function getOutboundImplementation(
    uint64, // destChainSelector
    bytes memory // extraArgs
  ) external view returns (address) {
    return address(this);
  }
}
