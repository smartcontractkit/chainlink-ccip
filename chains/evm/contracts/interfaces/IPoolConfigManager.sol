// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IPoolV2} from "./IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";

/// @notice Transfer validation and resolved configuration views for multi-token pools.
/// @dev Lane configuration resolves through token, then policy x chain, then policy default records.
/// Rate-limit consumption accepts calls only from the token's TokenAdminRegistry pool.
/// Views can return values without token registration or a record for the token and chain.
/// For example, the global pause applies to all tokens and chains.
interface IPoolConfigManager {
  /// @notice The token operation used on send. Burn flavors mint on receive; Lock uses a lock box in both directions.
  enum LockOrBurnFlavor {
    Burn, //          burn(amount), as in BurnMintTokenPool.
    BurnFrom, //      burnFrom(address(this), amount), as in BurnFromMintTokenPool.
    BurnWithFrom, //  burn(address(this), amount), as in BurnWithFromMintTokenPool.
    BurnToAddress, // Transfer to an address that can never move the tokens, as in BurnToAddressMintTokenPool.
    Lock //           ILockBox.deposit and ILockBox.withdraw, as in LockReleaseTokenPool.
  }

  /// @notice Resolved outbound configuration, pool data and token fee for a transfer.
  // solhint-disable-next-line gas-struct-packing
  struct OutboundLane {
    bytes remoteToken; //                   The token on the remote chain, abi encoded for EVM chains.
    bytes destPoolData; //                  The pool data for the destination pool.
    uint8 localDecimals; // ──────────────╮ The decimals of the local token.
    LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the local token out.
    address lockOrBurnAddress; // ────────╯ BurnToAddress target, Lock box, zero otherwise.
    uint256 feeAmount; //                   The token fee. The pool keeps it and locks or burns the rest.
    address hook; //                        The lane's IAdvancedPoolHooks, called preflight. Zero means no hook.
  }

  /// @notice Resolved inbound transfer. The pool must validate its supported flavor.
  struct ResolvedInboundTransfer {
    uint256 localAmount; //                 The local token amount to be used for the inbound transfer.
    LockOrBurnFlavor lockOrBurnFlavor; // ╮ The token's type, used to decide if release or mint during inbound.
    address lockOrBurnAddress; // ────────╯ Used for lockbox address during inbound.
    address hook; //                        The postflight hook to call. Zero means no hook.
  }

  /// @notice Token bucket state in local token units, with the ABI layout of RateLimiter.TokenBucket.
  struct TokenBucket {
    uint128 tokens; // ────╮ The tokens in the bucket at this block.
    uint32 lastUpdated; // │ The time of the last refill, in seconds.
    bool isEnabled; // ────╯ True if the rate limit is enabled.
    uint128 capacity; // ──╮ The maximum number of tokens in the bucket.
    uint128 rate; // ──────╯ The tokens per second that refill the bucket.
  }

  /// @notice Current outbound and inbound buckets for one finality class.
  struct RateLimiterState {
    TokenBucket outbound; // The outbound bucket.
    TokenBucket inbound; //  The inbound bucket.
  }

  /// @notice Identifies one token and remote chain lane in batch queries.
  struct LaneKey {
    address localToken; //         The local token.
    uint64 remoteChainSelector; // The remote chain.
  }

  /// @notice Resolved fee parameters for one token and destination.
  struct LaneFee {
    uint256 feeUSDCents; //        The flat fee in USD cents.
    uint32 destGasOverhead; // ──╮ The destination token execution gas overhead.
    uint32 destBytesOverhead; // │ The destination token data byte overhead.
    uint16 tokenFeeBps; //       │ The fee deducted from the transferred asset, in basis points.
    bool isEnabled; // ──────────╯ Whether the pool fee is enabled. False requests FeeQuoter fallback.
  }

  /// @notice Current buckets for both finality classes for one token and remote chain.
  struct LaneRateLimiterState {
    RateLimiterState defaultFinality; // The default finality buckets.
    RateLimiterState fastFinality; //    The FTF buckets.
  }

  /// @notice Resolved connection and chain records and effective pause state for one token and remote chain.
  struct LaneConnection {
    bool hasConnection; // ──────────────────────────╮ True if a connection record applies to the token and remote chain.
    bool isOutboundPaused; //                        │ Whether outbound transfers are paused.
    bool isInboundPaused; //                         │ Whether inbound transfers are paused.
    bool omitSourceTokenAddressOnSend; //            │ True sends pool data without the source token.
    bool allowSourceTokenAddressMissingOnReceive; // │ True accepts pool data without the source token.
    uint32 destGasOverhead; //                       │ The destination token execution gas overhead.
    uint32 destBytesOverhead; // ────────────────────╯ The destination token data byte overhead.
    bytes remoteToken; //                             The encoded remote token, or empty bytes if no lane exists.
    bytes[] remotePools; //                           The encoded trusted remote pools, or an empty list if none exist.
  }

  /// @notice Resolved verification configuration for one token and remote chain.
  struct LaneVerification {
    address hook; //           The lane's IAdvancedPoolHooks, or zero if no hook is configured or no lane exists.
    address[] outboundCCVs; // The configured outbound CCVs, or empty if no lane exists.
    address[] inboundCCVs; //  The configured inbound CCVs, or empty if no lane exists.
  }

