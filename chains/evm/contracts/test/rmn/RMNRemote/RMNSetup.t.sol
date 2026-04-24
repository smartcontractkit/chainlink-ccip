// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMN} from "../../../rmn/RMN.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract RMNRemoteSetup is BaseTest {
  RMN public s_rmn;

  bytes16 internal constant CURSE_SUBJ_1 = bytes16(keccak256("subject 1"));
  bytes16 internal constant CURSE_SUBJ_2 = bytes16(keccak256("subject 2"));
  bytes16[] internal s_curseSubjects;

  function setUp() public virtual override {
    super.setUp();
    s_rmn = new RMN(new address[](0));
    s_curseSubjects = [CURSE_SUBJ_1, CURSE_SUBJ_2];
  }
}
