// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

interface ICrossChainVerifierV1 is IERC165 {
  /// @notice Verification of the message, in any way the OffRamp wants. This could be using a signature, a quorum
  /// of signatures, using native interop, or some ZK light client. Any proof required for the verification is supplied
  /// through the ccvData parameter.
  /// @param originalCaller The original caller of verifyMessage, which is passed as input to enable proxy patterns.
  /// @param message The message to be verified. For efficiency, the messageID is also supplied, which acts as a small
  /// payload that once verified means the entire message is verified. Every component of the message is part of the
  /// message ID through hashing the struct. The entire message is provided to be able to act differently for different
  /// message properties.
  /// @param messageId A convenient 32 byte hash of the entire message. It can be recomputed from the passed in message
  /// at the cost of a not-insignificant amount of gas. Any CCV MUST include the messageID or the entire message struct
  /// as part of its proof.
  /// @param ccvData All the data that is specific to the CCV. This often means it contains some sort of proof, but it
  /// can also contain certain metadata like a nonce that's specific to the CCV. If any metadata like that exists and is
  /// important to the security of the CCV, it MUST be verified as well using the proof. A recommended way to do this is
  /// to encode a proof and the metadata separately in the ccvData and then concatenate the messageId with this metadata
  /// to get the payload that will be verified. In the case of a simple signature verification this means that the CCV
  /// offchain system must sign the concatenated (messageId, ccvMetaData) and not just the messageId. If no metadata
  /// is required, simply signing the messageId is enough.
  function verifyMessage(
    address originalCaller,
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageId,
    bytes memory ccvData
  ) external;

  /// @notice Quotes the fee for a CCIP message to a destination chain.
  /// @dev This takes EVM2AnyMessage (instead of MessageV1) because
  /// the router client API that user contracts interact with (IRouterClient.getFee)
  /// exposes EVM2AnyMessage. The on-ramp can translate to MessageV1 internally
  /// where required (e.g., verifier hooks), but using EVM2AnyMessage here keeps the
  /// interface aligned with what clients construct and pass to the router.
  /// @param originalCaller The original caller of getFee.
  /// @param destChainSelector The destination chain selector of the message.
  /// @param message The message to be sent.
  /// @param extraArgs Opaque extra args that can be used by the fee quoter
  function getFee(
    address originalCaller,
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256);

  /// @notice Message sending, verifier hook.
  /// @param originalCaller The original caller of forwardToVerifier.
  /// @param message Decoded MessageV1 structure for the message being sent.
  /// @param messageId The message ID of the message being sent.
  /// @param feeToken Fee token used for this message.
  /// @param feeTokenAmount Amount of fee token provided.
  /// @param verifierArgs Opaque verifier-specific arguments from the sender.
  /// @return verifierData Verifier-specific return data blob.
  function forwardToVerifier(
    address originalCaller,
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    address feeToken,
    uint256 feeTokenAmount,
    bytes calldata verifierArgs
  ) external returns (bytes memory verifierData);

  /// @notice Returns the storage location identifier for this CCV. This is a string that uniquely identifies the
  /// storage location. This can be an address, a URL, or any other identifier that makes sense for the CCV. The format
  /// of the string is up to the CCV implementer, but it should be something that can be easily parsed and used by the
  /// integrator. This is used by the executor(s) to know where to look for the proof data that the CCV has produced.
  function getStorageLocation() external view returns (string memory);
}
