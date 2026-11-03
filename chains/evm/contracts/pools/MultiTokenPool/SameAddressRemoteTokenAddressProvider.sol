// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRemoteTokenAddressProvider} from "../../interfaces/IRemoteTokenAddressProvider.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

/// @notice Gives each token its own address as the remote token on every remote chain.
/// @dev A connection policy uses this provider only if each of its tokens has the same address on each chain of the
/// policy. The answer has the EVM encoding, so these chains must be EVM chains.
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
