// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../interfaces/IRouter.sol";

import {Router} from "../Router.sol";
import {BaseTest} from "./BaseTest.t.sol";

/// @notice Mix-in for test suites that exercise real message routing. Overrides the BaseTest router hook to deploy
/// real Router contracts. Kept in a separate file so that only suites that route for real compile Router.
contract RouterFixture is BaseTest {
  function _setUpRouters() internal virtual override {
    s_sourceRouter = IRouter(address(new Router(s_weth, address(s_mockRMNRemote))));
    vm.label(address(s_sourceRouter), "sourceRouter");
    s_destRouter = IRouter(address(new Router(s_destWeth, address(s_mockRMNRemote))));
    vm.label(address(s_destRouter), "destRouter");
  }
}
