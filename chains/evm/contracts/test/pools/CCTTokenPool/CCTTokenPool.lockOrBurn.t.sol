// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {CCTTokenPoolSetup} from "./CCTTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool_lockOrBurn is CCTTokenPoolSetup {
  function test_lockOrBurn() public {
    uint256 burnAmount = 1_000e18;

    // The onRamp calls lockOrBurn, and _lockOrBurn burns from msg.sender (the onRamp).
    // Give the onRamp tokens to burn.
    deal(address(s_cctPool), s_allowedOnRamp, burnAmount);

    uint256 supplyBefore = IERC20(address(s_cctPool)).totalSupply();

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: address(s_cctPool), amount: burnAmount
    });

    vm.expectEmit();
    emit IERC20.Transfer(s_allowedOnRamp, address(0), burnAmount);

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: address(s_cctPool), sender: s_allowedOnRamp, amount: burnAmount
    });

    s_cctPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_cctPool)
      })
    );

    assertEq(0, IERC20(address(s_cctPool)).balanceOf(s_allowedOnRamp));
    assertEq(supplyBefore - burnAmount, IERC20(address(s_cctPool)).totalSupply());
  }

  function test_lockOrBurn_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.startPrank(s_allowedOnRamp);
    vm.expectRevert(TokenPool.CursedByRMN.selector);

    s_cctPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: 1e18,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_cctPool)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = 8838833;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_cctPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: 1,
        remoteChainSelector: wrongChainSelector,
        localToken: address(s_cctPool)
      })
    );
  }
}
