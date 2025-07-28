// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../libraries/Internal.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCBridgeMigrator} from "../../../../pools/USDC/USDCBridgeMigrator.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_releaseOrMint is SiloedUSDCTokenPoolSetup {

  function setUp() public virtual override {
    super.setUp();
    
    // Set up the allowed token pool proxies for testing
    address[] memory tokenPoolProxies = new address[](1);
    tokenPoolProxies[0] = s_routerAllowedOffRamp;
    bool[] memory allowed = new bool[](1);
    allowed[0] = true;
    s_usdcTokenPool.setAllowedTokenPoolProxies(tokenPoolProxies, allowed);

    // Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      rebalancer: OWNER
    });
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
      abi.encodeWithSelector(bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
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

    vm.stopPrank();
  }

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
      abi.encodeWithSelector(bytes4(keccak256("isOffRamp(uint64,address)")), SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
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

    vm.expectRevert(abi.encodeWithSelector(SiloedLockReleaseTokenPool.InsufficientLiquidity.selector, amount / 2, amount));
    s_usdcTokenPool.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

}
