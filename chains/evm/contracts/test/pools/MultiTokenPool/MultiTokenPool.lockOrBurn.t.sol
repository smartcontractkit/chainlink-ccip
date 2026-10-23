// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPoolHooks} from "../../../interfaces/IMultiTokenPoolHooks.sol";
import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_lockOrBurn is MultiTokenPoolSetup {
  function test_lockOrBurn() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _lockOrBurnIn(amount);

    vm.startPrank(s_allowedOnRamp);

    vm.expectCall(address(s_token), abi.encodeWithSelector(bytes4(keccak256("burn(uint256)")), amount));

    vm.expectEmit();
    emit MultiTokenPool.LockedOrBurned(DEST_CHAIN_SELECTOR, address(s_token), s_allowedOnRamp, amount);

    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) =
      s_pool.lockOrBurn(lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    assertEq(destTokenAmount, amount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_remoteToken));
    assertEq(lockOrBurnOut.destPoolData, DEST_POOL_DATA);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_ChargesTokenFee() public {
    uint256 amount = 20_000e18;
    uint256 feeAmount = amount / 20;
    deal(address(s_token), address(s_pool), amount);

    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.feeAmount = feeAmount;
    s_poolConfigManager.setOutboundLane(lane);

    vm.startPrank(s_allowedOnRamp);
    (, uint256 destTokenAmount) = s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    assertEq(destTokenAmount, amount - feeAmount);
    // The pool keeps the fee and burns the rest.
    assertEq(s_token.balanceOf(address(s_pool)), feeAmount);
    assertTrue(s_poolConfigManager.lastChargeTokenFee());
  }

  function test_lockOrBurn_LegacyPathChargesNoFee() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);

    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.feeAmount = amount / 20;
    s_poolConfigManager.setOutboundLane(lane);

    vm.startPrank(s_allowedOnRamp);
    Pool.LockOrBurnOutV1 memory lockOrBurnOut = s_pool.lockOrBurn(_lockOrBurnIn(amount));

    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_remoteToken));
    assertFalse(s_poolConfigManager.lastChargeTokenFee());
    assertEq(s_poolConfigManager.lastRequestedFinalityConfig(), FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_lockOrBurn_CallsPreflightHook() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);

    address hook = makeAddr("hook");
    vm.etch(hook, bytes("fake bytecode"));
    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.hook = hook;
    s_poolConfigManager.setOutboundLane(lane);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _lockOrBurnIn(amount);

    vm.startPrank(s_allowedOnRamp);
    vm.expectCall(
      hook,
      abi.encodeCall(
        IMultiTokenPoolHooks.preflightCheck, (lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "args", amount)
      )
    );
    s_pool.lockOrBurn(lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "args");
  }

  function test_lockOrBurn_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.startPrank(s_allowedOnRamp);
    vm.expectRevert(MultiTokenPool.CursedByRMN.selector);
    s_pool.lockOrBurn(_lockOrBurnIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_lockOrBurn_RevertWhen_CallerIsNotARampOnRouter() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPool.CallerIsNotARampOnRouter.selector, STRANGER));
    s_pool.lockOrBurn(_lockOrBurnIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_lockOrBurn_RevertWhen_UnsupportedLockOrBurnFlavor() public {
    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Lock;
    s_poolConfigManager.setOutboundLane(lane);

    vm.startPrank(s_allowedOnRamp);
    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPool.UnsupportedLockOrBurnFlavor.selector, address(s_token), IPoolConfigManager.LockOrBurnFlavor.Lock
      )
    );
    s_pool.lockOrBurn(_lockOrBurnIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }
}
