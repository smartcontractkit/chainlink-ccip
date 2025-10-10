// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RouterSetup} from "../../onRamp/OnRamp/RouterSetup.t.sol";

contract Router_constructor is RouterSetup {
  function test_Constructor() public view {
    assertEq("Router 1.2.0", s_sourceRouter.typeAndVersion());
    assertEq(OWNER, s_sourceRouter.owner());
  }
}
