// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_releaseOrMint is SiloedLockReleaseTokenPoolSetup {
  function setUp() public override {
    super.setUp();

    s_token.approve(address(s_lockBox), type(uint256).max);
    IERC20(address(s_token)).approve(address(s_siloLockBox), type(uint256).max);
  }

  function test_ReleaseOrMint_SiloedChain() public {
    uint256 amount = 10e18;

    deal(address(s_token), address(s_siloedLockReleaseTokenPool), amount);

    vm.startPrank(s_allowedOnRamp);

    // Lock funds so that they can be released without underflowing the internal accounting.
    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_siloLockBox)), amount);

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_siloLockBox), OWNER, amount);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: amount,
        localToken: address(s_token),
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(s_token.balanceOf(address(s_siloLockBox)), 0);
  }

  function test_ReleaseOrMint_UnsiloedChain() public {
    uint256 amount = 10e18;

    deal(address(s_token), address(s_siloedLockReleaseTokenPool), amount);
    vm.startPrank(s_allowedOnRamp);

    // Lock funds so they can be released later.
    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_lockBox)), amount);

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_lockBox), OWNER, amount);

    s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: amount,
        localToken: address(s_token),
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(s_token.balanceOf(address(s_lockBox)), 0);
  }

  function test_ReleaseOrMintV2_SiloedChain() public {
    uint256 amount = 10e18;
    address recipient = makeAddr("recipient");

    deal(address(s_token), address(s_siloedLockReleaseTokenPool), amount);

    vm.startPrank(s_allowedOnRamp);

    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    vm.startPrank(s_allowedOffRamp);

    Pool.ReleaseOrMintOutV1 memory output = s_siloedLockReleaseTokenPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: recipient,
        sourceDenominatedAmount: amount,
        localToken: address(s_token),
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_siloedDestPoolAddress),
        sourcePoolData: "",
        offchainTokenData: ""
      }),
      0
    );

    assertEq(output.destinationAmount, amount);
    assertEq(s_token.balanceOf(recipient), amount);
    assertEq(s_token.balanceOf(address(s_siloLockBox)), 0);
  }
}
