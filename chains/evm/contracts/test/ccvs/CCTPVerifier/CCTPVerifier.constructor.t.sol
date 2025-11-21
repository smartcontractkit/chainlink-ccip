// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {CCTPMessageTransmitterProxy} from "../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CCTPVerifier_constructor is CCTPVerifierSetup {
  uint16[] internal s_customCCIPFinalities = new uint16[](1);

  function setUp() public override {
    super.setUp();

    s_customCCIPFinalities[0] = CCIP_FAST_FINALITY_THRESHOLD;
  }

  function test_constructor() public {
    vm.expectEmit();
    emit CCTPVerifier.StaticConfigSet(
      address(s_mockTokenMessenger), address(s_messageTransmitterProxy), address(s_USDCToken), LOCAL_DOMAIN_IDENTIFIER
    );

    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );

    // Check allowance of the token messenger on the USDC token.
    assertEq(s_USDCToken.allowance(address(s_cctpVerifier), address(s_mockTokenMessenger)), type(uint256).max);

    // Check the static configuration.
    CCTPVerifier.StaticConfig memory staticConfig = s_cctpVerifier.getStaticConfig();
    assertEq(staticConfig.tokenMessenger, address(s_mockTokenMessenger));
    assertEq(staticConfig.messageTransmitterProxy, address(s_messageTransmitterProxy));
    assertEq(staticConfig.usdcToken, address(s_USDCToken));
    assertEq(staticConfig.localDomainIdentifier, LOCAL_DOMAIN_IDENTIFIER);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_TokenMessengerIsZero() public {
    vm.expectRevert(CCTPVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      ITokenMessenger(address(0)),
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_MessageTransmitterProxyIsZero() public {
    vm.expectRevert(CCTPVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      s_mockTokenMessenger,
      CCTPMessageTransmitterProxy(address(0)),
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_USDCTokenIsZero() public {
    vm.expectRevert(CCTPVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      IERC20(address(0)),
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }

  function test_constructor_RevertWhen_InvalidTokenMessengerVersion() public {
    vm.mockCall(
      address(s_mockTokenMessenger), abi.encodeCall(s_mockTokenMessenger.messageBodyVersion, ()), abi.encode(0)
    );
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidTokenMessengerVersion.selector, 1, 0));
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }

  function test_constructor_RevertWhen_InvalidMessageTransmitterVersion() public {
    vm.mockCall(address(s_mockMessageTransmitter), abi.encodeCall(s_mockMessageTransmitter.version, ()), abi.encode(0));
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidMessageTransmitterVersion.selector, 1, 0));
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }

  function test_constructor_RevertWhen_InvalidMessageTransmitterOnProxy() public {
    vm.mockCall(
      address(s_messageTransmitterProxy),
      abi.encodeCall(s_messageTransmitterProxy.i_cctpTransmitter, ()),
      abi.encode(address(0))
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        CCTPVerifier.InvalidMessageTransmitterOnProxy.selector, address(s_mockMessageTransmitter), address(0)
      )
    );
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      STORAGE_LOCATION,
      CCTPVerifier.FinalityConfig({
        defaultCCTPFinalityThreshold: CCTP_STANDARD_FINALITY_THRESHOLD,
        defaultCCTPFinalityBps: CCTP_STANDARD_FINALITY_BPS,
        customCCIPFinalities: s_customCCIPFinalities,
        customCCTPFinalityThresholds: new uint32[](1),
        customCCTPFinalityBps: new uint16[](1)
      }),
      CCTPVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN})
    );
  }
}
