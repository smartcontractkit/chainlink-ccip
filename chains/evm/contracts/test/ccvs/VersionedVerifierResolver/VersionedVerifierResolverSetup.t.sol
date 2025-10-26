// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract VersionedVerifierResolverSetup is BaseTest {
  bytes4 internal constant INITIAL_VERSION = bytes4(0x12345678);
  bytes4 internal constant UNKNOWN_VERSION = bytes4(0x87654321);
  uint64 internal constant INITIAL_DEST_CHAIN_SELECTOR = 1;
  uint64 internal constant UNKNOWN_DEST_CHAIN_SELECTOR = 2;

  VersionedVerifierResolver internal s_versionedVerifierResolver;
  address internal s_initialVerifier;

  function setUp() public virtual override {
    super.setUp();

    s_initialVerifier = makeAddr("MockVerifier");

    s_versionedVerifierResolver = new VersionedVerifierResolver();

    VersionedVerifierResolver.InboundImplementationArgs[] memory inboundImplementations =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    inboundImplementations[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION, verifier: s_initialVerifier});
    s_versionedVerifierResolver.applyInboundImplementationUpdates(new bytes4[](0), inboundImplementations);

    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImplementations =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outboundImplementations[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR,
      verifier: s_initialVerifier
    });
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(new uint64[](0), outboundImplementations);
  }
}
