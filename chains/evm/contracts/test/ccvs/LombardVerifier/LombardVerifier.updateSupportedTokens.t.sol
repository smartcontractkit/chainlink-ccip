// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {MockLombardAdapter} from "../../mocks/MockLombardAdapter.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract LombardVerifier_updateSupportedTokens is LombardVerifierSetup {
  function test_updateSupportedTokens_AddToken() public {
    address newToken = address(new BurnMintERC20("New Token", "NEW", 18, 0, 0));
    address localAdapter = address(0);

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: newToken, localAdapter: localAdapter});

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenSet(newToken, localAdapter);

    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertTrue(s_lombardVerifier.isSupportedToken(newToken), "Token should be in supported tokens");
    // Check the token if approval is set to max uint256.
    uint256 allowance =
      BurnMintERC20(newToken).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(allowance, type(uint256).max, "Allowance should be max uint256");
  }

  function test_updateSupportedTokens_AddTokenWithAdapter() public {
    uint256 countBefore = s_lombardVerifier.getSupportedTokens().length;
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapter = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenSet(address(newToken), address(adapter));

    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertTrue(s_lombardVerifier.isSupportedToken(address(newToken)), "Token should be in supported tokens");
    // Check token->adapter approval is set to max uint256.
    uint256 allowance = newToken.allowance(address(s_lombardVerifier), address(adapter));
    assertEq(allowance, type(uint256).max, "Allowance should be max uint256");

    uint256 countAfter = s_lombardVerifier.getSupportedTokens().length;
    assertEq(countAfter, countBefore + 1, "Supported tokens count should have increased by 1");
  }

  function test_updateSupportedTokens_RemoveToken() public {
    // First add a token.
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(0)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Now remove it.
    address[] memory tokensToRemove = new address[](1);
    tokensToRemove[0] = address(newToken);

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenRemoved(address(newToken));

    s_lombardVerifier.updateSupportedTokens(tokensToRemove, new LombardVerifier.SupportedTokenArgs[](0));

    assertFalse(s_lombardVerifier.isSupportedToken(address(newToken)), "Token should not be in supported tokens");

    // Verify the token's allowance was reset to 0.
    uint256 allowance =
      BurnMintERC20(newToken).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(allowance, 0, "Token allowance should be reset to 0");
  }

  function test_updateSupportedTokens_RemoveTokenWithAdapter() public {
    // First add a token with an adapter.
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapter = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Verify token->adapter allowance was set to max.
    uint256 allowanceBefore = newToken.allowance(address(s_lombardVerifier), address(adapter));
    assertEq(allowanceBefore, type(uint256).max, "Allowance should be max uint256 before removal");

    // Now remove the token.
    address[] memory tokensToRemove = new address[](1);
    tokensToRemove[0] = address(newToken);

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenRemoved(address(newToken));

    s_lombardVerifier.updateSupportedTokens(tokensToRemove, new LombardVerifier.SupportedTokenArgs[](0));

    assertFalse(s_lombardVerifier.isSupportedToken(address(newToken)), "Token should not be in supported tokens");

    // Verify token->adapter allowance was reset to 0.
    uint256 allowanceAfter = newToken.allowance(address(s_lombardVerifier), address(adapter));
    assertEq(allowanceAfter, 0, "Allowance should be reset to 0 after removal");
  }

  function test_updateSupportedTokens_RotateAdapter_RevokesOldAdapterAllowance() public {
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapterA = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));
    MockLombardAdapter adapterB = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    // Add token with adapter A.
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] =
      LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapterA)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(newToken.allowance(address(s_lombardVerifier), address(adapterA)), type(uint256).max);

    // Rotate to adapter B.
    tokensToAdd[0] =
      LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapterB)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(newToken.allowance(address(s_lombardVerifier), address(adapterA)), 0, "Old adapter allowance should be 0");
    assertEq(
      newToken.allowance(address(s_lombardVerifier), address(adapterB)),
      type(uint256).max,
      "New adapter allowance should be max"
    );
  }

  function test_updateSupportedTokens_SwitchFromNoAdapterToAdapter_RevokesBridgeAllowance() public {
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapter = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    // Add token without adapter (approved to bridge).
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(0)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(newToken.allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge())), type(uint256).max);

    // Switch to adapter.
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(
      newToken.allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge())),
      0,
      "Bridge allowance should be 0"
    );
    assertEq(
      newToken.allowance(address(s_lombardVerifier), address(adapter)),
      type(uint256).max,
      "Adapter allowance should be max"
    );
  }

  function test_updateSupportedTokens_SwitchFromAdapterToNoAdapter_RevokesAdapterAllowance() public {
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapter = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    // Add token with adapter.
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(newToken.allowance(address(s_lombardVerifier), address(adapter)), type(uint256).max);

    // Switch to no adapter (approved to bridge).
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(0)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(newToken.allowance(address(s_lombardVerifier), address(adapter)), 0, "Adapter allowance should be 0");
    assertEq(
      newToken.allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge())),
      type(uint256).max,
      "Bridge allowance should be max"
    );
  }

  function test_updateSupportedTokens_SameAdapter_NoRedundantRevoke() public {
    BurnMintERC20 newToken = new BurnMintERC20("New Token", "NEW", 18, 0, 0);
    MockLombardAdapter adapter = new MockLombardAdapter(address(s_lombardVerifier.i_bridge()), address(newToken));

    // Add token with adapter.
    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Re-set the same adapter. Allowance should remain max (no revoke).
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertEq(
      newToken.allowance(address(s_lombardVerifier), address(adapter)),
      type(uint256).max,
      "Allowance should remain max when adapter is unchanged"
    );
  }

  function test_updateSupportedTokens_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.updateSupportedTokens(new address[](0), new LombardVerifier.SupportedTokenArgs[](0));
  }
}
