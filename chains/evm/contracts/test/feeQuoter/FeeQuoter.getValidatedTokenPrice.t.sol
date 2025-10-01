// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Internal} from "../../libraries/Internal.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_getValidatedTokenPrice is FeeQuoterSetup {
  function test_GetValidatedTokenPrice() public view {
    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedInitialPriceUpdates, (Internal.PriceUpdates));
    address token = priceUpdates.tokenPriceUpdates[0].sourceToken;

    uint224 tokenPrice = s_feeQuoter.getValidatedTokenPrice(token);

    assertEq(priceUpdates.tokenPriceUpdates[0].usdPerToken, tokenPrice);
  }

  // Reverts

  function test_RevertWhen_TokenNotSupported() public {
    vm.expectRevert(abi.encodeWithSelector(FeeQuoter.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_feeQuoter.getValidatedTokenPrice(DUMMY_CONTRACT_ADDRESS);
  }
}
