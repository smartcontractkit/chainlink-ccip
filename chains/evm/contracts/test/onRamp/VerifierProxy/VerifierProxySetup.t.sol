// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VerifierProxy} from "../../../onRamp/VerifierProxy.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

contract VerifierProxySetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  bytes32 internal s_metadataHash;

  VerifierProxy internal s_verifierProxy;
  MockVerifier internal s_mockVerifierOne;

  function setUp() public virtual override {
    super.setUp();

    s_verifierProxy = new VerifierProxy(
      VerifierProxy.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      VerifierProxy.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
    s_mockVerifierOne = new MockVerifier();
  }
}
