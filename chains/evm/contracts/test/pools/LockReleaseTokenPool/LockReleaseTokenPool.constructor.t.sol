// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract LockReleaseTokenPool_constructor is BaseTest {
  function test_constructor() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({name: "T", symbol: "T", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER}),
      OWNER,
      OWNER
    );
    ERC20LockBox lockBox = new ERC20LockBox(address(token));

    LockReleaseTokenPool pool = new LockReleaseTokenPool(
      IERC20(address(token)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(lockBox)
    );

    assertEq(address(pool.getToken()), address(token));
    assertEq(pool.typeAndVersion(), "LockReleaseTokenPool 2.0.0-dev");
  }

  function test_constructor_RevertWhen_InvalidToken() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({name: "T", symbol: "T", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER}),
      OWNER,
      OWNER
    );
    ERC20LockBox lockBox = new ERC20LockBox(address(token));

    CrossChainToken invalidToken = new CrossChainToken(
      BaseERC20.ConstructorParams({name: "IT", symbol: "IT", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER}),
      OWNER,
      OWNER
    );

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, address(invalidToken)));
    new LockReleaseTokenPool(
      IERC20(address(invalidToken)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(lockBox)
    );
  }
}
