// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../../pools/ERC20LockBox.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract SiloedUSDCTokenPoolSetup is USDCSetup {
  SiloedUSDCTokenPool internal s_usdcTokenPool;
  SiloedUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;

  address internal s_lockBox;

  function setUp() public virtual override {
    super.setUp();

    s_lockBox = address(new ERC20LockBox(address(s_tokenAdminRegistry)));

    // Mock the isAdministrator function to return true so that the owner can configure allowed callers for the lock box.
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("isAdministrator(address,address)", address(s_USDCToken), OWNER),
      abi.encode(true)
    );

    s_usdcTokenPool = new SiloedUSDCTokenPool(
      s_USDCToken,
      6, // localTokenDecimals
      new address[](0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_router), // router
      s_lockBox // lockBox
    );

    _poolApplyChainUpdates(address(s_usdcTokenPool));

    // Mock the getPool function to return the address of the token pool
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getPool(address)", address(s_USDCToken)),
      abi.encode(address(s_usdcTokenPool))
    );

    // Allow the router to call the releaseOrMint function for the token pool 
    ERC20LockBox.AllowedCallerConfigArgs[] memory allowedCallers = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    allowedCallers[0] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_USDCToken),
      caller: address(s_routerAllowedOffRamp),
      allowed: true
    });
    allowedCallers[1] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_USDCToken),
      caller: address(s_routerAllowedOffRamp),
      allowed: true
    });
    ERC20LockBox(s_lockBox).configureAllowedCallers(allowedCallers);

    // Allow the router to call the releaseOrMint function
    s_usdcTokenPoolTransferLiquidity = new SiloedUSDCTokenPool(
      s_USDCToken,
      6, // localTokenDecimals
      new address[](0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_router), // router
      s_lockBox // lockBox
    );
  }
}
