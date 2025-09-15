// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRampV1} from "../interfaces/ICCVOffRampV1.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
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
    uint64, // sourceChainSelector
    address, // originalCaller
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageHash,
    bytes calldata ccvData,
    Internal.MessageExecutionState originalState
  ) external {
    (bytes memory ccvArgs, bytes32[] memory rs, bytes32[] memory ss) =
      abi.decode(ccvData, (bytes, bytes32[], bytes32[]));

    _validateSignatures(keccak256(bytes.concat(messageHash, ccvArgs)), rs, ss);

    uint64 nonce = abi.decode(ccvArgs, (uint64));

    // Nonce changes per state transition (these only apply for ordered messages):
    // UNTOUCHED -> FAILURE  nonce bump.
    // UNTOUCHED -> SUCCESS  nonce bump.
    // FAILURE   -> SUCCESS  no nonce bump.
    // UNTOUCHED messages MUST be executed in order always.
    // If nonce == 0 then out of order execution is allowed.
    if (nonce != 0) {
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        // If a nonce is not incremented, that means it was skipped, and we can ignore the message.
        if (!INonceManager(i_nonceManager).incrementInboundNonce(message.sourceChainSelector, nonce, message.sender)) {
          revert InvalidNonce(nonce);
        }
      }
    }
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICCVOffRampV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
