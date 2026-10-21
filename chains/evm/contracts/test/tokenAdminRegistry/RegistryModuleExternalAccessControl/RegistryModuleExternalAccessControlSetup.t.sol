// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {MockExternalAccessControlRegistry} from "../../mocks/MockExternalAccessControlRegistry.sol";
import {MockExternalAccessControlledToken} from "../../mocks/MockExternalAccessControlledToken.sol";

contract RegistryModuleExternalAccessControlSetup is BaseTest {
  error AuthorityUnavailable();

  RegistryModuleExternalAccessControl internal s_registryModuleExternalAccessControl;
  TokenAdminRegistry internal s_tokenAdminRegistry;
  MockExternalAccessControlRegistry internal s_accessControlRegistry;
  address internal s_token;

  function setUp() public virtual override {
    super.setUp();

    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_accessControlRegistry = new MockExternalAccessControlRegistry(OWNER);
    s_token = address(new MockExternalAccessControlledToken(address(s_accessControlRegistry)));
    s_registryModuleExternalAccessControl = new RegistryModuleExternalAccessControl(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleExternalAccessControl));
  }

  function _deployToken(
    address owner,
    address ccipAdmin
  ) internal returns (CrossChainToken) {
    return new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "Test",
        symbol: "TST",
        decimals: 18,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: address(0),
        ccipAdmin: ccipAdmin
      }),
      owner,
      owner
    );
  }

  function _assertPendingAdministrators(
    address[] memory tokens,
    address expectedAdmin
  ) internal view {
    for (uint256 i = 0; i < tokens.length; ++i) {
      assertEq(expectedAdmin, s_tokenAdminRegistry.getTokenConfig(tokens[i]).pendingAdministrator);
      assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(tokens[i]).administrator);
    }
  }

  function _expectRegistrations(
    address[] memory tokens
  ) internal {
    for (uint256 i = 0; i < tokens.length; ++i) {
      vm.expectEmit(address(s_registryModuleExternalAccessControl));
      emit RegistryModuleExternalAccessControl.AdministratorRegistered(tokens[i], OWNER);
    }
  }
}
