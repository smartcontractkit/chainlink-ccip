// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validateTVMAddress is Test {
  function validateTVMAddress(
    bytes memory encoded
  ) public pure {
    Internal._validateTVMAddress(encoded);
  }

  function test_validateTVMAddress_ValidAddress() public {
    bytes memory validTvmAddress = hex"11ff1234567890123456789012345678901234567890123456789012345678901234abcd";
    this.validateTVMAddress(validTvmAddress);
  }

  function test_validateTVMAddress_RevertWhen_ShortLength() public {
    bytes memory shortAddress = new bytes(35);
    // Fill with non-zero data to avoid zero address check
    for (uint256 i = 0; i < 35; i++) {
      shortAddress[i] = 0x01;
    }
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, shortAddress));
    this.validateTVMAddress(shortAddress);
  }

  function test_validateTVMAddress_RevertWhen_LongLength() public {
    bytes memory longAddress = new bytes(37);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, longAddress));
    this.validateTVMAddress(longAddress);
  }

  function test_validateTVMAddress_MinimalNonZeroAccountId() public {
    bytes memory tvmAddress = new bytes(36);
    tvmAddress[0] = 0x00; // flags
    tvmAddress[1] = 0x00; // workchain_id
    tvmAddress[2] = 0x01; // non-zero account_id
    this.validateTVMAddress(tvmAddress);
  }

  function test_validateTVMAddress_RevertWhen_ZeroAccountId() public {
    bytes memory invalidTVMAddress = hex"11ff000000000000000000000000000000000000000000000000000000000000000012ab";
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, invalidTVMAddress));
    this.validateTVMAddress(invalidTVMAddress);
  }
}
