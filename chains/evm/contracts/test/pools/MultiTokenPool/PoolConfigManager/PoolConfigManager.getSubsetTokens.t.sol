// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getSubsetTokens is PoolConfigManagerSetup {
  address internal s_tokenC;

  /// @dev The DEFAULT connection policy restricts the other chain to subset tokens. Tokens B and C use the DEFAULT
  /// policies.
  function setUp() public virtual override {
    super.setUp();
    _setConnectionPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    // Remote pool policies have no default records, so the other chain needs an explicit record.
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    s_tokenC = _deployToken("C");
    _registerToken(s_tokenC);
  }

  /// @notice Returns the onlySubsetTokens flag and the subset token list. A record with no restriction and no
  /// subset tokens returns false and an empty list.
  function test_getSubsetTokens() public {
    s_poolConfigManager.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0), _getSubsetTokenArgs(_getTokensBAndC())
    );

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, 0, type(uint64).max);
    assertTrue(onlySubsetTokens);
    assertEq(_getTokensBAndC(), subsetTokens);

    (onlySubsetTokens, subsetTokens) =
      s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertFalse(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }

  function test_getSubsetTokens_Pagination() public {
    s_poolConfigManager.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0), _getSubsetTokenArgs(_getTokensBAndC())
    );

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, 1, 1);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_tokenC), subsetTokens);

    (onlySubsetTokens, subsetTokens) = s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, 2, 1);
    assertTrue(onlySubsetTokens);
    assertEq(0, subsetTokens.length);
  }

  /// @dev Subset tokens of the DEFAULT connection policy record for the other chain.
  function _getSubsetTokenArgs(
    address[] memory localTokens
  ) internal pure returns (ConnectionPolicyStore.SubsetTokenArgs[] memory args) {
    args = new ConnectionPolicyStore.SubsetTokenArgs[](localTokens.length);
    for (uint256 i = 0; i < localTokens.length; ++i) {
      args[i] = ConnectionPolicyStore.SubsetTokenArgs({
        policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, localToken: localTokens[i]
      });
    }
    return args;
  }

  function _getTokensBAndC() internal view returns (address[] memory localTokens) {
    localTokens = new address[](2);
    localTokens[0] = s_tokenB;
    localTokens[1] = s_tokenC;
    return localTokens;
  }
}
