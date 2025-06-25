// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {INonceManager} from "../../../interfaces/INonceManager.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";

import {NonceManager} from "../../../NonceManager.sol";
import {Router} from "../../../Router.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropHelper} from "../../helpers/OffRampOverSuperchainInteropHelper.sol";

import {ICrossL2Inbox} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {BaseTest} from "../../BaseTest.t.sol";
import {TokenSetup} from "../../TokenSetup.t.sol";
import {MaybeRevertMessageReceiver} from "../../helpers/receivers/MaybeRevertMessageReceiver.sol";
import {MessageInterceptorHelper} from "../../helpers/MessageInterceptorHelper.sol";
import {MockCrossL2Inbox} from "../../helpers/MockCrossL2Inbox.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {WETH9} from "@chainlink/contracts/src/v0.8/vendor/canonical-weth/WETH9.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract OffRampOverSuperchainInteropSetup is BaseTest, TokenSetup {
  uint64 internal constant SOURCE_CHAIN_SELECTOR_1 = 5009297550715157269;
  uint64 internal constant SOURCE_CHAIN_SELECTOR_2 = 6433500567565415381;
  uint64 internal constant SOURCE_CHAIN_SELECTOR_3 = 4051577828743386545;

  uint256 internal constant SOURCE_CHAIN_ID_1 = 111;
  uint256 internal constant SOURCE_CHAIN_ID_2 = 222;
  uint256 internal constant SOURCE_CHAIN_ID_3 = 333;

  address internal constant ON_RAMP_ADDRESS_1 = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;
  address internal constant ON_RAMP_ADDRESS_2 = 0xaA3f843Cf8E33B1F02dd28303b6bD87B1aBF8AE4;
  address internal constant ON_RAMP_ADDRESS_3 = 0x71830C37Cb193e820de488Da111cfbFcC680a1b9;

  bytes internal constant ON_RAMP_ENCODED_1 = abi.encode(ON_RAMP_ADDRESS_1);
  bytes internal constant ON_RAMP_ENCODED_2 = abi.encode(ON_RAMP_ADDRESS_2);
  bytes internal constant ON_RAMP_ENCODED_3 = abi.encode(ON_RAMP_ADDRESS_3);

  uint32 internal constant PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS = 500;

  MockCrossL2Inbox internal s_mockCrossL2Inbox;
  OffRampOverSuperchainInterop internal s_offRamp;
  OffRampOverSuperchainInteropHelper internal s_offRampHelper;
  NonceManager internal s_nonceManager;
  Router internal s_router;
  MessageInterceptorHelper internal s_messageInterceptor;

  IAny2EVMMessageReceiver internal s_receiver;
  MaybeRevertMessageReceiver internal s_reverting_receiver;

  address[] internal s_allowedTransmitters;
  address internal s_transmitter1 = makeAddr("transmitter1");
  address internal s_transmitter2 = makeAddr("transmitter2");
  address internal s_transmitter3 = makeAddr("transmitter3");

  uint64 internal s_sequenceNumber;

  function setUp() public virtual override(BaseTest, TokenSetup) {
    BaseTest.setUp();
    TokenSetup.setUp();

    // Deploy mock CrossL2Inbox
    s_mockCrossL2Inbox = new MockCrossL2Inbox();

    // Deploy core contracts
    s_nonceManager = new NonceManager(new address[](0));
    s_router = new Router(address(new WETH9()), address(s_mockRMNRemote));
    s_messageInterceptor = new MessageInterceptorHelper();

    // Deploy receivers
    s_receiver = new MaybeRevertMessageReceiver(false);
    s_reverting_receiver = new MaybeRevertMessageReceiver(true);

    // Setup allowed transmitters
    s_allowedTransmitters = new address[](2);
    s_allowedTransmitters[0] = s_transmitter1;
    s_allowedTransmitters[1] = s_transmitter2;

    // Deploy OffRamp
    _deployOffRamp();
  }

  function _deployOffRamp() internal {
    OffRampOverSuperchainInterop.SourceChainConfigArgs[] memory sourceChainConfigs = 
      new OffRampOverSuperchainInterop.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = OffRampOverSuperchainInterop.SourceChainConfigArgs({
      router: s_router,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      isEnabled: true,
      onRamp: ON_RAMP_ENCODED_1
    });

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainIdConfigs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    chainIdConfigs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: SOURCE_CHAIN_ID_1
    });

    // Deploy both regular and helper versions
    s_offRamp = new OffRampOverSuperchainInterop(
      OffRampOverSuperchainInterop.StaticConfig({
        chainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        crossL2Inbox: s_mockCrossL2Inbox,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        nonceManager: address(s_nonceManager)
      }),
      OffRampOverSuperchainInterop.DynamicConfig({
        feeQuoter: makeAddr("mockFeeQuoter"),
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS,
        messageInterceptor: address(0)
      }),
      sourceChainConfigs,
      s_allowedTransmitters,
      chainIdConfigs
    );

    s_offRampHelper = new OffRampOverSuperchainInteropHelper(
      OffRampOverSuperchainInterop.StaticConfig({
        chainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        crossL2Inbox: s_mockCrossL2Inbox,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        nonceManager: address(s_nonceManager)
      }),
      OffRampOverSuperchainInterop.DynamicConfig({
        feeQuoter: makeAddr("mockFeeQuoter"),
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS,
        messageInterceptor: address(0)
      }),
      sourceChainConfigs,
      s_allowedTransmitters,
      chainIdConfigs
    );

    // Authorize offRamp in NonceManager
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_offRamp);
    // Already pranked as OWNER in BaseTest.setUp()
    s_nonceManager.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({
        addedCallers: authorizedCallers,
        removedCallers: new address[](0)
      })
    );

    // Setup router
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      offRamp: address(s_offRamp)
    });
    s_router.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);
  }

  function _generateAny2EVMMessage(
    uint64 sourceChainSelector,
    address, // onRamp
    uint64 sequenceNumber,
    uint64 nonce,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bool hasData
  ) internal view returns (Internal.Any2EVMRampMessage memory) {
    bytes memory data = hasData ? abi.encode("test data") : bytes("");
    bytes32 messageId = keccak256(abi.encode(sourceChainSelector, sequenceNumber, nonce));

    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: messageId,
        sourceChainSelector: sourceChainSelector,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: sequenceNumber,
        nonce: nonce
      }),
      sender: abi.encode(OWNER),
      data: data,
      receiver: address(s_receiver),
      gasLimit: 200_000,
      tokenAmounts: _convertToInternal(tokenAmounts)
    });
  }

  function _convertToInternal(
    Client.EVMTokenAmount[] memory tokenAmounts
  ) internal pure returns (Internal.Any2EVMTokenTransfer[] memory) {
    Internal.Any2EVMTokenTransfer[] memory internal_ = new Internal.Any2EVMTokenTransfer[](tokenAmounts.length);
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      internal_[i] = Internal.Any2EVMTokenTransfer({
        sourcePoolAddress: abi.encode(address(0)),
        destTokenAddress: tokenAmounts[i].token,
        extraData: "",
        amount: tokenAmounts[i].amount,
        destGasAmount: 50_000
      });
    }
    return internal_;
  }

  function _createIdentifier(
    address origin,
    uint256 blockNumber,
    uint256 logIndex,
    uint256 timestamp,
    uint256 chainId
  ) internal pure returns (Identifier memory) {
    return Identifier({
      origin: origin,
      blockNumber: blockNumber,
      logIndex: logIndex,
      timestamp: timestamp,
      chainId: chainId
    });
  }

  function _encodeLogData(
    uint64 destChainSelector,
    uint64 sequenceNumber,
    Internal.Any2EVMRampMessage memory message
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      SuperchainInterop.SENT_MESSAGE_LOG_SELECTOR,
      abi.encode(destChainSelector, sequenceNumber),
      abi.encode(message)
    );
  }

  function _createExecutionReport(
    Internal.Any2EVMRampMessage memory message,
    Identifier memory identifier,
    bytes[] memory offchainTokenData
  ) internal pure returns (SuperchainInterop.ExecutionReport memory) {
    bytes memory logData = _encodeLogData(
      message.header.destChainSelector,
      message.header.sequenceNumber,
      message
    );

    return SuperchainInterop.ExecutionReport({
      logData: logData,
      identifier: identifier,
      offchainTokenData: offchainTokenData
    });
  }

  function _setupMultipleOffRamps() internal {
    OffRampOverSuperchainInterop.SourceChainConfigArgs[] memory sourceChainConfigs = 
      new OffRampOverSuperchainInterop.SourceChainConfigArgs[](3);
    sourceChainConfigs[0] = OffRampOverSuperchainInterop.SourceChainConfigArgs({
      router: s_router,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      isEnabled: true,
      onRamp: ON_RAMP_ENCODED_1
    });
    sourceChainConfigs[1] = OffRampOverSuperchainInterop.SourceChainConfigArgs({
      router: s_router,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_2,
      isEnabled: false,
      onRamp: ON_RAMP_ENCODED_2
    });
    sourceChainConfigs[2] = OffRampOverSuperchainInterop.SourceChainConfigArgs({
      router: s_router,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_3,
      isEnabled: true,
      onRamp: ON_RAMP_ENCODED_3
    });

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainIdConfigs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](3);
    chainIdConfigs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: SOURCE_CHAIN_ID_1
    });
    chainIdConfigs[1] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: SOURCE_CHAIN_ID_2
    });
    chainIdConfigs[2] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_3,
      chainId: SOURCE_CHAIN_ID_3
    });

    // Already pranked as OWNER in BaseTest.setUp()
    s_offRamp.applySourceChainConfigUpdates(sourceChainConfigs);
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainIdConfigs);
  }

  function _setMockCrossL2InboxValidMessage(
    Internal.Any2EVMRampMessage memory message,
    address // onRamp
  ) internal {
    bytes memory logData = _encodeLogData(
      message.header.destChainSelector,
      message.header.sequenceNumber,
      message
    );
    bytes32 msgHash = keccak256(logData);
    s_mockCrossL2Inbox.setValidMessage(msgHash, true);
  }
}