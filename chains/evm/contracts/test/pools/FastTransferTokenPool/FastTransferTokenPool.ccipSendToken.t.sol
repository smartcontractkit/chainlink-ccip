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
    uint256 expectedFastTransferFee = params.amount * params.fastFeeBpsExpected / 10_000;
    uint256 expectedFillerFee = expectedFastTransferFee; // All fee goes to filler in basic tests
    uint256 expectedPoolFee = 0; // No pool fee in basic tests
    bytes32 fillId = s_pool.computeFillId(
      params.mockMessageId, SOURCE_CHAIN_SELECTOR, params.amount - expectedFastTransferFee, 18, params.receiver
    );

    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, params.chainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: params.chainSelector,
      fillId: fillId,
      settlementId: params.mockMessageId,
      sourceAmountNetFee: params.amount - expectedFastTransferFee,
      sourceDecimals: SOURCE_DECIMALS,
      fillerFee: expectedFillerFee,
      poolFee: expectedPoolFee,
      destinationPool: destPoolAddress,
      receiver: params.receiver
    });

    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes32 settlementId = s_pool.ccipSendToken{value: 1 ether}(
      params.chainSelector, params.amount, expectedFastTransferFee, params.receiver, address(0), ""
    );

    assertEq(settlementId, params.mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - params.amount);
    assertEq(s_token.balanceOf(address(s_pool)), params.amount);
  }

  function test_ccipSendToken_NativeFee() public {
    TestParams memory params = _setupTestParams(DEST_CHAIN_SELECTOR);

    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );

    _executeTest(params, extraArgs);
  }

  function test_ccipSendToken_NativeFee_ToSVM() public {
    TestParams memory params = _setupTestParams(SVM_CHAIN_SELECTOR);
    _executeTest(params, s_svmExtraArgsBytesEncoded);
  }

  function test_ccipSendToken_APTOS_WithSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_WITH_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_APTOS, SETTLEMENT_GAS_OVERHEAD, "");

    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );

    _executeTest(params, extraArgs);
  }

  function test_ccipSendToken_APTOS_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_ZERO_GAS")));
    TestParams memory params = _setupTestParams(testChainSelector);

    bytes memory customExtraArgs = abi.encode("aptosCustomExtraArgs");
    _setupChainConfig(testChainSelector, Internal.CHAIN_FAMILY_SELECTOR_APTOS, 0, customExtraArgs);
    _executeTest(params, customExtraArgs);
  }

  function test_ccipSendToken_WithERC20FeeToken() public {
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

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: DEST_CHAIN_SELECTOR,
      fillId: s_pool.computeFillId(
        fakeMessageId, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT - expectedFastTransferFee, 18, abi.encode(RECEIVER)
      ), //expected fill id
      settlementId: fakeMessageId,
      sourceAmountNetFee: SOURCE_AMOUNT - expectedFastTransferFee, // expected amount net fee
      sourceDecimals: SOURCE_DECIMALS,
      fillerFee: expectedFastTransferFee,
      poolFee: 0,
      destinationPool: destPoolAddress,
      receiver: abi.encode(RECEIVER)
    });

    bytes32 settlementId = s_pool.ccipSendToken(
      DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, expectedFastTransferFee, abi.encode(RECEIVER), feeToken, ""
    );

    assertTrue(settlementId != bytes32(0));
    assertEq(s_token.balanceOf(OWNER), balanceBefore - SOURCE_AMOUNT - quote.ccipSettlementFee);
    // When using ERC20 fee token, both transfer amount and settlement fee go to the pool
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore + SOURCE_AMOUNT + quote.ccipSettlementFee);
  }

  function test_ccipSendToken_WithPoolFee() public {
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 25; // 0.25%

    TestParams memory params = _setupTestParams(DEST_CHAIN_SELECTOR);
    params.fastFeeBpsExpected = fillerFeeBps + poolFeeBps; // Total fee

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    _setupMocks(params.mockMessageId);

    uint256 fillerFeeAmount = (SOURCE_AMOUNT * fillerFeeBps) / 10_000;
    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 totalFastTransferFee = fillerFeeAmount + poolFeeAmount;
    uint256 amountNetTotalFee = SOURCE_AMOUNT - totalFastTransferFee;

    bytes32 fillId =
      s_pool.computeFillId(params.mockMessageId, SOURCE_CHAIN_SELECTOR, amountNetTotalFee, 18, params.receiver);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: params.chainSelector,
      fillId: fillId,
      settlementId: params.mockMessageId,
      sourceAmountNetFee: amountNetTotalFee,
      sourceDecimals: SOURCE_DECIMALS,
      fillerFee: fillerFeeAmount,
      poolFee: poolFeeAmount,
      destinationPool: destPoolAddress,
      receiver: params.receiver
    });

    uint256 balanceBefore = s_token.balanceOf(OWNER);

    bytes32 settlementId = s_pool.ccipSendToken{value: 1 ether}(
      params.chainSelector, params.amount, totalFastTransferFee, params.receiver, address(0), ""
    );

    assertEq(settlementId, params.mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - params.amount);
    assertEq(s_token.balanceOf(address(s_pool)), params.amount);
  }

  function test_ccipSendToken_FeeQuote_WithPoolFee() public {
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

  function test_ccipSendToken_EqualFeeSplit() public {
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

  function test_ccipSendToken_FeeValidation_Success_WhenFeeWithinLimit() public {
    // Setup: Calculate expected fee and set max limit higher
    uint256 expectedFee = (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000; // 1% of 100 ether = 1 ether
    uint256 maxFeeLimit = expectedFee + 0.5 ether; // Set limit higher than expected fee

    // Setup mocks
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    bytes memory receiver = abi.encodePacked(address(0x5));
    uint256 balanceBefore = s_token.balanceOf(OWNER);

    // Should succeed - fee is within limit
    bytes32 settlementId =
      s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, maxFeeLimit, receiver, address(0), "");

    // Verify transaction completed successfully
    assertEq(settlementId, mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - SOURCE_AMOUNT);
  }

  function test_ccipSendToken_FeeValidation_Success_WhenFeeEqualsLimit() public {
    // Setup: Calculate expected fee and set max limit equal to it
    uint256 expectedFee = (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000; // 1% of 100 ether = 1 ether
    uint256 maxFeeLimit = expectedFee; // Set limit exactly equal to expected fee

    // Setup mocks
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    bytes memory receiver = abi.encodePacked(address(0x5));
    uint256 balanceBefore = s_token.balanceOf(OWNER);

    // Should succeed - fee equals limit (boundary condition)
    bytes32 settlementId =
      s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, maxFeeLimit, receiver, address(0), "");

    // Verify transaction completed successfully
    assertEq(settlementId, mockMessageId);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - SOURCE_AMOUNT);
  }

  function test_ccipSendToken_RevertWhen_ReceiverIsEmptyBytes() public {
    // Setup: Empty receiver address
    bytes memory emptyReceiver = "";

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, emptyReceiver));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, emptyReceiver, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_ReceiverExceedsMaxLength() public {
    // Setup: Receiver longer than 64 bytes (65 bytes)
    bytes memory oversizedReceiver = new bytes(65);
    // Fill with non-zero data to ensure it's not rejected for being all zeros
    for (uint256 i = 0; i < 65; i++) {
      oversizedReceiver[i] = bytes1(uint8(i + 1));
    }

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, oversizedReceiver));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, oversizedReceiver, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_ReceiverIsAllZeros() public {
    // Setup: 20-byte receiver address filled with zeros (typical Ethereum address length)
    bytes memory zeroReceiver20 = new bytes(20);
    // bytes constructor already initializes to zeros

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, zeroReceiver20));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, zeroReceiver20, address(0), "");

    // Setup: 32-byte receiver address filled with zeros (one full word)
    bytes memory zeroReceiver32 = new bytes(32);

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, zeroReceiver32));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, zeroReceiver32, address(0), "");

    // Setup: 40-byte receiver address filled with zeros
    bytes memory zeroReceiver40 = new bytes(40);

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, zeroReceiver40));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, zeroReceiver40, address(0), "");

    // Setup: 64-byte receiver address filled with zeros (maximum allowed length)
    bytes memory zeroReceiver64 = new bytes(64);

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidReceiver.selector, zeroReceiver64));
    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, zeroReceiver64, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_FeeExceedsUserMaxLimit() public {
    // Setup: Calculate expected fee and set max limit lower
    uint256 expectedFee = (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000; // 1% of 100 ether = 1 ether
    uint256 maxFeeLimit = expectedFee - 0.1 ether; // Set limit lower than expected fee

    // Setup mocks (needed for fee calculation)
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    bytes memory receiver = abi.encodePacked(address(0x5));

    // Should revert with QuoteFeeExceedsUserMaxLimit
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.QuoteFeeExceedsUserMaxLimit.selector, expectedFee, maxFeeLimit
      )
    );

    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, maxFeeLimit, receiver, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_FeeExceedsUserMaxLimit_WithPoolFee() public {
    // Setup: Configure both filler and pool fees
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 50; // 0.5%
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    // Calculate expected total fee
    uint256 expectedTotalFee = (SOURCE_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000; // 1.25% of 100 ether = 1.25 ether
    uint256 maxFeeLimit = expectedTotalFee - 0.05 ether; // Set limit slightly lower than expected fee

    // Setup mocks
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    bytes memory receiver = abi.encodePacked(address(0x5));

    // Should revert with QuoteFeeExceedsUserMaxLimit
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.QuoteFeeExceedsUserMaxLimit.selector, expectedTotalFee, maxFeeLimit
      )
    );

    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, maxFeeLimit, receiver, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_FeeIncreasedAfterQuote_FrontRun() public {
    // Setup: Start with low fee configuration
    uint16 initialFillerFeeBps = 50; // 0.5%
    uint16 initialPoolFeeBps = 25; // 0.25%
    _updateConfigWithPoolFee(initialFillerFeeBps, initialPoolFeeBps);

    // Setup mocks for fee calculation
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));

    bytes memory receiver = abi.encodePacked(address(0x5));

    // Step 1: User queries the fee with initial low configuration
    IFastTransferPool.Quote memory initialQuote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, receiver, address(0), "");

    uint256 expectedInitialFee = (SOURCE_AMOUNT * (initialFillerFeeBps + initialPoolFeeBps)) / 10_000; // 0.75% of 100 ether = 0.75 ether
    assertEq(initialQuote.fastTransferFee, expectedInitialFee, "Initial fee calculation incorrect");

    // User decides to use the queried fee as their maximum acceptable fee
    uint256 userMaxFeeLimit = initialQuote.fastTransferFee;

    // Step 2: Configuration gets updated to higher fees (front-running scenario)
    uint16 newFillerFeeBps = 150; // 1.5% (3x increase)
    uint16 newPoolFeeBps = 100; // 1.0% (4x increase)
    _updateConfigWithPoolFee(newFillerFeeBps, newPoolFeeBps);

    // Calculate what the new fee would be after the configuration update
    uint256 expectedNewFee = (SOURCE_AMOUNT * (newFillerFeeBps + newPoolFeeBps)) / 10_000; // 2.5% of 100 ether = 2.5 ether

    // Verify the new fee is indeed higher than user's limit
    assertGt(expectedNewFee, userMaxFeeLimit, "New fee should be higher than user's limit");

    // Setup mocks for the actual send
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    // Step 3: User calls ccipSendToken with their originally queried fee as max limit
    // This should now revert because the actual fee (with updated config) exceeds their limit
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.QuoteFeeExceedsUserMaxLimit.selector, expectedNewFee, userMaxFeeLimit
      )
    );

    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, userMaxFeeLimit, receiver, address(0), "");
  }

  function test_ccipSendToken_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.ccipSendToken{value: 1 ether}(
      DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, 1 ether, abi.encode(RECEIVER), address(0), ""
    );
  }

  function test_ccipSendToken_RevertWhen_ConfigUpdatedBetweenQuoteAndSend_EdgeCase() public {
    // Setup: Start with zero pool fee, only filler fee
    uint16 initialFillerFeeBps = 100; // 1%
    uint16 initialPoolFeeBps = 0; // 0%
    _updateConfigWithPoolFee(initialFillerFeeBps, initialPoolFeeBps);

    // Setup mocks for fee calculation
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));

    bytes memory receiver = abi.encodePacked(address(0x5));

    // Step 1: User queries fee with no pool fee
    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, receiver, address(0), "");

    uint256 expectedInitialFee = (SOURCE_AMOUNT * initialFillerFeeBps) / 10_000; // 1% of 100 ether = 1 ether
    assertEq(quote.fastTransferFee, expectedInitialFee, "Initial fee should only include filler fee");

    // User sets their max fee limit to exactly the queried amount
    uint256 userMaxFeeLimit = quote.fastTransferFee;

    // Step 2: Pool fee gets introduced (admin adds pool fee)
    uint16 newFillerFeeBps = 100; // Keep same filler fee
    uint16 newPoolFeeBps = 50; // Add 0.5% pool fee
    _updateConfigWithPoolFee(newFillerFeeBps, newPoolFeeBps);

    // Now total fee includes both filler + pool fee
    uint256 expectedNewTotalFee = (SOURCE_AMOUNT * (newFillerFeeBps + newPoolFeeBps)) / 10_000; // 1.5% of 100 ether = 1.5 ether

    // Verify the new fee exceeds user's limit
    assertGt(expectedNewTotalFee, userMaxFeeLimit, "New total fee should exceed user's original limit");

    // Setup mocks for the actual send
    bytes32 mockMessageId = keccak256("mockMessageId");
    _setupMocks(mockMessageId);

    // Step 3: User's transaction should fail due to newly introduced pool fee
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.QuoteFeeExceedsUserMaxLimit.selector, expectedNewTotalFee, userMaxFeeLimit
      )
    );

    s_pool.ccipSendToken{value: 1 ether}(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, userMaxFeeLimit, receiver, address(0), "");
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

  // This test achieves higher input space coverage than setting 1 non-zero byte at random index
  function testFuzz_ccipSendToken_ValidReceiver(
    uint8 receiverLength,
    bytes32 receiverHead,
    bytes32 receiverTail
  ) public {
    receiverLength = uint8(bound(receiverLength, 1, 64));
    // Combine the 2 halves into bytes of length 64
    bytes memory validReceiver = abi.encodePacked(receiverHead, receiverTail);
    // Set bytes array length to target receiver length
    assembly {
      mstore(validReceiver, receiverLength)
    }

    // Throw out all-zero receiver
    vm.assume(keccak256(validReceiver) != keccak256(new bytes(receiverLength)));

    TestParams memory params = TestParams({
      chainSelector: DEST_CHAIN_SELECTOR,
      fastFeeBpsExpected: FAST_FEE_FILLER_BPS,
      amount: SOURCE_AMOUNT,
      mockMessageId: keccak256("mockMessageId"),
      receiver: validReceiver
    });
    bytes memory extraArgs = Client._argsToBytes(
      Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
    );
    _executeTest(params, extraArgs);
  }
}
