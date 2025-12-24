// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPMessageTransmitterProxy} from "../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTPVerifier_constructor is CCTPVerifierSetup {
  function test_constructor() public {
    vm.expectEmit();
    emit CCTPVerifier.StaticConfigSet(
      address(s_mockTokenMessenger), address(s_messageTransmitterProxy), address(s_USDCToken), LOCAL_DOMAIN_IDENTIFIER
    );

    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );

    // Check allowance of the token messenger on the USDC token.
    assertEq(s_USDCToken.allowance(address(s_cctpVerifier), address(s_mockTokenMessenger)), type(uint256).max);

    // Check the static configuration.
    (address gotTokenMessenger, address gotMessageTransmitter, address gotUSDCToken, uint64 gotDomainId) =
      s_cctpVerifier.getStaticConfig();
    assertEq(gotTokenMessenger, address(s_mockTokenMessenger));
    assertEq(gotMessageTransmitter, address(s_messageTransmitterProxy));
    assertEq(gotUSDCToken, address(s_USDCToken));
    assertEq(gotDomainId, LOCAL_DOMAIN_IDENTIFIER);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_TokenMessengerIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      ITokenMessenger(address(0)),
      s_messageTransmitterProxy,
      s_USDCToken,
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_MessageTransmitterProxyIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      s_mockTokenMessenger,
      CCTPMessageTransmitterProxy(address(0)),
      s_USDCToken,
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_USDCTokenIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      IERC20(address(0)),
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
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
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );
  }

  function test_constructor_RevertWhen_InvalidMessageTransmitterVersion() public {
    vm.mockCall(address(s_mockMessageTransmitter), abi.encodeCall(s_mockMessageTransmitter.version, ()), abi.encode(0));
    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidMessageTransmitterVersion.selector, 1, 0));
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
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
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: FEE_AGGREGATOR, allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_FeeAggregatorIsZero() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new CCTPVerifier(
      s_mockTokenMessenger,
      s_messageTransmitterProxy,
      s_USDCToken,
      s_storageLocations,
      CCTPVerifier.DynamicConfig({
        feeAggregator: address(0), allowlistAdmin: ALLOWLIST_ADMIN, fastFinalityBps: CCTP_FAST_FINALITY_BPS
      }),
      address(s_mockRMNRemote)
    );
  }
}
