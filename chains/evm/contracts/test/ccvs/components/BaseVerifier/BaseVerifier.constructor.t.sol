// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierTestHelper} from "../../../helpers/BaseVerifierTestHelper.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_constructor is BaseVerifierSetup {
  function test_constructor_RevertWhen_ZeroAddressNotAllowed_RMNIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new BaseVerifierTestHelper(s_storageLocations, address(0), BASE_VERIFIER_TEST_VERSION_TAG);
  }

  function test_constructor_RevertWhen_VersionTagCannotBeZero() public {
    vm.expectRevert(BaseVerifier.VersionTagCannotBeZero.selector);
    new BaseVerifierTestHelper(s_storageLocations, address(s_mockRMNRemote), bytes4(0));
  }
}
