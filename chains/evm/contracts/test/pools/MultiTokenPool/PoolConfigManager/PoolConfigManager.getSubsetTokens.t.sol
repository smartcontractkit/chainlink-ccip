// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getSubsetTokens is PoolConfigManagerSetup {
  address internal s_tokenC;

  /// @dev The DEFAULT connection policy opens the other chain only for its subset tokens. Token C uses the DEFAULT
  /// policy of every type, the same as token B.
  function setUp() public virtual override {
    super.setUp();
    _setConnectionPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    // The merged store has no remote pool policy default record, so the other chain gets a chain record.
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    s_tokenC = _deployToken("C");
    _registerToken(s_tokenC);
  }

  /// @notice A connection policy record returns its flag and its subset tokens. A record with the flag off and no
  /// subset tokens returns false and an empty list.
  function test_getSubsetTokens() public {
    s_poolConfigManager.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0), _getSubsetTokenArgs(_getTokensBAndC())
    );

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR);
    assertTrue(onlySubsetTokens);
    assertEq(_getTokensBAndC(), subsetTokens);

    (onlySubsetTokens, subsetTokens) = s_poolConfigManager.getSubsetTokens(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
    assertFalse(onlySubsetTokens);
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
