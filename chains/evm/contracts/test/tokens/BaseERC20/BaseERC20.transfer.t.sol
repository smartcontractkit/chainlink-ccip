// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";
import {IERC20Errors} from "@openzeppelin/contracts@5.3.0/interfaces/draft-IERC6093.sol";

contract BaseERC20_transfer is BaseERC20Setup {
  function test_transfer() public {
    address receiver = makeAddr("receiver");
    uint256 amount = 1000e18;

    s_baseERC20.transfer(receiver, amount);

    assertEq(amount, s_baseERC20.balanceOf(receiver));
    assertEq(PRE_MINT - amount, s_baseERC20.balanceOf(OWNER));
  }

  // Reverts

  function test_transfer_RevertWhen_InvalidRecipient_TransferToSelf() public {
    vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InvalidReceiver.selector, address(s_baseERC20)));
    s_baseERC20.transfer(address(s_baseERC20), 100e18);
  }

  function test_approve_RevertWhen_InvalidRecipient_ApproveToSelf() public {
    vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InvalidSpender.selector, address(s_baseERC20)));
    s_baseERC20.approve(address(s_baseERC20), 100e18);
  }
}
