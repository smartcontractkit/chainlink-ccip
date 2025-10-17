// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";

contract ReentrantCCV is ICrossChainVerifierV1 {
  OffRamp internal immutable i_offRamp;

  constructor(
    address offRamp
  ) {
    i_offRamp = OffRamp(offRamp);
  }

  function forwardToVerifier(
    address,
    MessageV1Codec.MessageV1 calldata,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external pure returns (bytes memory) {
    return "";
  }

  function getFee(
    address, // originalSender
    uint64, // destChainSelector
    Client.EVM2AnyMessage memory, // message
    bytes memory, // extraArgs
    uint16 // finalityConfig
  ) external pure returns (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) {
    return (0, 0, 0);
  }

  function verifyMessage(
    address, // originalCaller
    MessageV1Codec.MessageV1 memory message,
    bytes32, // messageHash
    bytes memory ccvData
  ) external override {
    // Create a dummy report to trigger reentrancy.
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(this);
    bytes[] memory ccvDataArray = new bytes[](1);
    ccvDataArray[0] = ccvData;

    // This should trigger the reentrancy guard.
    i_offRamp.execute(MessageV1Codec._encodeMessageV1(message), ccvs, ccvDataArray);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(ICrossChainVerifierV1).interfaceId;
  }

  function getStorageLocation() external pure override returns (string memory) {
    return "reentrant://ccv";
  }
}
