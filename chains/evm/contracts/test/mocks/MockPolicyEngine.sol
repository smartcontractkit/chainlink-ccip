// SPDX-License-Identifier: BUSL-1.1
// solhint-disable one-contract-per-file
pragma solidity ^0.8.24;

import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Mock implementation of IPolicyEngine for testing.
contract MockPolicyEngine is IPolicyEngine {
  using EnumerableSet for EnumerableSet.AddressSet;

  error MockPolicyEngineRejection(string reason);
  error NoPayloadRecorded();

  EnumerableSet.AddressSet internal s_attachedTargets;
  bool internal s_shouldRevert;
  string internal s_revertReason;
  IPolicyEngine.Payload internal s_lastPayload;
  bool internal s_payloadRecorded;

  function attach() external override {
    if (s_attachedTargets.contains(msg.sender)) {
      revert TargetAlreadyAttached(msg.sender);
    }
    s_attachedTargets.add(msg.sender);
    emit TargetAttached(msg.sender);
  }

  function detach() external override {
    if (!s_attachedTargets.contains(msg.sender)) {
      revert TargetNotAttached(msg.sender);
    }
    s_attachedTargets.remove(msg.sender);
    emit TargetDetached(msg.sender);
  }

  function run(
    Payload calldata payload
  ) external override {
    s_lastPayload = payload;
    s_payloadRecorded = true;

    if (s_shouldRevert) {
      revert MockPolicyEngineRejection(s_revertReason);
    }
  }

  function check(
    Payload calldata
  ) external view override {
    if (s_shouldRevert) {
      revert MockPolicyEngineRejection(s_revertReason);
    }
  }

  /// @notice Sets whether run() should revert.
  function setShouldRevert(
    bool shouldRevert,
    string memory revertReason
  ) external {
    s_shouldRevert = shouldRevert;
    s_revertReason = revertReason;
  }

  /// @notice Returns whether a target is attached.
  function isAttached(
    address target
  ) external view returns (bool) {
    return s_attachedTargets.contains(target);
  }

  /// @notice Returns the last payload passed to run().
  function getLastPayload() external view returns (Payload memory) {
    if (!s_payloadRecorded) revert NoPayloadRecorded();
    return s_lastPayload;
  }

  /// @notice Clears the recorded payload.
  function clearLastPayload() external {
    s_payloadRecorded = false;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "MockPolicyEngine 1.0.0";
  }

  function setExtractor(
    bytes4,
    address
  ) external override {}

  function setExtractors(
    bytes4[] calldata,
    address
  ) external override {}

  function getExtractor(
    bytes4
  ) external pure override returns (address) {
    return address(0);
  }

  function setPolicyMapper(
    address,
    address
  ) external override {}

  function getPolicyMapper(
    address
  ) external pure override returns (address) {
    return address(0);
  }

  function addPolicy(
    address,
    bytes4,
    address,
    bytes32[] calldata
  ) external override {}

  function addPolicyAt(
    address,
    bytes4,
    address,
    bytes32[] calldata,
    uint256
  ) external override {}

  function removePolicy(
    address,
    bytes4,
    address
  ) external override {}

  function getPolicies(
    address,
    bytes4
  ) external pure override returns (address[] memory) {
    return new address[](0);
  }

  function setPolicyConfiguration(
    address,
    uint256,
    bytes4,
    bytes calldata
  ) external override {}

  function getPolicyConfigVersion(
    address
  ) external pure override returns (uint256) {
    return 0;
  }

  function setDefaultPolicyAllow(
    bool
  ) external override {}

  function setTargetDefaultPolicyAllow(
    address,
    bool
  ) external override {}
}

contract MockPolicyEngineRevertingDetach {
  error DetachNotSupported();

  function attach() external {}

  function detach() external pure {
    revert DetachNotSupported();
  }

  function run(
    IPolicyEngine.Payload calldata
  ) external {}

  function check(
    IPolicyEngine.Payload calldata
  ) external view {}
}

contract MockPolicyEngineNoDetach {
  function attach() external {}

  function run(
    IPolicyEngine.Payload calldata
  ) external {}

  function check(
    IPolicyEngine.Payload calldata
  ) external view {}
}
