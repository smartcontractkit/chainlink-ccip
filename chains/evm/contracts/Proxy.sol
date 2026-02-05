// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {FeeTokenHandler} from "./libraries/FeeTokenHandler.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice Proxy forwards calls to a target contract.
contract Proxy is ITypeAndVersion, Ownable2StepMsgSender {
  string public constant override typeAndVersion = "Proxy 2.0.0-dev";

  error ZeroAddressNotAllowed();

  event TargetUpdated(address indexed oldTarget, address indexed newTarget);
  event FeeAggregatorUpdated(address indexed oldFeeAggregator, address indexed newFeeAggregator);

  /// @notice The address of the target contract.
  address internal s_target;

  /// @notice The address fees are sent to.
  address internal s_feeAggregator;

  constructor(
    address target,
    address feeAggregator
  ) {
    _setTarget(target);
    _setFeeAggregator(feeAggregator);
  }

  // ================================================================
  // │                           Target                             │
  // ================================================================

  /// @notice Returns the address of the target contract.
  function getTarget() external view virtual returns (address) {
    return s_target;
  }

  /// @notice Sets the address of the target contract.
  /// @param target The address of the new target contract.
  function setTarget(
    address target
  ) external virtual onlyOwner {
    _setTarget(target);
  }

  /// @dev Internal method that allows for reuse in constructor.
  function _setTarget(
    address target
  ) internal virtual {
    if (target == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    address oldTarget = s_target;
    s_target = target;
    emit TargetUpdated(oldTarget, target);
  }

  // ================================================================
  // │                        Fee Aggregator                        │
  // ================================================================

  /// @notice Returns the address of the fee aggregator.
  function getFeeAggregator() external view virtual returns (address) {
    return s_feeAggregator;
  }

  /// @notice Sets the address of the fee aggregator.
  /// @param feeAggregator The address of the new fee aggregator contract.
  /// @dev FeeTokenHandler will revert if feeAggregator is zero when withdrawing fees.
  /// @dev A zero address fee aggregator is valid, and intentionally reverts calls to withdraw fee tokens.
  function setFeeAggregator(
    address feeAggregator
  ) external virtual onlyOwner {
    _setFeeAggregator(feeAggregator);
  }

  /// @dev Internal method that allows for reuse in constructor.
  /// @dev FeeTokenHandler will revert if feeAggregator is zero when withdrawing fees.
  /// @dev A zero address fee aggregator is valid, and intentionally reverts calls to withdraw fee tokens.
  function _setFeeAggregator(
    address feeAggregator
  ) internal virtual {
    address oldFeeAggregator = s_feeAggregator;
    s_feeAggregator = feeAggregator;
    emit FeeAggregatorUpdated(oldFeeAggregator, feeAggregator);
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external virtual {
    FeeTokenHandler._withdrawFeeTokens(feeTokens, s_feeAggregator);
  }

  // ================================================================
  // │                           Fallback                           │
  // ================================================================

  /// @notice The fallback function forwards all calls to the target contract via a call.
  /// solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external virtual {
    address target = s_target;
    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())

      // Forward the call to the target contract.
      let success := call(gas(), target, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
