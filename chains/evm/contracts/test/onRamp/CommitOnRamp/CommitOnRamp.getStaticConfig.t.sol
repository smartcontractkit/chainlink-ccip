// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_getStaticConfig is CommitOnRampSetup {
  function test_GetStaticConfig() public view {
    CommitOnRamp.StaticConfig memory s = s_commitOnRamp.getStaticConfig();
    assertEq(s.rmnRemote, address(s_mockRMNRemote));
    assertEq(s.nonceManager, address(s_nonceManager));
  }
}
