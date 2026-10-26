// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../libraries/FeeTokenHandler.sol";
import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_withdrawFeeTokens is MultiTokenPoolSetup {
  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 20e18;
    address recipient = makeAddr("fee_recipient");
    deal(address(s_token), address(s_pool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit(address(s_pool));
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount);

    s_poolConfigManager.withdrawFeeTokens(s_pool, feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_withdrawFeeTokens_RevertWhen_CallerIsNotPoolConfigManager() public {
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectRevert(abi.encodeWithSelector(MultiTokenPool.CallerIsNotPoolConfigManager.selector, OWNER));
    s_pool.withdrawFeeTokens(feeTokens, makeAddr("fee_recipient"));
  }
}
