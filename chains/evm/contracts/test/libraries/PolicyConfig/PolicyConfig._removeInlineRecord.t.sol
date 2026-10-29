// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__removeInlineRecord is PolicyConfigSetup {
  function setUp() public override {
    super.setUp();
    s_helper.setInlineRecord(s_token, CHAIN_SELECTOR);
  }

  function test__removeInlineRecord() public {
    s_helper.removeInlineRecord(s_token, CHAIN_SELECTOR);

    assertEq(0, s_helper.getInlineChainSelectors(s_token).length);
  }

  // Reverts

  function test__removeInlineRecord_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_token, OTHER_CHAIN_SELECTOR));
    s_helper.removeInlineRecord(s_token, OTHER_CHAIN_SELECTOR);
  }
}
