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
    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_sourceTokens[0];
    feeTokens[1] = s_sourceTokens[1];

    FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = _generateFeeQuoterDestChainConfigArgs();

    FeeQuoter.StaticConfig memory staticConfig =
      FeeQuoter.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: MAX_MSG_FEES_JUELS});
    s_feeQuoter = new FeeQuoterHelper(
      staticConfig,
      priceUpdaters,
      feeTokens,
      s_feeQuoterTokenTransferFeeConfigArgs,
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      destChainConfigArgs
    );

    assertEq(s_feeQuoter.getStaticConfig().linkToken, staticConfig.linkToken);
    assertEq(s_feeQuoter.getStaticConfig().maxFeeJuelsPerMsg, staticConfig.maxFeeJuelsPerMsg);

    assertEq(feeTokens, s_feeQuoter.getFeeTokens());
    assertEq(priceUpdaters, s_feeQuoter.getAllAuthorizedCallers());

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
      new address[](0),
      s_feeQuoterTokenTransferFeeConfigArgs,
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      new FeeQuoter.DestChainConfigArgs[](0)
    );
  }

  function test_constructor_RevertWhen_InvalidMaxFeeJuelsPerMsg() public {
    vm.expectRevert(FeeQuoter.InvalidStaticConfig.selector);

    s_feeQuoter = new FeeQuoterHelper(
      FeeQuoter.StaticConfig({linkToken: s_sourceTokens[0], maxFeeJuelsPerMsg: 0}),
      new address[](0),
      new address[](0),
      s_feeQuoterTokenTransferFeeConfigArgs,
      s_feeQuoterPremiumMultiplierWeiPerEthArgs,
      new FeeQuoter.DestChainConfigArgs[](0)
    );
  }
}
