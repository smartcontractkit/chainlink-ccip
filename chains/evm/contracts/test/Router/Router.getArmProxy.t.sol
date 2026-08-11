// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../Router.sol";
import {RouterFixture} from "../RouterFixture.t.sol";

contract Router_getArmProxy is RouterFixture {
  function test_getArmProxy() public view {
    assertEq(Router(payable(address(s_sourceRouter))).getArmProxy(), address(s_mockRMNRemote));
  }
}
