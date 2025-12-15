// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library CCTPHelper {
  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageHeader {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    bytes32 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageBody {
    uint32 version;
    bytes32 burnToken;
    bytes32 mintRecipient;
    uint256 amount;
    bytes32 messageSender;
    uint256 maxFee;
    uint256 feeExecuted;
    uint256 expirationBlock;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageHookData {
    bytes4 verifierVersion;
    bytes32 messageId;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessage {
    CCTPMessageHeader header;
    CCTPMessageBody body;
    CCTPMessageHookData hookData;
  }

  function _encodeCCTPMessage(
    CCTPMessage memory cctpMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      _encodeCCTPMessageHeader(cctpMessage.header),
      _encodeCCTPMessageBody(cctpMessage.body),
      _encodeCCTPMessageHookData(cctpMessage.hookData)
    );
  }

  function _encodeCCTPMessageHeader(
    CCTPMessageHeader memory header
  ) private pure returns (bytes memory) {
    return abi.encodePacked(
      header.version,
      header.sourceDomain,
      header.destinationDomain,
      header.nonce,
      header.sender,
      header.recipient,
      header.destinationCaller,
      header.minFinalityThreshold,
      header.finalityThresholdExecuted
    );
  }

  function _encodeCCTPMessageBody(
    CCTPMessageBody memory body
  ) private pure returns (bytes memory) {
    return abi.encodePacked(
      body.version,
      body.burnToken,
      body.mintRecipient,
      body.amount,
      body.messageSender,
      body.maxFee,
      body.feeExecuted,
      body.expirationBlock
    );
  }

  function _encodeCCTPMessageHookData(
    CCTPMessageHookData memory hookData
  ) private pure returns (bytes memory) {
    return abi.encodePacked(hookData.verifierVersion, hookData.messageId);
  }
}
