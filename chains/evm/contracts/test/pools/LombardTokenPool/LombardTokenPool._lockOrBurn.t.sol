// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_lockOrBurn is LombardTokenPoolSetup {
  function test_lockOrBurn_ForwardsToVerifier() public {
    uint256 amount = 1e18;
    vm.prank(OWNER);
    s_token.transfer(address(s_pool), amount);

    vm.prank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory out, uint256 destAmount) = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encodePacked(address(0xDEAD)),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(destAmount, amount);
    assertEq(out.destTokenAddress, abi.encode(s_remoteToken));
    assertEq(out.destPoolData, abi.encode(uint8(DEFAULT_TOKEN_DECIMALS)));
    assertEq(s_token.balanceOf(VERIFIER), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_setVerifierUpdatesAllowance() public {
    address newVerifier = makeAddr("newVerifier");

    vm.prank(OWNER);
    s_pool.setVerifier(newVerifier);

    assertEq(s_pool.s_verifier(), newVerifier);
    assertEq(s_token.allowance(address(s_pool), newVerifier), type(uint256).max);
    assertEq(s_token.allowance(address(s_pool), VERIFIER), 0);
  }

  function test_setVerifierRevertsOnZero() public {
    vm.prank(OWNER);
    vm.expectRevert(LombardTokenPool.ZeroVerifierNotAllowed.selector);
    s_pool.setVerifier(address(0));
  }
}
