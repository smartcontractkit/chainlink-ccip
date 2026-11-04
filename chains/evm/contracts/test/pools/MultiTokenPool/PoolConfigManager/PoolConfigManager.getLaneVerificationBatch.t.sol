// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token A has an inline security record for the destination chain with a hook, one outbound CCV, and one
/// inbound CCV.
contract PoolConfigManager_getLaneVerificationBatch is PoolConfigManagerSetup {
  address internal s_hook = makeAddr("hook");
  address[] internal s_outboundCCVs;
  address[] internal s_inboundCCVs;

  function setUp() public virtual override {
    super.setUp();
    s_outboundCCVs = _arrayOf(makeAddr("outboundCCV"));
    s_inboundCCVs = _arrayOf(makeAddr("inboundCCV"));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, s_outboundCCVs, s_inboundCCVs, s_hook);
  }

  /// @notice Returns the resolved security config for each lane. Missing lanes return zero values even when a
  /// security record is configured.
  function test_getLaneVerificationBatch() public {
    _setSecurityPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, s_outboundCCVs, s_inboundCCVs, s_hook);

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    (address[] memory hooks, address[][] memory laneOutboundCCVs, address[][] memory laneInboundCCVs) =
      s_poolConfigManager.getLaneVerificationBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(s_hook, hooks[0]);
    assertEq(s_outboundCCVs, laneOutboundCCVs[0]);
    assertEq(s_inboundCCVs, laneInboundCCVs[0]);
    // Token B uses the empty policy default record on the destination chain.
    assertEq(address(0), hooks[1]);
    assertEq(0, laneOutboundCCVs[1].length);
    assertEq(0, laneInboundCCVs[1].length);
    // Token B has no lane to the other chain.
    assertEq(address(0), hooks[2]);
    assertEq(0, laneOutboundCCVs[2].length);
    assertEq(0, laneInboundCCVs[2].length);
  }

  /// @notice This view ignores requested finality. It returns the same security config for default finality, FTF,
  /// and requests that FinalityCodec would reject.
  function test_getLaneVerificationBatch_FinalityHasNoEffect() public view {
    bytes4[3] memory finalities =
      [_getFastFinalityConfig(), FinalityCodec.WAIT_FOR_SAFE_FLAG, FinalityCodec._encodeBlockDepthAndSafeFlag(10)];
    for (uint256 i = 0; i < finalities.length; ++i) {
      _assertSecurityRecordOfTokenA(finalities[i]);
    }
  }

  function _assertSecurityRecordOfTokenA(
    bytes4 requestedFinalityConfig
  ) internal view {
    (address[] memory hooks, address[][] memory laneOutboundCCVs, address[][] memory laneInboundCCVs) =
      s_poolConfigManager.getLaneVerificationBatch(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), requestedFinalityConfig);
    assertEq(s_hook, hooks[0]);
    assertEq(s_outboundCCVs, laneOutboundCCVs[0]);
    assertEq(s_inboundCCVs, laneInboundCCVs[0]);
  }
}
