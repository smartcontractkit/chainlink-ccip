// SPDX-License-Identifier: BUSL-1.1
// solhint-disable one-contract-per-file
pragma solidity ^0.8.24;

import {IHasUID} from "../../interfaces/IHasUID.sol";
import {IResolveTokenAddressFromUID} from "../../interfaces/IResolveTokenAddressFromUID.sol";

/// @notice A token with a fixed UID. It can also be deployed outside a factory.
contract MockTokenWithUID is IHasUID {
  bytes32 private immutable i_uid;

  constructor(
    bytes32 tokenUid
  ) {
    i_uid = tokenUid;
  }

  /// @inheritdoc IHasUID
  function uid() external view returns (bytes32) {
    return i_uid;
  }
}

/// @notice A token factory that deploys one MockTokenWithUID for each UID.
contract MockTokenWithUIDFactory is IResolveTokenAddressFromUID {
  error TokenAlreadyDeployed(bytes32 uid);

  mapping(bytes32 uid => address token) private s_tokens;

  /// @notice Deploys a token for a UID.
  /// @param uid The UID of the new token.
  /// @return token The new token.
  function deploy(
    bytes32 uid
  ) external returns (address token) {
    if (s_tokens[uid] != address(0)) revert TokenAlreadyDeployed(uid);
    token = address(new MockTokenWithUID(uid));
    s_tokens[uid] = token;
    return token;
  }

  /// @inheritdoc IResolveTokenAddressFromUID
  function tokenAddress(
    bytes32 uid
  ) external view returns (address) {
    return s_tokens[uid];
  }
}
