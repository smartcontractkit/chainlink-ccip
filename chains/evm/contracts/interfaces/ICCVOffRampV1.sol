// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

interface ICCVOffRampV1 is IERC165 {
  /// @notice Verification of the message, in any way the OffRamp wants. This could be using a signature, a quorum
  /// of signatures, using native interop, or some ZK light client. Any proof required for the verification is supplied
  /// through the ccvData parameter.
  /// @param message The message to be verified. For efficiency, the messageID is also supplied, which acts as a small
  /// payload that once verified means the entire message is verified. Every component of the message is part of the
  /// message ID through hashing the struct. The entire message is provided to be able to act differently for different
  /// message properties.
  /// @param messageId A convenient 32 byte hash of the entire message. It can be recomputed from the passed in message
  /// at the cost of a not-insignificant amount of gas. Any CCV MUST verify this as part of this call.
  /// @param ccvData All the data that is specific to the CCV. This often means it contains some sort of proof, but it
  /// can also contain certain metadata like a nonce that's specific to the CCV. If any metadata like that exists and is
  /// important to the security of the CCV, it MUST be verified as well using the proof. A recommended way to do this is
  /// to encode a proof and the metadata separately in the ccvData and then concatenate the messageId with this metadata
  /// to get the payload that will be verified. In the case of a simple signature verification this means that the CCV
  /// offchain system must sign the concatenated (messageId, ccvMetaData) and not just the messageId. If no metadata
  /// is required, simply signing the messageId is enough.
  function verifyMessage(MessageV1Codec.MessageV1 memory message, bytes32 messageId, bytes memory ccvData) external;
}
