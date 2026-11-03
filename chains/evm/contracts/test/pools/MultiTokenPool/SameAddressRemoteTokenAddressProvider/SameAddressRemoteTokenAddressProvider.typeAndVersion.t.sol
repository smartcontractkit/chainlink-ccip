// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {
  SameAddressRemoteTokenAddressProvider
} from "../../../../pools/MultiTokenPool/SameAddressRemoteTokenAddressProvider.sol";
import {Test} from "forge-std/Test.sol";

contract SameAddressRemoteTokenAddressProvider_typeAndVersion is Test {
  function test_typeAndVersion() public {
    SameAddressRemoteTokenAddressProvider provider = new SameAddressRemoteTokenAddressProvider();
    assertEq("SameAddressRemoteTokenAddressProvider 2.1.0-dev", provider.typeAndVersion());
  }
}
