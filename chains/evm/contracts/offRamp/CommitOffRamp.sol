// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../interfaces/ICCVOffRampV1.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {SignatureQuorumVerifier} from "./components/SignatureQuorumVerifier.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

contract CommitOffRamp is ICCVOffRampV1, SignatureQuorumVerifier, ITypeAndVersion {
  error InvalidCCVData();

  string public constant override typeAndVersion = "CommitOffRamp 1.7.0-dev";

  uint256 internal constant SIGNATURE_LENGTH_BYTES = 2;

  function verifyMessage(MessageV1Codec.MessageV1 calldata, bytes32 messageHash, bytes calldata ccvData) external view {
    if (ccvData.length < SIGNATURE_LENGTH_BYTES) {
      revert InvalidCCVData();
    }

    uint256 signatureLength = uint16(bytes2(ccvData[:SIGNATURE_LENGTH_BYTES]));
    if (ccvData.length < SIGNATURE_LENGTH_BYTES + signatureLength) {
      revert InvalidCCVData();
    }

    // Even though the current version of this contract only expects signatures to be included in the ccvData, bounding
    // it to the given length allows potential forward compatibility with future formats that supply more data.
    _validateSignatures(messageHash, ccvData[SIGNATURE_LENGTH_BYTES:SIGNATURE_LENGTH_BYTES + signatureLength]);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
