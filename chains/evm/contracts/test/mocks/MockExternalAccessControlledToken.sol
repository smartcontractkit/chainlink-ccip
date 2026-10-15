// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IExternalAccessControlRegistry} from "../../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

/// @notice Token mock that returns an external AccessControl contract.
contract MockExternalAccessControlledToken is IExternalAccessControlRegistry {
  address private immutable i_accessControlRegistry;

  /// @param accessControlRegistry The AccessControl contract that holds the roles for the token.
  constructor(
    address accessControlRegistry
  ) {
    i_accessControlRegistry = accessControlRegistry;
  }

  /// @inheritdoc IExternalAccessControlRegistry
  // solhint-disable-next-line func-name-mixedcase
  function ACCESS_CONTROLLED_REGISTRY() external view returns (address) {
    return i_accessControlRegistry;
  }
}
