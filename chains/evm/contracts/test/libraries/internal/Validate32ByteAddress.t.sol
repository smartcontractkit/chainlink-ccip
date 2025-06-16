// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validate32ByteAddress is Test {
  function validate32ByteAddress(bytes memory encodedAddress, uint256 minValue) public pure {
    Internal._validate32ByteAddress(encodedAddress, minValue);
  }

  function test_validate32ByteAddress_succeeds_onValidAddress() public {
    bytes memory valid32ByteAddress =
      abi.encode(uint256(0x1234567890123456789012345678901234567890123456789012345678901234));
    this.validate32ByteAddress(valid32ByteAddress, 0);
  }

  function test_validate32ByteAddress_reverts_onInvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, invalidAddress));
    this.validate32ByteAddress(invalidAddress, 0);
  }

  function test_validate32ByteAddress_reverts_onAddressBelowMinValue() public {
    uint256 minValue = 1000;
    bytes memory belowMinAddress = abi.encode(minValue - 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, belowMinAddress));
    this.validate32ByteAddress(belowMinAddress, minValue);
  }

  function test_validate32ByteAddress_succeeds_onBoundaryMinValue() public {
    uint256 minValue = 1000;
    bytes memory exactMinAddress = abi.encode(minValue);
    this.validate32ByteAddress(exactMinAddress, minValue);
  }

  function test_validate32ByteAddress_reverts_onAptosPrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(Internal.APTOS_PRECOMPILE_SPACE - 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, precompileAddress));
    this.validate32ByteAddress(precompileAddress, Internal.APTOS_PRECOMPILE_SPACE);
  }
}
