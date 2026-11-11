// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getLaneVerificationBatch is SecurityPolicyStoreSetup {
  /// @notice getLaneVerificationBatch returns the hook and the two CCV lists of a lane, and empty for a lane with no
  /// security record.
  function test_getLaneVerificationBatch() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), _arrayOf(s_ccv2), s_hook);

    IPoolConfigManager.LaneKey[] memory lanes = _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR);
    (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs) =
      s_store.getLaneVerificationBatch(lanes);

    assertEq(s_hook, hooks[0]);
    assertEq(abi.encode(_ccvsOf(s_ccv1, s_ccv2)), abi.encode(outboundCCVs[0]));
    assertEq(abi.encode(_arrayOf(s_ccv2)), abi.encode(inboundCCVs[0]));

    // A lane that does not exist gives an empty entry, because the store does not resolve the lane.
    (hooks, outboundCCVs, inboundCCVs) = s_store.getLaneVerificationBatch(_lanesOf(s_tokenB, DEST_CHAIN_SELECTOR));
    assertEq(address(0), hooks[0]);
    assertEq(0, outboundCCVs[0].length);
    assertEq(0, inboundCCVs[0].length);
  }

  /// @notice A lane without an inline record uses the chain record of its policy, and a lane without a chain record
  /// uses the policy default record.
  function test_getLaneVerificationBatch_PolicyRecords() public {
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setSecurityPolicyConfig(
      DEFAULT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _arrayOf(s_ccv2), _arrayOf(s_ccv1), address(0)
    );
    _setTokenSecurityPolicy(s_tokenA, DEFAULT_POLICY);

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs) =
      s_store.getLaneVerificationBatch(lanes);

    assertEq(s_hook, hooks[0]);
    assertEq(abi.encode(_arrayOf(s_ccv1)), abi.encode(outboundCCVs[0]));
    assertEq(abi.encode(_arrayOf(s_ccv2)), abi.encode(inboundCCVs[0]));
    assertEq(address(0), hooks[1]);
    assertEq(abi.encode(_arrayOf(s_ccv2)), abi.encode(outboundCCVs[1]));
    assertEq(abi.encode(_arrayOf(s_ccv1)), abi.encode(inboundCCVs[1]));
  }
}
