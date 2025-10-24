// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getFee is TokenPoolV2Setup {
  function test_getFee_DefaultFinality() public {
    Client.EVM2AnyMessage memory message = _buildMessage();
    uint16 defaultFeeBps = 250; // 2.50%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultFinalityFeeUSDCents: 75,
      customFinalityFeeUSDCents: 125,
      defaultFinalityTransferFeeBps: defaultFeeBps,
      customFinalityTransferFeeBps: 400,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    vm.startPrank(OWNER);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(usdFeeCents, feeConfig.defaultFinalityFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, defaultFeeBps);
  }

  function test_getFee_CustomFinality() public {
    Client.EVM2AnyMessage memory message = _buildMessage();
    uint16 customFeeBps = 400; // 4%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultFinalityFeeUSDCents: 80,
      customFinalityFeeUSDCents: 150,
      defaultFinalityTransferFeeBps: 0,
      customFinalityTransferFeeBps: customFeeBps,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    vm.startPrank(OWNER);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, message, 5, "");

    assertEq(usdFeeCents, feeConfig.customFinalityFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, customFeeBps);
  }

  function test_getFee_DisabledConfig() public view {
    Client.EVM2AnyMessage memory message = _buildMessage();

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(usdFeeCents, 0);
    assertEq(destGasOverhead, 0);
    assertEq(destBytesOverhead, 0);
    assertEq(tokenFeeBps, 0);
  }

  function _buildMessage() internal pure returns (Client.EVM2AnyMessage memory message) {
    message.receiver = abi.encode(address(0xBEEF));
    message.data = "";
    message.tokenAmounts = new Client.EVMTokenAmount[](0);
    message.feeToken = address(0);
    message.extraArgs = "";
    return message;
  }
}
