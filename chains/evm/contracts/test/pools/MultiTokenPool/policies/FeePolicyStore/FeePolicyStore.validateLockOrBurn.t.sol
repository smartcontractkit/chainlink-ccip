// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_validateLockOrBurn is FeePolicyStoreSetup {
  uint256 internal constant AMOUNT = 1e18;

  function test_validateLockOrBurn() public view {
    assertEq(
      (AMOUNT * _getFeeConfig().finalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, MIN_DEST_BYTES_OVERHEAD
      )
    );
  }

  function test_validateLockOrBurn_FastFinality() public view {
    assertEq(
      (AMOUNT * _getFeeConfig().fastFinalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), true, MIN_DEST_BYTES_OVERHEAD
      )
    );
  }

  function test_validateLockOrBurn_PolicyChainRecord() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));

    assertEq(
      (AMOUNT * _getOtherFeeConfig().finalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, MIN_DEST_BYTES_OVERHEAD
      )
    );
  }

  function test_validateLockOrBurn_NoChargeTokenFee() public view {
    assertEq(
      0,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false, MIN_DEST_BYTES_OVERHEAD
      )
    );
    assertEq(
      0,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), false, MIN_DEST_BYTES_OVERHEAD
      )
    );
  }

  function test_validateLockOrBurn_ZeroFees() public {
    FeePolicyStore.FeeConfig memory zeroFeeConfig;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, zeroFeeConfig));

    assertEq(
      0,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, MIN_DEST_BYTES_OVERHEAD
      )
    );
    assertEq(
      0,
      s_feeStore.validateLockOrBurn(
        s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), true, MIN_DEST_BYTES_OVERHEAD
      )
    );
  }

  function test_validateLockOrBurn_RoundsDown() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig();
    feeConfig.finalityTransferFeeBps = 3_333;
    feeConfig.fastFinalityTransferFeeBps = 6_667;
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, feeConfig));

    assertEq(
      0, s_feeStore.validateLockOrBurn(s_token, DEST_CHAIN_SELECTOR, 3, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, 0)
    );
    assertEq(
      2, s_feeStore.validateLockOrBurn(s_token, DEST_CHAIN_SELECTOR, 3, FinalityCodec._encodeBlockDepth(10), true, 0)
    );
  }

  function test_validateLockOrBurn_InlineDestOverheadRecord() public {
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getOtherDestOverheadConfig())
    );
    uint32 destBytesOverhead = _getOtherDestOverheadConfig().destBytesOverhead;

    // The inline byte limit overrides the policy default.
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, destBytesOverhead
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector,
        s_token,
        DEST_CHAIN_SELECTOR,
        uint256(destBytesOverhead) + 1,
        destBytesOverhead
      )
    );
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, uint256(destBytesOverhead) + 1
    );
  }

  function testFuzz_validateLockOrBurn_PoolDataAtDestBytesOverhead(
    uint32 destBytesOverhead
  ) public {
    destBytesOverhead = uint32(bound(destBytesOverhead, MIN_DEST_BYTES_OVERHEAD, type(uint32).max));
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig = _getDestOverheadConfig();
    destOverheadConfig.destBytesOverhead = destBytesOverhead;
    _setInlineDestOverheadConfig(_getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, destOverheadConfig));

    // The byte limit is inclusive, whether or not a token fee is charged.
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, destBytesOverhead
    );
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false, destBytesOverhead
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector,
        s_token,
        DEST_CHAIN_SELECTOR,
        uint256(destBytesOverhead) + 1,
        destBytesOverhead
      )
    );
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, uint256(destBytesOverhead) + 1
    );
  }

  // Reverts

  function test_validateLockOrBurn_RevertWhen_FeeRecordNotFound() public {
    // Both records are missing; report the missing fee record first.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.validateLockOrBurn(
      s_otherToken, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, 0
    );

    // Disabling token fees must not bypass the missing fee record check.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.validateLockOrBurn(
      s_otherToken, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false, 0
    );
  }

  function test_validateLockOrBurn_RevertWhen_DestOverheadRecordNotFound() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.validateLockOrBurn(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, 0);

    // Destination overheads are still required when no token fee is charged.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.validateLockOrBurn(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false, 0);
  }

  function test_validateLockOrBurn_RevertWhen_PoolDataTooLarge() public {
    uint256 destPoolDataLength = uint256(MIN_DEST_BYTES_OVERHEAD) + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector,
        s_token,
        DEST_CHAIN_SELECTOR,
        destPoolDataLength,
        MIN_DEST_BYTES_OVERHEAD
      )
    );
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true, destPoolDataLength
    );

    // The V1 path must enforce the same byte limit without charging a token fee.
    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector,
        s_token,
        DEST_CHAIN_SELECTOR,
        destPoolDataLength,
        MIN_DEST_BYTES_OVERHEAD
      )
    );
    s_feeStore.validateLockOrBurn(
      s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false, destPoolDataLength
    );
  }
}
