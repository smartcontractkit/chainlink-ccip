// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMN} from "../../../rmn/RMN.sol";
import {RMNRemoteSetup} from "./RMNSetup.t.sol";

contract RMNRemote_getVersionedConfig is RMNRemoteSetup {
  function test_getVersionedConfig_ReturnsEmptyConfig() public view {
    (uint32 version, RMN.Config memory config) = s_rmn.getVersionedConfig();

    assertEq(version, 0);
    assertEq(config.rmnHomeContractConfigDigest, bytes32(0));
    assertEq(config.signers.length, 0);
    assertEq(config.fSign, 0);
  }
}
