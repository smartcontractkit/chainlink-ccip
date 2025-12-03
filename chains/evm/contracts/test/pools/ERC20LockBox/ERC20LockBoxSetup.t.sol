// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";

import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract ERC20LockBoxSetup is BaseTest {
  IERC20 internal s_token;
  ERC20LockBox internal s_erc20LockBox;
  TokenAdminRegistry internal s_tokenAdminRegistry;
  address internal s_allowedCaller = makeAddr("allowed_caller");
  address internal s_recipient = makeAddr("recipient");
  address internal s_tokenPool = makeAddr("tokenPool");

  function setUp() public virtual override {
    super.setUp();
    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);
    deal(address(s_token), s_allowedCaller, type(uint256).max);

    // Deploy and configure the token admin registry for the token
    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_tokenAdminRegistry.proposeAdministrator(address(s_token), address(OWNER));
    s_tokenAdminRegistry.acceptAdminRole(address(s_token));

    vm.mockCall(s_tokenPool, abi.encodeWithSignature("isSupportedToken(address)", address(s_token)), abi.encode(true));

    // Set the token pool for the token
    s_tokenAdminRegistry.setPool(address(s_token), s_tokenPool);

    // Mock the owner of the token pool to be the owner of the token so that we can test the allowed caller configuration
    vm.mockCall(s_tokenPool, abi.encodeWithSignature("owner()"), abi.encode(OWNER));

    // Deploy the ERC20 lock box
    s_erc20LockBox = new ERC20LockBox(address(s_tokenAdminRegistry));

    // Configure the allowed caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] =
      ERC20LockBox.AllowedCallerConfigArgs({token: address(s_token), caller: s_allowedCaller, allowed: true});
    s_erc20LockBox.configureAllowedCallers(configArgs);
  }

  function _depositTokens(
    uint256 amount
  ) internal {
    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);
    s_erc20LockBox.deposit(address(s_token), amount);
    vm.stopPrank();
  }
}
