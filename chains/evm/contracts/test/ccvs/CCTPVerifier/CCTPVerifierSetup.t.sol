// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPMessageTransmitterProxy} from "../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {MockE2EUSDCTransmitterCCTPV2} from "../../mocks/MockE2EUSDCTransmitterCCTPV2.sol";
import {MockUSDCTokenMessenger} from "../../mocks/MockUSDCTokenMessenger.sol";

import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

contract CCTPVerifierSetup is BaseVerifierSetup {
  // solhint-disable-next-line gas-struct-packing
  struct CCTPMessage {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    bytes32 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
    bytes messageBody;
  }

  CCTPVerifier internal s_cctpVerifier;
  MockUSDCTokenMessenger internal s_mockTokenMessenger;
  MockE2EUSDCTransmitterCCTPV2 internal s_mockMessageTransmitter;
  CCTPMessageTransmitterProxy internal s_messageTransmitterProxy;
  IBurnMintERC20 internal s_USDCToken;

  uint16 internal constant BPS_DIVIDER = 10_000;

  uint32 internal constant CCTP_STANDARD_FINALITY_THRESHOLD = 2000;
  uint16 internal constant CCTP_STANDARD_FINALITY_BPS = 0;

  uint16 internal constant CCIP_FAST_FINALITY_THRESHOLD = 1;
  uint32 internal constant CCTP_FAST_FINALITY_THRESHOLD = 1000;
  uint16 internal constant CCTP_FAST_FINALITY_BPS = 2; // 0.02%

  uint32 internal constant DEST_DOMAIN_IDENTIFIER = 9999;
  uint32 internal constant LOCAL_DOMAIN_IDENTIFIER = 8888;
  bytes32 internal constant ALLOWED_CALLER_ON_DEST = keccak256("allowedCallerOnDest");
  bytes32 internal constant ALLOWED_CALLER_ON_SOURCE = keccak256("allowedCallerOnSource");

  function setUp() public virtual override {
    super.setUp();

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
    CCTPVerifier.DomainUpdate[] memory domains = new CCTPVerifier.DomainUpdate[](1);
    domains[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipient: bytes32(0),
      domainIdentifier: DEST_DOMAIN_IDENTIFIER,
      destChainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_cctpVerifier.setDomains(domains);
  }

  function _generateCCTPMessage(
    CCTPMessage memory cctpMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      cctpMessage.version,
      cctpMessage.sourceDomain,
      cctpMessage.destinationDomain,
      cctpMessage.nonce,
      cctpMessage.sender,
      cctpMessage.recipient,
      cctpMessage.destinationCaller,
      cctpMessage.minFinalityThreshold,
      cctpMessage.finalityThresholdExecuted,
      cctpMessage.messageBody
    );
  }
}
