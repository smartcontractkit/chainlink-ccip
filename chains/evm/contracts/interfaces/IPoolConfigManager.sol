// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IPoolV2} from "./IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";

/// @notice Transfer validation and resolved configuration views for multi-token pools.
/// @dev Rate-limit consumption is restricted to authorized callers.
interface IPoolConfigManager {
  /// @notice The token operation used on send. Burn flavors mint on receive; Lock releases tokens on receive.
  enum LockOrBurnFlavor {
    Burn, //          burn(amount).
    BurnFrom, //      burnFrom(address(this), amount).
    BurnWithFrom, //  burn(address(this), amount).
    BurnToAddress, // Transfer to an address that can never move the tokens.
    Lock //           Lock tokens on send and release them on receive.
  }

  /// @notice Resolved outbound configuration, pool data and token fee for a transfer.
  // solhint-disable-next-line gas-struct-packing
  struct OutboundLane {
    bytes remoteToken; //                   The token on the remote chain, abi encoded for EVM chains.
    bytes destPoolData; //                  The pool data for the destination pool.
    uint8 localDecimals; // ──────────────╮ The decimals of the local token.
    LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the local token out.
    address lockOrBurnAddress; // ────────╯ BurnToAddress target, Lock box, zero otherwise.
    uint256 feeAmount; //                   The token fee, in local token units.
    address hook; //                        The lane's IMultiTokenPoolHooks, called preflight. Zero means no hook.
  }

  /// @notice Resolved inbound transfer configuration.
  struct ResolvedInboundTransfer {
    uint256 localAmount; //                 The local token amount to be used for the inbound transfer.
    LockOrBurnFlavor lockOrBurnFlavor; // ╮ The token's type, used to decide if release or mint during inbound.
    address lockOrBurnAddress; // ────────╯ Used for lockbox address during inbound.
    address hook; //                        The postflight hook to call. Zero means no hook.
  }

  /// @notice Token bucket state in local token units.
  struct TokenBucket {
    uint128 tokens; // ────╮ The tokens in the bucket at this block.
    uint32 lastUpdated; // │ The time of the last refill, in seconds.
    bool isEnabled; // ────╯ True if the rate limit is enabled.
    uint128 capacity; // ──╮ The maximum number of tokens in the bucket.
    uint128 rate; // ──────╯ The tokens per second that refill the bucket.
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
    bool isEnabled; // ──────────╯ Whether pool fees are enabled. False requests default fees.
  }

  /// @notice Current buckets for both finality classes for one token and remote chain.
  struct LaneRateLimiterState {
    TokenBucket outbound; //             The outbound bucket for default finality.
    TokenBucket inbound; //              The inbound bucket for default finality.
    TokenBucket fastFinalityOutbound; // The outbound FTF bucket.
    TokenBucket fastFinalityInbound; //  The inbound FTF bucket.
  }

  /// @notice Resolved connection and effective pause state for one token and remote chain.
  struct LaneConnection {
    bool hasConnection; // ──────────────────────────╮ True if the token has a connection to the remote chain.
    bool isOutboundPaused; //                        │ Whether outbound transfers are paused.
    bool isInboundPaused; //                         │ Whether inbound transfers are paused.
    bool isGloballyPaused; //                        │ True if the global pause is on.
    bool isTokenOutboundPaused; //                   │ True if the token is paused outbound.
    bool isTokenInboundPaused; //                    │ True if the token is paused inbound.
    bool isChainOutboundPaused; //                   │ True if the remote chain is paused outbound.
    bool isChainInboundPaused; //                    │ True if the remote chain is paused inbound.
    bool omitSourceTokenAddressOnSend; //            │ True sends pool data without the source token.
    bool allowSourceTokenAddressMissingOnReceive; // ╯ True accepts pool data without the source token.
    bytes remoteToken; //                             The encoded remote token. Empty if it cannot be resolved.
    bytes[] remotePools; //                           The encoded trusted remote pools, or an empty list if none exist.
  }

