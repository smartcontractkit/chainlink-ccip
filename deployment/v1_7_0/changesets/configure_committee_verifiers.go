package changesets

import (
	"fmt"
	"strconv"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/operations/fetch_signing_keys"
)

// TopologyCommitteePopulator implements lanes.CommitteeConfigPopulator by
// populating committee verifier contracts from a registry and signing keys
// from NOP topology + JD.
type TopologyCommitteePopulator struct {
	contractRegistry *adapters.CommitteeVerifierContractRegistry
	topology         *offchain.EnvironmentTopology

	signingKeysOnce sync.Once
	signingKeys     fetch_signing_keys.SigningKeysByNOP
	signingKeysErr  error
}

// NewTopologyCommitteePopulator creates a populator that encapsulates the
// topology-aware committee verifier population. It is intended to be
// created per-invocation with the relevant topology config.
func NewTopologyCommitteePopulator(
	contractRegistry *adapters.CommitteeVerifierContractRegistry,
	topology *offchain.EnvironmentTopology,
) *TopologyCommitteePopulator {
	return &TopologyCommitteePopulator{
		contractRegistry: contractRegistry,
		topology:         topology,
	}
}

// Deprecated: Use ConfigureChainsForLanesFromTopology this is the canonical way to configure lanes for 2.0
func (r *TopologyCommitteePopulator) PopulateCommitteeConfig(
	e deployment.Environment,
	chainSelector uint64,
	inputs []lanes.CommitteeVerifierInput,
) ([]lanes.CommitteeVerifierConfig[datastore.AddressRef], error) {
	if r.topology == nil || r.topology.NOPTopology == nil {
		return nil, fmt.Errorf("TopologyCommitteePopulator: topology/NOPTopology must not be nil")
	}
	if r.contractRegistry == nil {
		return nil, fmt.Errorf("TopologyCommitteePopulator: contractRegistry must not be nil")
	}
	r.signingKeysOnce.Do(func() {
		committeeNOPs := committeeNOPAliases(r.topology.NOPTopology)
		r.signingKeys, r.signingKeysErr = fetchSigningKeysForNOPsFiltered(e, r.topology.NOPTopology.NOPs, func(nop offchain.NOPConfig) bool {
			if _, ok := committeeNOPs[nop.Alias]; !ok {
				return false
			}
			return len(nop.SignerAddressByFamily) == 0
		})
	})
	if r.signingKeysErr != nil {
		return nil, fmt.Errorf("failed to fetch signing keys: %w", r.signingKeysErr)
	}

	result := make([]lanes.CommitteeVerifierConfig[datastore.AddressRef], 0, len(inputs))
	for _, input := range inputs {
		remoteChains := make(map[uint64]lanes.CommitteeVerifierRemoteChainConfig, len(input.RemoteChains))
		for remoteChainSelector, rc := range input.RemoteChains {
			signatureConfig, err := getSignatureConfigForLane(
				e, r.topology, input.CommitteeQualifier,
				chainSelector, remoteChainSelector, r.signingKeys,
			)
			if err != nil {
				return nil, fmt.Errorf(
					"failed to get signature config for lane local chain %d -> remote chain %d: %w",
					chainSelector, remoteChainSelector, err,
				)
			}
			remoteChains[remoteChainSelector] = lanes.CommitteeVerifierRemoteChainConfig{
				AllowlistEnabled:          rc.AllowlistEnabled,
				AddedAllowlistedSenders:   rc.AddedAllowlistedSenders,
				RemovedAllowlistedSenders: rc.RemovedAllowlistedSenders,
				FeeUSDCents:               rc.FeeUSDCents,
				GasForVerification:        rc.GasForVerification,
				PayloadSizeBytes:          rc.PayloadSizeBytes,
				SignatureConfig: lanes.CommitteeVerifierSignatureQuorumConfig{
					Signers:   signatureConfig.Signers,
					Threshold: signatureConfig.Threshold,
				},
			}
		}

		adapter, err := r.contractRegistry.GetByChain(chainSelector)
		if err != nil {
			return nil, fmt.Errorf("no committee verifier contract adapter for chain %d: %w", chainSelector, err)
		}

		contracts, err := adapter.ResolveCommitteeVerifierContracts(e.DataStore, chainSelector, input.CommitteeQualifier)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to resolve committee verifier contracts for chain %d qualifier %q: %w",
				chainSelector, input.CommitteeQualifier, err,
			)
		}

		result = append(result, lanes.CommitteeVerifierConfig[datastore.AddressRef]{
			CommitteeVerifier: contracts,
			RemoteChains:      remoteChains,
		})
	}

	return result, nil
}

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

func committeeNOPAliases(nopTopology *offchain.NOPTopology) map[string]struct{} {
	aliases := make(map[string]struct{})
	for _, committee := range nopTopology.Committees {
		for _, chainCfg := range committee.ChainConfigs {
			for _, alias := range chainCfg.NOPAliases {
				aliases[alias] = struct{}{}
			}
		}
	}
	return aliases
}
