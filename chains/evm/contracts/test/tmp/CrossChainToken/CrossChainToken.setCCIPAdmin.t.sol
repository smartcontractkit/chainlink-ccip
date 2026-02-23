// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainTokenSetup} from "./CrossChainTokenSetup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";

contract CrossChainToken_setCCIPAdmin is CrossChainTokenSetup {
  function test_setCCIPAdmin() public {
    address newAdmin = makeAddr("newAdmin");

    vm.expectEmit();
    emit BaseERC20.CCIPAdminTransferred(OWNER, newAdmin);

    s_crossChainToken.setCCIPAdmin(newAdmin);

    assertEq(newAdmin, s_crossChainToken.getCCIPAdmin());
  }

  // Reverts

  function test_setCCIPAdmin_RevertWhen_MissingRole() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, STRANGER, s_crossChainToken.DEFAULT_ADMIN_ROLE()
      )
    );
    s_crossChainToken.setCCIPAdmin(STRANGER);
  }
}
