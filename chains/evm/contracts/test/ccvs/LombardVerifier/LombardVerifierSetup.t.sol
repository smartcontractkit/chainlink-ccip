// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";
import {IMailbox} from "../../../interfaces/lombard/IMailbox.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract MockLombardBridge is IBridgeV2 {
  address public s_mailbox;
  bytes32 public s_lastPayloadHash;

  constructor() {
    s_mailbox = address(new MockLombardMailbox());
  }

  function mailbox() external view override returns (address) {
    return s_mailbox;
  }

  function MSG_VERSION() external pure override returns (uint8) {
    return 1;
  }

  function deposit(
    bytes32,
    address,
    address,
    bytes32,
    uint256,
    bytes32,
    bytes calldata optionalMessage
  ) external payable override returns (uint256, bytes32) {
    s_lastPayloadHash = keccak256(abi.encode(block.timestamp, optionalMessage));

    MockLombardMailbox(s_mailbox).setMessageId(optionalMessage);

    return (0, s_lastPayloadHash);
  }
}

contract MockLombardMailbox is IMailbox {
  bool public s_shouldSucceed = true;
  bytes internal s_optionalMessage = abi.encode(bytes32(0));

  function setMessageId(
    bytes calldata optionalMessage
  ) external {
    s_optionalMessage = optionalMessage;
  }

  function setShouldSucceed(
    bool shouldSucceed
  ) external {
    s_shouldSucceed = shouldSucceed;
  }

  function deliverAndHandle(
    bytes calldata,
    bytes calldata
  ) external view override returns (bytes32, bool, bytes memory) {
    return (bytes32(0), s_shouldSucceed, s_optionalMessage);
  }
}

contract LombardVerifierSetup is BaseVerifierSetup {
  LombardVerifier internal s_lombardVerifier;
  MockLombardBridge internal s_mockBridge;
  MockLombardMailbox internal s_mockMailbox;
  BurnMintERC20 internal s_testToken;

  bytes32 internal constant LOMBARD_CHAIN_ID = bytes32(uint256(10000));
  bytes32 internal constant ALLOWED_CALLER = bytes32(uint256(0x123456));
  uint256 internal constant TRANSFER_AMOUNT = 1e18;

  function setUp() public virtual override {
    super.setUp();

    // Deploy mock bridge and get its mailbox.
    s_mockBridge = new MockLombardBridge();
    s_mockMailbox = MockLombardMailbox(s_mockBridge.s_mailbox());

    // Deploy verifier.
    s_lombardVerifier = new LombardVerifier(IBridgeV2(address(s_mockBridge)), STORAGE_LOCATION);

    // Deploy test token and add it as a supported token.
    s_testToken = new BurnMintERC20("Test Token", "TEST", 18, 0, 0);
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(s_testToken), localAdapter: address(0)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Set up remote chain config with the router.
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    s_lombardVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    // Set the path for the destination chain.
    s_lombardVerifier.setPath(DEST_CHAIN_SELECTOR, LOMBARD_CHAIN_ID, ALLOWED_CALLER);

    // Mock the router to return true for the valid offRamp.
    vm.mockCall(
      address(s_router), abi.encodeCall(IRouter.isOffRamp, (DEST_CHAIN_SELECTOR, s_offRamp)), abi.encode(true)
    );

    // Mock the router to return the valid onRamp.
    vm.mockCall(address(s_router), abi.encodeCall(IRouter.getOnRamp, (DEST_CHAIN_SELECTOR)), abi.encode(s_onRamp));
  }

  /// @notice Creates a MessageV1 with a token transfer for forwardToVerifier tests.
  function _createForwardMessage(
    address sourceToken,
    uint256 amount,
    address receiver
  ) internal returns (MessageV1Codec.MessageV1 memory, bytes32) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(sourceToken),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: abi.encodePacked(receiver),
      extraData: ""
    });

    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      executionGasLimit: GAS_LIMIT * 2,
      ccipReceiveGasLimit: GAS_LIMIT,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encodePacked(s_onRamp),
      offRampAddress: abi.encodePacked(s_offRamp),
      sender: abi.encodePacked(OWNER),
      receiver: abi.encodePacked(receiver),
      destBlob: "",
      tokenTransfer: tokenTransfer,
      data: ""
    });

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    return (message, messageId);
  }
}