  /// @notice Resolved token-level configuration, independent of lane selection and pause state.
  // solhint-disable-next-line gas-struct-packing
  struct TokenConfig {
    bool isRegistered; // ────────────────╮ True if the token is registered.
    uint8 localDecimals; //               │ The decimals of the local token. Zero if the token is not registered.
    LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the local token out. Burn if not registered.
    bytes4 allowedFinalityConfig; //      │ The allowed finality config. Decode with FinalityCodec.
    address lockOrBurnAddress; // ────────╯ BurnToAddress target, Lock box, zero otherwise.
    address additionalPoolDataProvider; //  The token's IAdditionalPoolDataProvider. Zero means none.
    uint64[] supportedChains; //            The remote chains covered by the token's resolved lanes.
    bytes32 connectionPolicyName; //        The connection policy of the token. Zero if it has none.
    bytes32 remotePoolPolicyName; //        The remote pool policy of the token. Zero if it has none.
    bytes32 rateLimitPolicyName; //         The rate limit policy of the token. Zero if it has none.
    bytes32 feePolicyName; //               The fee policy of the token. Zero if it has none.
    bytes32 destOverheadPolicyName; //      The destination overhead policy of the token. Zero if it has none.
    bytes32 securityPolicyName; //          The security policy of the token. Zero if it has none.
  }

  /// @notice Validates an outbound transfer, consumes its rate limit after the token fee and builds the pool data.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount before the token fee.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param tokenArgs The token arguments of the message. The sender chooses them.
  /// @param chargeTokenFee Whether to deduct a token fee from the transfer amount.
  /// @return lane The resolved outbound configuration, pool data, token fee and hook to call before moving tokens.
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    bool chargeTokenFee
  ) external returns (OutboundLane memory lane);

  /// @notice Checks the inbound lane, the source token and the additional pool data. Converts the amount to local
  /// decimals, and consumes the inbound rate limit.
  /// @param input The original inbound transfer.
  /// @param requestedFinalityConfig The requested finality.
  /// @return transfer The converted amount, token profile and postflight hook.
  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata input,
    bytes4 requestedFinalityConfig
  ) external returns (ResolvedInboundTransfer memory transfer);

  /// @notice Returns the resolved fee parameters for the requested finality.
  /// @dev Reverts if a token does not allow the requested finality, even when it has no lane.
  /// @param lanes The token and remote chain lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @return fees The fee parameters per lane, aligned with lanes by index. Lanes that do not exist, lanes without an
  /// enabled fee record and lanes without a dest overhead record return zeros with isEnabled false.
  function getFeeBatch(
    LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view returns (LaneFee[] memory fees);

  /// @notice Returns the resolved lane fee configurations.
  /// @param lanes The token and remote chain lanes to query.
  /// @return feeConfigs The fee configuration per lane, aligned with lanes by index. Lanes that do not exist, lanes
  /// without an enabled fee record and lanes without a dest overhead record return a disabled zero configuration.
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

  /// @notice Returns the resolved connections and effective pause state.
  /// @param lanes The token and remote chain lanes to query.
  /// @return connections The connection state per lane, aligned with lanes by index.
  function getLaneConnectionBatch(
    LaneKey[] calldata lanes
  ) external view returns (LaneConnection[] memory connections);

  /// @notice Returns the resolved hook and outbound and inbound CCV lists.
  /// @param lanes The token and remote chain lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @return hooks The lane's IMultiTokenPoolHooks per lane, aligned with lanes by index. Zero if no hook is configured
  /// or no lane exists.
  /// @return outboundCCVs The configured outbound CCVs per lane, aligned with lanes by index. Empty if no lane exists.
  /// @return inboundCCVs The configured inbound CCVs per lane, aligned with lanes by index. Empty if no lane exists.
  function getLaneVerificationBatch(
    LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view returns (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs);

  /// @notice Returns the token's resolved configuration.
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
