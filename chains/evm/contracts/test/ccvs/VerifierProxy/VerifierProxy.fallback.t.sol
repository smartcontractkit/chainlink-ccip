// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VerifierProxySetup} from "./VerifierProxySetup.t.sol";

interface IMockVerifier {
  function getValue() external returns (uint8);
}

contract VerifierProxy_fallback is VerifierProxySetup {
  function test_fallback() public {
    bytes memory revertData = "0x12345678";
    vm.mockCallRevert(s_verifierProxy.s_verifier(), abi.encodeWithSelector(IMockVerifier.getValue.selector), revertData);
    vm.expectRevert(revertData);
    IMockVerifier(address(s_verifierProxy)).getValue();

    uint8 expectedValue = 1;
    vm.mockCall(s_verifierProxy.s_verifier(), abi.encodeWithSelector(IMockVerifier.getValue.selector), abi.encode(1));
    uint8 value = IMockVerifier(address(s_verifierProxy)).getValue();
    assertEq(value, expectedValue);
  }
}
