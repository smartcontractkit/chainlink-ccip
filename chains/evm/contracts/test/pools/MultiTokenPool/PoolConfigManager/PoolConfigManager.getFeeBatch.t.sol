// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getFeeBatch is PoolConfigManagerSetup {
  /// @notice Each result uses its resolved fee and overhead records. Token B uses policy fees and inline overheads.
  /// Token A uses inline fees and policy default overheads.
  function test_getFeeBatch() public {
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getFeeConfig(100, 10));
    _setInlineDestOverheadConfig(s_tokenB, DEST_CHAIN_SELECTOR, 200_000, 256);
    _setInlineFeeConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getFeeConfig(300, 30));
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR});

    IPoolConfigManager.LaneFee[] memory fees =
      s_poolConfigManager.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(2, fees.length);
    assertEq(
      abi.encode(
        IPoolConfigManager.LaneFee({
          feeUSDCents: 100, destGasOverhead: 200_000, destBytesOverhead: 256, tokenFeeBps: 10, isEnabled: true
        })
      ),
      abi.encode(fees[0])
    );
    assertEq(
      abi.encode(
        IPoolConfigManager.LaneFee({
          feeUSDCents: 300,
          destGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
          destBytesOverhead: MIN_DEST_BYTES_OVERHEAD,
          tokenFeeBps: 30,
          isEnabled: true
        })
      ),
      abi.encode(fees[1])
    );
  }

  /// @notice Zero fees remain enabled when both fee and destination overhead records exist. The result includes the
  /// configured overheads.
  function test_getFeeBatch_ZeroFeesAreEnabled() public view {
    IPoolConfigManager.LaneFee memory fee = _getFee(s_tokenB, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(
      abi.encode(
        IPoolConfigManager.LaneFee({
          feeUSDCents: 0,
          destGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
          destBytesOverhead: MIN_DEST_BYTES_OVERHEAD,
          tokenFeeBps: 0,
          isEnabled: true
        })
      ),
      abi.encode(fee)
    );
  }

  /// @notice Missing fee or destination overhead records return zero values with isEnabled set to false.
  /// MultiTokenPool.getFee rejects this result with LaneFeeNotFound. Each token still has a lane and a nonzero
  /// record of the other type.
  function test_getFeeBatch_MissingRecordGivesZeros() public {
    _setInlineDestOverheadConfig(s_tokenB, DEST_CHAIN_SELECTOR, 200_000, 256);
    _setTokenFeePolicy(s_tokenB, bytes32(0));
    _setInlineFeeConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getFeeConfig(300, 30));
    _setTokenDestOverheadPolicy(s_tokenA, bytes32(0));
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR});

    IPoolConfigManager.LaneFee[] memory fees =
      s_poolConfigManager.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory zeroFee;
    assertTrue(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).hasConnection);
    assertTrue(_getLaneConnection(s_tokenA, DEST_CHAIN_SELECTOR).hasConnection);
    assertEq(abi.encode(zeroFee), abi.encode(fees[0]));
    assertEq(abi.encode(zeroFee), abi.encode(fees[1]));
  }

  /// @notice Fee and overhead records alone do not create a lane. The manager returns zero values for token B on
  /// the other chain.
  function test_getFeeBatch_NoLaneIsZeroed() public {
    _setInlineFeeConfig(s_tokenB, OTHER_CHAIN_SELECTOR, _getFeeConfig(100, 10));
    _setInlineDestOverheadConfig(s_tokenB, OTHER_CHAIN_SELECTOR, 200_000, 256);
    IPoolConfigManager.LaneKey[] memory lanes = _lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR);
    assertFalse(_getLaneConnection(s_tokenB, OTHER_CHAIN_SELECTOR).hasConnection);
    IPoolConfigManager.LaneFee memory storeFee = s_feeStore.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG)[0];
    assertEq(100, storeFee.feeUSDCents);
    assertTrue(storeFee.isEnabled);

    IPoolConfigManager.LaneFee memory fee =
      s_poolConfigManager.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG)[0];

    IPoolConfigManager.LaneFee memory zeroFee;
    assertEq(abi.encode(zeroFee), abi.encode(fee));
  }

  /// @notice Fee quotes do not require a remote pool record. Sending still reverts in lockOrBurn if that record is
  /// missing.
  function test_getFeeBatch_WithoutRemotePoolRecord() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(100, 10);
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, feeConfig);
    _setTokenRemotePoolPolicy(s_tokenB, bytes32(0));

    IPoolConfigManager.LaneFee memory fee = _getFee(s_tokenB, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(feeConfig.finalityFeeUSDCents, fee.feeUSDCents);
    assertTrue(fee.isEnabled);
  }

  /// @notice FTF requests use the FTF fee fields when the token profile allows that finality.
  function test_getFeeBatch_FastFinality() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(100, 10);
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, feeConfig);
    _setAllowedFinalityConfig(s_tokenB, _getFastFinalityConfig());

    IPoolConfigManager.LaneFee memory fee = _getFee(s_tokenB, DEST_CHAIN_SELECTOR, _getFastFinalityConfig());
    assertEq(feeConfig.fastFinalityFeeUSDCents, fee.feeUSDCents);
    assertEq(feeConfig.fastFinalityTransferFeeBps, fee.tokenFeeBps);
  }

  // Reverts

  /// @notice An unsupported finality request reverts even when the lane does not exist, as in TokenPool.
  function test_getFeeBatch_RevertWhen_InvalidRequestedFinalityWithoutLane() public {
    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, _getFastFinalityConfig(), bytes4(0))
    );
    s_poolConfigManager.getFeeBatch(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), _getFastFinalityConfig());
  }

  /// @dev The pool asks for the fee of one lane.
  function _getFee(
    address localToken,
    uint64 destChainSelector,
    bytes4 requestedFinalityConfig
  ) internal view returns (IPoolConfigManager.LaneFee memory) {
    return s_poolConfigManager.getFeeBatch(_lanesOf(localToken, destChainSelector), requestedFinalityConfig)[0];
  }
}
