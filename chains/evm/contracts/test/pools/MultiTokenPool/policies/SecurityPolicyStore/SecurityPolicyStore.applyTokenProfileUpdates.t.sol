// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_applyTokenProfileUpdates is SecurityPolicyStoreSetup {
  /// @notice A set of a new token registers it. The token has no policy, and a deregistration clears its decimals.
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

  // Reverts

  /// @notice The decimals of a registered token are fixed while it is registered.
  function test_applyTokenProfileUpdates_RevertWhen_InvalidDecimalArgs() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = _getTokenProfile(s_tokenA);
    profiles[0].localDecimals = DEFAULT_TOKEN_DECIMALS + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidDecimalArgs.selector, DEFAULT_TOKEN_DECIMALS, DEFAULT_TOKEN_DECIMALS + 1
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
}
