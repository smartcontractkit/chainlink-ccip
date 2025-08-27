// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";

import {BaseTest} from "../../BaseTest.t.sol";

contract CCVAggregatorHelper is CCVAggregator {
  constructor(
    CCVAggregator.StaticConfig memory staticConfig
  ) CCVAggregator(staticConfig) {}

  // Expose internal functions for testing
  function setExecutionState(
    uint64 sourceChainSelector,
    uint64 sequenceNumber,
    Internal.MessageExecutionState newState
  ) external {
    _setExecutionState(sourceChainSelector, sequenceNumber, newState);
  }

  function getSequenceNumberBitmap(uint64 sourceChainSelector, uint64 sequenceNumber) external view returns (uint256) {
    return _getSequenceNumberBitmap(sourceChainSelector, sequenceNumber);
  }

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

contract CCVAggregatorSetup is BaseTest {
  CCVAggregatorHelper internal s_agg;
  address internal s_defaultCCV;
  address internal s_tokenAdminRegistry;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_defaultCCV = makeAddr("defaultCCV");
    s_tokenAdminRegistry = makeAddr("tokenAdminRegistry");

    s_agg = new CCVAggregatorHelper(
      CCVAggregator.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: s_tokenAdminRegistry
      })
    );

    // Apply initial source chain configuration
    _applySourceConfig(
      s_sourceRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), true, new address[](1), new address[](0)
    );
  }

  function _applySourceConfig(
    IRouter router,
    uint64 sourceChainSelector,
    bytes memory onRamp,
    bool isEnabled,
    address[] memory defaultCCVs,
    address[] memory laneMandatedCCVs
  ) internal {
    defaultCCVs[0] = s_defaultCCV;

    CCVAggregator.SourceChainConfigArgs[] memory updates = new CCVAggregator.SourceChainConfigArgs[](1);
    updates[0] = CCVAggregator.SourceChainConfigArgs({
      router: router,
      sourceChainSelector: sourceChainSelector,
      isEnabled: isEnabled,
      onRamp: onRamp,
      defaultCCV: defaultCCVs,
      laneMandatedCCVs: laneMandatedCCVs
    });
    s_agg.applySourceChainConfigUpdates(updates);
  }
}
