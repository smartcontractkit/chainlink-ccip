// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Struct definitions for CCIP policy engine payloads.
/// @dev Zero-dependency library for external policy engines to decode policy data.
library CCIPPolicyEnginePayloads {
  // bytes4(keccak256("PoolHookOutboundPolicyDataV1"))
  bytes4 public constant POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG = 0x12bebcb8;

  // bytes4(keccak256("PoolHookInboundPolicyDataV1"))
  bytes4 public constant POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG = 0x44d1de78;

  struct PoolHookOutboundPolicyDataV1 {
    address originalSender; // ─────────────╮ The original sender of the tx on the source chain.
    uint16 blockConfirmationRequested; //   │ The requested block confirmations.
    uint64 remoteChainSelector; // ─────────╯ The destination chain selector.
    bytes receiver; //                        The recipient of the tokens on the destination chain, abi-encoded.
    /// @dev NOTE: This is the user-entered transfer amount, excluding any source pool fees.
    /// It may not reflect the actual tokens sent to the destination.
    uint256 amount; //                        The amount of tokens, denominated in the source token's decimals.
    address localToken; //                    The token address on source chain.
    bytes tokenArgs; //                       Additional token-specific arguments.
  }

  struct PoolHookInboundPolicyDataV1 {
    bytes originalSender; //                  The original sender of the tx on the source chain, abi-encoded.
    uint16 blockConfirmationRequested; // ──╮ The requested block confirmations.
    uint64 remoteChainSelector; //          │ The source chain selector.
    address receiver; // ───────────────────╯ The recipient of the tokens on this chain.
    uint256 amount; //                        The amount of tokens, denominated in the source token's decimals.
    address localToken; //                    The token address on dest chain.
    bytes sourcePoolAddress; //               The source pool address, abi-encoded.
    bytes sourcePoolData; //                  The data received from the source pool.
    bytes offchainTokenData; //               The offchain attestation data.
    uint256 localAmount; //                   The amount of tokens, denominated in the local token's decimals.
  }
}
