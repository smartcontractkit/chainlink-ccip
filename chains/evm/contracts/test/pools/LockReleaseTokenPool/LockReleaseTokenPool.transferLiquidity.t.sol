// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {LockReleaseTokenPoolSetup} from "./LockReleaseTokenPoolSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract LockReleaseTokenPool_transferLiquidity is LockReleaseTokenPoolSetup {
  LockReleaseTokenPool internal s_oldLockReleaseTokenPool;
  uint256 internal s_amount = 100000;

  function setUp() public virtual override {
    super.setUp();

    s_oldLockReleaseTokenPool = new LockReleaseTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(s_lockBox)
    );

    // Configure old pool as allowed caller in the lockBox.
    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = address(s_oldLockReleaseTokenPool);
    s_lockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    deal(address(s_token), address(s_lockBox), s_amount);
  }

  function test_transferLiquidity() public {
    uint256 balancePre = s_token.balanceOf(address(s_lockReleaseTokenPool));

    s_oldLockReleaseTokenPool.setRebalancer(address(s_lockReleaseTokenPool));

    vm.expectEmit();
    emit LockReleaseTokenPool.LiquidityTransferred(address(s_oldLockReleaseTokenPool), s_amount);

    s_lockReleaseTokenPool.transferLiquidity(address(s_oldLockReleaseTokenPool), s_amount);

    assertEq(s_token.balanceOf(address(s_lockReleaseTokenPool)), balancePre + s_amount);
  }

  function test_transferLiquidity_Uint256MaxAmountInput() public {
    deal(address(s_token), address(s_oldLockReleaseTokenPool), s_amount);
    uint256 balancePre = s_token.balanceOf(address(s_lockReleaseTokenPool));

    s_oldLockReleaseTokenPool.setRebalancer(address(s_lockReleaseTokenPool));

    uint256 balanceOldPool = s_token.balanceOf(address(s_oldLockReleaseTokenPool));

    vm.expectEmit();
    emit LockReleaseTokenPool.LiquidityTransferred(address(s_oldLockReleaseTokenPool), balanceOldPool);

    s_lockReleaseTokenPool.transferLiquidity(address(s_oldLockReleaseTokenPool), type(uint256).max);

    assertEq(s_token.balanceOf(address(s_lockReleaseTokenPool)), balancePre + balanceOldPool);
  }

  function test_transferLiquidity_RevertWhen_InsufficientBalance() public {
    uint256 balancePre = s_token.balanceOf(address(s_lockReleaseTokenPool));

    s_oldLockReleaseTokenPool.setRebalancer(address(s_lockReleaseTokenPool));

    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.InsufficientBalance.selector, s_amount + 1, s_amount));
    s_lockReleaseTokenPool.transferLiquidity(address(s_oldLockReleaseTokenPool), s_amount + 1);

    assertEq(s_token.balanceOf(address(s_lockReleaseTokenPool)), balancePre);
  }
}
