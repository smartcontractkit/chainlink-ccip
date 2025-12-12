// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/interfaces/IERC20.sol";

contract SiloedLockReleaseTokenPool_releaseOrMint is SiloedLockReleaseTokenPoolSetup {
  uint256 internal constant AMOUNT = 10e18;
  uint16 internal constant CUSTOM_FINALITY = 1;

  function setUp() public override {
    super.setUp();

    IERC20(address(s_token)).approve(address(s_lockBox), type(uint256).max);

    s_lockBox.deposit(address(s_token), AMOUNT);
    s_lockBox.deposit(address(s_token), AMOUNT);
  }

  function test_ReleaseOrMint_SiloedChain() public {
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.startPrank(s_allowedOnRamp);

    // Lock funds so that they can be released without underflowing the internal accounting
    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(SILOED_CHAIN_SELECTOR), AMOUNT);

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_lockBox), OWNER, AMOUNT);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: AMOUNT,
        localToken: address(s_token),
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(SILOED_CHAIN_SELECTOR), 0);
  }

  function test_ReleaseOrMint_UnsiloedChain() public {
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);
    vm.startPrank(s_allowedOnRamp);

    // Lock funds for unsiloed chain so they can be released later
    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(SOURCE_CHAIN_SELECTOR), AMOUNT);
    assertEq(s_siloedLockReleaseTokenPool.getUnsiloedLiquidity(), AMOUNT);

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_lockBox), OWNER, AMOUNT);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: AMOUNT,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(SOURCE_CHAIN_SELECTOR), 0);
    assertEq(s_siloedLockReleaseTokenPool.getUnsiloedLiquidity(), 0);
  }

  // Reverts

  function test_releaseOrMint_V2_UsesCustomFinalityAndNetLiquidity() public {
    uint16 feeBps = 500;
    uint256 expectedLockedAmount = AMOUNT - (AMOUNT * feeBps) / 10_000;

    uint256 startingLockBoxBalance = s_token.balanceOf(address(s_lockBox));

    _setTokenTransferFee(SOURCE_CHAIN_SELECTOR, feeBps);
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.startPrank(s_allowedOnRamp);
    (, uint256 lockedAmount) = s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(lockedAmount, expectedLockedAmount);
    assertEq(s_token.balanceOf(address(s_lockBox)), startingLockBoxBalance + expectedLockedAmount);
    assertEq(s_siloedLockReleaseTokenPool.getUnsiloedLiquidity(), expectedLockedAmount);

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit TokenPool.CustomBlockConfirmationInboundRateLimitConsumed({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: address(s_token),
      amount: expectedLockedAmount
    });

    vm.expectEmit();
    emit IERC20.Transfer(address(s_lockBox), OWNER, expectedLockedAmount);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: expectedLockedAmount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      }),
      CUSTOM_FINALITY
    );

    assertEq(s_siloedLockReleaseTokenPool.getUnsiloedLiquidity(), 0);
    assertEq(s_token.balanceOf(address(s_lockBox)), startingLockBoxBalance);
    assertEq(s_token.balanceOf(address(s_siloedLockReleaseTokenPool)), AMOUNT - expectedLockedAmount);
  }

  function test_ReleaseOrMint_RevertsWhen_InsufficientLiquidity_SiloedChain() public {
    uint256 releaseAmount = AMOUNT;
    uint256 liquidityAmount = releaseAmount - 1;

    s_siloedLockReleaseTokenPool.provideSiloedLiquidity(SILOED_CHAIN_SELECTOR, liquidityAmount);

    // Since amount to release is greater than provided liquidity, the function should revert
    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InsufficientLiquidity.selector, liquidityAmount, releaseAmount)
    );

    vm.startPrank(s_allowedOffRamp);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: releaseAmount,
        localToken: address(s_token),
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );
  }

  function test_ReleaseOrMint_RevertsWhen_InsufficientLiquidity_UnsiloedChain() public {
    uint256 releaseAmount = AMOUNT;
    uint256 liquidityAmount = releaseAmount - 1;

    // Call the provide liquidity function which provides to unsiloed chains.
    s_siloedLockReleaseTokenPool.provideLiquidity(liquidityAmount);

    // Since amount to release is greater than provided liquidity, the function should revert
    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InsufficientLiquidity.selector, liquidityAmount, releaseAmount)
    );

    vm.startPrank(s_allowedOffRamp);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: releaseAmount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );
  }
}
