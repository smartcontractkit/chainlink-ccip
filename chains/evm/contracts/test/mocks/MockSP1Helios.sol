// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ISP1Helios} from "../../interfaces/succinct/ISP1Helios.sol";

contract MockSP1Helios is ISP1Helios {
  mapping(uint256 blockNumber => bytes32 blockHash) internal s_executionBlockHashes;
  mapping(uint256 blockNumber => bytes32 receiptsRoot) internal s_executionReceiptsRoots;

  function setAnchor(
    uint256 blockNumber,
    bytes32 blockHash,
    bytes32 receiptsRoot
  ) external {
    s_executionBlockHashes[blockNumber] = blockHash;
    s_executionReceiptsRoots[blockNumber] = receiptsRoot;
  }

  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32) {
    return s_executionBlockHashes[blockNumber];
  }

  function executionReceiptsRoots(
    uint256 blockNumber
  ) external view returns (bytes32) {
    return s_executionReceiptsRoots[blockNumber];
  }
}
