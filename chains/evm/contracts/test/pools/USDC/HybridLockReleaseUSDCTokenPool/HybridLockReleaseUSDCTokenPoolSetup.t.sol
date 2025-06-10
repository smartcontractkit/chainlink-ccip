// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintERC677} from "@chainlink/contracts/src/v0.8/shared/token/ERC677/BurnMintERC677.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract HybridLockReleaseUSDCTokenPoolSetup is USDCSetup {
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPool;
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;
  CCTPMessageTransmitterProxy internal s_cctpMessageTransmitterProxyForTransferLiquidity;
  address[] internal s_allowedList;

  address internal s_previousPool = makeAddr("previousPool");

  function setUp() public virtual override {
    super.setUp();

    vm.mockCall(
      s_previousPool,
      abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector),
      abi.encode(Pool.ReleaseOrMintOutV1({destinationAmount: 1}))
    );

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      s_mockUSDCV2,
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool // previousPool
    );

    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](1);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_usdcTokenPool), allowed: true});

    // Set the token pool as an allowed caller on BOTH transmitter proxies
    // one for CCTP v1 and one for V2
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);

    s_cctpMessageTransmitterProxyForTransferLiquidity = new CCTPMessageTransmitterProxy(s_mockUSDC, s_mockUSDCV2);

    s_usdcTokenPoolTransferLiquidity = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      s_mockUSDCV2,
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool // previousPool
    );

    allowedCallerParams[0].caller = address(s_usdcTokenPoolTransferLiquidity);
    s_cctpMessageTransmitterProxyForTransferLiquidity.configureAllowedCallers(allowedCallerParams);

    BurnMintERC677(address(s_token)).grantMintAndBurnRoles(address(s_usdcTokenPool));

    _poolApplyChainUpdates(address(s_usdcTokenPool));

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_SELECTOR,
      mintRecipient: bytes32(0),
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCaller"),
      enabled: true
    });

    s_usdcTokenPool.setDomains(domains);

    s_usdcTokenPool.setLiquidityProvider(DEST_CHAIN_SELECTOR, OWNER);
    s_usdcTokenPool.setLiquidityProvider(SOURCE_CHAIN_SELECTOR, OWNER);
  }
}
