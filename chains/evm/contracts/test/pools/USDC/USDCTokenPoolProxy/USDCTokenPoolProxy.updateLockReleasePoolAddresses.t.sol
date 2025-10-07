// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

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

    _enableERC165InterfaceChecks(s_newLockReleasePool1, type(IPoolV1).interfaceId);

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

    _enableERC165InterfaceChecks(s_newLockReleasePool1, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newLockReleasePool2, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newLockReleasePool3, type(IPoolV1).interfaceId);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector1), s_newLockReleasePool1);
    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector2), s_newLockReleasePool2);
    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector3), s_newLockReleasePool3);
  }

  function test_updateLockReleasePoolAddresses_ZeroAddress() public {
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteChainSelector1;
    address[] memory lockReleasePools = new address[](1);
    lockReleasePools[0] = address(0);

    // Since the address is address(0) the IERC165 check is not performed and no revert should occur.
    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

    assertEq(s_usdcTokenPoolProxy.getLockReleasePoolAddress(s_remoteChainSelector1), address(0));
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

  function test_updateLockReleasePoolAddresses_RevertWhen_V1PoolDoesNotSupportIPoolV1() public {
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteChainSelector1;
    address[] memory lockReleasePools = new address[](1);
    lockReleasePools[0] = s_newLockReleasePool1;

    changePrank(OWNER);
    // Should revert because the pool does not support the IPoolV1 interface
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newLockReleasePool1));
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);

    // Should not revert because the pool supports the IPoolV1 interface
    _enableERC165InterfaceChecks(s_newLockReleasePool1, type(IPoolV1).interfaceId);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(remoteChainSelectors, lockReleasePools);
  }

  function _enableERC165InterfaceChecks(address pool, bytes4 interfaceId) internal {
    vm.mockCall(
      address(pool), abi.encodeWithSelector(IERC165.supportsInterface.selector, interfaceId), abi.encode(true)
    );

    vm.mockCall(
      address(pool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId),
      abi.encode(true)
    );
  }
}
