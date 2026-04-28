// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPool_applyPauseUpdates is AdvancedPoolHooksSetup {
  uint256 internal constant AMOUNT = 100e18;

  function setUp() public override {
    super.setUp();
    // BaseTest starts a persistent OWNER prank; stop so we can `vm.prank` explicitly per test.
    vm.stopPrank();
  }

  function test_applyPauseUpdates_globalOutbound_blocksLockOrBurn() public {
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(true, false, new uint64[](0), new uint8[](0));

    vm.expectRevert(abi.encodeWithSelector(TokenPool.OutboundPoolPaused.selector, DEST_CHAIN_SELECTOR));
    vm.prank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(AMOUNT), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 0);
  }

  function test_applyPauseUpdates_globalInbound_blocksReleaseOrMint() public {
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, true, new uint64[](0), new uint8[](0));

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InboundPoolPaused.selector, DEST_CHAIN_SELECTOR));
    vm.prank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(_buildReleaseOrMintIn(AMOUNT), AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_applyPauseUpdates_laneOutboundOnly_blocksOutboundNotInbound() public {
    uint64[] memory chains = new uint64[](1);
    chains[0] = DEST_CHAIN_SELECTOR;
    uint8[] memory flags = new uint8[](1);
    flags[0] = 1; // outbound only

    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.OutboundPoolPaused.selector, DEST_CHAIN_SELECTOR));
    vm.prank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(AMOUNT), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 0);

    vm.prank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(_buildReleaseOrMintIn(AMOUNT), AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_applyPauseUpdates_laneInboundOnly_blocksInboundNotOutbound() public {
    uint64[] memory chains = new uint64[](1);
    chains[0] = DEST_CHAIN_SELECTOR;
    uint8[] memory flags = new uint8[](1);
    flags[0] = 2; // PAUSE_LANE_INBOUND

    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InboundPoolPaused.selector, DEST_CHAIN_SELECTOR));
    vm.prank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(_buildReleaseOrMintIn(AMOUNT), AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    vm.prank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(AMOUNT), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 0);
  }

  function test_applyPauseUpdates_pauseAdmin() public {
    address pauseAdmin = makeAddr("pauseAdmin");
    vm.prank(OWNER);
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), address(0), address(0), pauseAdmin);

    vm.prank(pauseAdmin);
    s_tokenPool.applyPauseUpdates(true, false, new uint64[](0), new uint8[](0));

    (bool ob,) = s_tokenPool.getGlobalPauseState();
    assertTrue(ob);
  }

  function test_applyPauseUpdates_RevertWhen_Unauthorized() public {
    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, STRANGER));
    vm.prank(STRANGER);
    s_tokenPool.applyPauseUpdates(true, false, new uint64[](0), new uint8[](0));
  }

  function test_applyPauseUpdates_RevertWhen_LengthMismatch() public {
    uint64[] memory chains = new uint64[](1);
    chains[0] = DEST_CHAIN_SELECTOR;
    vm.expectRevert(TokenPool.PauseUpdatesLengthMismatch.selector);
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, new uint8[](0));
  }

  function test_applyPauseUpdates_RevertWhen_InvalidPauseFlags() public {
    uint64[] memory chains = new uint64[](1);
    chains[0] = DEST_CHAIN_SELECTOR;
    uint8[] memory flags = new uint8[](1);
    flags[0] = 4;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidPauseFlags.selector, uint8(4)));
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);
  }

  function test_applyPauseUpdates_RevertWhen_NonExistentChain() public {
    uint64 badSelector = 99999;
    uint64[] memory chains = new uint64[](1);
    chains[0] = badSelector;
    uint8[] memory flags = new uint8[](1);
    flags[0] = 1;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, badSelector));
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);
  }

  function test_applyPauseUpdates_clearsLaneFlags() public {
    uint64[] memory chains = new uint64[](1);
    chains[0] = DEST_CHAIN_SELECTOR;
    uint8[] memory flags = new uint8[](1);
    flags[0] = 3; // outbound + inbound

    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);
    assertEq(s_tokenPool.getLanePauseFlags(DEST_CHAIN_SELECTOR), flags[0]);

    flags[0] = 0;
    vm.prank(OWNER);
    s_tokenPool.applyPauseUpdates(false, false, chains, flags);
    assertEq(s_tokenPool.getLanePauseFlags(DEST_CHAIN_SELECTOR), 0);
  }

  function _buildLockOrBurnIn(
    uint256 amount
  ) internal view returns (Pool.LockOrBurnInV1 memory lockOrBurnIn) {
    return Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });
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
      sourcePoolData: abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)),
      offchainTokenData: ""
    });
  }
}
