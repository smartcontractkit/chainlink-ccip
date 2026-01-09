// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getRequiredCCVs is USDCTokenPoolProxySetup {
  function test_getRequiredCCVs_CCTPCCVRequired() public view {
    address[] memory requiredCCVs = s_usdcTokenPoolProxy.getRequiredCCVs(
      address(0), s_remoteCCTPChainSelector, 0, 0, "", IPoolV2.MessageDirection.Outbound
    );

    assertEq(requiredCCVs.length, 1);
    assertEq(requiredCCVs[0], s_cctpVerifier);
  }

  function test_getRequiredCCVs_DefaultCCVsRequired() public {
    // Set up a mechanism that doesn't require a CCV (e.g., CCTP_V1).
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;

    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    address[] memory requiredCCVs = s_usdcTokenPoolProxy.getRequiredCCVs(
      address(0), s_remoteLockReleaseChainSelector, 0, 0, "", IPoolV2.MessageDirection.Outbound
    );

    assertEq(requiredCCVs.length, 1);
    assertEq(requiredCCVs[0], address(0));
  }

  function test_getRequiredCCVs_RevertWhen_NoLockOrBurnMechanismSet() public {
    uint64 unknownChainSelector = 898989;
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.NoLockOrBurnMechanismSet.selector, unknownChainSelector));
    s_usdcTokenPoolProxy.getRequiredCCVs(address(0), unknownChainSelector, 0, 0, "", IPoolV2.MessageDirection.Outbound);
  }
}
