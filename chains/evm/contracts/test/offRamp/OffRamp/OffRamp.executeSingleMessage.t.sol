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

  function test_executeSingleMessage_PassActualTokenAmountToReceiver() public {
    uint256 amount = 100;
    address pool = makeAddr("pool");
    address sourceToken = makeAddr("sourceToken");
    address destToken = address(new BurnMintERC20("destToken", "destToken", 18, 0, 0));
    address tokenReceiver = makeAddr("tokenReceiver");

    (
      MessageV1Codec.MessageV1 memory message,
      bytes32 messageId,
      address[] memory ccvs,
      bytes[] memory verifierResults,
      address verifierImpl
    ) = _setupMessageWithTokenTransfer(pool, sourceToken, destToken, tokenReceiver, amount, "test data", 200_000);

    _mockVerifierCalls(message, messageId, verifierResults, verifierImpl);
    _mockPoolCalls(pool, tokenReceiver, destToken, message, amount, amount);

    // Expect the routeMessage call with amount = 0 because the balance of the receiver didn't actually change.
    // releaseOrMint was mocked and did not actually update token state.
    // Since the receiver is not the token pool, the actual balance diff should be passed to the receiver.
    _mockReceiverCalls(address(bytes20(message.receiver)));
    _expectRouteMessage(message, messageId, destToken, 0);

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_PassAmountReturnedByPoolToReceiver() public {
    address pool = makeAddr("pool");
    uint256 amount = 100;
    address sourceToken = makeAddr("sourceToken");
    address destToken = address(new BurnMintERC20("destToken", "destToken", 18, 0, 0));

    (
      MessageV1Codec.MessageV1 memory message,
      bytes32 messageId,
      address[] memory ccvs,
      bytes[] memory verifierResults,
      address verifierImpl
    ) = _setupMessageWithTokenTransfer(pool, sourceToken, destToken, pool, amount, "test data", 200_000);

    _mockVerifierCalls(message, messageId, verifierResults, verifierImpl);
    // Since the pool is the receiver, we expect the amount returned by the pool to be passed to the receiver.
    Pool.ReleaseOrMintOutV1 memory releaseOrMintOut = _mockPoolCalls(pool, pool, destToken, message, amount, amount);

    _mockReceiverCalls(address(bytes20(message.receiver)));
    _expectRouteMessage(message, messageId, destToken, releaseOrMintOut.destinationAmount);

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
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
    // Make this NOT a token-only transfer so receiver CCVs are queried.
    message.ccipReceiveGasLimit = 100_000;

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
    uint256 amount = 100;

    MessageV1Codec.MessageV1 memory message = _getMessage();

    // Modify message with token transfer.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(pool),
      sourceTokenAddress: abi.encode(sourceToken),
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
        IPoolV2.getRequiredCCVs, (token, SOURCE_CHAIN_SELECTOR, amount, 0, "", IPoolV2.MessageDirection.Inbound)
      ),
      abi.encode(_arrayOf(poolRequiredCCV))
    );
    // Mock token receiver balance check.
    vm.mockCall(address(token), abi.encodeWithSelector(IERC20.balanceOf.selector, tokenReceiver), abi.encode(amount));

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
      defaultCCVs: _arrayOf(s_defaultCCV),
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

  function test_executeSingleMessage_RevertWhen_InvalidDestTokenAddressLength() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes memory invalidDestToken = new bytes(19); // not 20 bytes

    MessageV1Codec.TokenTransferV1[] memory transfers = new MessageV1Codec.TokenTransferV1[](1);
    transfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1,
      sourcePoolAddress: abi.encode(makeAddr("pool")),
      sourceTokenAddress: abi.encode(makeAddr("sourceToken")),
      destTokenAddress: invalidDestToken,
      tokenReceiver: abi.encodePacked(makeAddr("receiver")),
      extraData: ""
    });
    message.tokenTransfer = transfers;

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidDestToken));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_InvalidTokenReceiverLength() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    bytes memory invalidTokenReceiver = new bytes(21); // not 20 bytes

    MessageV1Codec.TokenTransferV1[] memory transfers = new MessageV1Codec.TokenTransferV1[](1);
    transfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1,
      sourcePoolAddress: abi.encode(makeAddr("pool")),
      sourceTokenAddress: abi.encode(makeAddr("sourceToken")),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: invalidTokenReceiver,
      extraData: ""
    });
    message.tokenTransfer = transfers;

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    address[] memory ccvs = _arrayOf(s_defaultCCV);
    bytes[] memory verifierResults = new bytes[](1);

    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidTokenReceiver));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_OptionalCCVQuorumNotReached() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    address receiver = address(bytes20(message.receiver));

    // Make this NOT a token-only transfer so receiver CCVs are queried.
    message.ccipReceiveGasLimit = 100_000;

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
    uint256 amount = 100;

    // Create message with multiple token transfers (invalid) - but we need to manually set this
    // since encoding would fail. We'll create a valid single token transfer first, then modify the array.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(makeAddr("pool1")),
      sourceTokenAddress: abi.encode(makeAddr("sourceToken1")),
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
    vm.mockCall(destToken, abi.encodeWithSelector(IERC20.balanceOf.selector, tokenReceiver), abi.encode(amount));

    s_offRamp.executeSingleMessage(message, messageId, ccvs, verifierResults);
  }

  function test_executeSingleMessage_RevertWhen_InvalidEVMAddress_TokenAddress() public {
    MessageV1Codec.MessageV1 memory message = _getMessage();
    uint256 amount = 100;

    // Create message with invalid token address length.
    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(makeAddr("pool")),
      sourceTokenAddress: abi.encode(makeAddr("sourceToken")),
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

  function _getMessage() internal returns (MessageV1Codec.MessageV1 memory message) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: 200_000,
      ccipReceiveGasLimit: 0,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encode(makeAddr("onRamp")),
      offRampAddress: abi.encodePacked(makeAddr("offRamp")),
      sender: abi.encode(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  function _setupMessageWithTokenTransfer(
    address pool,
    address sourceToken,
    address destToken,
    address tokenReceiver,
    uint256 amount,
    string memory data,
    uint32 ccipReceiveGasLimit
  )
    internal
    returns (
      MessageV1Codec.MessageV1 memory message,
      bytes32 messageId,
      address[] memory ccvs,
      bytes[] memory verifierResults,
      address verifierImpl
    )
  {
    message = _getMessage();
    ccvs = _arrayOf(s_defaultCCV);
    verifierResults = new bytes[](1);
    verifierImpl = makeAddr("verifierImpl");

    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(pool),
      sourceTokenAddress: abi.encode(sourceToken),
      destTokenAddress: abi.encodePacked(destToken),
      tokenReceiver: abi.encodePacked(tokenReceiver),
      extraData: ""
    });
    message.tokenTransfer = tokenAmounts;
    message.data = bytes(data);
    message.ccipReceiveGasLimit = ccipReceiveGasLimit;
    messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    return (message, messageId, ccvs, verifierResults, verifierImpl);
  }

  function _mockVerifierCalls(
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageId,
    bytes[] memory verifierResults,
    address verifierImpl
  ) internal {
    vm.mockCall(
      s_defaultCCV,
      abi.encodeCall(ICrossChainVerifierResolver.getInboundImplementation, (verifierResults[0])),
      abi.encode(verifierImpl)
    );

    vm.mockCall(
      verifierImpl,
      abi.encodeCall(ICrossChainVerifierV1.verifyMessage, (message, messageId, verifierResults[0])),
      abi.encode("")
    );
  }

  function _mockPoolCalls(
    address pool,
    address tokenReceiver,
    address destToken,
    MessageV1Codec.MessageV1 memory message,
    uint256 amount,
    uint256 destinationAmount
  ) internal returns (Pool.ReleaseOrMintOutV1 memory) {
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (destToken)), abi.encode(pool));

    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    vm.mockCall(
      pool,
      abi.encodeCall(
        IPoolV2.getRequiredCCVs, (destToken, SOURCE_CHAIN_SELECTOR, amount, 0, "", IPoolV2.MessageDirection.Inbound)
      ),
      abi.encode(new address[](0))
    );

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: message.sender,
      receiver: tokenReceiver,
      sourceDenominatedAmount: message.tokenTransfer[0].amount,
      localToken: destToken,
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(pool),
      sourcePoolData: message.tokenTransfer[0].extraData,
      offchainTokenData: ""
    });

    Pool.ReleaseOrMintOutV1 memory releaseOrMintOut = Pool.ReleaseOrMintOutV1({destinationAmount: destinationAmount});

    vm.mockCall(pool, abi.encodeCall(IPoolV2.releaseOrMint, (releaseOrMintIn, 0)), abi.encode(releaseOrMintOut));

    return releaseOrMintOut;
  }

  function _mockReceiverCalls(
    address receiver
  ) internal {
    vm.mockCall(receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      receiver, abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)), abi.encode(true)
    );
  }

  function _expectRouteMessage(
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageId,
    address destToken,
    uint256 tokenAmount
  ) internal {
    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](1);
    destTokenAmounts[0] = Client.EVMTokenAmount({token: destToken, amount: tokenAmount});

    Client.Any2EVMMessage memory any2EVMMessage = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: message.sender,
      data: message.data,
      destTokenAmounts: destTokenAmounts
    });

    vm.expectCall(
      address(s_sourceRouter),
      abi.encodeCall(
        IRouter.routeMessage,
        (any2EVMMessage, GAS_FOR_CALL_EXACT_CHECK, message.ccipReceiveGasLimit, address(bytes20(message.receiver)))
      )
    );
  }
}
