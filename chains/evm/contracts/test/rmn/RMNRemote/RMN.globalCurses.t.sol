// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {GLOBAL_CURSE_SUBJECT} from "../../../rmn/RMN.sol";
import {RMNRemoteSetup} from "./RMNSetup.t.sol";

contract RMNRemote_global_curses is RMNRemoteSetup {
  function test_isCursed_globalCurseSubject() public {
    bytes16 randSubject = bytes16(keccak256("random subject"));
    assertFalse(s_rmn.isCursed());
    assertFalse(s_rmn.isCursed(randSubject));

    s_rmn.curse(GLOBAL_CURSE_SUBJECT);
    assertTrue(s_rmn.isCursed());
    assertTrue(s_rmn.isCursed(randSubject));

    s_rmn.uncurse(GLOBAL_CURSE_SUBJECT);
    assertFalse(s_rmn.isCursed());
    assertFalse(s_rmn.isCursed(randSubject));
  }
}
