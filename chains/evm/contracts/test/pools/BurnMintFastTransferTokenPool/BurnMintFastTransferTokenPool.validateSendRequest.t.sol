// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {BurnMintFastTransferTokenPool} from "../../../pools/BurnMintFastTransferTokenPool.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_validateSendRequest is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant CCIP_SEND_FEE = 1 ether;
  bytes32 internal constant MESSAGE_ID = keccak256("messageId");

  function setUp() public virtual override {
    super.setUp();
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(CCIP_SEND_FEE)
    );
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(MESSAGE_ID)
    );
    deal(address(s_token), OWNER, TRANSFER_AMOUNT * 10);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_validateSendRequest_Success() public view {
    // From the setup, we have a pool with the default chain config and no allowlist.
    // This should not revert - all validations pass
    s_pool.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed status for destination chain
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));
    // Should revert with CursedByRMN error
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_RevertWhen_ChainNotAllowed() public {
    // Unregistered chain selector
    uint64 unsupportedChainSelector = 999999;
    // Should revert with ChainNotAllowed error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChainSelector));
    s_pool.getCcipSendTokenFee(
      unsupportedChainSelector,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_RevertWhen_TransferAmountExceedsMaxFillAmount() public {
    // Amount exceeding the max fill amount per request (FILL_AMOUNT_MAX = 1000 ether)
    uint256 excessiveAmount = FILL_AMOUNT_MAX + 1;

    // Should revert with TransferAmountExceedsMaxFillAmount error
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.TransferAmountExceedsMaxFillAmount.selector, DEST_CHAIN_SELECTOR, excessiveAmount
      )
    );
    s_pool.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      excessiveAmount,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_WithAllowlistedSender() public {
    address allowlistedSender = makeAddr("allowlistedSender");

    // Create a pool with allowlist enabled
    address[] memory allowlist = new address[](1);
    allowlist[0] = allowlistedSender;

    BurnMintFastTransferTokenPool poolWithAllowlist = new BurnMintFastTransferTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      allowlist,
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      SOURCE_CHAIN_SELECTOR
    );
    s_token.grantMintAndBurnRoles(address(poolWithAllowlist));
    // Setup chain and lane config for the new pool
    _setupPoolConfiguration(poolWithAllowlist);

    deal(address(s_token), allowlistedSender, TRANSFER_AMOUNT * 2);

    vm.stopPrank();
    vm.startPrank(allowlistedSender);
    s_token.approve(address(poolWithAllowlist), type(uint256).max);
    deal(allowlistedSender, CCIP_SEND_FEE); // Ensure sender has enough ETH for the fee
    // Should succeed with allowlisted sender
    poolWithAllowlist.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_RevertWhen_SenderNotAllowlisted() public {
    vm.stopPrank();
    // Create a pool with allowlist enabled
    address[] memory allowlist = new address[](1);
    allowlist[0] = OWNER; // Only OWNER is allowed

    BurnMintFastTransferTokenPool poolWithAllowlist = new BurnMintFastTransferTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      allowlist,
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      SOURCE_CHAIN_SELECTOR
    );

    // Setup chain and lane config for the new pool
    _setupPoolConfiguration(poolWithAllowlist);

    address unauthorizedSender = makeAddr("unauthorizedSender");
    vm.prank(unauthorizedSender);
    // Should revert with SenderNotAllowed error (from _checkAllowList)
    vm.expectRevert(abi.encodeWithSelector(TokenPool.SenderNotAllowed.selector, unauthorizedSender));
    poolWithAllowlist.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function test_validateSendRequest_RevertWhen_RMNCurseSpecificChain() public {
    uint64 anotherChainSelector = 54321;

    // Add configuration for another chain
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_remoteBurnMintPool);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: anotherChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(3)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_pool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Add lane config for the new chain
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: anotherChainSelector,
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    // Mock RMN to return cursed status only for the new chain
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(anotherChainSelector))),
      abi.encode(true)
    );

    // Original chain should still work
    s_pool.getCcipSendTokenFee(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );

    // New chain should revert due to curse
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.getCcipSendTokenFee(
      anotherChainSelector,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );
  }

  function _setupPoolConfiguration(
    BurnMintFastTransferTokenPool pool
  ) internal {
    // Setup chain configuration
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_remoteBurnMintPool);

    TokenPool.ChainUpdate[] memory chainsToAdd = new TokenPool.ChainUpdate[](1);
    chainsToAdd[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_remoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    pool.applyChainUpdates(new uint64[](0), chainsToAdd);

    // Setup lane configuration
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }
}
