// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__moveToken is PolicyConfigSetup {
  function setUp() public override {
    super.setUp();
    s_helper.setPolicyRecord(POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
    s_helper.setPolicyRecord(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
    s_helper.moveToken(bytes32(0), POLICY);
  }

  function test__moveToken() public {
    assertEq(1, s_helper.getTokenCount(POLICY));

    s_helper.moveToken(POLICY, OTHER_POLICY);
    assertEq(0, s_helper.getTokenCount(POLICY));
    assertEq(1, s_helper.getTokenCount(OTHER_POLICY));

    // A move to zero takes the token out of its policy.
    s_helper.moveToken(OTHER_POLICY, bytes32(0));
    assertEq(0, s_helper.getTokenCount(OTHER_POLICY));
  }

  function test__moveToken_SamePolicy() public {
    s_helper.moveToken(POLICY, POLICY);
    assertEq(1, s_helper.getTokenCount(POLICY));
  }

  function test__moveToken_MultipleTokens() public {
    s_helper.moveToken(bytes32(0), POLICY);
    assertEq(2, s_helper.getTokenCount(POLICY));

    s_helper.moveToken(POLICY, OTHER_POLICY);
    assertEq(1, s_helper.getTokenCount(POLICY));
    assertEq(1, s_helper.getTokenCount(OTHER_POLICY));
  }

  // Reverts

  function test__moveToken_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, bytes32("UNKNOWN_POLICY")));
    s_helper.moveToken(POLICY, "UNKNOWN_POLICY");
  }
}
