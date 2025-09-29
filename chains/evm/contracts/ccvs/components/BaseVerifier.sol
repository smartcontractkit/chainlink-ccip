// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";
import {IRMNRemote} from "../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

abstract contract BaseVerifier is ICrossChainVerifierV1, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.AddressSet;
  using SafeERC20 for IERC20;

  error CursedByRMN(uint64 destChainSelector);
  error InvalidDestChainConfig(uint64 destChainSelector);
  error InvalidAllowListRequest(uint64 destChainSelector);
  error SenderNotAllowed(address sender);
  error CallerIsNotARampOnRouter(address caller);

  event FeeTokenWithdrawn(address indexed receiver, address indexed feeToken, uint256 amount);
  event DestChainConfigSet(uint64 indexed destChainSelector, address router, bool allowlistEnabled);
  event AllowListSendersAdded(uint64 indexed destChainSelector, address[] senders);
  event AllowListSendersRemoved(uint64 indexed destChainSelector, address[] senders);
  event StorageLocationUpdated(string oldLocation, string newLocation);

  struct DestChainConfig {
    bool allowlistEnabled; // ─╮ True if the allowlist is enabled.
    IRouter router; // ────────╯ Local router to use for messages going to this dest chain.
    EnumerableSet.AddressSet allowedSendersList; // The list of addresses allowed to send messages.
  }

  struct DestChainConfigArgs {
    IRouter router; // ──────────╮ Local router to use for messages going to this dest chain.
    uint64 destChainSelector; // │ Destination chain selector.
    bool allowlistEnabled; // ───╯ True if the allowlist is enabled.
  }

  /// @dev Struct to hold the allowlist configuration args per dest chain.
  struct AllowlistConfigArgs {
    uint64 destChainSelector; // ─╮ Destination chain selector.
    bool allowlistEnabled; // ────╯ True if the allowlist is enabled.
    address[] addedAllowlistedSenders; // list of senders to be added to the allowedSendersList.
    address[] removedAllowlistedSenders; // list of senders to be removed from the allowedSendersList.
  }

  /// @dev The rmn contract.
  IRMNRemote internal immutable i_rmnRemote;

  /// @dev The destination chain specific configs.
  mapping(uint64 destChainSelector => DestChainConfig destChainConfig) private s_destChainConfigs;

  /// @dev The storage location for off-chain components to read from. Implementations of the BaseVerifier should
  /// implement a way to update this value if needed.
  string internal s_storageLocation;

  constructor(
    string memory storageLocation
  ) {
    s_storageLocation = storageLocation;

    emit StorageLocationUpdated("", storageLocation);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function getStorageLocation() external view override returns (string memory) {
    return s_storageLocation;
  }

  /// @notice get ChainConfig configured for the DestinationChainSelector.
  /// @param destChainSelector The destination chain selector.
  /// @return allowlistEnabled boolean indicator to specify if allowlist check is enabled.
  /// @return router address of the local router.
  /// @return allowedSendersList list of addresses that are allowed to send messages to the destination chain.
  function getDestChainConfig(
    uint64 destChainSelector
  ) external view returns (bool allowlistEnabled, address router, address[] memory allowedSendersList) {
    DestChainConfig storage config = _getDestChainConfig(destChainSelector);
    allowlistEnabled = config.allowlistEnabled;
    router = address(config.router);
    allowedSendersList = config.allowedSendersList.values();
    return (allowlistEnabled, router, allowedSendersList);
  }

  function _getDestChainConfig(
    uint64 destChainSelector
  ) internal view returns (DestChainConfig storage) {
    return s_destChainConfigs[destChainSelector];
  }

  /// @notice Internal version of applyDestChainConfigUpdates.
  /// @dev the function that calls this has to ensure proper access control is in place.
  function _applyDestChainConfigUpdates(
    DestChainConfigArgs[] memory destChainConfigArgs
  ) internal {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArgs[i].destChainSelector;

      if (destChainSelector == 0) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
      // The router can be zero to pause the destination chain
      destChainConfig.router = destChainConfigArg.router;
      destChainConfig.allowlistEnabled = destChainConfigArg.allowlistEnabled;

      emit DestChainConfigSet(destChainSelector, address(destChainConfigArg.router), destChainConfig.allowlistEnabled);
    }
  }

  function _assertSenderIsAllowed(uint64 destChainSelector, address sender, address verifierCaller) internal view {
    DestChainConfig storage destChainConfig = _getDestChainConfig(destChainSelector);
    // CCVs should query the CCVProxy address from the router, this allows for CCVProxy updates without touching CCVs
    // CCVProxy address may be zero intentionally to pause, which should stop all messages.
    if (verifierCaller != destChainConfig.router.getOnRamp(destChainSelector)) {
      revert CallerIsNotARampOnRouter(verifierCaller);
    }

    if (destChainConfig.allowlistEnabled) {
      if (!destChainConfig.allowedSendersList.contains(sender)) {
        revert SenderNotAllowed(sender);
      }
    }
  }

  /// @notice Updates the allowlist for the destination chain.
  /// @param allowlistConfigArgsItems Array of AllowlistConfigArguments where each item is for a destChainSelector.
  function _applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) internal {
    for (uint256 i = 0; i < allowlistConfigArgsItems.length; ++i) {
      AllowlistConfigArgs memory allowlistConfigArgs = allowlistConfigArgsItems[i];

      DestChainConfig storage destChainConfig = s_destChainConfigs[allowlistConfigArgs.destChainSelector];
      destChainConfig.allowlistEnabled = allowlistConfigArgs.allowlistEnabled;

      if (allowlistConfigArgs.addedAllowlistedSenders.length > 0) {
        if (allowlistConfigArgs.allowlistEnabled) {
          for (uint256 j = 0; j < allowlistConfigArgs.addedAllowlistedSenders.length; ++j) {
            address toAdd = allowlistConfigArgs.addedAllowlistedSenders[j];
            if (toAdd == address(0)) {
              revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
            }
            destChainConfig.allowedSendersList.add(toAdd);
          }

          emit AllowListSendersAdded(allowlistConfigArgs.destChainSelector, allowlistConfigArgs.addedAllowlistedSenders);
        } else {
          revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
        }
      }

      for (uint256 j = 0; j < allowlistConfigArgs.removedAllowlistedSenders.length; ++j) {
        destChainConfig.allowedSendersList.remove(allowlistConfigArgs.removedAllowlistedSenders[j]);
      }

      if (allowlistConfigArgs.removedAllowlistedSenders.length > 0) {
        emit AllowListSendersRemoved(
          allowlistConfigArgs.destChainSelector, allowlistConfigArgs.removedAllowlistedSenders
        );
      }
    }
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @param feeAggregator The address to withdraw the fee tokens to.
  function _withdrawFeeTokens(address[] calldata feeTokens, address feeAggregator) internal {
    for (uint256 i = 0; i < feeTokens.length; ++i) {
      IERC20 feeToken = IERC20(feeTokens[i]);
      uint256 feeTokenBalance = feeToken.balanceOf(address(this));

      if (feeTokenBalance > 0) {
        feeToken.safeTransfer(feeAggregator, feeTokenBalance);

        emit FeeTokenWithdrawn(feeAggregator, address(feeToken), feeTokenBalance);
      }
    }
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) external pure virtual override returns (bool) {
    return interfaceId == type(ICrossChainVerifierV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
