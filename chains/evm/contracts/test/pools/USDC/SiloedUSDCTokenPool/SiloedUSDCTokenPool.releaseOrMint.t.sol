// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/BurnMintWithLockReleaseFlagTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SiloedUSDCTokenPool_releaseOrMint is SiloedUSDCTokenPoolSetup {
  bytes internal s_originalSender = abi.encode(makeAddr("sender"));
  bytes internal s_sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
  address internal s_recipient = makeAddr("recipient");
  uint256 internal constant DEFAULT_LIQUIDITY = 1000e6;

  function setUp() public virtual override {
    super.setUp();

    // Provide default liquidity to the source lockbox
    deal(address(s_USDCToken), address(s_sourceLockBox), DEFAULT_LIQUIDITY);

    // Mock the router's isOffRamp function to return true for the allowed off ramp
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(
        bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp
      ),
      abi.encode(true)
    );
  }

  function test_releaseOrMint_Success() public {
    uint256 amount = DEFAULT_LIQUIDITY;
    uint256 localAmount = amount;
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: localToken,
      sender: s_routerAllowedOffRamp,
      recipient: s_recipient,
      amount: localAmount
    });

    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    assertEq(result.destinationAmount, localAmount);
    assertEq(s_USDCToken.balanceOf(s_recipient), localAmount);
    assertEq(s_USDCToken.balanceOf(address(s_sourceLockBox)), 0);
  }

  function test_releaseOrMintV2_Success() public {
    uint256 amount = DEFAULT_LIQUIDITY;
    uint256 localAmount = amount;
    address localToken = address(s_USDCToken);

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn, 0);

    assertEq(result.destinationAmount, localAmount);
    assertEq(s_USDCToken.balanceOf(s_recipient), localAmount);
    assertEq(s_USDCToken.balanceOf(address(s_sourceLockBox)), 0);
  }

  function test_releaseOrMint_SubtractsFromExcludedTokens() public {
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);

    // Set up the circle migrator address
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);
    s_usdcTokenPool.setLockedUSDCToBurn(SOURCE_CHAIN_SELECTOR, DEFAULT_LIQUIDITY);

    // Exclude some tokens from burn (liquidity already provided in setUp)
    uint256 excludedAmount = 500e6;
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, excludedAmount);
    vm.stopPrank();

    // Calling releaseOrMint before the burn event should subtract from total token balance and excluded tokens.
    uint256 releaseAmount = 200e6; // Amount to release (less than excluded amount)
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: releaseAmount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.startPrank(s_routerAllowedOffRamp);
    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    uint256 availableTokens = s_USDCToken.balanceOf(address(s_sourceLockBox));
    assertEq(availableTokens, DEFAULT_LIQUIDITY - releaseAmount);
    uint256 newExcludedAmount = s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR);
    assertEq(newExcludedAmount, excludedAmount - releaseAmount);

    vm.startPrank(OWNER);
    s_usdcTokenPool.setLockedUSDCToBurn(SOURCE_CHAIN_SELECTOR, availableTokens);

    // Execute the migration to mark the chain as migrated
    vm.startPrank(circleMigrator);
    s_usdcTokenPool.burnLockedUSDC();

    // Verify the chain is now migrated and tokens are still excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), newExcludedAmount);

    vm.startPrank(s_routerAllowedOffRamp);
    result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    assertEq(result.destinationAmount, releaseAmount);
    assertEq(s_USDCToken.balanceOf(s_recipient), releaseAmount * 2); // There were two releases

    // Verify that the excluded tokens were reduced by the release amount
    uint256 remainingExcludedTokens = newExcludedAmount - releaseAmount;
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), remainingExcludedTokens);
  }

  function test_releaseOrMint_RevertWhen_InsufficientLiquidity_ProposedChainHasNoExcludedLiquidity() public {
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);

    uint256 releaseAmount = 100e6;
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: releaseAmount,
      localToken: address(s_USDCToken),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.startPrank(s_routerAllowedOffRamp);
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.InsufficientLiquidity.selector, 0, releaseAmount));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);
  }

  function test_releaseOrMint_RevertWhen_InsufficientLiquidity_MigratedChainHasNoExcludedLiquidity() public {
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);

    uint256 excludedAmount = 200e6;
    s_usdcTokenPool.setLockedUSDCToBurn(SOURCE_CHAIN_SELECTOR, DEFAULT_LIQUIDITY);
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, excludedAmount);

    vm.startPrank(circleMigrator);
    s_usdcTokenPool.burnLockedUSDC();

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: excludedAmount,
      localToken: address(s_USDCToken),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    // Consume the full excluded reserve with a valid post-migration in-flight message.
    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), 0);

    // Simulate unexpected liquidity appearing in the lockbox after migration.
    uint256 unexpectedLiquidity = 100e6;
    deal(address(s_USDCToken), address(s_sourceLockBox), unexpectedLiquidity);

    releaseOrMintIn.sourceDenominatedAmount = unexpectedLiquidity;

    vm.startPrank(s_routerAllowedOffRamp);
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.InsufficientLiquidity.selector, 0, unexpectedLiquidity));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);
  }

  // Reverts

  function test_releaseOrMint_RevertWhen_InsufficientLiquidity_InsufficientExcludedTokens() public {
    // Propose a CCTP migration to enable excluded tokens tracking.
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
    s_usdcTokenPool.setLockedUSDCToBurn(SOURCE_CHAIN_SELECTOR, DEFAULT_LIQUIDITY);

    // Exclude only a small amount of tokens from burn.
    uint256 excludedAmount = 100e6;
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, excludedAmount);
    vm.stopPrank();

    // Try to release more than the excluded amount.
    uint256 releaseAmount = 200e6;

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: releaseAmount,
      localToken: address(s_USDCToken),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    // Should revert because excluded tokens (100e6) < release amount (200e6).
    vm.expectRevert(
      abi.encodeWithSelector(SiloedUSDCTokenPool.InsufficientLiquidity.selector, excludedAmount, releaseAmount)
    );
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);
  }

  function test_releaseOrMint_RevertWhen_ChainNotSupported() public {
    uint256 amount = DEFAULT_LIQUIDITY;
    address localToken = address(s_USDCToken);
    uint64 unsupportedChain = 999999999; // Chain that's not configured

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: unsupportedChain,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChain));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);
  }

  function test_releaseOrMint_RevertWhen_NotAllowedTokenPoolProxy() public {
    uint256 amount = DEFAULT_LIQUIDITY;
    address unauthorizedProxy = makeAddr("unauthorizedProxy");

    // Mock the router's isOffRamp function to return true for the unauthorized proxy
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, unauthorizedProxy),
      abi.encode(true)
    );

    vm.startPrank(unauthorizedProxy);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: amount,
      localToken: address(s_USDCToken),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedProxy));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);
  }
}
