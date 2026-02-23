// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

contract BaseERC20_setCCIPAdmin is BaseERC20Setup {
  function test_setCCIPAdmin() public {
    address newAdmin = makeAddr("newAdmin");

    vm.expectEmit();
    emit BaseERC20.CCIPAdminTransferred(OWNER, newAdmin);

    s_baseERC20.setCCIPAdmin(newAdmin);

    assertEq(newAdmin, s_baseERC20.getCCIPAdmin());
  }

  function test_setCCIPAdmin_TransferChain() public {
    address admin2 = makeAddr("admin2");
    address admin3 = makeAddr("admin3");

    s_baseERC20.setCCIPAdmin(admin2);
    assertEq(admin2, s_baseERC20.getCCIPAdmin());

    vm.startPrank(admin2);
    s_baseERC20.setCCIPAdmin(admin3);
    assertEq(admin3, s_baseERC20.getCCIPAdmin());
  }

  // Reverts

  function test_setCCIPAdmin_RevertWhen_OnlyCCIPAdmin() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(BaseERC20.OnlyCCIPAdmin.selector);
    s_baseERC20.setCCIPAdmin(STRANGER);
  }
}
