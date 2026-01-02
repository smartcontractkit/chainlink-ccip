// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_updateLockOrBurnMechanisms is USDCTokenPoolProxySetup {
  function test_updateLockOrBurnMechanisms() public {
    uint64 testChainSelector = 12345;
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV1Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    USDCTokenPoolProxy.LockOrBurnMechanism cctpV2Mechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    USDCTokenPoolProxy.LockOrBurnMechanism ccvMechanism = USDCTokenPoolProxy.LockOrBurnMechanism.CCV;

    uint64[] memory chainSelectors = new uint64[](3);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    chainSelectors[1] = DEST_CHAIN_SELECTOR;
    chainSelectors[2] = testChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](3);
    mechanisms[0] = cctpV1Mechanism;
    mechanisms[1] = cctpV2Mechanism;
    mechanisms[2] = ccvMechanism;

    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(SOURCE_CHAIN_SELECTOR, cctpV1Mechanism);
    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(DEST_CHAIN_SELECTOR, cctpV2Mechanism);
    vm.expectEmit();
    emit USDCTokenPoolProxy.LockOrBurnMechanismUpdated(testChainSelector, ccvMechanism);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(SOURCE_CHAIN_SELECTOR)), uint8(cctpV1Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(DEST_CHAIN_SELECTOR)), uint8(cctpV2Mechanism));
    assertEq(uint8(s_usdcTokenPoolProxy.getLockOrBurnMechanism(testChainSelector)), uint8(ccvMechanism));
  }

  // Reverts

  function test_updateLockOrBurnMechanisms_RevertWhen_MismatchedArrayLengths() public {
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](2);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.MismatchedArrayLengths.selector));
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_updateLockOrBurnMechanisms_RevertWhen_MustSetPoolForMechanism_CCTPV1PoolNotSet() public {
    uint64 testChainSelector = 99999;

    // Remove the CCTP V1 pool.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = testChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        testChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1
      )
    );
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_updateLockOrBurnMechanisms_RevertWhen_MustSetPoolForMechanism_CCTPV2PoolNotSet() public {
    uint64 testChainSelector = 99999;

    // Remove the CCTP V2 pool.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = testChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        testChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2
      )
    );
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_updateLockOrBurnMechanisms_RevertWhen_MustSetPoolForMechanism_CCVPoolNotSet() public {
    uint64 testChainSelector = 99999;

    // Remove the CCTP V2 with CCV pool.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = testChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCV;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        testChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.CCV
      )
    );
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }
}
