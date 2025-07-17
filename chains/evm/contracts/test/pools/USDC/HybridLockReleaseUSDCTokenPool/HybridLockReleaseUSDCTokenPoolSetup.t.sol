// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintERC677} from "@chainlink/contracts/src/v0.8/shared/token/ERC677/BurnMintERC677.sol";

import {ERC20LockBox} from "../../../../pools/ERC20LockBox.sol";
import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract HybridLockReleaseUSDCTokenPoolSetup is USDCSetup {
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPool;
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;
  CCTPMessageTransmitterProxy internal s_cctpMessageTransmitterProxyForTransferLiquidity;
  address[] internal s_allowedList;

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

    s_tokenAdminRegistry.proposeAdministrator(address(s_USDCToken), address(OWNER));
    s_tokenAdminRegistry.acceptAdminRole(address(s_USDCToken));
    s_tokenAdminRegistry.setPool(address(s_USDCToken), address(s_usdcTokenPool));

    ERC20LockBox.AllowedCallerConfigArgs[] memory allowedCallers = new ERC20LockBox.AllowedCallerConfigArgs[](2);

    allowedCallers[0] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_USDCToken),
      caller: address(s_usdcTokenPool),
      allowed: true
    });
    allowedCallers[1] =
      ERC20LockBox.AllowedCallerConfigArgs({token: address(s_USDCToken), caller: address(this), allowed: true});

    ERC20LockBox(s_lockBox).configureAllowedCallers(allowedCallers);

    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_usdcTokenPool), allowed: true});
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);
    s_cctpMessageTransmitterProxyForTransferLiquidity = new CCTPMessageTransmitterProxy(s_mockUSDC);
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
    allowedCallerParams[0].caller = address(s_usdcTokenPoolTransferLiquidity);
    s_cctpMessageTransmitterProxyForTransferLiquidity.configureAllowedCallers(allowedCallerParams);

    BurnMintERC677(address(s_USDCToken)).grantMintAndBurnRoles(address(s_usdcTokenPool));

    _poolApplyChainUpdates(address(s_usdcTokenPool));

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_SELECTOR,
      mintRecipient: bytes32(0),
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCaller"),
      enabled: true,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2
    });

    s_usdcTokenPool.setDomains(domains);

    s_usdcTokenPool.setLiquidityProvider(DEST_CHAIN_SELECTOR, OWNER);
    s_usdcTokenPool.setLiquidityProvider(SOURCE_CHAIN_SELECTOR, OWNER);
  }
}
