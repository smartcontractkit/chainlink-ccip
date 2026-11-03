// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token A has an inline connection record to the destination chain. Token B uses the DEFAULT connection
/// policy, which serves the destination chain. An inline connection record always gives its token a lane, so no
/// inline connection record exists on a chain without a lane.
contract PoolConfigManager_getInlineConnectionConfigs is PoolConfigManagerSetup {
  /// @dev The merged store has no remote pool policy default record, so the other chain gets a chain record. Every
  /// test that opens an inline connection on the other chain needs it.
  function setUp() public virtual override {
    super.setUp();
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
  }

  /// @notice The view returns each inline connection record of the token in the shape of
  /// applyInlineConnectionConfigUpdates, in the order of the sets. A remote token in another encoding shows unchanged.
  function test_getInlineConnectionConfigs() public {
    bytes memory otherRemoteToken = abi.encodePacked(keccak256("otherRemoteToken"));
    _setInlineConnectionConfig(s_tokenA, OTHER_CHAIN_SELECTOR, otherRemoteToken);

    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory expectedConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](2);
    expectedConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenA)
    });
    expectedConfigs[1] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenA, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: otherRemoteToken
    });
    assertEq(abi.encode(expectedConfigs), abi.encode(s_poolConfigManager.getInlineConnectionConfigs(s_tokenA)));
  }
}
