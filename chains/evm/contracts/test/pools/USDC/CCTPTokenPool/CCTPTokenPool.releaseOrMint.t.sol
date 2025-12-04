// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {CCTPTokenPoolSetup} from "./CCTPTokenPoolSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPTokenPool_releaseOrMint is CCTPTokenPoolSetup {
  function test_releaseOrMint() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: abi.encode(makeAddr("originalSender")),
      receiver: makeAddr("receiver"),
      sourceDenominatedAmount: 1000000000000000000,
      localToken: address(s_USDCToken),
      sourcePoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG),
      sourcePoolAddress: abi.encode(DEST_CHAIN_USDC_POOL),
      offchainTokenData: ""
    });

    Pool.ReleaseOrMintOutV1 memory expectedOut =
      Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.sourceDenominatedAmount});

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(s_USDCToken),
      sender: address(s_allowedCaller),
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    vm.startPrank(s_allowedCaller);
    Pool.ReleaseOrMintOutV1 memory actualOut = s_cctpTokenPool.releaseOrMint(releaseOrMintIn, 0);

    assertEq(actualOut.destinationAmount, expectedOut.destinationAmount);
  }

  function test_releaseOrMint_RevertWhen_IPoolV1NotSupported() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: abi.encode(makeAddr("originalSender")),
      receiver: makeAddr("receiver"),
      sourceDenominatedAmount: 1000000000000000000,
      localToken: address(s_USDCToken),
      sourcePoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_1_TAG),
      sourcePoolAddress: abi.encode(DEST_CHAIN_USDC_POOL),
      offchainTokenData: ""
    });

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(CCTPTokenPool.IPoolV1NotSupported.selector));
    s_cctpTokenPool.releaseOrMint(releaseOrMintIn);
  }

  function test_releaseOrMint_RevertWhen_InvalidCaller() public {
    address invalidCaller = makeAddr("invalidCaller");
    vm.startPrank(invalidCaller);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: abi.encode(makeAddr("originalSender")),
      receiver: makeAddr("receiver"),
      sourceDenominatedAmount: 1000000000000000000,
      localToken: address(s_USDCToken),
      sourcePoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG),
      sourcePoolAddress: abi.encode(DEST_CHAIN_USDC_POOL),
      offchainTokenData: ""
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, invalidCaller));
    s_cctpTokenPool.releaseOrMint(releaseOrMintIn, 0);
  }

  function test_releaseOrMint_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = DEST_CHAIN_SELECTOR + 1;

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: wrongChainSelector,
      originalSender: abi.encode(makeAddr("originalSender")),
      receiver: makeAddr("receiver"),
      sourceDenominatedAmount: 1000000000000000000,
      localToken: address(s_USDCToken),
      sourcePoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG),
      sourcePoolAddress: abi.encode(DEST_CHAIN_USDC_POOL),
      offchainTokenData: ""
    });

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_cctpTokenPool.releaseOrMint(releaseOrMintIn, 0);
  }
}
