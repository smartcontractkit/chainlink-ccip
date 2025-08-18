// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOffRamp} from "../interfaces/ICCVOffRamp.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";

import {Internal} from "../libraries/Internal.sol";
import {SignatureQuorumVerifier} from "../ocr/SignatureQuorumVerifier.sol";

contract CommitOffRamp is ICCVOffRamp, SignatureQuorumVerifier {
  error ZeroAddressNotAllowed();

  error InvalidNonce(uint64 nonce);

  address internal immutable i_nonceManager;

  constructor(
    address nonceManager
  ) {
    if (nonceManager == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_nonceManager = nonceManager;
  }

  function validateReport(
    bytes calldata rawMessage,
    bytes32 messageHash,
    bytes calldata ccvBlob,
    bytes calldata proof,
    Internal.MessageExecutionState originalState
  ) external {
    (bytes32 configDigest, uint64 nonce) = abi.decode(ccvBlob, (bytes32, uint64));

    _validateOCRSignatures(messageHash, configDigest, keccak256(ccvBlob), proof);

    // Nonce changes per state transition (these only apply for ordered messages):
    // UNTOUCHED -> FAILURE  nonce bump.
    // UNTOUCHED -> SUCCESS  nonce bump.
    // FAILURE   -> SUCCESS  no nonce bump.
    // UNTOUCHED messages MUST be executed in order always.
    // If nonce == 0 then out of order execution is allowed.
    if (nonce != 0) {
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        Internal.Any2EVMMessage memory message = abi.decode(rawMessage, (Internal.Any2EVMMessage));

        // If a nonce is not incremented, that means it was skipped, and we can ignore the message.
        if (
          !INonceManager(i_nonceManager).incrementInboundNonce(message.header.sourceChainSelector, nonce, message.sender)
        ) revert InvalidNonce(nonce);
      }
    }
  }
}
