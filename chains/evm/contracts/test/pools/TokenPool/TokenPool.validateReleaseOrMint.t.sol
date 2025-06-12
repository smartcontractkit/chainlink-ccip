// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolSetup} from "./TokenPoolSetup.t.sol";

contract TokenPool_validateReleaseOrMint is TokenPoolSetup {
  uint256 internal constant AMOUNT = 100;

  function setUp() public override {
    super.setUp();

    vm.startPrank(s_allowedOffRamp);
  }

  function test_validateReleaseOrMint_Success() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: address(s_token),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(address(0x456)),
      offchainTokenData: ""
    });

    // Should not revert
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT);
  }

  function test_validateReleaseOrMint_RateLimitLocalAmount() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: address(s_token),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(address(0x456)),
      offchainTokenData: ""
    });

    // Pretend the local amount is 10x the source amount and assert the rate limit is applied on the local amount
    uint256 localAmount = AMOUNT * 10;

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: releaseOrMintIn.localToken,
      amount: localAmount
    });

    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, localAmount);
  }

  function test_validateReleaseOrMint_InvalidToken() public {
    address wrongToken = address(0x456);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: wrongToken, // Invalid token address
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(address(0x456)),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, wrongToken));
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT);
  }

  function test_validateReleaseOrMint_CursedByRMN() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: address(s_token),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(0x456),
      offchainTokenData: ""
    });

    // Mock RMN to be cursed
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(DEST_CHAIN_SELECTOR))),
      abi.encode(true)
    );

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT);
  }

  function test_validateReleaseOrMint_InvalidOffRamp() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: address(s_token),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(address(0x456)),
      offchainTokenData: ""
    });

    // Mock router to return false for isOffRamp
    vm.mockCall(
      address(s_sourceRouter),
      abi.encodeWithSelector(IRouter.isOffRamp.selector, DEST_CHAIN_SELECTOR, s_allowedOffRamp),
      abi.encode(false)
    );

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotARampOnRouter.selector, s_allowedOffRamp));
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT);
  }

  function test_validateReleaseOrMint_InvalidSourcePool() public {
    address invalidPool = address(0x789);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      localToken: address(s_token),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      sourcePoolAddress: abi.encode(invalidPool),
      sourceDenominatedAmount: AMOUNT,
      sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
      receiver: address(0x123),
      originalSender: abi.encode(address(0x456)),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, abi.encode(invalidPool)));
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT);
  }
}
