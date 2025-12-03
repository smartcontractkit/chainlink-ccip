// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {BurnWithFromMintTokenPool} from "../../../pools/BurnWithFromMintTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintSetup} from "../BurnMintTokenPool/BurnMintSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnWithFromMintTokenPoolSetup is BurnMintSetup {
  BurnWithFromMintTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnWithFromMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}

contract BurnWithFromMintTokenPool_lockOrBurn is BurnWithFromMintTokenPoolSetup {
  function test_constructor() public view {
    assertEq(address(s_token), address(s_pool.getToken()));
    assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq(type(uint256).max, s_token.allowance(address(s_pool), address(s_pool)));
  }

  function test_lockOrBurn() public {
    uint256 amount = 20_000e18;

    deal(address(s_token), address(s_pool), amount);
    assertEq(s_token.balanceOf(address(s_pool)), amount);

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: amount
    });

    vm.expectEmit();
    emit IERC20.Transfer(address(s_pool), address(0), amount);

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: amount
    });

    bytes4 expectedSignature = bytes4(keccak256("burn(address,uint256)"));
    vm.expectCall(address(s_token), abi.encodeWithSelector(expectedSignature, address(s_pool), amount));

    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: amount,
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
