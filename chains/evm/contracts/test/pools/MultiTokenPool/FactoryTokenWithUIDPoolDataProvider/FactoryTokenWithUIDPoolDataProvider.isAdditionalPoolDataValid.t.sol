// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {
  FactoryTokenWithUIDPoolDataProvider
} from "../../../../pools/MultiTokenPool/FactoryTokenWithUIDPoolDataProvider.sol";
import {MockTokenWithUIDFactory} from "../../../mocks/MockTokenWithUIDFactory.sol";
import {FactoryTokenWithUIDPoolDataProviderSetup} from "./FactoryTokenWithUIDPoolDataProviderSetup.t.sol";

contract FactoryTokenWithUIDPoolDataProvider_isAdditionalPoolDataValid is FactoryTokenWithUIDPoolDataProviderSetup {
  function test_isAdditionalPoolDataValid() public view {
    assertTrue(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, abi.encode(UID)));
    assertTrue(s_provider.isAdditionalPoolDataValid(s_otherToken, SOURCE_CHAIN_SELECTOR, abi.encode(OTHER_UID)));
  }

  /// @dev The source chain has its own factory, so the token with the same UID has another address there.
  function test_isAdditionalPoolDataValid_DataFromSourceChain() public {
    MockTokenWithUIDFactory sourceTokenFactory = new MockTokenWithUIDFactory();
    address sourceToken = sourceTokenFactory.deploy(UID);
    FactoryTokenWithUIDPoolDataProvider sourceProvider = new FactoryTokenWithUIDPoolDataProvider(sourceTokenFactory);
    bytes memory additionalPoolData = sourceProvider.getAdditionalPoolData(sourceToken, DEST_CHAIN_SELECTOR);

    assertNotEq(sourceToken, s_token);
    assertTrue(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, additionalPoolData));
    assertFalse(s_provider.isAdditionalPoolDataValid(s_otherToken, SOURCE_CHAIN_SELECTOR, additionalPoolData));
  }

  function test_isAdditionalPoolDataValid_UIDOfOtherToken() public view {
    assertFalse(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, abi.encode(OTHER_UID)));
    assertFalse(s_provider.isAdditionalPoolDataValid(s_otherToken, SOURCE_CHAIN_SELECTOR, abi.encode(UID)));
  }

  function test_isAdditionalPoolDataValid_UIDWithoutToken() public view {
    assertFalse(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, abi.encode(UNKNOWN_UID)));
    // The factory returns zero for a UID without a token. A zero local token must not match it.
    assertFalse(s_provider.isAdditionalPoolDataValid(address(0), SOURCE_CHAIN_SELECTOR, abi.encode(UNKNOWN_UID)));
  }

  /// @dev A source pool that does not send the UID gives empty data.
  function test_isAdditionalPoolDataValid_EmptyData() public view {
    assertFalse(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, ""));
  }

  /// @dev Each case starts with the UID of s_token, so only the length makes it invalid.
  function test_isAdditionalPoolDataValid_WrongLength() public view {
    bytes[] memory additionalPoolData = new bytes[](4);
    additionalPoolData[0] = abi.encodePacked(bytes31(UID));
    additionalPoolData[1] = abi.encodePacked(UID, bytes1(0));
    additionalPoolData[2] = abi.encode(UID, UID);
    additionalPoolData[3] = abi.encodePacked(UID, abi.encode(UID));

    for (uint256 i = 0; i < additionalPoolData.length; ++i) {
      assertFalse(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, additionalPoolData[i]));
    }
  }

  function testFuzz_isAdditionalPoolDataValid_WrongLength(
    bytes calldata additionalPoolData
  ) public view {
    vm.assume(additionalPoolData.length != 32);

    assertFalse(s_provider.isAdditionalPoolDataValid(s_token, SOURCE_CHAIN_SELECTOR, additionalPoolData));
  }
}
