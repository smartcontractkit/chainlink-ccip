// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LombardVerifier_setPath is LombardVerifierSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 999;
  bytes32 internal constant NEW_LOMBARD_CHAIN_ID = bytes32(uint256(42));
  bytes32 internal constant NEW_ALLOWED_CALLER = bytes32(uint256(0xabcdef));

  function test_setPath() public {
    vm.expectEmit();
    emit LombardVerifier.PathSet(NEW_CHAIN_SELECTOR, NEW_LOMBARD_CHAIN_ID, NEW_ALLOWED_CALLER);

    s_lombardVerifier.setPath(NEW_CHAIN_SELECTOR, NEW_LOMBARD_CHAIN_ID, NEW_ALLOWED_CALLER);

    LombardVerifier.Path memory path = s_lombardVerifier.getPath(NEW_CHAIN_SELECTOR);
    assertEq(path.lChainId, NEW_LOMBARD_CHAIN_ID);
    assertEq(path.allowedCaller, NEW_ALLOWED_CALLER);

    // Verify chain is added to supported chains.
    uint64[] memory supportedChains = s_lombardVerifier.getSupportedChains();
    bool found = false;
    for (uint256 i = 0; i < supportedChains.length; i++) {
      if (supportedChains[i] == NEW_CHAIN_SELECTOR) {
        found = true;
        break;
      }
    }
    assertTrue(found, "Chain selector should be in supported chains");
  }

  function test_setPath_RevertWhen_ZeroLombardChainId() public {
    vm.expectRevert(LombardVerifier.ZeroLombardChainId.selector);
    s_lombardVerifier.setPath(NEW_CHAIN_SELECTOR, bytes32(0), NEW_ALLOWED_CALLER);
  }

  function test_setPath_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.setPath(NEW_CHAIN_SELECTOR, NEW_LOMBARD_CHAIN_ID, NEW_ALLOWED_CALLER);
  }
}
