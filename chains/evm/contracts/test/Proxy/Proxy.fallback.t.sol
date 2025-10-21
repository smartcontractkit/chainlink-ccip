// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ProxySetup} from "./ProxySetup.t.sol";

interface IMockTarget {
  function getValue(
    address caller
  ) external returns (uint8);
}

contract Proxy_fallback is ProxySetup {
  function test_fallback() public {
    address underlyingTarget = s_proxy.getTarget();

    // This value will be passed into the proxy call and it should be overwritten.
    address callerArg = makeAddr("CallerArg");

    // Send from expectedCallerOverride, so we can expect this address to be passed into the underlying target.
    address expectedCallerOverride = makeAddr("ExpectedCallerOverride");
    vm.startPrank(expectedCallerOverride);

    assertTrue(address(s_proxy) != callerArg, "topLevelCaller should not be the proxy itself");

    bytes memory revertData = "0x12345678";
    vm.mockCallRevert(
      underlyingTarget,
      abi.encodeWithSelector(IMockTarget.getValue.selector, address(expectedCallerOverride)),
      revertData
    );

    vm.expectRevert(revertData);
    IMockTarget(address(s_proxy)).getValue(callerArg);

    // We expect a call to the underlying target with the callerArg replaced by expectedCallerOverride.
    // The return value is mocked to be `expectedValue` and should be bubbled up without modifying it.
    uint8 expectedValue = 1;
    vm.mockCall(
      underlyingTarget,
      abi.encodeWithSelector(IMockTarget.getValue.selector, address(expectedCallerOverride)),
      abi.encode(expectedValue)
    );
    vm.expectCall(
      underlyingTarget, abi.encodeWithSelector(IMockTarget.getValue.selector, address(expectedCallerOverride))
    );

    uint8 value = IMockTarget(address(s_proxy)).getValue(callerArg);
    assertEq(value, expectedValue);
  }
}
