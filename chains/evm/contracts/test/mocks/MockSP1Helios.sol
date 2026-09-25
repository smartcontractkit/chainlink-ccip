// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ISP1Helios} from "../../interfaces/succinct/ISP1Helios.sol";

contract MockSP1Helios is ISP1Helios {
  bytes32 public lightClientVkey;
  bytes32 public executionHeaderVkey;

  mapping(uint256 blockNumber => bytes32 blockHash) internal s_executionBlockHashes;

  function setVkeys(
    bytes32 newLightClientVkey,
    bytes32 newExecutionHeaderVkey
  ) external {
    lightClientVkey = newLightClientVkey;
    executionHeaderVkey = newExecutionHeaderVkey;
  }

  function setExecutionBlockHash(
    uint256 blockNumber,
    bytes32 blockHash
  ) external {
    s_executionBlockHashes[blockNumber] = blockHash;
  }

  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32) {
    return s_executionBlockHashes[blockNumber];
  }
}
