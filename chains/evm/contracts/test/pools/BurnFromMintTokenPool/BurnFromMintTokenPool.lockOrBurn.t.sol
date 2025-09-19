// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnFromMintTokenPoolSetup} from "./BurnFromMintTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnFromMintTokenPool_lockOrBurn is BurnFromMintTokenPoolSetup {
  function test_constructor() public view {
    assertEq(address(s_token), address(s_pool.getToken()));
    assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq(type(uint256).max, s_token.allowance(address(s_pool), address(s_pool)));
  }

  function test_lockOrBurn() public {
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

    bytes4 expectedSignature = bytes4(keccak256("burnFrom(address,uint256)"));
    vm.expectCall(address(s_token), abi.encodeWithSelector(expectedSignature, address(s_pool), burnAmount));

    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  // Should not burn tokens if cursed.
  function test_lockOrBurn_RevertWhen_PoolBurnRevertNotHealthy() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    uint256 before = s_token.balanceOf(address(s_pool));
    vm.startPrank(s_allowedOnRamp);

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: 1e5,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_pool)), before);
  }

  function test_lockOrBurn_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = 8838833;
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: 1,
        localToken: address(s_token),
        remoteChainSelector: wrongChainSelector,
        sourcePoolAddress: _generateSourceTokenData().sourcePoolAddress,
        sourcePoolData: _generateSourceTokenData().extraData,
        offchainTokenData: ""
      })
    );
  }
}
