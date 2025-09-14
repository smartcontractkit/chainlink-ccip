pragma solidity ^0.8.0;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";

import {CCIPClientExample} from "../../../applications/CCIPClientExample.sol";
import {Client} from "../../../libraries/Client.sol";
import {OnRampSetup} from "../../onRamp/OnRamp/OnRampSetup.t.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {ERC165Checker} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/ERC165Checker.sol";

contract CCIPClientExample_sanity is OnRampSetup {
  function test_ImmutableExamples() public {
    CCIPClientExample exampleContract = new CCIPClientExample(s_sourceRouter, IERC20(s_sourceFeeToken));
    deal(OWNER, 100 ether);
    deal(s_sourceFeeToken, OWNER, 100 ether);
    exampleContract.provideNativeToken{value: 50 ether}();
    IERC20(s_sourceFeeToken).approve(address(exampleContract), 50 ether);
    exampleContract.provideFeeToken(50 ether);

    // feeToken approval works
    assertEq(IERC20(s_sourceFeeToken).allowance(address(exampleContract), address(s_sourceRouter)), 2 ** 256 - 1);

    // Can enable chain
    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: 300_000});
    bytes memory encodedExtraArgs = Client._argsToBytes(extraArgs);
    exampleContract.enableRemoteChain(DEST_CHAIN_SELECTOR, encodedExtraArgs, new address[](0), new address[](0), 0);
    assertEq(exampleContract.getRemoteChainConfig(DEST_CHAIN_SELECTOR).extraArgsBytes, encodedExtraArgs);

    address toAddress = makeAddr("toAddress");

    // Can send data pay native
    exampleContract.sendData(
      CCIPClientExample.PaymentMethod.NativeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello")
    );

    // Can send data pay feeToken
    exampleContract.sendData(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello")
    );

    // Can send data tokens
    address sourceToken = s_sourceTokens[1];
    assertEq(
      address(s_onRamp.getPoolBySourceToken(DEST_CHAIN_SELECTOR, IERC20(sourceToken))),
      address(s_sourcePoolByToken[sourceToken])
    );
    deal(sourceToken, OWNER, 100 ether);
    IERC20(sourceToken).approve(address(exampleContract), 1 ether);
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: sourceToken, amount: 1 ether});
    exampleContract.sendDataAndTokens(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello"), tokenAmounts
    );
    // Tokens transferred from owner to router then burned in pool.
    assertEq(IERC20(sourceToken).balanceOf(OWNER), 99 ether);
    assertEq(IERC20(sourceToken).balanceOf(address(s_sourceRouter)), 0);

    // Can send just tokens
    IERC20(sourceToken).approve(address(exampleContract), 1 ether);
    exampleContract.sendTokens(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), tokenAmounts
    );

    // Can receive
    assertTrue(ERC165Checker.supportsInterface(address(exampleContract), type(IAny2EVMMessageReceiver).interfaceId));

    // Can withdraw native token
    uint256 preNativeBalance = address(OWNER).balance;
    exampleContract.withdrawNativeToken(OWNER, 10 ether);
    assertEq(address(OWNER).balance, preNativeBalance + 10 ether);

    // Can withdraw fee token
    uint256 preFeeBalance = IERC20(s_sourceFeeToken).balanceOf(OWNER);
    exampleContract.withdrawFeeToken(OWNER, 10 ether);
    assertEq(IERC20(s_sourceFeeToken).balanceOf(OWNER), preFeeBalance + 10 ether);

    /////////////////////
    // STRANGER CHECKS //
    /////////////////////
    vm.startPrank(STRANGER);

    // Stranger can't withdraw (doesn't have funds on contract)
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientNativeTokenBalance.selector, 0));
    exampleContract.withdrawNativeToken(STRANGER, 10 ether);
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientFeeTokenBalance.selector, 0));
    exampleContract.withdrawFeeToken(STRANGER, 10 ether);

    // Stranger can't sendDataPayNative, no native funds allocated on client
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientNativeTokenBalance.selector, 0));
    exampleContract.sendData(
      CCIPClientExample.PaymentMethod.NativeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello")
    );

    // Stranger can't sendDataPayFeeToken, no fee token funds allocated on client
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientFeeTokenBalance.selector, 0));
    exampleContract.sendData(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello")
    );

    // Stranger can't sendDataAndTokens, no fee token funds allocated on client
    deal(sourceToken, STRANGER, 1 ether);
    IERC20(sourceToken).approve(address(exampleContract), 1 ether);
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientFeeTokenBalance.selector, 0));
    exampleContract.sendDataAndTokens(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello"), tokenAmounts
    );

    // Stranger can't sendTokens, no fee token funds allocated on client
    vm.expectRevert(abi.encodeWithSelector(CCIPClientExample.InsufficientFeeTokenBalance.selector, 0));
    exampleContract.sendTokens(
      CCIPClientExample.PaymentMethod.FeeToken, DEST_CHAIN_SELECTOR, abi.encode(toAddress), tokenAmounts
    );
  }
}
