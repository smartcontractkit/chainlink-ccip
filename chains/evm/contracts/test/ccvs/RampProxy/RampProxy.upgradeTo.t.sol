// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxySetup, MockImplementation} from "./RampProxySetup.t.sol";
import {Proxiable} from "../../../ccvs/components/Proxiable.sol";

contract RampProxy_upgradeTo is RampProxySetup {
  function test_upgradeTo_ViaProxy() public {
    MockImplementation(address(s_rampProxy)).upgradeTo(address(s_newMockImpl));
    uint8 value = MockImplementation(address(s_rampProxy)).getValue();
    assertEq(value, NEW_EXPECTED_VALUE);
  }

  function test_upgradeTo_ViaImpl() public {
    s_currMockImpl.upgradeTo(address(s_newMockImpl));
    uint8 value = MockImplementation(address(s_rampProxy)).getValue();
    assertEq(value, NEW_EXPECTED_VALUE);
  }

  function test_fallback_RevertWhen_NewContractNotProxiable() public {
    vm.expectRevert(abi.encodeWithSelector(Proxiable.NewContractNotProxiable.selector));
    MockImplementation(address(s_rampProxy)).upgradeTo(address(s_notProxiable));
  }
}
