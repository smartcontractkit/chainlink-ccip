// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierTestHelper} from "../../../helpers/BaseVerifierTestHelper.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_constructor is BaseVerifierSetup {
  function test_constructor_RevertWhen_ZeroAddressNotAllowed_RMNIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new BaseVerifierTestHelper(STORAGE_LOCATION, address(0));
  }
}
