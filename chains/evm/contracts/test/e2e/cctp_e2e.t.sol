// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouter} from "../../interfaces/IRouter.sol";
import {ITokenMessenger} from "../../pools/USDC/interfaces/ITokenMessenger.sol";

import {Router} from "../../Router.sol";
import {CCTPVerifier} from "../../ccvs/CCTPVerifier.sol";
import {VersionedVerifierResolver} from "../../ccvs/VersionedVerifierResolver.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";
import {Internal} from "../../libraries/Internal.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {CCTPTokenPool} from "../../pools/USDC/CCTPTokenPool.sol";
import {USDCTokenPoolProxy} from "../../pools/USDC/USDCTokenPoolProxy.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MockE2EUSDCTransmitterCCTPV2} from "../mocks/MockE2EUSDCTransmitterCCTPV2.sol";
import {MockUSDCTokenMessenger} from "../mocks/MockUSDCTokenMessenger.sol";
import {e2e} from "./e2e.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract cctp_e2e is e2e {
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

  uint32 private constant CCTP_VERSION = 1;
  uint16 private constant CCTP_FAST_FINALITY_BPS = 2; // 0.02%

  uint32 private constant SOURCE_DOMAIN = 1;
  uint32 private constant DEST_DOMAIN = 2;

  address private s_feeAggregator = makeAddr("feeAggregator");
  address private s_allowlistAdmin = makeAddr("allowlistAdmin");

  struct CCTPSetup {
    address router;
    address tokenAdminRegistry;
    VersionedVerifierResolver verifierResolver;
    CCTPVerifier verifier;
    CCTPTokenPool tokenPool;
    MockUSDCTokenMessenger tokenMessenger;
    MockE2EUSDCTransmitterCCTPV2 messageTransmitter;
    CCTPMessageTransmitterProxy messageTransmitterProxy;
    USDCTokenPoolProxy tokenPoolProxy;
    IERC20 token;
  }

  CCTPSetup internal s_sourceCCTPSetup;
  CCTPSetup internal s_destCCTPSetup;

  function setUp() public override {
    super.setUp();

    s_sourceCCTPSetup =
      _deployCCTPSetup(address(s_sourceRouter), address(s_mockRMNRemote), address(s_tokenAdminRegistry), SOURCE_DOMAIN);
    s_destCCTPSetup =
      _deployCCTPSetup(address(s_sourceRouter), address(s_mockRMNRemote), address(s_tokenAdminRegistry), DEST_DOMAIN);

    _connectCCTPSetups();

    // Deal some USDC to the OWNER.
    deal(address(s_sourceCCTPSetup.token), OWNER, 1000e6); // 1000 USDC.

    s_sourcePoolByToken[address(s_sourceCCTPSetup.token)] = address(s_sourceCCTPSetup.tokenPoolProxy);
    s_destPoolByToken[address(s_destCCTPSetup.token)] = address(s_destCCTPSetup.tokenPoolProxy);
    s_destTokenBySourceToken[address(s_sourceCCTPSetup.token)] = address(s_destCCTPSetup.token);
    s_destPoolBySourceToken[address(s_sourceCCTPSetup.token)] = address(s_destCCTPSetup.tokenPoolProxy);
    s_extraDataByToken[address(s_sourceCCTPSetup.token)] =
      abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG);

    // Apply off ramp updates on the source router.
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);
  }

  function test_e2e() public override {
    uint256 amount = 1e6;

    vm.pauseGasMetering();
    uint64 expectedMsgNum = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber + 1;
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), type(uint256).max);
    IERC20(address(s_sourceCCTPSetup.token)).approve(address(s_sourceRouter), type(uint256).max);

    // Specify the CCTP verifier so that it is the only verifier included.
    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = address(s_sourceCCTPSetup.verifierResolver);
    bytes[] memory userCCVArgs = new bytes[](1);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "cctp e2e test data",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: userCCVAddresses,
          ccvArgs: userCCVArgs,
          blockConfirmations: 0,
          gasLimit: GAS_LIMIT,
          executor: address(0),
          executorArgs: "",
          tokenReceiver: "",
          tokenArgs: ""
        })
      )
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: address(s_sourceCCTPSetup.token), amount: amount}); // 1 USDC.

    (bytes32 messageId, bytes memory encodedMessage, OnRamp.Receipt[] memory receipts, bytes[] memory verifierBlobs) =
    _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      msgNum: expectedMsgNum,
      originalSender: OWNER
    });
    receipts[receipts.length - 1].issuer = address(s_sourceRouter);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_sourceCCTPSetup.token),
      amount,
      address(s_sourceCCTPSetup.verifier),
      bytes32(abi.encode(OWNER)),
      DEST_DOMAIN,
      s_sourceCCTPSetup.tokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      bytes32(abi.encode(address(s_destCCTPSetup.messageTransmitterProxy))),
      0,
      2000,
      abi.encodePacked(s_sourceCCTPSetup.verifier.versionTag(), messageId)
    );

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: expectedMsgNum,
      messageId: messageId,
      feeToken: s_sourceFeeToken,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });

    vm.resumeGasMetering();
    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);
    vm.pauseGasMetering();

    assertEq(s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber, expectedMsgNum);

    CCTPMessage memory cctpMessage = CCTPMessage({
      header: CCTPMessageHeader({
        version: CCTP_VERSION,
        sourceDomain: SOURCE_DOMAIN,
        destinationDomain: DEST_DOMAIN,
        nonce: bytes32(0),
        sender: bytes32(abi.encode(s_sourceCCTPSetup.tokenMessenger)),
        recipient: bytes32(abi.encode(address(s_destCCTPSetup.tokenMessenger))),
        destinationCaller: bytes32(abi.encode(address(s_destCCTPSetup.messageTransmitterProxy))),
        minFinalityThreshold: 2000,
        finalityThresholdExecuted: 2000
      }),
      body: CCTPMessageBody({
        version: CCTP_VERSION,
        burnToken: bytes32(abi.encode(address(s_sourceCCTPSetup.token))),
        mintRecipient: bytes32(abi.encode(OWNER)),
        amount: amount,
        messageSender: bytes32(abi.encode(s_sourceCCTPSetup.verifier)),
        maxFee: 0,
        feeExecuted: 0,
        expirationBlock: block.number + 1000
      }),
      hookData: CCTPMessageHookData({verifierVersion: s_sourceCCTPSetup.verifier.versionTag(), messageId: messageId})
    });

    // Default CCV is applied because receiver is an EOA.
    // Therefore, we need to include a verifier result for the mock verifier here as well.
    address[] memory ccvAddresses = new address[](2);
    ccvAddresses[0] = address(s_destCCTPSetup.verifierResolver);
    ccvAddresses[1] = address(s_destVerifier);
    bytes[] memory verifierResults = new bytes[](2);
    verifierResults[0] =
      abi.encodePacked(s_sourceCCTPSetup.verifier.versionTag(), _encodeCCTPMessage(cctpMessage), new bytes(65));
    verifierResults[1] = "";

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      messageNumber: expectedMsgNum,
      messageId: messageId,
      state: Internal.MessageExecutionState.SUCCESS,
      returnData: ""
    });

    vm.resumeGasMetering();
    s_offRamp.execute(encodedMessage, ccvAddresses, verifierResults);
  }

  function _deployCCTPSetup(
    address router,
    address rmn,
    address tokenAdminRegistry,
    uint32 localDomainIdentifier
  ) internal returns (CCTPSetup memory) {
    CCTPSetup memory setup;

    // Deploy all required contracts.
    setup.router = router;
    setup.tokenAdminRegistry = tokenAdminRegistry;
    setup.verifierResolver = new VersionedVerifierResolver();
    setup.token = new BurnMintERC20("USD Coin", "USDC", 6, 0, 0);
    setup.tokenPool =
      new CCTPTokenPool(IERC20(address(setup.token)), 6, rmn, router, address(setup.verifierResolver), new address[](0));
    setup.messageTransmitter =
      new MockE2EUSDCTransmitterCCTPV2(CCTP_VERSION, localDomainIdentifier, address(setup.token));
    setup.tokenMessenger = new MockUSDCTokenMessenger(CCTP_VERSION, address(setup.messageTransmitter));
    setup.messageTransmitterProxy = new CCTPMessageTransmitterProxy(setup.tokenMessenger);
    setup.verifier = new CCTPVerifier(
      setup.tokenMessenger,
      setup.messageTransmitterProxy,
      IERC20(address(setup.token)),
      new string[](0),
      CCTPVerifier.DynamicConfig({
        feeAggregator: s_feeAggregator,
        allowlistAdmin: s_allowlistAdmin,
        fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      rmn
    );
    setup.tokenPoolProxy = new USDCTokenPoolProxy(
      IERC20(address(setup.token)),
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: address(0),
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(setup.tokenPool)
      }),
      router,
      address(setup.verifierResolver)
    );

    return setup;
  }

  function _connectCCTPSetups() internal {
    // Set outbound implementation on the source verifier resolver.
    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImpls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outboundImpls[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      verifier: address(s_sourceCCTPSetup.verifier)
    });
    s_sourceCCTPSetup.verifierResolver.applyOutboundImplementationUpdates(outboundImpls);

    // Set inbound implementation on the dest verifier resolver.
    VersionedVerifierResolver.InboundImplementationArgs[] memory inboundImpls =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    inboundImpls[0] = VersionedVerifierResolver.InboundImplementationArgs({
      version: s_destCCTPSetup.verifier.versionTag(),
      verifier: address(s_destCCTPSetup.verifier)
    });
    s_destCCTPSetup.verifierResolver.applyInboundImplementationUpdates(inboundImpls);

    // Apply dest chain config updates on the source CCTP verifier.
    CCTPVerifier.DestChainConfigArgs[] memory destChainConfigArgs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = BaseVerifier.DestChainConfigArgs({
      router: IRouter(s_sourceCCTPSetup.router),
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_sourceCCTPSetup.verifier.applyDestChainConfigUpdates(destChainConfigArgs);

    // Set the destination domain on the source CCTP verifier.
    CCTPVerifier.SetDomainArgs[] memory destDomainArgs = new CCTPVerifier.SetDomainArgs[](1);
    destDomainArgs[0] = CCTPVerifier.SetDomainArgs({
      allowedCallerOnDest: bytes32(abi.encode(address(s_destCCTPSetup.messageTransmitterProxy))),
      allowedCallerOnSource: bytes32(abi.encode(makeAddr("allowedCallerOnSource"))), // No need to set in this context.
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: DEST_DOMAIN,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_sourceCCTPSetup.verifier.setDomains(destDomainArgs);

    // Set the source domain on the dest CCTP verifier.
    CCTPVerifier.SetDomainArgs[] memory sourceDomainArgs = new CCTPVerifier.SetDomainArgs[](1);
    sourceDomainArgs[0] = CCTPVerifier.SetDomainArgs({
      allowedCallerOnDest: bytes32(abi.encode(makeAddr("allowedCallerOnDest"))), // No need to set in this context.
      allowedCallerOnSource: bytes32(abi.encode(address(s_sourceCCTPSetup.verifier))),
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: SOURCE_DOMAIN,
      chainSelector: SOURCE_CHAIN_SELECTOR,
      enabled: true
    });
    s_destCCTPSetup.verifier.setDomains(sourceDomainArgs);

    // Set authorized callers on the message transmitter proxy on destination.
    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_destCCTPSetup.verifier), allowed: true});
    s_destCCTPSetup.messageTransmitterProxy.configureAllowedCallers(allowedCallerParams);

    // Apply chain updates on the destination token pool.
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_sourceCCTPSetup.tokenPoolProxy));
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(s_sourceCCTPSetup.token)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_destCCTPSetup.tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Apply chain updates on the source token pool.
    remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_destCCTPSetup.tokenPoolProxy));
    chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(s_destCCTPSetup.token)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_sourceCCTPSetup.tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Apply authorized callers on the source token pool.
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_sourceCCTPSetup.tokenPoolProxy);
    s_sourceCCTPSetup.tokenPool.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    // Apply authorized callers on the dest token pool.
    authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_destCCTPSetup.tokenPoolProxy);
    s_destCCTPSetup.tokenPool.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    // Update lock or burn mechanism on the source token pool proxy.
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2_WITH_CCV;
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = DEST_CHAIN_SELECTOR;
    s_sourceCCTPSetup.tokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    // Register CCTP token pool proxy on source token admin registry.
    TokenAdminRegistry(s_sourceCCTPSetup.tokenAdminRegistry).proposeAdministrator(
      address(s_sourceCCTPSetup.token), OWNER
    );
    TokenAdminRegistry(s_sourceCCTPSetup.tokenAdminRegistry).acceptAdminRole(address(s_sourceCCTPSetup.token));
    TokenAdminRegistry(s_sourceCCTPSetup.tokenAdminRegistry).setPool(
      address(s_sourceCCTPSetup.token), address(s_sourceCCTPSetup.tokenPoolProxy)
    );

    // Register CCTP token pool proxy on dest token admin registry.
    TokenAdminRegistry(s_destCCTPSetup.tokenAdminRegistry).proposeAdministrator(address(s_destCCTPSetup.token), OWNER);
    TokenAdminRegistry(s_destCCTPSetup.tokenAdminRegistry).acceptAdminRole(address(s_destCCTPSetup.token));
    TokenAdminRegistry(s_destCCTPSetup.tokenAdminRegistry).setPool(
      address(s_destCCTPSetup.token), address(s_destCCTPSetup.tokenPoolProxy)
    );

    // Grant burn and mint roles on the source token to the source token messenger.
    BurnMintERC20(address(s_sourceCCTPSetup.token)).grantMintAndBurnRoles(address(s_sourceCCTPSetup.tokenMessenger));

    // Grant burn and mint roles on the dest token to the dest message transmitter.
    BurnMintERC20(address(s_destCCTPSetup.token)).grantMintAndBurnRoles(address(s_destCCTPSetup.messageTransmitter));
  }

  function _encodeCCTPMessage(
    CCTPMessage memory cctpMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      _encodeCCTPMessageHeader(cctpMessage.header),
      _encodeCCTPMessageBody(cctpMessage.body),
      _encodeCCTPMessageHookData(cctpMessage.hookData)
    );
  }

  function _encodeCCTPMessageHeader(
    CCTPMessageHeader memory header
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
    CCTPMessageBody memory body
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
    CCTPMessageHookData memory hookData
  ) private pure returns (bytes memory) {
    return abi.encodePacked(hookData.verifierVersion, hookData.messageId);
  }
}
