pragma solidity ^0.8.25;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

abstract contract HyperEVMLinker is Ownable2StepMsgSender {
  // keccak256("HyperCore deployer")
  bytes32 private constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  error LinkerAddressCannotBeZero();

  event HyperEVMLinkerSet(address indexed hyperEVMLinker);

  /// @notice Sets the hyperEVMLinker address.
  /// @param newHyperEVMLinker The address of the hyperEVMLinker.
  function setHyperEVMLinker(
    address newHyperEVMLinker
  ) external onlyOwner {
    if (newHyperEVMLinker == address(0)) {
      revert LinkerAddressCannotBeZero();
    }

    assembly {
      sstore(HYPER_EVM_LINKER_SLOT, newHyperEVMLinker)
    }

    emit HyperEVMLinkerSet(newHyperEVMLinker);
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
