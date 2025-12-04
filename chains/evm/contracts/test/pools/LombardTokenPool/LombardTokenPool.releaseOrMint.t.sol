// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";

import {MockMailbox} from "../../mocks/MockMailbox.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_releaseOrMint is LombardTokenPoolSetup {
  function setUp() public virtual override {
    super.setUp();
    vm.startPrank(s_allowedOffRamp);
  }

  bytes32 internal constant PAYLOAD_HASH = bytes32("payload-hash");

  function test_releaseOrMint_V1() public {
    MockMailbox mailbox = new MockMailbox();
    mailbox.setResult(PAYLOAD_HASH, true, "");
    s_bridge.setMailbox(address(mailbox));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encodePacked(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: address(0xBEEF),
      sourceDenominatedAmount: 5e17,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_remotePool),
      sourcePoolData: abi.encode(PAYLOAD_HASH),
      offchainTokenData: abi.encode(bytes("rawPayload"), bytes("proof"))
    });

    vm.expectEmit();
    emit TokenPool.ReleasedOrMinted({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: s_allowedOffRamp,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    Pool.ReleaseOrMintOutV1 memory out = s_pool.releaseOrMint(releaseOrMintIn);

    assertEq(out.destinationAmount, releaseOrMintIn.sourceDenominatedAmount);
    assertEq(mailbox.s_lastRawPayload(), bytes("rawPayload"));
  }

  function test_releaseOrMint_V1_RevertWhen_ExecutionError() public {
    MockMailbox mailbox = new MockMailbox();
    mailbox.setResult(PAYLOAD_HASH, false, "");
    s_bridge.setMailbox(address(mailbox));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encodePacked(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: address(0xBEEF),
      sourceDenominatedAmount: 1,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_remotePool),
      sourcePoolData: abi.encode(PAYLOAD_HASH),
      offchainTokenData: abi.encode(bytes("raw"), bytes("proof"))
    });

    vm.expectRevert(LombardTokenPool.ExecutionError.selector);
    s_pool.releaseOrMint(releaseOrMintIn);
  }

  function test_releaseOrMint_V1_RevertWhen_HashMismatch() public {
    MockMailbox mailbox = new MockMailbox();
    mailbox.setResult(bytes32("different"), true, "");
    s_bridge.setMailbox(address(mailbox));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encodePacked(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: address(0xBEEF),
      sourceDenominatedAmount: 1,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_remotePool),
      sourcePoolData: abi.encode(PAYLOAD_HASH),
      offchainTokenData: abi.encode(bytes("raw"), bytes("proof"))
    });

    vm.expectRevert(LombardTokenPool.HashMismatch.selector);
    s_pool.releaseOrMint(releaseOrMintIn);
  }
}
