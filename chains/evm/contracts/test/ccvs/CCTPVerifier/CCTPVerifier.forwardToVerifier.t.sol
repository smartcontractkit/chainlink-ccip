// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";

contract CCTPVerifier_forwardToVerifier is CCTPVerifierSetup {
  function setUp() public override {
    super.setUp();

    // Send transfer amount to the verifier, mocking a transfer from the token pool.
    deal(address(s_USDCToken), address(s_cctpVerifier), TRANSFER_AMOUNT);
    assertEq(s_USDCToken.balanceOf(address(s_cctpVerifier)), TRANSFER_AMOUNT);
  }

  function test_forwardToVerifier_MintRecipientFromMessage() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 0, address(s_USDCToken), TRANSFER_AMOUNT, s_tokenReceiver
    );

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      address(s_cctpVerifier),
      abi.decode(s_tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      0,
      CCTP_STANDARD_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );

    vm.startPrank(s_onRamp);
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_MintRecipientFromDomain() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 0, address(s_USDCToken), TRANSFER_AMOUNT, s_tokenReceiver
    );

    // Set a custom mint recipient for the domain.
    bytes32 customMintRecipient = bytes32(uint256(uint160(makeAddr("customMintRecipient"))));
    CCTPVerifier.SetDomainArgs[] memory domainUpdates = new CCTPVerifier.SetDomainArgs[](1);
    domainUpdates[0] = CCTPVerifier.SetDomainArgs({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: customMintRecipient,
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_cctpVerifier.setDomains(domainUpdates);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      address(s_cctpVerifier),
      customMintRecipient,
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      0,
      CCTP_STANDARD_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );

    vm.startPrank(s_onRamp);
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_CustomFinality() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );
    uint256 expectedMaxFee = TRANSFER_AMOUNT * CCTP_FAST_FINALITY_BPS / BPS_DIVIDER;

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      address(s_cctpVerifier),
      abi.decode(s_tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      uint32(expectedMaxFee),
      CCTP_FAST_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );

    vm.startPrank(s_onRamp);
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_CustomMaxFee() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    uint256 customMaxFee = 5e6; // 5 USDC
    bytes memory verifierArgs = abi.encode(customMaxFee);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      address(s_cctpVerifier),
      abi.decode(s_tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      uint32(customMaxFee),
      CCTP_FAST_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );

    vm.startPrank(s_onRamp);
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, verifierArgs);
  }

  function test_forwardToVerifier_RevertWhen_CursedByRMN() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 0, address(s_USDCToken), TRANSFER_AMOUNT, s_tokenReceiver
    );

    _setMockRMNChainCurse(message.destChainSelector, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.destChainSelector));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, STRANGER));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_SenderIsNotAllowed() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    // Enable allowlist, adding owner as the only allowed sender.
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, allowedSenders, new address[](0));
    s_cctpVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.startPrank(s_onRamp);
    vm.expectRevert(
      abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, abi.decode(message.sender, (address)))
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_DestinationNotSupported() public {
    uint64 unknownDestChainSelector = 99999;
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      unknownDestChainSelector,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.RemoteChainNotSupported.selector, unknownDestChainSelector));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_UnknownDomain() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    // Disable domain.
    CCTPVerifier.SetDomainArgs[] memory domainUpdates = new CCTPVerifier.SetDomainArgs[](1);
    domainUpdates[0] = CCTPVerifier.SetDomainArgs({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: false
    });
    s_cctpVerifier.setDomains(domainUpdates);

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.UnknownDomain.selector, DEST_CHAIN_SELECTOR));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidTokenTransferLength() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    // Message has to be updated here because message encoding will fail with multiple token transfers.
    message.tokenTransfer = new MessageV1Codec.TokenTransferV1[](2);
    message.tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: TRANSFER_AMOUNT,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encode(address(s_USDCToken)),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: s_tokenReceiver,
      extraData: "extra data"
    });
    message.tokenTransfer[1] = MessageV1Codec.TokenTransferV1({
      amount: TRANSFER_AMOUNT,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encode(address(s_USDCToken)),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: s_tokenReceiver,
      extraData: "extra data"
    });

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidTokenTransferLength.selector, 2));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidToken() public {
    address invalidToken = makeAddr("invalidToken");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      invalidToken,
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidToken.selector, abi.encode(invalidToken)));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidReceiver() public {
    bytes memory tokenReceiver = new bytes(33); // 33 bytes is too long for a bytes32.
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidReceiver.selector, tokenReceiver));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidVerifierArgsLength() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      TRANSFER_AMOUNT,
      s_tokenReceiver
    );

    // verifierArgs is too long (64 bytes)
    bytes memory verifierArgs = abi.encode(uint256(1), uint256(2));

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidVerifierArgsLength.selector, 64));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, verifierArgs);
  }
}
