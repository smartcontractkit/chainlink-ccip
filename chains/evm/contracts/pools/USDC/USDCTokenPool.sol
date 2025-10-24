// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMessageTransmitter} from "./interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";

import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "./CCTPMessageTransmitterProxy.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@4.8.3/utils/structs/EnumerableSet.sol";

/// @notice This pool mints and burns USDC tokens through the Cross Chain Transfer Protocol (CCTP).
/*
 OnRamp
   |
   | lockOrBurn()
   v
+------------------+    depositForBurn()    +------------------+    burn(from, localAmount)    +-----------+
| USDC Token Pool  | ---------------------> | Token Messenger  | ----------------------------> |   USDC    |
+------------------+                        +------------------+                               +-----------+

 OffRamp
   |
   | releaseOrMint()
   v
 USDC Token Pool
   |
   | receiveMessage()
   v
+------------------------+    receiveMessage()   +---------------------+    mint(amount, recipient)    +----------+
| CCTP Transmitter Proxy | --------------------> | Message Transmitter | ----------------------------> |   USDC   |
+------------------------+                       +---------------------+                               +----------+
*/
/// @dev This specific pool is used for CCTP V1. The CCTP V2 pool is a separate contract, which inherits many of the
/// state management from this contract, only overriding the functions absolutely necessary for supporting CCTP V2.
contract USDCTokenPool is TokenPool, ITypeAndVersion, AuthorizedCallers {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;

  event DomainsSet(DomainUpdate[]);
  event ConfigSet(address tokenMessenger);

  error UnknownDomain(uint64 domain);
  error UnlockingUSDCFailed();
  error InvalidConfig();
  error InvalidDomain(DomainUpdate domain);
  error InvalidMessageVersion(uint32 expected, uint32 got);
  error InvalidTokenMessengerVersion(uint32 expected, uint32 got);
  error InvalidNonce(uint64 expected, uint64 got);
  error InvalidSourceDomain(uint32 expected, uint32 got);
  error InvalidDestinationDomain(uint32 expected, uint32 got);
  error InvalidReceiver(bytes receiver);
  error InvalidTransmitterInProxy();
  error InvalidPreviousPool();
  error InvalidMessageLength(uint256 length);

  // This data is supplied from offchain and contains everything needed to mint the USDC tokens on the destination chain
  // through CCTP.
  struct MessageAndAttestation {
    bytes message;
    bytes attestation;
  }

  // solhint-disable-next-line gas-struct-packing
  struct DomainUpdate {
    bytes32 allowedCaller; // Address allowed to mint on the domain (destination MessageTransmitterProxy)
    bytes32 mintRecipient; // Address to mint to on the destination chain
    uint32 domainIdentifier; // Unique domain ID
    uint64 destChainSelector; // The destination chain for this domain
    bool enabled; // Whether the domain is enabled
    bool useLegacySourcePoolDataFormat; // Whether to use the legacy source pool data format
  }

  /// @notice The version of the USDC message format that this pool supports. Version 0 is the legacy version of CCTP.
  uint32 public immutable i_supportedUSDCVersion;

  // The local USDC config.
  ITokenMessenger public immutable i_tokenMessenger;
  CCTPMessageTransmitterProxy public immutable i_messageTransmitterProxy;
  uint32 public immutable i_localDomainIdentifier;

  /// A domain is a USDC representation of a destination chain.
  /// @dev Zero is a valid domain identifier.
  /// @dev The address to mint on the destination chain is the corresponding USDC pool.
  /// @dev The allowedCaller represents the contract authorized to call receiveMessage on the destination CCTP message
  /// transmitter. For dest pool version 1.6.1, this is the MessageTransmitterProxy of the destination chain. For dest
  /// pool version 1.5.1, this is the destination chain's token pool.
  // solhint-disable-next-line gas-struct-packing
  struct Domain {
    bytes32 allowedCaller; //                 Address allowed to mint on the domain
    bytes32 mintRecipient; //                 Address to mint to on the destination chain
    uint32 domainIdentifier; // ────────────╮ Unique domain ID
    bool enabled; //                        | Whether the domain is enabled
    bool useLegacySourcePoolDataFormat; // ─╯ Whether to use the legacy source pool data format for chains that
      // have not yet been updated to the new source pool data format.
  }

  // A mapping of CCIP chain identifiers to destination domains
  mapping(uint64 chainSelector => Domain CCTPDomain) internal s_chainToDomain;

  /// @dev The authorized callers are set as empty since the USDCTokenPoolProxy is the only authorized caller,
  /// but cannot be deployed until after this contract is deployed. The allowed callers are set after deployment.
  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    uint32 supportedUSDCVersion
  ) TokenPool(token, 6, allowlist, rmnProxy, router) AuthorizedCallers(new address[](0)) {
    // The version of the USDC message format that this pool supports. Version 0 is the legacy version of CCTP.
    i_supportedUSDCVersion = supportedUSDCVersion;

    // The token messenger, which is used for outgoing messages (burn operations), has a corresponding message transmitter
    // that is used for incoming messages (releaseOrMint).
    if (address(tokenMessenger) == address(0)) revert InvalidConfig();
    IMessageTransmitter transmitter = IMessageTransmitter(tokenMessenger.localMessageTransmitter());

    uint32 transmitterVersion = transmitter.version();

    // Check that the message transmitter version is supported by this version of the token pool.
    if (transmitterVersion != i_supportedUSDCVersion) {
      revert InvalidMessageVersion(i_supportedUSDCVersion, transmitterVersion);
    }

    // Check that the token messenger version is supported by this version of the token pool.
    uint32 tokenMessengerVersion = tokenMessenger.messageBodyVersion();

    // If the token messenger version is not supported, revert.
    if (tokenMessengerVersion != i_supportedUSDCVersion) {
      revert InvalidTokenMessengerVersion(i_supportedUSDCVersion, tokenMessengerVersion);
    }

    // Check that the message transmitter proxy is configured to use the correct message transmitter for
    // incoming messages (releaseOrMint).
    if (cctpMessageTransmitterProxy.i_cctpTransmitter() != transmitter) revert InvalidTransmitterInProxy();

    i_tokenMessenger = tokenMessenger;
    i_messageTransmitterProxy = cctpMessageTransmitterProxy;
    i_localDomainIdentifier = transmitter.localDomain();

    // Allow the token messenger to burn tokens on behalf of this pool.
    i_token.safeIncreaseAllowance(address(i_tokenMessenger), type(uint256).max);

    emit ConfigSet(address(tokenMessenger));
  }

  /// @notice Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "USDCTokenPool 1.6.x-dev";
  }

  /// @notice Burn tokens from the pool to initiate cross-chain transfer.
  /// @notice Outgoing messages (burn operations) are routed via `i_tokenMessenger.depositForBurnWithCaller`.
  /// The allowedCaller is preconfigured per destination domain and token pool version refer Domain struct.
  /// @dev Emits ITokenMessenger.DepositForBurn event.
  /// @dev Assumes caller has validated the destinationReceiver.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    Domain memory domain = s_chainToDomain[lockOrBurnIn.remoteChainSelector];
    if (!domain.enabled) revert UnknownDomain(lockOrBurnIn.remoteChainSelector);

    if (lockOrBurnIn.receiver.length != 32) {
      revert InvalidReceiver(lockOrBurnIn.receiver);
    }

    bytes32 decodedReceiver;
    // For EVM chains, the mintRecipient is not used, but is needed for Solana, where the mintRecipient will
    // be a PDA owned by the pool, and will forward the tokens to its final destination after minting.
    if (domain.mintRecipient != bytes32(0)) {
      decodedReceiver = domain.mintRecipient;
    } else {
      decodedReceiver = abi.decode(lockOrBurnIn.receiver, (bytes32));
    }

    // Since this pool is the msg sender of the CCTP transaction, only this contract
    // is able to call replaceDepositForBurn. Since this contract does not implement
    // replaceDepositForBurn, the tokens cannot be maliciously re-routed to another address.
    uint64 nonce = i_tokenMessenger.depositForBurnWithCaller(
      lockOrBurnIn.amount, domain.domainIdentifier, decodedReceiver, address(i_token), domain.allowedCaller
    );

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    bytes memory sourcePoolData;
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory sourceTokenDataPayload =
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: nonce, sourceDomain: i_localDomainIdentifier});

    // The useLegacySourcePoolDataFormat flag is set to false for chains that have been updated to the new source pool
    // data format. When the lane is updated, the flag should be set to false.
    if (domain.useLegacySourcePoolDataFormat) {
      // Since not all lanes will be updated to the new source pool data format simultaneously, it is important to support
      // the legacy format until such a time as the lane can support it. Otherwise, the destination pool would not be able
      // to parse the source pool data and all messages originating from this updated token pool would be rejected.

      // It is safe to have the legacy format still be supported temporarily, as the USDCTokenPoolProxy will convert
      // the legacy format to the new format before releaseOrMint() is called. Once all lanes in CCIP are updated to
      // the new format and CCTP V2, this branch can be safely removed.
      sourcePoolData = abi.encode(sourceTokenDataPayload);
    } else {
      sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(sourceTokenDataPayload);
    }

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: sourcePoolData
    });
  }

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned onRamp for the given chain on the Router.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    _validateCaller();
  }

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned offRamp for the given chain on the Router.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    _validateCaller();
  }

  /// @inheritdoc TokenPool
  /// @dev This function proxies the message to the message transmitter, which will mint the tokens through CCTP's contracts.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    // Decode the source pool data from its raw bytes into a SourceTokenDataPayloadV0 struct that can be
    // more easily validated. Since the USDCTokenPoolProxy that sits between this pool and the offRamp will convert
    // the legacy format to the new format, this operation is safe to perform, as a message originating from a legacy
    // pool will be converted to the new format before this decoding function is called.
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory sourceTokenDataPayload =
      USDCSourcePoolDataCodec._decodeSourceTokenDataPayloadV1(releaseOrMintIn.sourcePoolData);

    _validateMessage(msgAndAttestation.message, sourceTokenDataPayload);

    // Proxy the message to the message transmitter, which will mint the tokens through CCTP's contracts.
    if (!i_messageTransmitterProxy.receiveMessage(msgAndAttestation.message, msgAndAttestation.attestation)) {
      revert UnlockingUSDCFailed();
    }

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.sourceDenominatedAmount});
  }

  /// @notice Validates the USDC encoded message against the given parameters.
  /// @param usdcMessage The USDC encoded message
  /// @param sourceTokenData The expected source chain token data to check against
  /// @dev Only supports version SUPPORTED_USDC_VERSION of the CCTP message format
  /// @dev Message format for USDC:
  ///     * Field                 Bytes      Type       Index
  ///     * version               4          uint32     0
  ///     * sourceDomain          4          uint32     4
  ///     * destinationDomain     4          uint32     8
  ///     * nonce                 8          uint64     12
  ///     * sender                32         bytes32    20
  ///     * recipient             32         bytes32    52
  ///     * destinationCaller     32         bytes32    84
  ///     * messageBody           dynamic    bytes      116
  function _validateMessage(
    bytes memory usdcMessage,
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory sourceTokenData
  ) internal view virtual {
    // 116 is the minimum length of a valid USDC message. Since destinationCaller must be checked for the
    // previous pool, this ensures it can be parsed correctly and that the message is not too short.
    // Since messageBody is dynamic and not always used, it is not checked.
    if (usdcMessage.length < 116) revert InvalidMessageLength(usdcMessage.length);

    uint32 version;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, so only the first 4 bytes
      // of the message remain when cast to uint32. We want the lower 4 bytes to be the version
      // when cast to uint32, so we add 4. If you added 32 (to skip the first word containing
      // the length), version would be in the upper 4 bytes of the corresponding slot, which
      // would not be as easily parsed into a uint32.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }
    // This token pool only supports version 0 of the CCTP message format
    // We check the version prior to loading the rest of the message
    // to avoid unexpected reverts due to out-of-bounds reads.
    if (version != i_supportedUSDCVersion) revert InvalidMessageVersion(i_supportedUSDCVersion, version);

    uint32 sourceDomain;
    uint32 destinationDomain;
    uint64 nonce;

    // solhint-disable-next-line no-inline-assembly
    assembly {
      sourceDomain := mload(add(usdcMessage, 8)) // 4 + 4 = 8
      destinationDomain := mload(add(usdcMessage, 12)) // 8 + 4 = 12
      nonce := mload(add(usdcMessage, 20)) // 12 + 8 = 20
    }

    if (sourceDomain != sourceTokenData.sourceDomain) {
      revert InvalidSourceDomain(sourceTokenData.sourceDomain, sourceDomain);
    }

    if (destinationDomain != i_localDomainIdentifier) {
      revert InvalidDestinationDomain(i_localDomainIdentifier, destinationDomain);
    }

    if (nonce != sourceTokenData.nonce) revert InvalidNonce(sourceTokenData.nonce, nonce);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Gets the CCTP domain for a given CCIP chain selector.
  /// @param chainSelector The CCIP chain selector to get the domain for.
  /// @return The CCTP domain for the given chain selector.
  function getDomain(
    uint64 chainSelector
  ) external view returns (Domain memory) {
    return s_chainToDomain[chainSelector];
  }

  /// @notice Sets the CCTP domain for a CCIP chain selector.
  /// @param domains The array of DomainUpdate structs to set.
  /// @dev Must verify mapping of selectors -> (domain, caller) offchain.
  function setDomains(
    DomainUpdate[] calldata domains
  ) external onlyOwner {
    for (uint256 i = 0; i < domains.length; ++i) {
      DomainUpdate memory domain = domains[i];
      if (domain.allowedCaller == bytes32(0) || domain.destChainSelector == 0) revert InvalidDomain(domain);

      s_chainToDomain[domain.destChainSelector] = Domain({
        allowedCaller: domain.allowedCaller,
        mintRecipient: domain.mintRecipient,
        domainIdentifier: domain.domainIdentifier,
        enabled: domain.enabled,
        useLegacySourcePoolDataFormat: domain.useLegacySourcePoolDataFormat
      });
    }
    emit DomainsSet(domains);
  }
}
