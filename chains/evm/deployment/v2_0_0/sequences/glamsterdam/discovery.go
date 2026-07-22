// Package glamsterdam contains the sequences backing the v2.0 "update gas config for
// Glamsterdam" changeset.
package glamsterdam

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
)

// DiscoverLanesToTargetInput is the input to DiscoverLanesToTarget.
type DiscoverLanesToTargetInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	// FeeQuoterAddressByChain maps each candidate source chain selector to its FeeQuoter address
	// on that chain. Chains whose FeeQuoter address couldn't be resolved from the datastore
	// should be omitted here and reported separately by the caller.
	FeeQuoterAddressByChain map[uint64]common.Address
}

// DiscoverLanesToTargetOutput is the output of DiscoverLanesToTarget.
type DiscoverLanesToTargetOutput struct {
	// LanesToUpdate are the candidate chain selectors with a confirmed lane pointed at the
	// target chain, sorted ascending for deterministic output.
	LanesToUpdate []uint64
	// NoLane are candidate chain selectors that were scanned but have no lane pointed at the
	// target chain, sorted ascending.
	NoLane []uint64
}

// DiscoverLanesToTarget checks, for every candidate chain in the input, whether its FeeQuoter
// has an enabled DestChainConfig pointed at the target chain — i.e. whether that chain has a
// lane sending into the Glamsterdam target chain and therefore needs its source-side gas config
// updated. This is a read-only scan; it makes one RPC call per candidate chain.
var DiscoverLanesToTarget = cldf_ops.NewSequence(
	"DiscoverLanesToTargetV2",
	semver.MustParse("2.0.0"),
	"Discovers which v2.0 chains have a lane pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input DiscoverLanesToTargetInput) (DiscoverLanesToTargetOutput, error) {
		var output DiscoverLanesToTargetOutput

		candidates := make([]uint64, 0, len(input.FeeQuoterAddressByChain))
		for sel := range input.FeeQuoterAddressByChain {
			candidates = append(candidates, sel)
		}
		sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })

		for _, src := range candidates {
			chain, ok := chains.EVMChains()[src]
			if !ok {
				return DiscoverLanesToTargetOutput{}, fmt.Errorf("chain with selector %d not found", src)
			}

			report, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: src,
				Address:       input.FeeQuoterAddressByChain[src],
				Args:          input.TargetChainSelector,
			})
			if err != nil {
				return DiscoverLanesToTargetOutput{}, fmt.Errorf(
					"failed to read FeeQuoter dest chain config for src %d, dst %d: %w", src, input.TargetChainSelector, err,
				)
			}

			if report.Output.IsEnabled {
				output.LanesToUpdate = append(output.LanesToUpdate, src)
			} else {
				output.NoLane = append(output.NoLane, src)
			}
		}

		return output, nil
	},
)
