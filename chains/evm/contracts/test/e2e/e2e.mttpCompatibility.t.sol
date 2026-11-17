// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";

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
    (bytes memory poolData, uint256 postFee,) =
      _send(s_a, s_b, 0, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    uint256 received = _deliverFromPool(
      s_b, s_a.localSelector, address(s_a.pool), dest6, postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, amount / 1e12, "18 to 6 floor scaling");
    assertEq(_balance(dest6, s_receiver), received, "exact minted amount");
  }

  /// @dev Each supported burn flavor moves the exact amount off the source pool. BurnToAddress sends it to the
  /// configured burn address instead of burning it.
  function test_mttp_burnFlavorsMoveExactAmount() public {
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "Burn flavor");

    vm.prank(OWNER);
    _setProfile(s_a.pcm, s_a.tokens[0], T1_DECIMALS, IPoolConfigManager.LockOrBurnFlavor.BurnFrom);
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "BurnFrom flavor");

    vm.prank(OWNER);
    _setProfileFull(
      s_a.pcm, s_a.tokens[0], T1_DECIMALS, IPoolConfigManager.LockOrBurnFlavor.BurnToAddress, address(0xdead)
    );
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "BurnToAddress flavor");
    assertEq(_balance(s_a.tokens[0], address(0xdead)), 10e18, "BurnToAddress moved the amount to the burn address");
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
