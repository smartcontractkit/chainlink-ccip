// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract LockReleaseTokenPool_constructor is BaseTest {
  function test_constructor() public {
    BurnMintERC20 token = new BurnMintERC20("T", "T", 18, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(token));

    LockReleaseTokenPool pool = new LockReleaseTokenPool(
      token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(lockBox)
    );

    assertEq(address(pool.getToken()), address(token));
    assertEq(pool.typeAndVersion(), "LockReleaseTokenPool 1.7.0-dev");
  }

  function test_constructor_RevertWhen_InvalidToken() public {
    BurnMintERC20 token = new BurnMintERC20("T", "T", 18, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(token));

    BurnMintERC20 invalidToken = new BurnMintERC20("IT", "IT", 18, 0, 0);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, address(invalidToken)));
    new LockReleaseTokenPool(
      invalidToken,
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(lockBox)
    );
  }
}
