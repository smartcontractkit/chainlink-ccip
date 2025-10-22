// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Proxy} from "../../Proxy.sol";
import {BaseTest} from "../BaseTest.t.sol";

contract ProxySetup is BaseTest {
  Proxy internal s_proxy;

  function setUp() public override {
    super.setUp();

    s_proxy = new Proxy(makeAddr("MockTarget"));
  }
}
