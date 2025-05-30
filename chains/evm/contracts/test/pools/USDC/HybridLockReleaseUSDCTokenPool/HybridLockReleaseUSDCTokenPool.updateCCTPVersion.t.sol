// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {HybridLockReleaseUSDCTokenPoolSetup} from "./HybridLockReleaseUSDCTokenPoolSetup.t.sol";

contract HybridLockReleaseUSDCTokenPool_updateCCTPVersion is HybridLockReleaseUSDCTokenPoolSetup {
  // Reverts
  function test_RevertWhen_MismatchedArrayLengths() public {
    uint64[] memory chainSelectors = new uint64[](2);
    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);

    vm.startPrank(OWNER);

    vm.expectRevert(TokenPool.MismatchedArrayLengths.selector);

    s_usdcTokenPool.updateCCTPVersion(chainSelectors, versions);
  }

  function test_RevertWhen_ChainNotSupportedByCCTP() public {
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = DEST_CHAIN_SELECTOR;

    // Mark the chain as using Lock/Release instead of CCTP
    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), chainSelectors);

    HybridLockReleaseUSDCTokenPool.CCTPVersion[] memory versions = new HybridLockReleaseUSDCTokenPool.CCTPVersion[](1);

    vm.startPrank(OWNER);

    // Expect to Revert because the chain is not compatible with CCTP already.
    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.ChainNotSupportedByCCTP.selector, DEST_CHAIN_SELECTOR)
    );

    s_usdcTokenPool.updateCCTPVersion(chainSelectors, versions);
  }
}
