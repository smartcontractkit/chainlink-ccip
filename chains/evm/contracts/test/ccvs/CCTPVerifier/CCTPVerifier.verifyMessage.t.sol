// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMessageTransmitter} from "../../../pools/USDC/interfaces/IMessageTransmitter.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CCTPVerifier_verifyMessage is CCTPVerifierSetup {
  CCTPVerifierSetup.CCTPMessage internal s_baseCCTPMessage;
  address internal s_mintRecipient;

  function setUp() public virtual override {
    super.setUp();

    s_mintRecipient = makeAddr("mintRecipient");

    s_baseCCTPMessage = CCTPVerifierSetup.CCTPMessage({
      header: CCTPVerifierSetup.CCTPMessageHeader({
        version: 1,
        sourceDomain: REMOTE_DOMAIN_IDENTIFIER,
        destinationDomain: LOCAL_DOMAIN_IDENTIFIER,
        nonce: bytes32(0),
        sender: bytes32(0),
        recipient: bytes32(abi.encode(s_mockTokenMessenger)),
        destinationCaller: bytes32(abi.encode(s_messageTransmitterProxy)),
        minFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        finalityThresholdExecuted: CCTP_STANDARD_FINALITY_THRESHOLD
      }),
      body: CCTPVerifierSetup.CCTPMessageBody({
        version: 1,
        burnToken: bytes32(abi.encode(s_USDCToken)),
        mintRecipient: bytes32(abi.encode(s_mintRecipient)),
        amount: s_transferAmount,
        messageSender: ALLOWED_CALLER_ON_SOURCE,
        maxFee: 1e6, // 1 USDC
        feeExecuted: 0,
        expirationBlock: block.number + 1000
      }),
      hookData: CCTPVerifierSetup.CCTPMessageHookData({
        verifierVersion: s_cctpVerifier.versionTag(),
        messageId: bytes32(0)
      })
    });

    // Set the domain for the source chain.
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_cctpVerifier.setDomains(domainUpdates);
  }

  function test_verifyMessage() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    s_baseCCTPMessage.hookData.messageId = messageHash;
    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);

    // Ensure that the mint recipient received the tokens.
    // Mock transmitter always just mints 1 token.
    assertEq(IERC20(address(s_USDCToken)).balanceOf(s_mintRecipient), 1);
  }

  function test_verifyMessage_RevertWhen_InvalidCCVData() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidCCVData.selector));
    s_cctpVerifier.verifyMessage(message, messageHash, "");
  }

  function test_verifyMessage_RevertWhen_InvalidCCVVersion_VersionPrefix() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );
    bytes4 invalidVersion = bytes4(uint32(0x01020304));

    s_baseCCTPMessage.hookData.verifierVersion = invalidVersion;
    bytes memory ccvData = _createCCVData(invalidVersion, s_baseCCTPMessage);

    vm.expectRevert(
      abi.encodeWithSelector(CCTPVerifier.InvalidCCVVersion.selector, s_cctpVerifier.versionTag(), invalidVersion)
    );
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidCCVVersion_AttestedVersion() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    bytes4 invalidVersion = bytes4(uint32(0x01020304));

    s_baseCCTPMessage.hookData.verifierVersion = invalidVersion;
    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    vm.expectRevert(
      abi.encodeWithSelector(CCTPVerifier.InvalidCCVVersion.selector, s_cctpVerifier.versionTag(), invalidVersion)
    );
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidMessageId.selector, messageHash, bytes32(0)));
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_UnknownDomain() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    s_baseCCTPMessage.hookData.messageId = messageHash;
    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    // Disable domain.
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: false
    });
    s_cctpVerifier.setDomains(domainUpdates);

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.UnknownDomain.selector, DEST_CHAIN_SELECTOR));
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidMessageSender() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    bytes32 invalidMessageSender = keccak256("invalidMessageSender");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    s_baseCCTPMessage.hookData.messageId = messageHash;
    s_baseCCTPMessage.body.messageSender = invalidMessageSender;
    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    vm.expectRevert(
      abi.encodeWithSelector(CCTPVerifier.InvalidMessageSender.selector, ALLOWED_CALLER_ON_SOURCE, invalidMessageSender)
    );
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_ReceiveMessageCallFailed() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageHash) = _createCCIPMessage(
      DEST_CHAIN_SELECTOR,
      SOURCE_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    s_baseCCTPMessage.hookData.messageId = messageHash;
    bytes memory ccvData = _createCCVData(s_cctpVerifier.versionTag(), s_baseCCTPMessage);

    vm.mockCall(
      address(s_mockMessageTransmitter),
      abi.encodeWithSelector(IMessageTransmitter.receiveMessage.selector),
      abi.encode(false)
    );
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.ReceiveMessageCallFailed.selector));
    s_cctpVerifier.verifyMessage(message, messageHash, ccvData);
  }
}
