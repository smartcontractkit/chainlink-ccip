// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../libraries/Internal.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_releaseOrMint is USDCTokenPoolCCTPV2Setup {
  // From https://github.com/circlefin/evm-cctp-contracts/blob/377c9bd813fb86a42d900ae4003599d82aef635a/src/messages/BurnMessage.sol#L57
  function _formatMessage(
    uint32 _version,
    bytes32 _burnToken,
    bytes32 _mintRecipient,
    uint256 _amount,
    bytes32 _messageSender
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(_version, _burnToken, _mintRecipient, _amount, _messageSender);
  }

  function testFuzz_releaseOrMint_Success(address recipient, uint256 amount) public {
    vm.assume(recipient != address(0) && recipient != address(s_USDCToken));
    amount = bound(amount, 0, _getInboundRateLimiterConfig().capacity);

    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xC11"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(recipient))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: _formatMessage(
        1,
        bytes32(uint256(uint160(address(s_USDCToken)))),
        bytes32(uint256(uint160(recipient))),
        amount,
        bytes32(uint256(uint160(OWNER)))
      )
    });

    bytes memory message = _generateUSDCMessageCCTPV2(usdcMessage);
    bytes memory attestation = bytes("attestation bytes");

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: 0,
          sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
          cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
          amount: amount,
          destinationDomain: DEST_DOMAIN_IDENTIFIER,
          mintRecipient: bytes32(0),
          burnToken: address(s_USDCToken),
          destinationCaller: bytes32(0),
          maxFee: 0,
          minFinalityThreshold: 0
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation}));

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: s_routerAllowedOffRamp,
      recipient: recipient,
      amount: amount
    });

    vm.expectCall(
      address(s_mockUSDCTransmitterCCTPV2),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, message, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  // https://etherscan.io/tx/0x8897ffb613c4d7823ab42c85df2410381e7028c40f0a924530db99412065e338
  function test_releaseOrMint_RealTx() public {
    bytes memory encodedUsdcMessage =
      hex"000000010000000500000000bc7e669b9f452229fdd08bac21b6617068f2fd023ad6e805d03afa88b4bb79aea65fc81d0fefa8860cb3b83f089b0224be8a6687b7ae49f594c0b9b4d7e9389300000000000000000000000028b5a0e9c621a5badaa536219b3a228c8168cf5d0000000000000000000000000000000000000000000000000000000000000000000007d0000007d000000001c6fa7af3bedbad3a3d65f36aabc97431b1bbe4c2d2f6e0e47ca60203452f5d61000000000000000000000000e7492c49f71841d0f55f4f22c2ee22f02437084000000000000000000000000000000000000000000000000000000017491105202c747e9f0b8a0bb74202136e08fb8463bb15d1ab1d6d3f916f547004d7c7522f0000000000000000000000000000000000000000000000000000000000989a720000000000000000000000000000000000000000000000000000000000989a7200000000000000000000000000000000000000000000000000000000015d0d4e";
    bytes memory attestation = bytes("attestation bytes");

    uint32 nonce = 4730;
    uint32 sourceDomain = 5;
    uint256 amount = 100;

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: nonce,
          sourceDomain: sourceDomain,
          cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
          amount: amount,
          destinationDomain: DEST_DOMAIN_IDENTIFIER,
          mintRecipient: bytes32(0),
          burnToken: address(s_USDCToken),
          destinationCaller: bytes32(0),
          maxFee: 0,
          minFinalityThreshold: 0
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    // The mocked receiver does not release the token to the pool, so we manually do it here
    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: encodedUsdcMessage, attestation: attestation}));

    vm.expectCall(
      address(s_mockUSDCTransmitterCCTPV2),
      abi.encodeWithSelector(MockE2EUSDCTransmitter.receiveMessage.selector, encodedUsdcMessage, attestation)
    );

    vm.startPrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: OWNER,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  // Reverts
  function test_releaseOrMint_RevertWhen_UnlockingUSDCFailed() public {
    vm.startPrank(s_routerAllowedOffRamp);
    s_mockUSDCTransmitterCCTPV2.setShouldSucceed(false);

    uint256 amount = 13255235235;

    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xC11"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDCTokenMessenger)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: _formatMessage(
        1,
        bytes32(uint256(uint160(address(s_USDCToken)))),
        bytes32(uint256(uint160(OWNER))),
        amount,
        bytes32(uint256(uint160(OWNER)))
      )
    });

    Internal.SourceTokenData memory sourceTokenData = Internal.SourceTokenData({
      sourcePoolAddress: abi.encode(SOURCE_CHAIN_USDC_POOL),
      destTokenAddress: abi.encode(address(s_usdcTokenPool)),
      extraData: abi.encode(
        USDCTokenPool.SourceTokenDataPayload({
          nonce: 0,
          sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
          cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
          amount: amount,
          destinationDomain: DEST_DOMAIN_IDENTIFIER,
          mintRecipient: bytes32(0),
          burnToken: address(s_USDCToken),
          destinationCaller: bytes32(0),
          maxFee: 0,
          minFinalityThreshold: 0
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
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }

  function test_releaseOrMint_RevertWhen_TokenMaxCapacityExceeded() public {
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
          cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
          amount: amount,
          destinationDomain: DEST_DOMAIN_IDENTIFIER,
          mintRecipient: bytes32(0),
          burnToken: address(s_USDCToken),
          destinationCaller: bytes32(0),
          maxFee: 0,
          minFinalityThreshold: 0
        })
      ),
      destGasAmount: USDC_DEST_TOKEN_GAS
    });

    bytes memory offchainTokenData =
      abi.encode(USDCTokenPool.MessageAndAttestation({message: bytes(""), attestation: bytes("")}));

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, capacity, amount, address(s_USDCToken))
    );

    s_usdcTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(OWNER),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_USDCToken),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: sourceTokenData.sourcePoolAddress,
        sourcePoolData: sourceTokenData.extraData,
        offchainTokenData: offchainTokenData
      })
    );
  }
}
