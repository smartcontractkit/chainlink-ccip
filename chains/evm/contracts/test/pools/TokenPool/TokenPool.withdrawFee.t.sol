// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_withdrawFee is TokenPoolV2Setup {
  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 20 ether;
    address recipient = makeAddr("fee_recipient");

    s_token.transfer(address(s_tokenPool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit();
    emit TokenPool.FeeTokenWithdrawn(recipient, address(s_token), feeAmount);

    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount);
    assertEq(s_token.balanceOf(address(s_tokenPool)), 0);
  }

  function test_withdrawFeeTokens_RevtwertsWhen_OnlyCallableByOwner() public {
    address recipient = makeAddr("fee_token_recipient");

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);
  }
}
