// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutionRootOracle} from "../../interfaces/IExecutionRootOracle.sol";
import {IRouter} from "../../interfaces/IRouter.sol";

import {SuccinctZKVerifier} from "../../ccvs/SuccinctZKVerifier.sol";
import {VersionedVerifierResolver} from "../../ccvs/VersionedVerifierResolver.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {MockReceiverV2} from "../../test/mocks/MockReceiverV2.sol";

import {Script, console2} from "forge-std/Script.sol";

/// @notice Deploys the ZK verifier in its destination role on Arbitrum Sepolia, plus a receiver that requires it.
/// @dev Reads the SP1Helios address from the HELIOS environment variable.
contract DeployDest is Script {
  address internal constant RMN_REMOTE = 0x6f23BA5C1d74083e5bDe7d709Bb3B62a899E4baB;
  address internal constant ROUTER = 0x8F95FA37c55eF7beFdf05f6abDeC551773E17Fb4;
  address internal constant SEPOLIA_ON_RAMP = 0xA94E45744553F4B2bea9DfB8979a02962B980732;
  uint64 internal constant SEPOLIA_SELECTOR = 16015286601757825753;
  bytes4 internal constant VERSION_TAG = bytes4(keccak256("SuccinctZKVerifier 2.0.0"));

  function run() external {
    address helios = vm.envAddress("HELIOS");

    vm.startBroadcast();

    string[] memory storageLocations = new string[](1);
    storageLocations[0] = "local://zk-ccv-poc";
    SuccinctZKVerifier verifier = new SuccinctZKVerifier(storageLocations, RMN_REMOTE, VERSION_TAG);

    BaseVerifier.RemoteChainConfigArgs[] memory remoteConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteConfigs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(ROUTER),
      remoteChainSelector: SEPOLIA_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: 0,
      gasForVerification: 1_500_000,
      payloadSizeBytes: 30_000
    });
    verifier.applyRemoteChainConfigUpdates(remoteConfigs);

    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceConfigs = new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      sourceChainSelector: SEPOLIA_SELECTOR,
      rootOracle: IExecutionRootOracle(helios),
      onRamp: SEPOLIA_ON_RAMP,
      maxHeaderChainLength: 1024
    });
    verifier.applySourceChainConfigUpdates(sourceConfigs);

    VersionedVerifierResolver resolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.InboundImplementationArgs[] memory inbound =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    inbound[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: VERSION_TAG, verifier: address(verifier)});
    resolver.applyInboundImplementationUpdates(inbound);

    address[] memory required = new address[](1);
    required[0] = address(resolver);
    MockReceiverV2 receiver = new MockReceiverV2(required, new address[](0), 0);

    vm.stopBroadcast();

    console2.log("dest verifier", address(verifier));
    console2.log("dest resolver", address(resolver));
    console2.log("dest receiver", address(receiver));
  }
}
