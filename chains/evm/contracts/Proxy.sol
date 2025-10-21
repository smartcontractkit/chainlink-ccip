// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice Proxy enables upgrades to Cross-Chain Verifiers (CCVs) and Executors without breaking existing references in token pools, receivers, and apps.
/// The address of this contract will be referenced in the following places:
///   - Senders of messages can specify required and optional CCVs + a desired executor as part of ccipSend extraArgs.
///   - Token pools will specify required CCVs on both source and destination.
///   - Receiver contracts will specify required and optional CCVs on destination.
///   - OnRamp will specify default and mandated CCVs + a default executor for each destination.
///   - OffRamp will specify default and mandated CCVs for each source.
/// Each of these references should be to a Proxy contract, not a CCV / Executor directly.
/// @dev On source, the OnRamp will forward requests (i.e. getFee, forwardToVerifier) through this contract to the required CCV / Executor.
/// The same applies on destination. The OffRamp will forward requests (i.e. verifyMessage) through this contract to the required CCV.
/// To support this proxy, all relevant interfaces must define originalCaller as the first arg to each method.
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

  /// @dev Allows for child contracts to modify the access control logic.
  function _onlyAllowedCaller() internal virtual {
    if (msg.sender != owner()) {
      revert OnlyCallableByOwner();
    }
  }

  /// @notice Returns the address of the target contract.
  function getTarget() external view virtual returns (address) {
    return s_target;
  }

  /// @notice Sets the address of the target contract.
  /// @param target The address of the new target contract.
  function setTarget(
    address target
  ) external virtual {
    _onlyAllowedCaller();
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
  /// @dev The first argument of the calldata is always overwritten with the caller of the proxy. This ensures that the
  /// target contract always sees the original caller of the proxy. It also means originalCaller should be the first
  /// argument of any method called via this proxy.
  /// solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external virtual {
    address target = s_target;
    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())
      // Overwrite calldata with the actual caller.
      // This prevents an attacker from spoofing a different caller.
      // The caller must be at calldata index 4 (skip function selector)
      mstore(4, caller())

      // Forward the call to the target contract.
      let success := call(gas(), target, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
