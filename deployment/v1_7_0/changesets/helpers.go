package changesets

import (
	"strconv"

	execcontract "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/fetch_signing_keys"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
)

func convertTopologyMonitoring(m *topology.MonitoringConfig) shared.MonitoringInput {
	if m == nil {
		return shared.MonitoringInput{}
	}
	return shared.MonitoringInput{
		Enabled: m.Enabled,
		Type:    m.Type,
		Beholder: shared.BeholderInput{
			InsecureConnection:       m.Beholder.InsecureConnection,
			CACertFile:               m.Beholder.CACertFile,
			OtelExporterGRPCEndpoint: m.Beholder.OtelExporterGRPCEndpoint,
			OtelExporterHTTPEndpoint: m.Beholder.OtelExporterHTTPEndpoint,
			LogStreamingEnabled:      m.Beholder.LogStreamingEnabled,
			MetricReaderInterval:     m.Beholder.MetricReaderInterval,
			TraceSampleRatio:         m.Beholder.TraceSampleRatio,
			TraceBatchTimeout:        m.Beholder.TraceBatchTimeout,
		},
	}
}

func buildNOPModes(nops []topology.NOPConfig) map[shared.NOPAlias]shared.NOPMode {
	nopModes := make(map[shared.NOPAlias]shared.NOPMode)
	for _, nop := range nops {
		mode := nop.GetMode()
		nopModes[shared.NOPAlias(nop.Alias)] = mode
	}
	return nopModes
}

func getAllNOPAliases(nops []topology.NOPConfig) []shared.NOPAlias {
	aliases := make([]shared.NOPAlias, len(nops))
	for i, nop := range nops {
		aliases[i] = shared.NOPAlias(nop.Alias)
	}
	return aliases
}

func getCommitteeChainSelectors(committee topology.CommitteeConfig) []uint64 {
	selectors := make([]uint64, 0, len(committee.ChainConfigs))
	for chainStr := range committee.ChainConfigs {
		if sel, err := strconv.ParseUint(chainStr, 10, 64); err == nil {
			selectors = append(selectors, sel)
		}
	}
	return selectors
}

func filterChains(input, allowed []uint64) []uint64 {
	allowedSet := make(map[uint64]bool, len(allowed))
	for _, c := range allowed {
		allowedSet[c] = true
	}

	filtered := make([]uint64, 0, len(input))
	for _, sel := range input {
		if allowedSet[sel] {
			filtered = append(filtered, sel)
		}
	}
	return filtered
}

func getExecutorDeployedChains(ds datastore.DataStore, qualifier string) []uint64 {
	if ds == nil {
		return nil
	}
	refs := ds.Addresses().Filter(
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(execcontract.ProxyType)),
	)
	seen := make(map[uint64]bool, len(refs))
	chains := make([]uint64, 0, len(refs))
	for _, ref := range refs {
		if !seen[ref.ChainSelector] {
			seen[ref.ChainSelector] = true
			chains = append(chains, ref.ChainSelector)
		}
	}
	return chains
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
