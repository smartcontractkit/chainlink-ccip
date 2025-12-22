// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IBridgeV3} from "../interfaces/lombard/IBridgeV3.sol";
import {IMailbox} from "../interfaces/lombard/IMailbox.sol";

import {Internal} from "../libraries/Internal.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableMap} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableMap.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

contract LombardVerifier is BaseVerifier, Ownable2StepMsgSender {
  using EnumerableMap for EnumerableMap.AddressToAddressMap;
  using EnumerableSet for EnumerableSet.UintSet;
  using SafeERC20 for IERC20Metadata;

  error ZeroBridge();
  error ZeroLombardChainId();
  error PathNotExist(uint64 remoteChainSelector);
  error ExecutionError();
  error InvalidMessageLength(uint256 expected, uint256 actual);
  error InvalidMessageId(bytes32 messageMessageId, bytes32 bridgeMessageId);
  error InvalidReceiver(bytes);
  error InvalidMessageVersion(uint8 expected, uint8 actual);
  error InvalidCCVVersion(bytes4 expected, bytes4 actual);
  error TokenNotSupported(address token);
  error MustTransferTokens();

  /// @param remoteChainSelector CCIP selector of destination chain.
  /// @param lChainId The chain id of destination chain by Lombard Multi Chain Id conversion.
  /// @param allowedCaller The address of TokenPool on destination chain allowed to handle GMP message.
  event PathSet(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);
  /// @param remoteChainSelector CCIP selector of destination chain.
  /// @param lChainId The chain id of destination chain by Lombard Multi Chain Id conversion.
  /// @param allowedCaller The address that's allowed to call the bridge on the destination chain.
  event PathRemoved(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);
  event SupportedTokenRemoved(address token);
  event SupportedTokenSet(address localToken, address localAdapter);

  struct Path {
    /// @notice The address that's allowed to call the bridge on the destination chain.
    bytes32 allowedCaller;
    /// @notice Lombard chain id of destination chain.
    bytes32 lChainId;
  }

  struct SupportedTokenArgs {
    /// @notice The local token address.
    address localToken;
    /// @notice The local adapter address. Can be zero address if no adapter is used.
    address localAdapter;
  }

  string public constant typeAndVersion = "LombardVerifier 1.7.0-dev";
  /// @notice Version tag used in the verifier payload to indicate the version of this verifier.
  bytes4 private constant VERSION_TAG_V1_7_0 = bytes4(keccak256("LombardVerifier 1.7.0"));
  /// @notice The size of the version tag in bytes.
  uint256 private constant VERSION_TAG_SIZE = 4;
  /// @notice The size of a bytes32 in bytes.
  uint256 private constant BYTES32_SIZE = 32;
  /// @notice The expected size of the bridged message (version tag + message ID).
  uint256 private constant BRIDGED_MESSAGE_SIZE = VERSION_TAG_SIZE + BYTES32_SIZE;
  /// @notice The size of the rawPayload length field in ccvData.
  uint256 private constant RAW_PAYLOAD_LENGTH_SIZE = 2;
  uint256 private constant PAYLOAD_START_INDEX = VERSION_TAG_SIZE + RAW_PAYLOAD_LENGTH_SIZE;

  /// @notice Supported bridge message version.
  uint8 internal constant SUPPORTED_BRIDGE_MSG_VERSION = 1;
  /// @notice The address of bridge contract.
  IBridgeV3 public immutable i_bridge;

  /// @notice Set of supported tokens for cross-chain transfers. Even if an adapter is used, the source token must be
  /// added to the supported tokens set, not the adapter.
  EnumerableMap.AddressToAddressMap internal s_supportedTokens;
  /// @notice Set of supported chains for cross-chain transfers.
  EnumerableSet.UintSet internal s_supportedChains;
  /// @notice Mapping of CCIP chain selector to chain specific config.
  mapping(uint64 chainSelector => Path path) internal s_chainSelectorToPath;

  constructor(
    IBridgeV3 bridge,
    string[] memory storageLocation,
    address rmn
  ) BaseVerifier(storageLocation, rmn) {
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
    bytes32 messageId,
    address,
    uint256,
    bytes calldata
  ) external returns (bytes memory verifierData) {
    _assertNotCursedByRMN(message.destChainSelector);
    // We only support token transfers.
    if (message.tokenTransfer.length == 0) {
      revert MustTransferTokens();
    }

    // Sender must be an abi enceded EVM address.
    _assertSenderIsAllowed(message.destChainSelector, abi.decode(message.sender, (address)));
    return _callDepositOnBridge(message.tokenTransfer[0], message.destChainSelector, message.sender, messageId);
  }

  function _callDepositOnBridge(
    MessageV1Codec.TokenTransferV1 calldata tokenTransfer,
    uint64 destChainSelector,
    bytes calldata sender,
    bytes32 messageId
  ) internal returns (bytes memory) {
    // The Lombard bridge assumes addresses fit in 32 bytes and therefore only supports up to 32 byte addresses.
    if (tokenTransfer.tokenReceiver.length > 32) {
      revert InvalidReceiver(tokenTransfer.tokenReceiver);
    }

    // Check if the token is supported. This CCV will only support Lombard tokens.
    address sourceToken = abi.decode(tokenTransfer.sourceTokenAddress, (address));
    if (!s_supportedTokens.contains(sourceToken)) {
      revert TokenNotSupported(sourceToken);
    }

    Path memory path = s_chainSelectorToPath[destChainSelector];
    if (path.allowedCaller == bytes32(0)) {
      revert PathNotExist(destChainSelector);
    }

    // For some tokens we need to override the source token with an adapter.
    address localAdapter = s_supportedTokens.get(sourceToken);
    if (localAdapter != address(0)) {
      sourceToken = localAdapter;
    }

    (, bytes32 payloadHash) = i_bridge.deposit({
      destinationChain: path.lChainId,
      token: sourceToken,
      sender: abi.decode(sender, (address)),
      // Left pad receiver to 32 bytes if not already 32 bytes.
      recipient: Internal._leftPadBytesToBytes32(tokenTransfer.tokenReceiver),
      amount: tokenTransfer.amount,
      destinationCaller: path.allowedCaller,
      optionalMessage: bytes.concat(VERSION_TAG_V1_7_0, messageId)
    });

    // Return raw bytes instead of abi.encode for gas efficiency.
    return bytes.concat(payloadHash);
  }

  /// @inheritdoc ICrossChainVerifierV1
  /// @dev ccvData format:
  /// [versionTag (4 bytes)][rawPayloadLength (2 bytes)][rawPayload (variable)][proofLength (2 bytes)][proof (variable)]
  function verifyMessage(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    bytes calldata ccvData
  ) external {
    _assertNotCursedByRMN(message.sourceChainSelector);
    _onlyOffRamp(message.sourceChainSelector);

    bytes4 versionPrefix = bytes4(ccvData[:VERSION_TAG_SIZE]);
    if (versionPrefix != VERSION_TAG_V1_7_0) {
      revert InvalidCCVVersion(VERSION_TAG_V1_7_0, versionPrefix);
    }

    uint256 rawPayloadLength = uint16(bytes2(ccvData[VERSION_TAG_SIZE:PAYLOAD_START_INDEX]));
    bytes calldata rawPayload = ccvData[PAYLOAD_START_INDEX:PAYLOAD_START_INDEX + rawPayloadLength];

    uint256 proofDataStartIndex = PAYLOAD_START_INDEX + rawPayloadLength;

    uint256 proofLength = uint16(bytes2(ccvData[proofDataStartIndex:proofDataStartIndex + RAW_PAYLOAD_LENGTH_SIZE]));
    uint256 proofStartIndex = proofDataStartIndex + RAW_PAYLOAD_LENGTH_SIZE;
    bytes calldata proof = ccvData[proofStartIndex:proofStartIndex + proofLength];

    (, bool executed, bytes memory bridgedMessage) = IMailbox(i_bridge.mailbox()).deliverAndHandle(rawPayload, proof);
    if (!executed) {
      revert ExecutionError();
    }
    // The bridged message is expected to be the version tag and message id.
    if (bridgedMessage.length != BRIDGED_MESSAGE_SIZE) {
      revert InvalidMessageLength(BRIDGED_MESSAGE_SIZE, bridgedMessage.length);
    }
    bytes4 version;
    bytes32 returnedMessageId;
    assembly {
      // Load version from first 4 bytes.
      version := mload(add(bridgedMessage, 0x20))
      // Load messageId from bytes 4-36.
      returnedMessageId := mload(add(bridgedMessage, 0x24))
    }
    if (version != VERSION_TAG_V1_7_0) {
      revert InvalidCCVVersion(VERSION_TAG_V1_7_0, version);
    }
    if (returnedMessageId != messageId) {
      revert InvalidMessageId(messageId, returnedMessageId);
    }
  }

  /// @notice Gets the list of supported tokens for cross-chain transfers.
  function getSupportedTokens() external view returns (address[] memory) {
    return s_supportedTokens.keys();
  }

  /// @notice Checks if a token is supported for cross-chain transfers.
  /// @param token The token address to check.
  /// @return True if the token is supported, false otherwise.
  function isSupportedToken(
    address token
  ) external view returns (bool) {
    return s_supportedTokens.contains(token);
  }

  /// @notice Update the supported tokens for cross-chain transfers. When adding a token, it approves the bridge to
  /// spend an unlimited amount of the token. When removing a token, it resets the bridge's allowance to zero.
  /// @param tokensToRemove Array of token addresses to remove from supported tokens.
  /// @param tokensToSet Array of token addresses to set to supported tokens.
  function updateSupportedTokens(
    address[] calldata tokensToRemove,
    SupportedTokenArgs[] calldata tokensToSet
  ) external onlyOwner {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      address tokenToRemove = tokensToRemove[i];
      if (s_supportedTokens.remove(tokenToRemove)) {
        IERC20Metadata(tokenToRemove).safeApprove(address(i_bridge), 0);
        emit SupportedTokenRemoved(tokenToRemove);
      }
    }

    for (uint256 i = 0; i < tokensToSet.length; ++i) {
      SupportedTokenArgs memory tokenToAdd = tokensToSet[i];
      // No-op if the token is already supported.
      s_supportedTokens.set(tokenToAdd.localToken, tokenToAdd.localAdapter);

      address entityToApprove = tokenToAdd.localAdapter != address(0) ? tokenToAdd.localAdapter : tokenToAdd.localToken;

      // Either the token or the adapter needs to be approved for bridge spend. Cannot use safeApprove due to potential
      // existing non-zero allowance.
      IERC20Metadata(entityToApprove).approve(address(i_bridge), type(uint256).max);

      emit SupportedTokenSet(tokenToAdd.localToken, tokenToAdd.localAdapter);
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

  /// @notice Gets the path for a given CCIP chain selector.
  /// @param remoteChainSelector CCIP chain selector of remote chain.
  /// @return Path struct containing lChainId and allowedCaller.
  function getPath(
    uint64 remoteChainSelector
  ) external view returns (Path memory) {
    return s_chainSelectorToPath[remoteChainSelector];
  }

  /// @notice Sets the lChainId and allowed caller for a CCIP chain selector.
  /// @param remoteChainSelector CCIP chain selector of remote chain.
  /// @param lChainId Lombard chain id of remote chain.
  /// @param allowedCaller The address of LombardVerifier on destination chain.
  function setPath(
    uint64 remoteChainSelector,
    bytes32 lChainId,
    bytes32 allowedCaller
  ) external onlyOwner {
    if (lChainId == bytes32(0)) {
      revert ZeroLombardChainId();
    }

    s_chainSelectorToPath[remoteChainSelector] = Path({lChainId: lChainId, allowedCaller: allowedCaller});
    s_supportedChains.add(uint256(remoteChainSelector));

    emit PathSet(remoteChainSelector, lChainId, allowedCaller);
  }

  /// @notice Removes the path for the given CCIP chain selectors. This disables any traffic to those chains.
  /// @param remoteChainSelectors CCIP chain selectors of destination chains.
  function removePaths(
    uint64[] memory remoteChainSelectors
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = remoteChainSelectors[i];
      Path memory path = s_chainSelectorToPath[remoteChainSelector];

      if (!s_supportedChains.remove(uint256(remoteChainSelector))) {
        revert PathNotExist(remoteChainSelector);
      }

      delete s_chainSelectorToPath[remoteChainSelector];

      emit PathRemoved(remoteChainSelector, path.lChainId, path.allowedCaller);
    }
  }

  function applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] calldata remoteChainConfigArgs
  ) external onlyOwner {
    _applyRemoteChainConfigUpdates(remoteChainConfigArgs);
  }
}
