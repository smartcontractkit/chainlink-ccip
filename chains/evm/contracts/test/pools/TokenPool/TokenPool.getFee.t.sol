// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getFee is TokenPoolV2Setup {
  function test_getFee_DefaultFinality() public {
    uint16 defaultFeeBps = 250; // 2.50%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 75,
      customBlockConfirmationFeeUSDCents: 125,
      defaultBlockConfirmationTransferFeeBps: defaultFeeBps,
      customBlockConfirmationTransferFeeBps: 400
    });

    vm.startPrank(OWNER);
    _applyFeeConfig(feeConfig);

    uint256 amount = 1_000e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), 0, "");

    assertEq(usdFeeCents, feeConfig.defaultBlockConfirmationFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, defaultFeeBps);
  }

  function test_getFee_CustomFinality() public {
    uint16 customFeeBps = 400; // 4%
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 80,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 0,
      customBlockConfirmationTransferFeeBps: customFeeBps
    });

    vm.startPrank(OWNER);
    _applyFeeConfig(feeConfig);

    uint256 amount = 1_500e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), 5, "");

    assertEq(usdFeeCents, feeConfig.customBlockConfirmationFeeUSDCents);
    assertEq(destGasOverhead, feeConfig.destGasOverhead);
    assertEq(destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(tokenFeeBps, customFeeBps);
  }

  function test_getFee_NoConfig() public view {
    uint256 amount = 777e6;
    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps) =
      s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), 0, "");

    assertEq(usdFeeCents, 0);
    assertEq(destGasOverhead, 0);
    assertEq(destBytesOverhead, 0);
    assertEq(tokenFeeBps, 0);
  }

  // Reverts

  function test_getFee_RevertWhen_InvalidMinBlockConfirmation() public {
    uint16 minBlockConfirmation = 10;

    vm.startPrank(OWNER);
    // Set custom block confirmation config with minimum of 10 blocks
    s_tokenPool.applyCustomBlockConfirmationConfigUpdates(
      minBlockConfirmation, new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](0)
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 75,
      customBlockConfirmationFeeUSDCents: 125,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200
    });
    _applyFeeConfig(feeConfig);
    vm.stopPrank();

    uint256 amount = 1_000e6;
    uint16 requestedBlockConfirmation = 5; // Less than minimum of 10

    vm.expectRevert(
      abi.encodeWithSelector(
        TokenPool.InvalidMinBlockConfirmation.selector, requestedBlockConfirmation, minBlockConfirmation
      )
    );
    s_tokenPool.getFee(address(s_token), DEST_CHAIN_SELECTOR, amount, address(0), requestedBlockConfirmation, "");
  }

  function _applyFeeConfig(
    IPoolV2.TokenTransferFeeConfig memory feeConfig
  ) internal {
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }
}
