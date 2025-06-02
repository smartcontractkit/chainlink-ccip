// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../../libraries/Internal.sol";
import {Pool} from "../../../../../libraries/Pool.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../../../pools/TokenPool.sol";

import {USDCTokenPoolCCTPV2} from "../../../../../pools/USDC/CCTPV2/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPool} from "../../../../../pools/USDC/USDCTokenPool.sol";
import {MockE2EUSDCTransmitter} from "../../../../mocks/MockE2EUSDCTransmitter.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_releaseOrMint is USDCTokenPoolCCTPV2Setup {
  // The threshold at which CCTP will attest to a message. A fast threshold only waites for tx-confirmation
  // while slow threshold requires finalization. Any minFinalityThreshold value below 1000 is treated as 1000, and any value above 1000 is treated as 2000 as defined at https://developers.circle.com/stablecoins/cctp-finality-and-fees
  uint32 public constant MIN_FINALITY_THRESHOLD_SLOW = 2000;
  uint32 public constant MIN_FINALITY_THRESHOLD_FAST = 1000;

  // From https://developers.circle.com/stablecoins/message-format#message-body
  function _formatMessage(
    uint32 _version,
    bytes32 _burnToken,
    bytes32 _mintRecipient,
    uint256 _amount,
    bytes32 _messageSender,
    uint256 _maxFee,
    uint256 _feeExecuted,
    uint256 _expirationBlock,
    bytes memory _hookData
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      _version, _burnToken, _mintRecipient, _amount, _messageSender, _maxFee, _feeExecuted, _expirationBlock, _hookData
    );
  }

  function testFuzz_ReleaseOrMint_Success(address recipient, uint256 amount) public {
    vm.assume(recipient != address(0) && recipient != address(s_token));
    amount = bound(amount, 0, _getInboundRateLimiterConfig().capacity);

    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(recipient))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: MIN_FINALITY_THRESHOLD_SLOW,
      finalityThresholdExecuted: MIN_FINALITY_THRESHOLD_SLOW,
      messageBody: _formatMessage(
        0, // version
        bytes32(uint256(uint160(address(OWNER)))), // burnToken
        bytes32(uint256(uint160(recipient))), // mintRecipient
        amount, // amount
        bytes32(uint256(uint160(OWNER))), // messageSender
        0, // maxFee
        0, // feeExecuted
        block.number + (1 days / 12), // expirationBlock 1-day in the future assuming a 12-second block time.
        "" // hookData
      )
    });

    bytes memory message = _generateUSDCMessageCCTPV2(usdcMessage);
    bytes memory attestation = bytes("attestation bytes");

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: uint64(0),
          sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
          cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation}));

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_token), address(s_usdcTokenPool), amount);

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: address(s_token),
      sender: s_routerAllowedOffRamp,
      recipient: recipient,
      amount: amount
    });

    vm.expectCall(
      address(s_mockUSDCTransmitter),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, message, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  // Reverts
  function test_RevertWhen_UnlockingUSDCFailed() public {
    vm.startPrank(s_routerAllowedOffRamp);
    s_mockUSDCTransmitter.setShouldSucceed(false);

    uint256 amount = 13255235235;

    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDC)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: MIN_FINALITY_THRESHOLD_SLOW,
      finalityThresholdExecuted: MIN_FINALITY_THRESHOLD_SLOW,
      messageBody: _formatMessage(
        0, // version
        bytes32(uint256(uint160(address(OWNER)))), // burnToken
        bytes32(uint256(uint160(OWNER))), // mintRecipient
        amount, // amount
        bytes32(uint256(uint160(OWNER))), // messageSender
        0, // maxFee
        0, // feeExecuted
        block.number + (1 days / 12), // expirationBlock 1-day in the future assuming a 12-second block time.
        "" // hookData
      )
    });

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: uint64(0),
          sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
          cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessageCCTPV2(usdcMessage), attestation: bytes("")})
    );

    vm.expectRevert(USDCTokenPool.UnlockingUSDCFailed.selector);

    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  function test_RevertWhen_TokenMaxCapacityExceeded() public {
    uint256 capacity = _getInboundRateLimiterConfig().capacity;
    uint256 amount = 10 * capacity;
    address recipient = address(1);
    vm.startPrank(s_routerAllowedOffRamp);

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: 1,
          sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
          cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: bytes(""), attestation: bytes("")}));

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, capacity, amount, address(s_token))
    );

    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  // https://etherscan.io/tx/0xa1d6676041ab7200e6904a44e9bb3acb49044f36881b96f19a5274080562cf95
  function test_incomingMessage_Success() public {
    bytes memory encodedUsdcMessage =
      hex"000000010000000d00000000c51f10ffe8670c7acc9c995bd05747d064960e1680dc69b904c2b2114144ef5000000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d00000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d0000000000000000000000000000000000000000000000000000000000000000000007d0000007d00000000100000000000000000000000029219dd400f2bf60e5a23d13be72b486d4038894000000000000000000000000b38a90f14b24ae81ec0b8f1373694f5b59811d8a00000000000000000000000000000000000000000000000000000045d964b800000000000000000000000000b38a90f14b24ae81ec0b8f1373694f5b59811d8a0000000000000000000000000000000000000000000000000000000059682f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
    bytes memory attestation = bytes("attestation bytes");

    uint32 sourceDomain = 13;
    uint256 amount = 100;

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: uint64(0),
          sourceDomain: sourceDomain,
          cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_token), address(s_usdcTokenPool), amount);

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: encodedUsdcMessage, attestation: attestation}));

    vm.expectCall(
      address(s_mockUSDCTransmitter),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, encodedUsdcMessage, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        amount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }
}
