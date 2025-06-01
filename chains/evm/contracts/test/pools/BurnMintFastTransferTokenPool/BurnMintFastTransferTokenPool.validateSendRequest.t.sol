// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {BurnMintFastTransferTokenPool} from "../../../pools/BurnMintFastTransferTokenPool.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_validateSendRequest is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant TRANSFER_AMOUNT = 100 ether;
  address internal constant RECEIVER = address(0x1234);
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

  function test_ValidateSendRequest_Success() public {
    // This should not revert - all validations pass
    bytes32 fillRequestId = s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), // native fee token
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      ""
    );

    assertTrue(fillRequestId != bytes32(0));
  }

  function test_ValidateSendRequest_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed status for destination chain
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    // Should revert with CursedByRMN error
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
    );
  }

  function test_ValidateSendRequest_RevertWhen_ChainNotAllowed() public {
    uint64 unsupportedChainSelector = 999999;

    // Should revert with ChainNotAllowed error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChainSelector));
    s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), unsupportedChainSelector, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
    );
  }

  function test_ValidateSendRequest_RevertWhen_SenderNotAllowlisted() public {
    vm.stopPrank();
    // Create a pool with allowlist enabled
    address[] memory allowlist = new address[](1);
    allowlist[0] = OWNER; // Only OWNER is allowed

    BurnMintFastTransferTokenPool poolWithAllowlist = new BurnMintFastTransferTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, allowlist, address(s_mockRMNRemote), address(s_sourceRouter)
    );

    // Setup chain and lane config for the new pool
    _setupPoolConfiguration(poolWithAllowlist);

    address unauthorizedSender = makeAddr("unauthorizedSender");
    deal(address(s_token), unauthorizedSender, TRANSFER_AMOUNT * 2);
    deal(unauthorizedSender, CCIP_SEND_FEE); // Ensure sender has enough ETH for the fee
    vm.prank(unauthorizedSender);
    s_token.approve(address(poolWithAllowlist), type(uint256).max);

    // Should revert with SenderNotAllowed error (from _checkAllowList)
    vm.expectRevert(abi.encodeWithSelector(TokenPool.SenderNotAllowed.selector, unauthorizedSender));
    vm.prank(unauthorizedSender);
    poolWithAllowlist.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
    );
  }

  function test_ValidateSendRequest_WithAllowlistedSender() public {
    vm.stopPrank();
    address allowlistedSender = makeAddr("allowlistedSender");

    // Create a pool with allowlist enabled
    address[] memory allowlist = new address[](1);
    allowlist[0] = allowlistedSender;

    BurnMintFastTransferTokenPool poolWithAllowlist = new BurnMintFastTransferTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, allowlist, address(s_mockRMNRemote), address(s_sourceRouter)
    );
    vm.prank(OWNER);
    s_token.grantMintAndBurnRoles(address(poolWithAllowlist));
    // Setup chain and lane config for the new pool
    _setupPoolConfiguration(poolWithAllowlist);

    deal(address(s_token), allowlistedSender, TRANSFER_AMOUNT * 2);
    vm.prank(allowlistedSender);
    s_token.approve(address(poolWithAllowlist), type(uint256).max);
    deal(allowlistedSender, CCIP_SEND_FEE); // Ensure sender has enough ETH for the fee
    // Should succeed with allowlisted sender
    vm.prank(allowlistedSender);
    bytes32 fillRequestId = poolWithAllowlist.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
    );

    assertTrue(fillRequestId != bytes32(0));
  }

  function test_ValidateSendRequest_RMNCurseSpecificChain() public {
    vm.stopPrank();
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

    vm.prank(OWNER);
    s_pool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Add lane config for the new chain
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: anotherChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    vm.prank(OWNER);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    // Mock RMN to return cursed status only for the new chain
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(anotherChainSelector))),
      abi.encode(true)
    );

    // Original chain should still work
    vm.prank(OWNER);
    bytes32 fillRequestId1 = s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
    );
    assertTrue(fillRequestId1 != bytes32(0));

    // New chain should revert due to curse
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    vm.prank(OWNER);
    s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      address(0), anotherChainSelector, TRANSFER_AMOUNT, abi.encode(RECEIVER), ""
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
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
    pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, addFillers, new address[](0));
  }
}
