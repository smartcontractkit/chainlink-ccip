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

    s_token.approve(address(s_usdcTokenPool), type(uint256).max);

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
    emit TokenPool.Released(s_routerAllowedOffRamp, recipient, amount);

    Pool.ReleaseOrMintOutV1 memory poolReturnDataV1 = s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        amount: amount,
        localToken: address(s_token),
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
      s_token.balanceOf(address(s_usdcTokenPool)),
      liquidityAmount - amount,
      "Incorrect remaining liquidity in TokenPool"
    );
    assertEq(s_token.balanceOf(recipient), amount, "Tokens not transferred to recipient");
  }

  // https://etherscan.io/tx/0xac9f501fe0b76df1f07a22e1db30929fd12524bc7068d74012dff948632f0883
  function test_incomingMessageWithPrimaryMechanism() public {
    bytes memory encodedUsdcMessage =
      hex"000000000000000300000000000000000000127a00000000000000000000000019330d10d9cc8751218eaf51e8885d058642e08a000000000000000000000000bd3fa81b58ba92a82136038b25adec7066af3155000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000004af08f56978be7dce2d1be3c65c005b41e79401c000000000000000000000000000000000000000000000000000000002057ff7a0000000000000000000000003a23f943181408eac424116af7b7790c94cb97a50000000000000000000000000000000000000000000000000000000000000000000000000000008274119237535fd659626b090f87e365ff89ebc7096bb32e8b0e85f155626b73ae7c4bb2485c184b7cc3cf7909045487890b104efb62ae74a73e32901bdcec91df1bb9ee08ccb014fcbcfe77b74d1263fd4e0b0e8de05d6c9a5913554364abfd5ea768b222f50c715908183905d74044bb2b97527c7e70ae7983c443a603557cac3b1c000000000000000000000000000000000000000000000000000000000000";
    bytes memory attestation = bytes("attestation bytes");

    uint32 nonce = 4730;
    uint32 sourceDomain = 3;
    uint256 amount = 100;

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: nonce, sourceDomain: sourceDomain})),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_token), address(s_usdcTokenPool), amount);

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
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  function test_incomingMessageWithCCTPV2() public {
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = HybridLockReleaseUSDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);

    bytes memory encodedUsdcMessage =
      hex"000000010000000d00000000c51f10ffe8670c7acc9c995bd05747d064960e1680dc69b904c2b2114144ef5000000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d00000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d0000000000000000000000000000000000000000000000000000000000000000000007d0000007d00000000100000000000000000000000029219dd400f2bf60e5a23d13be72b486d4038894000000000000000000000000b38a90f14b24ae81ec0b8f1373694f5b59811d8a00000000000000000000000000000000000000000000000000000045d964b800000000000000000000000000b38a90f14b24ae81ec0b8f1373694f5b59811d8a0000000000000000000000000000000000000000000000000000000059682f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
    bytes memory attestation = bytes("attestation bytes");

    uint32 sourceDomain = 13;
    uint256 amount = 100;

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_token), address(s_usdcTokenPool), amount);

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: encodedUsdcMessage, attestation: attestation}));

    vm.expectCall(
      address(s_mockUSDCTransmitterV2),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, encodedUsdcMessage, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
        sourcePoolData: abi.encode(sourceDomain),
        offchainTokenData: offchainTokenData
      })
    );
  }

  // Reverts
  function test_RevertWhen_UnlockingUSDCFailed() public {
    vm.startPrank(OWNER);

    // Set the transmitter to revert on purpose to simulate a reversion
    s_mockUSDCTransmitterV2.setShouldSucceed(false);

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = HybridLockReleaseUSDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);
    vm.startPrank(s_routerAllowedOffRamp);
    s_mockUSDCTransmitter.setShouldSucceed(false);

    uint256 amount = 13255235235;

    // Format a USDC Message to sent to the transmitter
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDCV2)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: abi.encodePacked(
        uint32(0), // version
        bytes32(uint256(uint160(address(OWNER)))), // burnToken
        bytes32(uint256(uint160(OWNER))), // mintRecipient
        amount, // amount
        bytes32(uint256(uint160(OWNER))), // messageSender
        uint32(0), // maxFee
        uint32(0), // feeExecuted
        block.number + (1 days / 12), // expirationBlock 1-day in the future assuming a 12-second block time.
        "" // hookData
      )
    });

    // Create some source token data
    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(SOURCE_DOMAIN_IDENTIFIER),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // Supply the message and attestation as offChainTokenData
    bytes memory offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );

    vm.expectRevert(USDCTokenPool.UnlockingUSDCFailed.selector);

    // attempts releaseOrMint and revert
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  // TODO: Break into smaller tests for each individual and re-use the usdc message generator function
  function test_RevertWhen_InvalidMessageValues() public {
    vm.startPrank(OWNER);

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);
    versions[0] = HybridLockReleaseUSDCTokenPool.CCTPVersion.VERSION_2;

    // Update the config of the pool to tell it to use CCTP V2 instead of V1
    s_usdcTokenPool.updateCCTPVersion(remoteChainSelectors, versions);
    vm.startPrank(s_routerAllowedOffRamp);
    s_mockUSDCTransmitter.setShouldSucceed(false);

    uint256 amount = 13255235235;

    // Format a USDC Message to sent to the transmitter
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDCV2)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: abi.encodePacked(
        uint32(0), // version
        bytes32(uint256(uint160(address(OWNER)))), // burnToken
        bytes32(uint256(uint160(OWNER))), // mintRecipient
        amount, // amount
        bytes32(uint256(uint160(OWNER))), // messageSender
        uint32(0), // maxFee
        uint32(0), // feeExecuted
        block.number + (1 days / 12), // expirationBlock 1-day in the future assuming a 12-second block time.
        "" // hookData
      )
    });

    // Create some source token data
    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(SOURCE_DOMAIN_IDENTIFIER),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // Change the message version to be incorrect and encode
    usdcMessage.version = 2;
    bytes memory offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 2));

    // attempts releaseOrMint and revert
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );

    // Change the source Domain Identifier to be incorrect and fix the version
    usdcMessage.version = 1;
    usdcMessage.sourceDomain = SOURCE_DOMAIN_IDENTIFIER + 1;
    offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidSourceDomain.selector, SOURCE_DOMAIN_IDENTIFIER, SOURCE_DOMAIN_IDENTIFIER + 1
      )
    );
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );

    // Fix the source domain identifier and change the destination domain identifier
    usdcMessage.sourceDomain = SOURCE_DOMAIN_IDENTIFIER;
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;
    offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, DEST_DOMAIN_IDENTIFIER + 1
      )
    );
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );

    // Fix the Destination domain identifier and change the finality threshold
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;
    usdcMessage.minFinalityThreshold = 1000;
    offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );
    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.InvalidMinFinalityThreshold.selector, 2000, 1000)
    );
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );

    // Fix the minimum threshold and change the executed threshold
    usdcMessage.minFinalityThreshold = 2000;
    usdcMessage.finalityThresholdExecuted = 1000;
    offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );
    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.InvalidExecutionFinalityThreshold.selector, 2000, 1000)
    );
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
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
        amount: amount,
        localToken: address(s_token),
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

    deal(address(s_burnMintERC20), address(s_pool), burnAmount);
    assertEq(s_burnMintERC20.balanceOf(address(s_pool)), burnAmount);

    vm.startPrank(s_burnMintOnRamp);

    // Burn on the source chain and use the Lock-Release Flag
    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_burnMintERC20)
      })
    );

    assertEq(
      bytes4(lockOrBurnOut.destPoolData), LOCK_RELEASE_FLAG, "Incorrect destPoolData, should be the LOCK_RELEASE_FLAG"
    );

    // Assert Burning
    assertEq(s_burnMintERC20.balanceOf(address(s_pool)), 0);
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
    s_token.approve(address(s_usdcTokenPool), type(uint256).max);
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
    emit TokenPool.Released(s_routerAllowedOffRamp, recipient, amount);

    // Release the tokens that were previously locked on mainnet
    Pool.ReleaseOrMintOutV1 memory poolReturnDataV1 = s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: lockOrBurnOut.destPoolData,
        offchainTokenData: ""
      })
    );

    // Assert the tokens were delivered to the recipient
    assertEq(poolReturnDataV1.destinationAmount, amount, "destinationAmount and actual amount transferred differ");
    assertEq(
      s_token.balanceOf(address(s_usdcTokenPool)),
      liquidityAmount - amount,
      "Incorrect remaining liquidity in TokenPool"
    );
    assertEq(s_token.balanceOf(recipient), amount, "Tokens not transferred to recipient");
  }
}
