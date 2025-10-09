// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPoolV1} from "./IPool.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";

/// @notice Shared public interface for multiple V2 pool types.
/// Each pool type handles a different child token model e.g. lock/release, mint/burn.
interface IPoolV2 is IPoolV1 {
  struct TokenTransferFeeConfig {
    uint32 destGasOverhead; // ──────────────╮ Gas charged to execute the token transfer on the destination chain.
    uint32 destBytesOverhead; //             │ Data availability bytes.
    uint32 feeUSDCents; //                   │ Fee to charge per token transfer, multiples of 0.01 USD.
    uint16 customFinalityTransferFeeBps; //  │ Fee in basis points for custom finality transfers [0-10_000].
    uint16 defaultFinalityTransferFeeBps; // │ Fee in basis points for default finality transfers [0-10_000].
    bool isEnabled; // ──────────────────────╯ Whether this token has custom transfer fees.
  }

  /// @notice Lock tokens into the pool or burn the tokens.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token arguments.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  /// @return destTokenAmount The amount of tokens that will be set in TokenTransferV1.amount to be released/mint on destination.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality,
    bytes calldata tokenArgs
  ) external returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount);

  /// @notice Releases or mints tokens on the destination chain.
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  /// @param finality The finality configuration from the CCIP message.
  /// @return releaseOrMintOut Encoded data fields describing the result of the release or mint.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 finality
  ) external returns (Pool.ReleaseOrMintOutV1 memory releaseOrMintOut);

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param localToken The address of the local token.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token arguments.
  /// @return requiredCCVs A set of addresses representing the required outbound CCVs.
  function getRequiredOutboundCCVs(
    address localToken,
    uint64 destChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns the set of required CCVs for incoming messages from a source chain.
  /// @param localToken The address of the local token.
  /// @param sourceChainSelector The chain selector of the source chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param finality The finality configuration from the CCIP message.
  /// @param sourcePoolData The data received from the source pool to process the release or mint.
  /// @return requiredCCVs A set of addresses representing the required inbound CCVs.
  function getRequiredInboundCCVs(
    address localToken,
    uint64 sourceChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata sourcePoolData
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns the fee overrides for transferring the pool's token to a destination chain.
  /// @notice localToken The address of the local token.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param message The message to be sent to the destination chain.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token argument from the CCIP message.
  /// @return feeConfig the fee configuration for transferring the token to the destination chain.
  function getTokenTransferFeeConfig(
    address localToken,
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (TokenTransferFeeConfig memory feeConfig);
}
