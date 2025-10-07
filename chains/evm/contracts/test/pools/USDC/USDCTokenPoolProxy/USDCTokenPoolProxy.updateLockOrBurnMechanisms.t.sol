// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_updateLockOrBurnMechanisms is USDCTokenPoolProxySetup {
  function test_updateLockOrBurnMechanisms() public {
    // Arrange: Define test constants
    uint64 testChainSelector = 12345;
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV1Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV2Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    USDCTokenPoolProxy.LockOrBurnMechanism lockReleaseMechanism = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    uint64[] memory chainSelectors = new uint64[](3);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    chainSelectors[1] = DEST_CHAIN_SELECTOR;
    chainSelectors[2] = testChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](3);
    mechanisms[0] = cctpV1Mechanism;
    mechanisms[1] = cctpV2Mechanism;
    mechanisms[2] = lockReleaseMechanism;

    changePrank(OWNER);
    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(SOURCE_CHAIN_SELECTOR, cctpV1Mechanism);
    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(DEST_CHAIN_SELECTOR, cctpV2Mechanism);
    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(testChainSelector, lockReleaseMechanism);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(SOURCE_CHAIN_SELECTOR)), uint8(cctpV1Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(DEST_CHAIN_SELECTOR)), uint8(cctpV2Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(testChainSelector)), uint8(lockReleaseMechanism));
  }

  // Reverts

  function test_updateLockOrBurnMechanisms_RevertWhen_MismatchedArrayLengths() public {
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](2);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;

    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.MismatchedArrayLengths.selector));
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }
}
