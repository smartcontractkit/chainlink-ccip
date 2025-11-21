// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_setDomains is CCTPVerifierSetup {
  function test_setDomains() public {
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });

    vm.expectEmit();
    emit CCTPVerifier.DomainsSet(domainUpdates);
    s_cctpVerifier.setDomains(domainUpdates);

    // Check the domains.
    CCTPVerifier.Domain memory domain = s_cctpVerifier.getDomain(DEST_CHAIN_SELECTOR);
    assertEq(domain.allowedCallerOnDest, ALLOWED_CALLER_ON_DEST);
    assertEq(domain.allowedCallerOnSource, ALLOWED_CALLER_ON_SOURCE);
    assertEq(domain.mintRecipientOnDest, bytes32(0));
    assertEq(domain.domainIdentifier, REMOTE_DOMAIN_IDENTIFIER);
    assertEq(domain.enabled, true);

    // Disable the domain.
    domainUpdates[0].enabled = false;
    vm.expectEmit();
    emit CCTPVerifier.DomainsSet(domainUpdates);
    s_cctpVerifier.setDomains(domainUpdates);

    // Check the domains again.
    domain = s_cctpVerifier.getDomain(DEST_CHAIN_SELECTOR);
    assertEq(domain.enabled, false);
  }

  function test_setDomains_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.setDomains(new CCTPVerifier.DomainUpdate[](0));
  }

  function test_setDomains_RevertWhen_InvalidDomain_AllowedCallerOnDestIsZero() public {
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: bytes32(0),
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidDomainUpdate.selector, domainUpdates[0]));
    s_cctpVerifier.setDomains(domainUpdates);
  }

  function test_setDomains_RevertWhen_InvalidDomain_AllowedCallerOnSourceIsZero() public {
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: bytes32(0),
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: DEST_CHAIN_SELECTOR,
      enabled: true
    });

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidDomainUpdate.selector, domainUpdates[0]));
    s_cctpVerifier.setDomains(domainUpdates);
  }

  function test_setDomains_RevertWhen_InvalidDomain_DestChainSelectorIsZero() public {
    CCTPVerifier.DomainUpdate[] memory domainUpdates = new CCTPVerifier.DomainUpdate[](1);
    domainUpdates[0] = CCTPVerifier.DomainUpdate({
      allowedCallerOnDest: ALLOWED_CALLER_ON_DEST,
      allowedCallerOnSource: ALLOWED_CALLER_ON_SOURCE,
      mintRecipientOnDest: bytes32(0),
      domainIdentifier: REMOTE_DOMAIN_IDENTIFIER,
      chainSelector: 0,
      enabled: true
    });

    vm.expectRevert(abi.encodeWithSelector(CCTPVerifier.InvalidDomainUpdate.selector, domainUpdates[0]));
    s_cctpVerifier.setDomains(domainUpdates);
  }
}
