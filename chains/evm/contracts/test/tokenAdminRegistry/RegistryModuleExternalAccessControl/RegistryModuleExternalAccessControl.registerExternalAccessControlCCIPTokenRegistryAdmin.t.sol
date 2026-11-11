// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExternalAccessControlRegistry} from "../../../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MockExternalAccessControlRegistry} from "../../mocks/MockExternalAccessControlRegistry.sol";
import {MockExternalAccessControlledToken} from "../../mocks/MockExternalAccessControlledToken.sol";
import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

contract RegistryModuleExternalAccessControl_registerExternalAccessControlCCIPTokenRegistryAdmin is
  RegistryModuleExternalAccessControlSetup
{
  address internal s_defaultAdmin;
  bytes32 internal s_tokenRegistryAdminRole;
  MockExternalAccessControlRegistry internal s_registry;
  address[] internal s_tokens;

  function setUp() public override {
    super.setUp();
    s_defaultAdmin = makeAddr("defaultAdmin");
    s_tokenRegistryAdminRole = s_registryModuleExternalAccessControl.CCIP_TOKEN_REGISTRY_ADMIN_ROLE();

    // OWNER has CCIP_TOKEN_REGISTRY_ADMIN_ROLE but not DEFAULT_ADMIN_ROLE to test the role independently.
    s_registry = _deployRegistryWithTokenRegistryAdmin(OWNER);
    for (uint256 i = 0; i < 3; ++i) {
      s_tokens.push(address(new MockExternalAccessControlledToken(address(s_registry))));
    }
  }

  /// @notice Deploys an AccessControl contract where s_defaultAdmin has DEFAULT_ADMIN_ROLE and tokenRegistryAdmin has
  /// CCIP_TOKEN_REGISTRY_ADMIN_ROLE. Leaves the prank set to OWNER.
  function _deployRegistryWithTokenRegistryAdmin(
    address tokenRegistryAdmin
  ) internal returns (MockExternalAccessControlRegistry) {
    MockExternalAccessControlRegistry registry = new MockExternalAccessControlRegistry(s_defaultAdmin);
    vm.startPrank(s_defaultAdmin);
    registry.grantRole(s_tokenRegistryAdminRole, tokenRegistryAdmin);
    vm.startPrank(OWNER);
    return registry;
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin() public {
    vm.expectCall(s_tokens[0], abi.encodeCall(IExternalAccessControlRegistry.ACCESS_CONTROLLED_REGISTRY, ()), 1);
    vm.expectCall(
      address(s_registry),
      abi.encodeCall(AccessControl.hasRole, (keccak256("CCIP_TOKEN_REGISTRY_ADMIN_ROLE"), OWNER)),
      1
    );
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_tokens[0], OWNER)), 1
    );
    address[] memory tokens = _arrayOf(s_tokens[0]);
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(tokens[0]);
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(tokens[0]).administrator);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_MultipleTokens() public {
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_UsesDefaultAdminRoleFromContract() public {
    bytes32 defaultAdminRole = keccak256("CUSTOM_DEFAULT_ADMIN_ROLE");
    vm.mockCall(address(s_registry), abi.encodeCall(s_registry.DEFAULT_ADMIN_ROLE, ()), abi.encode(defaultAdminRole));
    vm.mockCall(
      address(s_registry),
      abi.encodeCall(AccessControl.getRoleAdmin, (s_tokenRegistryAdminRole)),
      abi.encode(defaultAdminRole)
    );

    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_EmptyTokens() public {
    vm.recordLogs();
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(new address[](0));
    assertEq(0, vm.getRecordedLogs().length);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_DuplicateTokens() public {
    address[] memory tokens = new address[](2);
    tokens[0] = s_tokens[0];
    tokens[1] = s_tokens[0];
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (tokens[0], OWNER)), 2
    );
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_DifferentRegistries() public {
    address otherRegistry = address(_deployRegistryWithTokenRegistryAdmin(OWNER));
    s_tokens.push(address(new MockExternalAccessControlledToken(otherRegistry)));
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_ReplacesDefaultAdminRegistration() public {
    vm.startPrank(s_defaultAdmin);
    s_registryModuleExternalAccessControl.registerExternalAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, s_defaultAdmin);

    vm.startPrank(OWNER);
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_ReRegisterAfterRoleRevoked() public {
    address newAdmin = makeAddr("newAdmin");
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    vm.startPrank(s_defaultAdmin);
    s_registry.grantRole(s_tokenRegistryAdminRole, newAdmin);
    s_registry.revokeRole(s_tokenRegistryAdminRole, OWNER);
    vm.startPrank(newAdmin);
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);

    vm.startPrank(OWNER);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector, OWNER, s_tokenRegistryAdminRole, s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);
  }

  // Reverts

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_RevertWhen_InvalidRoleAdmin() public {
    address otherRegistry = address(_deployRegistryWithTokenRegistryAdmin(OWNER));
    address lastToken = address(new MockExternalAccessControlledToken(otherRegistry));
    s_tokens.push(lastToken);
    bytes32[2] memory roleAdmins = [keccak256("OTHER_ADMIN_ROLE"), s_tokenRegistryAdminRole];
    for (uint256 i = 0; i < roleAdmins.length; ++i) {
      vm.mockCall(
        otherRegistry, abi.encodeCall(AccessControl.getRoleAdmin, (s_tokenRegistryAdminRole)), abi.encode(roleAdmins[i])
      );
      vm.expectRevert(
        abi.encodeWithSelector(
          RegistryModuleExternalAccessControl.InvalidRoleAdmin.selector,
          s_tokenRegistryAdminRole,
          roleAdmins[i],
          lastToken
        )
      );
      s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
      _assertPendingAdministrators(s_tokens, address(0));
    }
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_UnauthorizedCaller()
    public
  {
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    vm.startPrank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        STRANGER,
        s_tokenRegistryAdminRole,
        s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_DefaultAdmin()
    public
  {
    vm.startPrank(s_defaultAdmin);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        s_defaultAdmin,
        s_tokenRegistryAdminRole,
        s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerExternalAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_UnauthorizedLastToken()
    public
  {
    address otherRegistry = address(_deployRegistryWithTokenRegistryAdmin(STRANGER));
    s_tokens.push(address(new MockExternalAccessControlledToken(otherRegistry)));
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        OWNER,
        s_tokenRegistryAdminRole,
        s_tokens[s_tokens.length - 1]
      )
    );
    s_registryModuleExternalAccessControl.registerExternalAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }
}
