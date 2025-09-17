// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../interfaces/ICCVOffRampV1.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Internal} from "../libraries/Internal.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {SignatureQuorumVerifier} from "./components/SignatureQuorumVerifier.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

contract CommitOffRamp is ICCVOffRampV1, SignatureQuorumVerifier, ITypeAndVersion {
  error ZeroAddressNotAllowed();

  error InvalidNonce(uint64 nonce);

  string public constant override typeAndVersion = "CommitOffRamp 1.7.0-dev";

  address internal immutable i_nonceManager;

  constructor(
    address nonceManager
  ) {
    if (nonceManager == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_nonceManager = nonceManager;
  }

  function verifyMessage(
    MessageV1Codec.MessageV1 calldata,
    bytes32 messageHash,
    bytes calldata ccvData,
    Internal.MessageExecutionState
  ) external {
    (bytes memory ccvArgs, bytes32[] memory rs, bytes32[] memory ss) =
      abi.decode(ccvData, (bytes, bytes32[], bytes32[]));

    _validateSignatures(keccak256(bytes.concat(messageHash, ccvArgs)), rs, ss);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
