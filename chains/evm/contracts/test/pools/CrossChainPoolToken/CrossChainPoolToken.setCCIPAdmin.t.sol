// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainPoolTokenSetup} from "./CrossChainPoolTokenSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CrossChainPoolToken_setCCIPAdmin is CrossChainPoolTokenSetup {
  function test_setCCIPAdmin() public {
    address newAdmin = makeAddr("newAdmin");

    vm.expectEmit();
    emit BaseERC20.CCIPAdminTransferred(OWNER, newAdmin);

    s_cctPool.setCCIPAdmin(newAdmin);

    assertEq(newAdmin, s_cctPool.getCCIPAdmin());
  }

  function test_setCCIPAdmin_ToZeroAddress() public {
    vm.expectEmit();
    emit BaseERC20.CCIPAdminTransferred(OWNER, address(0));

    s_cctPool.setCCIPAdmin(address(0));

    assertEq(address(0), s_cctPool.getCCIPAdmin());
  }

  // Reverts

  function test_setCCIPAdmin_RevertWhen_CalledByStranger() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctPool.setCCIPAdmin(STRANGER);
  }

  function test_setCCIPAdmin_RevertWhen_CalledByCCIPAdmin() public {
    address ccipAdmin = makeAddr("ccipAdmin");
    s_cctPool.setCCIPAdmin(ccipAdmin);

    vm.startPrank(ccipAdmin);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctPool.setCCIPAdmin(makeAddr("anotherAdmin"));
  }
}
