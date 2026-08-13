// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {VerifierTestHelper} from "../VerifierTestHelper.sol";
import {VerifierTestHelperSetup} from "./VerifierTestHelperSetup.t.sol";

contract VerifierTestHelper_constructor is VerifierTestHelperSetup {
  function test_constructor() public view {
    assertEq(s_verifier.typeAndVersion(), "VerifierTestHelper 2.0.0");
    assertEq(s_verifier.getTestRouter(), s_testRouter);
    assertEq(s_verifier.versionTag(), VERSION_TAG);
  }

  function test_constructor_RevertWhen_TestRouterIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new VerifierTestHelper(address(0), s_testToken, s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);
  }

  function test_constructor_RevertWhen_TestTokenIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new VerifierTestHelper(s_testRouter, address(0), s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);
  }
}
