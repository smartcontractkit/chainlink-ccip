// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {
  SameAddressRemoteTokenAddressProvider
} from "../../../../pools/MultiTokenPool/SameAddressRemoteTokenAddressProvider.sol";
import {Test} from "forge-std/Test.sol";

contract SameAddressRemoteTokenAddressProvider_getRemoteTokenAddress is Test {
  SameAddressRemoteTokenAddressProvider internal s_provider;

  function setUp() public {
    s_provider = new SameAddressRemoteTokenAddressProvider();
  }

  function testFuzz_getRemoteTokenAddress(
    address localToken,
    uint64 remoteChainSelector
  ) public view {
    assertEq(abi.encode(localToken), s_provider.getRemoteTokenAddress(localToken, remoteChainSelector));
  }
}
