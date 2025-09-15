// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOnRampV1} from "../../interfaces/ICCVOnRampV1.sol";

import {RampProxy} from "../../RampProxy.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {RampProxySetup} from "./RampProxySetup.t.sol";

contract RampProxy_fallback is RampProxySetup {
  function test_fallback() public {
    MessageV1Codec.MessageV1 memory message;
    bytes memory data = ICCVOnRampV1(address(s_rampProxy)).forwardToVerifier(
      REMOTE_CHAIN_SELECTOR, address(this), message, "", address(0), 0, ""
    );
    assertEq(data, EXPECTED_VERIFIER_RESULT);
  }

  function test_fallback_CommitOnRamp() public {}

  function test_fallback_CommitOffRamp() public {}

  function test_fallback_RevertWhen_RemoteChainNotSupported() public {
    vm.expectRevert(
      abi.encodeWithSelector(RampProxy.RemoteChainNotSupported.selector, UNSUPPORTED_REMOTE_CHAIN_SELECTOR)
    );
    MessageV1Codec.MessageV1 memory message;
    bytes memory data = ICCVOnRampV1(address(s_rampProxy)).forwardToVerifier(
      UNSUPPORTED_REMOTE_CHAIN_SELECTOR, address(this), message, "", address(0), 0, ""
    );
  }
}
