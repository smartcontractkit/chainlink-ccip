// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintTokenPoolV2_releaseOrMint is BurnMintSetup {
  function test_releaseOrMint() public {
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
