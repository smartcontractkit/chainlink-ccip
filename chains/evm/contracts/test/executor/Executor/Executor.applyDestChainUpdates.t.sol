// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Executor_applyDestChainUpdates is ExecutorSetup {
  function _generateRemoteChainConfig(
    uint64 destChainSelector
  ) internal pure returns (Executor.RemoteChainConfigArgs memory) {
    return Executor.RemoteChainConfigArgs({
      destChainSelector: destChainSelector,
      config: Executor.RemoteChainConfig({
        usdCentsFee: DEFAULT_EXEC_FEE_USD_CENTS,
        baseExecGas: DEFAULT_EXEC_GAS,
        destAddressLengthBytes: EVM_ADDRESS_LENGTH,
        enabled: true
      })
    });
  }

  function test_applyDestChainUpdates_AddNewChain() public {
    uint64 newChainSelector = DEST_CHAIN_SELECTOR + 1;
    Executor.RemoteChainConfigArgs[] memory newRemote = new Executor.RemoteChainConfigArgs[](1);
    newRemote[0] = _generateRemoteChainConfig(newChainSelector);

    vm.expectEmit();
    emit Executor.DestChainAdded(newChainSelector, newRemote[0].config);
    s_executor.applyDestChainUpdates(new uint64[](0), newRemote);

    Executor.RemoteChainConfigArgs[] memory currentDestChains = s_executor.getDestChains();
    assertEq(2, currentDestChains.length);

    bool found = false;
    for (uint256 i = 0; i < currentDestChains.length; ++i) {
      if (currentDestChains[i].destChainSelector == newChainSelector) {
        assertEq(newRemote[0].destChainSelector, currentDestChains[i].destChainSelector);
        assertEq(newRemote[0].config.usdCentsFee, currentDestChains[i].config.usdCentsFee);
        assertEq(newRemote[0].config.baseExecGas, currentDestChains[i].config.baseExecGas);
        assertEq(newRemote[0].config.destAddressLengthBytes, currentDestChains[i].config.destAddressLengthBytes);
        assertEq(newRemote[0].config.enabled, currentDestChains[i].config.enabled);
        found = true;
        break;
      }
    }
    assertTrue(found);
  }

  function test_applyDestChainUpdates_AddExistingChain() public {
    Executor.RemoteChainConfigArgs[] memory newDests = new Executor.RemoteChainConfigArgs[](1);
    newDests[0] = _generateRemoteChainConfig(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit Executor.DestChainAdded(DEST_CHAIN_SELECTOR, newDests[0].config);

    s_executor.applyDestChainUpdates(new uint64[](0), newDests);

    Executor.RemoteChainConfigArgs[] memory currentDestChains = s_executor.getDestChains();
    assertEq(1, currentDestChains.length);
  }

  function test_applyDestChainUpdates_RemoveExistingChain() public {
    uint64[] memory destsToRemove = new uint64[](1);
    destsToRemove[0] = DEST_CHAIN_SELECTOR;

    vm.expectEmit();
    emit Executor.DestChainRemoved(DEST_CHAIN_SELECTOR);
    s_executor.applyDestChainUpdates(destsToRemove, new Executor.RemoteChainConfigArgs[](0));

    Executor.RemoteChainConfigArgs[] memory currentDestChains = s_executor.getDestChains();
    assertEq(0, currentDestChains.length);
  }

  function test_applyDestChainUpdates_RemoveNonexistentChain() public {
    uint64[] memory destsToRemove = new uint64[](1);
    destsToRemove[0] = 999;

    vm.recordLogs();
    s_executor.applyDestChainUpdates(destsToRemove, new Executor.RemoteChainConfigArgs[](0));
    assertEq(0, vm.getRecordedLogs().length);

    Executor.RemoteChainConfigArgs[] memory currentDestChains = s_executor.getDestChains();
    assertEq(1, currentDestChains.length);
  }

  function test_applyDestChainUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_executor.applyDestChainUpdates(new uint64[](0), new Executor.RemoteChainConfigArgs[](0));
  }

  function test_applyDestChainUpdates_RevertWhen_InvalidDestChain() public {
    uint64 invalidDest = 0;
    Executor.RemoteChainConfigArgs[] memory newDests = new Executor.RemoteChainConfigArgs[](1);
    newDests[0] = _generateRemoteChainConfig(invalidDest);

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, invalidDest));
    s_executor.applyDestChainUpdates(new uint64[](0), newDests);
  }
}
