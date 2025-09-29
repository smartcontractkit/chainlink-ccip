// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice VerifierProxy enables upgrades to Cross-Chain Verifiers (CCVs) without breaking existing references in token pools, receivers, and apps.
/// The address of this contract will be referenced in the following places:
///   - Senders of messages will specify required and optional CCVs as part of ccipSend extraArgs.
///   - Token pools will specify required CCVs on both source and destination.
///   - Receiver contracts will specify required and optional CCVs on destination.
///   - OnRamp will specify default and mandated CCVs for each destination.
///   - OffRamp will specify default and mandated CCVs for each source.
/// Each of these references should be to a VerifierProxy contract, not a Verifier contract directly.
/// @dev On source, the OnRamp will forward requests (i.e. getFee, forwardToVerifier) through this contract to the required Verifier.
/// The same applies on destination. The OffRamp will forward requests (i.e. verifyMessage) through this contract to the required Verifier.
/// To support this proxy, all future Verifier interfaces must have originalCaller defined as the first arg to each method.
contract VerifierProxy is Ownable2StepMsgSender {
  error ZeroAddressNotAllowed();

  event VerifierUpdated(address indexed oldVerifier, address indexed newVerifier);

  /// @notice The address of the verifier contract.
  address internal s_verifier;

  constructor(
    address verifierAddress
  ) {
    _setVerifier(verifierAddress);
  }

  /// @dev Allows for child contracts to modify the access control logic.
  function _onlyAllowedCaller() internal virtual {
    if (msg.sender != owner()) {
      revert OnlyCallableByOwner();
    }
  }

  /// @notice Returns the address of the verifier contract.
  function getVerifier() external view virtual returns (address) {
    return s_verifier;
  }

  /// @notice Sets the address of the verifier contract.
  /// @param verifierAddress The address of the new verifier contract.
  function setVerifier(
    address verifierAddress
  ) external virtual {
    _onlyAllowedCaller();
    _setVerifier(verifierAddress);
  }

  /// @dev Internal method that allows for reuse in constructor.
  function _setVerifier(
    address verifierAddress
  ) internal virtual {
    if (verifierAddress == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    address oldVerifier = s_verifier;
    s_verifier = verifierAddress;
    emit VerifierUpdated(oldVerifier, verifierAddress);
  }

  /// @notice The fallback function forwards all calls to the verifier contract via a call.
  /// @dev The first argument of the calldata is always overwritten with the caller of the proxy. This ensures that the
  /// verifier contract always sees the original caller of the proxy. It also means originalCaller should be the first
  /// argument of any method called via this proxy.
  /// solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external virtual {
    address verifierAddress = s_verifier;
    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())
      // Overwrite calldata with the actual caller.
      // This prevents an attacker from spoofing a different caller.
      // The caller must be at calldata index 4 (skip function selector)
      mstore(4, caller())

      // Forward the call to the verifier contract.
      let success := call(gas(), verifierAddress, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
