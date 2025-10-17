// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2Helper} from "../../../helpers/USDCTokenPoolCCTPV2Helper.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPoolCCTPV2Setup is USDCSetup {
  USDCTokenPoolCCTPV2Helper internal s_usdcTokenPool;

  function setUp() public virtual override {
    super.setUp();

    s_usdcTokenPool = new USDCTokenPoolCCTPV2Helper(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );

    // Set the on and offramps as authorized callers for the pool.
    s_usdcTokenPool.grantRole(s_usdcTokenPool.AUTHORIZED_CALLER_ROLE(), address(s_routerAllowedOnRamp));
    s_usdcTokenPool.grantRole(s_usdcTokenPool.AUTHORIZED_CALLER_ROLE(), address(s_routerAllowedOffRamp));

    // Set the pool as an authorized caller for the message transmitter proxy.
    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_usdcTokenPool), allowed: true});
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);

    // Apply the chain updates to the pool.
    _poolApplyChainUpdates(address(s_usdcTokenPool));

    // Set the domain for the pool.
    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_SELECTOR,
      mintRecipient: bytes32(0),
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCallerDestChain"),
      enabled: true,
      useLegacySourcePoolDataFormat: false
    });
    s_usdcTokenPool.setDomains(domains);
  }
}
