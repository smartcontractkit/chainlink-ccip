// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validate32ByteAddress is Test {
  function test_validate32ByteAddress_ValidAddress() public pure {
    bytes memory valid32ByteAddress =
      abi.encode(uint256(0x1234567890123456789012345678901234567890123456789012345678901234));
    Internal._validate32ByteAddress(valid32ByteAddress, 0);
  }

  function test_validate32ByteAddress_BoundaryMinValue() public pure {
    uint256 minValue = 100;
    bytes memory exactMinAddress = abi.encode(minValue);
    Internal._validate32ByteAddress(exactMinAddress, minValue);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validate32ByteAddress_RevertWhen_InvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, invalidAddress));
    Internal._validate32ByteAddress(invalidAddress, 0);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validate32ByteAddress_RevertWhen_AddressBelowMinValue() public {
    uint256 minValue = 100;
    bytes memory belowMinAddress = abi.encode(minValue - 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, belowMinAddress));
    Internal._validate32ByteAddress(belowMinAddress, minValue);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validate32ByteAddress_RevertWhen_AptosPrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(Internal.APTOS_PRECOMPILE_SPACE - 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, precompileAddress));
    Internal._validate32ByteAddress(precompileAddress, Internal.APTOS_PRECOMPILE_SPACE);
  }
}
