// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPoolV1} from "./IPool.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";

/// @notice Shared public interface for multiple V2 pool types.
/// Each pool type handles a different child token model e.g. lock/release, mint/burn.
interface IPoolV2 is IPoolV1 {
  struct TokenTransferFeeConfig {
    uint32 destGasOverhead; // ───────────────────────╮ Gas charged to execute the token transfer on the destination chain.
    uint32 destBytesOverhead; //                      │ Data availability bytes.
    uint32 defaultBlockConfirmationFeeUSDCents; //    │ Fee to charge for token transfer with default block confirmation, multiples of 0.01 USD.
    uint32 customBlockConfirmationFeeUSDCents; //     │ Fee to charge for token transfer with custom block confirmation, multiples of 0.01 USD.
    //                                                │ The following two fee is deducted from the transferred asset, not added on top.
    uint16 defaultBlockConfirmationTransferFeeBps; // │ Fee in basis points for default finality transfers [0-10_000].
    uint16 customBlockConfirmationTransferFeeBps; // ─╯ Fee in basis points for custom finality transfers [0-10_000].
  }

  enum MessageDirection {
    Outbound,
    Inbound
  }

  /// @notice Lock tokens into the pool or burn the tokens.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token arguments.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  /// @return destTokenAmount The amount of tokens that will be set in TokenTransferV1.amount to be released/mint on destination.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount);

  /// @notice Releases or mints tokens on the destination chain.
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @return releaseOrMintOut Encoded data fields describing the result of the release or mint.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) external returns (Pool.ReleaseOrMintOutV1 memory releaseOrMintOut);

  /// @notice Returns the set of required CCVs for transfers in a given direction.
  /// @param localToken The address of the local token.
  /// @param remoteChainSelector The chain selector of the remote chain.
  /// @param amount The amount of tokens to be transferred.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param extraData Direction-specific payload forwarded by the caller (e.g. token args or source pool data).
  /// @param direction Whether CCVs are required for outbound (source -> remote) or inbound (remote -> destination) transfers.
  /// @return requiredCCVs A set of addresses representing the required CCVs.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16 blockConfirmationRequested,
    bytes calldata extraData,
    MessageDirection direction
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns the fee overrides for transferring the pool's token to a destination chain.
  /// @notice localToken The address of the local token.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param message The message to be sent to the destination chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token argument from the CCIP message.
  /// @return feeConfig the fee configuration for transferring the token to the destination chain.
  function getTokenTransferFeeConfig(
    address localToken,
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external view returns (TokenTransferFeeConfig memory feeConfig);

  /// @notice Returns the pool fee parameters that will apply to a transfer.
  /// @param localToken The local asset being transferred.
  /// @param destChainSelector The destination lane selector.
  /// @param amount The amount of tokens being bridged on this lane.
  /// @param feeToken The token used to pay feeUSDCents.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Opaque token arguments supplied by the caller.
  /// @return feeUSDCents Flat fee charged in USD cents (crumbs) for this transfer.
  /// @return destGasOverhead Destination gas charged for accounting in the cost model.
  /// @return destBytesOverhead Destination calldata size attributed to the transfer.
  /// @return tokenFeeBps Bps charged in token units. Value of zero implies no in-token fee.
  function getFee(
    address localToken,
    uint64 destChainSelector,
    uint256 amount,
    address feeToken,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external view returns (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps);

  /// @notice Gets the token address on the remote chain.
  /// @param remoteChainSelector Remote chain selector.
  /// @dev To support non-evm chains, this value is encoded into bytes.
  function getRemoteToken(
    uint64 remoteChainSelector
  ) external view returns (bytes memory);
}
