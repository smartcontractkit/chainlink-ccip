// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../../ccvs/RampProxy.sol";
import {Proxiable} from "../../../ccvs/components/Proxiable.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract MockImplementation is Proxiable {
  error Failed();
  
  uint8 private s_value;

  constructor(uint8 value) {
    s_value = value;
  }

  function getValue() external view returns (uint8) {
    return s_value;
  }

  function revertWithError() external pure {
    revert Failed();
  }

  function upgradeTo(address newAddress) external {
    updateCodeAddress(newAddress);
  }
}

contract NotProxiable {
  function proxiableUUID() public pure returns (bytes32) {
    return keccak256("NOT_PROXIABLE");
  }
}

contract RampProxySetup is BaseTest {
  RampProxy internal s_rampProxy;
  MockImplementation internal s_currMockImpl;
  MockImplementation internal s_newMockImpl;
  NotProxiable internal s_notProxiable;

  uint8 internal constant CURR_EXPECTED_VALUE = 1;
  uint8 internal constant NEW_EXPECTED_VALUE = 2;

  function setUp() public override {
    super.setUp();

    s_currMockImpl = new MockImplementation(CURR_EXPECTED_VALUE);
    s_newMockImpl = new MockImplementation(NEW_EXPECTED_VALUE);
    s_rampProxy = new RampProxy(address(s_currMockImpl));
    s_notProxiable = new NotProxiable();
  }
}
