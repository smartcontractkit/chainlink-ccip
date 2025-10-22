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
    uint32 defaultFinalityFeeUSDCents; //    │ Fee to charge for default finality token transfer, multiples of 0.01 USD.
    uint32 customFinalityFeeUSDCents; //     │ Fee to charge for custom finality token transfer, multiples of 0.01 USD.
    //                                       │ The following two fee is deducted from the transferred asset, not added on top.
    uint16 defaultFinalityTransferFeeBps; // │ Fee in basis points for default finality transfers [0-10_000].
    uint16 customFinalityTransferFeeBps; //  │ Fee in basis points for custom finality transfers [0-10_000].
    bool isEnabled; // ──────────────────────╯ Whether this token has custom transfer fees.
  }

  struct TokenTransferFeeDetails {
    // Zero values across the fee fields indicate that no fee is presently applied (either the config is disabled or the
    // configured rate/crumb is zero for the provided inputs).
    uint16 tokenFeeBps; // ╮ Basis points charged in token terms.
    uint32 usdFeeCents; // ╯ USD-denominated flat fee associated with the transfer.
    uint256 tokenFeeAmount; // Amount deducted from the transfer in token units.
    uint256 destinationAmount; // Amount that will arrive on the destination chain after fees denominated in the source token's decimals.
  }

  enum CCVDirection {
    Outbound,
    Inbound
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

  /// @notice Returns the set of required CCVs for transfers in a given direction.
  /// @param localToken The address of the local token.
  /// @param remoteChainSelector The chain selector of the remote chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param finality The finality configuration from the CCIP message.
  /// @param extraData Direction-specific payload forwarded by the caller (e.g. token args or source pool data).
  /// @param direction Whether CCVs are required for outbound (source -> remote) or inbound (remote -> destination) transfers.
  /// @return requiredCCVs A set of addresses representing the required CCVs.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata extraData,
    CCVDirection direction
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

  /// @notice Returns the token transfer fee details for a proposed transfer.
  /// @param localToken The address of the local token.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param amount The amount of tokens being transferred.
  /// @param finality The requested finality configuration.
  /// @param tokenArgs Additional token arguments supplied by the caller.
  /// @return feeDetails A struct describing the applied fees and destination amount. Zero-valued fields indicate that
  /// no fee is deducted for the requested parameters (e.g. disabled config or zero rates).
  function getTokenTransferFeeDetails(
    address localToken,
    uint64 destChainSelector,
    uint256 amount,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (TokenTransferFeeDetails memory feeDetails);
}
