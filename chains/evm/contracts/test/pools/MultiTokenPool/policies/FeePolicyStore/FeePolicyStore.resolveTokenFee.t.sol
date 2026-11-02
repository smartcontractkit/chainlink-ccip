// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_resolveTokenFee is FeePolicyStoreSetup {
  uint256 internal constant AMOUNT = 1e18;

  function test_resolveTokenFee() public view {
    assertEq(
      (AMOUNT * _getFeeConfig().finalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true)
    );
  }

  function test_resolveTokenFee_FastFinality() public view {
    assertEq(
      (AMOUNT * _getFeeConfig().fastFinalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), true)
    );
  }

  function test_resolveTokenFee_PolicyChainRecord() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));

    // The policy chain record serves the chain instead of the policy default record.
    assertEq(
      (AMOUNT * _getOtherFeeConfig().finalityTransferFeeBps) / BPS_DIVIDER,
      s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true)
    );
  }

  function test_resolveTokenFee_NoChargeTokenFee() public view {
    assertEq(
      0, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false)
    );
    assertEq(
      0, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), false)
    );
  }

  function test_resolveTokenFee_DisabledFeeRecord() public {
    FeePolicyStore.FeeConfig memory disabledFeeConfig;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabledFeeConfig));

    assertEq(
      0, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true)
    );
    assertEq(
      0, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec._encodeBlockDepth(10), true)
    );
  }

  function test_resolveTokenFee_RoundsDown() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig();
    feeConfig.finalityTransferFeeBps = 3_333;
    feeConfig.fastFinalityTransferFeeBps = 6_667;
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, feeConfig));

    assertEq(0, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, 3, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true));
    assertEq(2, s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, 3, FinalityCodec._encodeBlockDepth(10), true));
  }

  // Reverts

  function test_resolveTokenFee_RevertWhen_FeeRecordNotFound() public {
    // s_otherToken has no record of either type, and the fee record is checked first.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.resolveTokenFee(s_otherToken, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true);

    // A send without a token fee also needs the record.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.resolveTokenFee(s_otherToken, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false);
  }

  function test_resolveTokenFee_RevertWhen_DestOverheadRecordNotFound() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true);

    // A send without a token fee also needs the record.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, false);

    // A send with a disabled fee record also needs the record.
    FeePolicyStore.FeeConfig memory disabledFeeConfig;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabledFeeConfig));

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true);
  }
}
