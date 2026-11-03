// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applySubsetTokenUpdates is PoolConfigManagerRolesSetup {
  address internal s_tokenC;

  /// @dev The DEFAULT connection policy opens the other chain only for its subset tokens. Token C uses the DEFAULT
  /// policy of every type, the same as token B.
  function setUp() public virtual override {
    super.setUp();
    _setConnectionPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    // A remote pool policy has no default record, so the DEFAULT policy needs a record on the other chain for the lane
    // check of the subset tokens that the tests add there.
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    s_tokenC = _deployToken("C");
    _registerToken(s_tokenC);
  }

  /// @notice Each subset token of the call gets a lane on the other chain. Each token has a record of every other type
  /// there, so the lane check passes for each add.
  function test_applySubsetTokenUpdates() public {
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](2);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenB);
    tokensToAdd[1] = _getSubsetTokenArgs(s_tokenC);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.SubsetTokenAdded(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, s_tokenB);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.SubsetTokenAdded(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, s_tokenC);
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);

    assertTrue(_getLaneConnection(s_tokenB, OTHER_CHAIN_SELECTOR).hasConnection);
    assertTrue(_getLaneConnection(s_tokenC, OTHER_CHAIN_SELECTOR).hasConnection);
  }

  /// @notice A removal closes a lane, so the lane check does not run for it. Token B has no fee record when it leaves
  /// the subset.
  function test_applySubsetTokenUpdates_RemovalsAreNotChecked() public {
    _addSubsetToken(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, s_tokenB);
    _setTokenFeePolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToRemove = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToRemove[0] = _getSubsetTokenArgs(s_tokenB);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.SubsetTokenRemoved(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, s_tokenB);
    s_poolConfigManager.applySubsetTokenUpdates(tokensToRemove, new ConnectionPolicyStore.SubsetTokenArgs[](0));

    assertFalse(_getLaneConnection(s_tokenB, OTHER_CHAIN_SELECTOR).hasConnection);
  }

  /// @notice CONFIG_ROLE, POLICY_ROLE and the policy role of the type each pass alone.
  function test_applySubsetTokenUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applySubsetTokenUpdates,
        (new ConnectionPolicyStore.SubsetTokenArgs[](0), new ConnectionPolicyStore.SubsetTokenArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_POLICY_ROLE())
    );
  }

  // Reverts

  /// @notice The lane check runs after the forward. An unknown token has no remote pool record, so a lane check before
  /// the forward would revert with RemotePoolRecordNotFound.
  function test_applySubsetTokenUpdates_RevertWhen_StoreRevertsBeforeLaneCheck() public {
    address unknownToken = makeAddr("unknownToken");
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(unknownToken);

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.TokenNotInPolicy.selector, unknownToken, DEFAULT_POLICY)
    );
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice A subset token add opens a lane, so the lane check reverts for the first type that the token misses.
  /// Token B misses the fee type and the security type.
  function test_applySubsetTokenUpdates_RevertWhen_FeeRecordNotFound() public {
    _setTokenFeePolicy(s_tokenB, bytes32(0));
    _setTokenSecurityPolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenB);

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR));
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice The security store runs the last part of the lane check. Token B has every record but a security record.
  function test_applySubsetTokenUpdates_RevertWhen_SecurityRecordNotFound() public {
    _setTokenSecurityPolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenB);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice The lane check runs for each added token, not only the first. Token C has every record. Token B, added
  /// second, misses a destination overhead record, so the revert names token B.
  function test_applySubsetTokenUpdates_RevertWhen_LaterTokenMissesRecord() public {
    _setTokenDestOverheadPolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](2);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenC);
    tokensToAdd[1] = _getSubsetTokenArgs(s_tokenB);

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applySubsetTokenUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applySubsetTokenUpdates,
        (new ConnectionPolicyStore.SubsetTokenArgs[](0), new ConnectionPolicyStore.SubsetTokenArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_POLICY_ROLE())
    );
  }

  /// @dev A subset token of the DEFAULT connection policy record for the other chain.
  function _getSubsetTokenArgs(
    address localToken
  ) internal pure returns (ConnectionPolicyStore.SubsetTokenArgs memory) {
    return ConnectionPolicyStore.SubsetTokenArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, localToken: localToken
    });
  }

  function _addSubsetToken(
    bytes32 policyName,
    uint64 remoteChainSelector,
    address localToken
  ) internal {
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = ConnectionPolicyStore.SubsetTokenArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, localToken: localToken
    });
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }
}
