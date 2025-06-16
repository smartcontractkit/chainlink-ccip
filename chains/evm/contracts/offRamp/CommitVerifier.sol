// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import {INonceManager} from "../interfaces/INonceManager.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";

import {Internal} from "../libraries/Internal.sol";
import {OCRVerifier} from "../ocr/OCRVerifier.sol";

contract CommitVerifier is IVerifier, OCRVerifier {
  error ZeroAddressNotAllowed();

  error InvalidNonce(uint64 nonce);

  address internal i_nonceManager;

  constructor(
    address nonceManager
  ) {
    if (nonceManager == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_nonceManager = nonceManager;
  }

  function validateReport(
    bytes calldata rawReport,
    bytes calldata ocrProof,
    uint256 verifierIndex,
    Internal.MessageExecutionState originalState
  ) external {
    _validateOCRSignatures(rawReport, ocrProof);

    Internal.Any2EVMMultiProofMessage memory message = abi.decode(rawReport, (Internal.Any2EVMMultiProofMessage));

    // Nonce changes per state transition (these only apply for ordered messages):
    // UNTOUCHED -> FAILURE  nonce bump.
    // UNTOUCHED -> SUCCESS  nonce bump.
    // FAILURE   -> SUCCESS  no nonce bump.
    // UNTOUCHED messages MUST be executed in order always.
    // If nonce == 0 then out of order execution is allowed.
    if (message.header.nonce != 0) {
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        // If a nonce is not incremented, that means it was skipped, and we can ignore the message.
        if (
          !INonceManager(i_nonceManager).incrementInboundNonce(
            message.header.sourceChainSelector, message.header.nonce, message.sender
          )
        ) revert InvalidNonce(message.header.nonce);
      }
    }
  }
}
