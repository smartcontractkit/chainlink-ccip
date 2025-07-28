// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

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
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    assertEq(uint8(s_usdcTokenPoolProxy.s_lockOrBurnMechanism(SOURCE_CHAIN_SELECTOR)), uint8(cctpV1Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.s_lockOrBurnMechanism(DEST_CHAIN_SELECTOR)), uint8(cctpV2Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.s_lockOrBurnMechanism(testChainSelector)), uint8(lockReleaseMechanism));
  }

  function test_updateLockOrBurnMechanisms_RevertWhen_NotOwner() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV1Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;

    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = cctpV1Mechanism;

    changePrank(makeAddr("notOwner"));
    vm.expectRevert();
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_updateLockOrBurnMechanisms_EmitsEvents() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV1Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV2Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;

    uint64[] memory chainSelectors = new uint64[](2);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    chainSelectors[1] = DEST_CHAIN_SELECTOR;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](2);
    mechanisms[0] = cctpV1Mechanism;
    mechanisms[1] = cctpV2Mechanism;

    changePrank(OWNER);
    vm.expectEmit(true, true, true, true);
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(SOURCE_CHAIN_SELECTOR, cctpV1Mechanism);
    vm.expectEmit(true, true, true, true);
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(DEST_CHAIN_SELECTOR, cctpV2Mechanism);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_updateLockOrBurnMechanisms_OverwriteExisting() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV1Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    USDCTokenPoolProxy.LockOrBurnMechanism lockReleaseMechanism = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    // First, set some initial mechanisms
    uint64[] memory initialChainSelectors = new uint64[](1);
    initialChainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory initialMechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    initialMechanisms[0] = cctpV1Mechanism;

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(initialChainSelectors, initialMechanisms);
    assertEq(uint8(s_usdcTokenPoolProxy.s_lockOrBurnMechanism(SOURCE_CHAIN_SELECTOR)), uint8(cctpV1Mechanism));

    // Now update the same chain selector with a different mechanism
    uint64[] memory updatedChainSelectors = new uint64[](1);
    updatedChainSelectors[0] = SOURCE_CHAIN_SELECTOR;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory updatedMechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    updatedMechanisms[0] = lockReleaseMechanism;

    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(updatedChainSelectors, updatedMechanisms);
    assertEq(uint8(s_usdcTokenPoolProxy.s_lockOrBurnMechanism(SOURCE_CHAIN_SELECTOR)), uint8(lockReleaseMechanism));
  }

  function test_updateLockOrBurnMechanisms_EmptyArrays() public {
    // Arrange: Define test constants
    uint64[] memory chainSelectors = new uint64[](0);
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](0);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
    // Should not revert and should not change any state
  }
} 