// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface ICCVOnRamp {
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256);

  /// @notice Message sending
  // TODO versioning?
  function forwardToVerifier(bytes memory rawMessage, uint256 verifierIndex) external returns (bytes memory);
}
