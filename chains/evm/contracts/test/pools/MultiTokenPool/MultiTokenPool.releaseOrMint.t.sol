// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPoolHooks} from "../../../interfaces/IMultiTokenPoolHooks.sol";
import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_releaseOrMint is MultiTokenPoolSetup {
  function test_releaseOrMint() public {
    uint256 sourceAmount = 20_000e18;
    // The pool config manager converts the amount to local decimals.
    uint256 amount = 20_000e6;
    bytes4 finality = FinalityCodec._encodeBlockDepth(5);
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _releaseOrMintIn(sourceAmount);
    // The mock matches only the input and the finality that the pool must forward.
    _mockValidateReleaseOrMint(releaseOrMintIn, finality, _defaultInboundTransfer(amount));

    vm.startPrank(s_allowedOffRamp);
    vm.expectCall(
      address(s_token), abi.encodeWithSelector(bytes4(keccak256("mint(address,uint256)")), s_receiver, amount)
    );
    vm.expectEmit();
    emit MultiTokenPool.ReleasedOrMinted(DEST_CHAIN_SELECTOR, address(s_token), s_allowedOffRamp, s_receiver, amount);

    Pool.ReleaseOrMintOutV1 memory releaseOrMintOut = s_pool.releaseOrMint(releaseOrMintIn, finality);

    assertEq(releaseOrMintOut.destinationAmount, amount);
    assertEq(s_token.balanceOf(s_receiver), amount);
  }

  function test_releaseOrMint_LegacyPathUsesDefaultFinality() public {
    uint256 amount = 20_000e18;
    // The legacy path forwards default finality, so only this mock matches.
    _mockValidateReleaseOrMint(
      _releaseOrMintIn(amount), FinalityCodec.WAIT_FOR_FINALITY_FLAG, _defaultInboundTransfer(amount)
    );

    vm.startPrank(s_allowedOffRamp);
    s_pool.releaseOrMint(_releaseOrMintIn(amount));

    assertEq(s_token.balanceOf(s_receiver), amount);
  }

  function test_releaseOrMint_CallsPostflightHook() public {
    uint256 sourceAmount = 20_000e18;
    uint256 amount = 20_000e6;
    bytes4 finality = FinalityCodec._encodeBlockDepth(5);
    address hook = makeAddr("hook");
    vm.etch(hook, bytes("fake bytecode"));
    IPoolConfigManager.ResolvedInboundTransfer memory transfer = _defaultInboundTransfer(amount);
    transfer.hook = hook;
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _releaseOrMintIn(sourceAmount);
    _mockValidateReleaseOrMint(releaseOrMintIn, finality, transfer);

    vm.startPrank(s_allowedOffRamp);
    // The hook gets the amount in local decimals and the requested finality.
    vm.expectCall(hook, abi.encodeCall(IMultiTokenPoolHooks.postflightCheck, (releaseOrMintIn, amount, finality)));
    s_pool.releaseOrMint(releaseOrMintIn, finality);
  }

  // Reverts

  function test_releaseOrMint_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.startPrank(s_allowedOffRamp);
    vm.expectRevert(MultiTokenPool.CursedByRMN.selector);
    s_pool.releaseOrMint(_releaseOrMintIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_releaseOrMint_RevertWhen_CallerIsNotARampOnRouter() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPool.CallerIsNotARampOnRouter.selector, STRANGER));
    s_pool.releaseOrMint(_releaseOrMintIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_releaseOrMint_RevertWhen_UnsupportedLockOrBurnFlavor() public {
    IPoolConfigManager.ResolvedInboundTransfer memory transfer = _defaultInboundTransfer(1e18);
    transfer.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Lock;
    transfer.lockOrBurnAddress = makeAddr("lockBox");
    _mockValidateReleaseOrMint(_releaseOrMintIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG, transfer);

    vm.startPrank(s_allowedOffRamp);
    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPool.UnsupportedLockOrBurnFlavor.selector, address(s_token), IPoolConfigManager.LockOrBurnFlavor.Lock
      )
    );
    s_pool.releaseOrMint(_releaseOrMintIn(1e18), FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
