// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IPoolV2} from "./IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";

/// @notice Transfer validation and resolved configuration views for multi-token pools.
/// @dev Lane configuration resolves through token, covering group chain, then group default records.
/// Rate-limit consumption accepts calls only from the token's TokenAdminRegistry pool.
interface IPoolConfigManager {
  error OverflowDetected(uint8 remoteDecimals, uint8 localDecimals, uint256 remoteAmount);

  /// @notice The token operation used on send. Burn flavors mint on receive; Lock uses a lock box in both directions.
  enum LockOrBurnFlavor {
    Burn, //          burn(amount), as in BurnMintTokenPool.
    BurnFrom, //      burnFrom(address(this), amount), as in BurnFromMintTokenPool.
    BurnWithFrom, //  burn(address(this), amount), as in BurnWithFromMintTokenPool.
    BurnToAddress, // Transfer to an address that can never move the tokens, as in BurnToAddressMintTokenPool.
    Lock //           ILockBox.deposit and ILockBox.withdraw, as in LockReleaseTokenPool.
  }

  /// @notice Resolved outbound configuration and token fee for a transfer.
  // solhint-disable-next-line gas-struct-packing
  struct OutboundLane {
    bytes remoteToken; //                   The token on the remote chain, abi encoded for EVM chains.
    bool omitSourceTokenAddressOnSend; // ╮ True sends pool data without the source token.
    uint8 localDecimals; //               │ The decimals of the local token.
    LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the local token out.
    address lockOrBurnAddress; // ────────╯ BurnToAddress target, Lock box, zero otherwise.
    uint256 feeAmount; //                   The token fee. The pool keeps it and locks or burns the rest.
    address hooks; //                       The lane's IAdvancedPoolHooks, called preflight. Zero means no hooks.
  }

  /// @notice Resolved inbound transfer. The pool must validate source-token identity and its supported flavor.
  struct ResolvedInboundTransfer {
    bytes remoteToken; // The expected encoded source token address.
    uint256 localAmount; // The converted amount charged to the inbound bucket and delivered to the receiver.
    LockOrBurnFlavor lockOrBurnFlavor; // The token's configured send operation.
    address lockOrBurnAddress; // BurnToAddress target, Lock box, zero otherwise.
    bool allowSourceTokenAddressMissingOnReceive; // Whether legacy pool data without a source token is accepted.
    address hooks; // The postflight hooks to call. Zero means no hooks.
  }

  /// @notice Token bucket state in local token units, with the ABI layout of RateLimiter.TokenBucket.
  struct TokenBucket {
    uint128 tokens; // ────╮ The tokens in the bucket at this block.
    uint32 lastUpdated; // │ The time of the last refill, in seconds.
    bool isEnabled; // ────╯ True if the rate limit is enabled.
    uint128 capacity; // ──╮ The maximum number of tokens in the bucket.
    uint128 rate; // ──────╯ The tokens per second that refill the bucket.
  }

  /// @notice Resolved token-level configuration, independent of lane selection and pause state.
  struct TokenConfig {
    bool isRegistered; //                 True if the token is registered.
    uint8 localDecimals; //               The decimals of the local token. Zero if the token is not registered.
    LockOrBurnFlavor lockOrBurnFlavor; // How the pool moves the local token out. Burn if the token is not registered.
    bytes4 allowedFinalityConfig; //      The allowed finality config. Decode with FinalityCodec.
    uint64[] supportedChains; //          The remote chains covered by the token's resolved lanes.
  }

