// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";
import {IBridgeV3} from "../../../interfaces/lombard/IBridgeV3.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {LombardVerifierSetup, MockLombardBridge} from "./LombardVerifierSetup.t.sol";

contract LombardVerifier_constructor is LombardVerifierSetup {
  function test_constructor() public view {
    assertEq(address(s_lombardVerifier.i_bridge()), address(s_mockBridge));
    assertEq(s_lombardVerifier.versionTag(), VERSION_TAG_V2_0_0);
  }

  function test_constructor_RevertWhen_ZeroBridge() public {
    vm.expectRevert(LombardVerifier.ZeroBridge.selector);
    new LombardVerifier(
      LombardVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      IBridgeV3(address(0)),
      s_storageLocations,
      address(s_mockRMNRemote),
      VERSION_TAG_V2_0_0
    );
  }

  function test_constructor_RevertWhen_InvalidMessageVersion() public {
    MockLombardBridge mockBridge = new MockLombardBridge();
    uint8 wrongVersion = 100;

    vm.mockCall(address(mockBridge), abi.encodeWithSelector(IBridgeV2.MSG_VERSION.selector), abi.encode(wrongVersion));

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidMessageVersion.selector, 2, wrongVersion));
    new LombardVerifier(
      LombardVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      IBridgeV3(address(mockBridge)),
      s_storageLocations,
      address(s_mockRMNRemote),
      VERSION_TAG_V2_0_0
    );
  }

  function test_constructor_RevertWhen_VersionTagCannotBeZero() public {
    vm.expectRevert(BaseVerifier.VersionTagCannotBeZero.selector);
    new LombardVerifier(
      LombardVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      IBridgeV3(address(s_mockBridge)),
      s_storageLocations,
      address(s_mockRMNRemote),
      bytes4(0)
    );
  }
}
