// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice Proxy forwards calls to a target contract.
contract Proxy is Ownable2StepMsgSender {
  error ZeroAddressNotAllowed();

  event TargetUpdated(address indexed oldTarget, address indexed newTarget);

  /// @notice The address of the target contract.
  address internal s_target;

  constructor(
    address target
  ) {
    _setTarget(target);
  }

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
