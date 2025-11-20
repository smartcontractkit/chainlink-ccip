// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IMessageTransmitter} from "../pools/USDC/interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "../pools/USDC/interfaces/ITokenMessenger.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {CCTPMessageTransmitterProxy} from "../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice The CCTPVerifier creates USDC burn messages on source and delivers them on destination.
/// @dev This verifier is for CCTP V2 and is not backwards compatible with CCTP V1.
contract CCTPVerifier is Ownable2StepMsgSender, BaseVerifier {
  using SafeERC20 for IERC20;

  error CustomFinalityArraysMustBeSameLength();
  error CustomFinalitiesMustBeStrictlyIncreasing();
  error InvalidCCVData();
  error InvalidCCVVersion(bytes4 expected, bytes4 got);
  error InvalidDomainUpdate(DomainUpdate args);
  error InvalidMessageTransmitterOnProxy(address expected, address got);
  error InvalidMessageTransmitterVersion(uint32 expected, uint32 got);
  error InvalidReceiver(bytes receiver);
  error InvalidTokenMessengerVersion(uint32 expected, uint32 got);
  error InvalidMessageId(bytes32 expected, bytes32 got);
  error InvalidMessageSender(bytes32 expected, bytes32 got);
  error InvalidMessageVersion(uint32 expected, uint32 got);
  error InvalidToken(address token);
  error InvalidTokenTransferLength(uint256 length);
  error MaxFeeExceedsUint32(uint256 maxFee);
  error MissingCustomFinalities();
  error OnlyCallableByOwnerOrAllowlistAdmin();
  error ReceiveMessageCallFailed();
  error UnknownDomain(uint64 destChainSelector);
  error UnsupportedFinality(uint32 finality);
  error ZeroAddressNotAllowed();

  event DomainsSet(DomainUpdate[] domains);
  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event StaticConfigSet(
    address tokenMessenger, address messageTransmitterProxy, address usdcToken, uint32 localDomainIdentifier
  );
  event FinalityConfigSet(FinalityConfig finalityConfig);

  /// @notice The static configuration.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    address tokenMessenger; // The address of the token messenger.
    address messageTransmitterProxy; // The address of the message transmitter proxy.
    address usdcToken; // The address of the USDC token.
    uint32 localDomainIdentifier; // The local domain identifier.
  }

  /// @notice The arguments required to update a remote domain.
  // solhint-disable-next-line gas-struct-packing
  struct DomainUpdate {
    bytes32 allowedCallerOnDest; // Address allowed to call receiveMessage on the domain (i.e. the MessageTransmitterProxy).
    bytes32 allowedCallerOnSource; // Address allowed to call depositForBurn on the domain (i.e. the TokenMessengerProxy).
    bytes32 mintRecipient; // Address to mint USDC to on the destination chain.
    uint32 domainIdentifier; // Unique domain ID used across CCTP.
    uint64 destChainSelector; // The corresponding CCIP destination chain selector for the domain.
    bool enabled; // Whether or not the domain is enabled.
  }

  /// @notice Parameters for _depositForBurn (stack too deep measure).
  // solhint-disable-next-line gas-struct-packing
  struct DepositForBurnParams {
    uint256 amount; // The amount of USDC to deposit for burn.
    bytes32 receiver; // The receiver of the minted USDC on the destination chain.
    bytes32 messageId; // The message ID of the CCIP message.
    bytes32 allowedCaller; // The allowed caller of the message transmitter on the destination chain.
    uint32 finality; // The finality of the CCIP message.
    uint32 domainIdentifier; // The domain identifier of the destination chain.
  }

  /// @notice A domain is a CCTP-specific representation of a destination chain.
  /// @dev Zero is a valid domain identifier.
  struct Domain {
    bytes32 allowedCallerOnDest; // Address allowed to call receiveMessage on the domain (i.e. the MessageTransmitterProxy).
    bytes32 allowedCallerOnSource; // Address allowed to call depositForBurn on the domain (i.e. the TokenMessengerProxy).
    bytes32 mintRecipient; // Address to mint USDC to on the destination chain.
    uint32 domainIdentifier; // ─╮ Unique domain ID used across CCTP.
    bool enabled; // ────────────╯ Whether or not the domain is enabled.
  }

  /// @notice Configures finality handling for this chain.
  struct FinalityConfig {
    uint32 defaultCCTPFinalityThreshold; // ──╮ CCTP finality threshold applied when CCIP finality is 0.
    uint16 defaultCCTPFinalityBps; // ────────╯ Basis points charged for standard CCTP transfers on destination.
    uint32[] customCCIPFinalities; // CCIP finality thresholds for custom finalities (enforced to be in ascending order).
    uint32[] customCCTPFinalityThresholds; // Corresponding CCTP finality thresholds for custom CCIP finalities.
    uint16[] customCCTPFinalityBps; // Basis points charged for CCTP transfers using custom finalities on destination.
  }

  /// @notice Dynamic configuration for this chain.
  struct DynamicConfig {
    address feeAggregator; // ──╮ Address to which fees are withdrawn.
    address allowlistAdmin; // ─╯ Address permitted to update the allowlist (in addition to the owner).
  }

  string public constant override typeAndVersion = "CCTPVerifier 1.7.0-dev";
  /// @notice The preimage is bytes4(keccak256("CCTPVerifier 1.7.0")).
  bytes4 private constant VERSION_TAG_V1_7_0 = 0x8e1d1a9d;
  /// @notice CCTP contracts use the number 1 to represent V2, as 0 represents V1.
  uint32 private constant SUPPORTED_USDC_VERSION = 1;
  /// @notice The division factor for basis points. This also represents the maximum bps fee.
  uint16 private constant BPS_DIVIDER = 10_000;
  /// @notice The length of a CCTP message, including the message body + hook data expected by this verifier.
  /// @dev Message format.
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * sourceDomain               4          uint32     4
  ///     * destinationDomain          4          uint32     8
  ///     * nonce                      32         bytes32   12
  ///     * sender                     32         bytes32   44
  ///     * recipient                  32         bytes32   76
  ///     * destinationCaller          32         bytes32   108
  ///     * minFinalityThreshold       4          uint32    140
  ///     * finalityThresholdExecuted  4          uint32    144
  ///     * messageBody                dynamic    bytes     148
  /// @dev Message body format.
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * burnToken                  32         bytes32    4
  ///     * mintRecipient              32         bytes32    36
  ///     * amount                     32         uint256    68
  ///     * messageSender              32         bytes32    100
  ///     * maxFee                     32         uint256    132
  ///     * feeExecuted                32         uint256    164
  ///     * expirationBlock            32         uint256    196
  ///     * hookData                   dynamic    bytes      228
  /// @dev Hook data format.
  ///     * Field                      Bytes      Type       Index
  ///     * verifierVersion            4          bytes4     0
  ///     * messageId                  32         bytes32    4
  /// @dev Total bytes = (4 * 3) + (32 * 4) + (4 * 2) + 4 + (32 * 7) + 4 + 32 = 412.
  uint256 private constant CCTP_MESSAGE_LENGTH = 412;
  /// @notice The length of a signature provided by the CCTP attestation service.
  /// @dev 65-byte ECDSA signature: v (32) + r (32) + s (1).
  uint256 private constant SIGNATURE_LENGTH = 65;
  /// @notice The number of bytes allocated to encoding the verifier version in the hook data.
  uint256 private constant VERIFIER_VERSION_LENGTH = 4;

  /// @notice The USDC token contract.
  IERC20 private immutable i_usdcToken;
  /// @notice The message transmitter proxy, which is used on destination as a non-upgradeable caller of all CCTP messages.
  /// @dev Instead of calling receiveMessage directly, we use a proxy to enable upgrades to the verifier without invalidating in-flight messages.
  /// CCTP messages define an address permitted to call receiveMessage, which will always be the message transmitter proxy.
  CCTPMessageTransmitterProxy private immutable i_messageTransmitterProxy;
  /// @notice The token messenger, which is used on source to send USDC over CCTP.
  /// @dev The token messenger calls into the message transmitter after burning USDC and forming the app-specific message body.
  ITokenMessenger private immutable i_tokenMessenger; // TODO: Update to TokenMessengerProxy when available.
  /// @notice The local domain identifier, i.e. a CCTP-specific identifier for the chain to which this contract is deployed.
  uint32 private immutable i_localDomainIdentifier;

  /// @notice A mapping of CCIP chain selectors to CCTP domain configurations.
  mapping(uint64 remoteChainSelector => Domain cctpDomain) private s_chainToDomain;
  /// @notice The dynamic configuration.
  DynamicConfig private s_dynamicConfig;
  /// @notice The finality configuration.
  FinalityConfig private s_finalityConfig;

  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy messageTransmitterProxy,
    IERC20 usdcToken,
    string memory storageLocation,
    FinalityConfig memory finalityConfig,
    DynamicConfig memory dynamicConfig
  ) BaseVerifier(storageLocation) {
    if (
      address(tokenMessenger) == address(0) || address(messageTransmitterProxy) == address(0)
        || address(usdcToken) == address(0)
    ) revert ZeroAddressNotAllowed();

    // Ensure that the token messenger is for CCTP.
    uint32 tokenMessengerVersion = tokenMessenger.messageBodyVersion();
    if (tokenMessengerVersion != SUPPORTED_USDC_VERSION) {
      revert InvalidTokenMessengerVersion(SUPPORTED_USDC_VERSION, tokenMessengerVersion);
    }

    // Ensure that the message transmitter is for CCTP.
    IMessageTransmitter messageTransmitter = IMessageTransmitter(tokenMessenger.localMessageTransmitter());
    uint32 messageTransmitterVersion = messageTransmitter.version();
    if (messageTransmitterVersion != SUPPORTED_USDC_VERSION) {
      revert InvalidMessageTransmitterVersion(SUPPORTED_USDC_VERSION, messageTransmitterVersion);
    }

    // Ensure that the message transmitter on the proxy is the same as the message transmitter on the token messenger.
    address messageTransmitterOnProxy = address(messageTransmitterProxy.i_cctpTransmitter());
    if (messageTransmitterOnProxy != address(messageTransmitter)) {
      revert InvalidMessageTransmitterOnProxy(address(messageTransmitter), messageTransmitterOnProxy);
    }

    // Set the immutable state variables.
    i_tokenMessenger = tokenMessenger;
    i_messageTransmitterProxy = messageTransmitterProxy;
    i_localDomainIdentifier = messageTransmitter.localDomain();
    i_usdcToken = usdcToken;

    // Approve the token messenger to burn the USDC token on behalf of this contract.
    // The USDC token pool will be responsible for forwarding USDC it receives from the router to this contract.
    i_usdcToken.safeIncreaseAllowance(address(i_tokenMessenger), type(uint256).max);

    emit StaticConfigSet(
      address(i_tokenMessenger), address(i_messageTransmitterProxy), address(i_usdcToken), i_localDomainIdentifier
    );

    _setDynamicConfig(dynamicConfig);
    _setFinalityConfig(finalityConfig);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    address, // feeToken
    uint256, // feeTokenAmount
    bytes calldata // verifierArgs
  ) external returns (bytes memory verifierReturnData) {
    // For EVM, sender is expected to be 20 bytes.
    address senderAddress = address(bytes20(message.sender));
    _assertSenderIsAllowed(message.destChainSelector, senderAddress);

    Domain storage domain = s_chainToDomain[message.destChainSelector];
    if (!domain.enabled) revert UnknownDomain(message.destChainSelector);

    // We expect exactly one token transfer per message.
    // The address of the token transferred must correspond to USDC.
    if (message.tokenTransfer.length != 1) revert InvalidTokenTransferLength(message.tokenTransfer.length);
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = message.tokenTransfer[0];

    address sourceTokenAddress = address(bytes20(tokenTransfer.sourceTokenAddress));
    if (sourceTokenAddress != address(i_usdcToken)) revert InvalidToken(sourceTokenAddress);

    if (message.tokenTransfer[0].tokenReceiver.length != 32) {
      revert InvalidReceiver(message.tokenTransfer[0].tokenReceiver);
    }

    bytes32 decodedReceiver;
    // For EVM chains, the mintRecipient is not used.
    // Solana requires it, as the mintRecipient will be a PDA owned by the pool.
    // The PDA will forward the tokens to their final destination after minting.
    if (domain.mintRecipient != bytes32(0)) {
      decodedReceiver = domain.mintRecipient;
    } else {
      decodedReceiver = abi.decode(tokenTransfer.tokenReceiver, (bytes32));
    }

    _depositForBurn(
      DepositForBurnParams({
        amount: tokenTransfer.amount,
        receiver: decodedReceiver,
        finality: message.finality,
        messageId: messageId,
        domainIdentifier: domain.domainIdentifier,
        allowedCaller: domain.allowedCallerOnDest
      })
    );

    return "";
  }

  /// @notice Deposits USDC tokens for burn into CCTP.
  /// @param params The parameters for the deposit.
  function _depositForBurn(
    DepositForBurnParams memory params
  ) private {
    (uint32 finalityThreshold, uint16 bps, bool found) = _getCCTPFinalityThresholdAndBps(params.finality);
    if (!found) revert UnsupportedFinality(params.finality);

    // The maximum fee, taken on destination, is a percentage of the total amount transferred.
    // We use bps to calculate the smallest possible value that we can set as the max fee.
    // The bps values configured for each finality threshold on this chain must mirror those used by CCTP.
    // CCTP defines different bps values for each chain.
    uint256 maxFee = params.amount * bps / BPS_DIVIDER;
    if (maxFee > type(uint32).max) revert MaxFeeExceedsUint32(maxFee);

    i_tokenMessenger.depositForBurnWithHook(
      params.amount,
      params.domainIdentifier,
      params.receiver,
      address(i_usdcToken),
      params.allowedCaller,
      uint32(maxFee),
      finalityThreshold,
      // The hook data includes the version tag and the message ID.
      // The version tag allows the destination verifier entity to route the message to the correct implementation.
      // Inclusion of the message ID ensures that the contents of the CCIP message can't be tampered with on destination.
      bytes.concat(VERSION_TAG_V1_7_0, params.messageId)
    );
  }

  /// @notice Returns the CCTP finality threshold and bps for the given CCIP finality.
  /// @param finality The CCIP finality.
  /// @return finalityThreshold The CCTP finality threshold.
  /// @return bps The bps charged on destinaton by CCTP for the finality threshold.
  /// @return found Whether the finality threshold and bps were found. We can't rely on 0 values because 0 could be valid.
  function _getCCTPFinalityThresholdAndBps(
    uint32 finality
  ) private view returns (uint32 finalityThreshold, uint16 bps, bool found) {
    if (finality == 0) {
      // Apply standard CCTP finality when CCIP finality is set to the default value of 0.
      return (s_finalityConfig.defaultCCTPFinalityThreshold, s_finalityConfig.defaultCCTPFinalityBps, true);
    } else {
      uint32[] memory customCCIPFinalities = s_finalityConfig.customCCIPFinalities;
      for (uint256 i = 0; i < customCCIPFinalities.length; ++i) {
        if (i == customCCIPFinalities.length - 1 || finality < customCCIPFinalities[i]) {
          // If we've reached the last custom finality available, we must use it no matter what.
          // If we've reached a finality that is greater than the requested finality, we will round up to it.
          // This mirrors the behavior of CCTP finality thresholds, which round up if the requested finality exceeds a threshold.
          return (s_finalityConfig.customCCTPFinalityThresholds[i], s_finalityConfig.customCCTPFinalityBps[i], true);
        }
      }
    }

    return (0, 0, false);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(MessageV1Codec.MessageV1 memory message, bytes32 messageHash, bytes calldata ccvData) external {
    // Message transmitter requires a minimum signature threshold of 1.
    // Our ccvData should therefore define at least 1 signature.
    if (ccvData.length < VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH + SIGNATURE_LENGTH) revert InvalidCCVData();

    bytes4 versionPrefix = bytes4(ccvData[:VERIFIER_VERSION_LENGTH]);
    if (versionPrefix != VERSION_TAG_V1_7_0) revert InvalidCCVVersion(VERSION_TAG_V1_7_0, versionPrefix);

    // The attested version is the first 4 bytes of the hook data, which occupies the last 36 bytes of the CCTP message.
    // We exclude the last 32 bytes of the hook data, which contains the message ID, to get the version.
    bytes4 attestedVersion = bytes4(
      ccvData[VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH - 36:VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH - 32]
    );
    if (attestedVersion != VERSION_TAG_V1_7_0) revert InvalidCCVVersion(VERSION_TAG_V1_7_0, attestedVersion);

    // The attested message ID should match the hash passed into this function.
    // If not, there is a mismatch between what was attested and what was computed within this transaction.
    bytes32 messageId =
      bytes32(ccvData[VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH - 32:VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH]);
    if (messageHash != messageId) revert InvalidMessageId(messageHash, messageId);

    // The messageSender property of the messageBody must align with the allowedCallerOnSource.
    // This check is critical to ensure that CCIP is unable to process burn messages generated by other systems.
    Domain storage sourceDomain = s_chainToDomain[message.sourceChainSelector];
    if (!sourceDomain.enabled) revert UnknownDomain(message.sourceChainSelector);
    // messageSender starts 148 (messageBody index) + 100 (index within messageBody) = 248 bytes after the verifier version tag.
    bytes32 messageSender = bytes32(ccvData[VERIFIER_VERSION_LENGTH + 248:VERIFIER_VERSION_LENGTH + 280]);
    if (messageSender != sourceDomain.allowedCallerOnSource) {
      revert InvalidMessageSender(sourceDomain.allowedCallerOnSource, messageSender);
    }

    // Call into CCTP via the message transmitter proxy.
    // CCTP will validate signatures against the message before minting USDC.
    // Attestation occupies all bytes following the CCTP message.
    if (
      !i_messageTransmitterProxy.receiveMessage(
        ccvData[VERIFIER_VERSION_LENGTH:VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH],
        ccvData[VERIFIER_VERSION_LENGTH + CCTP_MESSAGE_LENGTH:]
      )
    ) {
      revert ReceiveMessageCallFailed();
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static configuration.
  /// @return staticConfig The static configuration.
  function getStaticConfig() external view returns (StaticConfig memory staticConfig) {
    return StaticConfig({
      tokenMessenger: address(i_tokenMessenger),
      messageTransmitterProxy: address(i_messageTransmitterProxy),
      usdcToken: address(i_usdcToken),
      localDomainIdentifier: i_localDomainIdentifier
    });
  }

  /// @notice Returns the dynamic configuration.
  /// @return dynamicConfig The dynamic configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The dynamic configuration.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The dynamic configuration.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) private {
    if (dynamicConfig.feeAggregator == address(0)) revert ZeroAddressNotAllowed();

    s_dynamicConfig = dynamicConfig;

    emit DynamicConfigSet(dynamicConfig);
  }

  /// @notice Returns the finality configuration for this chain.
  /// @return finalityConfig The finality configuration.
  function getFinalityConfig() external view returns (FinalityConfig memory finalityConfig) {
    return s_finalityConfig;
  }

  /// @notice Sets the finality configuration for this chain.
  /// @dev This function validates that custom finality values are sorted in ascending order.
  /// @param finalityConfig The finality configuration.
  function setFinalityConfig(
    FinalityConfig memory finalityConfig
  ) external onlyOwner {
    _setFinalityConfig(finalityConfig);
  }

  /// @notice Sets the finality configuration for this chain.
  /// @dev This function validates that custom finality values are sorted in ascending order.
  /// @param finalityConfig The finality configuration.
  function _setFinalityConfig(
    FinalityConfig memory finalityConfig
  ) private {
    // We require at least one custom finality, otherwise CCTP fast finality won't be supported.
    if (finalityConfig.customCCIPFinalities.length == 0) revert MissingCustomFinalities();

    // All arrays must be the same length.
    if (
      finalityConfig.customCCIPFinalities.length != finalityConfig.customCCTPFinalityThresholds.length
        || finalityConfig.customCCIPFinalities.length != finalityConfig.customCCTPFinalityBps.length
    ) {
      revert CustomFinalityArraysMustBeSameLength();
    }

    // Validates that custom CCIP finalities are in ascending order.
    // We can initialize previous finality to 0, as 0 is the default finality value.
    // It always maps to standard CCTP transfer and should never be listed as a custom finality.
    uint32 previousFinality = 0;
    for (uint256 i = 0; i < finalityConfig.customCCIPFinalities.length; ++i) {
      if (finalityConfig.customCCIPFinalities[i] <= previousFinality) {
        revert CustomFinalitiesMustBeStrictlyIncreasing();
      }
      previousFinality = finalityConfig.customCCIPFinalities[i];
    }

    s_finalityConfig = finalityConfig;
    emit FinalityConfigSet(finalityConfig);
  }

  /// @notice Gets the CCTP domain for a given CCIP chain selector.
  /// @param chainSelector The CCIP chain selector corresponding to the domain.
  /// @return domain The CCTP domain corresponding to the given chain selector.
  function getDomain(
    uint64 chainSelector
  ) external view returns (Domain memory) {
    return s_chainToDomain[chainSelector];
  }

  /// @notice Sets the CCTP domain for a CCIP chain selector.
  /// @param domains The array of DomainUpdate structs to set.
  /// @dev Must validate mapping of selectors -> (domain, caller) prior to calling this function.
  function setDomains(
    DomainUpdate[] calldata domains
  ) external onlyOwner {
    for (uint256 i = 0; i < domains.length; ++i) {
      DomainUpdate memory domain = domains[i];
      if (
        domain.allowedCallerOnDest == bytes32(0) || domain.allowedCallerOnSource == bytes32(0)
          || domain.destChainSelector == 0
      ) {
        revert InvalidDomainUpdate(domain);
      }

      s_chainToDomain[domain.destChainSelector] = Domain({
        allowedCallerOnDest: domain.allowedCallerOnDest,
        allowedCallerOnSource: domain.allowedCallerOnSource,
        mintRecipient: domain.mintRecipient,
        domainIdentifier: domain.domainIdentifier,
        enabled: domain.enabled
      });
    }

    emit DomainsSet(domains);
  }

  /// @notice Updates destination chain configurations.
  /// @param destChainConfigArgs Array of destination chain configurations.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] calldata destChainConfigArgs
  ) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  /// @notice Updates senders that are allowed to use this verifier.
  /// @param allowlistConfigArgsItems Array of AllowListConfigArgs, where each item is for a destChainSelector.
  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external {
    if (msg.sender != owner() && msg.sender != s_dynamicConfig.allowlistAdmin) {
      revert OnlyCallableByOwnerOrAllowlistAdmin();
    }

    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  /// @notice Updates the storage location identifier.
  /// @param newLocation The new storage location identifier.
  function updateStorageLocation(
    string memory newLocation
  ) external onlyOwner {
    _setStorageLocation(newLocation);
  }

  /// @notice Exposes the version tag.
  function versionTag() public pure returns (bytes4) {
    return VERSION_TAG_V1_7_0;
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @dev This function can be permissionless as just transfers tokens to a trusted address.
  /// @param feeTokens The fee tokens to withdraw.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
    _withdrawFeeTokens(feeTokens, s_dynamicConfig.feeAggregator);
  }
}
