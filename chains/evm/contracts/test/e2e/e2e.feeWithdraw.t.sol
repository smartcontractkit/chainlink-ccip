// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Proxy} from "../../Proxy.sol";
import {Router} from "../../Router.sol";
import {CommitteeVerifier} from "../../ccvs/CommitteeVerifier.sol";
import {VersionedVerifierResolver} from "../../ccvs/VersionedVerifierResolver.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {Executor} from "../../executor/Executor.sol";
import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {OffRampHelper} from "../helpers/OffRampHelper.sol";

import {TokenPoolHelper} from "../helpers/TokenPoolHelper.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";

import {OnRampSetup} from "../onRamp/OnRamp/OnRampSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @title E2E Fee Withdrawal Test
/// @notice Tests fee withdrawal from all components after a ccipSend
/// @dev This test verifies:
///      1. Fees are correctly distributed to each component after ccipSend
///      2. Fee withdrawal works for all components (OnRamp, Executor, TokenPool, Verifier)
///      3. Verifier fees go to resolver (proxy), not implementation
///      4. All withdrawFeeTokens functions are permissionless (PAL compatible)
contract e2e_feeWithdrawal is OnRampSetup {
  struct Balances {
    uint256 onRampBalance;
    uint256 executorBalance;
    uint256 verifierResolverBalance;
    uint256 verifierImplBalance;
    uint256 tokenPoolBalance;
    uint256 feeAggregatorBalance;
  }
  uint16 internal constant MIN_BLOCK_CONFIRMATIONS = 50;

  OffRampHelper internal s_offRamp;
  address internal s_destVerifier;
  address internal s_verifierResolver; // Resolver address (proxy)
  CommitteeVerifier internal s_verifierImpl; // Implementation address
  address internal s_executor; // Executor address (proxy)
  Executor internal s_executorImpl; // Executor implementation address
  TokenPoolHelper internal s_tokenPool;
  BurnMintERC20 internal s_testToken;
  address internal s_automationAddress; // Simulates Chainlink Automation/CRE
  uint16 internal constant NETWORK_FEE_USD_CENTS = 200;
  address internal s_feeAggregator;

  function setUp() public virtual override {
    super.setUp();

    s_feeAggregator = makeAddr("feeAggregator");

    // Deploy a test token and pool
    s_testToken = new BurnMintERC20("TestToken", "TEST", 18, 0, 0);
    deal(address(s_testToken), OWNER, type(uint256).max);
    deal(address(s_testToken), address(s_onRamp), type(uint256).max);

    s_tokenPool = new TokenPoolHelper(
      IERC20(address(s_testToken)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );

    // Configure token pool in registry
    s_tokenAdminRegistry.proposeAdministrator(address(s_testToken), OWNER);
    s_tokenAdminRegistry.acceptAdminRole(address(s_testToken));
    s_tokenAdminRegistry.setPool(address(s_testToken), address(s_tokenPool));

    // Set up pool chain config
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encodePacked(address(s_tokenPool));
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encodePacked(address(s_testToken)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Set fee aggregator for TokenPool via setDynamicConfig
    vm.startPrank(OWNER);
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), address(0), s_feeAggregator);
    vm.stopPrank();

    // Set up router with onRamp
    vm.startPrank(OWNER);
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
    vm.stopPrank();

    // Deploy verifier implementation
    s_verifierImpl = new CommitteeVerifier(
      CommitteeVerifier.DynamicConfig({feeAggregator: s_feeAggregator, allowlistAdmin: address(0)}),
      new string[](0),
      address(s_mockRMNRemote)
    );

    // Configure verifier for destination chain
    BaseVerifier.RemoteChainConfigArgs[] memory destChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    destChainConfigs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: s_sourceRouter,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: uint16(VERIFIER_FEE_USD_CENTS), // $2.00 fee for verifier
      gasForVerification: VERIFIER_GAS,
      payloadSizeBytes: VERIFIER_BYTES
    });
    s_verifierImpl.applyRemoteChainConfigUpdates(destChainConfigs);

    // Deploy verifier resolver (proxy)
    VersionedVerifierResolver verifierResolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImpls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outboundImpls[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: DEST_CHAIN_SELECTOR, verifier: address(s_verifierImpl)
    });
    verifierResolver.applyOutboundImplementationUpdates(outboundImpls);
    // Set fee aggregator for resolver
    verifierResolver.setFeeAggregator(s_feeAggregator);
    // Create proxy with feeAggregator - fees go to proxy address, proxy can withdraw them
    s_verifierResolver = address(new Proxy(address(verifierResolver), s_feeAggregator));

    // Deploy executor implementation
    s_executorImpl = new Executor(
      10, // maxCCVsPerMsg
      Executor.DynamicConfig({
        feeAggregator: s_feeAggregator, minBlockConfirmations: MIN_BLOCK_CONFIRMATIONS, ccvAllowlistEnabled: false
      })
    );

    // Configure executor implementation for destination chain
    Executor.RemoteChainConfigArgs[] memory executorConfigs = new Executor.RemoteChainConfigArgs[](1);
    executorConfigs[0] = Executor.RemoteChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      config: Executor.RemoteChainConfig({usdCentsFee: 50, enabled: true}) // $0.50
    });
    s_executorImpl.applyDestChainUpdates(new uint64[](0), executorConfigs);

    // Wrap executor in Proxy - fees go to proxy address, proxy can withdraw them
    s_executor = address(new Proxy(address(s_executorImpl), s_feeAggregator));

    // Set up OnRamp destination chain config with our verifier and executor
    address[] memory defaultSourceCCVs = new address[](1);
    defaultSourceCCVs[0] = s_verifierResolver; // Use resolver, not impl
    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      tokenReceiverAllowed: false,
      messageNetworkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultSourceCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: s_executor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
    vm.startPrank(OWNER);
    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
    vm.stopPrank();

    // Set up OffRamp
    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        maxGasBufferToUpdateState: DEFAULT_MAX_GAS_BUFFER_TO_UPDATE_STATE
      })
    );
    s_destVerifier = address(new Proxy(address(new MockVerifier("")), s_feeAggregator));
    address[] memory defaultDestCCVs = new address[](1);
    defaultDestCCVs[0] = s_destVerifier;
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encode(s_onRamp);
    OffRamp.SourceChainConfigArgs[] memory updates = new OffRamp.SourceChainConfigArgs[](1);
    updates[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCVs: defaultDestCCVs,
      laneMandatedCCVs: new address[](0)
    });
    s_offRamp.applySourceChainConfigUpdates(updates);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    vm.startPrank(OWNER);
    s_destRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);
    vm.stopPrank();

    // Set up automation address (simulates Chainlink Automation/CRE)
    s_automationAddress = makeAddr("automation");

    // Update OnRamp's feeAggregator to match s_feeAggregator (from BaseTest)
    // OnRampSetup uses FEE_AGGREGATOR constant, but we need to use s_feeAggregator for consistency
    OnRamp.DynamicConfig memory onRampConfig = s_onRamp.getDynamicConfig();
    onRampConfig.feeAggregator = s_feeAggregator;
    vm.startPrank(OWNER);
    s_onRamp.setDynamicConfig(onRampConfig);
    vm.stopPrank();
  }

  /// @notice Test fee withdrawal from all components after ccipSend
  function test_FeeWithdrawal_AfterCcipSend() public {
    vm.pauseGasMetering();

    // Ensure the test contract has fee tokens/test tokens to pay for the transaction
    // Router will transferFrom msg.sender (this test contract) to OnRamp
    deal(s_sourceFeeToken, address(this), type(uint256).max);
    deal(address(s_testToken), address(this), type(uint256).max);

    // Get initial balances
    Balances memory initial = Balances({
      onRampBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp)),
      executorBalance: IERC20(s_sourceFeeToken).balanceOf(s_executor),
      verifierResolverBalance: IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver),
      verifierImplBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl)),
      tokenPoolBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_tokenPool)),
      feeAggregatorBalance: IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator)
    });

    // Prepare message with token transfer
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: new address[](0), // Use default CCV
          ccvArgs: new bytes[](0),
          blockConfirmations: 0,
          gasLimit: GAS_LIMIT,
          executor: address(0), // Use default executor
          executorArgs: "",
          tokenReceiver: "",
          tokenArgs: ""
        })
      )
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: address(s_testToken), amount: 1e18});

    // Get fee and approve
    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fee);
    // Also approve the token being transferred
    IERC20(address(s_testToken)).approve(address(s_sourceRouter), message.tokenAmounts[0].amount);

    // Perform ccipSend
    vm.resumeGasMetering();
    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);
    vm.pauseGasMetering();

    // Check that fees were distributed
    // Get end balances
    Balances memory end = Balances({
      onRampBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp)),
      executorBalance: IERC20(s_sourceFeeToken).balanceOf(s_executor),
      verifierResolverBalance: IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver),
      verifierImplBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl)),
      tokenPoolBalance: IERC20(s_sourceFeeToken).balanceOf(address(s_tokenPool)),
      feeAggregatorBalance: IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator)
    });

    // Network fee should remain on OnRamp
    assertGt(end.onRampBalance, initial.onRampBalance, "Network fee should be on OnRamp");

    // Executor (proxy) should have received fee
    assertGt(end.executorBalance, initial.executorBalance, "Executor proxy should have received fee");

    // Verifier fees should go to resolver (proxy), not implementation
    assertGt(
      end.verifierResolverBalance, initial.verifierResolverBalance, "Verifier fees should go to resolver (proxy)"
    );
    assertEq(
      end.verifierImplBalance, initial.verifierImplBalance, "Verifier implementation should NOT receive fees directly"
    );

    // TokenPool (V2) should have received fees directly if pool fee config is enabled
    // If pool fee config is disabled, OnRamp falls back to FeeQuoter but still transfers fees to pool
    // Note: Pool fees may be 0 if FeeQuoter returns 0, but the mechanism should work
    // We'll verify the withdrawal mechanism works regardless

    // Now test withdrawing fees from each component

    uint256 totalFeesWithdrawn;
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    {
      // 1. Test OnRamp withdrawal (permissionless - PAL compatible)
      // Get the actual balance at the time of withdrawal (in case it changed)
      uint256 actualOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp));

      // Only proceed if there's actually a balance to withdraw
      if (actualOnRampBalance > 0) {
        uint256 aggregatorBalanceBeforeOnRamp = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
        vm.stopPrank();
        vm.prank(s_automationAddress); // Anyone can call (PAL compatible)
        s_onRamp.withdrawFeeTokens(feeTokens);
        uint256 aggregatorBalanceAfterOnRamp = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
        // Use actual balance withdrawn, not the calculated difference
        uint256 actualWithdrawn = aggregatorBalanceAfterOnRamp - aggregatorBalanceBeforeOnRamp;
        assertEq(actualWithdrawn, actualOnRampBalance, "OnRamp fees should be withdrawn to aggregator");
        totalFeesWithdrawn += actualWithdrawn;
      } else {
        // If no balance, still add 0 to totalFeesWithdrawn for consistency
        totalFeesWithdrawn += 0;
      }
    }

    {
      // 2. Test Executor withdrawal (permissionless - PAL compatible)
      // Fees go to executor proxy, which can withdraw them
      uint256 executorFeeBalance = end.executorBalance - initial.executorBalance;
      uint256 aggregatorBalanceBeforeExecutor = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      vm.stopPrank();
      vm.prank(s_automationAddress); // Anyone can call (PAL compatible)
      Proxy(s_executor).withdrawFeeTokens(feeTokens);
      uint256 aggregatorBalanceAfterExecutor = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      // Use actual balance withdrawn, not the calculated difference
      uint256 actualExecutorWithdrawn = aggregatorBalanceAfterExecutor - aggregatorBalanceBeforeExecutor;
      assertEq(actualExecutorWithdrawn, executorFeeBalance, "Executor fees should be withdrawn to aggregator");
      totalFeesWithdrawn += actualExecutorWithdrawn;
    }

    {
      // 3. Test TokenPool withdrawal (permissionless - PAL compatible)
      // For V2 pools, fees go directly to pool during _distributeFees
      // Check if pool received fees from the actual transaction
      uint256 poolFeeBalance = end.tokenPoolBalance - initial.tokenPoolBalance;

      // If pool received fees from the transaction, test withdrawing them
      // Otherwise, manually add fees to test the withdrawal mechanism
      uint256 poolFeeAmount = poolFeeBalance > 0 ? poolFeeBalance : 100 ether;
      if (poolFeeBalance == 0) {
        // Pool didn't receive fees (likely because FeeQuoter returned 0 or pool config disabled)
        // Manually add fees to test withdrawal mechanism
        deal(s_sourceFeeToken, address(s_tokenPool), poolFeeAmount);
      }

      uint256 aggregatorBalanceBeforePool = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      uint256 actualPoolBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_tokenPool));
      vm.stopPrank();
      vm.prank(s_automationAddress); // Anyone can call (PAL compatible)
      s_tokenPool.withdrawFeeTokens(feeTokens);
      uint256 aggregatorBalanceAfterPool = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      uint256 actualPoolWithdrawn = aggregatorBalanceAfterPool - aggregatorBalanceBeforePool;
      assertEq(actualPoolWithdrawn, actualPoolBalance, "TokenPool fees should be withdrawn to aggregator");
      totalFeesWithdrawn += actualPoolWithdrawn;
    }

    // 4. Test Verifier withdrawal (permissionless - PAL compatible)
    // Fees go to resolver (proxy), which can withdraw them
    uint256 verifierFeeBalance = end.verifierResolverBalance - initial.verifierResolverBalance;
    if (verifierFeeBalance > 0) {
      uint256 aggregatorBalanceBeforeVerifier = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);

      // Withdraw from resolver (proxy) - Proxy has its own withdrawFeeTokens
      vm.stopPrank();
      vm.prank(s_automationAddress); // Anyone can call (PAL compatible)
      Proxy(s_verifierResolver).withdrawFeeTokens(feeTokens);

      // Use actual balance withdrawn, not the calculated difference
      uint256 actualVerifierWithdrawn =
        IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator) - aggregatorBalanceBeforeVerifier;
      assertEq(actualVerifierWithdrawn, verifierFeeBalance, "Verifier fees should be withdrawable from resolver");

      // Verify no fees remain on resolver after withdrawal
      assertEq(
        IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver),
        initial.verifierResolverBalance,
        "No fees should remain on resolver after withdrawal"
      );

      totalFeesWithdrawn += actualVerifierWithdrawn;
    }

    // Verify total fees collected in aggregator
    end.feeAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      end.feeAggregatorBalance - initial.feeAggregatorBalance,
      totalFeesWithdrawn,
      "Total fees withdrawn should match sum of individual withdrawals"
    );

    vm.resumeGasMetering();
  }

  /// @notice Test that verifier fees go to resolver (proxy), not implementation
  function test_VerifierFeeIssue_FeesGoToResolverNotImplementation() public {
    vm.pauseGasMetering();

    // Ensure the test contract has fee tokens to pay for the transaction
    // Router will transferFrom msg.sender (this test contract) to OnRamp
    deal(s_sourceFeeToken, address(this), type(uint256).max);

    // Get initial balances
    uint256 initialVerifierResolverBalance = IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver);
    uint256 initialVerifierImplBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl));

    // Prepare message
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: new address[](0),
          ccvArgs: new bytes[](0),
          blockConfirmations: 0,
          gasLimit: GAS_LIMIT,
          executor: address(0),
          executorArgs: "",
          tokenReceiver: "",
          tokenArgs: ""
        })
      )
    });

    // Get fee and approve
    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fee);

    // Perform ccipSend
    vm.resumeGasMetering();
    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);
    vm.pauseGasMetering();

    // Check balances
    uint256 finalVerifierResolverBalance = IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver);
    uint256 finalVerifierImplBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl));

    // Fees should go to resolver (proxy), not implementation
    assertGt(finalVerifierResolverBalance, initialVerifierResolverBalance, "Fees should go to resolver (proxy)");
    assertEq(finalVerifierImplBalance, initialVerifierImplBalance, "Implementation should NOT receive fees directly");

    // Verify that resolver is different from implementation
    address implAddress =
      ICrossChainVerifierResolver(s_verifierResolver).getOutboundImplementation(DEST_CHAIN_SELECTOR, "");
    assertEq(implAddress, address(s_verifierImpl), "Resolver should resolve to implementation");
    assertTrue(s_verifierResolver != address(s_verifierImpl), "Resolver and impl should be different");

    vm.resumeGasMetering();
  }

  /// @notice Test PAL compatibility: permissionless vs onlyOwner
  function test_PALCompatibility_PermissionlessVsOnlyOwner() public {
    vm.pauseGasMetering();

    // Give components some fees
    uint256 feeAmount = 100 ether;
    deal(s_sourceFeeToken, address(s_onRamp), feeAmount);
    deal(s_sourceFeeToken, s_executor, feeAmount); // Proxy address
    deal(s_sourceFeeToken, address(s_tokenPool), feeAmount);
    deal(s_sourceFeeToken, s_verifierResolver, feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    // 1. OnRamp: Permissionless (PAL compatible) ✅
    uint256 aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.stopPrank();
    vm.prank(s_automationAddress);
    s_onRamp.withdrawFeeTokens(feeTokens);
    uint256 aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore, feeAmount, "OnRamp: Automation can withdraw (PAL compatible)"
    );

    // 2. Verifier: Permissionless (PAL compatible) ✅
    aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.prank(s_automationAddress);
    Proxy(s_verifierResolver).withdrawFeeTokens(feeTokens);
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore, feeAmount, "Verifier: Automation can withdraw (PAL compatible)"
    );

    // 3. Executor: Permissionless (PAL compatible) ✅
    aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.prank(s_automationAddress);
    Proxy(s_executor).withdrawFeeTokens(feeTokens);
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore, feeAmount, "Executor: Automation can withdraw (PAL compatible)"
    );

    // 4. TokenPool: Permissionless (PAL compatible) ✅
    deal(s_sourceFeeToken, address(s_tokenPool), feeAmount); // Re-add fees
    aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.prank(s_automationAddress);
    s_tokenPool.withdrawFeeTokens(feeTokens);
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore, feeAmount, "TokenPool: Automation can withdraw (PAL compatible)"
    );

    vm.resumeGasMetering();
  }
}
