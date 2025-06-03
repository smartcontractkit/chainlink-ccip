// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {HybridLockReleaseUSDCTokenPoolSetup} from "./HybridLockReleaseUSDCTokenPoolSetup.t.sol";

contract HybridLockReleaseUSDCTokenPool_lockOrBurn is HybridLockReleaseUSDCTokenPoolSetup {
  function test_onLockReleaseMechanism() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));

    // Mark the destination chain as supporting CCTP, so use L/R instead.
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = DEST_CHAIN_SELECTOR;

    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);

    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(DEST_CHAIN_SELECTOR),
      "Lock/Release mech not configured for outgoing message to DEST_CHAIN_SELECTOR"
    );

    uint256 amount = 1e6;

    s_token.transfer(address(s_usdcTokenPool), amount);

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_usdcTokenPool)), amount, "Incorrect token amount in the tokenPool");
  }

  function test_PrimaryMechanism() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;

    vm.startPrank(OWNER);

    // Mark outgoing messages as using CCTP V1 primary mechanism
    USDCTokenPool.CCTPVersion[] memory versions = new USDCTokenPool.CCTPVersion[](1);
    versions[0] = USDCTokenPool.CCTPVersion.VERSION_1;
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = DEST_CHAIN_SELECTOR;
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);

    s_token.transfer(address(s_usdcTokenPool), amount);

    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: amount
    });

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      s_mockUSDC.s_nonce(),
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      receiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
  }

  function test_PrimaryMechanism_CCTPV2() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;

    vm.startPrank(OWNER);

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = DEST_CHAIN_SELECTOR;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = USDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);

    s_token.transfer(address(s_usdcTokenPool), amount);

    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    // The event signature for CCTP V2 is different from V1
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      receiver,
      expectedDomain.domainIdentifier,
      s_mockUSDCV2.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      0, // maxFee (should be zero for slow transfer)
      2000, // minFinalityThreshold (Finalized)
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_onLockReleaseMechanism_thenSwitchToPrimary() public {
    // Test Enabling the LR mechanism and sending an outgoing message
    test_PrimaryMechanism();

    // Disable the LR mechanism so that primary CCTP is used and then attempt to send a message
    uint64[] memory destChainRemoves = new uint64[](1);
    destChainRemoves[0] = DEST_CHAIN_SELECTOR;

    vm.startPrank(OWNER);

    vm.expectEmit();
    emit HybridLockReleaseUSDCTokenPool.LockReleaseDisabled(DEST_CHAIN_SELECTOR);

    s_usdcTokenPool.updateChainSelectorMechanisms(destChainRemoves, new uint64[](0));

    // Send an outgoing message
    test_PrimaryMechanism();
  }

  // Reverts
  function test_RevertWhen_WhileMigrationPause() public {
    // Mark the destination chain as supporting CCTP, so use L/R instead.
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = DEST_CHAIN_SELECTOR;

    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);

    // Create a fake migration proposal
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), DEST_CHAIN_SELECTOR);

    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));

    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(DEST_CHAIN_SELECTOR),
      "Lock Release mech not configured for outgoing message to DEST_CHAIN_SELECTOR"
    );

    uint256 amount = 1e6;

    s_token.transfer(address(s_usdcTokenPool), amount);

    vm.startPrank(s_routerAllowedOnRamp);

    // Expect the lockOrBurn to fail because a pending CCTP-Migration has paused outgoing messages on CCIP
    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.LanePausedForCCTPMigration.selector, DEST_CHAIN_SELECTOR)
    );

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_InvalidReceiver() public {
    vm.startPrank(OWNER);

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = DEST_CHAIN_SELECTOR;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = USDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);

    vm.startPrank(s_routerAllowedOnRamp);

    // Generate a 33-byte long string to use as invalid receiver for CCTP V2
    // because only a 32-byte receiver can be manually decoded correctly.
    bytes memory invalidReceiver = abi.encode(keccak256("0xCLL"), "A");

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidReceiver.selector, invalidReceiver));
    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: invalidReceiver,
        amount: 1000,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_UnknownDomain() public {
    vm.startPrank(OWNER);

    uint64 wrongDomain = DEST_CHAIN_SELECTOR + 1;

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = wrongDomain;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = USDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);

    // We need to setup the wrong chainSelector so it reaches the domain check
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: wrongDomain, onRamp: s_routerAllowedOnRamp});
    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: wrongDomain,
      remotePoolAddresses: new bytes[](0),
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_usdcTokenPool.applyChainUpdates(new uint64[](0), chainUpdates);

    uint256 amount = 1000;
    vm.startPrank(s_routerAllowedOnRamp);
    deal(address(s_token), s_routerAllowedOnRamp, amount);
    s_token.approve(address(s_usdcTokenPool), amount);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.UnknownDomain.selector, wrongDomain));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(address(0)),
        amount: amount,
        remoteChainSelector: wrongDomain,
        localToken: address(s_token)
      })
    );
  }
}
