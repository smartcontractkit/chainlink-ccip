// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getTokenConfigBatch is PoolConfigManagerSetup {
  address internal s_tokenC;

  /// @dev Token C uses the policy of each type that is named after its type.
  function setUp() public virtual override {
    super.setUp();
    _addPolicyOfEachType();

    s_tokenC = _deployToken("C");
    _setTokenProfile(_getTokenProfileArgs(s_tokenC));
    _setTokenRemotePoolPolicy(s_tokenC, REMOTE_POOL_POLICY);
    _setTokenRateLimitPolicy(s_tokenC, RATE_LIMIT_POLICY);
    _setTokenFeePolicy(s_tokenC, FEE_POLICY);
    _setTokenDestOverheadPolicy(s_tokenC, DEST_OVERHEAD_POLICY);
    _setTokenSecurityPolicy(s_tokenC, SECURITY_POLICY);
    _setTokenConnectionPolicy(s_tokenC, CONNECTION_POLICY);
  }

  /// @notice Returns the profile and enabled chains for each registered token. Unregistered tokens return zero
  /// values.
  function test_getTokenConfigBatch() public {
    address unknownToken = makeAddr("unknownToken");
    address[] memory localTokens = new address[](2);
    localTokens[0] = s_tokenA;
    localTokens[1] = unknownToken;

    IPoolConfigManager.TokenConfig[] memory configs = s_poolConfigManager.getTokenConfigBatch(localTokens);

    assertTrue(configs[0].isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, configs[0].localDecimals);
    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.Burn), uint8(configs[0].lockOrBurnFlavor));
    assertEq(FinalityCodec.WAIT_FOR_FINALITY_FLAG, configs[0].allowedFinalityConfig);
    assertEq(address(0), configs[0].lockOrBurnAddress);
    assertEq(address(0), configs[0].additionalPoolDataProvider);
    assertEq(1, configs[0].supportedChains.length);
    assertEq(DEST_CHAIN_SELECTOR, configs[0].supportedChains[0]);

    assertFalse(configs[1].isRegistered);
    assertEq(0, configs[1].localDecimals);
    assertEq(0, configs[1].supportedChains.length);
  }

  /// @notice Returns each token's policy names. Missing assignments and unregistered tokens return zero names.
  function test_getTokenConfigBatch_PolicyNames() public {
    address[] memory localTokens = new address[](4);
    localTokens[0] = s_tokenA;
    localTokens[1] = s_tokenB;
    localTokens[2] = s_tokenC;
    localTokens[3] = makeAddr("unknownToken");

    IPoolConfigManager.TokenConfig[] memory configs = s_poolConfigManager.getTokenConfigBatch(localTokens);

    _assertPolicyNames(
      configs[0], bytes32(0), DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY
    );
    _assertPolicyNames(
      configs[1], DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY
    );
    _assertTokenCPolicyNames(configs[2]);
    _assertPolicyNames(configs[3], bytes32(0), bytes32(0), bytes32(0), bytes32(0), bytes32(0), bytes32(0));
  }

  /// @dev Checks the names of the policies that the setup gives token C.
  function _assertTokenCPolicyNames(
    IPoolConfigManager.TokenConfig memory config
  ) internal pure {
    _assertPolicyNames(
      config,
      CONNECTION_POLICY,
      REMOTE_POOL_POLICY,
      RATE_LIMIT_POLICY,
      FEE_POLICY,
      DEST_OVERHEAD_POLICY,
      SECURITY_POLICY
    );
  }

  function _assertPolicyNames(
    IPoolConfigManager.TokenConfig memory config,
    bytes32 connectionPolicyName,
    bytes32 remotePoolPolicyName,
    bytes32 rateLimitPolicyName,
    bytes32 feePolicyName,
    bytes32 destOverheadPolicyName,
    bytes32 securityPolicyName
  ) internal pure {
    assertEq(connectionPolicyName, config.connectionPolicyName);
    assertEq(remotePoolPolicyName, config.remotePoolPolicyName);
    assertEq(rateLimitPolicyName, config.rateLimitPolicyName);
    assertEq(feePolicyName, config.feePolicyName);
    assertEq(destOverheadPolicyName, config.destOverheadPolicyName);
    assertEq(securityPolicyName, config.securityPolicyName);
  }
}
