// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applySubsetTokenUpdates is PoolConfigManagerRolesSetup {
  address internal s_tokenC;

  /// @dev The DEFAULT connection policy restricts the other chain to subset tokens. Tokens B and C use the DEFAULT
  /// policies.
  function setUp() public virtual override {
    super.setUp();
    _setConnectionPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    // A remote pool policy has no default record, so the DEFAULT policy needs a record on the other chain for the lane
    // check of the subset tokens that the tests add there.
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    s_tokenC = _deployToken("C");
    _registerToken(s_tokenC);
  }

  /// @notice Adding subset tokens enables their lanes on the other chain. Each token has all required records.
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

  /// @notice Removing a subset token does not check lane completeness. Token B can leave the subset despite having
  /// no fee record.
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

  /// @notice Each allowed role authorizes the call on its own.
  function test_applySubsetTokenUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applySubsetTokenUpdates,
        (new ConnectionPolicyStore.SubsetTokenArgs[](0), new ConnectionPolicyStore.SubsetTokenArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  // Reverts

  /// @notice The connection store rejects a token not assigned to the policy before the manager checks for a remote
  /// pool record.
  function test_applySubsetTokenUpdates_RevertWhen_StoreRevertsBeforeLaneCheck() public {
    address unknownToken = makeAddr("unknownToken");
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(unknownToken);

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.TokenNotInPolicy.selector, unknownToken, DEFAULT_POLICY)
    );
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice Lane validation reports the first missing record. Token B has no fee or security record, so the fee
  /// check reverts first.
  function test_applySubsetTokenUpdates_RevertWhen_FeeRecordNotFound() public {
    _setTokenFeePolicy(s_tokenB, bytes32(0));
    _setTokenSecurityPolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenB);

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR));
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice The security check runs last. Token B has every required record except its security record.
  function test_applySubsetTokenUpdates_RevertWhen_SecurityRecordNotFound() public {
    _setTokenSecurityPolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    tokensToAdd[0] = _getSubsetTokenArgs(s_tokenB);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);
  }

  /// @notice Every added token is checked. Token C has all required records. Token B is added second and has no
  /// destination overhead record, so the error identifies token B.
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

  /// @notice Calls without an allowed role revert.
  function test_applySubsetTokenUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applySubsetTokenUpdates,
        (new ConnectionPolicyStore.SubsetTokenArgs[](0), new ConnectionPolicyStore.SubsetTokenArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
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
