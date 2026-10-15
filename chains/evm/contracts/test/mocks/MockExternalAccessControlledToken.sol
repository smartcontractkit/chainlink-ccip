// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IExternalAccessControlRegistry} from "../../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

contract MockExternalAccessControlledToken is IExternalAccessControlRegistry {
  address private immutable i_accessControlledRegistry;

  constructor(
    address accessControlledRegistry
  ) {
    i_accessControlledRegistry = accessControlledRegistry;
  }

  /// @inheritdoc IExternalAccessControlRegistry
  // solhint-disable-next-line func-name-mixedcase
  function ACCESS_CONTROLLED_REGISTRY() external view returns (address) {
    return i_accessControlledRegistry;
  }
}
