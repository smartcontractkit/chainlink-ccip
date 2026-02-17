// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {IBridgeV3} from "../../../interfaces/lombard/IBridgeV3.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MockLombardBridge} from "../../mocks/MockLombardBridge.sol";
import {MockLombardMailbox} from "../../mocks/MockLombardMailbox.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract LombardVerifierSetup is BaseVerifierSetup {
  bytes4 internal constant VERSION_TAG_V1_7_0 = bytes4(keccak256("LombardVerifier 1.7.0"));

  LombardVerifier internal s_lombardVerifier;
  MockLombardBridge internal s_mockBridge;
  MockLombardMailbox internal s_mockMailbox;
  BurnMintERC20 internal s_testToken;

  bytes32 internal constant LOMBARD_CHAIN_ID = bytes32(uint256(10000));
  bytes32 internal constant ALLOWED_CALLER = bytes32(uint256(0x123456));
  uint256 internal constant TRANSFER_AMOUNT = 1e18;

  function setUp() public virtual override {
    super.setUp();

    s_mockBridge = new MockLombardBridge();
    s_mockMailbox = MockLombardMailbox(s_mockBridge.s_mailbox());
    // Set default execution result matching the version tag format.
    s_mockMailbox.setMessageId(abi.encodePacked(VERSION_TAG_V1_7_0, bytes32(0)));

    s_lombardVerifier = new LombardVerifier(
      LombardVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      IBridgeV3(address(s_mockBridge)),
      s_storageLocations,
      address(s_mockRMNRemote)
    );

    // Deploy test token and add it as a supported token.
    s_testToken = new BurnMintERC20("Test Token", "TEST", 18, 0, 0);
    deal(address(s_testToken), address(s_lombardVerifier), TRANSFER_AMOUNT);
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(s_testToken), localAdapter: address(0)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Set up remote chain config with the router.
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](2);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    remoteChainConfigs[1] = _getRemoteChainConfig(s_router, SOURCE_CHAIN_SELECTOR, false);
    s_lombardVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    // Set the path for the destination chain.
    s_lombardVerifier.setPath(DEST_CHAIN_SELECTOR, LOMBARD_CHAIN_ID, ALLOWED_CALLER);

    // Mock the router to return true for the valid offRamp.
    vm.mockCall(
      address(s_router), abi.encodeCall(IRouter.isOffRamp, (DEST_CHAIN_SELECTOR, s_offRamp)), abi.encode(true)
    );

    // Mock the router to return the valid onRamp.
    vm.mockCall(address(s_router), abi.encodeCall(IRouter.getOnRamp, (DEST_CHAIN_SELECTOR)), abi.encode(s_onRamp));
    vm.mockCall(
      address(s_router), abi.encodeCall(IRouter.isOffRamp, (SOURCE_CHAIN_SELECTOR, s_offRamp)), abi.encode(true)
    );
  }

  function _createForwardMessage(
    address sourceToken,
    address receiver
  ) internal returns (MessageV1Codec.MessageV1 memory, bytes32) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: TRANSFER_AMOUNT,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encode(sourceToken),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: abi.encodePacked(receiver),
      extraData: ""
    });

    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: GAS_LIMIT * 2,
      ccipReceiveGasLimit: GAS_LIMIT,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encode(s_onRamp),
      offRampAddress: abi.encodePacked(s_offRamp),
      sender: abi.encode(OWNER),
      receiver: abi.encodePacked(receiver),
      destBlob: "",
      tokenTransfer: tokenTransfer,
      data: ""
    });

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    return (message, messageId);
  }

  /// @notice Generates a valid rawPayload for use in verifyMessage tests.
  /// @dev The rawPayload structure matches what Lombard bridge expects:
  /// [version (4 bytes)][abi.encode(destinationChain, nonce, sender, recipient, destinationCaller, msgBody)]
  /// where msgBody layout (read by assembly in _validatePayload):
  ///   byte 0:       version (1 byte)
  ///   bytes 1..32:  token (32 bytes)
  ///   bytes 33..64: unused (32 bytes)
  ///   bytes 65..96: recipient (32 bytes)
  ///   bytes 97..128: amount (32 bytes)
  /// @param destToken The destination token address.
  /// @param tokenReceiver The token receiver address.
  /// @param amount The amount to transfer.
  /// @return rawPayload The encoded payload.
  function _generateValidRawPayload(
    bytes memory destToken,
    bytes memory tokenReceiver,
    uint256 amount
  ) internal pure returns (bytes memory) {
    // Create msgBody matching the assembly offsets in _validatePayload:
    //   mload(msgBody + 0x21) => bytes 1..32  = token
    //   mload(msgBody + 0x61) => bytes 65..96 = recipient
    //   mload(msgBody + 0x81) => bytes 97..128 = amount
    bytes memory msgBody =
      abi.encodePacked(bytes1(0), bytes32(destToken), bytes32(0), bytes32(tokenReceiver), bytes32(amount));

    // Encode the full payload structure
    bytes memory encodedData = abi.encode(
      bytes32(LOMBARD_CHAIN_ID), // destinationChain
      uint256(1), // nonce
      bytes32(uint256(uint160(OWNER))), // sender
      address(0), // recipient (not used in validation)
      address(0), // destinationCaller (not used in validation)
      msgBody
    );

    // Prepend version tag (4 bytes)
    return abi.encodePacked(bytes4(0x01000000), encodedData);
  }
}
