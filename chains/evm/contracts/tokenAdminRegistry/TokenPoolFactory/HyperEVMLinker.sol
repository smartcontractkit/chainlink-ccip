pragma solidity ^0.8.25;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @title HyperEVMLinker
/// @notice This contract is used to facilitate linking of an ERC20 from HyperEVM to HyperCore by storing the linker
/// address in the required slot.
/// @dev This contract is built in accordance with the Hyperliquid standard at the following URL:
/// https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/hyperevm/hypercore-less-than-greater-than-hyperevm-transfers#linking-core-and-evm-spot-assets
abstract contract HyperEVMLinker is Ownable2StepMsgSender {
  // In order to bridge to HyperCore, factory-deployed contracts must store the address of a finalizer linker at
  // storage slot keccak256("HyperCore deployer")
  bytes32 internal constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  error LinkerAddressCannotBeZero();

  event HyperEVMLinkerSet(address indexed hyperEVMLinker);

  /// @notice Sets the hyperEVMLinker address.
  /// @param newLinker The address of the hyperEVMLinker.
  function setHyperEVMLinker(
    address newLinker
  ) external onlyOwner {
    if (newLinker == address(0)) {
      revert LinkerAddressCannotBeZero();
    }

    assembly {
      sstore(HYPER_EVM_LINKER_SLOT, newLinker)
    }

    emit HyperEVMLinkerSet(newLinker);
  }

  /// @notice Gets the hyperEVMLinker address.
  /// @return hyperEVMLinker The address of the hyperEVMLinker.
  function getHyperEVMLinker() public view returns (address) {
    address hyperEVMLinker;
    assembly {
      hyperEVMLinker := sload(HYPER_EVM_LINKER_SLOT)
    }
    return hyperEVMLinker;
  }
}
