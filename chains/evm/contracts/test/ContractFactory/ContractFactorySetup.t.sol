// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {ContractFactory} from "../../ContractFactory.sol";
import {BaseTest} from "../BaseTest.t.sol";

// We need to use a real contract because mockCallRevert causes issues with CREATE2.
// Calling mockCallRevert against a predicted address creates the contract, which breaks the following CREATE2 call.
contract Storage is Ownable2StepMsgSender {
  error InvalidValue();

  uint256 private s_value;

  constructor(
    uint256 value
  ) {
    _setValue(value);
  }

  function setValue(
    uint256 value
  ) external onlyOwner {
    _setValue(value);
  }

  function _setValue(
    uint256 value
  ) internal {
    if (value == 0) {
      revert InvalidValue();
    }
    s_value = value;
  }

  function getValue() external view returns (uint256) {
    return s_value;
  }
}

contract ContractFactorySetup is BaseTest {
  ContractFactory internal s_contractFactory;
  bytes32 internal constant SALT = keccak256("SALT");
  address internal s_allowedCaller;
  address internal s_invalidCaller;

  function setUp() public virtual override {
    super.setUp();

    s_allowedCaller = makeAddr("ALLOWED_CALLER");
    s_invalidCaller = makeAddr("INVALID_CALLER");

    address[] memory allowList = new address[](1);
    allowList[0] = s_allowedCaller;
    s_contractFactory = new ContractFactory(allowList);
  }

  function getStorageCreationCode(
    uint256 value
  ) public pure returns (bytes memory) {
    return abi.encodePacked(type(Storage).creationCode, abi.encode(value));
  }
}
