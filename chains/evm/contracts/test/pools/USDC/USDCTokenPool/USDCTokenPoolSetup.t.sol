// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolHelper} from "../../../helpers/USDCTokenPoolHelper.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";
import {MockUSDCTokenMessenger} from "../../../mocks/MockUSDCTokenMessenger.sol";

import {BurnMintERC677} from "@chainlink/contracts/src/v0.8/shared/token/ERC677/BurnMintERC677.sol";

contract USDCTokenPoolSetup is USDCSetup {
  USDCTokenPoolHelper internal s_usdcTokenPool;
  USDCTokenPoolHelper internal s_usdcTokenPoolWithAllowList;
  address[] internal s_allowedList;

  function setUp() public virtual override {
    super.setUp();

    s_mockUSDCTransmitter = new MockE2EUSDCTransmitter(0, DEST_DOMAIN_IDENTIFIER, address(s_USDCToken));
    s_mockUSDC = new MockUSDCTokenMessenger(0, address(s_mockUSDCTransmitter));
    s_mockLegacyUSDC = new MockUSDCTokenMessenger(0, address(s_mockUSDCTransmitter));
    s_cctpMessageTransmitterProxy = new CCTPMessageTransmitterProxy(s_mockUSDC);

    BurnMintERC677(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockUSDCTransmitter));
    BurnMintERC677(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockUSDC));

    s_usdcTokenPool = new USDCTokenPoolHelper(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );

    // Allow the usdcTokenPool to be used as a token pool proxy
    address[] memory allowedTokenPoolProxies = new address[](3);
    allowedTokenPoolProxies[0] = address(OWNER);
    allowedTokenPoolProxies[1] = address(s_routerAllowedOnRamp);
    allowedTokenPoolProxies[2] = address(s_routerAllowedOffRamp);

    bool[] memory allowed = new bool[](5);
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
    s_usdcTokenPoolWithAllowList = new USDCTokenPoolHelper(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      s_allowedList,
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
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
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1
    });

    s_usdcTokenPool.setDomains(domains);
    s_usdcTokenPoolWithAllowList.setDomains(domains);
  }
}
