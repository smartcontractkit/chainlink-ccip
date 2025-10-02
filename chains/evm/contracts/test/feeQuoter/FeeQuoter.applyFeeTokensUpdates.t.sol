// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FeeQuoter_applyFeeTokensUpdates is FeeQuoterSetup {
  function testFuzz_applyFeeTokensUpdates_SetPremiumMultiplier(
    FeeQuoter.PremiumMultiplierWeiPerEthArgs memory feeTokenAdds
  ) public {
    FeeQuoter.PremiumMultiplierWeiPerEthArgs[] memory premiumMultiplierWeiPerEthArgs =
      new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](1);
    premiumMultiplierWeiPerEthArgs[0] = feeTokenAdds;

    vm.expectEmit();
    emit FeeQuoter.PremiumMultiplierWeiPerEthUpdated(feeTokenAdds.token, feeTokenAdds.premiumMultiplierWeiPerEth);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), premiumMultiplierWeiPerEthArgs);

    assertEq(feeTokenAdds.premiumMultiplierWeiPerEth, s_feeQuoter.getPremiumMultiplierWeiPerEth(feeTokenAdds.token));
  }

  function test_applyFeeTokensUpdates_singleToken() public {
    FeeQuoter.PremiumMultiplierWeiPerEthArgs[] memory feeTokenAdds = new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](1);
    feeTokenAdds[0] = s_feeQuoterPremiumMultiplierWeiPerEthArgs[0];
    feeTokenAdds[0].token = vm.addr(1);

    vm.expectEmit();
    emit FeeQuoter.PremiumMultiplierWeiPerEthUpdated(vm.addr(1), feeTokenAdds[0].premiumMultiplierWeiPerEth);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokenAdds);

    assertEq(
      s_feeQuoterPremiumMultiplierWeiPerEthArgs[0].premiumMultiplierWeiPerEth,
      s_feeQuoter.getPremiumMultiplierWeiPerEth(vm.addr(1))
    );
  }

  function test_applyFeeTokensUpdates_multipleTokens() public {
    FeeQuoter.PremiumMultiplierWeiPerEthArgs[] memory feeTokenAdds = new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](2);
    feeTokenAdds[0] = s_feeQuoterPremiumMultiplierWeiPerEthArgs[0];
    feeTokenAdds[0].token = vm.addr(1);
    feeTokenAdds[1].token = vm.addr(2);

    vm.expectEmit();
    emit FeeQuoter.PremiumMultiplierWeiPerEthUpdated(vm.addr(1), feeTokenAdds[0].premiumMultiplierWeiPerEth);
    vm.expectEmit();
    emit FeeQuoter.PremiumMultiplierWeiPerEthUpdated(vm.addr(2), feeTokenAdds[1].premiumMultiplierWeiPerEth);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokenAdds);

    assertEq(feeTokenAdds[0].premiumMultiplierWeiPerEth, s_feeQuoter.getPremiumMultiplierWeiPerEth(vm.addr(1)));
    assertEq(feeTokenAdds[1].premiumMultiplierWeiPerEth, s_feeQuoter.getPremiumMultiplierWeiPerEth(vm.addr(2)));
  }

  function test_applyFeeTokensUpdates() public {
    FeeQuoter.PremiumMultiplierWeiPerEthArgs[] memory feeTokens = new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](1);
    feeTokens[0].token = s_sourceTokens[1];
    feeTokens[0].premiumMultiplierWeiPerEth = 1e18;

    address[] memory feeTokenAddresses = new address[](1);
    feeTokenAddresses[0] = feeTokens[0].token;

    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(feeTokens[0].token, feeTokens[0].premiumMultiplierWeiPerEth);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokens);
    assertEq(s_feeQuoter.getFeeTokens().length, 3);
    assertEq(s_feeQuoter.getFeeTokens()[2], feeTokens[0].token);

    // add same feeToken is no-op
    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokens);
    assertEq(s_feeQuoter.getFeeTokens().length, 3);
    assertEq(s_feeQuoter.getFeeTokens()[2], feeTokens[0].token);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenRemoved(feeTokenAddresses[0]);

    s_feeQuoter.applyFeeTokensUpdates(feeTokenAddresses, new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](0));
    assertEq(s_feeQuoter.getFeeTokens().length, 2);

    // removing already removed feeToken is no-op and does not emit an event
    vm.recordLogs();

    s_feeQuoter.applyFeeTokensUpdates(feeTokenAddresses, new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](0));
    assertEq(s_feeQuoter.getFeeTokens().length, 2);

    vm.assertEq(vm.getRecordedLogs().length, 0);

    // Removing and adding the same fee token is allowed and emits both events
    // Add it first
    s_feeQuoter.applyFeeTokensUpdates(new address[](0), feeTokens);

    vm.expectEmit();
    emit FeeQuoter.FeeTokenRemoved(feeTokenAddresses[0]);
    vm.expectEmit();
    emit FeeQuoter.FeeTokenAdded(feeTokens[0].token, feeTokens[0].premiumMultiplierWeiPerEth);

    s_feeQuoter.applyFeeTokensUpdates(feeTokenAddresses, feeTokens);
  }

  // Reverts

  function test_applyFeeTokensUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_feeQuoter.applyFeeTokensUpdates(new address[](0), new FeeQuoter.PremiumMultiplierWeiPerEthArgs[](0));
  }
}
