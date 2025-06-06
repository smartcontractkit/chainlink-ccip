// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolSetup} from "./USDCTokenPoolSetup.t.sol";

contract USDCTokenPool_setMintRecipientOverride is USDCTokenPoolSetup {
  function test_SetMintRecipientOverride_Success() public {
    bytes32 receiver = bytes32(uint256(uint160(STRANGER)));
    uint256 amount = 1;
    s_token.transfer(address(s_usdcTokenPool), amount);
    vm.startPrank(s_routerAllowedOnRamp);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_SELECTOR);

    vm.startPrank(OWNER);

    bytes32 mintRecipient = keccak256(abi.encodePacked(address(this)));

    uint64[] memory chainSelectors = new uint64[](1);
    bytes32[] memory mintRecipients = new bytes32[](1);

    chainSelectors[0] = DEST_CHAIN_SELECTOR;
    mintRecipients[0] = mintRecipient;

    // Set the mint recipient override
    vm.expectEmit();
    emit USDCTokenPool.MintRecipientOverrideSet(chainSelectors[0], mintRecipients[0]);

    s_usdcTokenPool.setMintRecipientOverrides(chainSelectors, mintRecipients);

    assertEq(s_usdcTokenPool.getMintRecipientOverride(chainSelectors[0]), mintRecipients[0]);

    // Expect the mint recipient override to be used in the deposit for burn event
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      s_mockUSDC.s_nonce(),
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      mintRecipient,
      expectedDomain.domainIdentifier,
      s_mockUSDC.DESTINATION_TOKEN_MESSENGER(),
      expectedDomain.allowedCaller
    );

    vm.startPrank(s_routerAllowedOnRamp);

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(receiver),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_RevertWhen_MismatchedArrayLengths() public {
    uint64[] memory chainSelectors = new uint64[](1);
    bytes32[] memory mintRecipients = new bytes32[](2);

    vm.expectRevert(TokenPool.MismatchedArrayLengths.selector);
    s_usdcTokenPool.setMintRecipientOverrides(chainSelectors, mintRecipients);
  }

  function test_RevertWhen_UnknownDomain() public {
    uint64[] memory chainSelectors = new uint64[](1);
    bytes32[] memory mintRecipients = new bytes32[](1);

    chainSelectors[0] = type(uint64).max;
    mintRecipients[0] = keccak256(abi.encodePacked(address(this)));

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.UnknownDomain.selector, chainSelectors[0]));
    s_usdcTokenPool.setMintRecipientOverrides(chainSelectors, mintRecipients);
  }
}
