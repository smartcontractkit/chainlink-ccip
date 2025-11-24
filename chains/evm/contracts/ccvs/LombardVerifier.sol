// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {IBridgeV2} from "../interfaces/lombard/IBridgeV2.sol";
import {IMailbox} from "../interfaces/lombard/IMailbox.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableMap} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableMap.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

abstract contract LombardVerifier is BaseVerifier, Ownable2StepMsgSender {
  using EnumerableMap for EnumerableMap.AddressToAddressMap;
  using EnumerableSet for EnumerableSet.UintSet;
  using SafeERC20 for IERC20Metadata;

  error ZeroBridge();
  error ZeroLombardChainId();
  error PathNotExist(uint64 remoteChainSelector);
  error ExecutionError();
  error HashMismatch();
  error InvalidReceiver(bytes);
  error InvalidMessageVersion(uint8 expected, uint8 actual);
  error InvalidAllowedCaller(bytes);
  error ChainNotSupported();
  error TokenNotSupported(address token);
  error RemoteTokenMismatch(bytes32 bridge, bytes32 pool);
  error MustTransferTokens();

  /// @param remoteChainSelector CCIP selector of destination chain
  /// @param lChainId The chain id of destination chain by Lombard Multi Chain Id conversion
  /// @param allowedCaller The address of TokenPool on destination chain allowed to handle GMP message
  event PathSet(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);

  /// @param remoteChainSelector CCIP selector of destination chain
  /// @param lChainId The chain id of destination chain by Lombard Multi Chain Id conversion
  /// @param allowedCaller The address of TokenPool on destination chain allowed to handle GMP message
  event PathRemoved(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);
  event SupportedTokenRemoved(address token);
  event SupportedTokenAdded(address localToken, address localAdapter);

  struct Path {
    /// @notice The address that's allowed to call the bridge on the destination chain.
    bytes32 allowedCaller;
    /// @notice Lombard chain id of destination chain.
    bytes32 lChainId;
  }

  struct SupportedTokenArgs {
    address localToken;
    address localAdapter; // Can be zero address if no adapter is used.
  }

  string public constant typeAndVersion = "LombardVerifier 1.7.0-dev";

  /// @notice Supported bridge message version.
  uint8 internal constant SUPPORTED_BRIDGE_MSG_VERSION = 1;
  /// @notice The address of bridge contract.
  IBridgeV2 public immutable i_bridge;

  /// @notice The address of the TokenAdminRegistry contract used to get the remote token addresses through the pool.
  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;

  /// @notice Set of supported tokens for cross-chain transfers. Even if an adapter is used, the source token must be
  /// added to the supported tokens set, not the adapter.
  EnumerableMap.AddressToAddressMap internal s_supportedTokens;
  /// @notice Set of supported chains for cross-chain transfers.
  EnumerableSet.UintSet internal s_supportedChains;
  /// @notice Mapping of CCIP chain selector to chain specific config.
  mapping(uint64 chainSelector => Path path) internal s_chainSelectorToPath;

  constructor(IBridgeV2 bridge, string memory storageLocation) BaseVerifier(storageLocation) {
    if (address(bridge) == address(0)) {
      revert ZeroBridge();
    }
    uint8 bridgeMsgVersion = bridge.MSG_VERSION();
    if (bridgeMsgVersion != SUPPORTED_BRIDGE_MSG_VERSION) {
      revert InvalidMessageVersion(SUPPORTED_BRIDGE_MSG_VERSION, bridgeMsgVersion);
    }

    i_bridge = bridge;
  }

  /// @inheritdoc ICrossChainVerifierV1
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external returns (bytes memory verifierData) {
    // We only support token transfers.
    if (message.tokenTransfer.length == 0) {
      revert MustTransferTokens();
    }
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = message.tokenTransfer[0];

    // The Lombard bridge assumes addresses fit in 32 bytes and therefore only supports up to 32 byte addresses.
    if (tokenTransfer.tokenReceiver.length > 32) {
      revert InvalidReceiver(tokenTransfer.tokenReceiver);
    }

    // left pad receiver to 32 bytes if not already 32 bytes.
    bytes32 decodedReceiver =
      bytes32(uint256(bytes32(tokenTransfer.tokenReceiver)) >> (256 - tokenTransfer.tokenReceiver.length * 8));
    address decodedSender = address(bytes20(message.sender));

    // Check if the token is supported. This CCV will only support Lombard tokens. Whether the token is supported on
    // the destination chain is validated through the token pool.
    address sourceToken = address(bytes20(message.tokenTransfer[0].sourceTokenAddress));
    // Even is an adapter is used, the source token must be supported.
    if (!s_supportedTokens.contains(sourceToken)) {
      revert TokenNotSupported(sourceToken);
    }

    Path memory path = s_chainSelectorToPath[message.destChainSelector];
    if (path.allowedCaller == bytes32(0)) {
      revert PathNotExist(message.destChainSelector);
    }

    address localAdapter = s_supportedTokens.get(sourceToken);
    // For some tokens we need to override the source token with an adapter.
    if (localAdapter != address(0)) {
      sourceToken = localAdapter;
    }

    // Check the remote token on the pool. Since this CCV supports multiple tokens, this ensures we don't duplicate
    // config and keep the pool as source of truth of supported destination chains.
    // The pool remote token is the source of truth for CCIP. The bridge token does not have to match, but since CCIP
    // is only aware of the pool remote token, this does not cause any issues.
    bytes32 poolDestToken = abi.decode(
      IPoolV2(i_tokenAdminRegistry.getPool(sourceToken)).getRemoteToken(message.destChainSelector), (bytes32)
    );
    if (poolDestToken == bytes32(0)) {
      revert ChainNotSupported();
    }

    // TODO call a function that takes in the message ID.
    (, bytes32 payloadHash) = i_bridge.deposit({
      destinationChain: path.lChainId, // Lombard chain id, not CCIP chain selector.
      token: sourceToken, // Either the source token or the adapter token.
      sender: decodedSender, // Sender decoded to address
      recipient: decodedReceiver, // Receiver decoded to bytes32
      amount: tokenTransfer.amount,
      destinationCaller: path.allowedCaller
    });

    return abi.encode(payloadHash);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(MessageV1Codec.MessageV1 memory, bytes32, bytes calldata ccvData) external {
    (bytes memory rawPayload, bytes memory proof) = abi.decode(ccvData, (bytes, bytes));

    // TODO verify message id matches payload message id.

    (, bool executed,) = IMailbox(i_bridge.mailbox()).deliverAndHandle(rawPayload, proof);
    if (!executed) {
      revert ExecutionError();
    }
  }

  /// @notice Gets the list of supported tokens for cross-chain transfers.
  function getSupportedTokens() external view returns (address[] memory) {
    return s_supportedTokens.keys();
  }

  /// @notice Update the supported tokens for cross-chain transfers. When adding a token, it approves the bridge to
  /// spend an unlimited amount of the token. When removing a token, it resets the bridge's allowance to zero.
  /// @param tokensToRemove Array of token addresses to remove from supported tokens.
  /// @param tokensToAdd Array of token addresses to add to supported tokens.
  function updateSupportedTokens(
    address[] calldata tokensToRemove,
    SupportedTokenArgs[] calldata tokensToAdd
  ) external onlyOwner {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      address tokenToRemove = tokensToRemove[i];
      if (s_supportedTokens.remove(tokenToRemove)) {
        IERC20Metadata(tokenToRemove).safeApprove(address(i_bridge), 0);
        emit SupportedTokenRemoved(tokenToRemove);
      }
    }

    for (uint256 i = 0; i < tokensToAdd.length; ++i) {
      SupportedTokenArgs memory tokenToAdd = tokensToAdd[i];
      if (s_supportedTokens.set(tokenToAdd.localToken, tokenToAdd.localAdapter)) {
        address entityToApprove =
          tokenToAdd.localAdapter != address(0) ? tokenToAdd.localAdapter : tokenToAdd.localToken;

        // Either the token or the adapter needs to be approved for bridge spend.
        IERC20Metadata(entityToApprove).approve(address(i_bridge), type(uint256).max);

        emit SupportedTokenAdded(tokenToAdd.localToken, tokenToAdd.localAdapter);
      }
    }
  }

  /// @notice Returns the list of supported chains.
  /// @return Array of supported CCIP chain selectors.
  function getSupportedChains() external view returns (uint64[] memory) {
    uint256 length = s_supportedChains.length();
    uint64[] memory chains = new uint64[](length);
    for (uint256 i = 0; i < length; ++i) {
      chains[i] = uint64(s_supportedChains.at(i));
    }
    return chains;
  }

  /// @notice Sets the lChainId and allowed caller for a CCIP chain selector.
  /// @param remoteChainSelector CCIP chain selector of remote chain
  /// @param lChainId Lombard chain id of remote chain
  /// @param allowedCaller The address of LombardVerifier on destination chain
  function setPath(uint64 remoteChainSelector, bytes32 lChainId, bytes calldata allowedCaller) external onlyOwner {
    if (lChainId == bytes32(0)) {
      revert ZeroLombardChainId();
    }

    if (allowedCaller.length != 32) {
      revert InvalidAllowedCaller(allowedCaller);
    }
    bytes32 decodedAllowedCaller = abi.decode(allowedCaller, (bytes32));

    s_chainSelectorToPath[remoteChainSelector] = Path({lChainId: lChainId, allowedCaller: decodedAllowedCaller});

    // No-op if already exists.
    s_supportedChains.add(uint256(remoteChainSelector));

    emit PathSet(remoteChainSelector, lChainId, decodedAllowedCaller);
  }

  /// @notice remove path mapping
  /// @param remoteChainSelectors CCIP chain selectors of destination chains
  function removePaths(
    uint64[] memory remoteChainSelectors
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = remoteChainSelectors[i];
      Path memory path = s_chainSelectorToPath[remoteChainSelector];

      if (path.allowedCaller == bytes32(0)) {
        revert PathNotExist(remoteChainSelector);
      }

      delete s_chainSelectorToPath[remoteChainSelector];

      s_supportedChains.remove(uint256(remoteChainSelector));

      emit PathRemoved(remoteChainSelector, path.lChainId, path.allowedCaller);
    }
  }
}
