// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_getTokenPolicyNamesBatch is FeePolicyStoreSetup {
  function test_getTokenPolicyNamesBatch() public {
    address[] memory localTokens = new address[](3);
    localTokens[0] = s_token;
    localTokens[1] = s_otherToken;
    localTokens[2] = makeAddr("unknownToken");

    (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) =
      s_feeStore.getTokenPolicyNamesBatch(localTokens);

    assertEq(3, feePolicyNames.length);
    assertEq(3, destOverheadPolicyNames.length);
    assertEq(FEE_POLICY, feePolicyNames[0]);
    assertEq(DEST_OVERHEAD_POLICY, destOverheadPolicyNames[0]);
    // A registered token without a policy and an unknown token both give zero.
    assertEq(bytes32(0), feePolicyNames[1]);
    assertEq(bytes32(0), destOverheadPolicyNames[1]);
    assertEq(bytes32(0), feePolicyNames[2]);
    assertEq(bytes32(0), destOverheadPolicyNames[2]);
  }
}
