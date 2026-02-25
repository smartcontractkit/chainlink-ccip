// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";

contract LombardVerifier_updateSupportedTokens is LombardVerifierSetup {
  function test_updateSupportedTokens_AddToken() public {
    address newToken = address(
      new CrossChainToken(
        BaseERC20.ConstructorParams({
          name: "New Token", symbol: "NEW", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
        }),
        OWNER,
        OWNER
      )
    );
    address localAdapter = address(0);

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: newToken, localAdapter: localAdapter});

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenSet(newToken, localAdapter);

    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertTrue(s_lombardVerifier.isSupportedToken(newToken), "Token should be in supported tokens");
    // Check the token if approval is set to max uint256.
    uint256 allowance =
      CrossChainToken(newToken).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(allowance, type(uint256).max, "Allowance should be max uint256");
  }

  function test_updateSupportedTokens_AddTokenWithAdapter() public {
    uint256 countBefore = s_lombardVerifier.getSupportedTokens().length;
    CrossChainToken newToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "New Token", symbol: "NEW", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
    // The adapter must be a valid ERC20 since the contract calls approve on it.
    CrossChainToken adapter = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "Adapter Token", symbol: "ADAPT", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenSet(address(newToken), address(adapter));

    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    assertTrue(s_lombardVerifier.isSupportedToken(address(newToken)), "Token should be in supported tokens");
    // Check the adapter if approval is set to max uint256.
    uint256 allowance =
      CrossChainToken(adapter).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(allowance, type(uint256).max, "Allowance should be max uint256");

    uint256 countAfter = s_lombardVerifier.getSupportedTokens().length;
    assertEq(countAfter, countBefore + 1, "Supported tokens count should have increased by 1");
  }

  function test_updateSupportedTokens_RemoveToken() public {
    // First add a token.
    CrossChainToken newToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "New Token", symbol: "NEW", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
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
      CrossChainToken(newToken).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(allowance, 0, "Token allowance should be reset to 0");
  }

  function test_updateSupportedTokens_RemoveTokenWithAdapter() public {
    // First add a token with an adapter.
    CrossChainToken newToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "New Token", symbol: "NEW", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
    CrossChainToken adapter = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "Adapter Token", symbol: "ADAPT", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: address(newToken), localAdapter: address(adapter)});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    // Verify adapter allowance was set to max.
    uint256 adapterAllowanceBefore =
      CrossChainToken(adapter).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(adapterAllowanceBefore, type(uint256).max, "Adapter allowance should be max uint256 before removal");

    // Now remove the token.
    address[] memory tokensToRemove = new address[](1);
    tokensToRemove[0] = address(newToken);

    vm.expectEmit();
    emit LombardVerifier.SupportedTokenRemoved(address(newToken));

    s_lombardVerifier.updateSupportedTokens(tokensToRemove, new LombardVerifier.SupportedTokenArgs[](0));

    assertFalse(s_lombardVerifier.isSupportedToken(address(newToken)), "Token should not be in supported tokens");

    // Verify the adapter's allowance was reset to 0 (not the token's).
    uint256 adapterAllowanceAfter =
      CrossChainToken(adapter).allowance(address(s_lombardVerifier), address(s_lombardVerifier.i_bridge()));
    assertEq(adapterAllowanceAfter, 0, "Adapter allowance should be reset to 0 after removal");
  }

  function test_updateSupportedTokens_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.updateSupportedTokens(new address[](0), new LombardVerifier.SupportedTokenArgs[](0));
  }
}
