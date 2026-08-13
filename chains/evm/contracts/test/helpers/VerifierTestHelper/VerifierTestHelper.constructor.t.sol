// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {VerifierTestHelper} from "../VerifierTestHelper.sol";
import {VerifierTestHelperSetup} from "./VerifierTestHelperSetup.t.sol";

contract VerifierTestHelper_constructor is VerifierTestHelperSetup {
  function test_constructor() public view {
    assertEq(s_verifier.typeAndVersion(), "VerifierTestHelper 2.0.0");
    assertEq(s_verifier.getTestRouter(), s_testRouter);
    assertEq(s_verifier.getTestToken(), s_testToken);
    assertEq(s_verifier.versionTag(), VERSION_TAG);
  }

  function test_constructor_RevertWhen_TestRouterIsZero() public {
    address testToken = s_verifier.getTestToken();

    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new VerifierTestHelper(address(0), testToken, s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);
  }

  function test_constructor_RevertWhen_TestTokenIsZero() public {
    address testRouter = s_verifier.getTestRouter();

    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new VerifierTestHelper(testRouter, address(0), s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);
  }
}
