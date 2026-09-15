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
    address otherOnRamp = makeAddr("otherOnRamp");
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](2);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: otherHelios, sourceChainSelector: SOURCE_CHAIN_SELECTOR, maxHeaderChainLength: 32, onRamp: otherOnRamp
    });
    // A zero light client pauses the chain.
    sourceChainConfigs[1] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: ISP1Helios(address(0)),
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      maxHeaderChainLength: 1,
      onRamp: s_onRamp
    });

    vm.expectEmit();
    emit SuccinctZKVerifier.SourceChainConfigSet(SOURCE_CHAIN_SELECTOR, address(otherHelios), otherOnRamp, 32);
    vm.expectEmit();
    emit SuccinctZKVerifier.SourceChainConfigSet(DEST_CHAIN_SELECTOR, address(0), s_onRamp, 1);

    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);

    SuccinctZKVerifier.SourceChainConfig memory config = s_zkVerifier.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(address(otherHelios), address(config.helios));
    assertEq(otherOnRamp, config.onRamp);
    assertEq(32, config.maxHeaderChainLength);

    config = s_zkVerifier.getSourceChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(address(0), address(config.helios));
    assertEq(s_onRamp, config.onRamp);
    assertEq(1, config.maxHeaderChainLength);
  }

  // Reverts

  function test_applySourceChainConfigUpdates_RevertWhen_InvalidSourceChainConfig_ZeroSelector() public {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: s_mockHelios, sourceChainSelector: 0, maxHeaderChainLength: 32, onRamp: s_onRamp
    });

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, 0));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_InvalidSourceChainConfig_ZeroMaxHeaderChainLength() public {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: s_mockHelios, sourceChainSelector: SOURCE_CHAIN_SELECTOR, maxHeaderChainLength: 0, onRamp: s_onRamp
    });

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_InvalidSourceChainConfig_ZeroOnRamp() public {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: s_mockHelios, sourceChainSelector: SOURCE_CHAIN_SELECTOR, maxHeaderChainLength: 32, onRamp: address(0)
    });

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidSourceChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.applySourceChainConfigUpdates(new SuccinctZKVerifier.SourceChainConfigArgs[](0));
  }
}
