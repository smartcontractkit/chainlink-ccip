// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract ERC20LockBoxSetup is BaseTest {
  IERC20 internal s_token;
  ERC20LockBox internal s_erc20LockBox;
  address internal s_allowedCaller = makeAddr("allowed_caller");
  address internal s_recipient = makeAddr("recipient");

  function setUp() public virtual override {
    super.setUp();
    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);
    deal(address(s_token), s_allowedCaller, type(uint256).max);

    // Deploy the ERC20 lock box (unsiloed selector = 0).
    s_erc20LockBox = new ERC20LockBox(address(s_token), 0);

    // Configure the allowed caller.
    address[] memory callers = new address[](1);
    callers[0] = s_allowedCaller;
    s_erc20LockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: callers, removedCallers: new address[](0)})
    );
  }

  function _depositTokens(
    uint256 amount
  ) internal {
    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);
    s_erc20LockBox.deposit(0, amount);
    vm.stopPrank();
  }
}
