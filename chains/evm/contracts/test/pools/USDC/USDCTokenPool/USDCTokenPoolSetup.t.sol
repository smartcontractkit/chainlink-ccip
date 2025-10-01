// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolHelper} from "../../../helpers/USDCTokenPoolHelper.sol";

import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";
import {MockUSDCTokenMessenger} from "../../../mocks/MockUSDCTokenMessenger.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract USDCTokenPoolSetup is USDCSetup {
  USDCTokenPoolHelper internal s_usdcTokenPool;
  USDCTokenPoolHelper internal s_usdcTokenPoolWithAllowList;
  address[] internal s_allowedList;

  function setUp() public virtual override {
    super.setUp();

    s_mockUSDCTransmitter = new MockE2EUSDCTransmitter(0, DEST_DOMAIN_IDENTIFIER, address(s_USDCToken));
    s_mockUSDCTokenMessenger = new MockUSDCTokenMessenger(0, address(s_mockUSDCTransmitter));
    s_cctpMessageTransmitterProxy = new CCTPMessageTransmitterProxy(s_mockUSDCTokenMessenger);

    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockUSDCTransmitter));
    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockUSDCTokenMessenger));

    s_usdcTokenPool = new USDCTokenPoolHelper(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );

    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_usdcTokenPool), allowed: true});
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);

    s_allowedList.push(vm.randomAddress());
    s_usdcTokenPoolWithAllowList = new USDCTokenPoolHelper(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      s_allowedList,
      address(s_mockRMNRemote),
      address(s_router)
    );

    // Set the owner as an authorized caller for the pools
    address[] memory authorizedCallers = new address[](3);
    authorizedCallers[0] = OWNER;
    authorizedCallers[1] = address(s_routerAllowedOnRamp);
    authorizedCallers[2] = address(s_routerAllowedOffRamp);
    s_usdcTokenPool.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );
    s_usdcTokenPoolWithAllowList.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    _poolApplyChainUpdates(address(s_usdcTokenPool));
    _poolApplyChainUpdates(address(s_usdcTokenPoolWithAllowList));

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_SELECTOR,
      mintRecipient: bytes32(0),
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCallerDestChain"),
      enabled: true
    });

    s_usdcTokenPool.setDomains(domains);
    s_usdcTokenPoolWithAllowList.setDomains(domains);
  }
}
