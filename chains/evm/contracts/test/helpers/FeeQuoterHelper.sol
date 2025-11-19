// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Client} from "../../libraries/Client.sol";

contract FeeQuoterHelper is FeeQuoter {
  constructor(
    StaticConfig memory staticConfig,
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    DestChainConfigArgs[] memory destChainConfigArgs
  ) FeeQuoter(staticConfig, priceUpdaters, feeTokens, tokenTransferFeeConfigArgs, destChainConfigArgs) {}

  function getTokenTransferCost(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] calldata tokenAmounts
  ) external view returns (uint256, uint32, uint32) {
    return _getTokenTransferCost(
      s_destChainConfigs[destChainSelector].defaultTokenFeeUSDCents,
      s_destChainConfigs[destChainSelector].defaultTokenDestGasOverhead,
      destChainSelector,
      tokenAmounts
    );
  }

  function parseEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint64 destChainSelector
  ) external view returns (Client.GenericExtraArgsV2 memory) {
    return _parseGenericExtraArgsFromBytes(
      extraArgs,
      s_destChainConfigs[destChainSelector].defaultTxGasLimit,
      s_destChainConfigs[destChainSelector].maxPerMsgGasLimit
    );
  }

  function parseSVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    DestChainConfig memory destChainConfig
  ) external pure returns (Client.SVMExtraArgsV1 memory) {
    return _parseSVMExtraArgsFromBytes(extraArgs, destChainConfig.maxPerMsgGasLimit);
  }

  function parseSuiExtraArgsFromBytes(
    bytes calldata extraArgs,
    DestChainConfig memory destChainConfig
  ) external pure returns (Client.SuiExtraArgsV1 memory) {
    return _parseSuiExtraArgsFromBytes(extraArgs, destChainConfig.maxPerMsgGasLimit);
  }

  function processChainFamilySelector(
    uint64 chainFamilySelector,
    bytes calldata messageReceiver,
    bytes calldata extraArgs
  ) external view returns (bytes memory, bool, bytes memory) {
    return _processChainFamilySelector(chainFamilySelector, messageReceiver, extraArgs);
  }

  function validateDestFamilyAddress(
    bytes4 chainFamilySelector,
    bytes memory destAddress,
    uint256 gasLimit
  ) external pure {
    _validateDestFamilyAddress(chainFamilySelector, destAddress, gasLimit);
  }
}
