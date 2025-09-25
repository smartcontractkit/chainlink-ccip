// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";

import {CCVAggregator} from "../../offRamp/CCVAggregator.sol";

contract CCVAggregatorHelper is CCVAggregator {
  constructor(
    CCVAggregator.StaticConfig memory staticConfig
  ) CCVAggregator(staticConfig) {}

  function ensureCCVQuorumIsReached(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality,
    address[] calldata ccvs
  ) external view returns (address[] memory, uint256[] memory) {
    return _ensureCCVQuorumIsReached(sourceChainSelector, receiver, tokenTransfer, finality, ccvs);
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
    MessageV1Codec.TokenTransferV1 memory sourceTokenAmount,
    bytes memory originalSender,
    address receiver,
    uint64 sourceChainSelector
  ) external returns (Client.EVMTokenAmount memory) {
    return _releaseOrMintSingleToken(sourceTokenAmount, originalSender, receiver, sourceChainSelector);
  }

  function getBalanceOfReceiver(address receiver, address token) external view returns (uint256) {
    return _getBalanceOfReceiver(receiver, token);
  }

  function __getCCVsForMessage(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality
  ) external view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold) {
    return _getCCVsForMessage(sourceChainSelector, receiver, tokenTransfer, finality);
  }
}
