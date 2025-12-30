// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract SiloedUSDCTokenPoolSetup is USDCSetup {
  SiloedUSDCTokenPool internal s_usdcTokenPool;
  SiloedUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;

  ERC20LockBox internal s_lockBox;
  ERC20LockBox internal s_sourceLockBox;
  ERC20LockBox internal s_destLockBox;

  function setUp() public virtual override {
    super.setUp();

    s_lockBox = new ERC20LockBox(address(s_USDCToken));
    s_sourceLockBox = new ERC20LockBox(address(s_USDCToken));
    s_destLockBox = new ERC20LockBox(address(s_USDCToken));

    s_usdcTokenPool = new SiloedUSDCTokenPool(
      s_USDCToken,
      6, // localTokenDecimals
      address(0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_router), // router
      address(s_lockBox) // lockBox
    );

    address[] memory authorizedCallers = new address[](3);
    authorizedCallers[0] = OWNER;
    authorizedCallers[1] = address(s_routerAllowedOnRamp);
    authorizedCallers[2] = address(s_routerAllowedOffRamp);
    s_usdcTokenPool.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_usdcTokenPool));

    s_tokenAdminRegistry.proposeAdministrator(address(s_USDCToken), OWNER);
    s_tokenAdminRegistry.acceptAdminRole(address(s_USDCToken));
    s_tokenAdminRegistry.setPool(address(s_USDCToken), address(s_usdcTokenPool));

    _poolApplyChainUpdates(address(s_usdcTokenPool));

    // Mock the getPool function to return the address of the token pool
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getPool(address)", address(s_USDCToken)),
      abi.encode(address(s_usdcTokenPool))
    );

    // Allow the router to call the releaseOrMint function
    s_usdcTokenPoolTransferLiquidity = new SiloedUSDCTokenPool(
      s_USDCToken,
      6, // localTokenDecimals
      address(0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_router), // router
      address(s_lockBox) // lockBox
    );

    s_usdcTokenPoolTransferLiquidity.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    _poolApplyChainUpdates(address(s_usdcTokenPoolTransferLiquidity));

    // Allow both pools to interact with the lockboxes.
    address[] memory allowedCallers = new address[](2);
    allowedCallers[0] = address(s_usdcTokenPool);
    allowedCallers[1] = address(s_usdcTokenPoolTransferLiquidity);
    ERC20LockBox(s_lockBox)
      .applyAuthorizedCallerUpdates(
        AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
      );
    ERC20LockBox(s_sourceLockBox)
      .applyAuthorizedCallerUpdates(
        AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
      );
    ERC20LockBox(s_destLockBox)
      .applyAuthorizedCallerUpdates(
        AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
      );

    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](2);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR, lockBox: address(s_sourceLockBox)
    });
    lockBoxes[1] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: DEST_CHAIN_SELECTOR, lockBox: address(s_destLockBox)
    });
    s_usdcTokenPool.configureLockBoxes(lockBoxes);
    s_usdcTokenPoolTransferLiquidity.configureLockBoxes(lockBoxes);
  }
}
