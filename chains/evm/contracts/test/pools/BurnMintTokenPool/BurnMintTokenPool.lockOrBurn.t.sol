// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintTokenPoolSetup is BurnMintSetup {
  BurnMintTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}

contract BurnMintTokenPool_lockOrBurn is BurnMintTokenPoolSetup {
  function test_constructor() public view {
    assertEq(address(s_token), address(s_pool.getToken()));
    assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
    assertEq(false, s_pool.getAllowListEnabled());
  }

  function test_lockOrBurn_() public {
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

  function test_lockOrBurn_WithFinalityParam() public {
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

    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_FeeNotApplied_LegacyLockOrBurn() public {
    uint16 minBlockConfirmation = 5;
    uint16 defaultFinalityTransferFeeBps = 100;
    uint16 customFinalityTransferFeeBps = 500;
    uint256 amount = 1000e18;

    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(minBlockConfirmation, new TokenPool.CustomFinalityRateLimitConfigArgs[](0));
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 50_000,
        destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
        defaultFinalityFeeUSDCents: 0,
        customFinalityFeeUSDCents: 0,
        defaultFinalityTransferFeeBps: defaultFinalityTransferFeeBps,
        customFinalityTransferFeeBps: customFinalityTransferFeeBps,
        isEnabled: true
      })
    });
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: bytes(""),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(s_tokenPool.getAccumulatedFees(), 0); // No fees should be accumulated
  }

  // Should not burn tokens if cursed.
  function test_lockOrBurn_RevertWhen_CursedByRMN() public {
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
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: 1,
        remoteChainSelector: wrongChainSelector,
        localToken: address(s_token)
      })
    );
  }
}
