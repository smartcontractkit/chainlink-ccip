// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract USDCTokenPoolCCTPV2_lockOrBurn is USDCTokenPoolCCTPV2Setup {
  // Base test case, included for PR gas comparisons as fuzz tests are excluded from forge snapshot due to being flaky.
  function test_lockOrBurn() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;
    s_USDCToken.transfer(address(s_usdcTokenPool), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      amount: amount
    });

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      amount,
      address(s_usdcTokenPool),
      receiver,
      expectedDomain.domainIdentifier,
      s_mockUSDCTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenDataPayload =
      abi.decode(poolReturnDataV1.destPoolData, (USDCTokenPool.SourceTokenDataPayload));
    assertEq(sourceTokenDataPayload.nonce, 0, "nonce is incorrect");
    assertEq(sourceTokenDataPayload.sourceDomain, DEST_DOMAIN_IDENTIFIER, "sourceDomain is incorrect");
    assertEq(
      uint8(sourceTokenDataPayload.cctpVersion), uint8(USDCTokenPool.CCTPVersion.CCTP_V2), "cctpVersion is incorrect"
    );
    assertEq(sourceTokenDataPayload.amount, amount, "amount is incorrect");
    assertEq(
      sourceTokenDataPayload.destinationDomain, expectedDomain.domainIdentifier, "destinationDomain is incorrect"
    );
    assertEq(sourceTokenDataPayload.mintRecipient, receiver, "mintRecipient is incorrect");
    assertEq(sourceTokenDataPayload.burnToken, address(s_USDCToken), "burnToken is incorrect");
    assertEq(sourceTokenDataPayload.destinationCaller, expectedDomain.allowedCaller, "destinationCaller is incorrect");
    assertEq(sourceTokenDataPayload.maxFee, s_usdcTokenPool.MAX_FEE(), "maxFee is incorrect");
    assertEq(
      sourceTokenDataPayload.minFinalityThreshold,
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      "minFinalityThreshold is incorrect"
    );
  }

  function test_lockOrBurn_MintRecipientOverride() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;
    s_USDCToken.transfer(address(s_usdcTokenPool), amount);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    // Set up a domain override with a custom mintRecipient
    bytes32 extraMintRecipient = bytes32("random_mint_recipient_123");
    USDCTokenPool.DomainUpdate[] memory updates = new USDCTokenPool.DomainUpdate[](1);
    updates[0] = USDCTokenPool.DomainUpdate({
      allowedCaller: expectedDomain.allowedCaller,
      mintRecipient: extraMintRecipient,
      domainIdentifier: expectedDomain.domainIdentifier,
      destChainSelector: DEST_CHAIN_SELECTOR,
      enabled: expectedDomain.enabled
    });
    vm.startPrank(OWNER);
    s_usdcTokenPool.setDomains(updates);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      amount: amount
    });

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      amount,
      address(s_usdcTokenPool),
      extraMintRecipient,
      expectedDomain.domainIdentifier,
      s_mockUSDCTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(nonce, 0);
  }

  function testFuzz_lockOrBurn_Success(bytes32 destinationReceiver, uint256 amount) public {
    vm.assume(destinationReceiver != bytes32(0));
    amount = bound(amount, 1, _getOutboundRateLimiterConfig().capacity);
    s_USDCToken.transfer(address(s_usdcTokenPool), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      amount: amount
    });

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      amount,
      address(s_usdcTokenPool),
      destinationReceiver,
      expectedDomain.domainIdentifier,
      s_mockUSDCTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(destinationReceiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(nonce, 0);
    assertEq(poolReturnDataV1.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
  }

  // Reverts
  function test_lockOrBurn_RevertWhen_UnknownDomain() public {
    uint64 wrongDomain = DEST_CHAIN_SELECTOR + 1;

    // We need to setup the wrong chainSelector so it reaches the domain check
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: wrongDomain, onRamp: s_routerAllowedOnRamp});
    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: wrongDomain,
      remotePoolAddresses: new bytes[](0),
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_usdcTokenPool.applyChainUpdates(new uint64[](0), chainUpdates);

    uint256 amount = 1000;
    vm.startPrank(s_routerAllowedOnRamp);
    deal(address(s_USDCToken), s_routerAllowedOnRamp, amount);
    s_USDCToken.approve(address(s_usdcTokenPool), amount);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.UnknownDomain.selector, wrongDomain));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(address(0)),
        amount: amount,
        remoteChainSelector: wrongDomain,
        localToken: address(s_USDCToken)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_CallerIsNotARampOnRouter() public {
    address randomAddress = makeAddr("RANDOM_ADDRESS");

    vm.startPrank(randomAddress);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, randomAddress));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(address(0)),
        amount: 0,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_InvalidReceiver() public {
    vm.startPrank(s_routerAllowedOnRamp);
    uint256 amount = 1000;

    // Generate a byte string that is not 32 bytes long
    bytes memory invalidReceiver = abi.encodePacked(keccak256("invalid_receiver"), keccak256("invalid_receiver"));

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidReceiver.selector, invalidReceiver));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: invalidReceiver,
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );
  }
}
