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
import {Internal} from "../../libraries/Internal.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {OffRampHelper} from "../helpers/OffRampHelper.sol";

import {TokenPoolHelper} from "../helpers/TokenPoolHelper.sol";
import {MockExecutor} from "../mocks/MockExecutor.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";

import {OnRampSetup} from "../onRamp/OnRamp/OnRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @title E2E Fee Withdrawal Test
/// @notice Tests fee withdrawal from all components after a ccipSend
/// @dev This test verifies:
///      1. Fees are correctly distributed to each component after ccipSend
///      2. Fee withdrawal works for all components (OnRamp, Executor, TokenPool, Verifier)
///      3. Verifier fee issue: fees go to resolver, not implementation
///      4. PAL compatibility: permissionless vs onlyOwner restrictions
contract e2e_feeWithdrawal is OnRampSetup {
  uint16 internal constant MIN_BLOCK_CONFIRMATIONS = 50;

  OffRampHelper internal s_offRamp;
  address internal s_destVerifier;
  address internal s_verifierResolver; // Resolver address (proxy)
  CommitteeVerifier internal s_verifierImpl; // Implementation address
  Executor internal s_executor;
  TokenPoolHelper internal s_tokenPool;
  BurnMintERC20 internal s_testToken;
  address internal s_automationAddress; // Simulates Chainlink Automation/CRE

  function setUp() public virtual override {
    super.setUp();

    // Deploy a test token and pool
    s_testToken = new BurnMintERC20("TestToken", "TEST", 18, 0, 0);
    deal(address(s_testToken), OWNER, type(uint256).max);
    deal(address(s_testToken), address(s_onRamp), type(uint256).max);

    s_tokenPool = new TokenPoolHelper(
      IERC20(address(s_testToken)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      s_feeAggregator
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

    // Set up router with onRamp
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

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
      feeUSDCents: uint16(VERIFIER_FEE_USD_CENTS), //TODO this is defined at 200, ie $2.00 but maybe I shouold define it here?
      gasForVerification: VERIFIER_GAS,
      payloadSizeBytes: VERIFIER_BYTES
    });
    s_verifierImpl.applyRemoteChainConfigUpdates(destChainConfigs);

    // Deploy verifier resolver (proxy)
    VersionedVerifierResolver verifierResolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImpls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outboundImpls[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      verifier: address(s_verifierImpl)
    });
    verifierResolver.applyOutboundImplementationUpdates(outboundImpls);
    s_verifierResolver = address(new Proxy(address(verifierResolver)));

    // Deploy executor
    s_executor = new Executor(
      10, // maxCCVsPerMsg
      Executor.DynamicConfig({
        feeAggregator: s_feeAggregator,
        minBlockConfirmations: MIN_BLOCK_CONFIRMATIONS,
        ccvAllowlistEnabled: false
      })
    );

    // Configure executor for destination chain
    Executor.RemoteChainConfigArgs[] memory executorConfigs = new Executor.RemoteChainConfigArgs[](1);
    executorConfigs[0] = Executor.RemoteChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      config: Executor.RemoteChainConfig({usdCentsFee: 50, enabled: true}) // $0.50
    });
    s_executor.applyDestChainUpdates(new uint64[](0), executorConfigs);

    // Set up OnRamp destination chain config with our verifier and executor
    address[] memory defaultSourceCCVs = new address[](1);
    defaultSourceCCVs[0] = s_verifierResolver; // Use resolver, not impl
    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultSourceCCVs,
      defaultExecutor: address(s_executor),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

    // Set up OffRamp
    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      })
    );
    s_destVerifier = address(new Proxy(address(new MockVerifier(""))));
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
    s_destRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    // Set up automation address (simulates Chainlink Automation/CRE)
    s_automationAddress = makeAddr("automation");

    // Update OnRamp's feeAggregator to match s_feeAggregator (from BaseTest)
    // OnRampSetup uses FEE_AGGREGATOR constant, but we need to use s_feeAggregator for consistency
    OnRamp.DynamicConfig memory onRampConfig = s_onRamp.getDynamicConfig();
    onRampConfig.feeAggregator = s_feeAggregator;
    vm.startPrank(OWNER);
    s_onRamp.setDynamicConfig(onRampConfig);
    vm.stopPrank();

    // Seed balances to avoid cold storage costs
    deal(s_sourceFeeToken, address(s_verifierImpl), 1);
    deal(s_sourceFeeToken, address(s_verifierResolver), 1);
    deal(s_sourceFeeToken, address(s_executor), 1);
  }

  /// @notice Test fee withdrawal from all components after ccipSend
  function test_FeeWithdrawal_AfterCcipSend() public {
    vm.pauseGasMetering();

    // Get initial balances
    uint256 initialOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp));
    uint256 initialExecutorBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_executor));
    uint256 initialVerifierResolverBalance = IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver);
    uint256 initialVerifierImplBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl));
    uint256 initialFeeAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);

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
    uint256 finalOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_onRamp));
    uint256 finalExecutorBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_executor));
    uint256 finalVerifierResolverBalance = IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver);
    uint256 finalVerifierImplBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_verifierImpl));

    // Network fee should remain on OnRamp
    assertGt(finalOnRampBalance, initialOnRampBalance, "Network fee should be on OnRamp");

    // Executor should have received fee
    assertGt(finalExecutorBalance, initialExecutorBalance, "Executor should have received fee");

    // FIXED: Verifier fees should go to implementation, not resolver
    // After fix: fees should go directly to implementation OR resolver should be able to withdraw them
    assertEq(
      finalVerifierResolverBalance,
      initialVerifierResolverBalance,
      "Verifier fees should NOT go to resolver (should go to implementation instead)"
    );
    assertGt(
      finalVerifierImplBalance, initialVerifierImplBalance, "Verifier implementation should receive fees directly"
    );

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
      // 2. Test Executor withdrawal (onlyOwner - NOT PAL compatible)
      uint256 executorFeeBalance = finalExecutorBalance - initialExecutorBalance;
      uint256 aggregatorBalanceBeforeExecutor = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      vm.startPrank(OWNER);
      s_executor.withdrawFeeTokens(feeTokens);
      uint256 aggregatorBalanceAfterExecutor = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      // Note: Using assertApproxEqAbs with tolerance of 1 wei due to integer division rounding
      // in fee conversion (OnRamp._getReceipts line 949 and 955). The executor fee is calculated
      // as: (baseFee * multiplier * 1e32 / price) + (execCost * 1e34 / price), where each
      // division truncates, potentially causing a 1 wei difference.
      assertApproxEqAbs(
        aggregatorBalanceAfterExecutor - aggregatorBalanceBeforeExecutor,
        executorFeeBalance,
        1,
        "Executor fees should be withdrawn to aggregator"
      );
      vm.stopPrank();
      totalFeesWithdrawn += executorFeeBalance;

      // FIXED: Test that automation CAN withdraw from Executor (PAL compatibility)
      // After fix: Executor should be permissionless or have a role-based access
      deal(s_sourceFeeToken, address(s_executor), executorFeeBalance); // Re-add fees for testing
      uint256 aggregatorBalanceBeforeExecutorAuto = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      vm.prank(s_automationAddress);
      s_executor.withdrawFeeTokens(feeTokens); // Should succeed after fix
      uint256 aggregatorBalanceAfterExecutorAuto = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      assertApproxEqAbs(
        aggregatorBalanceAfterExecutorAuto - aggregatorBalanceBeforeExecutorAuto,
        executorFeeBalance,
        1,
        "Executor: Automation should be able to withdraw (PAL compatible after fix)"
      );
    }

    {
      // 3. Test TokenPool withdrawal (onlyOwner - NOT PAL compatible)
      // First, give token pool some fees (simulating V1 pool behavior where fees stay on OnRamp)
      // For V2 pools, fees go directly to pool, but we'll test withdrawal anyway
      uint256 poolFeeAmount = 100 ether;
      deal(s_sourceFeeToken, address(s_tokenPool), poolFeeAmount);
      uint256 aggregatorBalanceBeforePool = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      vm.startPrank(OWNER);
      s_tokenPool.withdrawFeeTokens(feeTokens);
      uint256 aggregatorBalanceAfterPool = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      assertEq(
        aggregatorBalanceAfterPool - aggregatorBalanceBeforePool,
        poolFeeAmount,
        "TokenPool fees should be withdrawn to aggregator"
      );
      vm.stopPrank();
      totalFeesWithdrawn += poolFeeAmount;

      // FIXED: Test that automation CAN withdraw from TokenPool (PAL compatibility)
      // After fix: TokenPool should be permissionless or have a role-based access
      deal(s_sourceFeeToken, address(s_tokenPool), poolFeeAmount); // Re-add fees for testing
      uint256 aggregatorBalanceBeforePoolAuto = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      vm.prank(s_automationAddress);
      s_tokenPool.withdrawFeeTokens(feeTokens); // Should succeed after fix
      uint256 aggregatorBalanceAfterPoolAuto = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      assertEq(
        aggregatorBalanceAfterPoolAuto - aggregatorBalanceBeforePoolAuto,
        poolFeeAmount,
        "TokenPool: Automation should be able to withdraw (PAL compatible after fix)"
      );
    }

    // 4. Test Verifier withdrawal (permissionless - PAL compatible)
    // FIXED: After fix, fees should go to implementation OR resolver should be able to withdraw
    uint256 verifierFeeBalance = finalVerifierImplBalance - initialVerifierImplBalance;
    if (verifierFeeBalance > 0) {
      // After fix: Either fees go directly to implementation (which can withdraw),
      // or resolver has withdrawal capability
      uint256 aggregatorBalanceBeforeVerifier = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);

      // Try to withdraw from implementation (if fees went there)
      vm.prank(s_automationAddress);
      s_verifierImpl.withdrawFeeTokens(feeTokens);

      uint256 aggregatorBalanceAfterVerifier = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
      assertApproxEqAbs(
        aggregatorBalanceAfterVerifier - aggregatorBalanceBeforeVerifier,
        verifierFeeBalance,
        1,
        "Verifier fees should be withdrawable (either from impl or resolver)"
      );

      // Verify no fees are stuck on resolver
      assertEq(
        IERC20(s_sourceFeeToken).balanceOf(s_verifierResolver),
        initialVerifierResolverBalance,
        "No fees should be stuck on resolver after fix"
      );
    }

    // Verify total fees collected in aggregator
    uint256 finalFeeAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    // Note: Using assertApproxEqAbs with tolerance of 1 wei due to cumulative rounding errors
    // from integer division in fee conversion calculations across all components.
    assertApproxEqAbs(
      finalFeeAggregatorBalance - initialFeeAggregatorBalance,
      totalFeesWithdrawn,
      1,
      "Total fees withdrawn should match sum of individual withdrawals"
    );

    vm.resumeGasMetering();
  }

  /// @notice Test that verifier fees go to implementation, not resolver (verifies the fix)
  function test_VerifierFeeIssue_FeesGoToImplementationNotResolver() public {
    vm.pauseGasMetering();

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

    // FIXED: Fees should go to implementation, not resolver
    assertEq(
      finalVerifierResolverBalance, initialVerifierResolverBalance, "Fees should NOT go to resolver (proxy) after fix"
    );
    assertGt(
      finalVerifierImplBalance, initialVerifierImplBalance, "Implementation SHOULD receive fees directly after fix"
    );

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
    deal(s_sourceFeeToken, address(s_executor), feeAmount);
    deal(s_sourceFeeToken, address(s_tokenPool), feeAmount);
    deal(s_sourceFeeToken, address(s_verifierImpl), feeAmount);

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
    s_verifierImpl.withdrawFeeTokens(feeTokens);
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore, feeAmount, "Verifier: Automation can withdraw (PAL compatible)"
    );

    // 3. Executor: Should be PAL compatible after fix ✅
    aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.prank(s_automationAddress);
    s_executor.withdrawFeeTokens(feeTokens); // Should succeed after fix
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore,
      feeAmount,
      "Executor: Automation should be able to withdraw (PAL compatible after fix)"
    );

    // 4. TokenPool: Should be PAL compatible after fix ✅
    deal(s_sourceFeeToken, address(s_tokenPool), feeAmount); // Re-add fees
    aggregatorBalanceBefore = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    vm.prank(s_automationAddress);
    s_tokenPool.withdrawFeeTokens(feeTokens); // Should succeed after fix
    aggregatorBalanceAfter = IERC20(s_sourceFeeToken).balanceOf(s_feeAggregator);
    assertEq(
      aggregatorBalanceAfter - aggregatorBalanceBefore,
      feeAmount,
      "TokenPool: Automation should be able to withdraw (PAL compatible after fix)"
    );

    vm.resumeGasMetering();
  }
}
