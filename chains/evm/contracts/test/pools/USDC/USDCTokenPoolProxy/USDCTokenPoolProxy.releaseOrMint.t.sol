// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";

import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract USDCTokenPoolProxy_releaseOrMint is USDCTokenPoolProxySetup {
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  bytes internal s_sourcePoolAddress = abi.encode(SOURCE_CHAIN_USDC_POOL);

  function test_releaseOrMint_LockReleaseFlag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 1234;
    bytes memory lockReleaseFlag = abi.encodePacked(USDCSourcePoolDataCodec.LOCK_RELEASE_FLAG);
    bytes memory originalSender = abi.encode(s_sender);
    bytes memory offchainTokenData = "";

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    address[] memory lockReleasePools = new address[](1);
    lockReleasePools[0] = s_lockReleasePool;

    _enableERC165InterfaceChecks(s_lockReleasePool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

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

    bytes memory sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: 0, sourceDomain: 0})
    );
    bytes memory offChainTokenData = "";

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

  function test_releaseOrMint_CCTPV2_CCVFlag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 5678;
    bytes memory originalSender = abi.encode(s_sender);

    bytes memory sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2CCV(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({sourceDomain: 0, depositHash: bytes32(hex"deafbeef")})
    );
    bytes memory offChainTokenData = "";

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

  function test_releaseOrMint_CCTPV2Flag() public {
    // Arrange: Prepare test data
    uint256 testAmount = 5678;
    bytes memory originalSender = abi.encode(s_sender);

    bytes memory sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({sourceDomain: 0, depositHash: bytes32(hex"deafbeef")})
    );
    bytes memory offChainTokenData = "";

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

  function test_releaseOrMint_LegacyFormat_MessageTransmitterProxyNotSupported() public {
    uint256 testAmount = 1e6;

    bytes memory legacySourcePoolDataBytes =
      abi.encode(USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: 12345, sourceDomain: 67890}));

    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: uint32(0),
      destinationDomain: uint32(0),
      nonce: uint64(0),
      sender: bytes32(0),
      recipient: bytes32(0),
      destinationCaller: bytes32(uint256(uint160(s_legacyCctpV1Pool))),
      messageBody: ""
    });

    bytes memory message = _generateUSDCMessage(usdcMessage);
    bytes memory attestation = bytes("attestation bytes");

    bytes memory offChainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation}));

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
      offchainTokenData: offChainTokenData
    });

    // Prepare expected output
    Pool.ReleaseOrMintOutV1 memory expectedOut = Pool.ReleaseOrMintOutV1({destinationAmount: testAmount});

    // Expect the legacy pool's releaseOrMint to be called with the converted format
    // The _generateNewReleaseOrMintIn function will convert the legacy format to the new format
    // and call the legacy pool with the converted data
    vm.mockCall(
      address(s_legacyCctpV1Pool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector), abi.encode(expectedOut)
    );

    vm.expectCall(
      address(s_legacyCctpV1Pool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector, releaseOrMintIn)
    );

    // Act: Call releaseOrMint on the proxy
    Pool.ReleaseOrMintOutV1 memory actualOut = s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    // Assert: The output matches
    assertEq(actualOut.destinationAmount, expectedOut.destinationAmount);

    vm.stopPrank();
  }

  function test_releaseOrMint_LegacyFormat_MessageTransmitterProxySupported() public {
    // Set the legacy pool address to zero to simulate a scenario where there are no legacy inflight messages
    USDCTokenPoolProxy.PoolAddresses memory updatedPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: address(0), // Set to zero to indicate no legacy pool
      cctpV1Pool: s_cctpV1Pool,
      cctpV2Pool: s_cctpV2Pool
    });

    _enableERC165InterfaceChecks(s_cctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_cctpV2Pool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(updatedPools);

    // Arrange: Prepare test data for legacy format (64 bytes)
    uint256 testAmount = 1e6;

    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory legacySourcePoolData =
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: 12345, sourceDomain: 67890});

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

    bytes memory invalidSourcePoolData = abi.encodePacked(bytes4(uint32(2)), uint32(0), bytes32(hex"deafbeef"));

    bytes memory offChainTokenData = "";

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
      sourcePoolData: invalidSourcePoolData,
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: offChainTokenData
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.InvalidMessageVersion.selector, bytes4(uint32(2))));
    s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);

    vm.stopPrank();
  }

  function test_releaseOrMint_RevertWhen_CallerIsNotARampOnRouter() public {
    address unauthorized = makeAddr("unauthorized");

    vm.startPrank(unauthorized);

    // Prepare input with CCTP_V1_FLAG in sourcePoolData (version 0 in offchainTokenData)
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: abi.encode(s_sender),
      receiver: s_receiver,
      sourceDenominatedAmount: 4321,
      localToken: address(s_USDCToken),
      sourcePoolData: "",
      sourcePoolAddress: s_sourcePoolAddress,
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CallerIsNotARampOnRouter.selector, unauthorized));
    s_usdcTokenPoolProxy.releaseOrMint(releaseOrMintIn);
  }

  function _enableERC165InterfaceChecks(address pool, bytes4 interfaceId) internal {
    vm.mockCall(
      address(pool), abi.encodeWithSelector(IERC165.supportsInterface.selector, interfaceId), abi.encode(true)
    );

    vm.mockCall(
      address(pool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId),
      abi.encode(true)
    );
  }
}
