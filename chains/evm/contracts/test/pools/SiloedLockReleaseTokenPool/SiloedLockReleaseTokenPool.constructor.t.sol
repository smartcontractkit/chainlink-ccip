// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract SiloedLockReleaseTokenPool_constructor is BaseTest {
  function test_constructor() public {
    BurnMintERC20 token = new BurnMintERC20("TKN", "T", DEFAULT_TOKEN_DECIMALS, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(token), 0);

    SiloedLockReleaseTokenPool pool = new SiloedLockReleaseTokenPool(
      token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(lockBox)
    );

    assertEq(address(pool.getToken()), address(token));
    assertEq(pool.getUnsiloedLiquidity(), 0);
    assertEq(pool.typeAndVersion(), "SiloedLockReleaseTokenPool 1.7.0-dev");
  }

  function test_constructor_RevertWhen_InvalidLockBoxLiquidityDomain() public {
    BurnMintERC20 token = new BurnMintERC20("TKN", "T", DEFAULT_TOKEN_DECIMALS, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(token), bytes32(uint256(5)));

    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InvalidLockBoxLiquidityDomain.selector, uint64(5))
    );
    new SiloedLockReleaseTokenPool(
      token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(lockBox)
    );
  }

  function test_constructor_RevertWhen_InvalidToken() public {
    BurnMintERC20 token = new BurnMintERC20("TKN", "T", DEFAULT_TOKEN_DECIMALS, 0, 0);
    BurnMintERC20 otherToken = new BurnMintERC20("OTH", "O", DEFAULT_TOKEN_DECIMALS, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(otherToken), 0);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, address(otherToken)));
    new SiloedLockReleaseTokenPool(
      token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(lockBox)
    );
  }
}
