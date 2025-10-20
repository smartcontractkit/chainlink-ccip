// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {FeeQuoterHelper} from "../helpers/FeeQuoterHelper.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_constructor is FeeQuoterSetup {
  function test_constructor() public virtual {
    address[] memory priceUpdaters = new address[](2);
    priceUpdaters[0] = STRANGER;
    priceUpdaters[1] = OWNER;

    FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = _generateFeeQuoterDestChainConfigArgs();

    FeeQuoter.StaticConfig memory staticConfig =
      FeeQuoter.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: MAX_MSG_FEES_JUELS});
    s_feeQuoter = new FeeQuoterHelper(
      staticConfig,
      priceUpdaters,
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      s_feeQuoterTokenTransferFeeConfigArgs,
      destChainConfigArgs
    );

    assertEq(s_feeQuoter.getStaticConfig().linkToken, staticConfig.linkToken);
    assertEq(s_feeQuoter.getStaticConfig().maxFeeJuelsPerMsg, staticConfig.maxFeeJuelsPerMsg);

    assertEq(priceUpdaters, s_feeQuoter.getAllAuthorizedCallers());

    _assertTokenPriceFeedConfigEquality(
      tokenPriceFeedUpdates[0].feedConfig, s_feeQuoter.getTokenPriceFeedConfig(s_sourceTokens[0])
    );

    _assertTokenPriceFeedConfigEquality(
      tokenPriceFeedUpdates[1].feedConfig, s_feeQuoter.getTokenPriceFeedConfig(s_sourceTokens[1])
    );

    assertEq(
      s_feeQuoterPremiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth,
      s_feeQuoter.getPremiumMultiplierWeiPerEth(s_feeQuoterPremiumMultiplierWeiPerEthArgs[0].token)
    );

    assertEq(
      s_feeQuoterPremiumMultiplierWeiPerEthArgs[1].premiumMultiplierWeiPerEth,
      s_feeQuoter.getPremiumMultiplierWeiPerEth(s_feeQuoterPremiumMultiplierWeiPerEthArgs[1].token)
    );

    FeeQuoter.TokenTransferFeeConfigArgs memory tokenTransferFeeConfigArg = s_feeQuoterTokenTransferFeeConfigArgs[0];
    for (uint256 i = 0; i < tokenTransferFeeConfigArg.tokenTransferFeeConfigs.length; ++i) {
      FeeQuoter.TokenTransferFeeConfigSingleTokenArgs memory tokenFeeArgs =
        s_feeQuoterTokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[i];

      _assertTokenTransferFeeConfigEqual(
        tokenFeeArgs.tokenTransferFeeConfig,
        s_feeQuoter.getTokenTransferFeeConfig(tokenTransferFeeConfigArg.destChainSelector, tokenFeeArgs.token)
      );
    }

    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      FeeQuoter.DestChainConfig memory expectedConfig = destChainConfigArgs[i].destChainConfig;
      uint64 destChainSelector = destChainConfigArgs[i].destChainSelector;

      _assertFeeQuoterDestChainConfigsEqual(expectedConfig, s_feeQuoter.getDestChainConfig(destChainSelector));
    }
  }

  function test_constructor_RevertWhen_InvalidLinkTokenEqZeroAddress() public {
    vm.expectRevert(FeeQuoter.InvalidStaticConfig.selector);

    s_feeQuoter = new FeeQuoterHelper(
      FeeQuoter.StaticConfig({linkToken: address(0), maxFeeJuelsPerMsg: MAX_MSG_FEES_JUELS}),
      new address[](0),
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      s_feeQuoterTokenTransferFeeConfigArgs,
      new FeeQuoter.DestChainConfigArgs[](0)
    );
  }

  function test_constructor_RevertWhen_InvalidMaxFeeJuelsPerMsg() public {
    vm.expectRevert(FeeQuoter.InvalidStaticConfig.selector);

    s_feeQuoter = new FeeQuoterHelper(
      FeeQuoter.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: 0}),
      new address[](0),
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      s_feeQuoterTokenTransferFeeConfigArgs,
      new FeeQuoter.DestChainConfigArgs[](0)
    );
  }
}
