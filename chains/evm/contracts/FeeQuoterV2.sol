// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoterV2} from "./interfaces/IFeeQuoterV2.sol";

import {FeeQuoter} from "./FeeQuoter.sol";
import {Client} from "./libraries/Client.sol";

contract FeeQuoterV2 is IFeeQuoterV2, FeeQuoter {
  constructor(
    StaticConfig memory staticConfig,
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    TokenPriceFeedUpdate[] memory tokenPriceFeeds,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs,
    DestChainConfigArgs[] memory destChainConfigArgs
  )
    FeeQuoter(
      staticConfig,
      priceUpdaters,
      feeTokens,
      tokenPriceFeeds,
      tokenTransferFeeConfigArgs,
      premiumMultiplierWeiPerEthArgs,
      destChainConfigArgs
    )
  {}

  function resolveTokenReceiver(
    bytes calldata extraArgs
  ) external pure returns (bytes memory tokenReceiver) {
    if (extraArgs.length < 4 || bytes4(extraArgs[:4]) != Client.SVM_EXTRA_ARGS_V1_TAG) {
      return (bytes(""));
    }

    return abi.encode(abi.decode(extraArgs[4:], (Client.SVMExtraArgsV1)).tokenReceiver);
  }
}
