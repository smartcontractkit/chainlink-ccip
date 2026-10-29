// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__setInlineRecord is PolicyConfigSetup {
  function test__setInlineRecord() public {
    s_helper.setInlineRecord(s_token, CHAIN_SELECTOR);

    uint256[] memory inlineChains = s_helper.getInlineChains(s_token);
    assertEq(1, inlineChains.length);
    assertEq(CHAIN_SELECTOR, inlineChains[0]);
  }

  function test__setInlineRecord_ExistingRecord() public {
    s_helper.setInlineRecord(s_token, CHAIN_SELECTOR);
    s_helper.setInlineRecord(s_token, CHAIN_SELECTOR);

    assertEq(1, s_helper.getInlineChains(s_token).length);
  }

  // Reverts

  function test__setInlineRecord_RevertWhen_LocalTokenNotRegistered() public {
    address token = makeAddr("unregisteredToken");

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, token));
    s_helper.setInlineRecord(token, CHAIN_SELECTOR);
  }

  function test__setInlineRecord_RevertWhen_ZeroChainSelectorNotAllowed() public {
    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_helper.setInlineRecord(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
  }
}
