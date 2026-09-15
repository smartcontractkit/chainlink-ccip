// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

contract SuccinctZKVerifier_versionTag is SuccinctZKVerifierSetup {
  function test_versionTag() public view {
    assertEq(VERSION_TAG_V0_0_1, s_zkVerifier.versionTag());
  }

  function test_versionTag_MatchesExpectedHash() public pure {
    assertEq(bytes4(keccak256("SuccinctZKVerifier 0.0.1-dev")), VERSION_TAG_V0_0_1);
  }
}
