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
    uint16 fastFeeBpsExpected;
    uint256 amount;
    bytes32 mockMessageId;
    bytes receiver;
  }

  function _setupTestParams(
    uint64 chainSelector
  ) internal pure returns (TestParams memory) {
    return TestParams({
      chainSelector: chainSelector,
      amount: SOURCE_AMOUNT,
      receiver: abi.encodePacked(address(0x5)),
      mockMessageId: keccak256("mockMessageId"),
      fastFeeBpsExpected: FAST_FEE_FILLER_BPS
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
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: settlementGas,
      chainFamilySelector: chainFamily,
      customExtraArgs: extraArgs
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
    s_pool.updateFillerAllowList(addFillers, new address[](0));

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
          fastTransferFillerFeeBps: params.fastFeeBpsExpected,
          fastTransferPoolFeeBps: 0,
          sourceDecimals: 18,
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
    bytes32 fillId = s_pool.computeFillId(params.mockMessageId, params.amount - expectedFee, 18, params.receiver);

    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, params.chainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: params.chainSelector,
      fillId: fillId,
      settlementId: params.mockMessageId,
      sourceAmountNetFee: params.amount - expectedFee,
      fastTransferFee: expectedFee,
      receiver: params.receiver
    });

    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes32 settlementId =
      s_pool.ccipSendToken{value: 1 ether}(params.chainSelector, params.amount, params.receiver, address(0), "");

    assertEq(settlementId, params.mockMessageId);
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

  function test_CcipSendToken_WithERC20FeeToken() public {
    address feeToken = address(s_token);
    uint256 balanceBefore = s_token.balanceOf(OWNER);
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));
    uint256 fakeFee = 1 ether;
    bytes32 fakeMessageId = keccak256("mockMessageId");
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(fakeFee));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(fakeMessageId)
    );

    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encode(RECEIVER), feeToken, "");

    uint256 expectedFastTransferFee = SOURCE_AMOUNT * FAST_FEE_FILLER_BPS / 10_000;
    uint256 expectedAmountNetFee = SOURCE_AMOUNT - expectedFastTransferFee;
    bytes32 expectedFillId = s_pool.computeFillId(fakeMessageId, expectedAmountNetFee, 18, abi.encode(RECEIVER));

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: DEST_CHAIN_SELECTOR,
      fillId: expectedFillId,
      settlementId: fakeMessageId,
      sourceAmountNetFee: expectedAmountNetFee,
      fastTransferFee: expectedFastTransferFee,
      receiver: abi.encode(RECEIVER)
    });

    bytes32 settlementId = s_pool.ccipSendToken(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encode(RECEIVER), feeToken, "");

    assertTrue(settlementId != bytes32(0));
    assertEq(s_token.balanceOf(OWNER), balanceBefore - SOURCE_AMOUNT - quote.ccipSettlementFee);
    // When using ERC20 fee token, both transfer amount and settlement fee go to the pool
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore + SOURCE_AMOUNT + quote.ccipSettlementFee);
  }

  function test_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encode(RECEIVER), address(0), "");
  }

  function test_CcipSendToken_WithPoolFee() public {
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 25; // 0.25%

    TestParams memory params = _setupTestParams(DEST_CHAIN_SELECTOR);
    params.fastFeeBpsExpected = fillerFeeBps + poolFeeBps; // Total fee

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    _setupMocks(params.mockMessageId);

    uint256 fillerFeeAmount = (SOURCE_AMOUNT * fillerFeeBps) / 10_000;
    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 totalFee = fillerFeeAmount + poolFeeAmount;
    uint256 amountNetTotalFee = SOURCE_AMOUNT - totalFee;

    bytes32 fillId = s_pool.computeFillId(params.mockMessageId, amountNetTotalFee, 18, params.receiver);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: params.chainSelector,
      fillId: fillId,
      settlementId: params.mockMessageId,
      sourceAmountNetFee: amountNetTotalFee,
      fastTransferFee: totalFee,
      receiver: params.receiver
    });

    uint256 balanceBefore = s_token.balanceOf(OWNER);

    bytes32 settlementId =
      s_pool.ccipSendToken{value: 1 ether}(params.chainSelector, params.amount, params.receiver, address(0), "");

    assertEq(settlementId, params.mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - params.amount);
    assertEq(s_token.balanceOf(address(s_pool)), params.amount);
  }

  function test_CcipSendToken_FeeQuote_WithPoolFee() public {
    uint16 fillerFeeBps = 50; // 0.5%
    uint16 poolFeeBps = 150; // 1.5%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    // Setup mocks for router calls
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));

    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encode(RECEIVER), address(0), "");

    uint256 expectedFillerFee = (SOURCE_AMOUNT * fillerFeeBps) / 10_000;
    uint256 expectedPoolFee = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 expectedTotalFee = expectedFillerFee + expectedPoolFee;

    assertEq(quote.fastTransferFee, expectedTotalFee);
  }

  function test_CcipSendToken_EqualFeeSplit() public {
    uint16 equalFee = 100; // 1% each for filler and pool

    TestParams memory params = _setupTestParams(DEST_CHAIN_SELECTOR);
    params.fastFeeBpsExpected = equalFee * 2; // Total fee is 2%

    // Update config to have equal fee split
    _updateConfigWithPoolFee(equalFee, equalFee);

    // Setup mocks for router calls
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));

    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encode(RECEIVER), address(0), "");

    uint256 expectedTotalFee = (SOURCE_AMOUNT * equalFee * 2) / 10_000;
    assertEq(quote.fastTransferFee, expectedTotalFee);
  }

  function _updateConfigWithPoolFee(uint16 fillerFeeBps, uint16 poolFeeBps) internal {
    vm.stopPrank();

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: fillerFeeBps,
      fastTransferPoolFeeBps: poolFeeBps,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    vm.prank(OWNER);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address[] memory fillersToAdd = new address[](1);
    fillersToAdd[0] = s_filler;
    vm.prank(OWNER);
    s_pool.updateFillerAllowList(fillersToAdd, new address[](0));

    vm.startPrank(OWNER);
  }
}
