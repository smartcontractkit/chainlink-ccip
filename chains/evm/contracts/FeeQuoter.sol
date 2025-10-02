// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoter} from "./interfaces/IFeeQuoter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "./libraries/Client.sol";
import {Internal} from "./libraries/Internal.sol";
import {Pool} from "./libraries/Pool.sol";
import {USDPriceWith18Decimals} from "./libraries/USDPriceWith18Decimals.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice The FeeQuoter contract responsibility is to:
///   - Store the current gas price in USD for a given destination chain.
///   - Store the price of a token in USD allowing the owner or priceUpdater to update this value.
///   - Manage chain specific fee calculations.
/// The authorized callers in the contract represent the fee price updaters.
contract FeeQuoter is AuthorizedCallers, IFeeQuoter, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.AddressSet;
  using USDPriceWith18Decimals for uint224;

  error TokenNotSupported(address token);
  error FeeTokenNotSupported(address token);
  error StaleGasPrice(uint64 destChainSelector, uint256 threshold, uint256 timePassed);
  error InvalidDestBytesOverhead(address token, uint32 destBytesOverhead);
  error MessageGasLimitTooHigh();
  error MessageComputeUnitLimitTooHigh();
  error DestinationChainNotEnabled(uint64 destChainSelector);
  error InvalidExtraArgsTag();
  error InvalidExtraArgsData();
  error SourceTokenDataTooLarge(address token);
  error InvalidDestChainConfig(uint64 destChainSelector);
  error MessageFeeTooHigh(uint256 msgFeeJuels, uint256 maxFeeJuelsPerMsg);
  error InvalidStaticConfig();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnsupportedNumberOfTokens(uint256 numberOfTokens, uint256 maxNumberOfTokensPerMsg);
  error InvalidChainFamilySelector(bytes4 chainFamilySelector);
  error InvalidTokenReceiver();
  error TooManySVMExtraArgsAccounts(uint256 numAccounts, uint256 maxAccounts);
  error InvalidSVMExtraArgsWritableBitmap(uint64 accountIsWritableBitmap, uint256 numAccounts);
  error TooManySuiExtraArgsReceiverObjectIds(uint256 numReceiverObjectIds, uint256 maxReceiverObjectIds);

  event FeeTokenAdded(address indexed feeToken);
  event FeeTokenRemoved(address indexed feeToken);
  event UsdPerUnitGasUpdated(uint64 indexed destChain, uint256 value, uint256 timestamp);
  event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);
  event TokenTransferFeeConfigUpdated(
    uint64 indexed destChainSelector, address indexed token, TokenTransferFeeConfig tokenTransferFeeConfig
  );
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector, address indexed token);
  event PremiumMultiplierWeiPerEthUpdated(address indexed token, uint64 premiumMultiplierWeiPerEth);
  event DestChainConfigUpdated(uint64 indexed destChainSelector, DestChainConfig destChainConfig);
  event DestChainAdded(uint64 indexed destChainSelector, DestChainConfig destChainConfig);

  /// @dev Struct that contains the static configuration.
  /// RMN depends on this struct, if changing, please notify the RMN maintainers.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint96 maxFeeJuelsPerMsg; // ─╮ Maximum fee that can be charged for a message.
    address linkToken; // ────────╯ LINK token address.
  }

  /// @dev Struct to hold the fee & validation configs for a destination chain.
  // solhint-disable gas-struct-packing
  struct DestChainConfig {
    bool isEnabled; // ────────────────────╮ Whether this destination chain is enabled.
    uint32 maxDataBytes; //                │ Maximum data payload size in bytes.
    uint32 maxPerMsgGasLimit; //           │ Maximum gas limit for messages targeting EVMs.
    uint32 destGasOverhead; //             │ Gas charged on top of the gasLimit to cover destination chain costs.
    uint8 destGasPerPayloadByteBase; //    │ Default dest-chain gas charged each byte of `data` payload.
    bytes4 chainFamilySelector; //         │ Selector that identifies the destination chain's family. Used to determine the correct validations to perform for the dest chain.
    // The following three properties are defaults, they can be overridden by setting the TokenTransferFeeConfig for a token.
    uint16 defaultTokenFeeUSDCents; //     │ Default token fee charged per token transfer.
    uint32 defaultTokenDestGasOverhead; // │ Default gas charged to execute a token transfer on the destination chain.
    uint32 defaultTxGasLimit; //           │ Default gas limit for a tx.
    uint32 networkFeeUSDCents; // ─────────╯ Flat network fee to charge for messages, multiples of 0.01 USD.
  }

  /// @dev Struct to hold the configs and its destination chain selector. Same as DestChainConfig but with the
  /// destChainSelector so that an array of these can be passed in the constructor and applyDestChainConfigUpdates.
  /// solhint-disable gas-struct-packing
  struct DestChainConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    DestChainConfig destChainConfig; // Config to update for the chain selector.
  }

  /// @dev No need to write for future non-EVMs that only support 1.7. This is a legacy struct for pre 1.7 pools.
  struct TokenTransferFeeConfig {
    uint32 feeUSDCents; // ──────╮ Minimum fee to charge per token transfer, multiples of 0.01 USD.
    uint32 destGasOverhead; //   │ Gas charged to execute the token transfer on the destination chain.
    //                           │ Data availability bytes that are returned from the source pool and sent to the dest
    uint32 destBytesOverhead; // │ pool. Must be >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES. Set as multiple of 32 bytes.
    bool isEnabled; // ──────────╯ Whether this token has custom transfer fees.
  }

  /// @dev Struct with token transfer fee configurations for a token, same as TokenTransferFeeConfig but with the token
  /// address included.
  struct TokenTransferFeeConfigSingleTokenArgs {
    address token; // Token address.
    TokenTransferFeeConfig tokenTransferFeeConfig; // Struct to hold the transfer fee configuration for token transfers.
  }

  /// @dev Struct with args for setting the token transfer fee configurations for a destination chain and a set of tokens.
  struct TokenTransferFeeConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    TokenTransferFeeConfigSingleTokenArgs[] tokenTransferFeeConfigs; // Array of token transfer fee configurations.
  }

  /// @dev Struct with a pair of destination chain selector and token address so that an array of these can be passed in
  /// the applyTokenTransferFeeConfigUpdates function to remove the token transfer fee configuration for a token.
  struct TokenTransferFeeConfigRemoveArgs {
    uint64 destChainSelector; // ─╮ Destination chain selector.
    address token; // ────────────╯ Token address.
  }

  /// @dev Struct with fee token configuration for a token.
  struct PremiumMultiplierWeiPerEthArgs {
    address token; // // ──────────────────╮ Token address.
    uint64 premiumMultiplierWeiPerEth; // ─╯ Multiplier for destination chain specific premiums.
  }

  string public constant override typeAndVersion = "FeeQuoter 1.6.3-dev";

  /// @dev The gas price per unit of gas for a given destination chain, in USD with 18 decimals. Multiple gas prices can
  /// be encoded into the same value. Each price takes {Internal.GAS_PRICE_BITS} bits. For example, if Optimism is the
  /// destination chain, gas price can include L1 base fee and L2 gas price. Logic to parse the price components is
  ///  chain-specific, and should live in OnRamp.
  /// @dev Price of 1e18 is 1 USD. Examples:
  ///     Very Expensive:   1 unit of gas costs 1 USD                  -> 1e18.
  ///     Expensive:        1 unit of gas costs 0.1 USD                -> 1e17.
  ///     Cheap:            1 unit of gas costs 0.000001 USD           -> 1e12.
  mapping(uint64 destChainSelector => Internal.TimestampedPackedUint224 price) private
    s_usdPerUnitGasByDestChainSelector;

  /// @dev The price, in USD with 18 decimals, per 1e18 of the smallest token denomination.
  /// @dev Price of 1e18 represents 1 USD per 1e18 token amount.
  ///     1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30.
  ///     1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18.
  ///     1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18.
  mapping(address token => Internal.TimestampedPackedUint224 price) private s_usdPerToken;

  /// @dev The multiplier for destination chain specific premiums that can be set by the owner or fee admin.
  mapping(address token => uint64 premiumMultiplierWeiPerEth) private s_premiumMultiplierWeiPerEth;

  /// @dev The destination chain specific fee configs.
  mapping(uint64 destChainSelector => DestChainConfig destChainConfig) internal s_destChainConfigs;

  /// @dev The token transfer fee config that can be set by the owner or fee admin.
  mapping(uint64 destChainSelector => mapping(address token => TokenTransferFeeConfig tranferFeeConfig)) internal
    s_tokenTransferFeeConfig;

  /// @dev Maximum fee that can be charged for a message. This is a guard to prevent massively overcharging due to
  /// misconfiguration.
  uint96 internal immutable i_maxFeeJuelsPerMsg;
  /// @dev The link token address.
  address internal immutable i_linkToken;

  /// @dev Subset of tokens which prices tracked by this registry which are fee tokens.
  EnumerableSet.AddressSet private s_feeTokens;

  constructor(
    StaticConfig memory staticConfig,
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs,
    DestChainConfigArgs[] memory destChainConfigArgs
  ) AuthorizedCallers(priceUpdaters) {
    if (staticConfig.linkToken == address(0) || staticConfig.maxFeeJuelsPerMsg == 0) {
      revert InvalidStaticConfig();
    }

    i_linkToken = staticConfig.linkToken;
    i_maxFeeJuelsPerMsg = staticConfig.maxFeeJuelsPerMsg;

    _applyFeeTokensUpdates(new address[](0), feeTokens);
    _applyDestChainConfigUpdates(destChainConfigArgs);
    _applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);
    _applyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs, new TokenTransferFeeConfigRemoveArgs[](0));
  }

  // ================================================================
  // │                     Price calculations                       │
  // ================================================================

  /// @inheritdoc IFeeQuoter
  /// @dev returns the price even if it's stale or zero.
  function getTokenPrice(
    address token
  ) public view override returns (Internal.TimestampedPackedUint224 memory) {
    return s_usdPerToken[token];
  }

  /// @notice Get the `tokenPrice` for a given token, checks if the price is valid.
  /// @param token The token to get the price for.
  /// @return tokenPrice The tokenPrice for the given token if it exists and is valid.
  function getValidatedTokenPrice(
    address token
  ) external view returns (uint224) {
    return _getValidatedTokenPrice(token);
  }

  /// @notice Get the `tokenPrice` for an array of tokens.
  /// @param tokens The tokens to get prices for.
  /// @return tokenPrices The tokenPrices for the given tokens.
  function getTokenPrices(
    address[] calldata tokens
  ) external view returns (Internal.TimestampedPackedUint224[] memory) {
    uint256 length = tokens.length;
    Internal.TimestampedPackedUint224[] memory tokenPrices = new Internal.TimestampedPackedUint224[](length);
    for (uint256 i = 0; i < length; ++i) {
      tokenPrices[i] = getTokenPrice(tokens[i]);
    }
    return tokenPrices;
  }

  /// @notice Get an encoded `gasPrice` for a given destination chain ID.
  /// The 224-bit result encodes necessary gas price components.
  /// - On L1 chains like Ethereum or Avax, the only component is the gas price.
  /// - On Optimistic Rollups, there are two components - the L2 gas price, and L1 base fee for data availability.
  /// - On future chains, there could be more or differing price components.
  /// @param destChainSelector The destination chain to get the price for.
  /// @return gasPrice The encoded gasPrice for the given destination chain ID.
  /// @dev Does not validate if the chain is enabled
  function getDestinationChainGasPrice(
    uint64 destChainSelector
  ) external view returns (Internal.TimestampedPackedUint224 memory) {
    return s_usdPerUnitGasByDestChainSelector[destChainSelector];
  }

  /// @notice Gets the fee token price and the gas price, both denominated in dollars.
  /// @param token The source token to get the price for.
  /// @param destChainSelector The destination chain to get the gas price for.
  /// @return tokenPrice The price of the feeToken in 1e18 dollars per base unit.
  /// @return gasPriceValue The price of gas in 1e18 dollars per base unit.
  function getTokenAndGasPrices(
    address token,
    uint64 destChainSelector
  ) external view returns (uint224 tokenPrice, uint224 gasPriceValue) {
    if (!s_destChainConfigs[destChainSelector].isEnabled) revert DestinationChainNotEnabled(destChainSelector);
    return (_getValidatedTokenPrice(token), s_usdPerUnitGasByDestChainSelector[destChainSelector].value);
  }

  /// @notice Convert a given token amount to target token amount.
  /// @dev this function assumes that no more than 1e59 dollars are sent as payment.
  /// If more is sent, the multiplication of feeTokenAmount and feeTokenValue will overflow.
  /// Since there isn't even close to 1e59 dollars in the world economy this is safe.
  /// @param fromToken The given token address.
  /// @param fromTokenAmount The given token amount.
  /// @param toToken The target token address.
  /// @return toTokenAmount The target token amount.
  function convertTokenAmount(
    address fromToken,
    uint256 fromTokenAmount,
    address toToken
  ) public view returns (uint256) {
    /// Example:
    /// fromTokenAmount:   1e18      // 1 ETH
    /// ETH:               2_000e18
    /// LINK:              5e18
    /// return:            1e18 * 2_000e18 / 5e18 = 400e18 (400 LINK)
    return (fromTokenAmount * _getValidatedTokenPrice(fromToken)) / _getValidatedTokenPrice(toToken);
  }

  /// @notice Gets the token price for a given token and reverts if the token is not supported.
  /// @param token The address of the token to get the price for.
  /// @return tokenPriceValue The token price.
  function _getValidatedTokenPrice(
    address token
  ) internal view returns (uint224) {
    Internal.TimestampedPackedUint224 memory tokenPrice = getTokenPrice(token);
    // Token price must be set at least once.
    if (tokenPrice.timestamp == 0 || tokenPrice.value == 0) revert TokenNotSupported(token);
    return tokenPrice.value;
  }

  // ================================================================
  // │                         Fee tokens                           │
  // ================================================================

  /// @inheritdoc IFeeQuoter
  function getFeeTokens() external view returns (address[] memory) {
    return s_feeTokens.values();
  }

  /// @notice Add and remove tokens from feeTokens set.
  /// @param feeTokensToRemove The addresses of the tokens which are no longer considered feeTokens.
  /// @param feeTokensToAdd The addresses of the tokens which are now considered fee tokens and can be used
  /// to calculate fees.
  function applyFeeTokensUpdates(
    address[] memory feeTokensToRemove,
    address[] memory feeTokensToAdd
  ) external onlyOwner {
    _applyFeeTokensUpdates(feeTokensToRemove, feeTokensToAdd);
  }

  /// @notice Add and remove tokens from feeTokens set.
  /// @param feeTokensToRemove The addresses of the tokens which are no longer considered feeTokens.
  /// @param feeTokensToAdd The addresses of the tokens which are now considered fee tokens.
  /// and can be used to calculate fees.
  function _applyFeeTokensUpdates(address[] memory feeTokensToRemove, address[] memory feeTokensToAdd) private {
    for (uint256 i = 0; i < feeTokensToRemove.length; ++i) {
      if (s_feeTokens.remove(feeTokensToRemove[i])) {
        emit FeeTokenRemoved(feeTokensToRemove[i]);
      }
    }
    for (uint256 i = 0; i < feeTokensToAdd.length; ++i) {
      if (s_feeTokens.add(feeTokensToAdd[i])) {
        emit FeeTokenAdded(feeTokensToAdd[i]);
      }
    }
  }

  // ================================================================
  // │                       Price updates                          │
  // ================================================================

  /// @inheritdoc IFeeQuoter
  function updatePrices(
    Internal.PriceUpdates calldata priceUpdates
  ) external override {
    // The caller must be a fee updater.
    _validateCaller();

    uint256 tokenUpdatesLength = priceUpdates.tokenPriceUpdates.length;

    for (uint256 i = 0; i < tokenUpdatesLength; ++i) {
      Internal.TokenPriceUpdate memory update = priceUpdates.tokenPriceUpdates[i];
      s_usdPerToken[update.sourceToken] =
        Internal.TimestampedPackedUint224({value: update.usdPerToken, timestamp: uint32(block.timestamp)});
      emit UsdPerTokenUpdated(update.sourceToken, update.usdPerToken, block.timestamp);
    }

    uint256 gasUpdatesLength = priceUpdates.gasPriceUpdates.length;

    for (uint256 i = 0; i < gasUpdatesLength; ++i) {
      Internal.GasPriceUpdate memory update = priceUpdates.gasPriceUpdates[i];
      s_usdPerUnitGasByDestChainSelector[update.destChainSelector] =
        Internal.TimestampedPackedUint224({value: update.usdPerUnitGas, timestamp: uint32(block.timestamp)});
      emit UsdPerUnitGasUpdated(update.destChainSelector, update.usdPerUnitGas, block.timestamp);
    }
  }

  // ================================================================
  // │                       Fee quoting                            │
  // ================================================================

  /// @inheritdoc IFeeQuoter
  /// @dev The function should always validate message.extraArgs, message.receiver and family-specific configs.
  function getValidatedFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    DestChainConfig memory destChainConfig = s_destChainConfigs[destChainSelector];
    if (!destChainConfig.isEnabled) revert DestinationChainNotEnabled(destChainSelector);
    if (!s_feeTokens.contains(message.feeToken)) revert FeeTokenNotSupported(message.feeToken);

    uint256 gasLimit = _validateMessageAndResolveGasLimitForDestination(destChainSelector, destChainConfig, message);

    // The below call asserts that feeToken is a supported token.
    uint224 feeTokenPrice = _getValidatedTokenPrice(message.feeToken);

    // Calculate premiumFee in USD with 18 decimals precision first.
    // If message-only and no token transfers, a flat network fee is charged.
    // If there are token transfers, premiumFee is calculated from token transfer fee.
    // If there are both token transfers and message, premiumFee is only calculated from token transfer fee.
    uint256 premiumFeeUSDWei = 0;
    uint32 tokenTransferGas = 0;
    uint32 tokenTransferBytesOverhead = 0;
    if (message.tokenAmounts.length > 0) {
      (premiumFeeUSDWei, tokenTransferGas, tokenTransferBytesOverhead) = _getTokenTransferCost(
        destChainConfig.defaultTokenFeeUSDCents,
        destChainConfig.defaultTokenDestGasOverhead,
        destChainSelector,
        message.tokenAmounts
      );
    } else {
      // Convert USD cents with 2 decimals to 18 decimals.
      premiumFeeUSDWei = uint256(destChainConfig.networkFeeUSDCents) * 1e16;
    }
    // Apply the premium multiplier for the fee token, making it 36 decimals
    premiumFeeUSDWei *= s_premiumMultiplierWeiPerEth[message.feeToken];

    uint256 destCallDataCost =
      (message.data.length + tokenTransferBytesOverhead) * destChainConfig.destGasPerPayloadByteBase;

    // We add the destination chain CCIP overhead (commit, exec), the token transfer gas, the calldata cost and the msg
    // gas limit to get the total gas the tx costs to execute on the destination chain.
    uint256 totalDestChainGas = destChainConfig.destGasOverhead + tokenTransferGas + destCallDataCost + gasLimit;
    uint224 packedGasPrice = s_usdPerUnitGasByDestChainSelector[destChainSelector].value;

    // Total USD fee is in 36 decimals, feeTokenPrice is in 18 decimals USD for 1e18 smallest token denominations.
    // The result is the fee in the feeTokens smallest denominations (e.g. wei for ETH).
    // uint112(packedGasPrice) = executionGasPrice
    return (totalDestChainGas * uint112(packedGasPrice) * 1e18 + premiumFeeUSDWei) / feeTokenPrice;
  }

  /// @notice Sets the fee configuration for a token.
  /// @param premiumMultiplierWeiPerEthArgs Array of PremiumMultiplierWeiPerEthArgs structs.
  function applyPremiumMultiplierWeiPerEthUpdates(
    PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs
  ) external onlyOwner {
    _applyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs);
  }

  /// @dev Sets the fee config.
  /// @param premiumMultiplierWeiPerEthArgs The multiplier for destination chain specific premiums.
  function _applyPremiumMultiplierWeiPerEthUpdates(
    PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs
  ) internal {
    for (uint256 i = 0; i < premiumMultiplierWeiPerEthArgs.length; ++i) {
      address token = premiumMultiplierWeiPerEthArgs[i].token;
      uint64 premiumMultiplierWeiPerEth = premiumMultiplierWeiPerEthArgs[i].premiumMultiplierWeiPerEth;
      s_premiumMultiplierWeiPerEth[token] = premiumMultiplierWeiPerEth;

      emit PremiumMultiplierWeiPerEthUpdated(token, premiumMultiplierWeiPerEth);
    }
  }

  /// @notice Gets the fee configuration for a token.
  /// @param token The token to get the fee configuration for.
  /// @return premiumMultiplierWeiPerEth The multiplier for destination chain specific premiums.
  function getPremiumMultiplierWeiPerEth(
    address token
  ) external view returns (uint64 premiumMultiplierWeiPerEth) {
    return s_premiumMultiplierWeiPerEth[token];
  }

  /// @notice Returns the token transfer cost parameters.
  /// A basis point fee is calculated from the USD value of each token transfer.
  /// For each individual transfer, this fee is between [minFeeUSD, maxFeeUSD].
  /// Total transfer fee is the sum of each individual token transfer fee.
  /// @dev Assumes that tokenAmounts are validated to be listed tokens elsewhere.
  /// @dev Splitting one token transfer into multiple transfers is discouraged, as it will result in a transferFee
  /// equal or greater than the same amount aggregated/de-duped.
  /// @param defaultTokenFeeUSDCents the default token fee in USD cents.
  /// @param defaultTokenDestGasOverhead the default token destination gas overhead.
  /// @param destChainSelector the destination chain selector.
  /// @param tokenAmounts token transfers in the message.
  /// @return tokenTransferFeeUSDWei total token transfer bps fee in USD with 18 decimals.
  /// @return tokenTransferGas total execution gas of the token transfers.
  /// @return tokenTransferBytesOverhead additional token transfer data passed to destination, e.g. USDC attestation.
  function _getTokenTransferCost(
    uint256 defaultTokenFeeUSDCents,
    uint32 defaultTokenDestGasOverhead,
    uint64 destChainSelector,
    Client.EVMTokenAmount[] calldata tokenAmounts
  ) internal view returns (uint256 tokenTransferFeeUSDWei, uint32 tokenTransferGas, uint32 tokenTransferBytesOverhead) {
    uint256 numberOfTokens = tokenAmounts.length;

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      TokenTransferFeeConfig memory transferFeeConfig =
        s_tokenTransferFeeConfig[destChainSelector][tokenAmounts[i].token];

      // If the token has no specific overrides configured, we use the global defaults.
      if (!transferFeeConfig.isEnabled) {
        tokenTransferFeeUSDWei += defaultTokenFeeUSDCents * 1e16;
        tokenTransferGas += defaultTokenDestGasOverhead;
        tokenTransferBytesOverhead += Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES;
        continue;
      }

      tokenTransferGas += transferFeeConfig.destGasOverhead;
      tokenTransferBytesOverhead += transferFeeConfig.destBytesOverhead;

      // Convert USD values with 2 decimals to 18 decimals.
      tokenTransferFeeUSDWei += uint256(transferFeeConfig.feeUSDCents) * 1e16;
    }

    return (tokenTransferFeeUSDWei, tokenTransferGas, tokenTransferBytesOverhead);
  }

  /// @notice Gets the transfer fee config for a given token.
  /// @param destChainSelector The destination chain selector.
  /// @param token The token address.
  /// @return tokenTransferFeeConfig The transfer fee config for the token.
  function getTokenTransferFeeConfig(
    uint64 destChainSelector,
    address token
  ) external view returns (TokenTransferFeeConfig memory tokenTransferFeeConfig) {
    return s_tokenTransferFeeConfig[destChainSelector][token];
  }

  /// @notice Sets the transfer fee config.
  /// @dev only callable by the owner or admin.
  function applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    TokenTransferFeeConfigRemoveArgs[] memory tokensToUseDefaultFeeConfigs
  ) external onlyOwner {
    _applyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs);
  }

  /// @notice internal helper to set the token transfer fee config.
  function _applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    TokenTransferFeeConfigRemoveArgs[] memory tokensToUseDefaultFeeConfigs
  ) internal {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      TokenTransferFeeConfigArgs memory tokenTransferFeeConfigArg = tokenTransferFeeConfigArgs[i];
      uint64 destChainSelector = tokenTransferFeeConfigArg.destChainSelector;

      for (uint256 j = 0; j < tokenTransferFeeConfigArg.tokenTransferFeeConfigs.length; ++j) {
        TokenTransferFeeConfig memory tokenTransferFeeConfig =
          tokenTransferFeeConfigArg.tokenTransferFeeConfigs[j].tokenTransferFeeConfig;
        address token = tokenTransferFeeConfigArg.tokenTransferFeeConfigs[j].token;

        if (tokenTransferFeeConfig.destBytesOverhead < Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
          revert InvalidDestBytesOverhead(token, tokenTransferFeeConfig.destBytesOverhead);
        }

        s_tokenTransferFeeConfig[destChainSelector][token] = tokenTransferFeeConfig;

        emit TokenTransferFeeConfigUpdated(destChainSelector, token, tokenTransferFeeConfig);
      }
    }

    // Remove the custom fee configs for the tokens that are in the tokensToUseDefaultFeeConfigs array.
    for (uint256 i = 0; i < tokensToUseDefaultFeeConfigs.length; ++i) {
      uint64 destChainSelector = tokensToUseDefaultFeeConfigs[i].destChainSelector;
      address token = tokensToUseDefaultFeeConfigs[i].token;
      delete s_tokenTransferFeeConfig[destChainSelector][token];
      emit TokenTransferFeeConfigDeleted(destChainSelector, token);
    }
  }

  // ================================================================
  // │             Validations & message processing                 │
  // ================================================================

  /// @notice Validates that the destAddress matches the expected format of the family.
  /// @param chainFamilySelector Tag to identify the target family.
  /// @param destAddress Dest address to validate.
  /// @dev precondition - assumes the family tag is correct and validated.
  function _validateDestFamilyAddress(
    bytes4 chainFamilySelector,
    bytes memory destAddress,
    uint256 gasLimit
  ) internal pure {
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_EVM) {
      return Internal._validateEVMAddress(destAddress);
    }
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SVM) {
      // SVM addresses don't have a precompile space at the first X addresses, instead we validate that if the gasLimit
      // is non-zero, the address must not be 0x0.
      return Internal._validate32ByteAddress(destAddress, gasLimit > 0 ? 1 : 0);
    }
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_APTOS) {
      return Internal._validate32ByteAddress(destAddress, Internal.APTOS_PRECOMPILE_SPACE);
    }
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_TVM) {
      return Internal._validateTVMAddress(destAddress);
    }
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SUI) {
      return Internal._validate32ByteAddress(destAddress, gasLimit > 0 ? Internal.APTOS_PRECOMPILE_SPACE : 0);
    }
    revert InvalidChainFamilySelector(chainFamilySelector);
  }

  /// @notice Parse and validate the SVM specific Extra Args Bytes.
  function _parseSVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint256 maxPerMsgGasLimit
  ) internal pure returns (Client.SVMExtraArgsV1 memory svmExtraArgs) {
    if (extraArgs.length == 0) {
      revert InvalidExtraArgsData();
    }

    bytes4 tag = bytes4(extraArgs[:4]);
    if (tag != Client.SVM_EXTRA_ARGS_V1_TAG) {
      revert InvalidExtraArgsTag();
    }

    svmExtraArgs = abi.decode(extraArgs[4:], (Client.SVMExtraArgsV1));

    if (svmExtraArgs.computeUnits > maxPerMsgGasLimit) {
      revert MessageComputeUnitLimitTooHigh();
    }

    return svmExtraArgs;
  }

  /// @notice Parse and validate the Sui specific Extra Args Bytes.
  function _parseSuiExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint256 maxPerMsgGasLimit
  ) internal pure returns (Client.SuiExtraArgsV1 memory suiExtraArgs) {
    if (extraArgs.length == 0) {
      revert InvalidExtraArgsData();
    }

    bytes4 tag = bytes4(extraArgs[:4]);
    if (tag != Client.SUI_EXTRA_ARGS_V1_TAG) {
      revert InvalidExtraArgsTag();
    }

    suiExtraArgs = abi.decode(extraArgs[4:], (Client.SuiExtraArgsV1));

    if (suiExtraArgs.gasLimit > maxPerMsgGasLimit) {
      revert MessageGasLimitTooHigh();
    }

    return suiExtraArgs;
  }

  /// @dev Convert the extra args bytes into a struct with validations against the dest chain config.
  /// @param extraArgs The extra args bytes.
  /// @return genericExtraArgs The GenericExtraArgs struct.
  function _parseGenericExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint32 defaultTxGasLimit,
    uint256 maxPerMsgGasLimit
  ) internal pure returns (Client.GenericExtraArgsV2 memory) {
    // Since GenericExtraArgs are simply a superset of EVMExtraArgsV1, we can parse them as such. For Aptos, this
    // technically means EVMExtraArgsV1 are processed like they would be valid, but they will always fail on the
    // allowedOutOfOrderExecution check below.
    Client.GenericExtraArgsV2 memory parsedExtraArgs =
      _parseUnvalidatedEVMExtraArgsFromBytes(extraArgs, defaultTxGasLimit);

    if (parsedExtraArgs.gasLimit > maxPerMsgGasLimit) revert MessageGasLimitTooHigh();

    return parsedExtraArgs;
  }

  /// @dev Convert the extra args bytes into a struct.
  /// @param extraArgs The extra args bytes.
  /// @param defaultTxGasLimit default tx gas limit to use in the absence of extra args.
  /// @return EVMExtraArgsV2 the extra args struct populated with either the given args or default values.
  function _parseUnvalidatedEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint64 defaultTxGasLimit
  ) private pure returns (Client.GenericExtraArgsV2 memory) {
    if (extraArgs.length == 0) {
      // If extra args are empty, generate default values.
      return Client.GenericExtraArgsV2({gasLimit: defaultTxGasLimit, allowOutOfOrderExecution: false});
    }

    bytes4 extraArgsTag = bytes4(extraArgs);
    bytes memory argsData = extraArgs[4:];

    if (extraArgsTag == Client.GENERIC_EXTRA_ARGS_V2_TAG) {
      return abi.decode(argsData, (Client.GenericExtraArgsV2));
    } else if (extraArgsTag == Client.EVM_EXTRA_ARGS_V1_TAG) {
      // EVMExtraArgsV1 originally included a second boolean (strict) field which has been deprecated.
      // Clients may still include it but it will be ignored.
      return Client.GenericExtraArgsV2({gasLimit: abi.decode(argsData, (uint256)), allowOutOfOrderExecution: false});
    }
    revert InvalidExtraArgsTag();
  }

  /// @notice Validate the forwarded message to ensure it matches the configuration limits (message length, number of
  /// tokens) and family-specific expectations (address format).
  /// @param destChainSelector The destination chain selector.
  /// @param destChainConfig The destination chain config.
  /// @param message The message to validate.
  /// @return gasLimit The gas limit to use for the message.
  function _validateMessageAndResolveGasLimitForDestination(
    uint64 destChainSelector,
    DestChainConfig memory destChainConfig,
    Client.EVM2AnyMessage calldata message
  ) internal view returns (uint256 gasLimit) {
    uint256 dataLength = message.data.length;
    uint256 numberOfTokens = message.tokenAmounts.length;

    // Check that payload is formed correctly.
    if (dataLength > uint256(destChainConfig.maxDataBytes)) {
      revert MessageTooLarge(uint256(destChainConfig.maxDataBytes), dataLength);
    }
    if (numberOfTokens > 1) {
      revert UnsupportedNumberOfTokens(numberOfTokens, 1);
    }

    // resolve gas limit and validate chainFamilySelector
    if (
      destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_EVM
        || destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_APTOS
        || destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_TVM
    ) {
      gasLimit = _parseGenericExtraArgsFromBytes(
        message.extraArgs, destChainConfig.defaultTxGasLimit, destChainConfig.maxPerMsgGasLimit
      ).gasLimit;

      _validateDestFamilyAddress(destChainConfig.chainFamilySelector, message.receiver, gasLimit);
    } else if (destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SUI) {
      Client.SuiExtraArgsV1 memory suiExtraArgsV1 =
        _parseSuiExtraArgsFromBytes(message.extraArgs, destChainConfig.maxPerMsgGasLimit);

      gasLimit = suiExtraArgsV1.gasLimit;

      _validateDestFamilyAddress(destChainConfig.chainFamilySelector, message.receiver, gasLimit);

      uint256 receiverObjectIdsLength = suiExtraArgsV1.receiverObjectIds.length;
      // The max payload size for SUI is heavily dependent on the receiver object ids passed into extra args and the number of
      // tokens. Below, token and account overhead will count towards maxDataBytes.
      uint256 suiExpandedDataLength = dataLength;

      // This abi.decode is safe because the address is validated above.
      if (abi.decode(message.receiver, (uint256)) == 0) {
        // When message receiver is zero, CCIP receiver is not invoked on SUI.
        // There should not be additional accounts specified for the receiver.
        if (receiverObjectIdsLength > 0) {
          revert TooManySuiExtraArgsReceiverObjectIds(receiverObjectIdsLength, 0);
        }
      } else {
        // The messaging accounts needed for CCIP receiver on SUI are:
        // message receiver,
        // plus remaining accounts specified in Sui extraArgs. Each account is 32 bytes.
        suiExpandedDataLength +=
          ((receiverObjectIdsLength + Client.SUI_MESSAGING_ACCOUNTS_OVERHEAD) * Client.SUI_ACCOUNT_BYTE_SIZE);
      }

      if (numberOfTokens > 0 && suiExtraArgsV1.tokenReceiver == bytes32(0)) {
        revert InvalidTokenReceiver();
      }
      if (receiverObjectIdsLength > Client.SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS) {
        revert TooManySuiExtraArgsReceiverObjectIds(
          receiverObjectIdsLength, Client.SUI_EXTRA_ARGS_MAX_RECEIVER_OBJECT_IDS
        );
      }

      suiExpandedDataLength += (numberOfTokens * Client.SUI_TOKEN_TRANSFER_DATA_OVERHEAD);

      // The token destBytesOverhead can be very different per token so we have to take it into account as well.
      for (uint256 i = 0; i < numberOfTokens; ++i) {
        uint256 destBytesOverhead =
          s_tokenTransferFeeConfig[destChainSelector][message.tokenAmounts[i].token].destBytesOverhead;

        // Pools get Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES by default, but if an override is set we use that instead.
        if (destBytesOverhead > 0) {
          suiExpandedDataLength += destBytesOverhead;
        } else {
          suiExpandedDataLength += Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES;
        }
      }

      if (suiExpandedDataLength > uint256(destChainConfig.maxDataBytes)) {
        revert MessageTooLarge(uint256(destChainConfig.maxDataBytes), suiExpandedDataLength);
      }
    } else if (destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SVM) {
      Client.SVMExtraArgsV1 memory svmExtraArgsV1 =
        _parseSVMExtraArgsFromBytes(message.extraArgs, destChainConfig.maxPerMsgGasLimit);

      gasLimit = svmExtraArgsV1.computeUnits;

      _validateDestFamilyAddress(destChainConfig.chainFamilySelector, message.receiver, gasLimit);

      uint256 accountsLength = svmExtraArgsV1.accounts.length;
      // The max payload size for SVM is heavily dependent on the accounts passed into extra args and the number of
      // tokens. Below, token and account overhead will count towards maxDataBytes.
      uint256 svmExpandedDataLength = dataLength;

      // This abi.decode is safe because the address is validated above.
      if (abi.decode(message.receiver, (uint256)) == 0) {
        // When message receiver is zero, CCIP receiver is not invoked on SVM.
        // There should not be additional accounts specified for the receiver.
        if (accountsLength > 0) {
          revert TooManySVMExtraArgsAccounts(accountsLength, 0);
        }
      } else {
        // The messaging accounts needed for CCIP receiver on SVM are:
        // message receiver, offRamp PDA signer,
        // plus remaining accounts specified in SVM extraArgs. Each account is 32 bytes.
        svmExpandedDataLength +=
          ((accountsLength + Client.SVM_MESSAGING_ACCOUNTS_OVERHEAD) * Client.SVM_ACCOUNT_BYTE_SIZE);
      }

      if (numberOfTokens > 0 && svmExtraArgsV1.tokenReceiver == bytes32(0)) {
        revert InvalidTokenReceiver();
      }
      if (accountsLength > Client.SVM_EXTRA_ARGS_MAX_ACCOUNTS) {
        revert TooManySVMExtraArgsAccounts(accountsLength, Client.SVM_EXTRA_ARGS_MAX_ACCOUNTS);
      }
      if (svmExtraArgsV1.accountIsWritableBitmap >> accountsLength != 0) {
        revert InvalidSVMExtraArgsWritableBitmap(svmExtraArgsV1.accountIsWritableBitmap, accountsLength);
      }

      svmExpandedDataLength += (numberOfTokens * Client.SVM_TOKEN_TRANSFER_DATA_OVERHEAD);

      // The token destBytesOverhead can be very different per token so we have to take it into account as well.
      for (uint256 i = 0; i < numberOfTokens; ++i) {
        uint256 destBytesOverhead =
          s_tokenTransferFeeConfig[destChainSelector][message.tokenAmounts[i].token].destBytesOverhead;

        // Pools get Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES by default, but if an override is set we use that instead.
        if (destBytesOverhead > 0) {
          svmExpandedDataLength += destBytesOverhead;
        } else {
          svmExpandedDataLength += Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES;
        }
      }

      if (svmExpandedDataLength > uint256(destChainConfig.maxDataBytes)) {
        revert MessageTooLarge(uint256(destChainConfig.maxDataBytes), svmExpandedDataLength);
      }
    } else {
      revert InvalidChainFamilySelector(destChainConfig.chainFamilySelector);
    }

    return gasLimit;
  }

  /// @inheritdoc IFeeQuoter
  /// @dev precondition - onRampTokenTransfers and sourceTokenAmounts lengths must be equal.
  function processMessageArgs(
    uint64 destChainSelector,
    address feeToken,
    uint256 feeTokenAmount,
    bytes calldata extraArgs,
    bytes calldata messageReceiver
  )
    external
    view
    returns (
      uint256 msgFeeJuels,
      bool isOutOfOrderExecution,
      bytes memory convertedExtraArgs,
      bytes memory tokenReceiver
    )
  {
    // Convert feeToken to link if not already in link.
    if (feeToken == i_linkToken) {
      msgFeeJuels = feeTokenAmount;
    } else {
      msgFeeJuels = convertTokenAmount(feeToken, feeTokenAmount, i_linkToken);
    }

    if (msgFeeJuels > i_maxFeeJuelsPerMsg) revert MessageFeeTooHigh(msgFeeJuels, i_maxFeeJuelsPerMsg);

    (convertedExtraArgs, isOutOfOrderExecution, tokenReceiver) =
      _processChainFamilySelector(destChainSelector, messageReceiver, extraArgs);

    return (msgFeeJuels, isOutOfOrderExecution, convertedExtraArgs, tokenReceiver);
  }

  /// @notice Parses the extra Args based on the chain family selector. Isolated into a separate function
  /// as it was the only way to prevent a stack too deep error, and makes future chain family additions easier.
  // solhint-disable-next-line chainlink-solidity/explicit-returns
  function _processChainFamilySelector(
    uint64 destChainSelector,
    bytes calldata messageReceiver,
    bytes calldata extraArgs
  ) internal view returns (bytes memory validatedExtraArgs, bool allowOutOfOrderExecution, bytes memory tokenReceiver) {
    // Since this function is called after getFee, which already validates the params, no validation is necessary.
    DestChainConfig memory destChainConfig = s_destChainConfigs[destChainSelector];
    // EVM and Aptos both use the same GenericExtraArgs, with EVM also supporting EVMExtraArgsV1 which is handled inside
    // the generic function.
    if (
      destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_EVM
        || destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_APTOS
        || destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_TVM
    ) {
      Client.GenericExtraArgsV2 memory parsedExtraArgs =
        _parseUnvalidatedEVMExtraArgsFromBytes(extraArgs, destChainConfig.defaultTxGasLimit);

      return (Client._argsToBytes(parsedExtraArgs), parsedExtraArgs.allowOutOfOrderExecution, messageReceiver);
    }
    if (destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SUI) {
      return (extraArgs, true, messageReceiver);
    }
    if (destChainConfig.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SVM) {
      // If extraArgs passes the parsing it's valid and can be returned unchanged.
      // ExtraArgs are required on SVM, meaning the supplied extraArgs are either invalid and we would have reverted
      // or we have valid extraArgs and we can return them without having to re-encode them.
      return (
        extraArgs,
        true,
        abi.encode(_parseSVMExtraArgsFromBytes(extraArgs, destChainConfig.maxPerMsgGasLimit).tokenReceiver)
      );
    }
    revert InvalidChainFamilySelector(destChainConfig.chainFamilySelector);
  }

  /// @inheritdoc IFeeQuoter
  function processPoolReturnData(
    uint64 destChainSelector,
    Internal.EVM2AnyTokenTransfer[] calldata onRampTokenTransfers,
    Client.EVMTokenAmount[] calldata sourceTokenAmounts
  ) external view returns (bytes[] memory destExecDataPerToken) {
    bytes4 chainFamilySelector = s_destChainConfigs[destChainSelector].chainFamilySelector;
    destExecDataPerToken = new bytes[](onRampTokenTransfers.length);
    for (uint256 i = 0; i < onRampTokenTransfers.length; ++i) {
      address sourceToken = sourceTokenAmounts[i].token;

      // Since the DON has to pay for the extraData to be included on the destination chain, we cap the length of the
      // extraData. This prevents gas bomb attacks on the NOPs. As destBytesOverhead accounts for both.
      // extraData and offchainData, this caps the worst case abuse to the number of bytes reserved for offchainData.
      uint256 destPoolDataLength = onRampTokenTransfers[i].extraData.length;
      if (destPoolDataLength > Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
        if (destPoolDataLength > s_tokenTransferFeeConfig[destChainSelector][sourceToken].destBytesOverhead) {
          revert SourceTokenDataTooLarge(sourceToken);
        }
      }

      // We pass '1' here so that SVM validation requires a non-zero token address.
      // The 'gasLimit' parameter isn't actually used for gas in this context; it simply
      // signals that the address must not be zero on SVM.
      _validateDestFamilyAddress(chainFamilySelector, onRampTokenTransfers[i].destTokenAddress, 1);
      FeeQuoter.TokenTransferFeeConfig memory tokenTransferFeeConfig =
        s_tokenTransferFeeConfig[destChainSelector][sourceToken];

      uint32 destGasAmount = tokenTransferFeeConfig.isEnabled
        ? tokenTransferFeeConfig.destGasOverhead
        : s_destChainConfigs[destChainSelector].defaultTokenDestGasOverhead;

      // The user will be billed either the default or the override, so we send the exact amount that we billed for
      // to the destination chain to be used for the token releaseOrMint and transfer.
      destExecDataPerToken[i] = abi.encode(destGasAmount);
    }
    return destExecDataPerToken;
  }

  // ================================================================
  // │                           Configs                            │
  // ================================================================

  /// @notice Returns the configured config for the dest chain selector.
  /// @param destChainSelector Destination chain selector to fetch config for.
  /// @return destChainConfig Config for the destination chain.
  function getDestChainConfig(
    uint64 destChainSelector
  ) external view returns (DestChainConfig memory) {
    return s_destChainConfigs[destChainSelector];
  }

  /// @notice Updates the destination chain specific config.
  /// @param destChainConfigArgs Array of source chain specific configs.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] memory destChainConfigArgs
  ) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  /// @notice Internal version of applyDestChainConfigUpdates.
  function _applyDestChainConfigUpdates(
    DestChainConfigArgs[] memory destChainConfigArgs
  ) internal {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArgs[i].destChainSelector;
      DestChainConfig memory destChainConfig = destChainConfigArg.destChainConfig;

      // destChainSelector must be non-zero, defaultTxGasLimit must be set, must be less than maxPerMsgGasLimit
      if (
        destChainSelector == 0 || destChainConfig.defaultTxGasLimit == 0
          || destChainConfig.defaultTxGasLimit > destChainConfig.maxPerMsgGasLimit
          || (
            destChainConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_EVM
              && destChainConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_SVM
              && destChainConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_APTOS
              && destChainConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_SUI
              && destChainConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_TVM
          )
      ) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      // If the chain family selector is zero, it indicates that the chain was never configured and we
      // are adding a new chain.
      if (s_destChainConfigs[destChainSelector].chainFamilySelector == 0) {
        emit DestChainAdded(destChainSelector, destChainConfig);
      } else {
        emit DestChainConfigUpdated(destChainSelector, destChainConfig);
      }

      s_destChainConfigs[destChainSelector] = destChainConfig;
    }
  }

  /// @notice Returns the static FeeQuoter config.
  /// @dev RMN depends on this function, if updated, please notify the RMN maintainers.
  /// @return staticConfig The static configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({maxFeeJuelsPerMsg: i_maxFeeJuelsPerMsg, linkToken: i_linkToken});
  }
}
