// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";

import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";

contract BurnMintTokenPool_typeAndVersion is TokenPoolSetup {
  function test_typeAndVersion() public {
    BurnMintTokenPool pool = new BurnMintTokenPool(
      IBurnMintERC20(address(s_token)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );
    assertEq(pool.typeAndVersion(), "BurnMintTokenPool 2.0.0-dev");
  }
}
