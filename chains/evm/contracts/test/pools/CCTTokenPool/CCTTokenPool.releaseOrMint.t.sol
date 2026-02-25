// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CCTTokenPoolSetup} from "./CCTTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool_releaseOrMint is CCTTokenPoolSetup {
  function test_releaseOrMint() public {
    uint256 amount = 1e18;
    address receiver = makeAddr("receiver_address");

    vm.startPrank(s_allowedOffRamp);

    vm.expectEmit();
    emit IERC20.Transfer(address(0), receiver, amount);

    s_cctPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: receiver,
        sourceDenominatedAmount: amount,
        localToken: address(s_cctPool),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );

    assertEq(IERC20(address(s_cctPool)).balanceOf(receiver), amount);
  }

  function test_releaseOrMint_RevertWhen_MaxSupplyExceeded() public {
    // Try to mint more than the remaining supply
    uint256 remaining = MAX_SUPPLY - IERC20(address(s_cctPool)).totalSupply();
    uint256 tooMuch = remaining + 1;

    vm.startPrank(s_allowedOffRamp);

    vm.expectRevert(
      abi.encodeWithSelector(BaseERC20.MaxSupplyExceeded.selector, IERC20(address(s_cctPool)).totalSupply() + tooMuch)
    );
    s_cctPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: tooMuch,
        localToken: address(s_cctPool),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );
  }

  function test_releaseOrMint_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.startPrank(s_allowedOffRamp);
    vm.expectRevert(TokenPool.CursedByRMN.selector);

    s_cctPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: 1e18,
        localToken: address(s_cctPool),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: "",
        offchainTokenData: ""
      })
    );
  }

  function test_releaseOrMint_RevertWhen_ChainNotAllowed() public {
    uint64 wrongChainSelector = 8838833;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_cctPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        sourceDenominatedAmount: 1,
        localToken: address(s_cctPool),
        remoteChainSelector: wrongChainSelector,
        sourcePoolAddress: _generateSourceTokenData().sourcePoolAddress,
        sourcePoolData: _generateSourceTokenData().extraData,
        offchainTokenData: ""
      })
    );
  }
}
