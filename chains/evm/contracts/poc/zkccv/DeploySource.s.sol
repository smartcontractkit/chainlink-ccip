// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../interfaces/IRouter.sol";

import {SuccinctZKVerifier} from "../../ccvs/SuccinctZKVerifier.sol";
import {VersionedVerifierResolver} from "../../ccvs/VersionedVerifierResolver.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";

import {Script, console2} from "forge-std/Script.sol";

/// @notice Deploys the ZK verifier in its source role on Sepolia against the ccv staging testnet lane.
contract DeploySource is Script {
  address internal constant RMN_REMOTE = 0xCA76dfFA5855546B98Fd5fD5823Fc771aAd8C161;
  address internal constant ROUTER = 0x784d49a71BB4C48eB7dA4cD7e6Ecb424f9b5EAB1;
  uint64 internal constant ARBITRUM_SEPOLIA_SELECTOR = 3478487238524512106;
  bytes4 internal constant VERSION_TAG = bytes4(keccak256("SuccinctZKVerifier 2.0.0"));

  function run() external {
    vm.startBroadcast();

    string[] memory storageLocations = new string[](1);
    storageLocations[0] = "local://zk-ccv-poc";
    SuccinctZKVerifier verifier = new SuccinctZKVerifier(storageLocations, RMN_REMOTE, VERSION_TAG);

    BaseVerifier.RemoteChainConfigArgs[] memory remoteConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteConfigs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(ROUTER),
      remoteChainSelector: ARBITRUM_SEPOLIA_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: 0,
      gasForVerification: 1_500_000,
      payloadSizeBytes: 30_000
    });
    verifier.applyRemoteChainConfigUpdates(remoteConfigs);

    VersionedVerifierResolver resolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.OutboundImplementationArgs[] memory outbound =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outbound[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: ARBITRUM_SEPOLIA_SELECTOR, verifier: address(verifier)
    });
    resolver.applyOutboundImplementationUpdates(outbound);

    vm.stopBroadcast();

    console2.log("source verifier", address(verifier));
    console2.log("source resolver", address(resolver));
  }
}