  /// @notice Resolved token-level configuration, independent of lane selection and pause state.
  // solhint-disable-next-line gas-struct-packing
  struct TokenConfig {
    bool isRegistered; // ────────────────╮ True if the token is registered.
    uint8 localDecimals; //               │ The decimals of the local token. Zero if the token is not registered.
    LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the local token out. Burn if not registered.
    bytes4 allowedFinalityConfig; //      │ The allowed finality config. Decode with FinalityCodec.
    address lockOrBurnAddress; // ────────╯ BurnToAddress target, Lock box, zero otherwise.
    address tokenIdentityMetadataProvider; // The token's ITokenIdentityMetadataProvider. Zero means none.
    uint64[] supportedChains; //            The remote chains covered by the token's resolved lanes.
  }

  /// @notice Validates an outbound transfer, consumes its rate limit after the token fee and builds the pool data.
  /// @dev Checks effective outbound pause, lane existence, requested finality and available rate-limit capacity.
  /// Pool data with the source token also carries the metadata from the token's ITokenIdentityMetadataProvider, or
  /// empty metadata if the token has none.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount before the token fee.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @return lane The resolved outbound configuration, pool data, token fee and hook to call before moving tokens.
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bool chargeTokenFee
  ) external returns (OutboundLane memory lane);

  /// @notice Checks the inbound lane, source token and token identity metadata, converts token transfer amount, and
  /// consumes inbound rate limit.
  /// @dev Checks effective inbound pause, lane existence, source pool, source pool data, source-token identity, token
  /// identity metadata and available rate-limit capacity. Source pool data without the source token is accepted only if
  /// the lane allows it. The token's ITokenIdentityMetadataProvider, if any, gets the metadata of every transfer, also
  /// empty metadata. A token without one ignores the metadata.
  /// @param input The original inbound transfer.
  /// @param requestedFinalityConfig The requested finality.
  /// @return transfer The converted amount, token profile and postflight hook.
  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata input,
    bytes4 requestedFinalityConfig
  ) external returns (ResolvedInboundTransfer memory transfer);

  /// @notice Returns the resolved CCV list for a transfer direction.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  /// @param direction The direction of the transfer.
  /// @return requiredCCVs The configured list, or empty if no lane exists.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Returns the resolved fee parameters for the requested finality.
  /// @param localToken The local token.
  /// @param destChainSelector The destination chain.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @dev Reverts with FinalityCodec.InvalidRequestedFinality if the token profile does not allow the requested
  /// finality, also for a token without a lane. This is the only PCM view that reverts.
  /// @return fee The resolved fee params, or zeros with isEnabled false if the fee record is disabled or the token
  /// has no lane or chain record.
  function getFee(
    address localToken,
    uint64 destChainSelector,
    bytes4 requestedFinalityConfig
  ) external view returns (LaneFee memory fee);

  /// @notice Returns the resolved fee parameters for the requested finality. Batched version of getFee.
  /// @param lanes The token and remote chain lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @return fees The fee parameters per lane, aligned with lanes by index. Lanes that do not exist, or whose fee
  /// record is disabled, return zeros with isEnabled false.
  function getFeeBatch(
    LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view returns (LaneFee[] memory fees);

  /// @notice Returns the resolved lane fee configurations.
  /// @param lanes The token and remote chain lanes to query.
  /// @return feeConfigs The fee configuration per lane, aligned with lanes by index. Lanes that do not exist return
  /// a disabled zero configuration.
  function getTokenTransferFeeConfigBatch(
    LaneKey[] calldata lanes
  ) external view returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs);

  /// @notice Returns the current buckets for both finality classes under each resolved lane.
  /// @param lanes The token and remote chain lanes to query.
  /// @return states The rate limiter state per lane, aligned with lanes by index. Lanes that do not exist return
  /// zero values.
  function getCurrentRateLimiterStateBatch(
    LaneKey[] calldata lanes
  ) external view returns (LaneRateLimiterState[] memory states);

  /// @notice Returns the resolved connection and chain records and the effective pause state.
  /// @dev A direction is paused if the global pause, token pause or chain pause applies to it.
  /// Pause flags do not depend on token registration or connection records.
  /// @param lanes The token and remote chain lanes to query.
  /// @return connections The connection state per lane, aligned with lanes by index.
  function getLaneConnectionBatch(
    LaneKey[] calldata lanes
  ) external view returns (LaneConnection[] memory connections);

  /// @notice Returns the resolved hook and outbound and inbound CCV lists.
  /// @param lanes The token and remote chain lanes to query.
  /// @return verifications The verification configuration per lane, aligned with lanes by index. Lanes that do
  /// not exist return zero values.
  function getLaneVerificationBatch(
    LaneKey[] calldata lanes
  ) external view returns (LaneVerification[] memory verifications);

  /// @notice Returns the token's resolved configuration.
  /// @dev Returns the resolved value based on policy setting / token override.
  /// @param localTokens The local tokens.
  /// @return configs The resolved token configurations, aligned with localTokens by index.
  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view returns (TokenConfig[] memory configs);

  /// @notice Returns one page of tokens registered with this pool config manager.
  /// @param startIndex The position of the first token to return.
  /// @param maxCount The maximum number of tokens to return.
  /// @return tokens The tokens. Fewer than maxCount only at the end of the list.
  function getAllConfiguredTokens(
    uint64 startIndex,
    uint64 maxCount
  ) external view returns (address[] memory tokens);

  /// @notice Returns the TokenAdminRegistry that this pool config manager uses.
  function getTokenAdminRegistry() external view returns (address tokenAdminRegistry);
}
