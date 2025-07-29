// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IVerifierRegistry} from "./interfaces/verifiers/IVerifierRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "./libraries/Client.sol";
import {Internal} from "./libraries/Internal.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract VerifierRegistry is IVerifierRegistry, ITypeAndVersion, Ownable2StepMsgSender {
  string public constant override typeAndVersion = "VerifierRegistry 1.7.0-dev";

  mapping(bytes32 verifierId => address) internal s_verifiers;

  function getVerifier(
    bytes32 verifierId
  ) external view returns (address) {
    return s_verifiers[verifierId];
  }

  function addVerifier(bytes32 verifierId, address verifierAddress) external {
    s_verifiers[verifierId] = verifierAddress;
  }
}
