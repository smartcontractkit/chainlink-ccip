package fetch_signing_keys

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type SigningKeysByNOP map[string]map[string]string

type FetchSigningKeysInput struct {
	NOPAliases []string
}

type FetchSigningKeysOutput struct {
	SigningKeysByNOP SigningKeysByNOP
}

type FetchSigningKeysDeps struct {
	JDClient shared.JDClient
	Logger   logger.Logger
	NodeIDs  []string
}

var FetchNOPSigningKeys = operations.NewOperation(
	"fetch-nop-signing-keys",
	semver.MustParse("1.0.0"),
	"Fetches signing keys for all specified NOPs from the job distributor in a single batch",
	func(b operations.Bundle, deps FetchSigningKeysDeps, input FetchSigningKeysInput) (FetchSigningKeysOutput, error) {
		ctx := b.GetContext()
		lggr := deps.Logger

		output := FetchSigningKeysOutput{
			SigningKeysByNOP: make(SigningKeysByNOP),
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
			if chainConfig.Ocr2Config == nil || chainConfig.Ocr2Config.OcrKeyBundle == nil {
				continue
			}

			nopAlias, ok := nodeIDToAlias[chainConfig.NodeId]
			if !ok {
				continue
			}

			signerAddress := chainConfig.Ocr2Config.OcrKeyBundle.OnchainSigningAddress
			if signerAddress == "" {
				continue
			}

			chainFamily, ok := shared.ProtoChainTypeToFamily[chainConfig.Chain.Type]
			if !ok {
				lggr.Debugw("Skipping unsupported chain type",
					"chainType", chainConfig.Chain.Type.String())
				continue
			}

			if output.SigningKeysByNOP[nopAlias] == nil {
				output.SigningKeysByNOP[nopAlias] = make(map[string]string)
			}
			output.SigningKeysByNOP[nopAlias][chainFamily] = fmt.Sprintf("0x%s", signerAddress)

			lggr.Debugw("Found signing address",
				"nopAlias", nopAlias,
				"nodeId", chainConfig.NodeId,
				"chainFamily", chainFamily,
				"signerAddress", signerAddress)
		}

		return output, nil
	},
)
