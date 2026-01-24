// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Struct definitions for CCIP policy engine payloads.
/// @dev Zero-dependency library for external policy engines to decode policy data.
library CCIPPolicyEnginePayloads {
  struct OutboundPolicyDataV1 {
    bytes receiver; //  The recipient of the tokens on the destination chain, abi-encoded.
    uint64 remoteChainSelector; // ─╮ The destination chain selector.
    address originalSender; // ─────╯ The original sender of the tx on the source chain.
    uint256 amount; //  The amount of tokens, denominated in the source token's decimals.
    address localToken; //  The token address on this chain.
    uint16 blockConfirmationRequested; //  The requested block confirmations.
    bytes tokenArgs; //  Additional token-specific arguments.
  }

  struct InboundPolicyDataV1 {
    bytes originalSender; //  The original sender of the tx on the source chain, abi-encoded.
    uint64 remoteChainSelector; // ─╮ The source chain selector.
    address receiver; // ───────────╯ The recipient of the tokens on this chain.
    uint256 amount; //  The amount of tokens, denominated in the source token's decimals.
    address localToken; //  The token address on this chain.
    bytes sourcePoolAddress; //  The source pool address, abi-encoded.
    bytes sourcePoolData; //  The data received from the source pool.
    bytes offchainTokenData; //  The offchain attestation data.
    uint256 localAmount; //  The amount of tokens, denominated in the local token's decimals.
    uint16 blockConfirmationRequested; //  The requested block confirmations.
  }
}
