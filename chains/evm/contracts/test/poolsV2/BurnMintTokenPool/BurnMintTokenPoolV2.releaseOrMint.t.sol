// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool as TokenPoolV1} from "../../../pools/TokenPool.sol";
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

  function test_RevertWhen_CursedByRMN() public {
    // Should not mint tokens if cursed.
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));
    uint256 before = s_token.balanceOf(OWNER);
    vm.startPrank(s_allowedOffRamp);

    vm.expectRevert(TokenPoolV1.CursedByRMN.selector);
    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: 1e5,
        localToken: address(s_token),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: _generateSourceTokenData().sourcePoolAddress,
        sourcePoolData: _generateSourceTokenData().extraData,
        offchainTokenData: ""
      })
    );

    assertEq(s_token.balanceOf(OWNER), before);
  }

  function test_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = 8838833;

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV1.ChainNotAllowed.selector, wrongChainSelector));
    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: 1,
        localToken: address(s_token),
        remoteChainSelector: wrongChainSelector,
        sourcePoolAddress: _generateSourceTokenData().sourcePoolAddress,
        sourcePoolData: _generateSourceTokenData().extraData,
        offchainTokenData: ""
      })
    );
  }
}
