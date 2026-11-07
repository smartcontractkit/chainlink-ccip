// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenProfileUpdates is PoolConfigManagerRolesSetup {
  address internal s_tokenC;
  address internal s_tokenD;

  function setUp() public virtual override {
    super.setUp();
    s_tokenC = _deployToken("C");
    s_tokenD = _deployToken("D");
  }

  /// @notice Registers multiple tokens in all three stores. The new tokens appear after those registered during
  /// setup.
  function test_applyTokenProfileUpdates() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](2);
    profiles[0] = _getTokenProfileArgs(s_tokenC);
    profiles[1] = _getTokenProfileArgs(s_tokenD);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenProfileSet(
      s_tokenC,
      DEFAULT_TOKEN_DECIMALS,
      IPoolConfigManager.LockOrBurnFlavor.Burn,
      address(0),
      FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      address(0)
    );
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenProfileSet(
      s_tokenD,
      DEFAULT_TOKEN_DECIMALS,
      IPoolConfigManager.LockOrBurnFlavor.Burn,
      address(0),
      FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      address(0)
    );
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.TokenProfileSet(s_tokenC);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.TokenProfileSet(s_tokenD);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenProfileSet(s_tokenC, DEFAULT_TOKEN_DECIMALS);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenProfileSet(s_tokenD, DEFAULT_TOKEN_DECIMALS);
    s_poolConfigManager.applyTokenProfileUpdates(new address[](0), profiles);

    address[] memory expectedTokens = new address[](4);
    expectedTokens[0] = s_tokenA;
    expectedTokens[1] = s_tokenB;
    expectedTokens[2] = s_tokenC;
    expectedTokens[3] = s_tokenD;
    assertEq(expectedTokens, s_poolConfigManager.getAllConfiguredTokens(0, 10));
    (bool isRegistered,,,) = s_securityStore.getTokenPolicies(s_tokenC);
    assertTrue(isRegistered);
    (isRegistered,,,) = s_securityStore.getTokenPolicies(s_tokenD);
    assertTrue(isRegistered);
    (isRegistered,,) = s_feeStore.getTokenPolicies(s_tokenC);
    assertTrue(isRegistered);
    (isRegistered,,) = s_feeStore.getTokenPolicies(s_tokenD);
    assertTrue(isRegistered);
  }

  /// @notice Updating a registered token replaces its profile in all three stores and preserves its policy
  /// assignments.
  function test_applyTokenProfileUpdates_ProfileSetKeepsPolicies() public {
    address additionalPoolDataProvider = makeAddr("additionalPoolDataProvider");
    // The connection store accepts a provider only if it has code. No transfer runs here, so nothing calls it.
    vm.etch(additionalPoolDataProvider, hex"00");
    PolicyConfig.TokenProfileArgs memory tokenProfile = _getTokenProfileArgs(s_tokenB);
    tokenProfile.allowedFinalityConfig = _getFastFinalityConfig();
    tokenProfile.additionalPoolDataProvider = additionalPoolDataProvider;

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenProfileSet(
      s_tokenB,
      DEFAULT_TOKEN_DECIMALS,
      IPoolConfigManager.LockOrBurnFlavor.Burn,
      address(0),
      _getFastFinalityConfig(),
      additionalPoolDataProvider
    );
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenProfileSet(s_tokenB, DEFAULT_TOKEN_DECIMALS);
    _setTokenProfile(tokenProfile);

    assertEq(abi.encode(tokenProfile), abi.encode(_getStoredTokenProfile(s_tokenB)));
    _assertPolicyNames(_getTokenConfig(s_tokenB), DEFAULT_POLICY);
  }

  /// @notice A token can be deregistered from all three stores after its policy assignments are removed.
  function test_applyTokenProfileUpdates_RemoveToken() public {
    _clearTokenPolicies(s_tokenB);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenRemoved(s_tokenB);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.TokenRemoved(s_tokenB);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenRemoved(s_tokenB);
    s_poolConfigManager.applyTokenProfileUpdates(_arrayOf(s_tokenB), new PolicyConfig.TokenProfileArgs[](0));

    assertEq(_arrayOf(s_tokenA), s_poolConfigManager.getAllConfiguredTokens(0, 10));
    assertFalse(s_connectionStore.getTokenProfiles(_arrayOf(s_tokenB))[0].isRegistered);
    (bool isRegistered,,,) = s_securityStore.getTokenPolicies(s_tokenB);
    assertFalse(isRegistered);
    (isRegistered,,) = s_feeStore.getTokenPolicies(s_tokenB);
    assertFalse(isRegistered);
  }

  /// @notice Either CONFIG_ROLE or TOKEN_MANAGER_ROLE authorizes the call on its own.
  function test_applyTokenProfileUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyTokenProfileUpdates, (new address[](0), new PolicyConfig.TokenProfileArgs[](0))
      ),
      _getProfileWriteRoles()
    );
  }

  // Reverts

  /// @notice The connection store runs first and rejects token C's supplied decimals. The security store would
  /// accept token C but reject token B's decimal change. The caller receives the connection store error.
  function test_applyTokenProfileUpdates_RevertWhen_InvalidDecimalArgs() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](2);
    profiles[0] = _getTokenProfileArgs(s_tokenC);
    profiles[0].localDecimals = 6;
    profiles[1] = _getTokenProfileArgs(s_tokenB);
    profiles[1].localDecimals = 8;

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidDecimalArgs.selector, s_tokenC, DEFAULT_TOKEN_DECIMALS, 6)
    );
    s_poolConfigManager.applyTokenProfileUpdates(new address[](0), profiles);
  }

  /// @notice Deregistration is atomic across all three stores. The fee store first rejects token B's remaining fee
  /// configuration. After those assignments are removed, the security store rejects its remaining policy assignments.
  /// Each error reaches the caller unchanged, and earlier store changes are reverted.
  function test_applyTokenProfileUpdates_RevertWhen_TokenHasConfig() public {
    _setTokenConnectionPolicy(s_tokenB, bytes32(0));
    _setTokenRemotePoolPolicy(s_tokenB, bytes32(0));

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.TokenHasConfig.selector, s_tokenB), address(s_feeStore));
    s_poolConfigManager.applyTokenProfileUpdates(_arrayOf(s_tokenB), new PolicyConfig.TokenProfileArgs[](0));

    _setTokenFeePolicy(s_tokenB, bytes32(0));
    _setTokenDestOverheadPolicy(s_tokenB, bytes32(0));
    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.TokenHasConfig.selector, s_tokenB), address(s_securityStore)
    );
    s_poolConfigManager.applyTokenProfileUpdates(_arrayOf(s_tokenB), new PolicyConfig.TokenProfileArgs[](0));

    assertTrue(s_connectionStore.getTokenProfiles(_arrayOf(s_tokenB))[0].isRegistered);
    (bool isRegistered,,) = s_feeStore.getTokenPolicies(s_tokenB);
    assertTrue(isRegistered);
  }

  /// @notice Calls without an allowed role revert.
  function test_applyTokenProfileUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyTokenProfileUpdates, (new address[](0), new PolicyConfig.TokenProfileArgs[](0))
      ),
      _getProfileWriteRoles()
    );
  }

  /// @dev Checks that the token has the same policy name for every type.
  function _assertPolicyNames(
    IPoolConfigManager.TokenConfig memory tokenConfig,
    bytes32 policyName
  ) internal pure {
    assertEq(policyName, tokenConfig.connectionPolicyName);
    assertEq(policyName, tokenConfig.remotePoolPolicyName);
    assertEq(policyName, tokenConfig.rateLimitPolicyName);
    assertEq(policyName, tokenConfig.feePolicyName);
    assertEq(policyName, tokenConfig.destOverheadPolicyName);
    assertEq(policyName, tokenConfig.securityPolicyName);
  }

  /// @dev Removes all policy assignments, starting with the connection policy.
  function _clearTokenPolicies(
    address localToken
  ) internal {
    _setTokenConnectionPolicy(localToken, bytes32(0));
    _setTokenRemotePoolPolicy(localToken, bytes32(0));
    _setTokenRateLimitPolicy(localToken, bytes32(0));
    _setTokenFeePolicy(localToken, bytes32(0));
    _setTokenDestOverheadPolicy(localToken, bytes32(0));
    _setTokenSecurityPolicy(localToken, bytes32(0));
  }
}
