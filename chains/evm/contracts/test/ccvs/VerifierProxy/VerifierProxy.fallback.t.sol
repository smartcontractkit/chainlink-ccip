// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VerifierProxySetup} from "./VerifierProxySetup.t.sol";

interface IMockVerifier {
  function getValue(
    address caller
  ) external returns (uint8);
}

contract VerifierProxy_fallback is VerifierProxySetup {
  function test_fallback() public {
    address underlyingVerifier = s_verifierProxy.getVerifier();

    // This value will be passed into the proxy call and it should be overwritten.
    address callerArg = makeAddr("CallerArg");

    // Send from expectedCallerOverride, so we can expect this address to be passed into the underlying verifier.
    address expectedCallerOverride = makeAddr("ExpectedCallerOverride");
    vm.startPrank(expectedCallerOverride);

    assertTrue(address(s_verifierProxy) != callerArg, "topLevelCaller should not be the proxy itself");

    bytes memory revertData = "0x12345678";
    vm.mockCallRevert(
      underlyingVerifier,
      abi.encodeWithSelector(IMockVerifier.getValue.selector, address(expectedCallerOverride)),
      revertData
    );

    vm.expectRevert(revertData);
    IMockVerifier(address(s_verifierProxy)).getValue(callerArg);

    // We expect a call to the underlying verifier with the callerArg replaced by expectedCallerOverride.
    // The return value is mocked to be `expectedValue` and should be bubbled up without modifying it.
    uint8 expectedValue = 1;
    vm.mockCall(
      underlyingVerifier,
      abi.encodeWithSelector(IMockVerifier.getValue.selector, address(expectedCallerOverride)),
      abi.encode(expectedValue)
    );
    vm.expectCall(
      underlyingVerifier, abi.encodeWithSelector(IMockVerifier.getValue.selector, address(expectedCallerOverride))
    );

    uint8 value = IMockVerifier(address(s_verifierProxy)).getValue(callerArg);
    assertEq(value, expectedValue);
  }
}
