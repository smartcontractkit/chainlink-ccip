// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LombardVerifier_removePaths is LombardVerifierSetup {
  function test_removePaths() public {
    // First verify the path exists.
    LombardVerifier.Path memory pathBefore = s_lombardVerifier.getPath(DEST_CHAIN_SELECTOR);
    assertEq(pathBefore.lChainId, LOMBARD_CHAIN_ID);
    assertEq(pathBefore.allowedCaller, ALLOWED_CALLER);

    uint64[] memory chainsToRemove = new uint64[](1);
    chainsToRemove[0] = DEST_CHAIN_SELECTOR;

    vm.expectEmit();
    emit LombardVerifier.PathRemoved(DEST_CHAIN_SELECTOR, LOMBARD_CHAIN_ID, ALLOWED_CALLER);

    s_lombardVerifier.removePaths(chainsToRemove);

    // Verify the path is removed.
    LombardVerifier.Path memory pathAfter = s_lombardVerifier.getPath(DEST_CHAIN_SELECTOR);
    assertEq(pathAfter.lChainId, bytes32(0));
    assertEq(pathAfter.allowedCaller, bytes32(0));

    // Verify chain is removed from supported chains.
    uint64[] memory supportedChains = s_lombardVerifier.getSupportedChains();
    for (uint256 i = 0; i < supportedChains.length; ++i) {
      assertNotEq(supportedChains[i], DEST_CHAIN_SELECTOR, "Chain selector should not be in supported chains");
    }
  }

  function test_removePaths_RevertWhen_PathNotExist() public {
    uint64 nonExistentChainSelector = 999999;
    uint64[] memory chainsToRemove = new uint64[](1);
    chainsToRemove[0] = nonExistentChainSelector;

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.PathNotExist.selector, nonExistentChainSelector));
    s_lombardVerifier.removePaths(chainsToRemove);
  }

  function test_removePaths_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.removePaths(new uint64[](1));
  }
}
