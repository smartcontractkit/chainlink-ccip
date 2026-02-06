// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../interfaces/ICrossChainVerifierV1.sol";
import {IRMNRemote} from "../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "../../libraries/Client.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

abstract contract BaseVerifier is ICrossChainVerifierV1, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.AddressSet;

  error CursedByRMN(uint64 destChainSelector);
  error InvalidRemoteChainConfig(uint64 remoteChainSelector);
  error DestGasCannotBeZero(uint64 destChainSelector);
  error InvalidAllowListRequest(uint64 destChainSelector);
  error SenderNotAllowed(address sender);
  error CallerIsNotARampOnRouter(address caller);
  error RemoteChainNotSupported(uint64 remoteChainSelector);
  error ZeroAddressNotAllowed();

  event RemoteChainConfigSet(uint64 indexed remoteChainSelector, address router, bool allowlistEnabled);
  event AllowListSendersAdded(uint64 indexed destChainSelector, address[] senders);
  event AllowListSendersRemoved(uint64 indexed destChainSelector, address[] senders);
  event AllowListStateChanged(uint64 indexed destChainSelector, bool allowlistEnabled);
  event StorageLocationsUpdated(string[] oldLocations, string[] newLocations);

  struct RemoteChainConfig {
    IRouter router; // ──────────╮ Local router to use for messages to/fom this chain.
    uint16 feeUSDCents; //       │ The fee in US dollar cents for messages to this remote chain. [0, $655.35]
    uint32 gasForVerification; //│ The gas to reserve for verification of messages on the remote chain.
    uint32 payloadSizeBytes; //  │ The size of the verification payload on the remote chain.
    bool allowlistEnabled; // ───╯ True if the allowlist is enabled.
    EnumerableSet.AddressSet allowedSendersList; // The list of addresses allowed to send messages.
  }

  struct RemoteChainConfigArgs {
    IRouter router; // ────────────╮ Local router to use for messages going to/from this chain.
    uint64 remoteChainSelector; // │ Remote chain selector.
    bool allowlistEnabled; //      │ True if the allowlist is enabled.
    uint16 feeUSDCents; // ────────╯ The fee in US dollar cents for messages to this remote chain.
    uint32 gasForVerification; // ─╮ The gas to reserve for verification of messages on the remote chain.
    uint32 payloadSizeBytes; // ───╯ The size of the verification payload on the remote chain.
  }

  /// @dev Struct to hold the allowlist configuration args per dest chain.
  struct AllowlistConfigArgs {
    uint64 destChainSelector; // ─╮ Destination chain selector.
    bool allowlistEnabled; // ────╯ True if the allowlist is enabled.
    address[] addedAllowlistedSenders; // list of senders to be added to the allowedSendersList.
    address[] removedAllowlistedSenders; // list of senders to be removed from the allowedSendersList.
  }

  /// @dev The rmn contract.
  IRMNRemote internal immutable i_rmn;

  /// @dev The remote chain specific configs.
  mapping(uint64 remoteChainSelector => RemoteChainConfig remoteChainConfig) private s_remoteChainConfigs;

  /// @dev The storage locations for off-chain components to read from. Implementations of the BaseVerifier should
  /// implement a way to update this value if needed.
  string[] internal s_storageLocations;

  constructor(
    string[] memory storageLocations,
    address rmnAddress
  ) {
    _setStorageLocations(storageLocations);

    if (rmnAddress == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    i_rmn = IRMNRemote(rmnAddress);
  }

  /// @notice Updates the storage locations.
  /// @param storageLocations The new storage locations.
  function _setStorageLocations(
    string[] memory storageLocations
  ) internal {
    uint256 oldLength = s_storageLocations.length;
    uint256 newLength = storageLocations.length;

    string[] memory oldLocations = getStorageLocations();

    // Clear existing array.
    for (uint256 i; i < oldLength; ++i) {
      s_storageLocations.pop();
    }

    // Add new elements into array.
    for (uint256 i; i < newLength; ++i) {
      s_storageLocations.push(storageLocations[i]);
    }

    emit StorageLocationsUpdated(oldLocations, storageLocations);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function getStorageLocations() public view virtual override returns (string[] memory) {
    return s_storageLocations;
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

  function _assertSenderIsAllowed(
    uint64 destChainSelector,
    address sender
  ) internal view virtual {
    RemoteChainConfig storage chainConfig = _getRemoteChainConfig(destChainSelector);
    if (address(chainConfig.router) == address(0)) {
      revert RemoteChainNotSupported(destChainSelector);
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
      revert RemoteChainNotSupported(sourceChainSelector);
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

      if (remoteChainConfig.allowlistEnabled != allowlistConfigArgs.allowlistEnabled) {
        remoteChainConfig.allowlistEnabled = allowlistConfigArgs.allowlistEnabled;

        emit AllowListStateChanged(allowlistConfigArgs.destChainSelector, allowlistConfigArgs.allowlistEnabled);
      }

      for (uint256 j = 0; j < allowlistConfigArgs.removedAllowlistedSenders.length; ++j) {
        remoteChainConfig.allowedSendersList.remove(allowlistConfigArgs.removedAllowlistedSenders[j]);
      }

      if (allowlistConfigArgs.removedAllowlistedSenders.length > 0) {
        emit AllowListSendersRemoved(
          allowlistConfigArgs.destChainSelector, allowlistConfigArgs.removedAllowlistedSenders
        );
      }

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
      revert RemoteChainNotSupported(destChainSelector);
    }
    return (
      s_remoteChainConfigs[destChainSelector].feeUSDCents,
      s_remoteChainConfigs[destChainSelector].gasForVerification,
      s_remoteChainConfigs[destChainSelector].payloadSizeBytes
    );
  }

  function _assertNotCursedByRMN(
    uint64 destChainSelector
  ) internal view virtual {
    if (i_rmn.isCursed(bytes16(uint128(destChainSelector)))) {
      revert CursedByRMN(destChainSelector);
    }
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) external pure virtual override returns (bool) {
    return interfaceId == type(ICrossChainVerifierV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
