// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {BurnMintMultiTokenPool} from "../../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {MultiTokenPoolSetup} from "../MultiTokenPool/MultiTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @dev The flavor-specific token movement of BurnMintMultiTokenPool. The generic lockOrBurn checks are covered in
/// MultiTokenPool.lockOrBurn.t.sol.
contract BurnMintMultiTokenPool_lockOrBurn is MultiTokenPoolSetup {
  function test_lockOrBurn_BurnFrom() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _mockFlavor(amount, IPoolConfigManager.LockOrBurnFlavor.BurnFrom, address(0));

    vm.startPrank(s_allowedOnRamp);
    vm.expectCall(
      address(s_token), abi.encodeWithSelector(bytes4(keccak256("burnFrom(address,uint256)")), address(s_pool), amount)
    );
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    // The first transfer sets the maximum allowance for the pool.
    assertEq(s_token.allowance(address(s_pool), address(s_pool)), type(uint256).max);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_BurnWithFrom() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _mockFlavor(amount, IPoolConfigManager.LockOrBurnFlavor.BurnWithFrom, address(0));

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
    _mockFlavor(amount, IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, burnAddress);

    vm.startPrank(s_allowedOnRamp);
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    // The tokens are sent to the configured burn address without reducing the total supply.
    assertEq(s_token.balanceOf(burnAddress), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_BurnToZeroAddress() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _mockFlavor(amount, IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, address(0));
    // Some tokens burn on a transfer to address(0). The test token reverts on it, so mock the transfer.
    vm.mockCall(address(s_token), abi.encodeCall(IERC20.transfer, (address(0), amount)), abi.encode(true));

    vm.startPrank(s_allowedOnRamp);
    vm.expectCall(address(s_token), abi.encodeCall(IERC20.transfer, (address(0), amount)));
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  // Reverts

  /// @dev Tokens sent to the pool would stay there and be withdrawn as fees.
  function test_lockOrBurn_RevertWhen_InvalidBurnAddress() public {
    uint256 amount = 20_000e18;
    deal(address(s_token), address(s_pool), amount);
    _mockFlavor(amount, IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, address(s_pool));

    vm.startPrank(s_allowedOnRamp);
    vm.expectRevert(
      abi.encodeWithSelector(BurnMintMultiTokenPool.InvalidBurnAddress.selector, address(s_token), address(s_pool))
    );
    s_pool.lockOrBurn(_lockOrBurnIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function _mockFlavor(
    uint256 amount,
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal {
    IPoolConfigManager.OutboundLane memory lane = _defaultOutboundLane();
    lane.lockOrBurnFlavor = flavor;
    lane.lockOrBurnAddress = lockOrBurnAddress;
    _mockValidateLockOrBurn(amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true, lane);
  }
}
