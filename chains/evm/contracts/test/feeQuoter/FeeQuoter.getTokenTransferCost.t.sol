// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Client} from "../../libraries/Client.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../../libraries/USDPriceWith18Decimals.sol";
import {FeeQuoterFeeSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_getTokenTransferCost is FeeQuoterFeeSetup {
  using USDPriceWith18Decimals for uint224;

  address internal s_selfServeTokenDefaultPricing = makeAddr("self-serve-token-default-pricing");

  function test_NoTokenTransferChargesZeroFee() public view {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    assertEq(0, feeUSDWei);
    assertEq(0, destGasOverhead);
    assertEq(0, destBytesOverhead);
  }

  function test_getTokenTransferCost_selfServeUsesDefaults() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_selfServeTokenDefaultPricing, 1000);

    // Get config to assert it isn't set
    FeeQuoter.TokenTransferFeeConfig memory transferFeeConfig =
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    assertFalse(transferFeeConfig.isEnabled);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    // Assert that the default values are used
    assertEq(uint256(DEFAULT_TOKEN_FEE_USD_CENTS) * 1e16, feeUSDWei);
    assertEq(DEFAULT_TOKEN_DEST_GAS_OVERHEAD, destGasOverhead);
    assertEq(DEFAULT_TOKEN_BYTES_OVERHEAD, destBytesOverhead);
  }

  function test_SmallTokenTransferChargesMinFeeAndGas() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1000);
    FeeQuoter.TokenTransferFeeConfig memory transferFeeConfig =
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    assertEq(_configUSDCentToWei(transferFeeConfig.feeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_ZeroAmountTokenTransferChargesMinFeeAndGas() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 0);
    FeeQuoter.TokenTransferFeeConfig memory transferFeeConfig =
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    assertEq(_configUSDCentToWei(transferFeeConfig.feeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_LargeTokenTransferChargesMaxFeeAndGas() public view {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1e36);
    FeeQuoter.TokenTransferFeeConfig memory transferFeeConfig =
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, message.tokenAmounts[0].token);

    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    assertEq(_configUSDCentToWei(transferFeeConfig.feeUSDCents), feeUSDWei);
    assertEq(transferFeeConfig.destGasOverhead, destGasOverhead);
    assertEq(transferFeeConfig.destBytesOverhead, destBytesOverhead);
  }

  function test_ZeroFeeConfigChargesMinFee() public {
    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs = _generateTokenTransferFeeConfigArgs(1, 1);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = s_sourceFeeToken;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = FeeQuoter.TokenTransferFeeConfig({
      feeUSDCents: 0,
      destGasOverhead: 0,
      destBytesOverhead: uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES),
      isEnabled: true
    });
    s_feeQuoter.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0)
    );

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1e36);
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    // if token charges 0 bps, it should cost minFee to transfer
    assertEq(
      _configUSDCentToWei(tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig.feeUSDCents),
      feeUSDWei
    );
    assertEq(0, destGasOverhead);
    assertEq(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES, destBytesOverhead);
  }

  function test_MixedTokenTransferFee() public view {
    address[3] memory testTokens = [s_sourceFeeToken, s_sourceRouter.getWrappedNative(), CUSTOM_TOKEN];
    FeeQuoter.TokenTransferFeeConfig[3] memory tokenTransferFeeConfigs = [
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[0]),
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[1]),
      s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[2])
    ];

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](3),
      feeToken: s_sourceRouter.getWrappedNative(),
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
    uint256 expectedTotalGas = 0;
    uint256 expectedTotalBytes = 0;

    for (uint256 i = 0; i < testTokens.length; ++i) {
      message.tokenAmounts[i] = Client.EVMTokenAmount({token: testTokens[i], amount: 1e14});
      FeeQuoter.TokenTransferFeeConfig memory tokenTransferFeeConfig =
        s_feeQuoter.getTokenTransferFeeConfig(DEST_CHAIN_SELECTOR, testTokens[i]);

      expectedTotalGas += tokenTransferFeeConfig.destGasOverhead == 0
        ? DEFAULT_TOKEN_DEST_GAS_OVERHEAD
        : tokenTransferFeeConfig.destGasOverhead;
      expectedTotalBytes += tokenTransferFeeConfig.destBytesOverhead == 0
        ? DEFAULT_TOKEN_BYTES_OVERHEAD
        : tokenTransferFeeConfig.destBytesOverhead;
    }
    (uint256 feeUSDWei, uint32 destGasOverhead, uint32 destBytesOverhead) =
      s_feeQuoter.getTokenTransferCost(DEST_CHAIN_SELECTOR, message.tokenAmounts);

    uint256 expectedFeeUSDWei = 0;
    for (uint256 i = 0; i < testTokens.length; ++i) {
      expectedFeeUSDWei += _configUSDCentToWei(
        tokenTransferFeeConfigs[i].feeUSDCents == 0
          ? DEFAULT_TOKEN_FEE_USD_CENTS
          : tokenTransferFeeConfigs[i].feeUSDCents
      );
    }

    assertEq(expectedFeeUSDWei, feeUSDWei, "wrong feeUSDWei 1");
    assertEq(expectedTotalGas, destGasOverhead, "wrong destGasOverhead 1");
    assertEq(expectedTotalBytes, destBytesOverhead, "wrong destBytesOverhead 1");
  }

  function _calcUSDValueFromTokenAmount(uint224 tokenPrice, uint256 tokenAmount) internal pure returns (uint256) {
    return (tokenPrice * tokenAmount) / 1e18;
  }
}
