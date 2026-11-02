// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {
  FactoryTokenWithUIDPoolDataProvider
} from "../../../../pools/MultiTokenPool/FactoryTokenWithUIDPoolDataProvider.sol";
import {BaseERC20} from "../../../../tokens/BaseERC20.sol";
import {MockTokenWithUID} from "../../../mocks/MockTokenWithUIDFactory.sol";
import {FactoryTokenWithUIDPoolDataProviderSetup} from "./FactoryTokenWithUIDPoolDataProviderSetup.t.sol";

contract FactoryTokenWithUIDPoolDataProvider_getAdditionalPoolData is FactoryTokenWithUIDPoolDataProviderSetup {
  function test_getAdditionalPoolData() public view {
    assertEq(abi.encode(UID), s_provider.getAdditionalPoolData(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, ""));
    assertEq(
      abi.encode(OTHER_UID), s_provider.getAdditionalPoolData(s_otherToken, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "")
    );
  }

  /// @dev The UID belongs to the token, so the finality and the token args do not change it.
  function testFuzz_getAdditionalPoolData_IgnoresFinalityAndTokenArgs(
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs
  ) public view {
    assertEq(
      abi.encode(UID),
      s_provider.getAdditionalPoolData(s_token, DEST_CHAIN_SELECTOR, requestedFinalityConfig, tokenArgs)
    );
  }

  // Reverts

  function test_getAdditionalPoolData_RevertWhen_TokenNotFromFactory() public {
    address token = address(new MockTokenWithUID(UNKNOWN_UID));

    vm.expectRevert(
      abi.encodeWithSelector(FactoryTokenWithUIDPoolDataProvider.TokenNotFromFactory.selector, token, UNKNOWN_UID)
    );
    s_provider.getAdditionalPoolData(token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  /// @dev The token has the UID of s_token, but the factory returns s_token for that UID.
  function test_getAdditionalPoolData_RevertWhen_TokenNotFromFactory_UIDOfOtherToken() public {
    address token = address(new MockTokenWithUID(UID));

    vm.expectRevert(
      abi.encodeWithSelector(FactoryTokenWithUIDPoolDataProvider.TokenNotFromFactory.selector, token, UID)
    );
    s_provider.getAdditionalPoolData(token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  /// @dev A plain ERC20 has no uid function, so the call to uid() reverts without a reason.
  function test_getAdditionalPoolData_RevertWhen_TokenWithoutUID() public {
    BaseERC20 token = new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Token without UID",
        symbol: "NOUID",
        decimals: DEFAULT_TOKEN_DECIMALS,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: address(0),
        ccipAdmin: OWNER
      })
    );

    vm.expectRevert(bytes(""));
    s_provider.getAdditionalPoolData(address(token), DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }
}
