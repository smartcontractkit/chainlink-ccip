// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ProxySetup} from "./ProxySetup.t.sol";

interface IMockTarget {
  function getValue() external returns (uint8);
}

contract Proxy_fallback is ProxySetup {
  function test_fallback() public {
    address underlyingTarget = s_proxy.getTarget();

    bytes memory revertData = "0x12345678";
    vm.mockCallRevert(underlyingTarget, abi.encodeWithSelector(IMockTarget.getValue.selector), revertData);

    vm.expectRevert(revertData);
    IMockTarget(address(s_proxy)).getValue();

    // The return value is mocked to be `expectedValue` and should be bubbled up without modifying it.
    uint8 expectedValue = 1;
    vm.mockCall(underlyingTarget, abi.encodeWithSelector(IMockTarget.getValue.selector), abi.encode(expectedValue));
    vm.expectCall(underlyingTarget, abi.encodeWithSelector(IMockTarget.getValue.selector));

    uint8 value = IMockTarget(address(s_proxy)).getValue();
    assertEq(value, expectedValue);
  }
}
