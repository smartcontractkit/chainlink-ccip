// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_setVerifier is LombardTokenPoolSetup {
  function test_setVerifier() public {
    address newVerifier = makeAddr("newVerifier");

    vm.prank(OWNER);
    s_pool.setVerifier(newVerifier);

    assertEq(s_pool.s_verifier(), newVerifier);
    assertEq(s_token.allowance(address(s_pool), newVerifier), type(uint256).max);
    assertEq(s_token.allowance(address(s_pool), VERIFIER), 0);
  }

  function test_setVerifier_RevertWhen_ZeroVerifierNotAllowed() public {
    vm.prank(OWNER);
    vm.expectRevert(LombardTokenPool.ZeroVerifierNotAllowed.selector);
    s_pool.setVerifier(address(0));
  }
}
