// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";
import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";

contract CCTPTokenPool_lockOrBurn is CCTPTokenPoolSetup {
  function test_lockOrBurn_PoolV1() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: 1,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(s_USDCToken),
      sender: address(s_routerAllowedOnRamp),
      amount: lockOrBurnIn.amount
    });

    vm.startPrank(s_routerAllowedOnRamp);
    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_cctpTokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(lockOrBurnOut.destTokenAddress, abi.encodePacked(address(s_USDCToken)));
    assertEq(lockOrBurnOut.destPoolData, abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG));
  }
}
