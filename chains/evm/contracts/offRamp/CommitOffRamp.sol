// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../interfaces/ICCVOffRampV1.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {SignatureQuorumVerifier} from "./components/SignatureQuorumVerifier.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

contract CommitOffRamp is ICCVOffRampV1, SignatureQuorumVerifier, ITypeAndVersion {
  string public constant override typeAndVersion = "CommitOffRamp 1.7.0-dev";

  function verifyMessage(MessageV1Codec.MessageV1 calldata, bytes32 messageHash, bytes calldata ccvData) external view {
    (bytes32[] memory rs, bytes32[] memory ss) = abi.decode(ccvData, (bytes32[], bytes32[]));

    _validateSignatures(messageHash, rs, ss);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
