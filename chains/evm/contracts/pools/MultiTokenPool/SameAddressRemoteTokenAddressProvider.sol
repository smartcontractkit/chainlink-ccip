// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRemoteTokenAddressProvider} from "../../interfaces/IRemoteTokenAddressProvider.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

/// @notice Returns the local token address as the remote token address.
/// @dev Only use this provider for EVM chains where each token has the same address as on the local chain.
contract SameAddressRemoteTokenAddressProvider is IRemoteTokenAddressProvider, ITypeAndVersion {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "SameAddressRemoteTokenAddressProvider 2.1.0-dev";
  }

  /// @inheritdoc IRemoteTokenAddressProvider
  function getRemoteTokenAddress(
    address localToken,
    uint64 // remoteChainSelector
  ) public pure virtual returns (bytes memory remoteToken) {
    return abi.encode(localToken);
  }
}
