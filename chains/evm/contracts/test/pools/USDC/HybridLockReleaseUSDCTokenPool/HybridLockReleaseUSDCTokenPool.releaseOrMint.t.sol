// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../libraries/Internal.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCBridgeMigrator} from "../../../../pools/USDC/USDCBridgeMigrator.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";

import {BurnMintWithLockReleaseFlagTokenPoolSetup} from
  "../../BurnMintWithLockReleaseFlagTokenPool/BurnMintWithLockReleaseFlagTokenPoolSetup.t.sol";
import {HybridLockReleaseUSDCTokenPoolSetup} from "./HybridLockReleaseUSDCTokenPoolSetup.t.sol";

contract HybridLockReleaseUSDCTokenPool_releaseOrMint is HybridLockReleaseUSDCTokenPoolSetup {
  function test_OnLockReleaseMechanism() public {
    address recipient = address(1234);

    // Designate the SOURCE_CHAIN as not using native-USDC, and so the L/R mechanism must be used instead
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = SOURCE_CHAIN_SELECTOR;

    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);

    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(SOURCE_CHAIN_SELECTOR),
      "Lock/Release mech not configured for incoming message from SOURCE_CHAIN_SELECTOR"
    );

    vm.startPrank(OWNER);
    s_usdcTokenPool.setLiquidityProvider(SOURCE_CHAIN_SELECTOR, OWNER);

    // Add 1e12 liquidity so that there's enough to release
    vm.startPrank(s_usdcTokenPool.getLiquidityProvider(SOURCE_CHAIN_SELECTOR));

    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);

    uint256 liquidityAmount = 1e12;
    s_usdcTokenPool.provideLiquidity(SOURCE_CHAIN_SELECTOR, liquidityAmount);

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: 1, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    uint256 amount = 1e6;

    vm.startPrank(s_routerAllowedOffRamp);

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: s_routerAllowedOffRamp,
      recipient: recipient,
      amount: amount
    });

    Pool.ReleaseOrMintOutV1 memory poolReturnDataV1 = s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
        offchainTokenData: ""
      })
    );

    assertEq(poolReturnDataV1.destinationAmount, amount, "destinationAmount and actual amount transferred differ");

    // Simulate the off-ramp forwarding tokens to the recipient on destination chain
    // s_token.transfer(recipient, amount);

    assertEq(
      s_USDCToken.balanceOf(address(s_usdcTokenPool)),
      liquidityAmount - amount,
      "Incorrect remaining liquidity in TokenPool"
    );
    assertEq(s_USDCToken.balanceOf(recipient), amount, "Tokens not transferred to recipient");
  }

  // https://etherscan.io/tx/0x8897ffb613c4d7823ab42c85df2410381e7028c40f0a924530db99412065e338
  function test_incomingMessageWithPrimaryMechanism() public {
    bytes memory encodedUsdcMessage =
      hex"000000010000000500000000bc7e669b9f452229fdd08bac21b6617068f2fd023ad6e805d03afa88b4bb79aea65fc81d0fefa8860cb3b83f089b0224be8a6687b7ae49f594c0b9b4d7e9389300000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d0000000000000000000000000000000000000000000000000000000000000000000007d0000007d000000001c6fa7af3bedbad3a3d65f36aabc97431b1bbe4c2d2f6e0e47ca60203452f5d61000000000000000000000000e7492c49f71841d0f55f4f22c2ee22f02437084000000000000000000000000000000000000000000000000000000017491105202c747e9f0b8a0bb74202136e08fb8463bb15d1ab1d6d3f916f547004d7c7522f0000000000000000000000000000000000000000000000000000000000989a720000000000000000000000000000000000000000000000000000000000989a7200000000000000000000000000000000000000000000000000000000015d0d4e";
    bytes memory attestation = bytes("attestation bytes");

    uint32 nonce = 4730;
    uint32 sourceDomain = 5;
    uint256 amount = 100;

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: nonce, sourceDomain: sourceDomain})),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: encodedUsdcMessage, attestation: attestation}));

    vm.expectCall(
      address(s_mockUSDCTransmitter),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, encodedUsdcMessage, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  function test_RevertWhen_WhileMigrationPause() public {
    address recipient = address(1234);

    // Designate the SOURCE_CHAIN as not using native-USDC, and so the L/R mechanism must be used instead
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = SOURCE_CHAIN_SELECTOR;

    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);

    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(SOURCE_CHAIN_SELECTOR),
      "Lock/Release mech not configured for incoming message from SOURCE_CHAIN_SELECTOR"
    );

    vm.startPrank(OWNER);

    vm.expectEmit();
    emit USDCBridgeMigrator.CCTPMigrationProposed(SOURCE_CHAIN_SELECTOR);

    // Propose the migration to CCTP
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: 1, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory sourcePoolDataLockRelease = abi.encode(LOCK_RELEASE_FLAG);

    uint256 amount = 1e6;

    vm.startPrank(s_routerAllowedOffRamp);

    // Expect revert because the lane is paused and no incoming messages should be allowed
    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.LanePausedForCCTPMigration.selector, SOURCE_CHAIN_SELECTOR)
    );

    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourcePoolDataLockRelease,
        offchainTokenData: ""
      })
    );
  }
}

