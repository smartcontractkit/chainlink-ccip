// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VerifierProxy} from "../../../ccvs/VerifierProxy.sol";
import {VerifierProxySetup} from "./VerifierProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VerifierProxy_setVerifier is VerifierProxySetup {
  function test_setVerifier() public {
    address newVerifier = makeAddr("NewVerifier");

    s_verifierProxy.setVerifier(newVerifier);

    assertEq(s_verifierProxy.getVerifier(), newVerifier);
  }

  function test_setVerifier_RevertWhen_ZeroAddressNotAllowed() public {
    vm.expectRevert(VerifierProxy.ZeroAddressNotAllowed.selector);
    s_verifierProxy.setVerifier(address(0));
  }

  function test_setVerifier_RevertWhen_NotOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_verifierProxy.setVerifier(makeAddr("NewVerifier"));
  }
}
