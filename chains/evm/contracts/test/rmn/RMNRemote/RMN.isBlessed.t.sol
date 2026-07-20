// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMN} from "../../../interfaces/IRMN.sol";

import {RMNRemoteSetup} from "./RMNSetup.t.sol";

contract RMNRemote_isBlessed is RMNRemoteSetup {
  function test_isBlessed_AlwaysReturnsTrue() public view {
    // Random data that has not been blessed, but isBlessed should always return true.
    IRMN.TaggedRoot memory taggedRoot = IRMN.TaggedRoot({commitStore: address(0), root: bytes32(0)});

    assertTrue(s_rmn.isBlessed(taggedRoot));
  }
}

