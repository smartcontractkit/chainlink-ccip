// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MockPolicyEngineRevertingDetach {
  function attach() external {}

  function detach() external pure {
    revert("detach not supported");
  }

  function run(
    IPolicyEngine.Payload calldata
  ) external {}

  function typeAndVersion() external pure returns (string memory) {
    return "MockPolicyEngineRevertingDetach 1.0.0";
  }
}

contract MockPolicyEngineNoDetach {
  function attach() external {}

  function run(
    IPolicyEngine.Payload calldata
  ) external {}

  function typeAndVersion() external pure returns (string memory) {
    return "MockPolicyEngineNoDetach 1.0.0";
  }
}

contract AdvancedPoolHooks_setPolicyEngine is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;
  MockPolicyEngine internal s_mockPolicyEngine2;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
    s_mockPolicyEngine2 = new MockPolicyEngine();
  }

  function test_setPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(0), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_ToZeroAddress() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(0));

    s_advancedPoolHooks.setPolicyEngine(address(0));

    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_Change() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertFalse(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(s_mockPolicyEngine2));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine2));

    assertEq(address(s_mockPolicyEngine2), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertTrue(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_SameValue() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.recordLogs();
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_getPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
  }

  function test_setPolicyEngine_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
  }

  function test_setPolicyEngine_RevertWhen_OldEngineDetachReverts() public {
    MockPolicyEngineRevertingDetach revertingEngine = new MockPolicyEngineRevertingDetach();
    s_advancedPoolHooks.setPolicyEngine(address(revertingEngine));
    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectRevert(
      abi.encodeWithSelector(
        AdvancedPoolHooks.PolicyEngineDetachFailed.selector,
        address(revertingEngine),
        abi.encodeWithSignature("Error(string)", "detach not supported")
      )
    );
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());
  }

  function test_setPolicyEngine_RevertWhen_OldEngineDoesNotImplementDetach() public {
    MockPolicyEngineNoDetach noDetachEngine = new MockPolicyEngineNoDetach();
    s_advancedPoolHooks.setPolicyEngine(address(noDetachEngine));
    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectRevert(
      abi.encodeWithSelector(AdvancedPoolHooks.PolicyEngineDetachFailed.selector, address(noDetachEngine), "")
    );
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());
  }
}

contract AdvancedPoolHooks_setPolicyEngineAllowFailedDetach is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;
  MockPolicyEngine internal s_mockPolicyEngine2;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
    s_mockPolicyEngine2 = new MockPolicyEngine();
  }

  function test_setPolicyEngineAllowFailedDetach() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(0), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_ToZeroAddress() public {
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(0));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(0));

    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_Change() public {
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertFalse(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(s_mockPolicyEngine2));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine2));

    assertEq(address(s_mockPolicyEngine2), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertTrue(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_SameValue() public {
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.recordLogs();
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_WhenOldEngineDetachReverts() public {
    MockPolicyEngineRevertingDetach revertingEngine = new MockPolicyEngineRevertingDetach();
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(revertingEngine));
    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(revertingEngine), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_WhenOldEngineDoesNotImplementDetach() public {
    MockPolicyEngineNoDetach noDetachEngine = new MockPolicyEngineNoDetach();
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(noDetachEngine));
    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(noDetachEngine), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
  }
}
