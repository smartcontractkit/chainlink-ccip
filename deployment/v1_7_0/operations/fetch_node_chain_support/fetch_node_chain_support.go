package fetch_node_chain_support

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type FetchNodeChainSupportInput struct {
	NOPAliases []string
}

type FetchNodeChainSupportOutput struct {
	SupportedChains shared.ChainSupportByNOP
}

type FetchNodeChainSupportDeps struct {
	JDClient shared.JDClient
	Logger   logger.Logger
	NodeIDs  []string
}

var FetchNodeChainSupport = operations.NewOperation(
	"fetch-node-chain-support",
	semver.MustParse("1.0.0"),
	"Fetches supported chain selectors for specified NOPs from the job distributor",
	func(b operations.Bundle, deps FetchNodeChainSupportDeps, input FetchNodeChainSupportInput) (FetchNodeChainSupportOutput, error) {
		ctx := b.GetContext()
		lggr := deps.Logger

		output := FetchNodeChainSupportOutput{
			SupportedChains: make(shared.ChainSupportByNOP),
		}

		if len(input.NOPAliases) == 0 {
			return output, nil
		}

		lookup, err := shared.FetchNodeLookup(ctx, deps.JDClient, deps.NodeIDs)
		if err != nil {
			return output, err
		}

		nodeIDs := make([]string, 0, len(input.NOPAliases))
		nodeIDToAlias := make(map[string]string)
		for _, nopAlias := range input.NOPAliases {
			node, ok := lookup.FindByName(nopAlias)
			if !ok {
				lggr.Warnw("Node not found for NOP alias",
					"nopAlias", nopAlias)
				continue
			}
			nodeIDs = append(nodeIDs, node.Id)
			nodeIDToAlias[node.Id] = nopAlias
		}

		if len(nodeIDs) == 0 {
			return output, nil
		}

		chainConfigsResp, err := deps.JDClient.ListNodeChainConfigs(ctx, &nodev1.ListNodeChainConfigsRequest{
			Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
				NodeIds: nodeIDs,
			},
		})
		if err != nil {
			return output, fmt.Errorf("failed to list chain configs: %w", err)
		}

		for _, chainConfig := range chainConfigsResp.ChainConfigs {
			nopAlias, ok := nodeIDToAlias[chainConfig.NodeId]
			if !ok {
				continue
			}

			chainFamily, ok := shared.ProtoChainTypeToFamily[chainConfig.Chain.Type]
			if !ok {
				lggr.Debugw("Skipping unsupported chain type",
					"chainType", chainConfig.Chain.Type.String())
				continue
			}

			chainDetails, err := chainsel.GetChainDetailsByChainIDAndFamily(chainConfig.Chain.Id, chainFamily)
			if err != nil {
				lggr.Warnw("Failed to get chain details from chain ID",
					"chainId", chainConfig.Chain.Id,
					"chainFamily", chainFamily,
					"error", err)
				continue
			}
			chainSelector := chainDetails.ChainSelector

			if output.SupportedChains[nopAlias] == nil {
				output.SupportedChains[nopAlias] = make([]uint64, 0)
			}
			output.SupportedChains[nopAlias] = append(output.SupportedChains[nopAlias], chainSelector)

			lggr.Debugw("Found supported chain",
				"nopAlias", nopAlias,
				"nodeId", chainConfig.NodeId,
				"chainId", chainConfig.Chain.Id,
				"chainFamily", chainFamily,
				"chainSelector", chainSelector)
		}

		return output, nil
	},
)
