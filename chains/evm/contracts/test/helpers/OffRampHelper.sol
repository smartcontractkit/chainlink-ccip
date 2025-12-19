// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

import {OffRamp} from "../../offRamp/OffRamp.sol";

contract OffRampHelper is OffRamp {
  constructor(
    OffRamp.StaticConfig memory staticConfig
  ) OffRamp(staticConfig) {}

  function ensureCCVQuorumIsReached(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality,
    address[] calldata ccvs,
    bool isTokenOnlyTransfer
  ) external view returns (address[] memory, uint256[] memory) {
    return _ensureCCVQuorumIsReached(sourceChainSelector, receiver, tokenTransfer, finality, ccvs, isTokenOnlyTransfer);
  }

  function getCCVsFromReceiver(
    uint64 sourceChainSelector,
    address receiver
  ) external view returns (address[] memory, address[] memory, uint8) {
    return _getCCVsFromReceiver(sourceChainSelector, receiver);
  }

  function getCCVsFromPool(
    address localToken,
    uint64 sourceChainSelector,
    uint256 amount,
    uint16 finality,
    bytes memory extraData
  ) external view returns (address[] memory) {
    return _getCCVsFromPool(localToken, sourceChainSelector, amount, finality, extraData);
  }

  function beforeExecuteSingleMessage(
    MessageV1Codec.MessageV1 memory message
  ) external returns (MessageV1Codec.MessageV1 memory) {
    return _beforeExecuteSingleMessage(message);
  }

  function releaseOrMintSingleToken(
    MessageV1Codec.TokenTransferV1 memory tokenTransfer,
    bytes memory originalSender,
    uint64 sourceChainSelector,
    uint16 finality
  ) external returns (Client.EVMTokenAmount memory, address) {
    return _releaseOrMintSingleToken(tokenTransfer, originalSender, sourceChainSelector, finality);
  }

  function getBalanceOfReceiver(address receiver, address token) external view returns (uint256) {
    return _getBalanceOfReceiver(receiver, token);
  }

  function __getCCVsForMessage(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality,
    bool isTokenOnlyTransfer
  ) external view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold) {
    return _getCCVsForMessage(sourceChainSelector, receiver, tokenTransfer, finality, isTokenOnlyTransfer);
  }

  function checkIsTokenOnlyTransfer(
    uint256 dataLength,
    uint256 ccipReceiveGasLimit,
    address receiver
  ) external view returns (bool) {
    return _isTokenOnlyTransfer(dataLength, ccipReceiveGasLimit, receiver);
  }
}
