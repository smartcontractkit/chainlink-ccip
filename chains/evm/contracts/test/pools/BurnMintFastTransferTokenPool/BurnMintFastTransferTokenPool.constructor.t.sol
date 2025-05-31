// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

contract BurnMintFastTransferTokenPool_constructor is BurnMintFastTransferTokenPoolSetup {
  function test_Constructor() public view {
    assertEq(address(s_burnMintERC20), address(s_pool.getToken()));
    assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
    assertEq(address(s_sourceRouter), s_pool.getRouter());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq("BurnMintFastTransferTokenPool 1.6.1", s_pool.typeAndVersion());
  }

  function test_SupportsInterface() public view {
    assertTrue(s_pool.supportsInterface(type(IFastTransferPool).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IAny2EVMMessageReceiver).interfaceId));
  }

  function test_GetDestChainConfig() public view {
    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, FAST_FEE_BPS);
    assertTrue(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, abi.encode(s_remoteBurnMintPool));
    assertEq(config.maxFillAmountPerRequest, FILL_AMOUNT_MAX);
  }

  function test_IsFillerAllowListed() public {
    assertTrue(s_pool.isAllowedFiller(DEST_CHAIN_SELECTOR, s_filler));
    assertFalse(s_pool.isAllowedFiller(DEST_CHAIN_SELECTOR, makeAddr("notFiller")));
  }

  function test_GetAllowListedFillers() public view {
    address[] memory allowlistedFillers = s_pool.getAllowedFillers(DEST_CHAIN_SELECTOR);
    assertEq(allowlistedFillers.length, 1);
    assertEq(allowlistedFillers[0], s_filler);
  }
}
