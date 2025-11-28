// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";
import {IRMNRemote} from "../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "../../libraries/Client.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

abstract contract BaseVerifier is ICrossChainVerifierV1, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.AddressSet;
  using SafeERC20 for IERC20;

  error CursedByRMN(uint64 destChainSelector);
  error InvalidRemoteChainConfig(uint64 remoteChainSelector);
  error DestGasCannotBeZero(uint64 destChainSelector);
  error InvalidAllowListRequest(uint64 destChainSelector);
  error SenderNotAllowed(address sender);
  error CallerIsNotARampOnRouter(address caller);
  error NotSupported(uint64 remoteChainSelector);

  event FeeTokenWithdrawn(address indexed receiver, address indexed feeToken, uint256 amount);
  event RemoteChainConfigSet(uint64 indexed destChainSelector, address router, bool allowlistEnabled);
  event AllowListSendersAdded(uint64 indexed destChainSelector, address[] senders);
  event AllowListSendersRemoved(uint64 indexed destChainSelector, address[] senders);
  event StorageLocationUpdated(string oldLocation, string newLocation);

  struct RemoteChainConfig {
    IRouter router; // ──────────╮ Local router to use for messages going to this dest chain.
    uint16 feeUSDCents; //       │ The fee in US dollar cents for messages to this dest chain. [0, $655.35]
    uint32 gasForVerification; //│ The gas to reserve for verification of messages on the dest chain.
    uint32 payloadSizeBytes; //  │ The size of the verification payload on the dest chain.
    bool allowlistEnabled; // ───╯ True if the allowlist is enabled.
    EnumerableSet.AddressSet allowedSendersList; // The list of addresses allowed to send messages.
  }

  struct RemoteChainConfigArgs {
    IRouter router; // ────────────╮ Local router to use for messages going to this dest chain.
    uint64 remoteChainSelector; // │ Remote chain selector.
    bool allowlistEnabled; //      │ True if the allowlist is enabled.
    uint16 feeUSDCents; // ────────╯ The fee in US dollar cents for messages to this dest chain.
    uint32 gasForVerification; // ─╮ The gas to reserve for verification of messages on the dest chain.
    uint32 payloadSizeBytes; // ───╯ The size of the verification payload on the dest chain.
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

  /// @dev The remote chain specific configs.
  mapping(uint64 remoteChainSelector => RemoteChainConfig remoteChainConfig) private s_remoteChainConfigs;

  /// @dev The storage location for off-chain components to read from. Implementations of the BaseVerifier should
  /// implement a way to update this value if needed.
  string internal s_storageLocation;

  constructor(
    string memory storageLocation
  ) {
    _setStorageLocation(storageLocation);
  }

  /// @notice Updates the storage location.
  /// @param storageLocation The new storage location.
  function _setStorageLocation(
    string memory storageLocation
  ) internal {
    string memory oldLocation = s_storageLocation;
    s_storageLocation = storageLocation;
    emit StorageLocationUpdated(oldLocation, storageLocation);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function getStorageLocation() external view virtual override returns (string memory) {
    return s_storageLocation;
  }

  /// @notice get ChainConfig configured for the remoteChainSelector.
  /// @param remoteChainSelector The remote chain selector.
  /// @return allowlistEnabled boolean indicator to specify if allowlist check is enabled.
  /// @return router address of the local router.
  /// @return allowedSendersList list of addresses that are allowed to send messages to the remote chain.
  function getRemoteChainConfig(
    uint64 remoteChainSelector
  ) external view virtual returns (bool allowlistEnabled, address router, address[] memory allowedSendersList) {
    RemoteChainConfig storage config = _getRemoteChainConfig(remoteChainSelector);
    allowlistEnabled = config.allowlistEnabled;
    router = address(config.router);
    allowedSendersList = config.allowedSendersList.values();
    return (allowlistEnabled, router, allowedSendersList);
  }

  function _getRemoteChainConfig(
    uint64 remoteChainSelector
  ) internal view virtual returns (RemoteChainConfig storage) {
    return s_remoteChainConfigs[remoteChainSelector];
  }

  /// @notice Internal version of applyRemoteChainConfigUpdates.
  /// @dev the function that calls this has to ensure proper access control is in place.
  function _applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] memory remoteChainConfigArgs
  ) internal virtual {
    for (uint256 i = 0; i < remoteChainConfigArgs.length; ++i) {
      RemoteChainConfigArgs memory remoteChainConfigArg = remoteChainConfigArgs[i];
      uint64 remoteChainSelector = remoteChainConfigArgs[i].remoteChainSelector;

      if (remoteChainSelector == 0) {
        revert InvalidRemoteChainConfig(remoteChainSelector);
      }

      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[remoteChainSelector];
      // The router can be zero to pause the remote chain
      remoteChainConfig.router = remoteChainConfigArg.router;
      remoteChainConfig.allowlistEnabled = remoteChainConfigArg.allowlistEnabled;
      remoteChainConfig.feeUSDCents = remoteChainConfigArg.feeUSDCents;
      // The call can never cost 0 gas.
      if (remoteChainConfigArg.gasForVerification == 0) {
        revert DestGasCannotBeZero(remoteChainSelector);
      }
      remoteChainConfig.gasForVerification = remoteChainConfigArg.gasForVerification;
      // The payload could be zero bytes if no offchain data is required.
      remoteChainConfig.payloadSizeBytes = remoteChainConfigArg.payloadSizeBytes;

      emit RemoteChainConfigSet(
        remoteChainSelector, address(remoteChainConfigArg.router), remoteChainConfig.allowlistEnabled
      );
    }
  }

  function _assertSenderIsAllowed(uint64 destChainSelector, address sender) internal view virtual {
    RemoteChainConfig storage chainConfig = _getRemoteChainConfig(destChainSelector);
    if (address(chainConfig.router) == address(0)) {
      revert NotSupported(destChainSelector);
    }
    // CCVs should query the OnRamp address from the router, this allows for OnRamp updates without touching CCVs
    // OnRamp address may be zero intentionally to pause, which should stop all messages.
    if (msg.sender != chainConfig.router.getOnRamp(destChainSelector)) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }

    if (chainConfig.allowlistEnabled) {
      if (!chainConfig.allowedSendersList.contains(sender)) {
        revert SenderNotAllowed(sender);
      }
    }
  }

  function _onlyOffRamp(
    uint64 sourceChainSelector
  ) internal view virtual {
    IRouter router = _getRemoteChainConfig(sourceChainSelector).router;
    if (address(router) == address(0)) {
      revert NotSupported(sourceChainSelector);
    }
    // Check ensures that only a configured offRamp can call the function.
    if (!router.isOffRamp(sourceChainSelector, msg.sender)) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }
  }

  /// @notice Updates the allowlist for the destination chain.
  /// @param allowlistConfigArgsItems Array of AllowlistConfigArguments where each item is for a destChainSelector.
  function _applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) internal virtual {
    for (uint256 i = 0; i < allowlistConfigArgsItems.length; ++i) {
      AllowlistConfigArgs memory allowlistConfigArgs = allowlistConfigArgsItems[i];

      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[allowlistConfigArgs.destChainSelector];
      remoteChainConfig.allowlistEnabled = allowlistConfigArgs.allowlistEnabled;

      if (allowlistConfigArgs.addedAllowlistedSenders.length > 0) {
        if (allowlistConfigArgs.allowlistEnabled) {
          for (uint256 j = 0; j < allowlistConfigArgs.addedAllowlistedSenders.length; ++j) {
            address toAdd = allowlistConfigArgs.addedAllowlistedSenders[j];
            if (toAdd == address(0)) {
              revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
            }
            remoteChainConfig.allowedSendersList.add(toAdd);
          }

          emit AllowListSendersAdded(allowlistConfigArgs.destChainSelector, allowlistConfigArgs.addedAllowlistedSenders);
        } else {
          revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
        }
      }

      for (uint256 j = 0; j < allowlistConfigArgs.removedAllowlistedSenders.length; ++j) {
        remoteChainConfig.allowedSendersList.remove(allowlistConfigArgs.removedAllowlistedSenders[j]);
      }

      if (allowlistConfigArgs.removedAllowlistedSenders.length > 0) {
        emit AllowListSendersRemoved(
          allowlistConfigArgs.destChainSelector, allowlistConfigArgs.removedAllowlistedSenders
        );
      }
    }
  }

  /// @inheritdoc ICrossChainVerifierV1
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory, // message
    bytes memory, // extraArgs
    uint16 // blockConfirmations
  ) external view virtual returns (uint16 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) {
    if (s_remoteChainConfigs[destChainSelector].router == IRouter(address(0))) {
      revert NotSupported(destChainSelector);
    }
    return (
      s_remoteChainConfigs[destChainSelector].feeUSDCents,
      s_remoteChainConfigs[destChainSelector].gasForVerification,
      s_remoteChainConfigs[destChainSelector].payloadSizeBytes
    );
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @param feeAggregator The address to withdraw the fee tokens to.
  function _withdrawFeeTokens(address[] calldata feeTokens, address feeAggregator) internal virtual {
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
