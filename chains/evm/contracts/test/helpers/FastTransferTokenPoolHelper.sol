  // SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {FastTransferTokenPoolAbstract} from "../../pools/FastTransferTokenPoolAbstract.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract FastTransferTokenPoolHelper is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  error NotAdmin();

  string public constant override typeAndVersion = "FastTransferTokenPoolHelper 1.6.1";

  address internal immutable i_admin;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router) {
    i_admin = msg.sender;
  }

  // Public wrappers for internal functions
  function handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) public {
    _handleTokenToTransfer(destinationChainSelector, sender, amount);
  }

  function transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 srcDecimals
  ) public returns (uint256) {
    return _transferFromFiller(sourceChainSelector, filler, receiver, srcAmount, srcDecimals);
  }

  // Implementation of abstract functions
  function _handleTokenToTransfer(uint64, address sender, uint256 amount) internal override {
    // For testing, we'll just transfer tokens from sender to this contract
    getToken().safeTransferFrom(sender, address(this), amount);
  }

  function _transferFromFiller(
    uint64,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8
  ) internal override returns (uint256) {
    getToken().safeTransferFrom(filler, receiver, srcAmount);
    return srcAmount;
  }

  function _settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    uint256 srcAmount,
    uint8 srcDecimal,
    uint256 fastTransferFee,
    address receiver
  ) internal override {
    // Call the common settlement logic from the abstract contract
    super._settle(
      sourceChainSelector, fillRequestId, sourcePoolAddress, srcAmount, srcDecimal, fastTransferFee, receiver
    );
  }

  /// @notice Validates settlement prerequisites - simple implementation for testing
  function _validateSettlement(uint64, bytes memory) internal view override {
    // For testing, we'll do minimal validation
    // Real implementations would check RMN curse and source pool validation
  }

  /// @notice Handles settlement when the request was not fast-filled
  function _handleNotFastFilled(uint64, uint256 settlementAmountLocal, address receiver) internal override {
    // For testing, just transfer tokens to receiver
    getToken().safeTransfer(receiver, settlementAmountLocal);
  }

  /// @notice Handles reimbursement when the request was fast-filled
  function _handleFastFilledReimbursement(address filler, uint256 settlementAmountLocal) internal override {
    // For testing, just transfer tokens to filler
    getToken().safeTransfer(filler, settlementAmountLocal);
  }

  function _checkAdmin() internal view override {
    if (msg.sender != i_admin) revert NotAdmin();
  }

  // TokenPool function implementations
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) external override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);
    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: _encodeLocalDecimals()
    });
  }

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) external override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);
    uint256 localAmount =
      _calculateLocalAmount(releaseOrMintIn.amount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData));
    getToken().safeTransfer(releaseOrMintIn.receiver, localAmount);
    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }
}
