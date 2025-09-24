// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {LOCK_RELEASE_FLAG} from "../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {BurnMintWithLockReleaseFlagTokenPoolSetup} from "./BurnMintWithLockReleaseFlagTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintWithLockReleaseFlagTokenPool_releaseOrMint is BurnMintWithLockReleaseFlagTokenPoolSetup {
  function test_releaseOrMint_LockReleaseFlagInSourcePoolData() public {
    uint256 amount = 1e19;
    address receiver = makeAddr("receiver_address");

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(0), receiver, amount);

    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: receiver,
        sourceDenominatedAmount: amount,
        localToken: address(s_token),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: abi.encode(LOCK_RELEASE_FLAG),
        offchainTokenData: ""
      })
    );

    assertEq(s_token.balanceOf(receiver), amount);
  }

  function test_releaseOrMint_EmptySourcePoolData() public {
    uint256 amount = 1e19;
    address receiver = makeAddr("receiver_address");

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(0), receiver, amount);

    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: receiver,
        sourceDenominatedAmount: amount,
        localToken: address(s_token),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(s_token.balanceOf(receiver), amount);
  }
}

// TODO: Full E2E tests
