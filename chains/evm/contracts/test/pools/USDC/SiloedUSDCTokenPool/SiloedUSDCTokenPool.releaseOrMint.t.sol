// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SiloedUSDCTokenPool_releaseOrMint is SiloedUSDCTokenPoolSetup {
  bytes internal s_originalSender = abi.encode(makeAddr("sender"));
  bytes internal s_sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
  address internal s_recipient = makeAddr("recipient");

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
    uint256 amount = 1000e6; // 1000 USDC (6 decimals)
    uint256 localAmount = amount; // Same amount since USDC has 6 decimals
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

    // Provide liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount);
    vm.stopPrank();

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
    assertEq(s_usdcTokenPool.getAvailableTokens(SOURCE_CHAIN_SELECTOR), 0);
  }

  function test_releaseOrMint_SubtractsFromExcludedTokens() public {
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

    uint256 releaseAmount = 200e6; // Amount to release (less than excluded amount)
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

    vm.startPrank(s_routerAllowedOffRamp);

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

    Pool.ReleaseOrMintOutV1 memory result = s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    assertEq(result.destinationAmount, releaseAmount);
    assertEq(s_USDCToken.balanceOf(s_recipient), releaseAmount);

    // Verify that the excluded tokens were reduced by the release amount
    uint256 remainingExcludedTokens = excludedAmount - releaseAmount;
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(SOURCE_CHAIN_SELECTOR), remainingExcludedTokens);

    vm.stopPrank();
  }

  // Reverts

  function test_releaseOrMint_RevertWhen_InsufficientLiquidity() public {
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);

    // Provide insufficient liquidity to the siloed pool
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(SOURCE_CHAIN_SELECTOR, amount / 2); // Only half the required amount
    vm.stopPrank();

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: s_originalSender,
      receiver: s_recipient,
      sourceDenominatedAmount: amount,
      localToken: localToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: s_sourcePoolAddress,
      sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
      offchainTokenData: ""
    });

    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InsufficientLiquidity.selector, amount / 2, amount)
    );
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_ChainNotSupported() public {
    uint256 amount = 1000e6;
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

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_NotAllowedTokenPoolProxy() public {
    uint256 amount = 1000e6;
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

    vm.stopPrank();
  }
}
