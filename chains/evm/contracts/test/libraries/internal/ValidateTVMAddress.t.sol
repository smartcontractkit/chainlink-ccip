// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_validateTVMAddress is Test {
  function test_validateTVMAddress_MasterchainAddress() public pure {
    // Ef9nROksb3HHdvu87ymeMb9285wkXNVvIHW4nGoiAmtBMZG9, -1:6744e92c6f71c776fbbcef299e31bf76f39c245cd56f2075b89c6a22026b4131
    bytes memory masterchainAddress = hex"11ff6744e92c6f71c776fbbcef299e31bf76f39c245cd56f2075b89c6a22026b413191bd";
    Internal._validateTVMAddress(masterchainAddress);
  }

  function test_validateTVMAddress_BasechainAddress() public pure {
    // EQAdp38Cabu7dshi6kJLJX32O9GssNTraBtoyarfv1U7kzWP, 0:1da77f0269bbbb76c862ea424b257df63bd1acb0d4eb681b68c9aadfbf553b93
    bytes memory basechainAddress = hex"11001da77f0269bbbb76c862ea424b257df63bd1acb0d4eb681b68c9aadfbf553b93358f";
    Internal._validateTVMAddress(basechainAddress);
  }

  function test_validateTVMAddress_MinimalNonZeroAccountId() public pure {
    bytes memory tvmAddress = new bytes(36);
    tvmAddress[0] = 0x00; // flags
    tvmAddress[1] = 0x00; // workchain_id
    tvmAddress[2] = 0x01; // non-zero account_id
    Internal._validateTVMAddress(tvmAddress);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateTVMAddress_RevertWhen_ShortLength() public {
    bytes memory shortAddress = new bytes(35);
    shortAddress[2] = 0x01; // account_id as non-zero to isolate length check
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, shortAddress));
    Internal._validateTVMAddress(shortAddress);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateTVMAddress_RevertWhen_LongLength() public {
    bytes memory longAddress = new bytes(37);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, longAddress));
    Internal._validateTVMAddress(longAddress);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test_validateTVMAddress_RevertWhen_ZeroAccountId() public {
    bytes memory invalidTVMAddress = hex"11ff000000000000000000000000000000000000000000000000000000000000000012ab";
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, invalidTVMAddress));
    Internal._validateTVMAddress(invalidTVMAddress);
  }
}
