// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_withdrawFee is AdvancedPoolHooksSetup {
  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 20 ether;

    s_token.transfer(address(s_tokenPool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit();
    emit TokenPool.FeeTokenWithdrawn(s_feeAggregator, address(s_token), feeAmount);

    s_tokenPool.withdrawFeeTokens(feeTokens);

    assertEq(s_token.balanceOf(s_feeAggregator), feeAmount);
    assertEq(s_token.balanceOf(address(s_tokenPool)), 0);
  }
}
