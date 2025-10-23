// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

interface IExecutor {
  /// @notice Returns the minimum number of block confirmations that's allowed to be requested. The actual waiting for
  /// the block confirmations is handled by the CCVs. This value is only here to gate the value a user can request from
  /// a verifier.
  function getMinBlockConfirmations() external view returns (uint16);

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  function getFee(
    uint64 destChainSelector,
    uint16 requestedBlockDepth,
    uint32 dataLength,
    uint8 numberOfTokens,
    Client.CCV[] memory ccvs,
    bytes memory extraArgs
  ) external view returns (uint16 usdCents, uint32 gasLimit, uint32 destBytesOverhead);
}
