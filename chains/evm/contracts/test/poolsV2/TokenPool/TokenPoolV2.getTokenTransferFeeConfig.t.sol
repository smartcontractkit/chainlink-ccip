// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getTokenTransferFeeConfig is TokenPoolV2Setup {
  function test_getTokenTransferFeeConfig() public {
    // Set up a fee config first.
    TokenPool.TokenTransferFeeConfig memory feeConfig = TokenPool.TokenTransferFeeConfig({
      destGasOverhead: 50000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      feeUSDCents: 100, // $1.00
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config
    Client.EVM2AnyMessage memory message;
    (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents) =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(isEnabled, true);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES);
    assertEq(feeUSDCents, feeConfig.feeUSDCents);
  }

  function test_getTokenTransferFeeConfig_DefaultConfig() public {
    // Test getting config when none is set (should return default values).
    Client.EVM2AnyMessage memory message;
    (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents) =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(isEnabled, false);
    assertEq(destGasOverhead, 0);
    assertEq(destBytesOverhead, 0);
    assertEq(feeUSDCents, 0);
  }

  function test_getTokenTransferFeeConfig_DeleteConfig() public {
    uint64[] memory toDelete = new uint64[](1);
    toDelete[0] = DEST_CHAIN_SELECTOR;
    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigDeleted(DEST_CHAIN_SELECTOR);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(new TokenPool.TokenTransferFeeConfigArgs[](0), toDelete);

    // Test getting the disabled config
    Client.EVM2AnyMessage memory message;
    (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents) =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(isEnabled, false);
    assertEq(destGasOverhead, 0);
    assertEq(destBytesOverhead, 0);
    assertEq(feeUSDCents, 0);
  }
}
