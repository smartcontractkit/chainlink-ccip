// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validateEVMAddress is Test {
  function validateEVMAddress(
    bytes memory encoded
  ) public pure {
    Internal._validateEVMAddress(encoded);
  }

  function test_validateEVMAddress_ValidAddress() public {
    address validAddress = 0x1234567890123456789012345678901234567890;
    this.validateEVMAddress(abi.encode(validAddress));
  }

  function test_validateEVMAddress_RevertWhen_InvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    this.validateEVMAddress(invalidAddress);
  }

  function test_validateEVMAddress_RevertWhen_PrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(address(0x01));
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, precompileAddress));
    this.validateEVMAddress(precompileAddress);
  }

  function test_validateEVMAddress_RevertWhen_OversizedAddress() public {
    bytes memory invalidAddress = abi.encode(uint256(type(uint160).max) + 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    this.validateEVMAddress(invalidAddress);
  }

  function test_validateEVMAddress_BoundaryAddresses() public {
    bytes memory lowerBoundary = abi.encode(Internal.EVM_PRECOMPILE_SPACE);
    this.validateEVMAddress(lowerBoundary);

    bytes memory upperBoundary = abi.encode(type(uint160).max);
    this.validateEVMAddress(upperBoundary);
  }
}
