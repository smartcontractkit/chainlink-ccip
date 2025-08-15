// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoterV2} from "./interfaces/IFeeQuoterV2.sol";

import {FeeQuoter} from "./FeeQuoter.sol";
import {Client} from "./libraries/Client.sol";
import {Internal} from "./libraries/Internal.sol";

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

  function processPoolReturnDataNew(
    uint64 destChainSelector,
    Internal.EVMTokenTransfer calldata tokenTransfer
  ) external view returns (bytes memory destExecDataPerToken) {
    bytes4 chainFamilySelector = s_destChainConfigs[destChainSelector].chainFamilySelector;

    // We pass '1' here so that SVM validation requires a non-zero token address.
    // The 'gasLimit' parameter isn't actually used for gas in this context; it simply
    // signals that the address must not be zero on SVM.
    _validateDestFamilyAddress(chainFamilySelector, tokenTransfer.destTokenAddress, 1);
    TokenTransferFeeConfig memory tokenTransferFeeConfig =
      s_tokenTransferFeeConfig[destChainSelector][tokenTransfer.sourceTokenAddress];

    uint32 destGasAmount = tokenTransferFeeConfig.isEnabled
      ? tokenTransferFeeConfig.destGasOverhead
      : s_destChainConfigs[destChainSelector].defaultTokenDestGasOverhead;

    // The user will be billed either the default or the override, so we send the exact amount that we billed for
    // to the destination chain to be used for the token releaseOrMint and transfer.
    return abi.encode(destGasAmount);
  }

  struct ModSecExtraArgs {
    uint256 gasLimit;
    bytes tokenReceiver;
    bytes destChainExtraArgs;
    bytes[] verifierExtraArgs;
  }

  function resolveTokenReceiver(
    bytes calldata extraArgs
  ) external pure returns (bytes memory tokenReceiver) {
    if (bytes4(extraArgs[:4]) != Client.SVM_EXTRA_ARGS_V1_TAG) {
      return (bytes(""));
    }

    return abi.encode(abi.decode(extraArgs[4:], (Client.SVMExtraArgsV1)).tokenReceiver);
  }
}
