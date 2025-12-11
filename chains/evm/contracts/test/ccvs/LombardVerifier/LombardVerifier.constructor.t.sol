// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {LombardVerifierSetup, MockLombardBridge} from "./LombardVerifierSetup.t.sol";

contract LombardVerifier_constructor is LombardVerifierSetup {
  function test_constructor() public view {
    assertEq(address(s_lombardVerifier.i_bridge()), address(s_mockBridge));
  }

  function test_constructor_RevertWhen_ZeroBridge() public {
    vm.expectRevert(LombardVerifier.ZeroBridge.selector);
    new LombardVerifier(IBridgeV2(address(0)), s_storageLocations, address(s_mockRMNRemote));
  }

  function test_constructor_RevertWhen_InvalidMessageVersion() public {
    MockLombardBridge mockBridge = new MockLombardBridge();
    uint8 wrongVersion = 100;

    vm.mockCall(address(mockBridge), abi.encodeWithSelector(IBridgeV2.MSG_VERSION.selector), abi.encode(wrongVersion));

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidMessageVersion.selector, 1, wrongVersion));
    new LombardVerifier(IBridgeV2(address(mockBridge)), s_storageLocations, address(s_mockRMNRemote));
  }
}
