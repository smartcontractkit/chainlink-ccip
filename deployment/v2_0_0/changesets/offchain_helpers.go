package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/operations/fetch_node_chain_support"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/shared"
)

func fetchNodeChainSupport(e deployment.Environment, nopAliases []string) (shared.ChainSupportByNOP, error) {
	if len(nopAliases) == 0 {
		return nil, nil
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		fetch_node_chain_support.FetchNodeChainSupport,
		fetch_node_chain_support.FetchNodeChainSupportDeps{
			JDClient: e.Offchain,
			Logger:   e.Logger,
			NodeIDs:  e.NodeIDs,
		},
		fetch_node_chain_support.FetchNodeChainSupportInput{
			NOPAliases: nopAliases,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch node chain support from JD: %w", err)
	}

	return report.Output.SupportedChains, nil
}

func convertTopologyMonitoring(m *offchain.MonitoringConfig) shared.MonitoringInput {
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

func buildNOPModes(nops []offchain.NOPConfig) map[shared.NOPAlias]shared.NOPMode {
	nopModes := make(map[shared.NOPAlias]shared.NOPMode)
	for _, nop := range nops {
		mode := nop.GetMode()
		nopModes[shared.NOPAlias(nop.Alias)] = mode
	}
	return nopModes
}

func filterCLModeNOPs(aliases []shared.NOPAlias, nops []offchain.NOPConfig) []shared.NOPAlias {
	modeByAlias := buildNOPModes(nops)
	filtered := make([]shared.NOPAlias, 0, len(aliases))
	for _, alias := range aliases {
		if mode, ok := modeByAlias[alias]; ok && mode == shared.NOPModeCL {
			filtered = append(filtered, alias)
		}
	}
	return filtered
}

func getAllNOPAliases(nops []offchain.NOPConfig) []shared.NOPAlias {
	aliases := make([]shared.NOPAlias, len(nops))
	for i, nop := range nops {
		aliases[i] = shared.NOPAlias(nop.Alias)
	}
	return aliases
}
