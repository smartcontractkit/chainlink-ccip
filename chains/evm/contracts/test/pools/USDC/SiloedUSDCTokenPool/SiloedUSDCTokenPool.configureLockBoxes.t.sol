// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SiloedUSDCTokenPool_configureLockBoxes is SiloedUSDCTokenPoolSetup {
  function test_configureLockBoxes() public {
    ERC20LockBox newSourceLockBox = new ERC20LockBox(address(s_USDCToken));
    ERC20LockBox newDestLockBox = new ERC20LockBox(address(s_USDCToken));

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = address(s_usdcTokenPool);

    newSourceLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );
    newDestLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](2);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR, lockBox: address(newSourceLockBox)
    });
    lockBoxes[1] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: DEST_CHAIN_SELECTOR, lockBox: address(newDestLockBox)
    });

    s_usdcTokenPool.configureLockBoxes(lockBoxes);

    assertEq(address(s_usdcTokenPool.getLockBox(SOURCE_CHAIN_SELECTOR)), address(newSourceLockBox));
    assertEq(address(s_usdcTokenPool.getLockBox(DEST_CHAIN_SELECTOR)), address(newDestLockBox));
  }

  function test_configureLockBoxes_RevertWhen_LockBoxCannotBeShared() public {
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](2);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR, lockBox: address(s_sourceLockBox)
    });
    lockBoxes[1] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: DEST_CHAIN_SELECTOR, lockBox: address(s_sourceLockBox)
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        SiloedUSDCTokenPool.LockBoxCannotBeShared.selector,
        SOURCE_CHAIN_SELECTOR,
        DEST_CHAIN_SELECTOR,
        address(s_sourceLockBox)
      )
    );
    s_usdcTokenPool.configureLockBoxes(lockBoxes);
  }
}
