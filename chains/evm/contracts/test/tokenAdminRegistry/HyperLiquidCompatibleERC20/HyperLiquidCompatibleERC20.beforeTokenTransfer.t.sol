// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_beforeTokenTransfer is HyperLiquidCompatibleERC20Setup {
  address public constant SPOT_BALANCE_PRECOMPILE = 0x0000000000000000000000000000000000000801;

  struct SpotBalance {
    uint64 total;
    uint64 hold;
    uint64 entryNtl;
  }

  function test_beforeTokenTransfer_Success_NonSpotBalancePrecompile() public {
    address recipient = makeAddr("recipient");
    uint256 amount = 100e18;

    // Transfer to a regular address should not trigger the spot balance check.
    s_hyperLiquidToken.transfer(recipient, amount);

    assertEq(s_hyperLiquidToken.balanceOf(recipient), amount);
  }

  function test_beforeTokenTransfer_SpotBalancePrecompile_WithSufficientBalance() public {
    uint256 amount = 100e18;

    // Since the amount of the remote token is 6 decimals, we need to have an amount that is greater than the local
    // amount when adjusted to 6 decimals.
    SpotBalance memory mockSpotBalance =
      SpotBalance({total: uint64(101e6), hold: uint64(500e6), entryNtl: uint64(200e6)});

    // Mock the staticcall to return success and the mock balance.
    vm.mockCall(
      SPOT_BALANCE_PRECOMPILE, abi.encode(s_hypercoreTokenSystemAddress, s_remoteTokenId), abi.encode(mockSpotBalance)
    );

    // Transfer to spot balance precompile should succeed with sufficient balance.
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, amount);

    assertEq(s_hyperLiquidToken.balanceOf(s_hypercoreTokenSystemAddress), amount);
  }

  function test_beforeTokenTransfer_Success_SpotBalancePrecompile_ExactBalance() public {
    uint256 amount = 100e18;

    // Mock the spot balance precompile to return exact balance needed
    SpotBalance memory mockSpotBalance =
      SpotBalance({total: uint64(100e6), hold: uint64(500e6), entryNtl: uint64(200e6)});

    vm.mockCall(
      SPOT_BALANCE_PRECOMPILE, abi.encode(s_hypercoreTokenSystemAddress, s_remoteTokenId), abi.encode(mockSpotBalance)
    );

    // Transfer to spot balance precompile should succeed with exact balance
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, amount);

    assertEq(s_hyperLiquidToken.balanceOf(s_hypercoreTokenSystemAddress), amount);
  }

  // Reverts

  function test_beforeTokenTransfer_RevertWhen_SpotBalancePrecompileCallFails() public {
    // Set the remote token first so that the second call reverts.
    s_hyperLiquidToken.setRemoteToken(s_remoteTokenId, 18);

    uint256 amount = 100e18;

    // Mock the staticcall to fail
    vm.mockCallRevert(SPOT_BALANCE_PRECOMPILE, bytes(""), "Mock failure");

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.HyperEVMTransferFailed.selector));
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, amount);
  }

  function test_beforeTokenTransfer_RevertWhen_InsufficientSpotBalance() public {
    s_hyperLiquidToken.setRemoteToken(s_remoteTokenId, 18);

    uint256 amount = 1001e18; // More than available balance

    // Mock the spot balance precompile to return insufficient balance.
    SpotBalance memory mockSpotBalance =
      SpotBalance({total: uint64(1000e6), hold: uint64(500e6), entryNtl: uint64(200e6)});

    vm.mockCall(
      SPOT_BALANCE_PRECOMPILE, abi.encode(s_hypercoreTokenSystemAddress, s_remoteTokenId), abi.encode(mockSpotBalance)
    );

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.InsufficientSpotBalance.selector));
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, amount);
  }

  function test_beforeTokenTransfer_RevertWhen_ZeroSpotBalance() public {
    // Set the remote token first so that the second call reverts.
    s_hyperLiquidToken.setRemoteToken(s_remoteTokenId, 18);

    uint256 amount = 100e18;

    // Mock the spot balance precompile to return zero balance
    SpotBalance memory mockSpotBalance = SpotBalance({total: uint64(0), hold: uint64(0), entryNtl: uint64(0)});

    vm.mockCall(
      SPOT_BALANCE_PRECOMPILE, abi.encode(s_hypercoreTokenSystemAddress, s_remoteTokenId), abi.encode(mockSpotBalance)
    );

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.InsufficientSpotBalance.selector));
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, amount);
  }

  function test_beforeTokenTransfer_RevertWhen_TransferZeroAmount() public {
    // Mock the spot balance precompile to return some balance.
    SpotBalance memory mockSpotBalance =
      SpotBalance({total: uint64(1000e6), hold: uint64(500e6), entryNtl: uint64(200e6)});

    vm.mockCall(
      SPOT_BALANCE_PRECOMPILE, abi.encode(s_hypercoreTokenSystemAddress, s_remoteTokenId), abi.encode(mockSpotBalance)
    );

    // Transfer zero amount should succeed.
    s_hyperLiquidToken.transfer(s_hypercoreTokenSystemAddress, 0);

    assertEq(s_hyperLiquidToken.balanceOf(s_hypercoreTokenSystemAddress), 0);
  }
}
