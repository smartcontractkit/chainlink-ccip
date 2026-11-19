// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Compatibility behavior: token decimals, burn flavors, and empty legacy pool data.
contract e2e_mttpCompatibility is e2e_mttpSetup {
  /// @dev An 18 decimal source to a 6 decimal destination rounds down. The destination mints the floored amount
  /// and the dust is dropped, not carried into the balance.
  function test_mttp_decimalScalingRoundsDown() public {
    vm.startPrank(OWNER);
    // Deploy a 6 decimal destination token and point A's T1 remote mapping at it.
    address dest6 = _deployExtraToken(s_b, "T1b6", 6);
    _registerExtraToken(s_b, dest6, 6, s_a.tokens[0]);
    s_a.remoteTokenProvider.setRemoteToken(s_a.tokens[0], s_b.localSelector, abi.encode(dest6));
    vm.stopPrank();

    uint256 amount = 1_234_567_890_123_456_789; // 1.234567890123456789 with 18 decimals.
    (bytes memory poolData, uint256 postFee, bytes memory destTokenAddress) =
      _send(s_a, s_b, 0, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    uint256 received = _deliverFromPool(
      s_b,
      s_a.localSelector,
      address(s_a.pool),
      abi.decode(destTokenAddress, (address)),
      postFee,
      poolData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, amount / 1e12, "18 to 6 floor scaling");
    assertEq(_balance(dest6, s_receiver), received, "exact minted amount");
  }

  /// @dev Each supported burn flavor moves the exact amount off the source pool and mints it on the destination.
  /// The source supply drops for the three burn flavors and stays flat for BurnToAddress, which moves the amount to
  /// an address that can never spend it.
  function test_mttp_burnFlavorsMoveExactAmount() public {
    _assertBurnFlavor(IPoolConfigManager.LockOrBurnFlavor.Burn, address(0), 10e18);
    _assertBurnFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnFrom, address(0), 10e18);
    _assertBurnFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnWithFrom, address(0), 10e18);
    _assertBurnFlavor(IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, address(0xdead), 0);
  }

  /// @dev Transfers 10e18 with one flavor and checks the source pool balance, the source supply drop, and where the
  /// amount went: burned for the burn flavors, held by the burn address for BurnToAddress.
  function _assertBurnFlavor(
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address burnAddress,
    uint256 expectedSupplyDrop
  ) internal {
    vm.prank(OWNER);
    if (burnAddress == address(0)) {
      _setProfile(s_a.pcm, s_a.tokens[0], T1_DECIMALS, flavor);
    } else {
      _setProfileFull(s_a.pcm, s_a.tokens[0], T1_DECIMALS, flavor, burnAddress);
    }

    address token = s_a.tokens[0];
    uint256 supplyBefore = IERC20(token).totalSupply();
    uint256 poolBefore = _balance(token, address(s_a.pool));
    uint256 burnBefore = _balance(token, burnAddress);

    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "destination amount");
    assertEq(_balance(token, address(s_a.pool)), poolBefore, "source pool balance unchanged");
    assertEq(supplyBefore - IERC20(token).totalSupply(), expectedSupplyDrop, "source supply drop");
    assertEq(_balance(token, burnAddress) - burnBefore, 10e18 - expectedSupplyDrop, "burn address balance");
  }

  /// @dev A migrated token still accepts empty legacy pool data from a pool it trusts.
  function test_mttp_emptyLegacyPoolDataAccepted() public {
    vm.prank(OWNER);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, address(s_b.pool), false, true);

    uint256 received = _deliverFromPool(
      s_a, s_b.localSelector, address(s_b.pool), s_a.tokens[0], 10e18, "", FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 10e18, "empty legacy pool data accepted when allowed");
  }
}
