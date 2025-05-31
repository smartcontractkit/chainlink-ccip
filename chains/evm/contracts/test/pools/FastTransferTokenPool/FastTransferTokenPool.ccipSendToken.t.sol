// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_ccipSendToken_Test is FastTransferTokenPoolSetup {
  struct TestParams {
    uint64 chainSelector;
    uint256 amount;
    bytes receiver;
    bytes32 mockMessageId;
    uint16 fastFeeBpsExpected;
  }

  function _setupTestParams(
    uint64 chainSelector
  ) internal pure returns (TestParams memory) {
    uint256 amount = 100e18;
    return TestParams({
      chainSelector: chainSelector,
      amount: amount,
      receiver: abi.encodePacked(address(0x5)),
      mockMessageId: keccak256("mockMessageId"),
      fastFeeBpsExpected: FAST_FEE_BPS
    });
  }

  function _setupChainConfig(
    uint64 chainSelector,
    bytes4 chainFamily,
    uint32 settlementGas,
    bytes memory extraArgs
  ) internal {
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: chainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: settlementGas,
      chainFamilySelector: chainFamily,
      customExtraArgs: extraArgs
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
    s_pool.updateFillerAllowList(chainSelector, addFillers, new address[](0));

    _setupRateLimiter(chainSelector);
  }

  function _setupRateLimiter(
    uint64 chainSelector
  ) internal {
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: chainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_pool.applyChainUpdates(new uint64[](0), chainUpdate);
  }

  function _setupMocks(
    bytes32 mockMessageId
  ) internal {
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
  }

  function _createMessage(
    TestParams memory params,
    bytes memory extraArgs
  ) internal view returns (Client.EVM2AnyMessage memory) {
    return Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmount: params.amount,
          sourceDecimals: 18,
          fastTransferFeeBps: params.fastFeeBpsExpected,
          receiver: params.receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: extraArgs
    });
  }

  function _executeTest(TestParams memory params, bytes memory extraArgs) internal {
    _setupMocks(params.mockMessageId);

    Client.EVM2AnyMessage memory message = _createMessage(params, extraArgs);
    uint256 expectedFee = params.amount * params.fastFeeBpsExpected / 10_000;

    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, params.chainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(
      params.mockMessageId, params.chainSelector, params.amount, expectedFee, params.receiver
    );

    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes32 fillRequestId =
      s_pool.ccipSendToken{value: 1 ether}(address(0), params.chainSelector, params.amount, params.receiver, "");

    assertEq(fillRequestId, params.mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - params.amount);
    assertEq(s_token.balanceOf(address(s_pool)), params.amount);
  }

  function test_CcipSendToken_NativeFee() public {
    TestParams memory params = _setupTestParams(DEST_CHAIN_SELECTOR);

    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );

    _executeTest(params, extraArgs);
  }

  function test_CcipSendToken_NativeFee_ToSVM() public {
    TestParams memory params = _setupTestParams(SVM_CHAIN_SELECTOR);
    _executeTest(params, s_svmExtraArgsBytesEncoded);
  }

  function test_CcipSendToken_APTOS_WithSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_WITH_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_APTOS, SETTLEMENT_GAS_OVERHEAD, "");

    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );

    _executeTest(params, extraArgs);
  }

  function test_CcipSendToken_APTOS_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_ZERO_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    bytes memory customExtraArgs = abi.encode("aptosCustomExtraArgs");
    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_APTOS, 0, customExtraArgs);
    _executeTest(params, customExtraArgs);
  }

  function test_CcipSendToken_SUI_WithSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("SUI_WITH_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_SUI, SETTLEMENT_GAS_OVERHEAD, "");

    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );

    _executeTest(params, extraArgs);
  }

  function test_CcipSendToken_SUI_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("SUI_ZERO_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    bytes memory customExtraArgs = abi.encode("suiCustomExtraArgs");
    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_SUI, 0, customExtraArgs);
    _executeTest(params, customExtraArgs);
  }
}
