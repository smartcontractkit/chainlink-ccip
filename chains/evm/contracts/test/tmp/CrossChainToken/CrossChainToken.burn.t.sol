// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CrossChainTokenSetup} from "./CrossChainTokenSetup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CrossChainToken_burn is CrossChainTokenSetup {
  address internal s_burner;

  function setUp() public virtual override {
    super.setUp();

    s_burner = makeAddr("burner");
    s_crossChainToken.grantRole(s_crossChainToken.BURN_MINT_ADMIN_ROLE(), OWNER);
    s_crossChainToken.grantRole(s_crossChainToken.BURNER_ROLE(), s_burner);

    // Transfer tokens to burner so they can burn
    s_crossChainToken.transfer(s_burner, 10_000e18);
  }

  function test_burn() public {
    uint256 burnAmount = 1000e18;
    uint256 balanceBefore = s_crossChainToken.balanceOf(s_burner);

    vm.startPrank(s_burner);

    vm.expectEmit();
    emit IERC20.Transfer(s_burner, address(0), burnAmount);

    s_crossChainToken.burn(burnAmount);

    assertEq(balanceBefore - burnAmount, s_crossChainToken.balanceOf(s_burner));
    assertEq(PRE_MINT - burnAmount, s_crossChainToken.totalSupply());
  }

  function test_burnFrom() public {
    uint256 burnAmount = 500e18;
    address tokenHolder = makeAddr("tokenHolder");

    // Transfer tokens to holder and approve burner
    vm.startPrank(OWNER);
    s_crossChainToken.transfer(tokenHolder, burnAmount);

    vm.startPrank(tokenHolder);
    s_crossChainToken.approve(s_burner, burnAmount);

    vm.startPrank(s_burner);

    vm.expectEmit();
    emit IERC20.Transfer(tokenHolder, address(0), burnAmount);

    s_crossChainToken.burnFrom(tokenHolder, burnAmount);

    assertEq(0, s_crossChainToken.balanceOf(tokenHolder));
  }

  function test_burn_AddressOverload() public {
    uint256 burnAmount = 500e18;
    address tokenHolder = makeAddr("tokenHolder");

    vm.startPrank(OWNER);
    s_crossChainToken.transfer(tokenHolder, burnAmount);

    vm.startPrank(tokenHolder);
    s_crossChainToken.approve(s_burner, burnAmount);

    vm.startPrank(s_burner);

    s_crossChainToken.burn(tokenHolder, burnAmount);

    assertEq(0, s_crossChainToken.balanceOf(tokenHolder));
  }

  // Reverts

  function test_burn_RevertWhen_MissingRole() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, STRANGER, s_crossChainToken.BURNER_ROLE()
      )
    );
    s_crossChainToken.burn(1e18);
  }

  function test_burnFrom_RevertWhen_MissingRole() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, STRANGER, s_crossChainToken.BURNER_ROLE()
      )
    );
    s_crossChainToken.burnFrom(OWNER, 1e18);
  }
}
