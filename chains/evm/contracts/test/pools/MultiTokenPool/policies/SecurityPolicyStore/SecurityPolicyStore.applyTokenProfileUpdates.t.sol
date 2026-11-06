// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyTokenProfileUpdates is SecurityPolicyStoreSetup {
  /// @notice Registering a token does not assign a policy. Deregistering it preserves its decimals.
  function test_applyTokenProfileUpdates_RegistersAndDeregistersToken() public {
    address newToken = makeAddr("newToken");
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = _getTokenProfile(newToken);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.TokenProfileSet(newToken, DEFAULT_TOKEN_DECIMALS);
    _applyTokenProfileUpdates(new address[](0), profiles);

    (bool isRegistered, uint8 localDecimals,,) = s_store.getTokenPolicies(newToken);
    assertTrue(isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.TokenRemoved(newToken);
    _applyTokenProfileUpdates(_arrayOf(newToken), new PolicyConfig.TokenProfileArgs[](0));

    (isRegistered, localDecimals,,) = s_store.getTokenPolicies(newToken);
    assertFalse(isRegistered);
    // The decimals stay after a deregistration.
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
  }

  /// @notice Deregistering a token preserves its rate limit buckets.
  /// Registering it again with the same configuration reuses those buckets.
  function test_applyTokenProfileUpdates_DeregistrationKeepsBuckets() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, DEFAULT_POLICY);

    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 400e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    _setTokenRateLimitPolicy(s_tokenA, bytes32(0));
    _setTokenSecurityPolicy(s_tokenA, bytes32(0));
    _applyTokenProfileUpdates(_arrayOf(s_tokenA), new PolicyConfig.TokenProfileArgs[](0));

    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = _getTokenProfile(s_tokenA);
    _applyTokenProfileUpdates(new address[](0), profiles);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);

    assertEq(CAPACITY - 400e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  // Reverts

  /// @notice The decimals of a registered token are fixed while it is registered.
  function test_applyTokenProfileUpdates_RevertWhen_InvalidDecimalArgs() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = _getTokenProfile(s_tokenA);
    profiles[0].localDecimals = DEFAULT_TOKEN_DECIMALS + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidDecimalArgs.selector, s_tokenA, DEFAULT_TOKEN_DECIMALS, DEFAULT_TOKEN_DECIMALS + 1
      )
    );
    _applyTokenProfileUpdates(new address[](0), profiles);
  }

  /// @notice A deregistration needs a token without a policy and without an inline record.
  function test_applyTokenProfileUpdates_RevertWhen_TokenHasConfig() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectRevert(abi.encodeWithSelector(SecurityPolicyStore.TokenHasConfig.selector, s_tokenA));
    _applyTokenProfileUpdates(_arrayOf(s_tokenA), new PolicyConfig.TokenProfileArgs[](0));
  }

  function test_applyTokenProfileUpdates_RevertWhen_ZeroAddressInvalid() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = _getTokenProfile(address(0));

    vm.expectRevert(SecurityPolicyStore.ZeroAddressInvalid.selector);
    _applyTokenProfileUpdates(new address[](0), profiles);
  }

  function test_applyTokenProfileUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    _applyTokenProfileUpdates(_arrayOf(unknownToken), new PolicyConfig.TokenProfileArgs[](0));
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyTokenProfileUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyTokenProfileUpdates(new address[](0), new PolicyConfig.TokenProfileArgs[](0));
  }
}
