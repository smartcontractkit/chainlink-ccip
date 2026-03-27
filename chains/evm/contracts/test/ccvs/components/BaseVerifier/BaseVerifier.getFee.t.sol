// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../../../libraries/Client.sol";
import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getFee is BaseVerifierSetup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
      s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(feeUSDCents, DEFAULT_CCV_FEE_USD_CENTS);
    assertEq(gasForVerification, DEFAULT_CCV_GAS_LIMIT);
    assertEq(payloadSizeBytes, DEFAULT_CCV_PAYLOAD_SIZE);
  }

  function test_getFee_WithWaitForSafeFlag() public {
    // Configure chain to allow the WAIT_FOR_SAFE flag.
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    configs[0].allowedFinalityConfig = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    s_baseVerifier.applyRemoteChainConfigUpdates(configs);

    Client.EVM2AnyMessage memory message;
    (uint256 feeUSDCents,,) = s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec.WAIT_FOR_SAFE_FLAG);
    assertEq(feeUSDCents, DEFAULT_CCV_FEE_USD_CENTS);
  }

  function test_getFee_WithAllowedBlockDepth() public {
    // Configure chain to allow up to block depth 10.
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    configs[0].allowedFinalityConfig = FinalityCodec._encodeBlockDepth(10);
    s_baseVerifier.applyRemoteChainConfigUpdates(configs);

    Client.EVM2AnyMessage memory message;
    // Request 10 blocks — meets the minimum of 10 (requesting at least the minimum is allowed).
    (uint256 feeUSDCents,,) =
      s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec._encodeBlockDepth(10));
    assertEq(feeUSDCents, DEFAULT_CCV_FEE_USD_CENTS);
  }

  function test_getFee_RevertWhen_BlockDepthBelowMinimum() public {
    // Configure chain to require at least block depth 10.
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    configs[0].allowedFinalityConfig = FinalityCodec._encodeBlockDepth(10);
    s_baseVerifier.applyRemoteChainConfigUpdates(configs);

    Client.EVM2AnyMessage memory message;
    // Request 5 blocks — below the minimum of 10.
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec._encodeBlockDepth(5),
        FinalityCodec._encodeBlockDepth(10)
      )
    );
    s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec._encodeBlockDepth(5));
  }

  function test_getFee_RevertWhen_BlockDepthRequestedButOnlyFinalityAllowed() public {
    // Default setup has finalityConfig = FinalityCodec.WAIT_FOR_FINALITY_FLAG → only WAIT_FOR_FINALITY accepted.
    Client.EVM2AnyMessage memory message;
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec._encodeBlockDepth(1),
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      )
    );
    s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec._encodeBlockDepth(1));
  }

  function test_getFee_RevertWhen_SafeFlagNotAllowed() public {
    // Default setup only allows finality (FinalityCodec.WAIT_FOR_FINALITY_FLAG) — WAIT_FOR_SAFE is not in allowed flags.
    Client.EVM2AnyMessage memory message;
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec.WAIT_FOR_SAFE_FLAG,
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      )
    );
    s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  function test_getFee_RevertWhen_DestinationNotSupported() public {
    uint64 wrongDestChainSelector = DEST_CHAIN_SELECTOR + 1;
    Client.EVM2AnyMessage memory message;

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.RemoteChainNotSupported.selector, wrongDestChainSelector));
    s_baseVerifier.getFee(wrongDestChainSelector, message, "", FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
