// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintWithLockReleaseFlagTokenPoolSetup} from "./BurnMintWithLockReleaseFlagTokenPoolSetup.t.sol";

import {LOCK_RELEASE_FLAG} from "../../../pools/USDC/SiloedUSDCTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintWithLockReleaseFlagTokenPool_lockOrBurn is BurnMintWithLockReleaseFlagTokenPoolSetup {
  function test_LockOrBurn_CorrectReturnData() public {
    uint256 burnAmount = 20_000e18;

    deal(address(s_token), address(s_pool), burnAmount);
    assertEq(s_token.balanceOf(address(s_pool)), burnAmount);

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: burnAmount
    });

    vm.expectEmit();
    emit IERC20.Transfer(address(s_pool), address(0), burnAmount);

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: burnAmount
    });

    bytes4 expectedSignature = bytes4(keccak256("burn(uint256)"));
    vm.expectCall(address(s_token), abi.encodeWithSelector(expectedSignature, burnAmount));

    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_pool)), 0);

    assertEq(bytes4(lockOrBurnOut.destPoolData), LOCK_RELEASE_FLAG);
  }
}
