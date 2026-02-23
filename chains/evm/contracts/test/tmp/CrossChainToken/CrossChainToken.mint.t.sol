// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {CrossChainTokenSetup} from "./CrossChainTokenSetup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CrossChainToken_mint is CrossChainTokenSetup {
  address internal s_minter;

  function setUp() public virtual override {
    super.setUp();

    s_minter = makeAddr("minter");
    s_crossChainToken.grantRole(s_crossChainToken.BURN_MINT_ADMIN_ROLE(), OWNER);
    s_crossChainToken.grantRole(s_crossChainToken.MINTER_ROLE(), s_minter);
  }

  function test_mint() public {
    address receiver = makeAddr("receiver");
    uint256 amount = 1000e18;

    vm.startPrank(s_minter);

    vm.expectEmit();
    emit IERC20.Transfer(address(0), receiver, amount);

    s_crossChainToken.mint(receiver, amount);

    assertEq(amount, s_crossChainToken.balanceOf(receiver));
    assertEq(PRE_MINT + amount, s_crossChainToken.totalSupply());
  }

  // Reverts

  function test_mint_RevertWhen_MaxSupplyExceeded() public {
    uint256 remaining = MAX_SUPPLY - s_crossChainToken.totalSupply();

    vm.startPrank(s_minter);

    vm.expectRevert(
      abi.encodeWithSelector(
        CrossChainToken.MaxSupplyExceeded.selector, s_crossChainToken.totalSupply() + remaining + 1
      )
    );
    s_crossChainToken.mint(makeAddr("receiver"), remaining + 1);
  }

  function test_mint_RevertWhen_InvalidRecipient() public {
    vm.startPrank(s_minter);

    vm.expectRevert(abi.encodeWithSelector(BaseERC20.InvalidRecipient.selector, address(s_crossChainToken)));
    s_crossChainToken.mint(address(s_crossChainToken), 1e18);
  }

  function test_mint_RevertWhen_MissingRole() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, STRANGER, s_crossChainToken.MINTER_ROLE()
      )
    );
    s_crossChainToken.mint(STRANGER, 1e18);
  }
}
