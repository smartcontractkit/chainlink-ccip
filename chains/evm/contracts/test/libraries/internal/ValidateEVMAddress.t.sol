// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validateEVMAddress is Test {
  function test_validateEVMAddress_ValidAddress() public pure {
    address validAddress = 0x1234567890123456789012345678901234567890;
    Internal._validateEVMAddress(abi.encode(validAddress));
  }

  function test_validateEVMAddress_BoundaryAddresses() public pure {
    bytes memory lowerBoundary = abi.encode(Internal.EVM_PRECOMPILE_SPACE);
    Internal._validateEVMAddress(lowerBoundary);
    bytes memory upperBoundary = abi.encode(type(uint160).max);
    Internal._validateEVMAddress(upperBoundary);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateEVMAddress_RevertWhen_InvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    Internal._validateEVMAddress(invalidAddress);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateEVMAddress_RevertWhen_PrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(address(0x01));
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, precompileAddress));
    Internal._validateEVMAddress(precompileAddress);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateEVMAddress_RevertWhen_OversizedAddress() public {
    bytes memory invalidAddress = abi.encode(uint256(type(uint160).max) + 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    Internal._validateEVMAddress(invalidAddress);
  }
}
