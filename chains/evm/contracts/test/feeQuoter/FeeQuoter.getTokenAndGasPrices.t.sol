// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Internal} from "../../libraries/Internal.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_getTokenAndGasPrices is FeeQuoterSetup {
  function test_getTokenAndGasPrices() public view {
    (uint224 feeTokenPrice, uint224 gasPrice) = s_feeQuoter.getTokenAndGasPrices(s_sourceFeeToken, DEST_CHAIN_SELECTOR);

    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));

    assertEq(feeTokenPrice, s_sourceTokenPrices[0]);
    assertEq(gasPrice, priceUpdates.gasPriceUpdates[0].usdPerUnitGas);
  }

  function test_getTokenAndGasPrices_ZeroGasPrice() public {
    uint64 zeroGasDestChainSelector = 345678;
    FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = _generateFeeQuoterDestChainConfigArgs();
    destChainConfigArgs[0].destChainSelector = zeroGasDestChainSelector;

    s_feeQuoter.applyDestChainConfigUpdates(destChainConfigArgs);
    Internal.GasPriceUpdate[] memory gasPriceUpdates = new Internal.GasPriceUpdate[](1);
    gasPriceUpdates[0] = Internal.GasPriceUpdate({destChainSelector: zeroGasDestChainSelector, usdPerUnitGas: 0});

    Internal.PriceUpdates memory priceUpdates =
      Internal.PriceUpdates({tokenPriceUpdates: new Internal.TokenPriceUpdate[](0), gasPriceUpdates: gasPriceUpdates});
    s_feeQuoter.updatePrices(priceUpdates);

    (, uint224 gasPrice) = s_feeQuoter.getTokenAndGasPrices(s_sourceFeeToken, zeroGasDestChainSelector);

    assertEq(gasPrice, 0);
  }

  function test_getTokenAndGasPrices_RevertWhen_UnsupportedChain() public {
    vm.expectRevert(abi.encodeWithSelector(FeeQuoter.DestinationChainNotEnabled.selector, DEST_CHAIN_SELECTOR + 1));
    s_feeQuoter.getTokenAndGasPrices(s_sourceTokens[0], DEST_CHAIN_SELECTOR + 1);
  }
}
