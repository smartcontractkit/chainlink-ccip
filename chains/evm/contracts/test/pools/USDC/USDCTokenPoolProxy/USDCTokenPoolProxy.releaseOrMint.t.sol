// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_releaseOrMint is USDCTokenPoolProxySetup {
  function test_releaseOrMint_LockReleaseFlag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 1234;
    address testSender = makeAddr("sender");
    address testReceiver = makeAddr("receiver");
    address localToken = address(s_USDCToken);
    bytes memory lockReleaseFlag = abi.encodePacked(LOCK_RELEASE_FLAG);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(testSender);
    bytes memory offchainTokenData = "";

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(Router.isOffRamp.selector, SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    // Prepare input with LOCK_RELEASE_FLAG in sourcePoolData
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: originalSender,
      receiver: testReceiver,
      sourceDenominatedAmount: testAmount,
      localToken: localToken,
      sourcePoolData: lockReleaseFlag,
      sourcePoolAddress: sourcePoolAddress,
      offchainTokenData: offchainTokenData
    });

    // Prepare expected output
    Pool.ReleaseOrMintOutV1 memory expectedOut = Pool.ReleaseOrMintOutV1({destinationAmount: testAmount});

    // Expect the lockReleasePool's releaseOrMint to be called and return expectedOut
    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector, releaseOrMintIn),
      abi.encode(expectedOut)
    );

    // Act: Call releaseOrMint on the proxy
    Pool.ReleaseOrMintOutV1 memory actualOut = s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    // Assert: The output matches
    assertEq(actualOut.destinationAmount, expectedOut.destinationAmount);

    vm.stopPrank();
  }

  function test_releaseOrMint_CCTPV2Flag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 5678;
    address testSender = makeAddr("sender");
    address testReceiver = makeAddr("receiver");
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(testSender);
    USDCMessage memory usdcMessage = USDCMessage({
      version: 1,
      sourceDomain: uint32(0),
      destinationDomain: uint32(0),
      nonce: uint64(0),
      sender: bytes32(0),
      recipient: bytes32(0),
      destinationCaller: bytes32(0),
      messageBody: ""
    });
    bytes memory sourcePoolData = "";
    bytes memory offChainTokenData = _generateUSDCMessage(usdcMessage);

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(Router.isOffRamp.selector, SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    // Prepare input with CCTP_V2_FLAG in sourcePoolData
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: originalSender,
      receiver: testReceiver,
      sourceDenominatedAmount: testAmount,
      localToken: localToken,
      // sourcePoolData should be a USDC Message where the version number is 1
      sourcePoolData: sourcePoolData,
      sourcePoolAddress: sourcePoolAddress,
      offchainTokenData: offChainTokenData
    });

    Pool.ReleaseOrMintOutV1 memory expectedOut = Pool.ReleaseOrMintOutV1({destinationAmount: testAmount});

    // Expect the cctpV2Pool's releaseOrMint to be called and return expectedOut
    vm.mockCall(
      address(s_cctpV2Pool),
      abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector, releaseOrMintIn),
      abi.encode(expectedOut)
    );

    Pool.ReleaseOrMintOutV1 memory actualOut = s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    assertEq(actualOut.destinationAmount, expectedOut.destinationAmount);

    vm.stopPrank();
  }

  function test_releaseOrMint_InvalidVersion() public {
    // Arrange: Prepare test data
    uint256 testAmount = 1234;
    address testSender = makeAddr("sender");
    address testReceiver = makeAddr("receiver");
    address localToken = address(s_USDCToken);
    bytes memory sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);
    bytes memory originalSender = abi.encode(testSender);
    bytes memory emptyMessageBody = new bytes(0);

    // Create a USDC message with version = 2 (not 0 or 1)
    USDCMessage memory usdcMessage = USDCMessage({
      version: uint32(2),
      sourceDomain: uint32(0),
      destinationDomain: uint32(0),
      nonce: uint64(0),
      sender: bytes32(0),
      recipient: bytes32(0),
      destinationCaller: bytes32(0),
      messageBody: emptyMessageBody
    });
    bytes memory sourcePoolData = "";
    bytes memory offChainTokenData = _generateUSDCMessage(usdcMessage);

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(Router.isOffRamp.selector, SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: originalSender,
      receiver: testReceiver,
      sourceDenominatedAmount: testAmount,
      localToken: localToken,
      sourcePoolData: sourcePoolData,
      sourcePoolAddress: sourcePoolAddress,
      offchainTokenData: offChainTokenData
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.InvalidMessageVersion.selector, 2));
    s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }
}
