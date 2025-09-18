// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

contract Proxiable {
    error NewContractNotProxiable();

    /// @notice Updates the implementation address to `newAddress`.
    /// @dev The new implementation must be proxiable, meaning it must implement `proxiableUUID` and return the same storage slot.
    function updateCodeAddress(address newAddress) internal {
        bytes32 slot = proxiableUUID();
        if (Proxiable(newAddress).proxiableUUID() != slot) {
            revert NewContractNotProxiable();
        }
        assembly {
            sstore(slot, newAddress)
        }
    }

    /// @notice Returns the storage slot that the implementation address is stored in.
    /// @dev Per UUPS, the address of the implementation is stored @ keccak256("PROXIABLE") = "0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7"
    function proxiableUUID() public pure returns (bytes32) {
        return 0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7;
    }
}
