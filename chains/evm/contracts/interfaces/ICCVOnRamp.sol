// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface ICCVOnRamp {
  function getFee(
    uint64 destChainSelector,
    bytes32 version,
    address caller,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256);

  /// @notice Message sending
  // TODO versioning?
  function forwardToVerifier(
    uint64 remoteChainSelector,
    bytes32 version,
    address caller,
    bytes memory rawMessage,
    uint256 verifierIndex
  ) external returns (bytes memory);
}
