// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_updateLockReleasePoolAddresses is USDCTokenPoolProxySetup {
  address internal s_newLockReleasePool1 = makeAddr("newLockReleasePool1");
  address internal s_newLockReleasePool2 = makeAddr("newLockReleasePool2");
  address internal s_newLockReleasePool3 = makeAddr("newLockReleasePool3");
  uint64 internal s_remoteChainSelector1 = 1001;
  uint64 internal s_remoteChainSelector2 = 1002;
  uint64 internal s_remoteChainSelector3 = 1003;

  function test_updateLockReleasePoolAddresses_Single() public {
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteChainSelector1;

    address[] memory lockReleasePools = new address[](1);
    lockReleasePools[0] = s_newLockReleasePool1;

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

    address retrievedPool = s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector1);
    assertEq(retrievedPool, s_newLockReleasePool1);
  }

  // Test successful multiple lock release pool address updates by owner
  function test_updateLockReleasePoolAddresses_Multiple() public {
    uint64[] memory remoteChainSelectors = new uint64[](3);
    remoteChainSelectors[0] = s_remoteChainSelector1;
    remoteChainSelectors[1] = s_remoteChainSelector2;
    remoteChainSelectors[2] = s_remoteChainSelector3;

    address[] memory lockReleasePools = new address[](3);
    lockReleasePools[0] = s_newLockReleasePool1;
    lockReleasePools[1] = s_newLockReleasePool2;
    lockReleasePools[2] = s_newLockReleasePool3;

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector1), s_newLockReleasePool1);
    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector2), s_newLockReleasePool2);
    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector3), s_newLockReleasePool3);
  }

  // Reverts

  function test_updateLockReleasePoolAddresses_RevertWhen_MismatchedArrayLengths() public {
    // Arrange: Define test constants with mismatched lengths
    uint64[] memory remoteChainSelectors = new uint64[](2);
    remoteChainSelectors[0] = s_remoteChainSelector1;
    remoteChainSelectors[1] = s_remoteChainSelector2;

    address[] memory lockReleasePools = new address[](1); // Different length
    lockReleasePools[0] = s_newLockReleasePool1;

    // Act & Assert: Should revert with MismatchedArrayLengths error
    changePrank(OWNER);
    vm.expectRevert(USDCTokenPoolProxy.MismatchedArrayLengths.selector);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);
  }
}
