// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

import {IGetCCIPAdmin} from "../../../interfaces/IGetCCIPAdmin.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract BaseERC20_constructor is BaseERC20Setup {
  function test_constructor() public view {
    assertEq("Base Token", s_baseERC20.name());
    assertEq("BASE", s_baseERC20.symbol());
    assertEq(DEFAULT_TOKEN_DECIMALS, s_baseERC20.decimals());
    assertEq(MAX_SUPPLY, s_baseERC20.maxSupply());
    assertEq(PRE_MINT, s_baseERC20.totalSupply());
    assertEq(PRE_MINT, s_baseERC20.balanceOf(OWNER));
    assertEq(OWNER, s_baseERC20.getCCIPAdmin());
  }

  function test_constructor_NoPreMint() public {
    uint8 decimals = 6;
    BaseERC20 token = new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "No Mint",
        symbol: "NM",
        decimals: decimals,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: address(0),
        ccipAdmin: OWNER
      })
    );

    assertEq(0, token.totalSupply());
    assertEq(decimals, token.decimals());
    assertEq(0, token.maxSupply());
  }

  function test_constructor_CcipAdminDefaultsToMsgSender() public {
    BaseERC20 token = new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Default Admin",
        symbol: "DA",
        decimals: 18,
        maxSupply: 0,
        preMint: 1000e18,
        preMintRecipient: OWNER,
        ccipAdmin: address(0)
      })
    );

    // ccipAdmin defaults to msg.sender (OWNER) when address(0) is passed; preMintRecipient is set explicitly to OWNER
    assertEq(OWNER, token.getCCIPAdmin());
    assertEq(1000e18, token.balanceOf(OWNER));
  }

  function test_constructor_SupportsInterface() public view {
    assertTrue(s_baseERC20.supportsInterface(type(IERC20).interfaceId));
    assertTrue(s_baseERC20.supportsInterface(type(IGetCCIPAdmin).interfaceId));
  }

  function test_constructor_PreMintRecipient() public {
    address recipient = makeAddr("preMintRecipient");
    uint256 amount = 500e18;

    BaseERC20 token = new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Recipient Token",
        symbol: "RT",
        decimals: 18,
        maxSupply: 0,
        preMint: amount,
        preMintRecipient: recipient,
        ccipAdmin: OWNER
      })
    );

    assertEq(amount, token.balanceOf(recipient));
    assertEq(0, token.balanceOf(OWNER));
    assertEq(amount, token.totalSupply());
  }

  function test_typeAndVersion() public view {
    assertEq("BaseERC20 2.0.0-dev", s_baseERC20.typeAndVersion());
  }

  function test_constructor_RevertWhen_PreMintAddressNotSet() public {
    vm.expectRevert(BaseERC20.PreMintAddressNotSet.selector);
    new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "No Recipient",
        symbol: "NR",
        decimals: 18,
        maxSupply: 0,
        preMint: 1e18,
        preMintRecipient: address(0),
        ccipAdmin: OWNER
      })
    );
  }

  function test_constructor_RevertWhen_PreMintRecipientSetWithZeroPreMint() public {
    address recipient = makeAddr("recipient");

    vm.expectRevert(abi.encodeWithSelector(BaseERC20.PreMintRecipientSetWithZeroPreMint.selector, recipient));
    new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Zero Mint",
        symbol: "ZM",
        decimals: 18,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: recipient,
        ccipAdmin: OWNER
      })
    );
  }

  function test_constructor_RevertWhen_MaxSupplyExceeded() public {
    uint256 maxSupply = 500e18;
    uint256 preMint = maxSupply + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseERC20.MaxSupplyExceeded.selector, preMint, maxSupply));
    new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Over",
        symbol: "OVR",
        decimals: 18,
        maxSupply: maxSupply,
        preMint: preMint,
        preMintRecipient: OWNER,
        ccipAdmin: OWNER
      })
    );
  }
}
