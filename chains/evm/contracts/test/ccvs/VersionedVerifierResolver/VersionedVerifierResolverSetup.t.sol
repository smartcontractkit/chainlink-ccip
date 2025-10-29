// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract VersionedVerifierResolverSetup is BaseTest {
  bytes4 internal constant INITIAL_VERSION_1 = bytes4(0x11111111);
  bytes4 internal constant INITIAL_VERSION_2 = bytes4(0x22222222);
  bytes4 internal constant UNKNOWN_VERSION = bytes4(0x98989898);
  uint64 internal constant INITIAL_DEST_CHAIN_SELECTOR_1 = 1;
  uint64 internal constant INITIAL_DEST_CHAIN_SELECTOR_2 = 2;
  uint64 internal constant UNKNOWN_DEST_CHAIN_SELECTOR = 999;

  VersionedVerifierResolver internal s_versionedVerifierResolver;
  address internal s_initialVerifier1;
  address internal s_initialVerifier2;

  function setUp() public virtual override {
    super.setUp();

    s_initialVerifier1 = makeAddr("InitialVerifier1");
    s_initialVerifier2 = makeAddr("InitialVerifier2");
    s_versionedVerifierResolver = new VersionedVerifierResolver();

    VersionedVerifierResolver.InboundImplementationArgs[] memory inboundImplementations =
      new VersionedVerifierResolver.InboundImplementationArgs[](2);
    inboundImplementations[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION_1, verifier: s_initialVerifier1});
    inboundImplementations[1] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION_2, verifier: s_initialVerifier2});
    s_versionedVerifierResolver.applyInboundImplementationUpdates(inboundImplementations);

    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImplementations =
      new VersionedVerifierResolver.OutboundImplementationArgs[](2);
    outboundImplementations[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR_1,
      verifier: s_initialVerifier1
    });
    outboundImplementations[1] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR_2,
      verifier: s_initialVerifier2
    });
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(outboundImplementations);
  }
}
