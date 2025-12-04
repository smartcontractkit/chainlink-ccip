// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @notice Minimal Lombard bridge interface needed by the token pool and verifier.
interface IBridgeV2 {
  /// @notice Message version supported by the bridge.
  function MSG_VERSION() external view returns (uint8);

  /// @notice Address of the mailbox contract used to deliver payloads on the destination chain.
  function mailbox() external view returns (address);

  /// @notice Returns the allowed destination token identifier for a given chain and local token.
  /// @param destinationChainId Lombard chain identifier.
  /// @param localToken Address of the token on the source chain (or adapter).
  function getAllowedDestinationToken(bytes32 destinationChainId, address localToken) external view returns (bytes32);

  /// @notice Initiates a deposit into the Lombard bridge.
  /// @param destinationChain Destination Lombard chain identifier.
  /// @param token Token or adapter address on the source chain.
  /// @param sender Sender address on the source chain.
  /// @param recipient Recipient address on the destination chain, left padded to 32 bytes.
  /// @param amount Amount of tokens to bridge.
  /// @param destinationCaller Address allowed to handle the bridged payload on destination.
  /// @param payload Optional opaque payload hashed by the bridge.
  /// @return nonce Bridge-assigned deposit nonce.
  /// @return payloadHash Hash of the payload emitted by the bridge.
  function deposit(
    bytes32 destinationChain,
    address token,
    address sender,
    bytes32 recipient,
    uint256 amount,
    bytes32 destinationCaller,
    bytes calldata payload
  ) external returns (uint64 nonce, bytes32 payloadHash);
}
