// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMNRemote} from "../../../rmn/RMNRemote.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract RMNRemoteSetup is BaseTest {
  RMNRemote public s_rmnRemote;

  bytes16 internal constant CURSE_SUBJ_1 = bytes16(keccak256("subject 1"));
  bytes16 internal constant CURSE_SUBJ_2 = bytes16(keccak256("subject 2"));
  bytes16[] internal s_curseSubjects;

  function setUp() public virtual override {
    super.setUp();
    s_rmnRemote = new RMNRemote(new address[](0));
    s_curseSubjects = [CURSE_SUBJ_1, CURSE_SUBJ_2];
  }
}
