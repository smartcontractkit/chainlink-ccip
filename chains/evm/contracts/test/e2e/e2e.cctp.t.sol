// SPDX-License-Identifier: BUSL-1.1
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
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {CCTPTokenPool} from "../../pools/USDC/CCTPTokenPool.sol";
import {USDCTokenPoolProxy} from "../../pools/USDC/USDCTokenPoolProxy.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {CCTPHelper} from "../helpers/CCTPHelper.sol";
import {OffRampHelper} from "../helpers/OffRampHelper.sol";
import {MockE2EUSDCTransmitterCCTPV2} from "../mocks/MockE2EUSDCTransmitterCCTPV2.sol";
import {MockUSDCTokenMessenger} from "../mocks/MockUSDCTokenMessenger.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";
import {OnRampSetup} from "../onRamp/OnRamp/OnRampSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract cctp_e2e is OnRampSetup {
  uint32 private constant CCTP_VERSION = 1;
  uint16 private constant CCTP_FAST_FINALITY_BPS = 2; // 0.02%

  uint32 private constant SOURCE_DOMAIN = 1;
  uint32 private constant DEST_DOMAIN = 2;

  OffRampHelper private s_offRamp;
  MockVerifier private s_defaultSourceVerifier;
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

    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      })
    );

    s_sourceCCTPSetup =
      _deployCCTPSetup(address(s_sourceRouter), address(s_mockRMNRemote), address(s_tokenAdminRegistry), SOURCE_DOMAIN);
    s_destCCTPSetup =
      _deployCCTPSetup(address(s_sourceRouter), address(s_mockRMNRemote), address(s_tokenAdminRegistry), DEST_DOMAIN);

    _connectCCTPSetups();

    s_defaultSourceVerifier = new MockVerifier("");

    // Deal some USDC to the OWNER.
    deal(address(s_sourceCCTPSetup.token), OWNER, 1000e6); // 1000 USDC.

    s_sourcePoolByToken[address(s_sourceCCTPSetup.token)] = address(s_sourceCCTPSetup.tokenPoolProxy);
    s_destPoolByToken[address(s_destCCTPSetup.token)] = address(s_destCCTPSetup.tokenPoolProxy);
    s_destTokenBySourceToken[address(s_sourceCCTPSetup.token)] = address(s_destCCTPSetup.token);
    s_destPoolBySourceToken[address(s_sourceCCTPSetup.token)] = address(s_destCCTPSetup.tokenPoolProxy);
    s_extraDataByToken[address(s_sourceCCTPSetup.token)] =
      abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG);

    // Apply off ramp and on ramp updates on the source router.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    // Set the default CCV on the off ramp.
    address[] memory defaultDestCCVs = new address[](1);
    defaultDestCCVs[0] = address(s_destCCTPSetup.verifierResolver);

    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encode(s_onRamp);

    OffRamp.SourceChainConfigArgs[] memory sourceChainUpdates = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainUpdates[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCVs: defaultDestCCVs,
      laneMandatedCCVs: new address[](0)
    });
    s_offRamp.applySourceChainConfigUpdates(sourceChainUpdates);

    // Set dest chain config on the on ramp.
    address[] memory defaultSourceCCVs = new address[](1);
    defaultSourceCCVs[0] = address(s_defaultSourceVerifier);
    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultSourceCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
    destChainConfigArgs[0].offRamp = abi.encodePacked(address(s_offRamp));
    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function test_cctp_e2e() public {
    uint256 amount = 1e6;

    vm.pauseGasMetering();
    uint64 expectedMsgNum = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber + 1;
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), type(uint256).max);
    IERC20(address(s_sourceCCTPSetup.token)).approve(address(s_sourceRouter), type(uint256).max);

    // Specify the CCTP verifier so that it is the only verifier included.
    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = address(s_sourceCCTPSetup.verifierResolver);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "cctp e2e test data",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: userCCVAddresses,
          ccvArgs: new bytes[](1),
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

    CCTPHelper.CCTPMessage memory cctpMessage = CCTPHelper.CCTPMessage({
      header: CCTPHelper.CCTPMessageHeader({
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
      body: CCTPHelper.CCTPMessageBody({
        version: CCTP_VERSION,
        burnToken: bytes32(abi.encode(address(s_sourceCCTPSetup.token))),
        mintRecipient: bytes32(abi.encode(OWNER)),
        amount: amount,
        messageSender: bytes32(abi.encode(s_sourceCCTPSetup.verifier)),
        maxFee: 0,
        feeExecuted: 0,
        expirationBlock: block.number + 1000
      }),
      hookData: CCTPHelper.CCTPMessageHookData({
        verifierVersion: s_sourceCCTPSetup.verifier.versionTag(),
        messageId: messageId
      })
    });

    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = address(s_destCCTPSetup.verifierResolver);
    bytes[] memory verifierResults = new bytes[](1);
    verifierResults[0] = abi.encodePacked(
      s_sourceCCTPSetup.verifier.versionTag(), CCTPHelper._encodeCCTPMessage(cctpMessage), new bytes(65)
    );

    MessageV1Codec.MessageV1 memory messageV1 = this._decodeMessageV1(encodedMessage);

    vm.expectCall(
      address(s_destCCTPSetup.verifier),
      abi.encodeCall(CCTPVerifier.verifyMessage, (messageV1, messageId, verifierResults[0]))
    );
    vm.expectCall(
      address(s_destCCTPSetup.messageTransmitter),
      abi.encodeCall(
        MockE2EUSDCTransmitterCCTPV2.receiveMessage, (CCTPHelper._encodeCCTPMessage(cctpMessage), new bytes(65))
      )
    );

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      messageNumber: expectedMsgNum,
      messageId: messageId,
      state: Internal.MessageExecutionState.SUCCESS,
      returnData: ""
    });

    vm.resumeGasMetering();
    s_offRamp.execute(encodedMessage, ccvAddresses, verifierResults, 0);
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

    // Apply remote chain config updates on the each CCTP verifier.
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigArgs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigArgs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(s_sourceCCTPSetup.router),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_sourceCCTPSetup.verifier.applyRemoteChainConfigUpdates(remoteChainConfigArgs);

    remoteChainConfigArgs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigArgs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(s_destCCTPSetup.router),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_destCCTPSetup.verifier.applyRemoteChainConfigUpdates(remoteChainConfigArgs);

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

  // External so we can pass `bytes memory` as calldata (for MessageV1Codec._decodeMessageV1).
  function _decodeMessageV1(
    bytes calldata encodedMessage
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encodedMessage);
  }
}
