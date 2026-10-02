// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";

contract LombardVerifier_versionTag is LombardVerifierSetup {
  function test_versionTag() public view {
    assertEq(s_lombardVerifier.versionTag(), VERSION_TAG_V2_0_0);
  }

  function test_versionTag_MatchesExpectedHash() public pure {
    assertEq(VERSION_TAG_V2_0_0, bytes4(keccak256("LombardVerifier 2.0.0")));
  }
}
