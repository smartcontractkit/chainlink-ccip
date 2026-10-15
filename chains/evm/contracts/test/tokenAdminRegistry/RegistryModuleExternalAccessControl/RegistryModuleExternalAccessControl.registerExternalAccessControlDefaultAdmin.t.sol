// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExternalAccessControlRegistry} from "../../../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MockExternalAccessControlRegistry} from "../../mocks/MockExternalAccessControlRegistry.sol";
import {MockExternalAccessControlledToken} from "../../mocks/MockExternalAccessControlledToken.sol";
import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

contract RegistryModuleExternalAccessControl_registerExternalAccessControlDefaultAdmin is
  RegistryModuleExternalAccessControlSetup
{
  address[] internal s_tokens;

  function setUp() public override {
    super.setUp();

    s_tokens.push(s_token);
    s_tokens.push(address(new MockExternalAccessControlledToken(address(s_accessControlsRegistry))));
    s_tokens.push(address(new MockExternalAccessControlledToken(address(s_accessControlsRegistry))));
  }

  function test_registerExternalAccessControlDefaultAdmin() public {
    assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(s_token).pendingAdministrator);

    bytes32 defaultAdminRole = s_accessControlsRegistry.DEFAULT_ADMIN_ROLE();

    vm.expectCall(s_token, abi.encodeCall(IExternalAccessControlRegistry.ACCESS_CONTROLLED_REGISTRY, ()), 1);
    vm.expectCall(
      address(s_accessControlsRegistry), abi.encodeCall(AccessControl.hasRole, (defaultAdminRole, OWNER)), 1
    );
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_token, OWNER)), 1
    );

    vm.expectEmit();
    emit RegistryModuleExternalAccessControl.AdministratorRegistered(s_token, OWNER);

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(_arrayOf(s_token));

    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_token).pendingAdministrator);
  }

  function test_registerExternalAccessControlDefaultAdmin_MultipleTokens() public {
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      vm.expectCall(
        address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_tokens[i], OWNER)), 1
      );
      vm.expectEmit();
      emit RegistryModuleExternalAccessControl.AdministratorRegistered(s_tokens[i], OWNER);
    }

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
  }

  function test_registerExternalAccessControlDefaultAdmin_EmptyTokens() public {
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeWithSelector(TokenAdminRegistry.proposeAdministrator.selector), 0
    );

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(new address[](0));
  }

  function test_registerExternalAccessControlDefaultAdmin_DuplicateTokens() public {
    address[] memory tokens = new address[](2);
    tokens[0] = s_token;
    tokens[1] = s_token;

    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_token, OWNER)), 2
    );
    vm.expectEmit();
    emit RegistryModuleExternalAccessControl.AdministratorRegistered(s_token, OWNER);
    vm.expectEmit();
    emit RegistryModuleExternalAccessControl.AdministratorRegistered(s_token, OWNER);

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(tokens);

    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_token).pendingAdministrator);
  }

  function test_registerExternalAccessControlDefaultAdmin_ReRegisterAfterRoleRevoked() public {
    address newAdmin = makeAddr("newAdmin");
    bytes32 defaultAdminRole = s_accessControlsRegistry.DEFAULT_ADMIN_ROLE();

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    // The default admin role moves to a new admin before the pending admin has accepted.
    s_accessControlsRegistry.grantRole(defaultAdminRole, newAdmin);
    s_accessControlsRegistry.revokeRole(defaultAdminRole, OWNER);

    vm.startPrank(newAdmin);
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(newAdmin, s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
  }

  function test_registerExternalAccessControlDefaultAdmin_DifferentRegistries() public {
    address otherRegistry = address(new MockExternalAccessControlRegistry(OWNER));
    s_tokens.push(address(new MockExternalAccessControlledToken(otherRegistry)));
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      s_tokenAdminRegistry.acceptAdminRole(s_tokens[i]);
      assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).administrator);
    }
  }

  function testFuzz_registerExternalAccessControlDefaultAdmin_BatchSize(
    uint8 count
  ) public {
    count = uint8(bound(count, 1, 16));
    address[] memory tokens = new address[](count);
    for (uint256 i = 0; i < count; ++i) {
      tokens[i] = address(new MockExternalAccessControlledToken(address(s_accessControlsRegistry)));
    }
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
  }

  // Reverts

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_RequiredRoleNotFound() public {
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    // An address without the default admin role cannot replace the pending admin.
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        STRANGER,
        s_accessControlsRegistry.DEFAULT_ADMIN_ROLE(),
        s_token
      )
    );

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_token).pendingAdministrator);
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_RequiredRoleNotFound_OwnerNotDefaultAdmin()
    public
  {
    // The last token uses an AccessControl contract where OWNER is not the default admin.
    MockExternalAccessControlRegistry otherAccessControlsRegistry =
      new MockExternalAccessControlRegistry(makeAddr("otherAdmin"));
    s_tokens.push(address(new MockExternalAccessControlledToken(address(otherAccessControlsRegistry))));

    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        OWNER,
        otherAccessControlsRegistry.DEFAULT_ADMIN_ROLE(),
        s_tokens[s_tokens.length - 1]
      )
    );

    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);

    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_AlreadyRegistered() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(_arrayOf(lastToken));
    s_tokenAdminRegistry.acceptAdminRole(lastToken);
    vm.expectRevert(abi.encodeWithSelector(TokenAdminRegistry.AlreadyRegistered.selector, lastToken));
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(lastToken).administrator);
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_OnlyRegistryModuleOrOwner() public {
    s_tokenAdminRegistry.removeRegistryModule(address(s_registryModuleExternalAccessControl));
    vm.expectRevert(
      abi.encodeWithSelector(
        TokenAdminRegistry.OnlyRegistryModuleOrOwner.selector, address(s_registryModuleExternalAccessControl)
      )
    );
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_TokenCallReverts() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    vm.mockCallRevert(
      lastToken,
      abi.encodeCall(IExternalAccessControlRegistry.ACCESS_CONTROLLED_REGISTRY, ()),
      abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_RegistryCallReverts() public {
    address registry = address(new MockExternalAccessControlRegistry(OWNER));
    s_tokens.push(address(new MockExternalAccessControlledToken(registry)));
    vm.mockCallRevert(
      registry,
      abi.encodeCall(AccessControl.hasRole, (bytes32(0), OWNER)),
      abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_RegistryHasNoCode() public {
    s_tokens.push(address(new MockExternalAccessControlledToken(address(0))));
    vm.expectRevert();
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlDefaultAdmin_RevertWhen_TokenHasNoCode() public {
    s_tokens.push(address(0));
    vm.expectRevert();
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }
}
