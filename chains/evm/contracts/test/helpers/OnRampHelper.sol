// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";

contract OnRampHelper is OnRamp {
  constructor(StaticConfig memory staticConfig, DynamicConfig memory dynamicConfig) OnRamp(staticConfig, dynamicConfig) {}

  /// @notice Exposes the internal _parseExtraArgsWithDefaults function for testing.
  function parseExtraArgsWithDefaults(
    uint64 destChainSelector,
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs
  ) external view returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return _parseExtraArgsWithDefaults(destChainSelector, destChainConfig, extraArgs);
  }

  /// @notice Exposes the internal _mergeCCVLists function for testing.
  function mergeCCVLists(
    address[] memory userRequestedOrDefaultCCVs,
    bytes[] memory userRequestedOrDefaultCCVArgs,
    address[] memory laneMandatedCCVs,
    address[] memory poolRequiredCCVs
  ) external pure returns (address[] memory ccvs, bytes[] memory ccvArgs) {
    return _mergeCCVLists(userRequestedOrDefaultCCVs, userRequestedOrDefaultCCVArgs, laneMandatedCCVs, poolRequiredCCVs);
  }

  function getCCVsForPool(
    uint64 destChainSelector,
    address token,
    uint256 amount,
    uint16 finality,
    bytes memory tokenArgs
  ) external view returns (address[] memory) {
    return _getCCVsForPool(destChainSelector, token, amount, finality, tokenArgs);
  }

  /// @notice Exposes the internal _getExecutionFee function for testing.
  function getExecutionFee(
    uint64 destChainSelector,
    uint256 dataLength,
    uint256 numberOfTokens,
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) external view returns (Receipt memory) {
    return _getExecutionFee(destChainSelector, dataLength, numberOfTokens, extraArgs);
  }

  function lockOrBurnSingleToken(
    Client.EVMTokenAmount memory tokenAndAmount,
    uint64 destChainSelector,
    bytes memory receiver,
    address originalSender,
    uint16 finality,
    bytes memory tokenArgs
  ) external returns (MessageV1Codec.TokenTransferV1 memory) {
    return _lockOrBurnSingleToken(tokenAndAmount, destChainSelector, receiver, originalSender, finality, tokenArgs);
  }

  function getReceipts(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) external view returns (Receipt[] memory verifierReceipts, uint32 gasLimitSum, uint256 feeUSDCentsSum) {
    return _getReceipts(destChainSelector, message, extraArgs);
  }

  function distributeFees(Client.EVM2AnyMessage calldata message, Receipt[] memory receipts) external {
    _distributeFees(message, receipts);
  }
}
