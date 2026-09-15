// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../ccvs/SuccinctZKVerifier.sol";

import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";

contract SuccinctZKVerifierHelper is SuccinctZKVerifier {
  constructor(
    DynamicConfig memory dynamicConfig,
    string[] memory storageLocations,
    address rmn,
    bytes4 versionTag
  ) SuccinctZKVerifier(dynamicConfig, storageLocations, rmn, versionTag) {}

  function validateLog(
    bytes memory receipt,
    uint256 logIndex,
    address onRamp,
    bytes32 messageId
  ) external pure {
    _validateLog(receipt, logIndex, onRamp, messageId);
  }

  function readBytes32(
    bytes memory encoded
  ) external pure returns (bytes32) {
    return _readBytes32(RLPReader.toRLPItem(encoded));
  }
}
