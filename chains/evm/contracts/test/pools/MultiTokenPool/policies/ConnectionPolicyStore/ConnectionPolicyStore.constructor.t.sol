// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_constructor is ConnectionPolicyStoreSetup {
  function test_constructor() public {
    vm.startPrank(OWNER);
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(s_poolConfigManager);
    ConnectionPolicyStore connectionStore = new ConnectionPolicyStore(_arrayOf(s_poolConfigManager));

    assertEq("ConnectionPolicyStore 2.1.0-dev", connectionStore.typeAndVersion());
    assertEq(OWNER, connectionStore.owner());
    assertEq(_arrayOf(s_poolConfigManager), connectionStore.getAllAuthorizedCallers());
    IPoolConfigManager.LaneConnection[] memory connections =
      connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertFalse(connections[0].isOutboundPaused);
    assertFalse(connections[0].isInboundPaused);
  }
}
