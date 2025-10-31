// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";
import {ERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/ERC20.sol";

contract TokenPoolV2_withdrawFee is TokenPoolV2Setup {
  function test_withdrawFee_SendsPoolTokenToRecipient() public {
    uint256 feeAmount = 20 ether;
    address recipient = makeAddr("fee_recipient");

    s_token.transfer(address(s_tokenPool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit(true, true, true, true, address(s_tokenPool));
    emit TokenPool.FeeTokenWithdrawn(recipient, address(s_token), feeAmount);
    vm.expectEmit(true, true, true, true, address(s_tokenPool));
    emit TokenPool.PoolFeeWithdrawn(recipient, feeAmount);

    s_tokenPool.withdrawFee(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount);
    assertEq(s_token.balanceOf(address(s_tokenPool)), 0);
    assertEq(s_tokenPool.getAccumulatedFees(), 0);
  }

  function test_withdrawFee_ForwardsNonPoolFeeTokensToRecipient() public {
    ERC20 feeToken = new ERC20("feeToken", "FEE");
    uint256 feeTokenAmount = 10 ether;
    address recipient = makeAddr("fee_token_recipient");

    deal(address(feeToken), address(s_tokenPool), feeTokenAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(feeToken);

    vm.expectEmit();
    emit TokenPool.FeeTokenWithdrawn(recipient, address(feeToken), feeTokenAmount);

    s_tokenPool.withdrawFee(feeTokens, recipient);

    assertEq(feeToken.balanceOf(recipient), feeTokenAmount);
    assertEq(feeToken.balanceOf(address(s_tokenPool)), 0);
  }

  function test_withdrawFee_RevertsWhen_CalledByNonOwner() public {
    address recipient = makeAddr("fee_token_recipient");

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);
    s_tokenPool.withdrawFee(feeTokens, recipient);
  }
}
