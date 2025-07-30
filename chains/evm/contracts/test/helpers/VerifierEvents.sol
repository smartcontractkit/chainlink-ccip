// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";

contract VerifierEvents {
  event CCIPMessageSent(
    uint64 indexed destChainSelector, uint64 indexed sequenceNumber, Internal.EVM2AnyCommitVerifierMessage message
  );

  function emitCCIPMessageSent(
    uint64 destChainSelector,
    uint64 sequenceNumber,
    Internal.EVM2AnyCommitVerifierMessage memory message
  ) external {
    emit CCIPMessageSent(destChainSelector, sequenceNumber, message);
  }

  function exposeAny2EVMMessage(
    Internal.Any2EVMMultiProofMessage memory message
  ) external pure {}
}
