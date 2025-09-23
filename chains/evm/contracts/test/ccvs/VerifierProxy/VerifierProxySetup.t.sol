// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VerifierProxy} from "../../../ccvs/VerifierProxy.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract VerifierProxySetup is BaseTest {
  VerifierProxy internal s_verifierProxy;

  function setUp() public override {
    super.setUp();

    s_verifierProxy = new VerifierProxy(makeAddr("MockImpl"));
  }
}
