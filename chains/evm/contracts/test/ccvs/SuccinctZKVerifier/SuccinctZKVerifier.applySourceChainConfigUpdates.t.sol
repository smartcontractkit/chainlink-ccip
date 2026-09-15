// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ISP1Helios} from "../../../interfaces/succinct/ISP1Helios.sol";

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_applySourceChainConfigUpdates is SuccinctZKVerifierSetup {
  function test_applySourceChainConfigUpdates() public {
    MockSP1Helios otherHelios = new MockSP1Helios();
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](2);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: otherHelios,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      lightClientVkey: keccak256("otherLightClientVkey"),
      executionHeaderVkey: keccak256("otherExecutionHeaderVkey")
    });
    // A zero light client pauses the chain, so no vkeys are needed.
    sourceChainConfigs[1] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: ISP1Helios(address(0)),
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      lightClientVkey: bytes32(0),
      executionHeaderVkey: bytes32(0)
    });

    vm.expectEmit();
    emit SuccinctZKVerifier.SourceChainConfigSet(SOURCE_CHAIN_SELECTOR, sourceChainConfigs[0]);
    vm.expectEmit();
    emit SuccinctZKVerifier.SourceChainConfigSet(DEST_CHAIN_SELECTOR, sourceChainConfigs[1]);

    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);

    SuccinctZKVerifier.SourceChainConfig memory config = s_zkVerifier.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(address(otherHelios), address(config.helios));
    assertEq(sourceChainConfigs[0].lightClientVkey, config.lightClientVkey);
    assertEq(sourceChainConfigs[0].executionHeaderVkey, config.executionHeaderVkey);

    config = s_zkVerifier.getSourceChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(address(0), address(config.helios));
  }

  // Reverts

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroSelector() public {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs = _getSourceChainConfigArgs();
    sourceChainConfigs[0].sourceChainSelector = 0;

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, 0));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroVkey() public {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs = _getSourceChainConfigArgs();
    sourceChainConfigs[0].lightClientVkey = bytes32(0);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);

    sourceChainConfigs[0].lightClientVkey = LIGHT_CLIENT_VKEY;
    sourceChainConfigs[0].executionHeaderVkey = bytes32(0);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.applySourceChainConfigUpdates(new SuccinctZKVerifier.SourceChainConfigArgs[](0));
  }

  function _getSourceChainConfigArgs()
    internal
    view
    returns (SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs)
  {
    sourceChainConfigs = new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: s_mockHelios,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      lightClientVkey: LIGHT_CLIENT_VKEY,
      executionHeaderVkey: EXECUTION_HEADER_VKEY
    });
    return sourceChainConfigs;
  }
}
