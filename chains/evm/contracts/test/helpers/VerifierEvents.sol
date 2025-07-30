// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";

contract VerifierEvents {
  event MessageExecuted(
    Internal.Any2EVMMultiProofMessage message
  );

  uint64 public s_numMessagesExecuted;
  uint64 public s_numMessagesReExecuted;
  mapping(bytes32 => bool) public s_messageExecuted;

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

  function executeMessage(
    Internal.Any2EVMMultiProofMessage memory message
  ) external {
    if (s_messageExecuted[_hashMessage(message)]) {
      s_numMessagesReExecuted++;
    } else {
      s_numMessagesExecuted++;
    }
    s_messageExecuted[_hashMessage(message)] = true;
    emit MessageExecuted(message);
  }

  function _hashMessage(Internal.Any2EVMMultiProofMessage memory message) internal pure returns (bytes32) {
    return keccak256(abi.encode(message));
  }
}
