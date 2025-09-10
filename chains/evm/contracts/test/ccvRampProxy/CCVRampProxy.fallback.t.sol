// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";

import {ICCVOnRamp} from "../../interfaces/ICCVOnRamp.sol";
import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {CCVRampProxySetup} from "./CCVRampProxySetup.t.sol";

contract CCVRampProxy_fallback is CCVRampProxySetup {
  function test_fallback() public {
    bytes memory data =
      ICCVOnRamp(address(s_ccvRampProxy)).forwardToVerifier(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, address(this), "", 0);
    assertEq(data, "");
  }

  function test_fallback_RevertWhen_UnknownRemoteChainSelector() public {
    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.UnknownRamp.selector, 2222, CCVRamp.V1));
    ICCVOnRamp(address(s_ccvRampProxy)).forwardToVerifier(2222, CCVRamp.V1, address(this), "", 0);
  }

  function test_fallback_RevertWhen_UnknownVersion() public {
    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.UnknownRamp.selector, REMOTE_CHAIN_SELECTOR, bytes32("V7")));
    ICCVOnRamp(address(s_ccvRampProxy)).forwardToVerifier(REMOTE_CHAIN_SELECTOR, bytes32("V7"), address(this), "", 0);
  }
}
