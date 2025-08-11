// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_releaseOrMint is USDCTokenPoolProxySetup {
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  bytes internal s_sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

  struct LegacySourcePoolData {
    uint64 nonce;
    uint32 sourceDomain;
  }

  function test_releaseOrMint_LockReleaseFlag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 1234;
    bytes memory lockReleaseFlag = abi.encodePacked(LOCK_RELEASE_FLAG);
    bytes memory originalSender = abi.encode(s_sender);
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
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      sourcePoolData: lockReleaseFlag,
      sourcePoolAddress: s_sourcePoolAddress,
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

  function test_releaseOrMint_CCTPV1Flag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 4321;
    bytes memory originalSender = abi.encode(s_sender);
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
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

    // Prepare input with CCTP_V1_FLAG in sourcePoolData (version 0 in offchainTokenData)
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: originalSender,
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      sourcePoolData: sourcePoolData,
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: offChainTokenData
    });

    Pool.ReleaseOrMintOutV1 memory expectedOut = Pool.ReleaseOrMintOutV1({destinationAmount: testAmount});

    // Expect the cctpV1Pool's releaseOrMint to be called and return expectedOut
    vm.mockCall(
      address(s_cctpV1Pool),
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
    bytes memory originalSender = abi.encode(s_sender);
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
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      // sourcePoolData should be a USDC Message where the version number is 1
      sourcePoolData: sourcePoolData,
      sourcePoolAddress: s_sourcePoolAddress,
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

  function test_releaseOrMint_LegacyFormat() public {
    // Arrange: Prepare test data for legacy format (64 bytes)
    uint256 testAmount = 1e6;

    LegacySourcePoolData memory legacySourcePoolData = LegacySourcePoolData({nonce: 12345, sourceDomain: 67890});

    bytes memory legacySourcePoolDataBytes = abi.encode(legacySourcePoolData);

    // Mock the CCTP V1 pool's i_localDomainIdentifier to return a test domain
    uint32 testLocalDomain = 12345;
    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(bytes4(keccak256("i_localDomainIdentifier()"))),
      abi.encode(testLocalDomain)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    // Prepare input with legacy 64-byte sourcePoolData
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: abi.encode(s_sender),
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      sourcePoolData: legacySourcePoolDataBytes, // 64 bytes: uint64 + uint32
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: ""
    });

    // Prepare expected output
    Pool.ReleaseOrMintOutV1 memory expectedOut = Pool.ReleaseOrMintOutV1({destinationAmount: testAmount});

    // Expect the legacy pool's releaseOrMint to be called with the converted format
    // The _generateNewReleaseOrMintIn function will convert the legacy format to the new format
    // and call the legacy pool with the converted data
    vm.mockCall(
      address(s_cctpV1Pool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector), abi.encode(expectedOut)
    );

    // Act: Call releaseOrMint on the proxy
    Pool.ReleaseOrMintOutV1 memory actualOut = s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    // Assert: The output matches
    assertEq(actualOut.destinationAmount, expectedOut.destinationAmount);

    vm.stopPrank();
  }

  // Reverts

  function test_releaseOrMint_InvalidVersion() public {
    // Arrange: Prepare test data
    uint256 testAmount = 1234;
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
      originalSender: abi.encode(s_sender),
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      sourcePoolData: sourcePoolData,
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: offChainTokenData
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.InvalidMessageVersion.selector, 2));
    s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_InvalidMessageLength() public {
    // Arrange: Prepare test data with offchainTokenData less than 4 bytes
    uint256 testAmount = 1234;

    // Create offchainTokenData with only 3 bytes (less than the required 4 bytes for version)
    bytes memory shortOffchainTokenData = new bytes(3);
    shortOffchainTokenData[0] = 0x01;
    shortOffchainTokenData[1] = 0x02;
    shortOffchainTokenData[2] = 0x03;

    // Mock the router's isOffRamp function to return true
    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(Router.isOffRamp.selector, SOURCE_CHAIN_SELECTOR, s_routerAllowedOffRamp),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOffRamp);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: abi.encode(s_sender),
      receiver: s_receiver,
      sourceDenominatedAmount: testAmount,
      localToken: address(s_USDCToken),
      sourcePoolData: "", // Not lock release flag, not legacy format
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: shortOffchainTokenData
    });

    // Expect revert with InvalidMessageLength error
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.InvalidMessageLength.selector, 3));
    s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }
}
