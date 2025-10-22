// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_validateReleaseOrMint is TokenPoolV2Setup {
  uint256 internal constant AMOUNT = 100e18;

  function test_validateReleaseOrMint() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), AMOUNT);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);

    assertEq(localAmount, AMOUNT);
  }

  function test_validateReleaseOrMint_NonZeroFinality() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    vm.expectEmit();
    emit TokenPool.CustomFinalityTransferInboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), AMOUNT);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 2);

    assertEq(localAmount, AMOUNT);
  }

  function test_validateReleaseOrMint_RateLimitLocalAmount() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Pretend the local amount is 10x the source amount and assert the rate limit is applied on the local amount.
    uint256 localAmount = AMOUNT * 10;

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: releaseOrMintIn.localToken,
      amount: localAmount
    });

    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, localAmount, 0);
  }

  function test_validateReleaseOrMint_InvalidToken() public {
    address wrongToken = address(0x456);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);
    releaseOrMintIn.localToken = wrongToken; // Invalid token address.

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, wrongToken));
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_CursedByRMN() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Mock RMN to be cursed
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(DEST_CHAIN_SELECTOR))),
      abi.encode(true)
    );

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_InvalidOffRamp() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Mock router to return false for isOffRamp
    vm.mockCall(
      address(s_sourceRouter),
      abi.encodeWithSelector(IRouter.isOffRamp.selector, DEST_CHAIN_SELECTOR, s_allowedOffRamp),
      abi.encode(false)
    );

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotARampOnRouter.selector, s_allowedOffRamp));
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_InvalidSourcePool() public {
    address invalidPool = address(0x789);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);
    releaseOrMintIn.sourcePoolAddress = abi.encode(invalidPool);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, abi.encode(invalidPool)));
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function _buildReleaseOrMintIn(
    uint256 amount
  ) internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: OWNER,
      sourceDenominatedAmount: amount,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourcePoolData: abi.encode(uint256(s_token.decimals())),
      offchainTokenData: ""
    });
  }
}
