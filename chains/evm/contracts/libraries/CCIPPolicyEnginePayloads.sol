// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Struct definitions for CCIP policy engine payloads.
/// @dev Zero-dependency library for external policy engines to decode policy data.
library CCIPPolicyEnginePayloads {
  struct OutboundPolicyDataV1 {
    bytes receiver; // Recipient of the tokens on the destination chain, abi encoded.
    uint64 remoteChainSelector; // ─╮ Destination chain selector.
    address originalSender; // ─────╯ Original sender of the tx on the source chain.
    uint256 amount; // Amount of tokens, denominated in the source token's decimals.
    address localToken; // Token address on this chain.
    uint16 blockConfirmationRequested; // Requested block confirmations.
    bytes tokenArgs; // Additional token-specific arguments.
  }

  struct InboundPolicyDataV1 {
    bytes originalSender; // Original sender of the tx on the source chain, abi encoded.
    uint64 remoteChainSelector; // ─╮ Source chain selector.
    address receiver; // ───────────╯ Recipient of the tokens on this chain.
    uint256 amount; // Amount of tokens, denominated in the source token's decimals.
    address localToken; // Token address on this chain.
    bytes sourcePoolAddress; // Source pool address, abi encoded.
    bytes sourcePoolData; // Data received from the source pool.
    bytes offchainTokenData; // Offchain attestation data.
    uint256 localAmount; // Amount of tokens, denominated in the local token's decimals.
    uint16 blockConfirmationRequested; // Requested block confirmations.
  }
}
