// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPool} from "../../../../interfaces/IMultiTokenPool.sol";

import {FeeTokenHandler} from "../../../../libraries/FeeTokenHandler.sol";
import {BurnMintMultiTokenPool} from "../../../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract PoolConfigManager_withdrawFeeTokens is PoolConfigManagerTransferSetup {
  address internal s_feeRecipient = makeAddr("feeRecipient");

  /// @notice Withdraws each fee token's full balance from a pool configured to use this manager.
  function test_withdrawFeeTokens() public {
    deal(s_tokenB, address(s_pool), 5e18, true);
    _setCaller(_grantNewCaller(s_poolConfigManager.FEE_COLLECTOR_ROLE()));

    vm.expectEmit(address(s_pool));
    emit FeeTokenHandler.FeeTokenWithdrawn(s_feeRecipient, s_tokenB, 5e18);
    s_poolConfigManager.withdrawFeeTokens(IMultiTokenPool(address(s_pool)), _arrayOf(s_tokenB), s_feeRecipient);

    assertEq(5e18, IERC20(s_tokenB).balanceOf(s_feeRecipient));
    assertEq(0, IERC20(s_tokenB).balanceOf(address(s_pool)));
  }

  // Reverts

  /// @notice Only FEE_COLLECTOR_ROLE can withdraw fees. Other roles, including DEFAULT_ADMIN_ROLE, and callers
  /// without roles are rejected.
  function test_withdrawFeeTokens_RevertWhen_MissingRole() public {
    bytes32[] memory allowedRoles = new bytes32[](1);
    allowedRoles[0] = s_poolConfigManager.FEE_COLLECTOR_ROLE();

    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.withdrawFeeTokens, (IMultiTokenPool(address(s_pool)), _arrayOf(s_tokenB), s_feeRecipient)
      ),
      allowedRoles
    );
  }

  /// @notice Rejects a pool configured to use another manager before calling withdrawFeeTokens. The error
  /// identifies both the pool and its manager, and the pool keeps its fees.
  function test_withdrawFeeTokens_RevertWhen_PoolConfigManagerMismatch() public {
    address otherPoolConfigManager = makeAddr("otherPoolConfigManager");
    BurnMintMultiTokenPool otherPool =
      new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(s_sourceRouter), otherPoolConfigManager);
    deal(s_tokenB, address(otherPool), 5e18, true);
    _setCaller(_grantNewCaller(s_poolConfigManager.FEE_COLLECTOR_ROLE()));

    vm.expectRevert(
      abi.encodeWithSelector(
        PoolConfigManager.PoolConfigManagerMismatch.selector, address(otherPool), otherPoolConfigManager
      )
    );
    s_poolConfigManager.withdrawFeeTokens(IMultiTokenPool(address(otherPool)), _arrayOf(s_tokenB), s_feeRecipient);

    assertEq(5e18, IERC20(s_tokenB).balanceOf(address(otherPool)));
  }
}
