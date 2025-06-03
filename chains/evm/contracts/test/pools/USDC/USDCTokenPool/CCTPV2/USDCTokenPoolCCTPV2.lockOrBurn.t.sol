// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {Router} from "../../../../../Router.sol";
import {Pool} from "../../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../../pools/TokenPool.sol";

import {USDCTokenPool} from "../../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_lockOrBurn is USDCTokenPoolCCTPV2Setup {
  // Base test case, included for PR gas comparisons as fuzz tests are excluded from forge snapshot due to being flaky.
  function test_LockOrBurn() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;
    s_token.transfer(address(s_usdcTokenPool), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      receiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
  }

  function testFuzz_LockOrBurn_Success(bytes32 destinationReceiver, uint256 amount) public {
    vm.assume(destinationReceiver != bytes32(0));
    amount = bound(amount, 1, _getOutboundRateLimiterConfig().capacity);
    s_token.transfer(address(s_usdcTokenPool), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      destinationReceiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(destinationReceiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
    assertEq(poolReturnDataV1.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
  }

  function testFuzz_LockOrBurnWithAllowList_Success(bytes32 destinationReceiver, uint256 amount) public {
    vm.assume(destinationReceiver != bytes32(0));
    amount = bound(amount, 1, _getOutboundRateLimiterConfig().capacity);
    s_token.transfer(address(s_usdcTokenPoolWithAllowList), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPoolWithAllowList.getDomain(DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_token),
      amount,
      address(s_usdcTokenPoolWithAllowList),
      destinationReceiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller,
      s_usdcTokenPoolWithAllowList.MAX_FEE(),
      s_usdcTokenPoolWithAllowList.FINALITY_THRESHOLD(),
      ""
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = s_usdcTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: s_allowedList[0],
        receiver: abi.encodePacked(destinationReceiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
    uint64 nonce = abi.decode(poolReturnDataV1.destPoolData, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
    assertEq(poolReturnDataV1.destTokenAddress, abi.encode(DEST_CHAIN_USDC_TOKEN));
  }

  // Reverts
  function test_RevertWhen_UnknownDomain() public {
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
    deal(address(s_token), s_routerAllowedOnRamp, amount);
    s_token.approve(address(s_usdcTokenPool), amount);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.UnknownDomain.selector, wrongDomain));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(address(0)),
        amount: amount,
        remoteChainSelector: wrongDomain,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_CallerIsNotARampOnRouter() public {
    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotARampOnRouter.selector, OWNER));

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(address(0)),
        amount: 0,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_LockOrBurnWithAllowList() public {
    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.SenderNotAllowed.selector, STRANGER));

    s_usdcTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: abi.encodePacked(address(0)),
        amount: 1000,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_InvalidReceiver() public {
    vm.startPrank(s_routerAllowedOnRamp);

    // Generate a 33-byte long string to use as invalid receiver
    bytes memory invalidReceiver = abi.encode(keccak256("0xCLL"), "A");

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidReceiver.selector, invalidReceiver));
    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: invalidReceiver,
        amount: 1000,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }
}
