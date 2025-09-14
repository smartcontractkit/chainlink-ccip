// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IRouterClient} from "../interfaces/IRouterClient.sol";

import {Client} from "../libraries/Client.sol";
import {CCIPReceiver} from "./CCIPReceiver.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

// @notice Example of a client which supports EVM/non-EVM chains.
// @dev If chain specific logic is required for different chain families (e.g. particular
// decoding the bytes sender for authorization checks), it may be required to point to a helper
// authorization contract unless all chain families are known up front.
// @dev If contract does not implement IAny2EVMMessageReceiver and IERC165,
// and tokens are sent to it, ccipReceive will not be called but tokens will be transferred.
// @dev If the client is upgradeable you have significantly more flexibility and
// can avoid storage based options like the below contract uses. However it's
// worth carefully considering how the trust assumptions of your client dapp will
// change if you introduce upgradability. An immutable dapp building on top of CCIP
// like the example below will inherit the trust properties of CCIP (i.e. the oracle network).
// @dev The receiver's are encoded offchain and passed as direct arguments to permit supporting
// new chain family receivers (e.g. a Solana encoded receiver address) without upgrading.
contract CCIPClientExample is CCIPReceiver, Ownable2StepMsgSender {
  using SafeERC20 for IERC20;

  enum PaymentMethod {
    NativeToken,
    FeeToken
  }

  error InvalidPaymentMethod(PaymentMethod method);
  error InvalidAddress(address);
  error InvalidAmount(uint256);
  error InvalidRemoteChain(uint64 remoteChainSelector);
  error InsufficientNativeTokenBalance(uint256 balance);
  error InsufficientFeeTokenBalance(uint256 balance);
  error NativeTokenTransferFailed(address to, uint256 amount);
  error FeeTokenTransferFailed(address token, address to, uint256 amount);

  event MessageSent(bytes32 messageId);
  event MessageReceived(bytes32 messageId);
  event NativeTokenDeposited(address from, uint256 amount);
  event FeeTokenDeposited(address from, uint256 amount);
  event NativeTokenWithdrawn(address withdrawer, address to, uint256 amount);
  event FeeTokenWithdrawn(address token, address withdrawer, address to, uint256 amount);

  /// @notice Configuration for a remote chain.
  /// @dev extraArgsBytes are added to a msg on source, CCV params are checked on dest.
  struct RemoteChainConfig {
    bytes extraArgsBytes;
    address[] requiredCCVs;
    address[] optionalCCVs;
    uint8 optionalThreshold;
  }

  // Current feeToken
  IERC20 public s_feeToken;
  // Below is a simplistic example (same params for all messages) of using storage to allow for new options without
  // upgrading the dapp. Note that extra args are chain family specific (e.g. gasLimit is EVM specific etc.).
  // and will always be backwards compatible i.e. upgrades are opt-in.
  // Offchain we can compute the V1 extraArgs:
  //    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: 300_000});
  //    bytes memory encodedV1ExtraArgs = Client._argsToBytes(extraArgs);
  // Then later compute V2 extraArgs, for example if a refund feature was added:
  //    Client.EVMExtraArgsV2 memory extraArgs = Client.EVMExtraArgsV2({gasLimit: 300_000, destRefundAddress: 0x1234});
  //    bytes memory encodedV2ExtraArgs = Client._argsToBytes(extraArgs);
  // and update storage with the new args.
  // If different options are required for different messages, for example different gas limits,
  // one can simply key based on (remoteChainSelector, messageType) instead of only chainSelector.
  mapping(uint64 remoteChainSelector => RemoteChainConfig remoteChainConfig) internal s_remoteChains;
  // Balance of native token held for each user.
  mapping(address => uint256) public s_nativeTokenBalances;
  // Balance of fee token held for each user.
  mapping(address => uint256) public s_feeTokenBalances;

  constructor(IRouterClient router, IERC20 feeToken) CCIPReceiver(address(router)) {
    s_feeToken = feeToken;
    s_feeToken.approve(address(router), type(uint256).max);
  }

  function getRemoteChainConfig(
    uint64 remoteChainSelector
  ) external view returns (RemoteChainConfig memory) {
    return s_remoteChains[remoteChainSelector];
  }

  function enableRemoteChain(
    uint64 remoteChainSelector,
    bytes memory extraArgs,
    address[] memory requiredCCVs,
    address[] memory optionalCCVs,
    uint8 optionalThreshold
  ) external onlyOwner {
    s_remoteChains[remoteChainSelector] = RemoteChainConfig({
      extraArgsBytes: extraArgs,
      requiredCCVs: requiredCCVs,
      optionalCCVs: optionalCCVs,
      optionalThreshold: optionalThreshold
    });
  }

  function disableRemoteChain(
    uint64 remoteChainSelector
  ) external onlyOwner {
    delete s_remoteChains[remoteChainSelector];
  }

  function ccipReceive(
    Client.Any2EVMMessage calldata message
  ) external virtual override onlyRouter validRemoteChain(message.sourceChainSelector) {
    // Extremely important to ensure only router calls this.
    // Tokens in message if any will be transferred to this contract.
    // TODO: Validate sender/origin chain and process message and/or tokens.
    _ccipReceive(message);
  }

  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal override {
    emit MessageReceived(message.messageId);
  }

  /// @notice sends data to receiver on dest chain.
  function sendData(
    PaymentMethod method,
    uint64 destChainSelector,
    bytes memory receiver,
    bytes memory data
  ) external validRemoteChain(destChainSelector) validPaymentMethod(method) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_remoteChains[destChainSelector].extraArgsBytes,
      feeToken: method == PaymentMethod.FeeToken ? address(s_feeToken) : address(0)
    });
    bytes32 messageId = _ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  /// @notice sends data to receiver on dest chain.
  function sendDataAndTokens(
    PaymentMethod method,
    uint64 destChainSelector,
    bytes memory receiver,
    bytes memory data,
    Client.EVMTokenAmount[] memory tokenAmounts
  ) external validRemoteChain(destChainSelector) validPaymentMethod(method) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_remoteChains[destChainSelector].extraArgsBytes,
      feeToken: method == PaymentMethod.FeeToken ? address(s_feeToken) : address(0)
    });
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }
    bytes32 messageId = _ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  /// @notice user sends tokens to a receiver.
  /// Approvals can be optimized with a whitelist of tokens and inf approvals if desired.
  function sendTokens(
    PaymentMethod method,
    uint64 destChainSelector,
    bytes memory receiver,
    Client.EVMTokenAmount[] memory tokenAmounts
  ) external validRemoteChain(destChainSelector) validPaymentMethod(method) {
    bytes memory data;
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_remoteChains[destChainSelector].extraArgsBytes,
      feeToken: method == PaymentMethod.FeeToken ? address(s_feeToken) : address(0)
    });
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }
    bytes32 messageId = _ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  /// @notice Return the CCVs required/optional for a source chain.
  function getCCVs(
    uint64 sourceChainSelector
  )
    external
    view
    override
    returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold)
  {
    RemoteChainConfig memory config = s_remoteChains[sourceChainSelector];
    return (config.requiredCCVs, config.optionalCCVs, config.optionalThreshold);
  }

  /// @notice Provide fee tokens to the contract.
  function provideFeeToken(
    uint256 amount
  ) external {
    s_feeTokenBalances[msg.sender] += amount;

    s_feeToken.safeTransferFrom(msg.sender, address(this), amount);
    emit FeeTokenDeposited(msg.sender, amount);
  }

  /// @notice Provide native token to the contract.
  function provideNativeToken() external payable {
    if (msg.value == 0) revert InvalidAmount(msg.value);
    s_nativeTokenBalances[msg.sender] += msg.value;

    emit NativeTokenDeposited(msg.sender, msg.value);
  }

  /// @notice Withdraw fee tokens to an address.
  function withdrawFeeToken(address to, uint256 amount) external {
    if (to == address(0)) revert InvalidAddress(to);
    if (amount == 0) revert InvalidAmount(amount);
    uint256 balance = s_feeTokenBalances[msg.sender];
    if (balance < amount) {
      revert InsufficientFeeTokenBalance(balance);
    }
    s_feeTokenBalances[msg.sender] = balance - amount;

    s_feeToken.safeTransfer(to, amount);
    emit FeeTokenWithdrawn(address(s_feeToken), msg.sender, to, amount);
  }

  /// @notice Withdraw native token to an address.
  function withdrawNativeToken(address to, uint256 amount) external {
    if (to == address(0)) revert InvalidAddress(to);
    if (amount == 0) revert InvalidAmount(amount);
    uint256 balance = s_nativeTokenBalances[msg.sender];
    if (balance < amount) {
      revert InsufficientNativeTokenBalance(balance);
    }
    s_nativeTokenBalances[msg.sender] = balance - amount;

    (bool success,) = to.call{value: amount}("");
    if (!success) {
      revert NativeTokenTransferFailed(to, amount);
    }
    emit NativeTokenWithdrawn(msg.sender, to, amount);
  }

  function _ccipSend(uint64 destChainSelector, Client.EVM2AnyMessage memory message) internal returns (bytes32) {
    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);
    uint256 value;
    if (message.feeToken == address(0)) {
      if (s_nativeTokenBalances[msg.sender] < fee) {
        revert InsufficientNativeTokenBalance(s_nativeTokenBalances[msg.sender]);
      }
      s_nativeTokenBalances[msg.sender] -= fee;
      value = fee;
    } else {
      if (s_feeTokenBalances[msg.sender] < fee) {
        revert InsufficientFeeTokenBalance(s_feeTokenBalances[msg.sender]);
      }
      s_feeTokenBalances[msg.sender] -= fee;
    }
    return IRouterClient(i_ccipRouter).ccipSend{value: value}(destChainSelector, message);
  }

  modifier validRemoteChain(
    uint64 remoteChainSelector
  ) {
    if (s_remoteChains[remoteChainSelector].extraArgsBytes.length == 0) revert InvalidRemoteChain(remoteChainSelector);
    _;
  }

  modifier validPaymentMethod(
    PaymentMethod method
  ) {
    if (method != PaymentMethod.NativeToken && method != PaymentMethod.FeeToken) {
      revert InvalidPaymentMethod(method);
    }
    _;
  }
}
