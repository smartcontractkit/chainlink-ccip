// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwner} from "../../../interfaces/IOwner.sol";

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

contract RegistryModuleExternalAccessControl_registerAdminViaOwner is RegistryModuleExternalAccessControlSetup {
  address[] internal s_tokens;

  function setUp() public override {
    super.setUp();
    for (uint256 i = 0; i < 3; ++i) {
      s_tokens.push(address(_deployToken(OWNER, STRANGER)));
    }
  }

  function test_registerAdminViaOwner() public {
    vm.expectCall(s_tokens[0], abi.encodeCall(IOwner.owner, ()), 1);
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (s_tokens[0], OWNER)), 1
    );
    address[] memory tokens = _arrayOf(s_tokens[0]);
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerAdminViaOwner(tokens);
    _assertPendingAdministrators(tokens, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(tokens[0]);
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(tokens[0]).administrator);
  }

  function test_registerAdminViaOwner_MultipleTokens() public {
    _expectRegistrations(s_tokens);
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerAdminViaOwner_EmptyTokens() public {
    vm.recordLogs();
    s_registryModuleExternalAccessControl.registerAdminViaOwner(new address[](0));
    assertEq(0, vm.getRecordedLogs().length);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAdminViaOwner_DuplicateTokens() public {
    address[] memory tokens = new address[](2);
    tokens[0] = s_tokens[0];
    tokens[1] = s_tokens[0];
    vm.expectCall(
      address(s_tokenAdminRegistry), abi.encodeCall(TokenAdminRegistry.proposeAdministrator, (tokens[0], OWNER)), 2
    );
    _expectRegistrations(tokens);
    s_registryModuleExternalAccessControl.registerAdminViaOwner(tokens);
    _assertPendingAdministrators(tokens, OWNER);
  }

  function test_registerAdminViaOwner_ReRegisterAfterOwnerTransfer() public {
    address newAdmin = makeAddr("newAdmin");
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      CrossChainToken(s_tokens[i]).beginDefaultAdminTransfer(newAdmin);
    }
    vm.warp(block.timestamp + 1);
    vm.startPrank(newAdmin);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      CrossChainToken(s_tokens[i]).acceptDefaultAdminTransfer();
    }
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);

    vm.startPrank(OWNER);
    vm.expectRevert(
      abi.encodeWithSelector(RegistryModuleExternalAccessControl.CanOnlySelfRegister.selector, newAdmin, s_tokens[0])
    );
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, newAdmin);
  }

  // Reverts

  function test_registerAdminViaOwner_RevertWhen_CanOnlySelfRegister_UnauthorizedCaller() public {
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    vm.startPrank(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(RegistryModuleExternalAccessControl.CanOnlySelfRegister.selector, OWNER, s_tokens[0])
    );
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, OWNER);
  }

  function test_registerAdminViaOwner_RevertWhen_CanOnlySelfRegister_UnauthorizedLastToken() public {
    address otherToken = address(_deployToken(STRANGER, STRANGER));
    s_tokens.push(otherToken);
    vm.expectRevert(
      abi.encodeWithSelector(RegistryModuleExternalAccessControl.CanOnlySelfRegister.selector, STRANGER, otherToken)
    );
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAdminViaOwner_RevertWhen_AlreadyRegistered() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    s_registryModuleExternalAccessControl.registerAdminViaOwner(_arrayOf(lastToken));
    s_tokenAdminRegistry.acceptAdminRole(lastToken);
    vm.expectRevert(abi.encodeWithSelector(TokenAdminRegistry.AlreadyRegistered.selector, lastToken));
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    for (uint256 i = 0; i < s_tokens.length; ++i) {
      assertEq(address(0), s_tokenAdminRegistry.getTokenConfig(s_tokens[i]).pendingAdministrator);
    }
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(lastToken).administrator);
  }

  function test_registerAdminViaOwner_RevertWhen_OnlyRegistryModuleOrOwner() public {
    s_tokenAdminRegistry.removeRegistryModule(address(s_registryModuleExternalAccessControl));
    vm.expectRevert(
      abi.encodeWithSelector(
        TokenAdminRegistry.OnlyRegistryModuleOrOwner.selector, address(s_registryModuleExternalAccessControl)
      )
    );
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAdminViaOwner_RevertWhen_AuthorityCallReverts() public {
    address lastToken = s_tokens[s_tokens.length - 1];
    vm.mockCallRevert(
      lastToken, abi.encodeCall(IOwner.owner, ()), abi.encodeWithSelector(AuthorityUnavailable.selector)
    );
    vm.expectRevert(AuthorityUnavailable.selector);
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }

  function test_registerAdminViaOwner_RevertWhen_TokenHasNoCode() public {
    s_tokens.push(address(0));
    vm.expectRevert();
    s_registryModuleExternalAccessControl.registerAdminViaOwner(s_tokens);
    _assertPendingAdministrators(s_tokens, address(0));
  }
}
