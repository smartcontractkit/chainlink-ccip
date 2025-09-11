// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";

import {CCVAggregator} from "../../offRamp/CCVAggregator.sol";

contract CCVAggregatorHelper is CCVAggregator {
  constructor(
    CCVAggregator.StaticConfig memory staticConfig
  ) CCVAggregator(staticConfig) {}

  function ensureCCVQuorumIsReached(
    uint64 sourceChainSelector,
    address receiver,
    address[] calldata ccvs,
    address[] memory tokenRequiredCCVs
  ) external view returns (address[] memory, uint256[] memory) {
    return _ensureCCVQuorumIsReached(sourceChainSelector, receiver, ccvs, tokenRequiredCCVs);
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
    bytes memory extraData
  ) external view returns (address[] memory) {
    return _getCCVsFromPool(localToken, sourceChainSelector, amount, extraData);
  }

  function trialExecute(
    Internal.Any2EVMMessage memory message
  ) external returns (Internal.MessageExecutionState, bytes memory) {
    return _trialExecute(message);
  }

  function beforeExecuteSingleMessage(
    Internal.Any2EVMMessage memory message
  ) external returns (Internal.Any2EVMMessage memory) {
    return _beforeExecuteSingleMessage(message);
  }

  function releaseOrMintSingleToken(
    Internal.TokenTransfer memory sourceTokenAmount,
    bytes memory originalSender,
    address receiver,
    uint64 sourceChainSelector
  ) external returns (Client.EVMTokenAmount memory) {
    return _releaseOrMintSingleToken(sourceTokenAmount, originalSender, receiver, sourceChainSelector);
  }

  function getBalanceOfReceiver(address receiver, address token) external view returns (uint256) {
    return _getBalanceOfReceiver(receiver, token);
  }
}