contract HybridLockReleaseUSDCTokenPool_releaseOrMint_E2ETest is
  HybridLockReleaseUSDCTokenPoolSetup,
  BurnMintWithLockReleaseFlagTokenPoolSetup
{
  function setUp() public override(HybridLockReleaseUSDCTokenPoolSetup, BurnMintWithLockReleaseFlagTokenPoolSetup) {
    HybridLockReleaseUSDCTokenPoolSetup.setUp();
    BurnMintWithLockReleaseFlagTokenPoolSetup.setUp();

    // Designate the SOURCE_CHAIN as not using native-USDC, and so the L/R mechanism must be used instead
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = SOURCE_CHAIN_SELECTOR;
    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);
  }

  function test_releaseOrMint_E2E() public {
    uint256 burnAmount = 20_000e18;

    deal(address(s_token), address(s_pool), burnAmount);
    assertEq(s_token.balanceOf(address(s_pool)), burnAmount);

    vm.startPrank(s_allowedOnRamp);

    // Burn on the source chain and use the Lock-Release Flag
    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(
      bytes4(lockOrBurnOut.destPoolData), LOCK_RELEASE_FLAG, "Incorrect destPoolData, should be the LOCK_RELEASE_FLAG"
    );

    // Assert Burning
    assertEq(s_token.balanceOf(address(s_pool)), 0);
    assertEq(bytes4(lockOrBurnOut.destPoolData), LOCK_RELEASE_FLAG);

    address recipient = address(1234);

    // Assert the chain configuration is correct
    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(SOURCE_CHAIN_SELECTOR),
      "Lock/Release mech not configured for incoming message from SOURCE_CHAIN_SELECTOR"
    );

    // Set the liquidity provider
    vm.startPrank(OWNER);
    s_usdcTokenPool.setLiquidityProvider(SOURCE_CHAIN_SELECTOR, OWNER);

    // Add 1e12 liquidity so that there's enough to release
    vm.startPrank(s_usdcTokenPool.getLiquidityProvider(SOURCE_CHAIN_SELECTOR));
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    uint256 liquidityAmount = 1e12;
    s_usdcTokenPool.provideLiquidity(SOURCE_CHAIN_SELECTOR, liquidityAmount);

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: 1, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    uint256 amount = 1e6;

    vm.startPrank(s_routerAllowedOffRamp);

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: s_routerAllowedOffRamp,
      recipient: recipient,
      amount: amount
    });

    // Release the tokens that were previously locked on mainnet
    Pool.ReleaseOrMintOutV1 memory poolReturnDataV1 = s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: lockOrBurnOut.destPoolData,
        offchainTokenData: ""
      })
    );

    // Assert the tokens were delivered to the recipient
    assertEq(poolReturnDataV1.destinationAmount, amount, "destinationAmount and actual amount transferred differ");
    assertEq(
      s_USDCToken.balanceOf(address(s_usdcTokenPool)),
      liquidityAmount - amount,
      "Incorrect remaining liquidity in TokenPool"
    );
    assertEq(s_USDCToken.balanceOf(recipient), amount, "Tokens not transferred to recipient");
  }
}
