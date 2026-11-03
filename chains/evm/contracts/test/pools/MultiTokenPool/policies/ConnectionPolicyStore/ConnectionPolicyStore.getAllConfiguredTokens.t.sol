// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getAllConfiguredTokens is ConnectionPolicyStoreSetup {
  function test_getAllConfiguredTokens() public view {
    address[] memory allTokens = new address[](2);
    allTokens[0] = s_token;
    allTokens[1] = s_otherToken;

    assertEq(allTokens, s_connectionStore.getAllConfiguredTokens(0, 2));
    assertEq(allTokens, s_connectionStore.getAllConfiguredTokens(0, type(uint64).max));
    assertEq(_arrayOf(s_token), s_connectionStore.getAllConfiguredTokens(0, 1));
    assertEq(_arrayOf(s_otherToken), s_connectionStore.getAllConfiguredTokens(1, 1));
    assertEq(_arrayOf(s_otherToken), s_connectionStore.getAllConfiguredTokens(1, 10));
  }

  function test_getAllConfiguredTokens_EmptyPage() public view {
    assertEq(0, s_connectionStore.getAllConfiguredTokens(0, 0).length);
    assertEq(0, s_connectionStore.getAllConfiguredTokens(2, 1).length);
    assertEq(0, s_connectionStore.getAllConfiguredTokens(type(uint64).max, type(uint64).max).length);
  }

  function test_getAllConfiguredTokens_Deregistration() public {
    address newToken = address(_deployToken("NEW", DEFAULT_TOKEN_DECIMALS));
    _registerToken(newToken);

    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    // Removal moves the last token into the vacated position.
    address[] memory allTokens = new address[](2);
    allTokens[0] = s_token;
    allTokens[1] = newToken;
    assertEq(allTokens, s_connectionStore.getAllConfiguredTokens(0, 10));
  }
}
