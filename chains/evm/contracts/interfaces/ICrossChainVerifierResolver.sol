// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

/// @notice Resolves and returns the appropriate verifier contract for the given outbound / inbound traffic.
interface ICrossChainVerifierResolver is IERC165 {
  /// @notice Returns the appropriate verifier contract based on the given ccvData.
  /// @dev The OffRamp is responsible for calling this function using the ccvData it receives from the executor.
  /// If the verifier specified by the executor is actually a resolver, the OffRamp will call this function to get the actual verifier contract.
  /// Verifiers can build resolvers that process the ccvData in accordance with how their verifier forms ccvData. For example, their verifier may
  /// prefix the ccvData with a version identifier, which the resolver can parse to determine the correct verifier contract.
  /// @param ccvData The ccvData formed by the verifier.
  /// @return verifierAddress The address of the verifier contract.
  function getInboundImplementation(
    bytes calldata ccvData
  ) external view returns (address);

  /// @notice Returns the appropriate verifier contract based on the given destChainSelector.
  /// @dev The OnRamp is responsible for calling this function using the destChainSelector specified by the sender.
  /// If the verifier specified by the sender is actually a resolver, the OnRamp will call this function to get the actual verifier contract.
  /// For example, resolvers can maintain a simple mapping of destChainSelector to verifier contract address.
  /// @param destChainSelector The destChainSelector for a message.
  /// @return verifierAddress The address of the verifier contract.
  function getOutboundImplementation(
    uint64 destChainSelector
  ) external view returns (address);
}
