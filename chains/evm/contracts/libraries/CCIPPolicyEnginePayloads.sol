// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Struct definitions for CCIP policy engine payloads.
library CCIPPolicyEnginePayloads {
  // bytes4(keccak256("PoolHookOutboundPolicyDataV1"))
  bytes4 public constant POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG = 0x12bebcb8;

  // bytes4(keccak256("PoolHookInboundPolicyDataV1"))
  bytes4 public constant POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG = 0x44d1de78;

  struct PoolHookOutboundPolicyDataV1 {
    address originalSender; // ─────────────╮ The The message sender on the source chain.
    uint16 blockConfirmationRequested; //   │ The block confirmation requested.
    uint64 remoteChainSelector; // ─────────╯ The destination chain selector.
    bytes receiver; //                        The recipient of the tokens on the destination chain, abi-encoded if destination chain is EVM.
    /// @dev NOTE: This is the user-entered transfer amount, excluding any source pool fees.
    /// It may not reflect the actual tokens sent to the destination.
    uint256 amount; //                        The amount of tokens, denominated in the source token's decimals.
    address localToken; //                    The token address on source chain.
    bytes tokenArgs; //                       Additional token-specific arguments.
  }

  struct PoolHookInboundPolicyDataV1 {
    bytes originalSender; //                  The The message sender on the source chain.
    uint16 blockConfirmationRequested; // ──╮ The block confirmation requested.
    uint64 remoteChainSelector; //          │ The source chain selector.
    address receiver; // ───────────────────╯ The recipient of the tokens on this chain.
    uint256 amount; //                        The amount of tokens to release or mint, denominated in the source token's decimals.
    address localToken; //                    The token address on dest chain.
    bytes sourcePoolAddress; //               The address of the source pool, abi encoded in the case of EVM chains.
    bytes sourcePoolData; //                  The data received from the source pool to process the release or mint.
    bytes offchainTokenData; //               The offchain data to process the release or mint.
    uint256 localAmount; //                   The local amount of tokens on dest chain to be released or minted.
  }
}
