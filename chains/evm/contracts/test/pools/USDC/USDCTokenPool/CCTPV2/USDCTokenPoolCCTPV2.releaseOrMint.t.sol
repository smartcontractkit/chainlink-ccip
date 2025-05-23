// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../../libraries/Internal.sol";
import {Pool} from "../../../../../libraries/Pool.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../../pools/USDC/cctpV2/USDCTokenPoolCCTPV2.sol";
import {MockE2EUSDCTransmitter} from "../../../../mocks/MockE2EUSDCTransmitter.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_releaseOrMint is USDCTokenPoolCCTPV2Setup {
  // TODO: Better define what this means
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
    return abi.encodePacked(_version, _burnToken, _mintRecipient, _amount, _messageSender, _maxFee,
    _feeExecuted, _expirationBlock, _hookData);
  }

  function testFuzz_ReleaseOrMint_Success(address recipient, uint256 amount) public {
    vm.assume(recipient != address(0) && recipient != address(s_token));
    amount = bound(amount, 0, _getInboundRateLimiterConfig().capacity);

    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 0,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(recipient))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: MIN_FINALITY_THRESHOLD_SLOW,
      finalityThresholdExecuted: MIN_FINALITY_THRESHOLD_SLOW,

      // TODO: Add Comments about each field
      messageBody: _formatMessage(
        0,
        bytes32(uint256(uint160(address(OWNER)))),
        bytes32(uint256(uint160(recipient))),
        amount,
        bytes32(uint256(uint160(OWNER))),
        0,
        0, 
        block.number + (1 days / 12), // TODO: Add Comment about being a block 24-hours in the future
        ""
      )
    });

    bytes memory message = _generateUSDCMessageCCTPV2(usdcMessage);
    bytes memory attestation = bytes("attestation bytes");

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        SOURCE_DOMAIN_IDENTIFIER
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation}));

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_token), address(s_usdcTokenPool), amount);

    vm.expectEmit();
    emit TokenPool.Minted(s_routerAllowedOffRamp, recipient, amount);

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
      version: 0,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDC)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: MIN_FINALITY_THRESHOLD_SLOW,
      finalityThresholdExecuted: MIN_FINALITY_THRESHOLD_SLOW,

      // TODO: Add Comments
      messageBody: _formatMessage(
        0,
        bytes32(uint256(uint160(address(OWNER)))),
        bytes32(uint256(uint160(OWNER))),
        amount,
        bytes32(uint256(uint160(OWNER))),
        0, 
        0, 
        block.number + (1 days / 12), // TODO: Comments
        "" 
      )
    });

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        SOURCE_DOMAIN_IDENTIFIER
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
      extraData: abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: 1, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})),
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
}
