// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2Helper} from "../../../helpers/USDCTokenPoolCCTPV2Helper.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPoolCCTPV2Setup is USDCSetup {
  USDCTokenPoolCCTPV2Helper internal s_usdcTokenPool;
  USDCTokenPoolCCTPV2Helper internal s_usdcTokenPoolWithAllowList;
  address[] internal s_allowedList;

  function setUp() public virtual override {
    super.setUp();

    s_usdcTokenPool = new USDCTokenPoolCCTPV2Helper(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );

    address[] memory allowedTokenPoolProxies = new address[](3);
    allowedTokenPoolProxies[0] = address(OWNER);
    allowedTokenPoolProxies[1] = address(s_routerAllowedOnRamp);
    allowedTokenPoolProxies[2] = address(s_routerAllowedOffRamp);

    bool[] memory allowed = new bool[](3);
    for(uint256 i = 0; i < allowedTokenPoolProxies.length; i++) {
      allowed[i] = true;
    }

    s_usdcTokenPool.setAllowedTokenPoolProxies(allowedTokenPoolProxies, allowed);

    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_usdcTokenPool), allowed: true});
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);

    s_allowedList.push(vm.randomAddress());
    s_usdcTokenPoolWithAllowList = new USDCTokenPoolCCTPV2Helper(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      s_allowedList,
      address(s_mockRMNRemote),
      address(s_router)
    );

    s_usdcTokenPoolWithAllowList.setAllowedTokenPoolProxies(allowedTokenPoolProxies, allowed);

    _poolApplyChainUpdates(address(s_usdcTokenPool));
    _poolApplyChainUpdates(address(s_usdcTokenPoolWithAllowList));

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_SELECTOR,
      mintRecipient: bytes32(0),
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCallerDestChain"),
      enabled: true,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2
    });

    s_usdcTokenPool.setDomains(domains);
    s_usdcTokenPoolWithAllowList.setDomains(domains);
  }
}
