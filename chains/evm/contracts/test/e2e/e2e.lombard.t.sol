// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../Router.sol";
import {CommitteeVerifier} from "../../ccvs/CommitteeVerifier.sol";
import {LombardVerifier} from "../../ccvs/LombardVerifier.sol";
import {VersionedVerifierResolver} from "../../ccvs/VersionedVerifierResolver.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {IBridgeV1} from "../../interfaces/lombard/IBridgeV1.sol";
import {IBridgeV2} from "../../interfaces/lombard/IBridgeV2.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {AdvancedPoolHooks} from "../../pools/AdvancedPoolHooks.sol";
import {LombardTokenPool} from "../../pools/Lombard/LombardTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {OffRampHelper} from "../helpers/OffRampHelper.sol";
import {MockLombardBridge} from "../mocks/MockLombardBridge.sol";
import {MockLombardMailbox} from "../mocks/MockLombardMailbox.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";
import {OnRampSetup} from "../onRamp/OnRamp/OnRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract e2e_lombard is OnRampSetup {
  // Lombard version tag used by VersionedVerifierResolver inbound routing and LombardVerifier.verifyMessage parsing.
  bytes4 internal constant LOMBARD_VERSION_TAG_V1_7_0 = bytes4(keccak256("LombardVerifier 1.7.0"));
  bytes32 internal constant LOMBARD_CHAIN_ID = bytes32(uint256(10_000));

  OffRampHelper internal s_offRamp;
  MockLombardBridge internal s_lombardBridge;

  // Committee verifier behind a resolver proxy.
  address internal s_committeeCCV;
  CommitteeVerifier internal s_sourceCommitteeVerifier;
  MockVerifier internal s_destCommitteeVerifier;

  // Lombard CCV required by the Lombard pool: versioned resolver.
  address internal s_lombardCCV;
  LombardVerifier internal s_sourceLombardVerifier;
  LombardVerifier internal s_destLombardVerifier;

  // Lombard pools (source/dest).
  LombardTokenPool internal s_sourceLombardPool;
  LombardTokenPool internal s_destLombardPool;

  function setUp() public virtual override {
    super.setUp();

    // Ensure Router can resolve the onRamp address for the verifier allowlist check.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    // ================================================================
    // │                     CommitteeVerifier                        │
    // ================================================================

    // CommitteeVerifier behind VersionedVerifierResolver.
    s_sourceCommitteeVerifier = new CommitteeVerifier(
      CommitteeVerifier.DynamicConfig({feeAggregator: address(1), allowlistAdmin: address(0)}),
      new string[](0),
      address(s_mockRMNRemote)
    );

    BaseVerifier.RemoteChainConfigArgs[] memory destChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    destChainConfigs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: s_sourceRouter,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_sourceCommitteeVerifier.applyRemoteChainConfigUpdates(destChainConfigs);

    VersionedVerifierResolver srcCommitteeResolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.OutboundImplementationArgs[] memory outboundImpls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    outboundImpls[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      verifier: address(s_sourceCommitteeVerifier)
    });
    srcCommitteeResolver.applyOutboundImplementationUpdates(outboundImpls);

    // On destination we use a mock verifier to bypass committee signature requirements, but still keep the CCV address
    // as the committee resolver (default dest = committee).
    s_destCommitteeVerifier = new MockVerifier("");
    VersionedVerifierResolver.InboundImplementationArgs[] memory committeeInbound =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    committeeInbound[0] = VersionedVerifierResolver.InboundImplementationArgs({
      version: s_sourceCommitteeVerifier.versionTag(),
      verifier: address(s_destCommitteeVerifier)
    });
    srcCommitteeResolver.applyInboundImplementationUpdates(committeeInbound);

    s_committeeCCV = address(srcCommitteeResolver);

    // ================================================================
    // │                     Lombard Verifier                         │
    // ================================================================

    s_lombardBridge = new MockLombardBridge();

    s_sourceLombardVerifier =
      new LombardVerifier(IBridgeV2(address(s_lombardBridge)), new string[](0), address(s_mockRMNRemote));
    s_destLombardVerifier =
      new LombardVerifier(IBridgeV2(address(s_lombardBridge)), new string[](0), address(s_mockRMNRemote));

    s_sourceLombardVerifier.applyRemoteChainConfigUpdates(destChainConfigs);
    s_sourceLombardVerifier.setPath(
      DEST_CHAIN_SELECTOR, LOMBARD_CHAIN_ID, bytes32(bytes20(address(s_destLombardVerifier)))
    );

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: s_sourceFeeToken, localAdapter: address(0)});
    s_sourceLombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Configure destination Lombard verifier for verifyMessage (offRamp authorization check uses sourceChainSelector).
    BaseVerifier.RemoteChainConfigArgs[] memory lombardSourceConfig = new BaseVerifier.RemoteChainConfigArgs[](1);
    lombardSourceConfig[0] = BaseVerifier.RemoteChainConfigArgs({
      router: s_destRouter,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    s_destLombardVerifier.applyRemoteChainConfigUpdates(lombardSourceConfig);

    VersionedVerifierResolver lombardResolver = new VersionedVerifierResolver();
    VersionedVerifierResolver.OutboundImplementationArgs[] memory lombardOutbound =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    lombardOutbound[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      verifier: address(s_sourceLombardVerifier)
    });
    lombardResolver.applyOutboundImplementationUpdates(lombardOutbound);

    VersionedVerifierResolver.InboundImplementationArgs[] memory lombardInbound =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    lombardInbound[0] = VersionedVerifierResolver.InboundImplementationArgs({
      version: LOMBARD_VERSION_TAG_V1_7_0,
      verifier: address(s_destLombardVerifier)
    });
    lombardResolver.applyInboundImplementationUpdates(lombardInbound);

    s_lombardCCV = address(lombardResolver);

    // ================================================================
    // │                        Lombard pool                          │
    // ================================================================

    AdvancedPoolHooks hooks = new AdvancedPoolHooks(new address[](0), 0);
    AdvancedPoolHooks.CCVConfigArg[] memory ccvConfigs = new AdvancedPoolHooks.CCVConfigArg[](2);

    address[] memory required = new address[](2);
    required[0] = s_lombardCCV;
    required[1] = address(0); // This means "require the default CCV(s) for this lane".

    ccvConfigs[0] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: required,
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: new address[](0),
      inboundCCVsToAddAboveThreshold: new address[](0)
    });
    ccvConfigs[1] = AdvancedPoolHooks.CCVConfigArg({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      outboundCCVsToAddAboveThreshold: new address[](0),
      inboundCCVs: required,
      inboundCCVsToAddAboveThreshold: new address[](0)
    });
    hooks.applyCCVConfigUpdates(ccvConfigs);

    // Replace the default LINK pools with Lombard pools for this e2e test.
    s_sourceLombardPool = new LombardTokenPool(
      IERC20Metadata(s_sourceFeeToken),
      s_lombardCCV,
      IBridgeV1(address(s_lombardBridge)),
      address(0),
      address(hooks),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      DEFAULT_TOKEN_DECIMALS
    );
    s_destLombardPool = new LombardTokenPool(
      IERC20Metadata(s_destFeeToken),
      s_lombardCCV,
      IBridgeV1(address(s_lombardBridge)),
      address(0),
      address(hooks),
      address(s_mockRMNRemote),
      address(s_destRouter),
      DEFAULT_TOKEN_DECIMALS
    );

    // Update TokenSetup mappings used by the OnRampSetup helper functions.
    s_sourcePoolByToken[s_sourceFeeToken] = address(s_sourceLombardPool);
    s_destPoolByToken[s_destFeeToken] = address(s_destLombardPool);
    s_destPoolBySourceToken[s_sourceFeeToken] = address(s_destLombardPool);

    // Update pools in the registry + configure remote pool/token for both directions.
    if (!s_tokenAdminRegistry.isAdministrator(s_sourceFeeToken, OWNER)) {
      s_tokenAdminRegistry.proposeAdministrator(s_sourceFeeToken, OWNER);
      s_tokenAdminRegistry.acceptAdminRole(s_sourceFeeToken);
    }
    s_tokenAdminRegistry.setPool(s_sourceFeeToken, address(s_sourceLombardPool));
    if (!s_tokenAdminRegistry.isAdministrator(s_destFeeToken, OWNER)) {
      s_tokenAdminRegistry.proposeAdministrator(s_destFeeToken, OWNER);
      s_tokenAdminRegistry.acceptAdminRole(s_destFeeToken);
    }
    s_tokenAdminRegistry.setPool(s_destFeeToken, address(s_destLombardPool));

    {
      bytes[] memory remotePools = new bytes[](1);
      remotePools[0] = abi.encode(address(s_destLombardPool));

      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        remotePoolAddresses: remotePools,
        remoteTokenAddress: abi.encode(s_destFeeToken),
        outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: _getInboundRateLimiterConfig()
      });
      s_sourceLombardPool.applyChainUpdates(new uint64[](0), chainUpdates);
    }
    {
      bytes[] memory remotePools = new bytes[](1);
      remotePools[0] = abi.encode(address(s_sourceLombardPool));

      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        remotePoolAddresses: remotePools,
        remoteTokenAddress: abi.encode(s_sourceFeeToken),
        outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: _getInboundRateLimiterConfig()
      });
      s_destLombardPool.applyChainUpdates(new uint64[](0), chainUpdates);
    }

    // ================================================================
    // │                           OffRamp                            │
    // ================================================================

    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      })
    );

    // Update OnRamp config to point at this OffRamp.
    address[] memory defaultSourceCCVs = new address[](1);
    defaultSourceCCVs[0] = s_committeeCCV;

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultSourceCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRamp))
    });
    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

    // Configure OffRamp to require Committee verifier (resolver) by default for this source chain.
    address[] memory defaultDestCCVs = new address[](1);
    defaultDestCCVs[0] = s_committeeCCV;

    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encodePacked(s_onRamp);

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

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_destRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    // Seed existing fee recipients to avoid first-transfer cold init costs.
    deal(s_sourceFeeToken, s_committeeCCV, 1);
    deal(s_sourceFeeToken, s_lombardCCV, 1);
    deal(s_sourceFeeToken, s_defaultExecutor, 1);
  }

  function test_e2e_Lombard() public {
    vm.pauseGasMetering();
    uint64 expectedMsgNum = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber + 1;

    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), type(uint256).max);

    // No CCVs are specified as the pool required both the Lombard CCV and the default CCV.
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceFeeToken, amount: 1e18});

    (bytes32 messageId, bytes memory encodedMessage, OnRamp.Receipt[] memory receipts, bytes[] memory verifierBlobs) =
    _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      msgNum: expectedMsgNum,
      originalSender: OWNER
    });

    // Committee verifier returns versionTag (first verifier blob).
    verifierBlobs[0] = abi.encodePacked(s_sourceCommitteeVerifier.versionTag());

    // Lombard verifier returns a payload hash from MockLombardBridge.deposit (second verifier blob).
    // MockLombardBridge hashes (block.timestamp, optionalMessage) where optionalMessage = versionTag || messageId.
    bytes memory optionalMessage = bytes.concat(LOMBARD_VERSION_TAG_V1_7_0, messageId);
    verifierBlobs[1] = abi.encodePacked(keccak256(abi.encode(block.timestamp, optionalMessage)));

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

    address[] memory ccvAddresses = new address[](2);
    ccvAddresses[0] = s_committeeCCV;
    ccvAddresses[1] = s_lombardCCV;
    bytes[] memory verifierResults = new bytes[](2);
    verifierResults[0] = abi.encodePacked(s_sourceCommitteeVerifier.versionTag());

    bytes memory fakePayload = bytes("fake payload data");
    bytes memory fakeProof = bytes("fake signature data");

    verifierResults[1] = bytes.concat(
      LOMBARD_VERSION_TAG_V1_7_0,
      bytes2(uint16(fakePayload.length)),
      fakePayload,
      bytes2(uint16(fakeProof.length)),
      fakeProof
    );

    MessageV1Codec.MessageV1 memory messageV1 = this._decodeMessageV1(encodedMessage);

    vm.expectCall(
      address(s_destCommitteeVerifier),
      abi.encodeCall(CommitteeVerifier.verifyMessage, (messageV1, messageId, verifierResults[0]))
    );
    vm.expectCall(
      address(s_destLombardVerifier),
      abi.encodeCall(LombardVerifier.verifyMessage, (messageV1, messageId, verifierResults[1]))
    );
    vm.expectCall(
      s_lombardBridge.mailbox(), abi.encodeCall(MockLombardMailbox.deliverAndHandle, (fakePayload, fakeProof))
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
    s_offRamp.execute(encodedMessage, ccvAddresses, verifierResults);
  }

  // External so we can pass `bytes memory` as calldata (for MessageV1Codec._decodeMessageV1).
  function _decodeMessageV1(
    bytes calldata encodedMessage
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encodedMessage);
  }
}
