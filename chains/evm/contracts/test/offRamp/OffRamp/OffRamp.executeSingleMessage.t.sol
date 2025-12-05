// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// forge-config: default.allow_internal_expect_revert = true

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract OffRamp_executeSingleMessage is OffRampSetup {
  struct TokenAddresses {
    address pool;
    address sourceToken;
    address destToken;
    address tokenReceiver;
  }

  struct MessageSetup {
    MessageV1Codec.MessageV1 message;
    bytes32 messageId;
    address[] ccvs;
    bytes[] verifierResults;
    address verifierImpl;
  }

  struct PoolParams {
    Pool.ReleaseOrMintOutV1 releaseOrMintOut;
    Pool.ReleaseOrMintInV1 releaseOrMintIn;
  }

  struct RouterParams {
    Client.EVMTokenAmount[] destTokenAmounts;
    Client.Any2EVMMessage any2EVMMessage;
  }

  function setUp() public virtual override {
    super.setUp();

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Apply off ramp updates on the receiver.
    Router.OffRamp[] memory offRamps = new Router.OffRamp[](1);
    offRamps[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRamps);

    // Mock validateReport for default message structure.
    bytes32 messageHash = keccak256(MessageV1Codec._encodeMessageV1(message));

    bytes memory defaultVerifierResults = abi.encode("mock verifier results");
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageHash, defaultVerifierResults)),
      abi.encode(true)
    );

    vm.startPrank(address(s_offRamp));
  }

  function _getMessage() internal returns (MessageV1Codec.MessageV1 memory message) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: 200_000,
      ccipReceiveGasLimit: 0,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encodePacked(makeAddr("onRamp")),
      offRampAddress: abi.encodePacked(makeAddr("offRamp")),
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  function test_executeSingleMessage_PassActualTokenAmountToReceiver() public {
    TokenAddresses memory tokenAddresses = TokenAddresses({
      pool: makeAddr("pool"),
      sourceToken: makeAddr("sourceToken"),
      destToken: address(new BurnMintERC20("destToken", "destToken", 18, 0, 0)),
      tokenReceiver: makeAddr("tokenReceiver")
    });

    MessageSetup memory msgSetup = MessageSetup({
      message: _getMessage(),
      messageId: bytes32(0), // Will be set after token transfer
      ccvs: _arrayOf(s_defaultCCV),
      verifierResults: new bytes[](1),
      verifierImpl: makeAddr("verifierImpl")
    });

    // Create token transfer.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(tokenAddresses.pool),
      sourceTokenAddress: abi.encodePacked(tokenAddresses.sourceToken),
      destTokenAddress: abi.encodePacked(tokenAddresses.destToken),
      tokenReceiver: abi.encodePacked(tokenAddresses.tokenReceiver),
      extraData: ""
    });
    msgSetup.message.tokenTransfer = tokenAmounts;
    msgSetup.message.data = "test data";
    msgSetup.message.ccipReceiveGasLimit = 200000;
    msgSetup.messageId = keccak256(MessageV1Codec._encodeMessageV1(msgSetup.message));

    // Mock CCV resolver to return address(0) for inbound implementation.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierResolver.getInboundImplementation, (msgSetup.verifierResults[0])),
      abi.encode(msgSetup.verifierImpl)
    );

    // Mock getRequiredCCVs on token pool.
    vm.mockCall(
      tokenAddresses.pool,
      abi.encodeCall(
        IPoolV2.getRequiredCCVs,
        (tokenAddresses.destToken, SOURCE_CHAIN_SELECTOR, 100, 0, "", IPoolV2.MessageDirection.Inbound)
      ),
      abi.encode(new address[](0))
    );

    // Mock verifyMessage call.
    vm.mockCall(
      msgSetup.verifierImpl,
      abi.encodeCall(
        ICrossChainVerifierV1.verifyMessage, (msgSetup.message, msgSetup.messageId, msgSetup.verifierResults[0])
      ),
      abi.encode("")
    );

    // Mock token admin registry to return the pool.
    vm.mockCall(
      s_tokenAdminRegistry,
      abi.encodeCall(ITokenAdminRegistry.getPool, (tokenAddresses.destToken)),
      abi.encode(tokenAddresses.pool)
    );

    // Mock supportsInterface calls. ERC165Checker checks IERC165 first, then the specific interface.
    vm.mockCall(
      tokenAddresses.pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true)
    );
    vm.mockCall(
      tokenAddresses.pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true)
    );

    // Mock releaseOrMint call.
    PoolParams memory poolParams = PoolParams({
      releaseOrMintOut: Pool.ReleaseOrMintOutV1({destinationAmount: 100}),
      releaseOrMintIn: Pool.ReleaseOrMintInV1({
        originalSender: msgSetup.message.sender,
        receiver: tokenAddresses.tokenReceiver,
        sourceDenominatedAmount: msgSetup.message.tokenTransfer[0].amount,
        localToken: tokenAddresses.destToken,
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(tokenAddresses.pool),
        sourcePoolData: msgSetup.message.tokenTransfer[0].extraData,
        offchainTokenData: ""
      })
    });
    vm.mockCall(
      tokenAddresses.pool,
      abi.encodeCall(IPoolV2.releaseOrMint, (poolParams.releaseOrMintIn, 0)),
      abi.encode(poolParams.releaseOrMintOut)
    );

    // Expect that the call to the receiver is correct.
    RouterParams memory routerParams;
    routerParams.destTokenAmounts = new Client.EVMTokenAmount[](1);
    // Expect the routeMessage call with amount = 0 because the balance of the receiver didn't actually change.
    // releaseOrMint was mocked and did not actually update token state.
    // Since the receiver is not the token pool, the actual balance diff should be passed to the receiver.
    routerParams.destTokenAmounts[0] = Client.EVMTokenAmount({token: tokenAddresses.destToken, amount: 0});
    routerParams.any2EVMMessage = Client.Any2EVMMessage({
      messageId: msgSetup.messageId,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: msgSetup.message.sender,
      data: msgSetup.message.data,
      destTokenAmounts: routerParams.destTokenAmounts
    });

    // Mock the supportsInterface check on the receiver.
    vm.mockCall(
      address(bytes20(msgSetup.message.receiver)),
      abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)),
      abi.encode(true)
    );
    vm.mockCall(
      address(bytes20(msgSetup.message.receiver)),
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );
    vm.expectCall(
      address(s_sourceRouter),
      abi.encodeCall(
        IRouter.routeMessage,
        (
          routerParams.any2EVMMessage,
          GAS_FOR_CALL_EXACT_CHECK,
          msgSetup.message.ccipReceiveGasLimit,
          address(bytes20(msgSetup.message.receiver))
        )
      )
    );
    s_offRamp.executeSingleMessage(msgSetup.message, msgSetup.messageId, msgSetup.ccvs, msgSetup.verifierResults);
  }

  function test_executeSingleMessage_PassAmountReturnedByPoolToReceiver() public {
    address pool = makeAddr("pool");

    TokenAddresses memory tokenAddresses = TokenAddresses({
      pool: pool,
      sourceToken: makeAddr("sourceToken"),
      destToken: address(new BurnMintERC20("destToken", "destToken", 18, 0, 0)),
      tokenReceiver: pool
    });

    MessageSetup memory msgSetup = MessageSetup({
      message: _getMessage(),
      messageId: bytes32(0), // Will be set after token transfer
      ccvs: _arrayOf(s_defaultCCV),
      verifierResults: new bytes[](1),
      verifierImpl: makeAddr("verifierImpl")
    });

    // Create token transfer.
    // Since the pool is the receiver, we expect the amount returned by the pool to be passed to the receiver.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(tokenAddresses.pool),
      sourceTokenAddress: abi.encodePacked(tokenAddresses.sourceToken),
      destTokenAddress: abi.encodePacked(tokenAddresses.destToken),
      tokenReceiver: abi.encodePacked(tokenAddresses.tokenReceiver),
      extraData: ""
    });
    msgSetup.message.tokenTransfer = tokenAmounts;
    msgSetup.message.data = "test data";
    msgSetup.message.ccipReceiveGasLimit = 200000;
    msgSetup.messageId = keccak256(MessageV1Codec._encodeMessageV1(msgSetup.message));

    // Mock CCV resolver to return address(0) for inbound implementation.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierResolver.getInboundImplementation, (msgSetup.verifierResults[0])),
      abi.encode(msgSetup.verifierImpl)
    );

    // Mock getRequiredCCVs on token pool.
    vm.mockCall(
      tokenAddresses.pool,
      abi.encodeCall(
        IPoolV2.getRequiredCCVs,
        (tokenAddresses.destToken, SOURCE_CHAIN_SELECTOR, 100, 0, "", IPoolV2.MessageDirection.Inbound)
      ),
      abi.encode(new address[](0))
    );

    // Mock verifyMessage call.
    vm.mockCall(
      msgSetup.verifierImpl,
      abi.encodeCall(
        ICrossChainVerifierV1.verifyMessage, (msgSetup.message, msgSetup.messageId, msgSetup.verifierResults[0])
      ),
      abi.encode("")
    );

    // Mock token admin registry to return the pool.
    vm.mockCall(
      s_tokenAdminRegistry,
      abi.encodeCall(ITokenAdminRegistry.getPool, (tokenAddresses.destToken)),
      abi.encode(tokenAddresses.pool)
    );

    // Mock supportsInterface calls. ERC165Checker checks IERC165 first, then the specific interface.
    vm.mockCall(
      tokenAddresses.pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true)
    );
    vm.mockCall(
      tokenAddresses.pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true)
    );

    // Mock releaseOrMint call.
    PoolParams memory poolParams = PoolParams({
      releaseOrMintOut: Pool.ReleaseOrMintOutV1({destinationAmount: 100}),
      releaseOrMintIn: Pool.ReleaseOrMintInV1({
        originalSender: msgSetup.message.sender,
        receiver: tokenAddresses.tokenReceiver,
        sourceDenominatedAmount: msgSetup.message.tokenTransfer[0].amount,
        localToken: tokenAddresses.destToken,
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(tokenAddresses.pool),
        sourcePoolData: msgSetup.message.tokenTransfer[0].extraData,
        offchainTokenData: ""
      })
    });
    vm.mockCall(
      tokenAddresses.pool,
      abi.encodeCall(IPoolV2.releaseOrMint, (poolParams.releaseOrMintIn, 0)),
      abi.encode(poolParams.releaseOrMintOut)
    );

    // Expect that the call to the receiver is correct.
    RouterParams memory routerParams;
    routerParams.destTokenAmounts = new Client.EVMTokenAmount[](1);
    routerParams.destTokenAmounts[0] =
      Client.EVMTokenAmount({token: tokenAddresses.destToken, amount: poolParams.releaseOrMintOut.destinationAmount});
    routerParams.any2EVMMessage = Client.Any2EVMMessage({
      messageId: msgSetup.messageId,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: msgSetup.message.sender,
      data: msgSetup.message.data,
      destTokenAmounts: routerParams.destTokenAmounts
    });

    // Mock the supportsInterface check on the receiver.
    vm.mockCall(
      address(bytes20(msgSetup.message.receiver)),
      abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)),
      abi.encode(true)
    );
    vm.mockCall(
      address(bytes20(msgSetup.message.receiver)),
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );

    vm.expectCall(
      address(s_sourceRouter),
      abi.encodeCall(
        IRouter.routeMessage,
        (
          routerParams.any2EVMMessage,
          GAS_FOR_CALL_EXACT_CHECK,
          msgSetup.message.ccipReceiveGasLimit,
          address(bytes20(msgSetup.message.receiver))
        )
      )
    );
    s_offRamp.executeSingleMessage(msgSetup.message, msgSetup.messageId, msgSetup.ccvs, msgSetup.verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_CanOnlySelfCall() public {
    vm.stopPrank();
    MessageV1Codec.MessageV1 memory message;

    vm.expectRevert(OffRamp.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(message, bytes32(0), new address[](0), new bytes[](0));
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_ReceiverCCV() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address receiver = makeAddr("receiver");
    address requiredCCV = makeAddr("requiredCCV");

    message.receiver = abi.encodePacked(receiver);

    // Set up receiver to require a specific CCV.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, _arrayOf(requiredCCV), new address[](0), 0);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Keep default CCV, but don't include the required CCV.
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.RequiredCCVMissing.selector, requiredCCV));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_PoolCCV() public {
    address poolRequiredCCV = makeAddr("poolRequiredCCV");
    address sourceToken = makeAddr("sourceToken");
    address token = makeAddr("token");
    address pool = makeAddr("pool");
    address tokenReceiver = makeAddr("tokenReceiver");
    uint256 tokenAmount = 100;

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with token transfer.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: tokenAmount,
      sourcePoolAddress: abi.encodePacked(pool),
      sourceTokenAddress: abi.encodePacked(sourceToken),
      destTokenAddress: abi.encodePacked(token),
      tokenReceiver: abi.encodePacked(tokenReceiver),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Keep default CCV, but don't include the pool required CCV.
    bytes[] memory verifierResults = new bytes[](1);

    // Mock token admin registry to return the pool.
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (token)), abi.encode(pool));

    // Mock pool supportsInterface for IPoolV2.
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool to require a specific CCV.
    vm.mockCall(
      pool,
      abi.encodeCall(
        IPoolV2.getRequiredCCVs, (token, SOURCE_CHAIN_SELECTOR, tokenAmount, 0, "", IPoolV2.MessageDirection.Inbound)
      ),
      abi.encode(_arrayOf(poolRequiredCCV))
    );
    // Mock token receiver balance check.
    vm.mockCall(
      address(token), abi.encodeWithSelector(IERC20.balanceOf.selector, tokenReceiver), abi.encode(tokenAmount)
    );

    vm.expectRevert(abi.encodeWithSelector(OffRamp.RequiredCCVMissing.selector, poolRequiredCCV));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_RequiredCCVMissing_LaneMandatedCCV() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address laneMandatedCCV = makeAddr("laneMandatedCCV");

    // Configure source chain with lane mandated CCV.
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encode(makeAddr("onRamp"));

    OffRamp.SourceChainConfigArgs[] memory sourceChainConfigArgs = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainConfigArgs[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCV: _arrayOf(s_defaultCCV),
      laneMandatedCCVs: _arrayOf(laneMandatedCCV)
    });

    vm.stopPrank();
    vm.startPrank(OWNER);
    s_offRamp.applySourceChainConfigUpdates(sourceChainConfigArgs);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Report doesn't include the lane mandated CCV.
    bytes[] memory verifierResults = new bytes[](1);

    vm.startPrank(address(s_offRamp));

    vm.expectRevert(abi.encodeWithSelector(OffRamp.RequiredCCVMissing.selector, laneMandatedCCV));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_OptionalCCVQuorumNotReached() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    address receiver = address(bytes20(message.receiver));

    address optionalCCV2 = makeAddr("optionalCCV2");
    address[] memory optionalCCVs = new address[](2);
    optionalCCVs[0] = s_defaultCCV; // This will be found in the report.
    optionalCCVs[1] = optionalCCV2; // This won't be found.

    uint8 optionalThreshold = 2; // Need 2 optional CCVs.

    // Set up receiver to return optional CCVs with threshold 2.
    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, new address[](0), optionalCCVs, optionalThreshold);

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV); // Report only includes one CCV, but threshold requires 2.
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.OptionalCCVQuorumNotReached.selector, optionalThreshold, ccvs.length)
    );

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_CCVValidationFails() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);
    bytes memory revertReason = "CCV validation failed";

    // Mock CCV validateReport to fail/revert.
    vm.mockCall(
      s_defaultCCV, abi.encodeCall(ICrossChainVerifierResolver.getInboundImplementation, ""), abi.encode(s_defaultCCV)
    );
    vm.mockCallRevert(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      revertReason
    );

    vm.expectRevert(revertReason);

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_InvalidNumberOfTokens() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    address destToken = makeAddr("destToken");
    address tokenReceiver = makeAddr("tokenReceiver");

    // Create message with multiple token transfers (invalid) - but we need to manually set this
    // since encoding would fail. We'll create a valid single token transfer first, then modify the array.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool1")),
      sourceTokenAddress: abi.encodePacked(makeAddr("sourceToken1")),
      destTokenAddress: abi.encodePacked(destToken),
      tokenReceiver: abi.encodePacked(tokenReceiver),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    // Encode the message with single token transfer first.
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    // Now manually create the invalid message with 2 token transfers for the executeSingleMessage call.
    MessageV1Codec.TokenTransferV1[] memory invalidTokenAmounts = new MessageV1Codec.TokenTransferV1[](2);
    invalidTokenAmounts[0] = tokenAmounts[0];
    invalidTokenAmounts[1] = tokenAmounts[0];
    message.tokenTransfer = invalidTokenAmounts;

    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidNumberOfTokens.selector, invalidTokenAmounts.length));
    vm.mockCall(destToken, abi.encodeWithSelector(IERC20.balanceOf.selector, tokenReceiver), abi.encode(100));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_InvalidEVMAddress_TokenAddress() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Create message with invalid token address length.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("sourceToken")),
      destTokenAddress: hex"1234", // Invalid length (not 20 bytes).
      tokenReceiver: abi.encodePacked(makeAddr("tokenReceiver")),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, tokenAmounts[0].destTokenAddress));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_InboundImplementationNotFound() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);

    // Mock CCV resolver to return address(0) for inbound implementation.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierResolver.getInboundImplementation, (verifierResults[0])),
      abi.encode(address(0))
    );

    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InboundImplementationNotFound.selector, s_defaultCCV, verifierResults[0])
    );

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }
}
