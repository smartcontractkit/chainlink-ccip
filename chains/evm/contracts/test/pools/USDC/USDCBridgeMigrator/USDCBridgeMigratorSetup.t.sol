// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../../pools/ERC20LockBox.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCBridgeMigratorSetup is USDCSetup {
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPool;
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;

  address internal s_lockBox;

  function setUp() public virtual override {
    super.setUp();

    s_lockBox = address(new ERC20LockBox(address(s_tokenAdminRegistry)));

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool,
      s_lockBox
    );

    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getPool(address)", address(s_USDCToken)),
      abi.encode(address(s_usdcTokenPool))
    );

    ERC20LockBox.AllowedCallerConfigArgs[] memory allowedCallers = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    allowedCallers[0] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_USDCToken),
      caller: address(s_usdcTokenPool),
      allowed: true
    });
    allowedCallers[1] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_USDCToken),
      caller: address(s_usdcTokenPoolTransferLiquidity),
      allowed: true
    });
    ERC20LockBox(s_lockBox).configureAllowedCallers(allowedCallers);

    s_usdcTokenPoolTransferLiquidity = new HybridLockReleaseUSDCTokenPool(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool,
      s_lockBox
    );
  }
}
