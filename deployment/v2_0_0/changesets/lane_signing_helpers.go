package changesets

import (
	"fmt"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/operations/fetch_signing_keys"
)

// deriveFamiliesFromSelectors returns the unique chain families for the given selectors.
func deriveFamiliesFromSelectors(selectors []uint64) []string {
	seen := make(map[string]struct{})
	for _, sel := range selectors {
		if family, err := chainsel.GetSelectorFamily(sel); err == nil {
			seen[family] = struct{}{}
		}
	}
	families := make([]string, 0, len(seen))
	for f := range seen {
		families = append(families, f)
	}
	return families
}

// fetchSigningKeysForNOPsByFamilies fetches JD signing keys for NOPs that are
// missing a configured signer address for any of the given families.
func fetchSigningKeysForNOPsByFamilies(e deployment.Environment, nops []offchain.NOPConfig, families []string) (fetch_signing_keys.SigningKeysByNOP, error) {
	return fetchSigningKeysForNOPsFiltered(e, nops, func(nop offchain.NOPConfig) bool {
		for _, family := range families {
			if nop.SignerAddressByFamily == nil || nop.SignerAddressByFamily[family] == "" {
				return true
			}
		}
		return false
	})
}

func fetchSigningKeysForNOPsFiltered(
	e deployment.Environment,
	nops []offchain.NOPConfig,
	include func(offchain.NOPConfig) bool,
) (fetch_signing_keys.SigningKeysByNOP, error) {
	if e.Offchain == nil {
		return nil, nil
	}

	aliases := make([]string, 0, len(nops))
	for _, nop := range nops {
		if include(nop) {
			aliases = append(aliases, nop.Alias)
		}
	}

	if len(aliases) == 0 {
		return nil, nil
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		fetch_signing_keys.FetchNOPSigningKeys,
		fetch_signing_keys.FetchSigningKeysDeps{
			JDClient: e.Offchain,
			Logger:   e.Logger,
			NodeIDs:  e.NodeIDs,
		},
		fetch_signing_keys.FetchSigningKeysInput{
			NOPAliases: aliases,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch signing keys from JD for NOPs %v: %w", aliases, err)
	}

	return report.Output.SigningKeysByNOP, nil
}

// getSignatureConfigForLane builds the committee verifier signature quorum
// config (signers + threshold) for a local→remote lane from topology + JD keys.
func getSignatureConfigForLane(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	committeeQualifier string,
	localSelector uint64,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (*adapters.CommitteeVerifierSignatureQuorumConfig, error) {
	committee, ok := topology.NOPTopology.Committees[committeeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found", committeeQualifier)
	}

	chainCfg, ok := committee.ChainConfigs[strconv.FormatUint(remoteSelector, 10)]
	if !ok {
		return nil, fmt.Errorf("chain selector %d not found in committee %q", remoteSelector, committeeQualifier)
	}

	localFamily, err := chainsel.GetSelectorFamily(localSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get selector family for selector %d: %w", localSelector, err)
	}

	signers := make([]string, 0, len(chainCfg.NOPAliases))
	for _, alias := range chainCfg.NOPAliases {
		signer, err := signerAddressForNOPAlias(e, topology, alias, localFamily, committeeQualifier, remoteSelector, signingKeysByNOP)
		if err != nil {
			return nil, err
		}
		signers = append(signers, signer)
	}

	return &adapters.CommitteeVerifierSignatureQuorumConfig{
		Threshold: chainCfg.Threshold,
		Signers:   signers,
	}, nil
}

func signerAddressForNOPAlias(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	alias string,
	localFamily string,
	committeeQualifier string,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (string, error) {
	nop, ok := topology.NOPTopology.GetNOP(alias)
	if !ok {
		return "", fmt.Errorf(
			"NOP alias %q not found for committee %q chain %d",
			alias, committeeQualifier, remoteSelector,
		)
	}

	if nop.SignerAddressByFamily != nil {
		if addr := nop.SignerAddressByFamily[localFamily]; addr != "" {
			return addr, nil
		}
	}

	if signer, ok := signerFromJDIfMissing(
		nop.SignerAddressByFamily,
		alias,
		localFamily,
		signingKeysByNOP,
	); ok {
		e.Logger.Debugw("Using signing address from JD",
			"nopAlias", alias,
			"chainFamily", localFamily,
			"signerAddress", signer,
		)
		return signer, nil
	}

	return "", fmt.Errorf(
		"NOP %q missing signer_address for family %s on committee %q chain %d",
		alias, localFamily, committeeQualifier, remoteSelector,
	)
}

func signerFromJDIfMissing(
	signerAddresses map[string]string,
	nopAlias string,
	family string,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (string, bool) {
	if signerAddresses != nil && signerAddresses[family] != "" {
		return "", false
	}

	if signingKeysByNOP == nil {
		return "", false
	}

	if signer := signingKeysByNOP[nopAlias][family]; signer != "" {
		return signer, true
	}

	return "", false
}
