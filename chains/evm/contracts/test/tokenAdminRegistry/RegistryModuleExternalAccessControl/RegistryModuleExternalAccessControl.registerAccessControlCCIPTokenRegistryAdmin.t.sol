// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

contract RegistryModuleExternalAccessControl_registerAccessControlCCIPTokenRegistryAdmin is
  RegistryModuleExternalAccessControlSetup
{
  address internal s_defaultAdmin;
  bytes32 internal s_tokenRegistryAdminRole;
  address[] internal s_tokens;

  function setUp() public override {
    super.setUp();
    s_defaultAdmin = makeAddr("defaultAdmin");
    s_tokenRegistryAdminRole = s_registryModuleExternalAccessControl.CCIP_TOKEN_REGISTRY_ADMIN_ROLE();

    // OWNER has CCIP_TOKEN_REGISTRY_ADMIN_ROLE but not DEFAULT_ADMIN_ROLE to test the role independently.
    vm.startPrank(s_defaultAdmin);
    for (uint256 i = 0; i < 3; ++i) {
      CrossChainToken token = _deployToken(s_defaultAdmin, STRANGER);
      token.grantRole(s_tokenRegistryAdminRole, OWNER);
      s_tokens.push(address(token));
    }
    vm.startPrank(OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin() public {
    vm.expectCall(
      s_tokens[0], abi.encodeCall(AccessControl.hasRole, (keccak256("CCIP_TOKEN_REGISTRY_ADMIN_ROLE"), OWNER)), 1
    );
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_tokens[0], OWNER)), 1
    );
    address[] memory tokens = _arrayOf(s_tokens[0]);
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(tokens[0]);
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(tokens[0]).administrator);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_MultipleTokens() public {
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_UsesDefaultAdminRoleFromContract() public {
    address token = s_tokens[0];
    bytes32 defaultAdminRole = keccak256("CUSTOM_DEFAULT_ADMIN_ROLE");
    vm.mockCall(token, abi.encodeCall(AccessControl(token).DEFAULT_ADMIN_ROLE, ()), abi.encode(defaultAdminRole));
    vm.mockCall(
      token, abi.encodeCall(AccessControl.getRoleAdmin, (s_tokenRegistryAdminRole)), abi.encode(defaultAdminRole)
    );

    address[] memory tokens = _arrayOf(token);
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_EmptyTokens() public {
    vm.recordLogs();
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(new address[](0));
    assertEq(0, vm.getRecordedLogs().length);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_DuplicateTokens() public {
    address[] memory tokens = new address[](2);
    tokens[0] = s_tokens[0];
    tokens[1] = s_tokens[0];
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (tokens[0], OWNER)), 2
    );
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(tokens);
    _assertPendingAdministrators(tokens, OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_ReplacesDefaultAdminRegistration() public {
    vm.stopPrank();
    vm.startPrank(s_defaultAdmin);
    s_registryModuleExternalAccessControl.registerAccessControlDefaultAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, s_defaultAdmin);

    vm.startPrank(OWNER);
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_ReRegisterAfterRoleRevoked() public {
    address newAdmin = makeAddr("newAdmin");
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    vm.startPrank(s_defaultAdmin);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      CrossChainToken(s_tokens[i]).grantRole(s_tokenRegistryAdminRole, newAdmin);
      CrossChainToken(s_tokens[i]).revokeRole(s_tokenRegistryAdminRole, OWNER);
    }
    vm.startPrank(newAdmin);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);

    vm.startPrank(OWNER);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector, OWNER, s_tokenRegistryAdminRole, s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_AcceptAfterRoleRevoked() public {
    address token = s_tokens[0];
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(_arrayOf(token));
    vm.startPrank(s_defaultAdmin);
    CrossChainToken(token).revokeRole(s_tokenRegistryAdminRole, OWNER);

    // TokenAdminRegistry does not check the role again when the pending admin accepts.
    vm.startPrank(OWNER);
    s_tokenAdminRegistry.acceptAdminRole(token);
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(token).administrator);
  }

  // Reverts

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_InvalidRoleAdmin() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    bytes32[2] memory roleAdmins = [keccak256("OTHER_ADMIN_ROLE"), s_tokenRegistryAdminRole];
    for (uint256 i = 0; i < roleAdmins.length; ++i) {
      vm.mockCall(
        lastToken, abi.encodeCall(AccessControl.getRoleAdmin, (s_tokenRegistryAdminRole)), abi.encode(roleAdmins[i])
      );
      vm.expectRevert(
        abi.encodeWithSelector(
          RegistryModuleExternalAccessControl.InvalidRoleAdmin.selector,
          s_tokenRegistryAdminRole,
          roleAdmins[i],
          lastToken
        )
      );
      s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
      _assertPendingAdministrators(s_tokens, address(0));
    }
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_RoleAdminCallReverts() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    vm.mockCallRevert(
      lastToken,
      abi.encodeCall(AccessControl.getRoleAdmin, (s_tokenRegistryAdminRole)),
      abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_DefaultAdminRoleCallReverts() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    vm.mockCallRevert(
      lastToken,
      abi.encodeCall(AccessControl(lastToken).DEFAULT_ADMIN_ROLE, ()),
      abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_UnauthorizedCaller()
    public
  {
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    vm.startPrank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        STRANGER,
        s_tokenRegistryAdminRole,
        s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_DefaultAdmin() public {
    vm.stopPrank();
    vm.startPrank(s_defaultAdmin);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector,
        s_defaultAdmin,
        s_tokenRegistryAdminRole,
        s_tokens[0]
      )
    );
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_RequiredRoleNotFound_UnauthorizedLastToken()
    public
  {
    address otherToken = address(_deployToken(STRANGER, STRANGER));
    s_tokens.push(otherToken);
    vm.expectRevert(
      abi.encodeWithSelector(
        RegistryModuleExternalAccessControl.RequiredRoleNotFound.selector, OWNER, s_tokenRegistryAdminRole, otherToken
      )
    );
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_AlreadyRegistered() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(_arrayOf(lastToken));
    s_tokenAdminRegistry.acceptAdminRole(lastToken);
    vm.expectRevert(abi.encodeWithSelector(TokenAdminRegistry.AlreadyRegistered.selector, lastToken));
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(lastToken).administrator);
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_OnlyRegistryModuleOrOwner() public {
    s_tokenAdminRegistry.removeRegistryModule(address(s_registryModuleExternalAccessControl));
    vm.expectRevert(
      abi.encodeWithSelector(
        TokenAdminRegistry.OnlyRegistryModuleOrOwner.selector, address(s_registryModuleExternalAccessControl)
      )
    );
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_AuthorityCallReverts() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    vm.mockCallRevert(
      lastToken,
      abi.encodeCall(AccessControl.hasRole, (s_tokenRegistryAdminRole, OWNER)),
      abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAccessControlCCIPTokenRegistryAdmin_RevertWhen_TokenHasNoCode() public {
    s_tokens.push(address(0));
    vm.expectRevert();
    s_registryModuleExternalAccessControl.registerAccessControlCCIPTokenRegistryAdmin(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }
}
