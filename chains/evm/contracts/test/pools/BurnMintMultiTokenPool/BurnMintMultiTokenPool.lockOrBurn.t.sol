// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolSetup} from "../MultiTokenPool/MultiTokenPoolSetup.t.sol";

/// @dev The flavor-specific token movement of BurnMintMultiTokenPool. The generic lockOrBurn checks are covered in
/// MultiTokenPool.lockOrBurn.t.sol.
contract BurnMintMultiTokenPool_lockOrBurn is MultiTokenPoolSetup {
  function test_lockOrBurn_BurnFrom() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _setFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnFrom, address(0));

    vm.startPrank(s_allowedOnRamp);
    vm.expectCall(
      address(s_token), abi.encodeWithSelector(bytes4(keccak256("burnFrom(address,uint256)")), address(s_pool), amount)
    );
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    // The pool approved itself for the maximum on first need.
    assertEq(s_token.allowance(address(s_pool), address(s_pool)), type(uint256).max);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_BurnWithFrom() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _setFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnWithFrom, address(0));

    vm.startPrank(s_allowedOnRamp);
    vm.expectCall(
      address(s_token), abi.encodeWithSelector(bytes4(keccak256("burn(address,uint256)")), address(s_pool), amount)
    );
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    assertEq(s_token.allowance(address(s_pool), address(s_pool)), type(uint256).max);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_BurnToAddress() public {
    uint256 amount = 20_000e18;
    address burnAddress = makeAddr("burnAddress");
    deal(address(s_token), address(s_pool), amount);
    _setFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, burnAddress);

    vm.startPrank(s_allowedOnRamp);
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    // The tokens moved to an address that can never move them, they are not burned.
    assertEq(s_token.balanceOf(burnAddress), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function _setFlavor(
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal {
    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.lockOrBurnFlavor = flavor;
    lane.lockOrBurnAddress = lockOrBurnAddress;
    s_poolConfigManager.setOutboundLane(lane);
  }
}