  /// @notice Validates an outbound transfer and consumes its rate limit after the token fee.
  /// @dev Checks effective outbound pause, lane existence, requested finality and available rate-limit capacity.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount before the token fee.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @return lane The resolved outbound configuration, token fee and hooks to call before moving tokens.
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bool chargeTokenFee
  ) external returns (OutboundLane memory lane);

  /// @notice Checks the inbound lane, converts the amount and consumes its rate limit.
  /// @param input The original inbound transfer.
  /// @param requestedFinalityConfig The requested finality.
  /// @return transfer The converted amount, expected source token, token profile and postflight hooks.
  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata input,
    bytes4 requestedFinalityConfig
  ) external returns (ResolvedInboundTransfer memory transfer);

  /// @notice Returns the resolved CCV list for a transfer direction.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @param direction The direction of the transfer.
  /// @return requiredCCVs The configured lists, or empty if no lane exists, aligned with localTokens by index.
  function getRequiredCCVs(
    address[] calldata localTokens,
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction
  ) external view returns (address[][] memory requiredCCVs);

  /// @notice Returns the resolved fee parameters for the requested finality.
  /// @param localTokens The local tokens.
  /// @param destChainSelector The destination chain.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @return feeUSDCents The flat fee in USD cents.
  /// @return destGasOverhead The destination token execution gas overhead.
  /// @return destBytesOverhead The destination token data byte overhead.
  /// @return tokenFeeBps The fee deducted from the transferred asset, in basis points.
  /// @return isEnabled Whether the pool fee is enabled. False requests FeeQuoter fallback.
  function getFee(
    address[] calldata localTokens,
    uint64 destChainSelector,
    bytes4 requestedFinalityConfig
  )
    external
    view
    returns (
      uint256[] memory feeUSDCents,
      uint32[] memory destGasOverhead,
      uint32[] memory destBytesOverhead,
      uint16[] memory tokenFeeBps,
      bool[] memory isEnabled
    );

  /// @notice Returns the resolved lane fee configuration.
  /// @param localTokens The local tokens.
  /// @param destChainSelector The destination chain.
  /// @return feeConfigs The configured fees, or a disabled zero configuration if no lane exists, aligned with
  /// localTokens by index.
  function getTokenTransferFeeConfig(
    address[] calldata localTokens,
    uint64 destChainSelector
  ) external view returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs);

  /// @notice Returns the token's current buckets for the requested finality class under the resolved lane.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @param fastFinality True for the FTF buckets, false for the default finality buckets.
  /// @return outbound The outbound buckets, aligned with localTokens by index.
  /// @return inbound The inbound buckets, aligned with localTokens by index.
  function getCurrentRateLimiterState(
    address[] calldata localTokens,
    uint64 remoteChainSelector,
    bool fastFinality
  ) external view returns (TokenBucket[] memory outbound, TokenBucket[] memory inbound);

  /// @notice Returns the resolved lane hooks.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @return hooks The hooks addresses, or zero if none are configured or no lane exists, aligned with localTokens by
  /// index.
  function getHooks(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (address[] memory hooks);

  /// @notice Returns the remote token from the resolved lane.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @return remoteTokens The encoded remote token addresses, or empty bytes if no lane exists, aligned with
  /// localTokens by index.
  function getRemoteToken(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (bytes[] memory remoteTokens);

  /// @notice Returns the trusted remote pools from the resolved lane.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @return remotePoolAddresses The encoded remote pool addresses, or an empty list if no lane or pools exist,
  /// aligned with localTokens by index.
  function getRemotePools(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (bytes[][] memory remotePoolAddresses);

  /// @notice Returns the effective outbound and inbound pause state for a token and remote chain.
  /// @param localTokens The local tokens.
  /// @param remoteChainSelector The remote chain.
  /// @return isOutboundPaused Whether outbound transfers are paused, aligned with localTokens by index.
  /// @return isInboundPaused Whether inbound transfers are paused, aligned with localTokens by index.
  function isTokenLanePaused(
    address[] calldata localTokens,
    uint64 remoteChainSelector
  ) external view returns (bool[] memory isOutboundPaused, bool[] memory isInboundPaused);

  /// @notice Returns the token's resolved configuration.
  /// @dev Returns the resolved value based on group setting / token override.
  /// @param localTokens The local tokens.
  /// @return configs The resolved token configurations, aligned with localTokens by index.
  function getTokenConfig(
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
