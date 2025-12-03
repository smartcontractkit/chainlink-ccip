// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";
import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPTokenPool_lockOrBurn is CCTPTokenPoolSetup {
  function test_lockOrBurn_PoolV1() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(s_USDCToken),
      sender: address(ALLOWED_CALLER),
      amount: lockOrBurnIn.amount
    });

    vm.startPrank(ALLOWED_CALLER);
    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_cctpTokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
    assertEq(lockOrBurnOut.destPoolData, abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG));
  }

  function test_lockOrBurn_PoolV2() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(s_USDCToken),
      sender: address(ALLOWED_CALLER),
      amount: lockOrBurnIn.amount
    });

    vm.startPrank(ALLOWED_CALLER);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) = s_cctpTokenPool.lockOrBurn(lockOrBurnIn, 0, "");

    assertEq(destTokenAmount, lockOrBurnIn.amount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
    assertEq(lockOrBurnOut.destPoolData, abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG));
  }

  function test_lockOrBurn_RevertWhen_InvalidCaller() public {
    address invalidCaller = makeAddr("invalidCaller");

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.startPrank(invalidCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, invalidCaller));
    s_cctpTokenPool.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = DEST_CHAIN_SELECTOR + 1;

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: wrongChainSelector,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.startPrank(ALLOWED_CALLER);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_cctpTokenPool.lockOrBurn(lockOrBurnIn);
  }
}
