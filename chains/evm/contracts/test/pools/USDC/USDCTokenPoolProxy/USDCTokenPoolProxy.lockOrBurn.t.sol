// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_lockOrBurn is USDCTokenPoolProxySetup {
  function setUp() public virtual override {
    super.setUp();

    // Configure lock or burn mechanisms for different chains
    uint64[] memory chainSelectors = new uint64[](3);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    chainSelectors[1] = DEST_CHAIN_SELECTOR;
    chainSelectors[2] = 12345; // Another test chain

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](3);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    mechanisms[2] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_lockOrBurn_CCTPV1() public {
    // Arrange: Define test constants
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 amount = 100;
    address localToken = address(s_USDCToken);
    bytes memory destPoolData = abi.encode(1, 2, 3);
    bytes memory destTokenAddress = abi.encode(localToken);

    // Set the DEST_CHAIN_SELECTOR to use CCTP V1 using the update function
    uint64[] memory selectors = new uint64[](1);
    selectors[0] = DEST_CHAIN_SELECTOR;
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechs = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechs[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(selectors, mechs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    // Mock the CCTP V1 pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: destPoolData});

    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);

    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_CCTPV2() public {
    // Arrange: Define test constants
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 amount = 200;
    address localToken = address(s_USDCToken);
    bytes memory destPoolData = abi.encode(4, 5, 6);
    bytes memory destTokenAddress = abi.encode(localToken);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    // Mock the CCTP V2 pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: destPoolData});

    vm.mockCall(
      address(s_cctpV2Pool),
      abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.startPrank(s_routerAllowedOnRamp);
    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_LockRelease() public {
    // Arrange: Define test constants
    uint64 testChainSelector = 12345;
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 amount = 300;
    address localToken = address(s_USDCToken);
    bytes memory destPoolData = abi.encode(7, 8, 9);
    bytes memory destTokenAddress = abi.encode(localToken);
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    // Mock the previous pool's releaseOrMint function to return the input amount
    // Add support for remoteChainSelector 12345 on the token pool proxy
    {
      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: testChainSelector,
        remotePoolAddresses: remotePoolAddresses,
        remoteTokenAddress: abi.encode(localToken),
        outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: _getInboundRateLimiterConfig()
      });
      s_usdcTokenPoolProxy.applyChainUpdates(new uint64[](0), chainUpdates);
    }

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(testChainSelector)),
      abi.encode(s_routerAllowedOnRamp)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: testChainSelector,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    // Mock the lock release pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: destPoolData});

    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_RevertWhen_InvalidLockOrBurnMechanism() public {
    // Arrange: Define test constants
    uint64 testChainSelector = 99999;
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 amount = 100;
    address localToken = address(s_USDCToken);
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    // Call applyChainUpdates to add support for remoteChainSelector 99999
    {
      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: testChainSelector,
        remotePoolAddresses: remotePoolAddresses,
        remoteTokenAddress: abi.encode(localToken),
        outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: _getInboundRateLimiterConfig()
      });
      s_usdcTokenPoolProxy.applyChainUpdates(new uint64[](0), chainUpdates);
    }

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(testChainSelector)),
      abi.encode(s_routerAllowedOnRamp)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: testChainSelector, // Chain with no configured mechanism
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism(0)
      )
    );

    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_InvalidReceiver() public {
    // Arrange: Define test constants
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 amount = 100;
    address localToken = address(s_USDCToken);
    bytes memory invalidReceiver = abi.encode(receiver, "extra"); // Invalid receiver format

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: invalidReceiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    vm.expectRevert();
    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_InvalidAmount() public {
    // Arrange: Define test constants
    address receiver = makeAddr("receiver");
    address sender = makeAddr("sender");
    uint256 invalidAmount = 0; // Invalid amount
    address localToken = address(s_USDCToken);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: invalidAmount,
      localToken: localToken
    });

    vm.expectRevert();
    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }
}
