// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoterV2} from "./interfaces/IFeeQuoterV2.sol";

import {FeeQuoter} from "./FeeQuoter.sol";
import {Internal} from "./libraries/Internal.sol";
import {Pool} from "./libraries/Pool.sol";

contract FeeQuoterV2 is FeeQuoter {
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
    Internal.EVMTokenTransfer[] calldata onRampTokenTransfers
  ) external view returns (bytes[] memory destExecDataPerToken) {
    bytes4 chainFamilySelector = s_destChainConfigs[destChainSelector].chainFamilySelector;
    destExecDataPerToken = new bytes[](onRampTokenTransfers.length);
    for (uint256 i = 0; i < onRampTokenTransfers.length; ++i) {
      Internal.EVMTokenTransfer memory tokenTransfer = onRampTokenTransfers[i];

      // Since the DON has to pay for the extraData to be included on the destination chain, we cap the length of the
      // extraData. This prevents gas bomb attacks on the NOPs. As destBytesOverhead accounts for both.
      // extraData and offchainData, this caps the worst case abuse to the number of bytes reserved for offchainData.
      uint256 destPoolDataLength = tokenTransfer.extraData.length;
      if (destPoolDataLength > Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
        if (
          destPoolDataLength
            > s_tokenTransferFeeConfig[destChainSelector][tokenTransfer.sourceTokenAddress].destBytesOverhead
        ) {
          revert SourceTokenDataTooLarge(tokenTransfer.sourceTokenAddress);
        }
      }

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
      destExecDataPerToken[i] = abi.encode(destGasAmount);
    }
    return destExecDataPerToken;
  }
}
