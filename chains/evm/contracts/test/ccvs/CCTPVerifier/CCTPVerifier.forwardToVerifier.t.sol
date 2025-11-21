// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CCTPVerifier_forwardToVerifier is CCTPVerifierSetup {
  function setUp() public override {
    super.setUp();

    // Send transfer amount to the verifier, mocking a transfer from the token pool.
    deal(address(s_USDCToken), address(s_cctpVerifier), s_transferAmount);
    assertEq(IERC20(address(s_USDCToken)).balanceOf(address(s_cctpVerifier)), s_transferAmount);
  }

  function test_forwardToVerifier_MintRecipientFromMessage() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 0, address(s_USDCToken), s_transferAmount, tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      s_transferAmount,
      address(s_cctpVerifier),
      abi.decode(tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      0,
      CCTP_STANDARD_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_MintRecipientFromDomain() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 0, address(s_USDCToken), s_transferAmount, tokenReceiver
    );

    // Set a custom mint recipient for the domain.
    bytes32 customMintRecipient = bytes32(uint256(uint160(makeAddr("customMintRecipient"))));
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: customMintRecipient,
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });
    s_cctpVerifier.setDomains(domainUpdates);

    vm.startPrank(s_onRamp);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      s_transferAmount,
      address(s_cctpVerifier),
      customMintRecipient,
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      0,
      CCTP_STANDARD_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_CustomFinality() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    uint256 expectedMaxFee = s_transferAmount * CCTP_FAST_FINALITY_BPS / BPS_DIVIDER;

    vm.startPrank(s_onRamp);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      s_transferAmount,
      address(s_cctpVerifier),
      abi.decode(tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      uint32(expectedMaxFee),
      CCTP_FAST_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_MultipleCustomFinalities() public {
    // Update finality config to have 2 custom finalities.
    uint32 cctpSlowFinalityThreshold = 5000;
    uint256 expectedMaxFee;
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    // Use a finality between 1 and 100 to ensure that we round up to the slower finality.
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, 50, address(s_USDCToken), s_transferAmount, tokenReceiver
    );
    {
      uint16 ccipSlowFinalityThreshold = 100;
      uint16 cctpSlowFinalityBps = 5; // 0.05%

      uint16[] memory customCCIPFinalities = new uint16[](2);
      customCCIPFinalities[0] = CCIP_FAST_FINALITY_THRESHOLD;
      customCCIPFinalities[1] = ccipSlowFinalityThreshold;
      uint32[] memory customCCTPFinalityThresholds = new uint32[](2);
      customCCTPFinalityThresholds[0] = CCTP_FAST_FINALITY_THRESHOLD;
      customCCTPFinalityThresholds[1] = cctpSlowFinalityThreshold;
      uint16[] memory customCCTPFinalityBps = new uint16[](2);
      customCCTPFinalityBps[0] = CCTP_FAST_FINALITY_BPS;
      customCCTPFinalityBps[1] = cctpSlowFinalityBps;
      CCTPVerifier.FinalityConfig memory finalityConfig = CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: customCCIPFinalities,
        customCCTPFinalityThresholds: customCCTPFinalityThresholds,
        customCCTPFinalityBps: customCCTPFinalityBps
      });
      s_cctpVerifier.setFinalityConfig(finalityConfig);
      expectedMaxFee = s_transferAmount * cctpSlowFinalityBps / BPS_DIVIDER;
    }

    vm.startPrank(s_onRamp);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      s_transferAmount,
      address(s_cctpVerifier),
      abi.decode(tokenReceiver, (bytes32)),
      REMOTE_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      ALLOWED_CALLER_ON_DEST,
      uint32(expectedMaxFee),
      cctpSlowFinalityThreshold,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, STRANGER));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_SenderIsNotAllowed() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    // Enable allowlist, adding owner as the only allowed sender.
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, allowedSenders, new address[](0));
    s_cctpVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, address(bytes20(message.sender))));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_DestinationNotSupported() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    uint64 unknownDestChainSelector = 99999;
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      unknownDestChainSelector,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.DestinationNotSupported.selector, unknownDestChainSelector));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_UnknownDomain() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

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

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.UnknownDomain.selector, DEST_CHAIN_SELECTOR));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidTokenTransferLength() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    // Message has to be updated here because message encoding will fail with multiple token transfers.
    message.tokenTransfer = new MessageV1Codec.TokenTransferV1[](2);
    message.tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: s_transferAmount,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(address(s_USDCToken)),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: tokenReceiver,
      extraData: "extra data"
    });
    message.tokenTransfer[1] = MessageV1Codec.TokenTransferV1({
      amount: s_transferAmount,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(address(s_USDCToken)),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: tokenReceiver,
      extraData: "extra data"
    });

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidTokenTransferLength.selector, 2));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidToken() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    address invalidToken = makeAddr("invalidToken");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      invalidToken,
      s_transferAmount,
      tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidToken.selector, invalidToken));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidReceiver() public {
    bytes memory tokenReceiver = abi.encodePacked(makeAddr("invalidReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      s_transferAmount,
      tokenReceiver
    );

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidReceiver.selector, tokenReceiver));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_RevertWhen_MaxFeeExceedsUint32() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    // Use a large amount that will exceed the uint32 max.
    uint256 largeAmount = 50000000000000; // 50 million USDC
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createCCIPMessage(
      SOURCE_CHAIN_SELECTOR,
      DEST_CHAIN_SELECTOR,
      CCIP_FAST_FINALITY_THRESHOLD,
      address(s_USDCToken),
      largeAmount,
      tokenReceiver
    );

    uint256 expectedMaxFee = largeAmount * CCTP_FAST_FINALITY_BPS / BPS_DIVIDER;

    vm.startPrank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.MaxFeeExceedsUint32.selector, expectedMaxFee));
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }
}
