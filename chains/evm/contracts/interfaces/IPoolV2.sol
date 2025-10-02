// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPoolV1} from "./IPool.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";

/// @notice Shared public interface for multiple V2 pool types.
/// Each pool type handles a different child token model e.g. lock/release, mint/burn.
interface IPoolV2 is IPoolV1 {
  /// @notice Lock tokens into the pool or burn the tokens.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param finality The finality configuration from the CCIP message.
  /// @param tokenArgs Additional token arguments.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  /// @return destTokenAmount The amount of tokens that will be set in TokenTransferV1.amount to be r/m on destination.
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
  /// @return isEnabled Whether this token has transfer fee overrides.
  /// @return destGasOverhead Gas charged to execute the token transfer on the destination chain.
  /// @return destBytesOverhead Data availability bytes.
  /// @return feeUSDCents Fee to charge per token transfer, multiples of 0.01 USD.
  function getTokenTransferFeeConfig(
    address localToken,
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint16 finality,
    bytes calldata tokenArgs
  ) external view returns (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents);

  /// @notice Withdraws all accumulated pool fees to the specified recipient.
  /// @dev For burn/mint pools, this transfers the entire token balance of the pool contract.
  /// lock/release pools should override this function with their own accounting mechanism.
  /// @param recipient The address to receive the withdrawn fees.
  function withdrawPoolFees(
    address recipient
  ) external;

  /// @notice Gets the accumulated pool fees that can be withdrawn.
  /// @dev burn/mint pools should return the contract's token balance since pool fees
  /// are minted directly to the pool contract (e.g., `return getToken().balanceOf(address(this))`).
  /// lock/release pools should implement their own accounting mechanism for pool fees
  /// by adding a storage variable (e.g., `s_accumulatedPoolFees`) since they cannot mint
  /// additional tokens for pool fee rewards.
  /// Note: Fee accounting can be obscured by sending tokens directly to the pool.
  /// This does not introduce security issues but will need to be handled operationally.
  /// @return The amount of accumulated pool fees available for withdrawal.
  function getAccumulatedPoolFees() external returns (uint256);
}
