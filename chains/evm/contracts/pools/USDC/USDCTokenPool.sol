// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../interfaces/IPool.sol";
import {IMessageTransmitter} from "./interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "./CCTPMessageTransmitterProxy.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

/// @notice This pool mints and burns USDC tokens through the Cross Chain Transfer
/// Protocol (CCTP).
contract USDCTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  event DomainsSet(DomainUpdate[]);
  event ConfigSet(address tokenMessenger);

  error UnknownDomain(uint64 domain);
  error UnlockingUSDCFailed();
  error InvalidConfig();
  error InvalidDomain(DomainUpdate domain);
  error InvalidMessageVersion(uint32 version);
  error InvalidTokenMessengerVersion(uint32 version);
  error InvalidNonce(uint64 expected, uint64 got);
  error InvalidSourceDomain(uint32 expected, uint32 got);
  error InvalidDestinationDomain(uint32 expected, uint32 got);
  error InvalidReceiver(bytes receiver);
  error InvalidTransmitterInProxy();
  error InvalidPreviousPool();
  error InvalidMessageLength(uint256 length);
  error InvalidMinFinalityThreshold(uint32 expected, uint32 got);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 got);

  // This data is supplied from offchain and contains everything needed
  // to receive the USDC tokens.
  struct MessageAndAttestation {
    bytes message;
    bytes attestation;
  }

  // A domain is a USDC representation of a chain.
  struct DomainUpdate {
    bytes32 allowedCaller; //       Address allowed to mint on the domain
    bytes32 mintRecipient; //       Address to mint to on the destination chain
    uint32 domainIdentifier; // ──╮ Unique domain ID
    uint64 destChainSelector; //  │ The destination chain for this domain
    bool enabled; // ─────────────╯ Whether the domain is enabled
  }

  struct SourceTokenDataPayload {
    uint64 nonce;
    uint32 sourceDomain;
  }

  string public constant override typeAndVersion = "USDCTokenPool 1.6.1-dev";

  // We restrict to the first version. New pool may be required for subsequent versions.
  uint32 public constant SUPPORTED_USDC_VERSION = 1;

  // CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  // TODO: Add Comment
  uint32 public constant FINALITY_THRESHOLD = 2000;

  // The local USDC config
  ITokenMessenger public immutable i_tokenMessenger;
  CCTPMessageTransmitterProxy public immutable i_messageTransmitterProxy;
  uint32 public immutable i_localDomainIdentifier;

  /// A domain is a USDC representation of a destination chain.
  /// @dev Zero is a valid domain identifier.
  /// @dev The address to mint on the destination chain is the corresponding USDC pool.
  /// @dev The allowedCaller represents the contract authorized to call receiveMessage on the destination CCTP message transmitter.
  /// For dest pool version 1.6.1, this is the MessageTransmitterProxy of the destination chain.
  /// For dest pool version 1.5.1, this is the destination chain's token pool.
  struct Domain {
    bytes32 allowedCaller; //      Address allowed to mint on the domain
    bytes32 mintRecipient; //      Address to mint to on the destination chain
    uint32 domainIdentifier; // ─╮ Unique domain ID
    bool enabled; // ────────────╯ Whether the domain is enabled
  }

  // A mapping of CCIP chain identifiers to destination domains
  mapping(uint64 chainSelector => Domain CCTPDomain) private s_chainToDomain;

  // In the event of an inflight message during a token pool migration, we need to route the message to the
  // previous pool to satisfy the allowedCaller. The currently in-use token pool must be set as an offRamp
  // in the router in order for the previous pool to accept the incoming call.
  address public immutable i_previousPool;
  address public immutable i_previousMessageTransmitterProxy;

  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address previousPool
  ) TokenPool(token, 6, allowlist, rmnProxy, router) {
    if (address(tokenMessenger) == address(0)) revert InvalidConfig();
    IMessageTransmitter transmitter = IMessageTransmitter(tokenMessenger.localMessageTransmitter());
    uint32 transmitterVersion = transmitter.version();
    if (transmitterVersion != SUPPORTED_USDC_VERSION) revert InvalidMessageVersion(transmitterVersion);
    uint32 tokenMessengerVersion = tokenMessenger.messageBodyVersion();
    if (tokenMessengerVersion != SUPPORTED_USDC_VERSION) revert InvalidTokenMessengerVersion(tokenMessengerVersion);
    if (cctpMessageTransmitterProxy.i_cctpTransmitter() != transmitter) revert InvalidTransmitterInProxy();

    i_tokenMessenger = tokenMessenger;
    i_messageTransmitterProxy = cctpMessageTransmitterProxy;
    i_localDomainIdentifier = transmitter.localDomain();
    i_token.safeIncreaseAllowance(address(i_tokenMessenger), type(uint256).max);

    // PreviousPool should not be current pool.
    if (previousPool == address(this)) {
      revert InvalidPreviousPool();
    }
    // If previousPool exists, it should be a valid token pool, we check it with supportsInterface.
    if (previousPool != address(0) && !IERC165(previousPool).supportsInterface(type(IPoolV1).interfaceId)) {
      revert InvalidPreviousPool();
    }

    if (previousPool != address(0)) {
      try USDCTokenPool(previousPool).i_messageTransmitterProxy() returns (CCTPMessageTransmitterProxy proxy) {
        i_previousMessageTransmitterProxy = address(proxy);
      } catch {
        revert InvalidPreviousPool();
      }
    } else {
      i_previousMessageTransmitterProxy = address(0);
    }

    i_previousPool = previousPool;

    emit ConfigSet(address(tokenMessenger));
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
    i_tokenMessenger.depositForBurn(
      lockOrBurnIn.amount,
      domain.domainIdentifier,
      decodedReceiver,
      address(i_token),
      domain.allowedCaller,
      MAX_FEE,
      FINALITY_THRESHOLD
    );

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    // As of CCTP v2, the nonce is not returned to this contract upon sending the message, and will instead be
    // acquired off-chain and included in the destination-message's offchainTokenData, so we set it to 0.
    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(SourceTokenDataPayload({nonce: uint64(0), sourceDomain: i_localDomainIdentifier}))
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// * sourceTokenData is part of the verified message and passed directly from
  /// the offRamp so it is guaranteed to be what the lockOrBurn pool released on the
  /// source chain. It contains (nonce, sourceDomain) which is guaranteed by CCTP
  /// to be unique.
  /// * offchainTokenData is untrusted (can be supplied by manual execution), but we assert
  /// that (nonce, sourceDomain) is equal to the message's (nonce, sourceDomain) and
  /// receiveMessage will assert that Attestation contains a valid attestation signature
  /// for that message, including its (nonce, sourceDomain). This way, the only
  /// non-reverting offchainTokenData that can be supplied is a valid attestation for the
  /// specific message that was sent on source.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);
    SourceTokenDataPayload memory sourceTokenDataPayload =
      abi.decode(releaseOrMintIn.sourcePoolData, (SourceTokenDataPayload));

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    // If the destinationCaller is the previous pool, indicating an inflight message during the migration, we need to
    // route the message to the previous pool to satisfy the allowedCaller.
    bytes32 destinationCallerBytes32;
    bytes memory messageBytes = msgAndAttestation.message;
    assembly {
      // destinationCaller is a 32-byte word starting at position 84 in messageBytes body, so add 32 to skip the 1st word
      // representing bytes length
      destinationCallerBytes32 := mload(add(messageBytes, 140)) // 108 + 32 = 140
    }
    address destinationCaller = address(uint160(uint256(destinationCallerBytes32)));

    // TODO: Fix this for previous pool. Consider checking the previous pool's transmitter proxy.
    if (i_previousPool != address(0) && destinationCaller == i_previousMessageTransmitterProxy) {
      // If the destinationCaller is the previous pool's message transmitter proxy, we can use this
      // as an indication that CCTP V1 was used to send the message, and route it to the previous pool for minting.
      return USDCTokenPool(i_previousPool).releaseOrMint(releaseOrMintIn);
    }

    // We call this after the destinationCaller check to ensure that the message is valid for CCTP V2. If it was called
    // before, then a V1 message which should be forwarded to the previous pool will be rejected.
    _validateMessage(msgAndAttestation.message, sourceTokenDataPayload);

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
  /// @param sourceTokenData The expected source chain CCTP identifier as provided by the CCIP-Source-Pool.
  /// @dev Only supports version SUPPORTED_USDC_VERSION of the CCTP V2 message format
  /// @dev Message format for USDC:
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * sourceDomain               4          uint32     4
  ///     * destinationDomain          4          uint32     8
  ///     * nonce                      32         bytes32   12
  ///     * sender                     32         bytes32   44
  ///     * recipient                  32         bytes32   76
  ///     * destinationCaller          32         bytes32   108
  ///     * minFinalityThreshold       32         uint32    140
  ///     * finalityThresholdExecuted  32         uint32    144
  ///     * messageBody                dynamic    bytes     148
  function _validateMessage(bytes memory usdcMessage, SourceTokenDataPayload memory sourceTokenData) internal view {
    // 116 is the minimum length of a valid USDC message. Since destinationCaller needs to be checked for the previous
    // pool, this ensures that it can be parsed correctly and that the message is not too short. Since messageBody is
    // dynamic and not always used, it is not checked.
    if (usdcMessage.length < 148) revert InvalidMessageLength(usdcMessage.length);

    uint32 version;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, meaning
      // we will only be left with the first 4 bytes of the message when we cast it to uint32. We want the lower 4 bytes
      // to be the version when casted to a uint32 , so we only add 4. If you added 32, attempting to skip the first word
      // containing the length, then version would be in the upper-4 bytes of the corresponding slot, which
      // would not be as easily parsed into a uint32.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }
    // This token pool only supports version 0 of the CCTP message format
    // We check the version prior to loading the rest of the message
    // to avoid unexpected reverts due to out-of-bounds reads.
    if (version != SUPPORTED_USDC_VERSION) revert InvalidMessageVersion(version);

    uint32 messageSourceDomain;
    uint32 destinationDomain;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;

    // solhint-disable-next-line no-inline-assembly
    assembly {
      messageSourceDomain := mload(add(usdcMessage, 8)) // 4 + 4 = 8
      destinationDomain := mload(add(usdcMessage, 12)) // 8 + 4 = 12
      minFinalityThreshold := mload(add(usdcMessage, 144)) // 140 + 4 = 144
      finalityThresholdExecuted := mload(add(usdcMessage, 148)) // 144 + 4 = 148
    }

    // Check that the source domain included in the CCTP Message matches the one forwarded by the source pool.
    if (messageSourceDomain != sourceTokenData.sourceDomain) {
      revert InvalidSourceDomain(sourceTokenData.sourceDomain, messageSourceDomain);
    }

    // Check that the destination domain in the CCTP message matches the immutable domain of this pool.
    if (destinationDomain != i_localDomainIdentifier) {
      revert InvalidDestinationDomain(i_localDomainIdentifier, destinationDomain);
    }

    // This pool only supports slow transfers on CCTP, so ensure that the message matches the same requirements.
    if (minFinalityThreshold != FINALITY_THRESHOLD) {
      revert InvalidMinFinalityThreshold(FINALITY_THRESHOLD, minFinalityThreshold);
    }

    if (finalityThresholdExecuted != FINALITY_THRESHOLD) {
      revert InvalidExecutionFinalityThreshold(FINALITY_THRESHOLD, finalityThresholdExecuted);
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Gets the CCTP domain for a given CCIP chain selector.
  function getDomain(
    uint64 chainSelector
  ) external view returns (Domain memory) {
    return s_chainToDomain[chainSelector];
  }

  /// @notice Sets the CCTP domain for a CCIP chain selector.
  /// @dev Must verify mapping of selectors -> (domain, caller) offchain.
  function setDomains(
    DomainUpdate[] calldata domains
  ) external onlyOwner {
    for (uint256 i = 0; i < domains.length; ++i) {
      DomainUpdate memory domain = domains[i];
      if (domain.allowedCaller == bytes32(0) || domain.destChainSelector == 0) revert InvalidDomain(domain);

      s_chainToDomain[domain.destChainSelector] = Domain({
        domainIdentifier: domain.domainIdentifier,
        mintRecipient: domain.mintRecipient,
        allowedCaller: domain.allowedCaller,
        enabled: domain.enabled
      });
    }
    emit DomainsSet(domains);
  }
}
