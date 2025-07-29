// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_setAllowedTokenPoolProxies is SiloedUSDCTokenPoolSetup {
  address PROXY_1 = makeAddr("PROXY_1");
  address PROXY_2 = makeAddr("PROXY_2");
  address PROXY_3 = makeAddr("PROXY_3");

  function test_setAllowedTokenPoolProxies_Success() public {
    // Arrange: Prepare arrays for setting allowed proxies
    address[] memory proxies = new address[](2);
    proxies[0] = PROXY_1;
    proxies[1] = PROXY_2;
    
    bool[] memory allowed = new bool[](2);
    allowed[0] = true;
    allowed[1] = true;

    // Act: Set allowed token pool proxies
    vm.startPrank(OWNER);
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies, allowed);
    vm.stopPrank();

    // Assert: Verify proxies are now allowed
    address[] memory allowedProxies = s_usdcTokenPool.getAllowedTokenPoolProxies();
    assertTrue(_containsAddress(allowedProxies, PROXY_1));
    assertTrue(_containsAddress(allowedProxies, PROXY_2));
  }

  function test_setAllowedTokenPoolProxies_RevertWhen_NotOwner() public {
    // Arrange: Prepare arrays
    address[] memory proxies = new address[](1);
    proxies[0] = PROXY_1;
    
    bool[] memory allowed = new bool[](1);
    allowed[0] = true;

    // Act & Assert: Expect revert when non-owner tries to set proxies
    vm.startPrank(STRANGER);
    vm.expectRevert();
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies, allowed);
    vm.stopPrank();
  }

  function test_setAllowedTokenPoolProxies_RevertWhen_MismatchedArrayLengths() public {
    // Arrange: Arrays with different lengths
    address[] memory proxies = new address[](2);
    proxies[0] = PROXY_1;
    proxies[1] = PROXY_2;
    
    bool[] memory allowed = new bool[](1);
    allowed[0] = true;

    // Act & Assert: Expect revert due to mismatched array lengths
    vm.startPrank(OWNER);
    vm.expectRevert();
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies, allowed);
    vm.stopPrank();
  }

  function test_setAllowedTokenPoolProxies_RevertWhen_ProxyAlreadyAllowed() public {
    // Arrange: First set a proxy as allowed
    address[] memory proxies1 = new address[](1);
    proxies1[0] = PROXY_1;
    
    bool[] memory allowed1 = new bool[](1);
    allowed1[0] = true;

    vm.startPrank(OWNER);
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies1, allowed1);

    // Act & Assert: Try to add the same proxy again
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.TokenPoolProxyAlreadyAllowed.selector, PROXY_1));
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies1, allowed1);
    vm.stopPrank();
  }

  function test_setAllowedTokenPoolProxies_RevertWhen_ProxyNotAllowed() public {
    // Arrange: Try to remove a proxy that was never added
    address[] memory proxies = new address[](1);
    proxies[0] = PROXY_1;
    
    bool[] memory allowed = new bool[](1);
    allowed[0] = false;

    // Act & Assert: Expect revert when trying to remove non-existent proxy
    vm.startPrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.TokenPoolProxyNotAllowed.selector, PROXY_1));
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies, allowed);
    vm.stopPrank();
  }

  function test_setAllowedTokenPoolProxies_RemoveProxy() public {
    // Arrange: First add a proxy
    address[] memory addProxies = new address[](1);
    addProxies[0] = PROXY_1;
    
    bool[] memory addAllowed = new bool[](1);
    addAllowed[0] = true;

    vm.startPrank(OWNER);
    s_usdcTokenPool.setAllowedTokenPoolProxies(addProxies, addAllowed);

    // Verify proxy was added
    address[] memory allowedProxiesBefore = s_usdcTokenPool.getAllowedTokenPoolProxies();
    assertTrue(_containsAddress(allowedProxiesBefore, PROXY_1));

    // Act: Remove the proxy
    address[] memory removeProxies = new address[](1);
    removeProxies[0] = PROXY_1;
    
    bool[] memory removeAllowed = new bool[](1);
    removeAllowed[0] = false;

    s_usdcTokenPool.setAllowedTokenPoolProxies(removeProxies, removeAllowed);
    vm.stopPrank();

    // Assert: Verify proxy was removed
    address[] memory allowedProxiesAfter = s_usdcTokenPool.getAllowedTokenPoolProxies();
    assertFalse(_containsAddress(allowedProxiesAfter, PROXY_1));
  }

  function test_setAllowedTokenPoolProxies_EmitsEvents() public {
    // Arrange: Prepare arrays
    address[] memory proxies = new address[](2);
    proxies[0] = PROXY_1;
    proxies[1] = PROXY_2;
    
    bool[] memory allowed = new bool[](2);
    allowed[0] = true;
    allowed[1] = true;

    // Act & Assert: Expect events to be emitted
    vm.startPrank(OWNER);
    vm.expectEmit(true, false, false, false);
    emit SiloedUSDCTokenPool.AllowedTokenPoolProxyAdded(PROXY_1);
    vm.expectEmit(true, false, false, false);
    emit SiloedUSDCTokenPool.AllowedTokenPoolProxyAdded(PROXY_2);
    s_usdcTokenPool.setAllowedTokenPoolProxies(proxies, allowed);
    vm.stopPrank();
  }

  /// @notice Helper function to check if an address is in an array
  function _containsAddress(address[] memory addresses, address target) internal pure returns (bool) {
    for (uint256 i = 0; i < addresses.length; i++) {
      if (addresses[i] == target) {
        return true;
      }
    }
    return false;
  }
} 