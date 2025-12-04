// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPTokenPool_lockOrBurn is CCTPTokenPoolSetup {
  function test_lockOrBurn() public {
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
      sender: address(s_allowedCaller),
      amount: lockOrBurnIn.amount
    });

    vm.startPrank(s_allowedCaller);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) =
      s_cctpTokenPool.lockOrBurn(lockOrBurnIn, 0, "");

    assertEq(destTokenAmount, lockOrBurnIn.amount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
    assertEq(lockOrBurnOut.destPoolData, abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG));
  }

  function test_lockOrBurn_RevertWhen_IPoolV1NotSupported() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      amount: 1000000000000000000,
      localToken: address(s_USDCToken),
      originalSender: makeAddr("originalSender"),
      receiver: abi.encode(makeAddr("receiver"))
    });

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(CCTPTokenPool.IPoolV1NotSupported.selector));
    s_cctpTokenPool.lockOrBurn(lockOrBurnIn);
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
    s_cctpTokenPool.lockOrBurn(lockOrBurnIn, 0, "");
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

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_cctpTokenPool.lockOrBurn(lockOrBurnIn, 0, "");
  }
}
