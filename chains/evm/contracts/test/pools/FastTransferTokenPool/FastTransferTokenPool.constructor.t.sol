// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

contract FastTransferTokenPool_constructor is FastTransferTokenPoolSetup {
  function test_Constructor() public view {
    assertEq(address(s_token), address(s_pool.getToken()));
    assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
    assertEq(address(s_sourceRouter), s_pool.getRouter());
    assertFalse(s_pool.getAllowListEnabled());
  }

  function test_SupportsInterface() public view {
    assertTrue(s_pool.supportsInterface(type(IFastTransferPool).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IAny2EVMMessageReceiver).interfaceId));
  }

  function test_GetDestChainConfig() public view {
    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, FAST_FEE_FILLER_BPS);
    assertTrue(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, destPoolAddress);
    assertEq(config.maxFillAmountPerRequest, MAX_FILL_AMOUNT_PER_REQUEST);
  }

  function test_IsFillerAllowListed() public {
    assertTrue(s_pool.isAllowedFiller(s_filler));
    assertFalse(s_pool.isAllowedFiller(makeAddr("notFiller")));
  }

  function test_GetAllowListedFillers() public view {
    address[] memory allowlistedFillers = s_pool.getAllowedFillers();
    assertEq(allowlistedFillers.length, 1);
    assertEq(allowlistedFillers[0], s_filler);
  }
}
