// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCTPMessageTransmitterProxy} from "../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {MockE2EUSDCTransmitterCCTPV2} from "../../mocks/MockE2EUSDCTransmitterCCTPV2.sol";
import {MockUSDCTokenMessenger} from "../../mocks/MockUSDCTokenMessenger.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract CCTPVerifierSetup is BaseVerifierSetup {
  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageHeader {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    bytes32 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageBody {
    uint32 version;
    bytes32 burnToken;
    bytes32 mintRecipient;
    uint256 amount;
    bytes32 messageSender;
    uint256 maxFee;
    uint256 feeExecuted;
    uint256 expirationBlock;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessageHookData {
    bytes4 verifierVersion;
    bytes32 messageId;
  }

  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessage {
    CCTPMessageHeader header;
    CCTPMessageBody body;
    CCTPMessageHookData hookData;
  }

  CCTPVerifier internal s_cctpVerifier;
  MockUSDCTokenMessenger internal s_mockTokenMessenger;
  MockE2EUSDCTransmitterCCTPV2 internal s_mockMessageTransmitter;
  CCTPMessageTransmitterProxy internal s_messageTransmitterProxy;
  IBurnMintERC20 internal s_USDCToken;

  bytes internal s_tokenReceiver;
  address internal s_tokenReceiverAddress;

  uint256 internal constant TRANSFER_AMOUNT = 10e6; // 10 USDC
  uint16 internal constant BPS_DIVIDER = 10_000;

  uint32 internal constant CCTP_STANDARD_FINALITY_THRESHOLD = 2000;
  uint16 internal constant CCTP_STANDARD_FINALITY_BPS = 0;

  uint16 internal constant CCIP_FAST_FINALITY_THRESHOLD = 1;
  uint32 internal constant CCTP_FAST_FINALITY_THRESHOLD = 1000;
  uint16 internal constant CCTP_FAST_FINALITY_BPS = 2; // 0.02%

  uint32 internal constant REMOTE_DOMAIN_IDENTIFIER = 9999;
  uint32 internal constant LOCAL_DOMAIN_IDENTIFIER = 8888;
  bytes32 internal constant ALLOWED_CALLER_ON_DEST = keccak256("allowedCallerOnDest");
  bytes32 internal constant ALLOWED_CALLER_ON_SOURCE = keccak256("allowedCallerOnSource");

  function setUp() public virtual override {
    super.setUp();

    s_tokenReceiverAddress = makeAddr("tokenReceiver");
    s_tokenReceiver = abi.encode(s_tokenReceiverAddress);

    BurnMintERC20 usdcToken = new BurnMintERC20("USD Coin", "USDC", 6, 0, 0);
    s_USDCToken = usdcToken;

    s_mockMessageTransmitter = new MockE2EUSDCTransmitterCCTPV2(1, LOCAL_DOMAIN_IDENTIFIER, address(s_USDCToken));
    s_mockTokenMessenger = new MockUSDCTokenMessenger(1, address(s_mockMessageTransmitter));
    s_messageTransmitterProxy = new CCTPMessageTransmitterProxy(s_mockTokenMessenger);

    uint16[] memory customCCIPFinalities = new uint16[](1);
    customCCIPFinalities[0] = CCIP_FAST_FINALITY_THRESHOLD;

    uint32[] memory customCCTPFinalityThresholds = new uint32[](1);
    customCCTPFinalityThresholds[0] = CCTP_FAST_FINALITY_THRESHOLD;

    uint16[] memory customCCTPFinalityBps = new uint16[](1);
    customCCTPFinalityBps[0] = CCTP_FAST_FINALITY_BPS;

    s_cctpVerifier = new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: customCCIPFinalities,
        customCCTPFinalityThresholds: customCCTPFinalityThresholds,
        customCCTPFinalityBps: customCCTPFinalityBps
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );

    // Apply dest chain config updates.
    CCTPVerifier.DestChainConfigArgs[] memory destChainConfigArgs = new CCTPVerifier.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = BaseVerifier.DestChainConfigArgs({
      router: s_router,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_cctpVerifier.applyDestChainConfigUpdates(destChainConfigArgs);

    // Set the domains.
    CCTPVerifier.SetDomainArgs[] memory domains = new CCTPVerifier.SetDomainArgs[](1);
    domains[0] = CCTPVerifier.SetDomainArgs({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_cctpVerifier.setDomains(domains);

    // Grant mint and burn roles to the token messenger and the message transmitter.
    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockTokenMessenger));
    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockMessageTransmitter));

    // Ensure that the verifier is allowed to call the message transmitter proxy.
    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_cctpVerifier), allowed: true});
    s_messageTransmitterProxy.configureAllowedCallers(allowedCallerParams);
  }

  function _createCCVData(
    bytes4 verifierVersion,
    CCTPVerifierSetup.CCTPMessage memory cctpMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      verifierVersion, // Prefix for routing.
      _encodeCCTPMessage(cctpMessage),
      new bytes(65) // Signature.
    );
  }

  function _encodeCCTPMessage(
    CCTPVerifierSetup.CCTPMessage memory cctpMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      _encodeCCTPMessageHeader(cctpMessage.header),
      _encodeCCTPMessageBody(cctpMessage.body),
      _encodeCCTPMessageHookData(cctpMessage.hookData)
    );
  }

  function _encodeCCTPMessageHeader(
    CCTPVerifierSetup.CCTPMessageHeader memory header
  ) private pure returns (bytes memory) {
    return abi.encodePacked(
      header.version,
      header.sourceDomain,
      header.destinationDomain,
      header.nonce,
      header.sender,
      header.recipient,
      header.destinationCaller,
      header.minFinalityThreshold,
      header.finalityThresholdExecuted
    );
  }

  function _encodeCCTPMessageBody(
    CCTPVerifierSetup.CCTPMessageBody memory body
  ) private pure returns (bytes memory) {
    return abi.encodePacked(
      body.version,
      body.burnToken,
      body.mintRecipient,
      body.amount,
      body.messageSender,
      body.maxFee,
      body.feeExecuted,
      body.expirationBlock
    );
  }

  function _encodeCCTPMessageHookData(
    CCTPVerifierSetup.CCTPMessageHookData memory hookData
  ) private pure returns (bytes memory) {
    return abi.encodePacked(hookData.verifierVersion, hookData.messageId);
  }

  function _createCCIPMessage(
    uint64 sourceChainSelector,
    uint64 destChainSelector,
    uint16 finality,
    address sourceTokenAddress,
    uint256 amount,
    bytes memory tokenReceiver
  ) internal returns (MessageV1Codec.MessageV1 memory, bytes32) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(sourceTokenAddress),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: tokenReceiver,
      extraData: "extra data"
    });

    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: sourceChainSelector,
      destChainSelector: destChainSelector,
      sequenceNumber: 1,
      executionGasLimit: 400_000,
      ccipReceiveGasLimit: 200_000,
      finality: finality,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      sender: abi.encodePacked(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: tokenTransfer,
      data: ""
    });

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    return (message, messageId);
  }
}
