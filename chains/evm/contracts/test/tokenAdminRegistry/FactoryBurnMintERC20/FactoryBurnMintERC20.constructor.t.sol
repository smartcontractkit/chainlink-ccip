// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {FactoryBurnMintERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/FactoryBurnMintERC20.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

contract FactoryBurnMintERC20_constructor is BurnMintERC20Setup {
  function test_Constructor() public {
    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    uint256 maxSupply = 1e33;

    s_burnMintERC20 = new FactoryBurnMintERC20(name, symbol, decimals, maxSupply, 1e18, s_alice);

    assertEq(name, s_burnMintERC20.name());
    assertEq(symbol, s_burnMintERC20.symbol());
    assertEq(decimals, s_burnMintERC20.decimals());
    assertEq(maxSupply, s_burnMintERC20.maxSupply());

    assertEq(s_burnMintERC20.balanceOf(s_alice), 1e18);
    assertEq(s_burnMintERC20.totalSupply(), 1e18);
    assertEq(s_burnMintERC20.typeAndVersion(), "FactoryBurnMintERC20 1.6.2");
  }

  function test_Constructor_When_MaxSupplyIsZero() public {
    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    uint256 maxSupply = 0;

    s_burnMintERC20 = new FactoryBurnMintERC20(name, symbol, decimals, maxSupply, 1e18, s_alice);

    assertEq(name, s_burnMintERC20.name());
    assertEq(symbol, s_burnMintERC20.symbol());
    assertEq(decimals, s_burnMintERC20.decimals());
    assertEq(maxSupply, s_burnMintERC20.maxSupply());

    assertEq(s_burnMintERC20.balanceOf(s_alice), 1e18);
    assertEq(s_burnMintERC20.totalSupply(), 1e18);
  }

  // Reverts

  function test_Constructor_RevertWhen_PreMintIsGreaterThanMaxSupply() public {
    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    uint256 maxSupply = 1e33;
    uint256 preMint = maxSupply + 1;

    vm.expectRevert(abi.encodeWithSelector(FactoryBurnMintERC20.MaxSupplyExceeded.selector, preMint));

    new FactoryBurnMintERC20(name, symbol, decimals, maxSupply, preMint, s_alice);
  }
}
