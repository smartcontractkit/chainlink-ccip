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

/// @notice The CCTPV2Verifier creates CCTP V2 messages on source and verifies them on destination.
/// @dev This verifier is not backwards compatible with CCTP V1.
contract CCTPV2Verifier is Ownable2StepMsgSender, BaseVerifier {
  using SafeERC20 for IERC20;

  error InvalidMessageTransmitterOnProxy(address expected, address got);
  error InvalidMessageTransmitterVersion(uint32 expected, uint32 got);
  error InvalidReceiver(bytes receiver);
  error InvalidTokenMessengerVersion(uint32 expected, uint32 got);
  error InvalidToken(address token);
  error InvalidTokenTransferLength(uint256 length);
  error OnlyCallableByOwnerOrAllowlistAdmin();
  error UnknownDomain(uint64 destChainSelector);
  error UnsupportedFinality(uint32 finality);
  error ZeroAddressNotAllowed();

  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event StaticConfigSet(
    address tokenMessenger, address messageTransmitterProxy, address usdcToken, uint32 localDomainIdentifier
  );

  /// @notice The arguments required to update a remote domain.
  // solhint-disable-next-line gas-struct-packing
  struct DomainUpdateArgs {
    bytes32 allowedCaller; // Address allowed to mint on the domain (i.e. the MessageTransmitterProxy on destination).
    bytes32 mintRecipient; // Address to mint USDC to on the destination chain.
    uint32 domainIdentifier; // Unique domain ID used across CCTP.
    uint64 destChainSelector; // The corresponding CCIP destination chain selector for the domain.
    bool enabled; // Whether or not the domain is enabled.
  }

  /// @notice Parameters for _depositForBurn.
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
    bytes32 allowedCaller; // Address allowed to mint on the domain (i.e. the MessageTransmitterProxy on destination).
    bytes32 mintRecipient; // Address to mint USDC to on the destination chain.
    uint32 domainIdentifier; // ─╮ Unique domain ID used across CCTP.
    bool enabled; // ────────────╯ Whether or not the domain is enabled.
  }

  /// @notice A custom finality maps any non-zero CCIP finality value into CCTP.
  struct CustomFinality {
    uint16 finality; // ─────────────╮ CCIP finality value.
    uint16 cctpFinalityThreshold; // | Corresponding CCTP finality threshold.
    uint16 cctpFinalityBps; // ──────╯ Basis points charged for the custom finality on destination.
  }

  /// @notice Configures finality handling for this chain.
  struct FinalityConfig {
    uint16 standardFinalityThreshold; // ──╮ CCTP finality threshold applied when CCIP finality is 0.
    uint16 standardFinalityBps; // ────────╯ Basis points charged for standard finality on destination.
    CustomFinality[] customFinalities; // Custom finality configurations for non-zero CCIP finality values.
  }

  /// @notice Dynamic configuration for this chain.
  struct DynamicConfig {
    address feeAggregator; // ──╮ Address to which fees are withdrawn.
    address allowlistAdmin; // ─╯ Address permitted to update the allowlist in addition to the owner.
  }

  string public constant override typeAndVersion = "CCTPV2Verifier 1.7.0-dev";
  /// @notice The preimage is bytes4(keccak256("CCTPV2Verifier 1.7.0")).
  bytes4 internal constant VERSION_TAG_V1_7_0 = 0xb4161002;
  /// @notice CCTP contracts use the number 1 to represent V2, as 0 represents V1.
  uint32 private constant SUPPORTED_USDC_VERSION = 1;
  /// @notice The division factor for basis points. This also represents the maximum bps fee.
  uint16 internal constant BPS_DIVIDER = 10_000;

  /// @notice The USDC token contract.
  IERC20 private immutable i_usdcToken;
  /// @notice The message transmitter proxy, which is used on destination as a non-upgradeable caller of all CCTP V2 messages.
  /// @dev Instead of calling receiveMessage directly, we use a proxy to enable upgrades to the verifier without invalidating in-flight messages.
  /// CCTP messages define an address permitted to call receiveMessage, which will always be the message transmitter proxy.
  CCTPMessageTransmitterProxy private immutable i_messageTransmitterProxy;
  /// @notice The token messenger, which is used on source to send USDC over CCTP V2.
  /// @dev The token messenger calls into the message transmitter after burning USDC and forming the app-specific message body.
  ITokenMessenger private immutable i_tokenMessenger;
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
    DynamicConfig memory dynamicConfig
  )
    // TODO: Construct with finality config?
    BaseVerifier(storageLocation)
  {
    if (
      address(tokenMessenger) == address(0) || address(messageTransmitterProxy) == address(0)
        || address(usdcToken) == address(0)
    ) revert ZeroAddressNotAllowed();

    // Ensure that the token messenger is for CCTP V2.
    uint32 tokenMessengerVersion = tokenMessenger.messageBodyVersion();
    if (tokenMessengerVersion != SUPPORTED_USDC_VERSION) {
      revert InvalidTokenMessengerVersion(SUPPORTED_USDC_VERSION, tokenMessengerVersion);
    }

    // Ensure that the message transmitter is for CCTP V2.
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
      address(tokenMessenger), address(messageTransmitterProxy), address(usdcToken), i_localDomainIdentifier
    );

    _setDynamicConfig(dynamicConfig);
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
    {
      address senderAddress = address(bytes20(message.sender));
      _assertSenderIsAllowed(message.destChainSelector, senderAddress);
    }

    Domain storage domain = s_chainToDomain[message.destChainSelector];
    if (!domain.enabled) revert UnknownDomain(message.destChainSelector);

    // We expect exactly one token transfer per message.
    // The address of the token transferred must correspond to USDC.
    if (message.tokenTransfer.length != 1) revert InvalidTokenTransferLength(message.tokenTransfer.length);
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = message.tokenTransfer[0];
    {
      address sourceTokenAddress = address(bytes20(tokenTransfer.sourceTokenAddress));
      if (sourceTokenAddress != address(i_usdcToken)) revert InvalidToken(sourceTokenAddress);
    }

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
        allowedCaller: domain.allowedCaller
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

    i_tokenMessenger.depositForBurnWithHook(
      params.amount,
      params.domainIdentifier,
      params.receiver,
      address(i_usdcToken),
      params.allowedCaller,
      // The maximum fee, taken on destination, is a percentage of the total amount transferred.
      // We use bps to calculate the smallest possible value that we can set as the max fee.
      // The bps values configured for each finality threshold on this chain must mirror those used by CCTP V2.
      // CCTP V2 defines different bps values for each chain.
      uint32(params.amount * bps / BPS_DIVIDER), // TODO: unsafe uint32 cast
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
  /// @return bps The bps charged on destinaton by CCTP for the finality threshold on this source chain.
  /// @return found Whether the finality threshold and bps were found. We can't rely on 0 values because 0 may be valid.
  function _getCCTPFinalityThresholdAndBps(
    uint32 finality
  ) internal view returns (uint32 finalityThreshold, uint16 bps, bool found) {
    if (finality == 0) {
      // Apply standard CCTP finality when CCIP finality is set to the default value of 0.
      return (s_finalityConfig.standardFinalityThreshold, s_finalityConfig.standardFinalityBps, true);
    } else {
      CustomFinality[] memory customFinalities = s_finalityConfig.customFinalities;
      for (uint256 i = 0; i < customFinalities.length; ++i) {
        if (i == customFinalities.length - 1 || finality < customFinalities[i].finality) {
          // If we've reached the last custom finality available, we must use it no matter what.
          // If we've reached a finality that is greater than the requested finality, we will round up to it.
          // This mirrors the behavior of CCTP finality thresholds, which round up if the requested finality exceeds a threshold.
          return (customFinalities[i].cctpFinalityThreshold, customFinalities[i].cctpFinalityBps, true);
        }
      }
    }

    return (0, 0, false);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageHash,
    bytes memory ccvData
  ) external pure {
    // TODO: Implement verification logic
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

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

  // TODO: setFinalityConfig, getFinalityConfig

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The dynamic configuration.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (dynamicConfig.feeAggregator == address(0)) revert ZeroAddressNotAllowed();

    s_dynamicConfig = dynamicConfig;

    emit DynamicConfigSet(dynamicConfig);
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
