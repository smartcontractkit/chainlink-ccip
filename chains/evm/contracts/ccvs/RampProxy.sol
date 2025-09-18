// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice RampProxy enables upgrades to CCVRamps without breaking existing references in token pools, receivers, and apps.
/// The address of this contract will be referenced in the following places:
///   - Users / apps will specify required and optional CCVs as part of ccipSend extraArgs.
///   - Token pools will specify required and optional CCVs on both source and destination.
///   - Receiver contracts will specify required and optional CCVs on destination.
///   - CCVProxy will specify default and mandated CCVs for each destination.
///   - CCVAggregator will specify default and mandated CCVs for each source.
/// Each of these references should be to a RampProxy contract, not a CCVRamp directly.
/// @dev On source, the CCVProxy will forward requests (i.e. getFee, forwardToVerifier) through this contract to the required CCVRamp.
/// The same applies on destination. The CCVAggregator will forward requests (i.e. verifyMessage) through this contract to the required CCVRamp.
/// This contract follows the Universal Upgradeable Proxy Standard (UUPS) & uses delegatecall so the CCVRamp implementation can access `msg.sender`
/// without needing to pass it through explicitly and thus modify the function signatures of the ICCVOnRamp and ICCVOffRamp interfaces.
contract RampProxy {
  constructor(address rampAddress) {
    // Per UUPS, the address of the implementation is stored @ keccak256("PROXIABLE") = "0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7"
    assembly {
        sstore(0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7, rampAddress)
    }
  }

  // The fallback function forwards all calls to the ramp contract via delegatecall.
  // solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external {
    assembly {
      let rampAddress := sload(0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7)
      // Store the calldata in memory.
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())

      // Forward the call to the ramp contract.
      let result := delegatecall(gas(), rampAddress, 0, calldatasize(), 0, 0)
      // Copy the returned data to memory.
      // Revert based on the result of the delegatecall.
      returndatacopy(0, 0, returndatasize())
      switch result
      case 0 {
        revert(0, returndatasize())
      }
      default {
        return(0, returndatasize())
      }
    }
  }
}
