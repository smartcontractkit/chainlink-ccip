// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {
  FactoryTokenWithUIDPoolDataProvider
} from "../../../../pools/MultiTokenPool/FactoryTokenWithUIDPoolDataProvider.sol";
import {BaseTest} from "../../../BaseTest.t.sol";
import {MockTokenWithUIDFactory} from "../../../mocks/MockTokenWithUIDFactory.sol";

contract FactoryTokenWithUIDPoolDataProviderSetup is BaseTest {
  bytes32 internal constant UID = keccak256("uid");
  bytes32 internal constant OTHER_UID = keccak256("otherUid");
  /// @dev A UID that the factory has no token for.
  bytes32 internal constant UNKNOWN_UID = keccak256("unknownUid");
  bytes4 internal constant WAIT_FOR_FINALITY = FinalityCodec.WAIT_FOR_FINALITY_FLAG;

  MockTokenWithUIDFactory internal s_tokenFactory;
  FactoryTokenWithUIDPoolDataProvider internal s_provider;
  address internal s_token;
  address internal s_otherToken;

  function setUp() public virtual override {
    super.setUp();
    s_tokenFactory = new MockTokenWithUIDFactory();
    s_token = s_tokenFactory.deploy(UID);
    s_otherToken = s_tokenFactory.deploy(OTHER_UID);
    s_provider = new FactoryTokenWithUIDPoolDataProvider(s_tokenFactory);
  }
}
