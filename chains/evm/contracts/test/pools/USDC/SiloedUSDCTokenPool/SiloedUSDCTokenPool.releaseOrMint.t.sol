// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_releaseOrMint is SiloedUSDCTokenPoolSetup {
  function setUp() public virtual override {
    super.setUp();

    // Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] =
      SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: SOURCE_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);
    vm.stopPrank();
  }

  function test_releaseOrMint_Success() public {
    // Arrange: Define test constants
    address recipient = makeAddr("recipient");
    uint256 amount = 1000e6; // 1000 USDC (6 decimals)
    uint256 localAmount = amount; // Same amount since USDC has 6 decimals
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));

    // Provide liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount);
    vm.stopPrank();

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(
        bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp
      ),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    // Act: Call releaseOrMint
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
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
      recipient: recipient,
      amount: localAmount
    });

    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    // Assert: Verify the result
    assertEq(result.destinationAmount, localAmount);
    assertEq(s_USDCToken.balanceOf(recipient), localAmount);
    assertEq(s_usdcTokenPool.getAvailableTokens(SOURCE_CHAIN_SELECTOR), 0);
  }

  function test_releaseOrMint_SubtractsFromExcludedTokens() public {
    // Arrange: Create a migration proposal and exclude tokens
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);

    // Set up the circle migrator address
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);

    // Provide liquidity and exclude some tokens from burn
    uint256 excludedAmount = 500e6;
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, 1000e6);
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, excludedAmount);
    vm.stopPrank();

    // Execute the migration to mark the chain as migrated
    vm.startPrank(circleMigrator);
    s_usdcTokenPool.burnLockedUSDC();
    vm.stopPrank();

    // Verify the chain is now migrated and tokens are excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), excludedAmount);

    // Arrange: Define test constants for releaseOrMint
    address recipient = makeAddr("recipient");
    uint256 releaseAmount = 200e6; // Amount to release (less than excluded amount)
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(
        bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp
      ),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    // Act: Call releaseOrMint - this should subtract from excluded tokens since the chain is migrated
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
      sourceDenominatedAmount: releaseAmount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    // Assert: Verify the result and that excluded tokens were reduced
    assertEq(result.destinationAmount, releaseAmount);
    assertEq(s_USDCToken.balanceOf(recipient), releaseAmount);

    // Verify that the excluded tokens were reduced by the release amount
    uint256 remainingExcludedTokens = excludedAmount - releaseAmount;
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), remainingExcludedTokens);

    vm.stopPrank();
  }

  // Reverts

  function test_releaseOrMint_RevertWhen_InsufficientLiquidity() public {
    // Arrange: Define test constants
    address recipient = makeAddr("recipient");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));

    // Provide insufficient liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount / 2); // Only half the required amount
    vm.stopPrank();

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(
        bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp
      ),
      abi.encode(true)
    );

    // Act & Assert: Call releaseOrMint with insufficient liquidity
    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InsufficientLiquidity.selector, amount / 2, amount)
    );
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_NotAllowedOffRamp() public {
    // Arrange: Define test constants
    address recipient = makeAddr("recipient");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));
    address unauthorizedCaller = makeAddr("unauthorized");

    // Provide liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount);
    vm.stopPrank();

    // Mock the router's isOffRamp function to return false for unauthorized caller
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, unauthorizedCaller),
      abi.encode(false)
    );

    // Act & Assert: Call releaseOrMint with unauthorized caller
    vm.startPrank(unauthorizedCaller);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedCaller));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_ChainNotSupported() public {
    // Arrange: Define test constants
    address recipient = makeAddr("recipient");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));
    uint64 unsupportedChain = 999999999; // Chain that's not configured

    // Act & Assert: Call releaseOrMint with unsupported chain
    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: unsupportedChain,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChain));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_NotAllowedTokenPoolProxy() public {
    // Arrange: Define test constants
    address recipient = makeAddr("recipient");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(makeAddr("sender"));
    address unauthorizedProxy = makeAddr("unauthorizedProxy");

    // Provide liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount);
    vm.stopPrank();

    // Mock the router's isOffRamp function to return true for the unauthorized proxy
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, unauthorizedProxy),
      abi.encode(true)
    );

    // Act & Assert: Call releaseOrMint with unauthorized proxy
    vm.startPrank(unauthorizedProxy);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedProxy));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }
}
