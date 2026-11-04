// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getSubsetTokens is ConnectionPolicyStoreSetup {
  function setUp() public override {
    super.setUp();
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token);
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_otherToken);
  }

  function test_getSubsetTokens() public view {
    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, 2);

    address[] memory expectedTokens = new address[](2);
    expectedTokens[0] = s_token;
    expectedTokens[1] = s_otherToken;
    assertTrue(onlySubsetTokens);
    assertEq(expectedTokens, subsetTokens);
  }

  function test_getSubsetTokens_Pagination() public view {
    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, 1);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_token), subsetTokens);

    (onlySubsetTokens, subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 1, 1);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_otherToken), subsetTokens);

    (onlySubsetTokens, subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 1, type(uint64).max);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_otherToken), subsetTokens);
  }

  function test_getSubsetTokens_EmptyPage() public view {
    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, 0);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);

    (onlySubsetTokens, subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 2, 1);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);

    (onlySubsetTokens, subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, type(uint64).max, type(uint64).max);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }

  function test_getSubsetTokens_EmptySubset() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, 1);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }

  function test_getSubsetTokens_UnknownRecord() public view {
    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(OTHER_POLICY, DEST_CHAIN_SELECTOR, 0, 1);
    assertFalse(onlySubsetTokens);
    assertEq(0, subsetTokens.length);

    (onlySubsetTokens, subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, 1);
    assertFalse(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }

  function test_getSubsetTokens_Unrestricted() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false));

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 1, 1);
    assertFalse(onlySubsetTokens);
    assertEq(_arrayOf(s_otherToken), subsetTokens);
  }

  function test_getSubsetTokens_Removal() public {
    s_connectionStore.applySubsetTokenUpdates(
      _getSubsetTokenArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token),
      new ConnectionPolicyStore.SubsetTokenArgs[](0)
    );

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, 1);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_otherToken), subsetTokens);

    (onlySubsetTokens, subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 1, 1);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }
}
